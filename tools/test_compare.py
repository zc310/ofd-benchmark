"""Regression tests using small synthetic PDFs, without changing result YAML."""

import contextlib
import io
from pathlib import Path
import sys
import tempfile
import unittest
from unittest.mock import patch

import fitz

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
