"""Release packaging/publishing tests without converters or GitHub access."""

import hashlib
import json
import os
from pathlib import Path
import re
import subprocess
import tempfile
import unittest
import zipfile

from package_benchmark_release import LIBRARIES, OFD_SAMPLES, PDF_SAMPLES, REPORTS, package


ROOT = Path(__file__).resolve().parent.parent
PUBLISH = ROOT / 'tools' / 'publish-benchmark-release.sh'


class ReleaseTests(unittest.TestCase):
    def setUp(self):
        self.tmp = tempfile.TemporaryDirectory(prefix='benchmark release ')
        self.addCleanup(self.tmp.cleanup)
        self.root = Path(self.tmp.name)
        self.output = self.root / 'release assets'
        for directory in ('docs', 'bin/output', 'bin/output_ofd'):
            (self.root / directory).mkdir(parents=True)
        for name in REPORTS:
            (self.root / 'docs' / name).write_text(f'report: {name}\n')

    def document(self, name, content=b'document', directory='output'):
        path = self.root / 'bin' / directory / name
        path.write_bytes(content)
        return path

    def bundle(self, tag='v2026.11.16'):
        return package(self.root, self.output, tag, 'commit123', 'https://example.org/run/1')

    def test_all_converters_are_classified_and_directions_are_separate(self):
        expected = {}
        for library, (ofd_to_pdf, pdf_to_ofd) in LIBRARIES.items():
            files = {}
            if ofd_to_pdf:
                source = self.document(f'hello_{ofd_to_pdf}.pdf', ofd_to_pdf.encode())
                files[f'{library}/ofd-to-pdf/hello.pdf'] = source.read_bytes()
            if pdf_to_ofd:
                for extension, direction in (('ofd', 'pdf-to-ofd'), ('pdf', 'roundtrip')):
                    source = self.document(f'hello_{pdf_to_ofd}.{extension}',
                                           f'{pdf_to_ofd}.{extension}'.encode(), 'output_ofd')
                    files[f'{library}/{direction}/hello.{extension}'] = source.read_bytes()
            expected[library] = files
        # Text and temporary outputs are not document attachments.
        self.document('hello_pdf2ofd-zc310.txt', directory='output_ofd')
        manifest = self.bundle()
        self.assertEqual(manifest['commit'], 'commit123')
        self.assertEqual(json.loads((self.output / 'manifest.json').read_text()), manifest)
        for library, files in expected.items():
            with zipfile.ZipFile(self.output / manifest['libraries'][library]['archive']) as bundle:
                self.assertEqual(set(bundle.namelist()), set(files))
                for name, content in files.items():
                    self.assertEqual(bundle.read(name), content)
            self.assertEqual(len(manifest['libraries'][library]['files']), len(files))
        for name in REPORTS:
            self.assertEqual((self.output / name).read_bytes(),
                             (self.root / 'docs' / name).read_bytes())
        lines = (self.output / 'SHA256SUMS').read_text().splitlines()
        self.assertEqual(len(lines), len(list(self.output.iterdir())) - 1)
        for line in lines:
            digest, name = line.split('  ', 1)
            self.assertEqual(digest, hashlib.sha256((self.output / name).read_bytes()).hexdigest())

    def test_release_notes_link_all_libraries_including_failed_ones(self):
        self.document('hello_go-zc310.pdf')
        manifest = self.bundle()
        notes = (self.output / 'RELEASE_NOTES.md').read_text()
        urls = {
            'zc310': 'https://github.com/zc310/ofd',
            'ofdgo': 'https://github.com/xiaoqidun/ofdgo',
            'ofdrw-java': 'https://github.com/ofdrw/ofdrw',
            'ofdrw-rs': 'https://github.com/lihongjie0209/ofdrw-rs',
            'ofdbox-folio': 'https://github.com/ofdbox/ofdbox-folio',
            'sofd': 'https://github.com/emfox/sOFD',
            'easyofd-python': 'https://pypi.org/project/easyofd/',
            'ofd-to-pdf-node': 'https://www.npmjs.com/package/@miconvert/ofd-to-pdf',
            'pdf2ofd-python': 'https://github.com/wanglrebe/pdf2ofd',
            'jsofd': 'https://github.com/Hufe921/jsOFD',
        }
        self.assertEqual(set(urls), set(LIBRARIES))
        for library, url in urls.items():
            record = manifest['libraries'][library]
            self.assertIn(
                f'| [{library}]({url}) | {len(record["files"])} | '
                f'{len(record["missing"])} | {record["archive"] or "—"} |', notes)

    def test_failed_and_empty_outputs_are_listed_without_empty_archives(self):
        self.document('hello_go-zc310.pdf')
        self.document('ano_go-zc310.pdf', b'')
        manifest = self.bundle()
        zc310 = manifest['libraries']['zc310']
        self.assertEqual(len(zc310['files']), 1)
        self.assertIn('ofd-to-pdf/ano.pdf', zc310['missing'])
        self.assertNotIn('ofd-to-pdf/hello.pdf', zc310['missing'])
        self.assertIn('roundtrip/hello.pdf', zc310['missing'])
        failed = manifest['libraries']['easyofd-python']
        self.assertIsNone(failed['archive'])
        self.assertEqual(len(failed['missing']), len(OFD_SAMPLES))
        self.assertEqual(len(list(self.output.glob('*.zip'))), 1)

    def test_unknown_converter_fails_instead_of_silently_dropping_documents(self):
        self.document('hello_new-converter.pdf')
        with self.assertRaisesRegex(ValueError, 'Unknown converter'):
            self.bundle()
        self.assertFalse(self.output.exists())

    def test_orphan_roundtrip_pdf_is_rejected(self):
        self.document('hello_pdf2ofd-zc310.pdf', directory='output_ofd')
        with self.assertRaisesRegex(ValueError, 'no source OFD'):
            self.bundle()

    def test_empty_run_is_not_published(self):
        with self.assertRaisesRegex(ValueError, 'No nonempty'):
            self.bundle()

    def test_reports_and_input_directories_are_required(self):
        self.document('hello_go-zc310.pdf')
        report = self.root / 'docs' / REPORTS[0]
        report.unlink()
        with self.assertRaisesRegex(ValueError, 'Missing report'):
            self.bundle()
        report.write_text('restored')
        (self.root / 'bin/output_ofd').rmdir()
        with self.assertRaisesRegex(ValueError, 'Missing benchmark output directory'):
            self.bundle()

    def test_existing_output_is_not_overwritten(self):
        self.document('hello_go-zc310.pdf')
        self.bundle()
        manifest = (self.output / 'manifest.json').read_bytes()
        with self.assertRaisesRegex(ValueError, 'Refusing to overwrite'):
            self.bundle()
        self.assertEqual((self.output / 'manifest.json').read_bytes(), manifest)

    def test_tag_with_slash_is_safe_for_asset_filenames(self):
        self.document('hello_go-zc310.pdf')
        manifest = self.bundle('v2026/11')
        self.assertEqual(manifest['tag'], 'v2026/11')
        self.assertTrue((self.output / 'zc310-v2026%2F11.zip').is_file())

    def test_mapping_and_sample_lists_match_benchmark_configuration(self):
        for index, path in enumerate((ROOT / 'main.go', ROOT / 'cmd/pdf2ofd-benchmark/main.go')):
            names = set(re.findall(r'Name:\s+"([^"]+)"', path.read_text()))
            self.assertEqual(names, {converters[index] for converters in LIBRARIES.values()
                                     if converters[index]})
        workflow = (ROOT / '.github/workflows/update-benchmarks.yml').read_text()
        loops = re.findall(r'for name in ([^;]+); do', workflow)
        self.assertEqual(set(loops[0].split()), set(OFD_SAMPLES))
        self.assertEqual(set(loops[1].split()), set(PDF_SAMPLES))

    def publish(self, **statuses):
        fake_bin = self.root / 'fake bin'
        fake_bin.mkdir(exist_ok=True)
        gh = fake_bin / 'gh'
        gh.write_text('''#!/usr/bin/env python3
import json, os, sys
with open(os.environ['GH_CALLS'], 'a') as stream:
    stream.write(json.dumps(sys.argv[1:]) + '\\n')
command = sys.argv[2].upper()
sys.exit(int(os.environ.get('GH_' + command + '_STATUS', '0')))
''')
        gh.chmod(0o755)
        log = self.root / 'gh-calls.jsonl'
        env = dict(os.environ, PATH=f'{fake_bin}:{os.environ["PATH"]}',
                   GH_CALLS=str(log), GH_VIEW_STATUS='1',
                   RELEASE_TAG='v2026.11.16', GITHUB_REPOSITORY='owner/repo')
        env.update({f'GH_{name.upper()}_STATUS': str(value) for name, value in statuses.items()})
        result = subprocess.run(['bash', str(PUBLISH), str(self.output)], env=env,
                                capture_output=True, text=True, timeout=10)
        calls = [json.loads(line) for line in log.read_text().splitlines()] if log.exists() else []
        return result, calls

    def test_publish_uploads_draft_before_publishing(self):
        self.document('hello_go-zc310.pdf')
        self.bundle()
        result, calls = self.publish()
        self.assertEqual(result.returncode, 0, result.stderr)
        self.assertEqual([call[1] for call in calls], ['view', 'create', 'edit'])
        self.assertIn('--draft', calls[1])
        self.assertIn('--verify-tag', calls[1])
        self.assertIn('--draft=false', calls[2])
        for path in self.output.iterdir():
            self.assertIn(str(path), calls[1])

    def test_existing_release_is_never_modified(self):
        self.document('hello_go-zc310.pdf')
        self.bundle()
        result, calls = self.publish(view=0)
        self.assertNotEqual(result.returncode, 0)
        self.assertIn('refusing to overwrite', result.stderr)
        self.assertEqual([call[1] for call in calls], ['view'])

    def test_upload_failure_does_not_publish(self):
        self.document('hello_go-zc310.pdf')
        self.bundle()
        result, calls = self.publish(create=1)
        self.assertNotEqual(result.returncode, 0)
        self.assertEqual([call[1] for call in calls], ['view', 'create'])

    def test_corrupted_assets_fail_before_github_access(self):
        self.document('hello_go-zc310.pdf')
        self.bundle()
        (self.output / 'zc310-v2026.11.16.zip').write_bytes(b'corrupted')
        result, calls = self.publish()
        self.assertNotEqual(result.returncode, 0)
        self.assertEqual(calls, [])


if __name__ == '__main__':
    unittest.main()
