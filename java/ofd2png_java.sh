#!/usr/bin/env bash
set -euo pipefail
script_dir="$(cd "$(dirname "$0")" && pwd)"
java -Djava.awt.headless=true -cp "$script_dir/lib/*:$script_dir" Ofd2Png "$@"
