#!/usr/bin/env bash
set -euo pipefail

# Install the same document fonts for local and CI benchmark runs.
font_dir="$HOME/.local/share/fonts/hanzi-fonts"
temp_dir="$(mktemp -d)"
trap 'rm -rf "$temp_dir"' EXIT

base_url="https://github.com/zispace/hanzi-fonts/releases/download/v1.0"
for name in doc-foundertype doc-zhongyi-etc latin-fonts; do
  curl --fail --location --show-error --retry 3 --connect-timeout 30 --max-time 600 \
    "$base_url/$name.zip" -o "$temp_dir/$name.zip"
  mkdir -p "$font_dir/$name"
  unzip -q -o "$temp_dir/$name.zip" -d "$font_dir/$name"
done

fc-cache -f "$font_dir"
echo "Fonts installed in $font_dir"
