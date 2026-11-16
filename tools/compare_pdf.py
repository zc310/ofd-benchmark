#!/usr/bin/env python3
"""Compare converted PDFs against the reference PDF (数科版式 reader output).

Usage:
    python3 tools/compare_pdf.py [dpi] [file...]

Renders each page of both PDFs and computes pixel-level similarity.
"""
import sys
import os
import numpy as np
import fitz
from benchmark_results import OFD_CONVERTER_ROWS, append, write_updates

def render(doc, page_idx, dpi):
    page = doc[page_idx]
    zoom = dpi / 72.0
    pix = page.get_pixmap(matrix=fitz.Matrix(zoom, zoom), alpha=False)
    img = np.frombuffer(pix.samples, dtype=np.uint8).reshape(pix.height, pix.width, pix.n)
    return img

# intro.pdf（数科导出的参考 PDF）第 14–19 页背景图丢失，像素对比跳过这些页
EXCLUDE_PAGES = {'intro': {14, 15, 16, 17, 18, 19}}

def resize_to(img, h, w):
    from PIL import Image
    im = Image.fromarray(img).resize((w, h), Image.LANCZOS)
    return np.asarray(im)

def similarity(ref, src):
    # Align sizes if they differ
    if ref.shape != src.shape:
        h = min(ref.shape[0], src.shape[0])
        w = min(ref.shape[1], src.shape[1])
        ref = ref[:h, :w]
        src = src[:h, :w]
    diff = np.abs(ref.astype(np.int16) - src.astype(np.int16)).sum(axis=2) / (3 * 255.0)
    mean_diff = diff.mean()
    identical = (diff < 0.01).mean()
    return mean_diff, identical

def compare(ref_path, src_path, dpi):
    name = os.path.splitext(os.path.basename(ref_path))[0]
    skip = EXCLUDE_PAGES.get(name, set())
    total_diff = 0.0
    total_ident = 0.0
    count = 0
    with fitz.open(ref_path) as ref_doc, fitz.open(src_path) as src_doc:
        if ref_doc.page_count != src_doc.page_count:
            raise ValueError(f'page count mismatch: reference={ref_doc.page_count}, output={src_doc.page_count}')
        for i in range(ref_doc.page_count):
            if i + 1 in skip:
                continue
            ref_img = render(ref_doc, i, dpi)
            src_img = render(src_doc, i, dpi)
            # A nearly empty PDF can score deceptively well because most pages are white.
            # Fail pages that lost virtually all foreground content instead of reporting a
            # misleadingly high whole-page similarity.
            ref_ink = np.count_nonzero(np.any(ref_img < 245, axis=2))
            src_ink = np.count_nonzero(np.any(src_img < 245, axis=2))
            if ref_ink >= 16 and src_ink < max(4, ref_ink * 0.01):
                raise ValueError(f'output page {i + 1} is nearly blank compared with reference')
            md, ident = similarity(ref_img, src_img)
            total_diff += md
            total_ident += ident
            count += 1

    if not count:
        raise ValueError('no comparable pages')
    mean_diff = total_diff / count
    pct = (1 - mean_diff) * 100
    ident = total_ident / count * 100
    return pct, mean_diff, ident

def main():
    dpi = 96
    args = sys.argv[1:]
    if args and args[0].isdigit():
        dpi = int(args[0])
        args = args[1:]

    files = args or ['hello', 'ano', '999', '1000-pages', 'intro', 'zsbk']
    base = os.path.dirname(os.path.dirname(os.path.abspath(__file__)))
    odir = os.path.join(base, 'bin', 'output')

    print(f"{'file':<12} {'convert':<16} {'ref_size':>10} {'pdf_size':>10} {'ratio':>6} {'similarity':>10} {'identical':>9}")
    print("-" * 88)
    updates = []
    for name in files:
        ref = os.path.join(base, 'testdata', 'suwell', f'{name}.pdf')
        ref_size = os.path.getsize(ref)
        for pdf in sorted(os.listdir(odir)):
            if not pdf.startswith(name + '_') or not pdf.endswith('.pdf'):
                continue
            conv = pdf[len(name)+1:-4]
            src = os.path.join(odir, pdf)
            src_size = os.path.getsize(src)
            ratio = src_size / ref_size if ref_size else 0
            pixel = 'FAILED'
            try:
                sim, md, ident = compare(ref, src, dpi)
                pixel = f'{sim:.2f}%'
                print(f"{name:<12} {conv:<16} {ref_size:>10} {src_size:>10} {ratio:>6.2f} {sim:>9.2f}% {ident:>8.2f}%")
            except Exception as e:
                print(f"{name:<12} {conv:<16} FAILED: {e}")
            row = OFD_CONVERTER_ROWS.get(conv)
            if row:
                append(updates, 'ofd-to-pdf', 'ofd-size-ratio', f'{name}.ofd', row, f'{ratio:.2f}')
                append(updates, 'ofd-to-pdf', 'ofd-pixel-similarity', f'{name}.ofd', row, pixel)
        print()
    write_updates(base, updates)

if __name__ == '__main__':
    main()
