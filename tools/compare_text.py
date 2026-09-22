#!/usr/bin/env python3
"""Compare extractable text of converted PDFs against the reference PDF (数科版式 reader).

Usage:
    python3 tools/compare_text.py [file...]

Extracts text from both PDFs (PyMuPDF) and compares character count + text similarity.

Similarity = Jaccard of 4-grams of whitespace-normalized text (fast, O(N)),
100% means identical extracted text word-for-word.
"""
import sys
import os
import re
import fitz

def extract(path):
    doc = fitz.open(path)
    text = ''.join(p.get_text() for p in doc)
    doc.close()
    return text

def normalize(s):
    s = re.sub(r'\s+', ' ', s)
    return s.strip()

def ngrams(s, n=4):
    return {s[i:i + n] for i in range(len(s) - n + 1)}

def jaccard(a, b):
    if not a and not b:
        return 1.0
    ga, gb = ngrams(a), ngrams(b)
    if not ga and not gb:
        return 1.0
    return len(ga & gb) / len(ga | gb)

def compare(ref_path, src_path):
    ref = extract(ref_path).strip()
    src = extract(src_path).strip()
    ratio = len(src) / len(ref) if ref else 0.0
    sim = jaccard(normalize(ref), normalize(src)) * 100
    return len(ref), len(src), ratio, sim

def main():
    args = sys.argv[1:]
    files = args or ['hello', 'ano', '999', '1000-pages', 'intro', 'zsbk']
    base = os.path.dirname(os.path.dirname(os.path.abspath(__file__)))
    odir = os.path.join(base, 'bin', 'output')

    print(f"{'file':<12} {'convert':<16} {'ref_chars':>10} {'pdf_chars':>10} {'ratio':>6} {'similarity':>10}")
    print("-" * 76)
    for name in files:
        ref = os.path.join(base, 'test', 'testdata', f'{name}.pdf')
        rlen, _, _, _ = compare(ref, ref)
        for pdf in sorted(os.listdir(odir)):
            if not pdf.startswith(name + '_'):
                continue
            conv = pdf[len(name)+1:-4]
            src = os.path.join(odir, pdf)
            slen, ratio, sim = 0, 0.0, 0.0
            try:
                rlen, slen, ratio, sim = compare(ref, src)
            except Exception as e:
                print(f"{name:<12} {conv:<16} ERROR: {e}")
                continue
            print(f"{name:<12} {conv:<16} {rlen:>10} {slen:>10} {ratio:>6.2f} {sim:>9.2f}%")
        print()

if __name__ == '__main__':
    main()