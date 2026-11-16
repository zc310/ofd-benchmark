"""Regression tests using small synthetic PDFs, without changing result YAML."""

import contextlib
import io
from pathlib import Path
import sys
import tempfile
import unittest
from unittest.mock import patch

import fitz
import numpy as np

import compare_pdf
import compare_pdf2ofd


def make_pdf(path, pages=2):
    with fitz.open() as doc:
        for _ in range(pages):
            page = doc.new_page(width=72, height=72)
            page.insert_text((10, 20), 'Test text')
        doc.save(path)


class PixelComparisonTests(unittest.TestCase):
    def test_matching_and_mismatched_page_counts(self):
        with tempfile.TemporaryDirectory() as tmp:
            ref = Path(tmp) / 'ref.pdf'
            make_pdf(ref)
            for pages in (1, 2, 3):
                src = Path(tmp) / f'output-{pages}.pdf'
                make_pdf(src, pages)
                for compare in (lambda a, b, dpi: compare_pdf.compare(a, b, dpi)[0],
                                compare_pdf2ofd.pixel_similarity):
                    with self.subTest(pages=pages, compare=compare):
                        if pages == 2:
                            self.assertAlmostEqual(compare(ref, src, 72), 100.0)
                        else:
                            with self.assertRaisesRegex(ValueError, 'page count mismatch'):
                                compare(ref, src, 72)

    def test_excluded_pages_do_not_bypass_page_count_check(self):
        with tempfile.TemporaryDirectory() as tmp:
            ref, src = Path(tmp) / 'ref.pdf', Path(tmp) / 'output.pdf'
            make_pdf(ref, 2)
            make_pdf(src, 1)
            with patch.dict(compare_pdf.EXCLUDE_PAGES, {'ref': {2}}):
                with self.assertRaisesRegex(ValueError, 'page count mismatch'):
                    compare_pdf.compare(ref, src, 72)
            make_pdf(src, 2)
            with patch.dict(compare_pdf.EXCLUDE_PAGES, {'ref': {1, 2}}):
                with self.assertRaisesRegex(ValueError, 'no comparable pages'):
                    compare_pdf.compare(ref, src, 72)

    def test_blank_output_does_not_receive_high_pixel_similarity(self):
        with tempfile.TemporaryDirectory() as tmp:
            ref, src = Path(tmp) / 'ref.pdf', Path(tmp) / 'blank.pdf'
            make_pdf(ref, 1)
            with fitz.open() as doc:
                doc.new_page(width=72, height=72)
                doc.save(src)
            with self.assertRaisesRegex(ValueError, 'nearly blank'):
                compare_pdf.compare(ref, src, 72)
            with fitz.open(ref) as ref_doc, fitz.open(src) as src_doc:
                ref_img = compare_pdf.render(ref_doc, 0, 72)
                src_img = compare_pdf.render(src_doc, 0, 72)
                self.assertLess(compare_pdf.content_similarity(ref_img, src_img), 90.0)

    def test_content_similarity_matches_identical_pages(self):
        with tempfile.TemporaryDirectory() as tmp:
            ref = Path(tmp) / 'ref.pdf'
            make_pdf(ref, 1)
            with fitz.open(ref) as doc:
                image = compare_pdf.render(doc, 0, 72)
            self.assertAlmostEqual(compare_pdf.content_similarity(image, image), 100.0)

    def test_cropped_images_are_rejected(self):
        ref = np.full((100, 100, 3), 255, dtype=np.uint8)
        ref[20:40, 80:95] = 0
        for score in (compare_pdf.content_similarity, compare_pdf.similarity):
            with self.subTest(score=score), self.assertRaisesRegex(ValueError, 'size mismatch'):
                score(ref, ref[:, :60])

    def test_single_pixel_rounding_is_padded_not_cropped(self):
        ref = np.full((10, 10, 3), 255, dtype=np.uint8)
        src = np.full((10, 11, 3), 255, dtype=np.uint8)
        src[:, -1] = 0
        self.assertEqual(compare_pdf.content_similarity(ref, src), 0.0)
        self.assertLess(compare_pdf.similarity(ref, src)[1], 1.0)

    def test_gray_background_cannot_improve_missing_content(self):
        ref = np.full((100, 100, 3), 255, dtype=np.uint8)
        ref[20:40, 20:40] = 0
        src = np.full_like(ref, 255)
        src[20:40, 20:30] = 0
        plain = compare_pdf.content_similarity(ref, src)
        src[src == 255] = 249
        self.assertLessEqual(compare_pdf.content_similarity(ref, src), plain)

    def test_tolerant_f1_rewards_shift_but_not_missing_line(self):
        ref = np.full((100, 100, 3), 255, dtype=np.uint8)
        ref[20:80, 50] = 0
        src = np.full_like(ref, 255)
        src[20:80, 51] = 0
        self.assertEqual(compare_pdf.foreground_f1(ref, src, 96), 100.0)
        src[:] = 255
        self.assertEqual(compare_pdf.foreground_f1(ref, src, 96), 0.0)
        self.assertEqual(compare_pdf.content_similarity(ref, src), 0.0)

    def test_extra_ink_is_penalized_by_color_and_foreground_scores(self):
        ref = np.full((100, 100, 3), 255, dtype=np.uint8)
        ref[20:40, 20:40] = 0
        src = ref.copy()
        src[60:80, 60:80] = 0
        self.assertLess(compare_pdf.content_similarity(ref, src), 100.0)
        self.assertLess(compare_pdf.foreground_f1(ref, src, 96), 100.0)

    def test_all_blank_document_has_no_content_scores(self):
        with tempfile.TemporaryDirectory() as tmp:
            blank = Path(tmp) / 'blank.pdf'
            with fitz.open() as doc:
                doc.new_page(width=72, height=72)
                doc.save(blank)
            result = compare_pdf.compare(blank, blank, 96)
            self.assertEqual(result[0], 100.0)
            self.assertIsNone(result[3])
            self.assertIsNone(result[4])

    def test_blank_pages_do_not_inflate_document_content_scores(self):
        with tempfile.TemporaryDirectory() as tmp:
            ref, src = Path(tmp) / 'ref.pdf', Path(tmp) / 'output.pdf'
            for path, width in ((ref, 20), (src, 10)):
                with fitz.open() as doc:
                    page = doc.new_page(width=100, height=100)
                    page.draw_rect(fitz.Rect(20, 20, 20 + width, 40), fill=(0, 0, 0), width=0)
                    doc.save(path)
            before = compare_pdf.compare(ref, src, 96)
            for path in (ref, src):
                with fitz.open(path) as doc:
                    for _ in range(9):
                        doc.new_page(width=100, height=100)
                    doc.saveIncr()
            after = compare_pdf.compare(ref, src, 96)
            self.assertEqual(before[3:], after[3:])
            blank = np.full((10, 10, 3), 255, dtype=np.uint8)
            self.assertIsNone(compare_pdf.content_similarity(blank, blank))
            self.assertIsNone(compare_pdf.foreground_f1(blank, blank, 96))

    def test_page_physical_size_mismatch_is_rejected(self):
        with tempfile.TemporaryDirectory() as tmp:
            ref, src = Path(tmp) / 'ref.pdf', Path(tmp) / 'output.pdf'
            make_pdf(ref, 1)
            with fitz.open() as doc:
                doc.new_page(width=40, height=72).insert_text((10, 20), 'Test')
                doc.save(src)
            with self.assertRaisesRegex(ValueError, 'page size mismatch'):
                compare_pdf.compare(ref, src, 96)

    def test_repeated_document_samples_only_first_three_pages(self):
        with tempfile.TemporaryDirectory() as tmp:
            ref, src = Path(tmp) / '1000-pages.pdf', Path(tmp) / 'output.pdf'
            make_pdf(ref, 4)
            make_pdf(src, 4)
            with patch.object(compare_pdf, 'render', wraps=compare_pdf.render) as render:
                result = compare_pdf.compare(ref, src, 96)
            self.assertEqual([call.args[1] for call in render.call_args_list], [0, 0, 1, 1, 2, 2])
            self.assertEqual(result[0], 100.0)
            self.assertEqual(result[3:], (100.0, 100.0))

            # Sampling must not hide missing pages after the sampled range.
            make_pdf(src, 5)
            with self.assertRaisesRegex(ValueError, 'page count mismatch'):
                compare_pdf.compare(ref, src, 96)

    def test_render_warnings_fail_instead_of_scoring_partial_output(self):
        with fitz.open() as doc:
            page = doc.new_page()
            page.insert_text((20, 20), 'Test')
            doc.update_stream(page.get_contents()[0], b'/Pattern cs /Missing scn 0 0 30 30 re f')
            with self.assertRaisesRegex(ValueError, 'incomplete render'):
                compare_pdf.render(doc, 0, 96)


class ComparisonScriptTests(unittest.TestCase):
    def run_script(self, module, base):
        with patch.object(module, '__file__', str(base / 'tools' / 'script.py')), \
             patch.object(sys, 'argv', ['script.py', '72', 'hello']), \
             patch.object(module, 'write_updates') as writer, \
             contextlib.redirect_stdout(io.StringIO()):
            module.main()
        writer.assert_called_once()
        return {(cell['converter'], cell['table']): cell['value']
                for cell in writer.call_args.args[1]}

    def test_ofd_to_pdf_continues_after_invalid_pdf(self):
        with tempfile.TemporaryDirectory() as tmp:
            base = Path(tmp)
            ref = base / 'testdata' / 'suwell' / 'hello.pdf'
            ref.parent.mkdir(parents=True)
            make_pdf(ref)
            output = base / 'bin' / 'output'
            output.mkdir(parents=True)
            (output / 'hello_go-ofdgo.pdf').write_bytes(b'invalid PDF')
            make_pdf(output / 'hello_go-zc310.pdf')
            make_pdf(output / 'hello_node-ofd2pdf.pdf', 1)
            updates = self.run_script(compare_pdf, base)
            table = 'ofd-pixel-similarity'
            self.assertEqual(updates['go-ofdgo', table], 'FAILED')
            self.assertEqual(updates['go-zc310', table], '100.00%')
            self.assertEqual(updates['@miconvert/ofd-to-pdf', table], 'FAILED')
            for table in ('ofd-content-similarity', 'ofd-foreground-f1'):
                self.assertEqual(updates['go-ofdgo', table], 'FAILED')
                self.assertEqual(updates['go-zc310', table], '100.00%')

    def test_round_trip_pixel_failure_does_not_block_text_or_next_output(self):
        with tempfile.TemporaryDirectory() as tmp:
            base = Path(tmp)
            ref = base / 'testdata' / 'suwell' / 'hello.pdf'
            ref.parent.mkdir(parents=True)
            make_pdf(ref)
            output = base / 'bin' / 'output_ofd'
            output.mkdir(parents=True)
            for converter in ('pdf2ofd-ofdrw', 'pdf2ofd-zc310', 'node-jsofd', 'python-pdf2ofd'):
                stem = output / f'hello_{converter}'
                stem.with_suffix('.ofd').write_bytes(b'placeholder OFD')
                stem.with_suffix('.txt').write_text('Test textTest text', encoding='utf-8')
                if converter == 'pdf2ofd-ofdrw':
                    stem.with_suffix('.pdf').write_bytes(b'invalid PDF')
                elif converter != 'python-pdf2ofd':
                    make_pdf(stem.with_suffix('.pdf'), 1 if converter == 'node-jsofd' else 2)
            updates = self.run_script(compare_pdf2ofd, base)
            table = 'pdf-pixel-similarity'
            self.assertEqual(updates['java-ofdrw', table], 'FAILED')
            self.assertEqual(updates['go-zc310', table], '100.00%')
            self.assertEqual(updates['node-jsofd', table], 'FAILED')
            self.assertEqual(updates['python-pdf2ofd', table], 'FAILED')
            for converter in ('java-ofdrw', 'go-zc310', 'node-jsofd', 'python-pdf2ofd'):
                self.assertEqual(updates[converter, 'pdf-text-extraction'], '18')


if __name__ == '__main__':
    unittest.main()
