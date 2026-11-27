#!/usr/bin/env python3
"""Quality comparison for the PDF -> OFD benchmark.

For each input PDF, each converter's OFD output is evaluated three ways:
  * OFD size ratio vs the original PDF size
  * round-trip fidelity: OFD -> PDF (via go-zc310), then pixel similarity and SSIM
    against the original PDF (same metric as compare_pdf.py; 1000-pages samples
    the first three pages while still validating the total page count)
  * text fidelity: characters extracted from the OFD (go-zc310 ofd-converter)
    vs the text extracted from the original PDF (4-gram Jaccard)

Usage:
    python3 tools/compare_pdf2ofd.py [dpi] [file...]
"""
import sys
import os
import re

import fitz
from compare_pdf import compare
from benchmark_paths import output_root
from benchmark_results import PDF_CONVERTER_ROWS, append, format_percentage, write_updates


def pixel_similarity(ref_path, src_path, dpi):
    return compare(ref_path, src_path, dpi)[0]


def normalize(s):
    return re.sub(r'\s+', ' ', s).strip()


def ngrams(s, n=4):
    return {s[i:i + n] for i in range(len(s) - n + 1)}


def text_similarity(ref_path, ofd_txt_path):
    with fitz.open(ref_path) as ref:
        ref_text = ''.join(p.get_text() for p in ref)
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
    odir = str(output_root() / 'output_ofd')

    print(f"{'file':<12} {'convert':<16} {'pdf_ref':>9} {'ofd_size':>10} {'size_rat':>7} {'pixel':>10} {'SSIM':>10} {'ref_ch':>9} {'ofd_ch':>9} {'txt_sim':>10}")
    print("-" * 106)
    updates = []
    for name in files:
        ref_pdf = next(
            (p for sub in ('suwell', 'standards')
             if os.path.exists(p := os.path.join(base, 'testdata', sub, f'{name}.pdf'))),
            os.path.join(base, 'testdata', 'suwell', f'{name}.pdf'),
        )
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
            pixel = 'FAILED'
            ssim = 'FAILED'
            rch, sch, sim = 'FAILED', 'FAILED', 'FAILED'
            if os.path.exists(rt_pdf):
                try:
                    scores = compare(ref_pdf, rt_pdf, dpi)
                    pixel = format_percentage(scores[0])
                    ssim = format_percentage(scores[5])
                except Exception as e:
                    print(f"{name:<12} {conv:<16} pixel FAILED: {e}")
            if os.path.exists(txt):
                try:
                    rch, sch, text_sim = text_similarity(ref_pdf, txt)
                    sim = format_percentage(text_sim)
                except Exception as e:
                    print(f"{name:<12} {conv:<16} text FAILED: {e}")
            print(f"{name:<12} {conv:<16} {ref_size:>9} {ofd_size:>10} {ratio:>6.2f} {pixel:>10} {ssim:>10} {rch:>9} {sch:>9} {sim:>10}")
            row = PDF_CONVERTER_ROWS.get(conv)
            if row:
                append(updates, 'pdf-to-ofd', 'pdf-size-ratio', f'{name}.pdf', row, f'{ratio:.2f}')
                append(updates, 'pdf-to-ofd', 'pdf-pixel-similarity', f'{name}.pdf', row, pixel)
                append(updates, 'pdf-to-ofd', 'pdf-ssim', f'{name}.pdf', row, ssim)
                append(updates, 'pdf-to-ofd', 'pdf-text-extraction', f'{name}.pdf', row, str(sch))
                append(updates, 'pdf-to-ofd', 'pdf-text-similarity', f'{name}.pdf', row, sim)
        print()
    write_updates(base, updates)


if __name__ == '__main__':
    main()
