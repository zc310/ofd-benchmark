#!/usr/bin/env python3
import sys
import base64
import warnings
warnings.filterwarnings("ignore")

def try_easyofd(ofd_path, pdf_path):
    from easyofd.ofd import OFD
    with open(ofd_path, 'rb') as f:
        ofdb64 = str(base64.b64encode(f.read()), 'utf-8')
    ofd = OFD()
    ofd.read(ofdb64, save_xml=False)
    pdf_bytes = ofd.to_pdf()
    ofd.del_data()
    with open(pdf_path, 'wb') as f:
        f.write(pdf_bytes)

LIBRARIES = [
    ("easyofd", try_easyofd),
]

def convert(ofd_path, pdf_path):
    errors = []
    for name, fn in LIBRARIES:
        try:
            fn(ofd_path, pdf_path)
            return name, None
        except Exception as e:
            errors.append(f"{name}: {e}")
    return None, "; ".join(errors)

if __name__ == '__main__':
    if len(sys.argv) < 3:
        print('Usage: ofd2pdf.py <input.ofd> <output.pdf>')
        sys.exit(1)

    lib, err = convert(sys.argv[1], sys.argv[2])
    if lib:
        print(f"LIB={lib}")
        sys.exit(0)
    else:
        print(f"FAILED: {err}", file=sys.stderr)
        sys.exit(1)
