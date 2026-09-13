#!/usr/bin/env python3
import sys
import warnings
warnings.filterwarnings("ignore")

if __name__ == '__main__':
    if len(sys.argv) < 3:
        print('Usage: ofd2pdf_ofdreader.py <input.ofd> <output.pdf>')
        sys.exit(1)

    try:
        from ofdreader import ofd_to_pdf
        ofd_to_pdf(sys.argv[1], sys.argv[2])
        print("LIB=ofdreader")
    except Exception as e:
        print(f"FAILED: {e}", file=sys.stderr)
        sys.exit(1)
