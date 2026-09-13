#!/bin/bash
# Node.js OFD to PDF converter using @miconvert/ofd-to-pdf
# Usage: ofd2pdf_node.sh <input.ofd> <output.pdf>

SCRIPT_DIR="$(cd "$(dirname "$0")" && pwd)"

node "$SCRIPT_DIR/ofd2pdf.mjs" "$1" "$2"
