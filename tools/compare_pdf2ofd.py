#!/usr/bin/env python3
"""Quality comparison for the PDF -> OFD benchmark.

For each input PDF, each converter's OFD output is evaluated three ways:
  * OFD size ratio vs the original PDF size
  * round-trip fidelity: OFD -> PDF (via go-zc310), then pixel similarity
    against the original PDF (same metric as compare_pdf.py)
  * text fidelity: characters extracted from the OFD (go-zc310 ofd-converter)
    vs the text extracted from the original PDF (4-gram Jaccard)

Usage:
    python3 tools/compare_pdf2ofd.py [dpi] [file...]
"""
import sys
import os
import re

import numpy as np
import fitz


def render(doc, page_idx, dpi):
    page = doc[page_idx]
    zoom = dpi / 72.0
    pix = page.get_pixmap(matrix=fitz.Matrix(zoom, zoom), alpha=False)
    return np.frombuffer(pix.samples, dtype=np.uint8).reshape(pix.height, pix.width, pix.n)


def pixel_similarity(ref_path, src_path, dpi):
    ref_doc = fitz.open(ref_path)
    src_doc = fitz.open(src_path)
    n = min(ref_doc.page_count, src_doc.page_count)
    total = 0.0
    for i in range(n):
        r = render(ref_doc, i, dpi)
        s = render(src_doc, i, dpi)
        if r.shape != s.shape:
            h = min(r.shape[0], s.shape[0])
            w = min(r.shape[1], s.shape[1])
            r = r[:h, :w]
            s = s[:h, :w]
        diff = np.abs(r.astype(np.int16) - s.astype(np.int16)).sum(axis=2) / (3 * 255.0)
        total += diff.mean()
    ref_doc.close()
    src_doc.close()
    return (1 - total / max(n, 1)) * 100


def normalize(s):
    return re.sub(r'\s+', ' ', s).strip()


def ngrams(s, n=4):
    return {s[i:i + n] for i in range(len(s) - n + 1)}


def text_similarity(ref_path, ofd_txt_path):
    ref = fitz.open(ref_path)
    ref_text = ''.join(p.get_text() for p in ref)
    ref.close()
    with open(ofd_txt_path, encoding='utf-8', errors='replace') as fh:
        src_text = fh.read()
    ref_n = normalize(ref_text)
    src_n = normalize(src_text)
    ref_chars = len(ref_text.strip())
    src_chars = len(src_text.strip())
    if not ref_n and not src_n:
        sim = 100.0
    elif not ref_n or not src_n:
        sim = 0.0
    else:
        ga, gb = ngrams(ref_n), ngrams(src_n)
        sim = len(ga & gb) / len(ga | gb) * 100
    return ref_chars, src_chars, sim


def main():
    dpi = 96
    args = sys.argv[1:]
    if args and args[0].isdigit():
        dpi = int(args[0])
        args = args[1:]

    files = args or ['hello', 'ano', '999', '1000-pages', 'intro', 'zsbk', 'GBT_33190-2016']
    base = os.path.dirname(os.path.dirname(os.path.abspath(__file__)))
    odir = os.path.join(base, 'bin', 'output_ofd')

    print(f"{'file':<12} {'convert':<16} {'pdf_ref':>9} {'ofd_size':>10} {'size_rat':>7} {'pixel':>9} {'ref_ch':>9} {'ofd_ch':>9} {'txt_sim':>9}")
    print("-" * 96)
    for name in files:
        ref_pdf = os.path.join(base, 'test', 'testdata', f'{name}.pdf')
        ref_size = os.path.getsize(ref_pdf)
        for ofd in sorted(os.listdir(odir)):
            if not ofd.endswith('.ofd'):
                continue
            stem = ofd[: -len('.ofd')]
            if not stem.startswith(name + '_'):
                continue
            conv = stem[len(name) + 1:]
            ofd_path = os.path.join(odir, ofd)
            ofd_size = os.path.getsize(ofd_path)
            ratio = ofd_size / ref_size if ref_size else 0.0
            rt_pdf = os.path.join(odir, stem + '.pdf')
            txt = os.path.join(odir, stem + '.txt')
            pixel = pixel_similarity(ref_pdf, rt_pdf, dpi) if os.path.exists(rt_pdf) else float('nan')
            rch, sch, sim = text_similarity(ref_pdf, txt) if os.path.exists(txt) else (0, 0, 0.0)
            print(f"{name:<12} {conv:<16} {ref_size:>9} {ofd_size:>10} {ratio:>6.2f} {pixel:>8.2f}% {rch:>9} {sch:>9} {sim:>8.2f}%")
        print()


if __name__ == '__main__':
    main()