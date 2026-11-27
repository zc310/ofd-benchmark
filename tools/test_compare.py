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
import compare_text
from benchmark_results import format_percentage


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
        for score in (compare_pdf.content_similarity, compare_pdf.similarity,
                      compare_pdf.ssim_similarity):
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
            self.assertEqual(result[5], 100.0)

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
            self.assertEqual(before[3:5], after[3:5])
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
            self.assertEqual(result[3:], (100.0, 100.0, 100.0))

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


class SSIMTests(unittest.TestCase):
    def test_identical_constant_and_textured_images(self):
        rng = np.random.default_rng(42)
        for image in (np.zeros((10, 10, 3), dtype=np.uint8),
                      np.full((10, 10, 3), 255, dtype=np.uint8),
                      rng.integers(0, 256, (150, 17, 3), dtype=np.uint8)):
            with self.subTest(shape=image.shape):
                self.assertAlmostEqual(compare_pdf.ssim_similarity(image, image), 100.0)

    def test_matches_direct_local_ssim_across_strip_boundary(self):
        # Independent, non-integral-image reference: all 7x7 windows at once.
        rng = np.random.default_rng(19)
        ref = rng.integers(0, 256, (150, 17, 3), dtype=np.uint8)
        src = np.clip(ref.astype(float) * 0.8 + 20, 0, 255).astype(np.uint8)
        x = np.lib.stride_tricks.sliding_window_view(ref.astype(float), (7, 7), axis=(0, 1))
        y = np.lib.stride_tricks.sliding_window_view(src.astype(float), (7, 7), axis=(0, 1))
        ux, uy = x.mean(axis=(-2, -1)), y.mean(axis=(-2, -1))
        vx, vy = x.var(axis=(-2, -1), ddof=1), y.var(axis=(-2, -1), ddof=1)
        cov = ((x - ux[..., None, None]) *
               (y - uy[..., None, None])).sum(axis=(-2, -1)) / 48
        c1, c2 = (0.01 * 255) ** 2, (0.03 * 255) ** 2
        expected = ((2 * ux * uy + c1) * (2 * cov + c2) /
                    ((ux ** 2 + uy ** 2 + c1) * (vx + vy + c2))).mean() * 100
        self.assertAlmostEqual(compare_pdf.ssim_similarity(ref, src), expected, places=9)

    def test_missing_structure_extra_ink_and_color_are_penalized(self):
        ref = np.full((50, 50, 3), 255, dtype=np.uint8)
        ref[10:40:4, 10:40] = 0
        missing, extra, color = ref.copy(), ref.copy(), ref.copy()
        missing[10:20] = 255
        extra[45:48, 10:40] = 0
        color[ref[:, :, 0] == 0] = (255, 0, 0)
        for src in (missing, extra, color):
            with self.subTest():
                self.assertLess(compare_pdf.ssim_similarity(ref, src), 100.0)
                self.assertAlmostEqual(compare_pdf.ssim_similarity(ref, src),
                                       compare_pdf.ssim_similarity(src, ref))

    def test_border_pixels_are_not_discarded(self):
        ref = np.full((10, 10, 3), 255, dtype=np.uint8)
        src = ref.copy()
        src[0, 0] = 0
        self.assertLess(compare_pdf.ssim_similarity(ref, src), 100.0)

    def test_rounding_is_padded_and_small_images_are_rejected(self):
        ref = np.full((10, 10, 3), 255, dtype=np.uint8)
        src = np.pad(ref, ((0, 1), (0, 1), (0, 0)), constant_values=255)
        self.assertEqual(compare_pdf.ssim_similarity(ref, src), 100.0)
        src[-1, :] = 0
        self.assertLess(compare_pdf.ssim_similarity(ref, src), 100.0)
        with self.assertRaisesRegex(ValueError, 'at least 7x7'):
            compare_pdf.ssim_similarity(ref[:6], ref[:6])

    def test_negative_ssim_is_reported_as_zero_percent(self):
        grid = np.indices((20, 20)).sum(axis=0) % 2 * 255
        ref = np.repeat(grid[:, :, None], 3, axis=2).astype(np.uint8)
        self.assertEqual(compare_pdf.ssim_similarity(ref, 255 - ref), 0.0)


class ComparisonScriptTests(unittest.TestCase):
    def test_percentage_precision(self):
        for value, expected in (
                (None, 'FAILED'),
                (0, '0.0000%'),
                (100, '100.0000%'),
                (98.123456, '98.1235%'),
                (99.996, '99.9960%'),
                (99.99996, '100.0000%')):
            with self.subTest(value=value):
                self.assertEqual(format_percentage(value), expected)

    def run_script(self, module, base):
        with patch.dict('os.environ', {'BENCHMARK_OUTPUT_DIR': str(base)}), \
             patch.object(module, '__file__', str(base / 'tools' / 'script.py')), \
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
            output = base / 'output'
            output.mkdir(parents=True)
            (output / 'hello_go-ofdgo.pdf').write_bytes(b'invalid PDF')
            make_pdf(output / 'hello_go-zc310.pdf')
            make_pdf(output / 'hello_rust-sofd.pdf')
            make_pdf(output / 'hello_rust-easyofd.pdf')  # Ignore retired converter output.
            make_pdf(output / 'hello_node-ofd2pdf.pdf', 1)
            updates = self.run_script(compare_pdf, base)
            self.assertNotIn('rust-easyofd', {converter for converter, _ in updates})
            table = 'ofd-pixel-similarity'
            self.assertEqual(updates['go-xiaoqidun-ofdgo', table], 'FAILED')
            self.assertEqual(updates['go-zc310-ofd', table], '100.0000%')
            self.assertEqual(updates['rust-sOFD', table], '100.0000%')
            self.assertEqual(updates['node-ofd-to-pdf', table], 'FAILED')
            for table in ('ofd-content-similarity', 'ofd-foreground-f1', 'ofd-ssim'):
                self.assertEqual(updates['go-xiaoqidun-ofdgo', table], 'FAILED')
                self.assertEqual(updates['go-zc310-ofd', table], '100.0000%')
                self.assertEqual(updates['rust-sOFD', table], '100.0000%')

    def test_sofd_text_results_are_written(self):
        with tempfile.TemporaryDirectory() as tmp:
            base = Path(tmp)
            ref = base / 'testdata' / 'suwell' / 'hello.pdf'
            ref.parent.mkdir(parents=True)
            make_pdf(ref)
            output = base / 'output'
            output.mkdir(parents=True)
            make_pdf(output / 'hello_rust-sofd.pdf')
            make_pdf(output / 'hello_rust-easyofd.pdf')  # Ignore retired converter output.
            with patch.dict('os.environ', {'BENCHMARK_OUTPUT_DIR': str(base)}), \
                 patch.object(compare_text, '__file__', str(base / 'tools' / 'script.py')), \
                 patch.object(sys, 'argv', ['script.py', 'hello']), \
                 patch.object(compare_text, 'write_updates') as writer, \
                 contextlib.redirect_stdout(io.StringIO()):
                compare_text.main()
            cells = {cell['table']: cell for cell in writer.call_args.args[1]}
            self.assertNotIn('rust-easyofd',
                             {cell['converter'] for cell in writer.call_args.args[1]})
            self.assertEqual(cells['ofd-text-similarity']['converter'], 'rust-sOFD')
            self.assertEqual(cells['ofd-text-similarity']['value'], '100.0000%')
            self.assertEqual(cells['ofd-text-extraction']['value'], '19')

    def test_1000_pages_text_extraction_is_written_for_entire_document(self):
        with tempfile.TemporaryDirectory() as tmp:
            base = Path(tmp)
            ref = base / 'testdata' / 'suwell' / '1000-pages.pdf'
            ref.parent.mkdir(parents=True)
            make_pdf(ref, 5)
            output = base / 'output'
            output.mkdir(parents=True)
            make_pdf(output / '1000-pages_go-zc310.pdf', 5)
            with patch.dict('os.environ', {'BENCHMARK_OUTPUT_DIR': str(base)}), \
                 patch.object(compare_text, '__file__', str(base / 'tools' / 'script.py')), \
                 patch.object(sys, 'argv', ['script.py', '1000-pages']), \
                 patch.object(compare_text, 'write_updates') as writer, \
                 contextlib.redirect_stdout(io.StringIO()):
                compare_text.main()
            writer.assert_called_once()
            cells = {cell['table']: cell for cell in writer.call_args.args[1]}
            extraction = cells['ofd-text-extraction']
            self.assertEqual(extraction['input'], '1000-pages.ofd')
            self.assertEqual(extraction['converter'], 'go-zc310-ofd')
            self.assertEqual(extraction['value'], '49')
            self.assertEqual(cells['ofd-text-similarity']['value'], '100.0000%')

    def test_round_trip_pixel_failure_does_not_block_text_or_next_output(self):
        with tempfile.TemporaryDirectory() as tmp:
            base = Path(tmp)
            ref = base / 'testdata' / 'suwell' / 'hello.pdf'
            ref.parent.mkdir(parents=True)
            make_pdf(ref)
            output = base / 'output_ofd'
            output.mkdir(parents=True)
            for converter in ('pdf2ofd-ofdrw', 'pdf2ofd-zc310', 'node-jsofd', 'python-pdf2ofd',
                              'pdf2ofd-ofdgo', 'pdf2ofd-easyofd-rust'):
                stem = output / f'hello_{converter}'
                stem.with_suffix('.ofd').write_bytes(b'placeholder OFD')
                stem.with_suffix('.txt').write_text('Test textTest text', encoding='utf-8')
                if converter == 'pdf2ofd-ofdrw':
                    stem.with_suffix('.pdf').write_bytes(b'invalid PDF')
                elif converter != 'python-pdf2ofd':
                    make_pdf(stem.with_suffix('.pdf'), 1 if converter == 'node-jsofd' else 2)
            updates = self.run_script(compare_pdf2ofd, base)
            self.assertNotIn('rust-easyofd', {converter for converter, _ in updates})
            table = 'pdf-pixel-similarity'
            self.assertEqual(updates['java-ofdrw', table], 'FAILED')
            self.assertEqual(updates['go-zc310-ofd', table], '100.0000%')
            self.assertEqual(updates['go-xiaoqidun-ofdgo', table], '100.0000%')
            self.assertEqual(updates['node-jsofd', table], 'FAILED')
            self.assertEqual(updates['python-pdf2ofd', table], 'FAILED')
            for converter in ('java-ofdrw', 'go-zc310-ofd', 'node-jsofd', 'python-pdf2ofd', 'go-xiaoqidun-ofdgo'):
                self.assertEqual(updates[converter, 'pdf-ssim'], updates[converter, table])
            for converter in ('java-ofdrw', 'go-zc310-ofd', 'node-jsofd', 'python-pdf2ofd', 'go-xiaoqidun-ofdgo'):
                self.assertEqual(updates[converter, 'pdf-text-extraction'], '18')


if __name__ == '__main__':
    unittest.main()
