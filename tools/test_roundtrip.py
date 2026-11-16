"""Round-trip timeout/failure handling without real document converters."""

import os
from pathlib import Path
import subprocess
import tempfile
import unittest


SCRIPT = Path(__file__).resolve().parent / 'roundtrip-pdf2ofd.sh'


class RoundTripTests(unittest.TestCase):
    def test_failures_and_timeouts_remove_partial_files_and_continue(self):
        with tempfile.TemporaryDirectory(prefix='roundtrip test ') as tmp:
            base = Path(tmp)
            converter = base / 'fake converter'
            converter.write_text('''#!/usr/bin/env bash
set -eu
format=$2
output=$4
input=$5
printf partial > "$output"
if [[ "$input" == *fail.ofd ]]; then exit 7; fi
if [[ "$input" == *hang.ofd && "$format" == pdf ]]; then sleep 5; fi
printf complete > "$output"
''')
            converter.chmod(0o755)
            for name in ('fail', 'hang', 'ok'):
                (base / f'{name}.ofd').write_bytes(b'placeholder OFD')
                (base / f'{name}.pdf').write_bytes(b'stale PDF')
                (base / f'{name}.txt').write_bytes(b'stale text')
            env = dict(os.environ, BENCHMARK_ROUNDTRIP_TIMEOUT='0.1s')
            result = subprocess.run(['bash', str(SCRIPT), str(converter), str(base)],
                                    env=env, capture_output=True, text=True, timeout=10)
            self.assertEqual(result.returncode, 0, result.stderr)
            for path in ('fail.pdf', 'fail.txt', 'hang.pdf'):
                self.assertFalse((base / path).exists(), path)
            for path in ('hang.txt', 'ok.pdf', 'ok.txt'):
                self.assertEqual((base / path).read_text(), 'complete', path)
            self.assertIn('exit 124', result.stderr)
            self.assertIn('exit 7', result.stderr)

    def test_invalid_timeout_fails_before_conversion(self):
        with tempfile.TemporaryDirectory() as tmp:
            for value in ('invalid', '0', '0m', '-1', '1m30s'):
                with self.subTest(value=value):
                    result = subprocess.run(['bash', str(SCRIPT), '/not/run', tmp],
                                            env=dict(os.environ, BENCHMARK_ROUNDTRIP_TIMEOUT=value),
                                            capture_output=True, text=True, timeout=5)
                    self.assertNotEqual(result.returncode, 0)


if __name__ == '__main__':
    unittest.main()
