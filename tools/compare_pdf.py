#!/usr/bin/env python3
"""Compare converted PDFs against the reference PDF (数科版式 reader output).

Usage:
    python3 tools/compare_pdf.py [dpi] [file...]

Renders each page of both PDFs and computes pixel-level similarity.
"""
import sys
import os
import io
import numpy as np
import fitz

def render(doc, page_idx, dpi):
    page = doc[page_idx]
    zoom = dpi / 72.0
    pix = page.get_pixmap(matrix=fitz.Matrix(zoom, zoom), alpha=False)
    img = np.frombuffer(pix.samples, dtype=np.uint8).reshape(pix.height, pix.width, pix.n)
    return img

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
    ref_doc = fitz.open(ref_path)
    src_doc = fitz.open(src_path)

    n = min(ref_doc.page_count, src_doc.page_count)
    total_diff = 0.0
    total_ident = 0.0
    for i in range(n):
        ref_img = render(ref_doc, i, dpi)
        src_img = render(src_doc, i, dpi)
        md, ident = similarity(ref_img, src_img)
        total_diff += md
        total_ident += ident

    ref_doc.close(); src_doc.close()

    mean_diff = total_diff / max(n, 1)
    pct = (1 - mean_diff) * 100
    ident = total_ident / max(n, 1) * 100
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
    for name in files:
        ref = os.path.join(base, 'test', 'testdata', f'{name}.pdf')
        ref_size = os.path.getsize(ref)
        for pdf in sorted(os.listdir(odir)):
            if not pdf.startswith(name + '_'):
                continue
            conv = pdf[len(name)+1:-4]
            src = os.path.join(odir, pdf)
            src_size = os.path.getsize(src)
            sim, md, ident = compare(ref, src, dpi)
            ratio = src_size / ref_size if ref_size else 0
            print(f"{name:<12} {conv:<16} {ref_size:>10} {src_size:>10} {ratio:>6.2f} {sim:>9.2f}% {ident:>8.2f}%")
        print()

if __name__ == '__main__':
    main()