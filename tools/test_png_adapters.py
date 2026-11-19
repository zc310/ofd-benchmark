"""Native CLI contracts: exact DPI and every page."""

from pathlib import Path
import tempfile
from types import SimpleNamespace
import unittest
from unittest.mock import patch
from zipfile import ZipFile

import native_cli_png


class NativeCLITests(unittest.TestCase):
    def test_manager_count_comes_from_ofd_page_tree(self):
        with tempfile.TemporaryDirectory() as tmp:
            source = Path(tmp) / 'input.ofd'
            with ZipFile(source, 'w') as package:
                package.writestr('OFD.xml', '<OFD xmlns="urn:ofd"><DocBody><DocRoot>/Doc_0/Document.xml</DocRoot></DocBody></OFD>')
                package.writestr('Doc_0/Document.xml', '<Document xmlns="urn:ofd"><Pages><Page ID="2"/><Page ID="3"/></Pages></Document>')
            output = Path(tmp) / 'output'
            output.mkdir()
            with patch('native_cli_png.subprocess.run') as run:
                native_cli_png.convert('ofdmanager', str(source), output)
            self.assertEqual(run.call_count, 2)
            for index, call in enumerate(run.call_args_list):
                command = call.args[0]
                self.assertEqual(command[command.index('--dpi') + 1], '200')
                self.assertEqual(command[command.index('--page') + 1], str(index))
                self.assertIn(str(output / f'page_{index + 1}.png'), command)

    def test_manager_rejects_multiple_documents_instead_of_truncating(self):
        with tempfile.TemporaryDirectory() as tmp:
            source = Path(tmp) / 'input.ofd'
            with ZipFile(source, 'w') as package:
                package.writestr('OFD.xml', '<OFD><DocRoot>A.xml</DocRoot><DocRoot>B.xml</DocRoot></OFD>')
            with self.assertRaisesRegex(ValueError, 'exactly one document'):
                native_cli_png.document_pages(source)

    def test_sofd_numeric_page_sequence(self):
        with tempfile.TemporaryDirectory() as tmp:
            output = Path(tmp)
            for index in range(12):
                (output / f'sample-{index}.png').write_bytes(str(index).encode())
            with patch('native_cli_png.subprocess.run',
                       return_value=SimpleNamespace(returncode=0, stdout='', stderr='')) as run:
                native_cli_png.convert('sofd', '/input/sample.ofd', output)
            self.assertIn('200', run.call_args.args[0])
            self.assertEqual((output / 'page_10.png').read_bytes(), b'9')
            self.assertEqual(len(list(output.iterdir())), 12)

    def test_sofd_missing_page_is_not_renumbered_away(self):
        with tempfile.TemporaryDirectory() as tmp:
            output = Path(tmp)
            for index in (0, 2):
                (output / f'sample-{index}.png').write_bytes(b'PNG')
            with patch('native_cli_png.subprocess.run',
                       return_value=SimpleNamespace(returncode=0, stdout='', stderr='')):
                with self.assertRaisesRegex(ValueError, 'incomplete'):
                    native_cli_png.convert('sofd', '/input/sample.ofd', output)

    def test_sofd_glyph_warning_is_retained_but_other_errors_fail(self):
        for warning, succeeds in (
            ('Warning: 2 characters had no glyph in any available font, so the page is missing them\n', True),
            ('error: malformed document\n', False),
        ):
            with self.subTest(warning=warning), tempfile.TemporaryDirectory() as tmp:
                output = Path(tmp)
                (output / 'sample-0.png').write_bytes(b'PNG')
                result = SimpleNamespace(returncode=1, stdout='', stderr=warning, args=['sofd'])
                with patch('native_cli_png.subprocess.run', return_value=result):
                    if succeeds:
                        native_cli_png.convert('sofd', '/input/sample.ofd', output)
                        self.assertTrue((output / 'page_1.png').exists())
                    else:
                        with self.assertRaises(Exception):
                            native_cli_png.convert('sofd', '/input/sample.ofd', output)


if __name__ == '__main__':
    unittest.main()
