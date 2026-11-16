#!/usr/bin/env bash
set -euo pipefail

if [[ $# -ne 2 ]]; then
  echo "Usage: sofd-convert.sh <input.ofd> <output.pdf>" >&2
  exit 2
fi

input=$1
output=$2
bin_dir="$(cd "$(dirname "$0")/../bin" && pwd)"
tmp_dir="$(mktemp -d)"
trap 'rm -rf "$tmp_dir"' EXIT

set +e
"$bin_dir/sofd-cli" --locale en-US convert "$input" --format pdf --output "$tmp_dir" \
  >"$tmp_dir/stdout" 2>"$tmp_dir/stderr"
status=$?
set -e
cat "$tmp_dir/stdout"
cat "$tmp_dir/stderr" >&2

input_name="$(basename "$input")"
input_stem="${input_name%.*}"
generated="$tmp_dir/$input_stem.pdf"
if [[ ! -s "$generated" ]] || [[ "$(head -c 5 "$generated")" != "%PDF-" ]]; then
  echo "sofd: no valid PDF was produced (exit status $status)" >&2
  exit 1
fi

mkdir -p "$(dirname "$output")"
mv "$generated" "$output"
if [[ $status -ne 0 ]]; then
  echo "sofd: retaining the generated PDF for visual comparison despite warnings (exit status $status)" >&2
fi
