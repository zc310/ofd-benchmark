"""PNG regression tests: no fixture downloads or result YAML writes."""

from pathlib import Path
import tempfile
import unittest
from unittest.mock import patch

import numpy as np
from PIL import Image

import compare_png


class PNGComparisonTests(unittest.TestCase):
    def setUp(self):
        self.tmp = tempfile.TemporaryDirectory()
        self.addCleanup(self.tmp.cleanup)
        self.ref = Path(self.tmp.name) / 'ref'
        self.out = Path(self.tmp.name) / 'out'
        self.ref.mkdir()
        self.out.mkdir()
        self.image = np.full((30, 40, 3), 255, dtype=np.uint8)
        self.image[5:25, 10:30] = 0

    def save_page(self, number, ref=None, src=None, dpi=200):
        Image.fromarray(self.image if ref is None else ref).save(
            self.ref / f'intro_{number}.png', dpi=(dpi, dpi))
        Image.fromarray(self.image if src is None else src).save(
            self.out / f'page_{number}.png')

    def compare(self):
        return compare_png.compare(self.ref, self.out, 'intro')

    def test_identical_and_numeric_page_order(self):
        for number in range(1, 16):
            self.save_page(number)
        self.assertEqual(self.compare(), (100.0, 100.0, 100.0))

    def test_intro_page_14_is_not_excluded(self):
        for number in range(1, 20):
            self.save_page(number)
        Image.fromarray(np.full_like(self.image, 255)).save(self.out / 'page_14.png')
        with self.assertRaisesRegex(ValueError, 'page 14:.*nearly blank'):
            self.compare()

    def test_missing_extra_and_nonconsecutive_pages(self):
        self.save_page(1)
        self.save_page(2)
        (self.out / 'page_2.png').unlink()
        with self.assertRaisesRegex(ValueError, 'page count mismatch'):
            self.compare()
        (self.ref / 'intro_1.png').rename(self.ref / 'intro_3.png')
        with self.assertRaisesRegex(ValueError, 'nonconsecutive'):
            self.compare()

    def test_size_mismatch_is_not_rescaled(self):
        self.save_page(1, src=self.image[:, :30])
        with self.assertRaisesRegex(ValueError, 'size mismatch'):
            self.compare()

    def test_one_pixel_rounding_is_allowed(self):
        self.save_page(1, src=np.pad(self.image, ((0, 1), (0, 1), (0, 0)), constant_values=255))
        self.assertEqual(self.compare(), (100.0, 100.0, 100.0))

    def test_reference_dpi_is_checked(self):
        self.save_page(1, dpi=96)
        with self.assertRaisesRegex(ValueError, 'not 200 DPI'):
            self.compare()

    def test_transparent_png_is_composited_on_white(self):
        self.save_page(1)
        rgba = np.zeros((30, 40, 4), dtype=np.uint8)
        rgba[5:25, 10:30, 3] = 255
        Image.fromarray(rgba).save(self.out / 'page_1.png')
        self.assertEqual(self.compare(), (100.0, 100.0, 100.0))

    def test_corrupt_output_and_empty_output_fail(self):
        with self.assertRaisesRegex(ValueError, 'no PNG'):
            self.compare()
        self.save_page(1)
        (self.out / 'page_1.png').write_bytes(b'not PNG')
        with self.assertRaisesRegex(ValueError, 'page 1'):
            self.compare()

    def test_blank_pages_have_no_content_score(self):
        blank = np.full_like(self.image, 255)
        self.save_page(1, ref=blank, src=blank)
        self.assertEqual(self.compare(), (100.0, None, None))

    def test_missing_output_resets_all_quality_cells(self):
        with patch('sys.argv', ['compare_png.py', '999']), \
                patch('compare_png.compare', side_effect=FileNotFoundError('missing output')), \
                patch('compare_png.write_updates') as write, \
                patch('builtins.print'):
            compare_png.main()
        updates = write.call_args.args[1]
        self.assertEqual(len(updates), len(compare_png.CONVERTERS) * 3)
        self.assertTrue(all(update['value'] == 'FAILED' for update in updates))
        for excluded in ('python-ofd2img', 'python-easyofd', 'node-ofd2pdf', 'rust-easyofd'):
            self.assertNotIn(excluded, compare_png.CONVERTERS)

    def test_no_write_does_not_change_yaml(self):
        with patch('sys.argv', ['compare_png.py', '999', '--no-write']), \
                patch('compare_png.compare', return_value=(100, 100, 100)), \
                patch('compare_png.write_updates') as write, \
                patch('builtins.print'):
            compare_png.main()
        write.assert_not_called()


if __name__ == '__main__':
    unittest.main()
