#!/bin/bash
# Java OFD to PDF converter using ofdrw
# Usage: ofd2pdf_java.sh <input.ofd> <output.pdf>

SCRIPT_DIR="$(cd "$(dirname "$0")" && pwd)"
CP=$(ls "$SCRIPT_DIR"/lib/*.jar | tr '\n' ':')

java -cp "$CP:$SCRIPT_DIR" Ofd2Pdf "$1" "$2"
