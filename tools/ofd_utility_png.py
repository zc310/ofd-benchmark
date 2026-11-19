#!/usr/bin/env python3
"""Adapt the pinned ofd-utility CLI to <input.ofd> <output-directory>, 200 DPI."""

from pathlib import Path
import re
import subprocess
import sys


def convert(input_file, output_dir):
    binary = Path(__file__).resolve().parents[1] / 'bin' / 'ofd-utility-cli'
    subprocess.run([str(binary), 'render', input_file, str(output_dir),
                    '--dpi', '200', '--format', 'png', '--prefix', 'page'], check=True)
    pattern = re.compile(r'page_doc(\d+)_page(\d+)\.png')
    pages = []
    for path in output_dir.iterdir():
        match = pattern.fullmatch(path.name)
        if not match:
            raise ValueError(f'unexpected output: {path}')
        pages.append((tuple(map(int, match.groups())), path))
    for index, (_, path) in enumerate(sorted(pages), 1):
        path.rename(output_dir / f'page_{index}.png')


if __name__ == '__main__':
    if len(sys.argv) != 3:
        sys.exit('Usage: ofd_utility_png.py <input.ofd> <output-directory>')
    try:
        convert(sys.argv[1], Path(sys.argv[2]))
    except (OSError, ValueError, subprocess.CalledProcessError) as error:
        sys.exit(f'ofd-utility PNG conversion failed: {error}')
