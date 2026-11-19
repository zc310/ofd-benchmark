#!/usr/bin/env python3
"""Compare native 200 DPI PNG output with direct Suwell exports."""

import argparse
from pathlib import Path
import re

import numpy as np
from PIL import Image

from benchmark_results import append, write_updates
from compare_pdf import align_images, content_similarity, foreground_f1, similarity

DPI = 200
CONVERTERS = ('go-zc310', 'go-ofdgo', 'go-ofdrw-rs', 'rust-ofdbox-folio',
              'rust-sofd', 'java-ofdrw', 'rust-ofdmanager', 'rust-ofd-utility')


def page_files(directory, prefix):
    pattern = re.compile(re.escape(prefix) + r'_([1-9]\d*)\.png')
    pages = {}
    for path in directory.iterdir():
        match = pattern.fullmatch(path.name)
        if not match or not path.is_file() or path.is_symlink():
            raise ValueError(f'unexpected page file: {path}')
        pages[int(match[1])] = path
    if not pages:
        raise ValueError(f'no PNG pages: {directory}')
    if sorted(pages) != list(range(1, len(pages) + 1)):
        raise ValueError(f'nonconsecutive page numbers: {directory}')
    return pages


def read_png(path, reference=False):
    with Image.open(path) as image:
        if image.format != 'PNG':
            raise ValueError(f'not a PNG: {path}')
        dpi = image.info.get('dpi')
        if reference and dpi and any(abs(value - DPI) > 1 for value in dpi):
            raise ValueError(f'reference is not {DPI} DPI: {path}: {dpi}')
        # Normalize alpha/palette/grayscale without resizing or cropping.
        rgba = image.convert('RGBA')
        white = Image.new('RGBA', rgba.size, (255, 255, 255, 255))
        return np.array(Image.alpha_composite(white, rgba).convert('RGB'))


def compare(reference_dir, output_dir, name):
    refs = page_files(Path(reference_dir), name)
    outputs = page_files(Path(output_dir), 'page')
    if refs.keys() != outputs.keys():
        raise ValueError(f'page count mismatch: reference={len(refs)}, output={len(outputs)}')
    full_scores, content_scores, f1_scores = [], [], []
    # Every page is compared, including intro pages 14-19. PDF exclusions and
    # repeated-page sampling intentionally do not apply to this PNG benchmark.
    for number in sorted(refs):
        try:
            ref = read_png(refs[number], reference=True)
            src = read_png(outputs[number])
            ref, src = align_images(ref, src)
            ref_ink = np.count_nonzero(np.any(ref < 245, axis=2))
            src_ink = np.count_nonzero(np.any(src < 245, axis=2))
            if ref_ink >= 16 and src_ink < max(4, ref_ink * 0.01):
                raise ValueError('output is nearly blank compared with reference')
            full_scores.append((1 - similarity(ref, src)[0]) * 100)
            content = content_similarity(ref, src)
            if content is not None:
                content_scores.append(content)
                f1_scores.append(foreground_f1(ref, src, DPI))
        except Exception as error:
            raise ValueError(f'page {number}: {error}') from error
    return (sum(full_scores) / len(full_scores),
            sum(content_scores) / len(content_scores) if content_scores else None,
            sum(f1_scores) / len(f1_scores) if f1_scores else None)


def main():
    parser = argparse.ArgumentParser(description=__doc__)
    parser.add_argument('files', nargs='*', default=['999', 'intro', 'zsbk'])
    parser.add_argument('--converters', nargs='+', choices=CONVERTERS, default=CONVERTERS)
    parser.add_argument('--no-write', action='store_true', help='do not update result YAML')
    args = parser.parse_args()
    base = Path(__file__).resolve().parents[1]
    updates = []
    print(f'OFD -> PNG: {DPI} DPI; all pages; direct Suwell reference PNGs')
    print(f'{"file":<10} {"converter":<20} {"full-page":>12} {"content":>12} {"fg-F1":>12}')
    for name in args.files:
        if name not in ('999', 'intro', 'zsbk'):
            parser.error(f'no configured reference for {name}')
        for converter in args.converters:
            values = ['FAILED'] * 3
            try:
                scores = compare(base / 'testdata' / 'suwell' / name,
                                 base / 'bin' / 'output_png' / f'{name}_{converter}', name)
                values = [f'{value:.2f}%' if value is not None else 'FAILED' for value in scores]
                print(f'{name:<10} {converter:<20} {values[0]:>12} {values[1]:>12} {values[2]:>12}')
            except Exception as error:
                print(f'{name:<10} {converter:<20} FAILED: {error}')
            for metric, value in zip(('pixel-similarity', 'content-similarity', 'foreground-f1'), values):
                append(updates, 'ofd-to-png', f'png-{metric}', name + '.ofd', converter, value)
    if not args.no_write:
        write_updates(str(base), updates)


if __name__ == '__main__':
    main()
