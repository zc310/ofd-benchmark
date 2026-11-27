"""Shared locations for benchmark outputs."""

import os
from pathlib import Path


def output_root():
    return Path(os.environ.get('BENCHMARK_OUTPUT_DIR') or '/tmp/ofd-benchmark').absolute()
