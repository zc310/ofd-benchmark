#!/bin/bash
# Java PDF to OFD converter using ofdrw (PDFConverter)
# Usage: pdf2ofd_java.sh <input.pdf> <output.ofd>

SCRIPT_DIR="$(cd "$(dirname "$0")" && pwd)"
CP=$(ls "$SCRIPT_DIR"/lib/*.jar | tr '\n' ':')

java -cp "$CP:$SCRIPT_DIR" Pdf2Ofd "$1" "$2"