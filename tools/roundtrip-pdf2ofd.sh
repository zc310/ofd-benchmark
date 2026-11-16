#!/usr/bin/env bash
set -euo pipefail

if [[ $# -ne 2 ]]; then
  echo "Usage: roundtrip-pdf2ofd.sh <ofd-converter> <output_ofd_dir>" >&2
  exit 2
fi

converter=$1
output_dir=$2
limit="${BENCHMARK_ROUNDTRIP_TIMEOUT:-15m}"
if [[ ! "$limit" =~ ^([0-9]+([.][0-9]+)?|[.][0-9]+)([smhd])?$ || ! "$limit" =~ [1-9] ]]; then
  echo "BENCHMARK_ROUNDTRIP_TIMEOUT must be a positive duration, e.g. 15m or 0.5s" >&2
  exit 2
fi
# Fail early on an invalid GNU timeout duration rather than hiding configuration
# errors as conversion failures. Each invocation has its own process group.
timeout --kill-after=5s "$limit" true
test -d "$output_dir"
shopt -s nullglob
for ofd in "$output_dir"/*.ofd; do
  stem="${ofd%.ofd}"
  for format in pdf txt; do
    output="$stem.$format"
    rm -f "$output"
    if timeout --kill-after=5s "$limit" "$converter" --format "$format" -o "$output" "$ofd"; then
      echo "Round-trip $format saved: $output"
    else
      status=$?
      rm -f "$output"
      echo "Round-trip $format failed: $ofd (exit $status; limit $limit)" >&2
    fi
  done
done
