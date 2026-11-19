from pathlib import Path
import tempfile
import unittest
from unittest.mock import patch

from ofd_utility_png import convert


class UtilityAdapterTests(unittest.TestCase):
    def test_numeric_page_and_document_order(self):
        with tempfile.TemporaryDirectory() as tmp:
            output = Path(tmp)
            for number in range(12):
                (output / f'page_doc0_page{number}.png').write_bytes(str(number).encode())
            (output / 'page_doc1_page0.png').write_bytes(b'next document')
            with patch('ofd_utility_png.subprocess.run') as run:
                convert('input.ofd', output)
            command = run.call_args.args[0]
            self.assertIn('200', command)
            self.assertTrue(run.call_args.kwargs['check'])
            self.assertEqual((output / 'page_10.png').read_bytes(), b'9')
            self.assertEqual((output / 'page_13.png').read_bytes(), b'next document')

    def test_unexpected_output_is_not_silently_ignored(self):
        with tempfile.TemporaryDirectory() as tmp:
            output = Path(tmp)
            (output / 'unexpected.png').write_bytes(b'PNG')
            with patch('ofd_utility_png.subprocess.run'):
                with self.assertRaisesRegex(ValueError, 'unexpected output'):
                    convert('input.ofd', output)


if __name__ == '__main__':
    unittest.main()
