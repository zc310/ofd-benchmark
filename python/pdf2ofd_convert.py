#!/usr/bin/env python3
"""PDF -> OFD converter using the pdf2ofd pip package.

Usage:
    python3 pdf2ofd_convert.py <input.pdf> <output.ofd>
"""
import sys

from pdf2ofd import convert

if __name__ == '__main__':
    if len(sys.argv) != 3:
        sys.stderr.write(f"Usage: {sys.argv[0]} <in.pdf> <out.ofd>\n")
        sys.exit(1)
    convert(sys.argv[1], sys.argv[2], verbose=False)