#!/usr/bin/env python3
"""200 DPI adapters for native sOFD and OFD Manager renderers."""

from pathlib import Path, PurePosixPath
import re
import subprocess
import sys
import xml.etree.ElementTree as ET
from zipfile import ZipFile


def document_pages(input_file):
    """Read the first document page count for OFD Manager's single-doc CLI."""
    with ZipFile(input_file) as package:
        ofd = ET.fromstring(package.read('OFD.xml'))
        roots = [node.text.strip() for node in ofd.iter()
                 if node.tag.rsplit('}', 1)[-1] == 'DocRoot' and node.text]
        if len(roots) != 1:
            raise ValueError('OFD Manager CLI adapter requires exactly one document')
        path = PurePosixPath(roots[0].lstrip('/'))
        if '..' in path.parts:
            raise ValueError('unsupported DocRoot path')
        doc = ET.fromstring(package.read(str(path)))
        pages = next(node for node in doc if node.tag.rsplit('}', 1)[-1] == 'Pages')
        count = sum(node.tag.rsplit('}', 1)[-1] == 'Page' for node in pages)
        if not count:
            raise ValueError('no OFD pages')
        return count


def convert(library, input_file, output):
    binary_dir = Path(__file__).resolve().parents[1] / 'bin'
    if library == 'ofdmanager':
        # The upstream CLI renders one page per invocation. Include repeated
        # parsing/font discovery in the measured end-to-end cost, do not crop.
        for index in range(document_pages(input_file)):
            subprocess.run([str(binary_dir / 'ofdmanager-cli'), 'render', input_file,
                            str(output / f'page_{index + 1}.png'),
                            '--dpi', '200', '--page', str(index)], check=True)
    elif library == 'sofd':
        result = subprocess.run(
            [str(binary_dir / 'sofd-cli'), '--locale', 'en-US', 'convert',
             input_file, '--format', 'png', '--dpi', '200', '--output', str(output)],
            capture_output=True, text=True)
        print(result.stdout, end='')
        print(result.stderr, end='', file=sys.stderr)
        # Like the existing PDF adapter, retain complete output on the CLI's
        # documented missing-glyph warning. Any other nonzero exit is a failure;
        # the benchmark driver still checks actual PNG headers and total pages.
        glyph_warning = re.fullmatch(
            r'Warning: \d+ characters had no glyph in any available font, so the page is missing them\s*',
            result.stderr)
        if result.returncode and not (result.returncode == 1 and glyph_warning):
            raise subprocess.CalledProcessError(result.returncode, result.args)
        pattern = re.compile(re.escape(Path(input_file).stem) + r'-(\d+)\.png')
        pages = {}
        for path in output.iterdir():
            match = pattern.fullmatch(path.name)
            if not match:
                raise ValueError(f'unexpected sOFD output: {path}')
            pages[int(match[1])] = path
        if not pages or sorted(pages) != list(range(len(pages))):
            raise ValueError('incomplete sOFD page sequence')
        for index, path in pages.items():
            path.rename(output / f'page_{index + 1}.png')
    else:
        raise ValueError(f'unknown native CLI: {library}')


if __name__ == '__main__':
    if len(sys.argv) != 4:
        sys.exit('Usage: native_cli_png.py <sofd|ofdmanager> <input.ofd> <output-directory>')
    try:
        convert(sys.argv[1], sys.argv[2], Path(sys.argv[3]))
    except Exception as error:
        sys.exit(f'native PNG conversion failed: {error}')
