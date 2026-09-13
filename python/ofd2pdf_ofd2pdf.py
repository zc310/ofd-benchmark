#!/usr/bin/env python3
"""OFD to PDF converter using ofd2pdf"""
import sys
import os

def main():
    if len(sys.argv) < 3:
        print("Usage: ofd2pdf_wrapper.py <input.ofd> <output.pdf>", file=sys.stderr)
        sys.exit(1)
    
    input_file = sys.argv[1]
    output_file = sys.argv[2]
    
    # Add ofd2pdf to path
    ofd2pdf_path = os.path.expanduser("~/.local/bin")
    os.environ["PATH"] = ofd2pdf_path + ":" + os.environ.get("PATH", "")
    
    # Run ofd2pdf
    cmd = f'ofd2pdf "{input_file}" -o "{output_file}"'
    exit_code = os.system(cmd)
    sys.exit(exit_code)

if __name__ == "__main__":
    main()
