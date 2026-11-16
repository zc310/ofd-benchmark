"""Package successful benchmark documents by library using only the standard library."""

import argparse
import hashlib
import json
import os
from pathlib import Path
import shutil
from urllib.parse import quote
import zipfile


# Output suffixes come from main.go and cmd/pdf2ofd-benchmark/main.go.
# Keep bindings/language implementations separate even when they share a name.
LIBRARIES = {
    'zc310': ('go-zc310', 'pdf2ofd-zc310'),
    'ofdgo': ('go-ofdgo', 'pdf2ofd-ofdgo'),
    'ofdrw-java': ('java-ofdrw', 'pdf2ofd-ofdrw'),
    'ofdrw-rs': ('go-ofdrw-rs', None),
    'ofdbox-folio': ('rust-ofdbox-folio', None),
    'sofd': ('rust-sofd', None),
    'easyofd-rust': ('rust-easyofd', 'pdf2ofd-easyofd-rust'),
    'easyofd-python': ('python-easyofd', None),
    'ofd-to-pdf-node': ('node-ofd2pdf', None),
    'pdf2ofd-python': (None, 'python-pdf2ofd'),
    'jsofd': (None, 'node-jsofd'),
}
LIBRARY_URLS = {
    'zc310': 'https://github.com/zc310/ofd',
    'ofdgo': 'https://github.com/xiaoqidun/ofdgo',
    'ofdrw-java': 'https://github.com/ofdrw/ofdrw',
    'ofdrw-rs': 'https://github.com/lihongjie0209/ofdrw-rs',
    'ofdbox-folio': 'https://github.com/ofdbox/ofdbox-folio',
    'sofd': 'https://github.com/emfox/sOFD',
    'easyofd-rust': 'https://github.com/easy-4-rust/easyofd-rust',
    'easyofd-python': 'https://pypi.org/project/easyofd/',
    'ofd-to-pdf-node': 'https://www.npmjs.com/package/@miconvert/ofd-to-pdf',
    'pdf2ofd-python': 'https://github.com/wanglrebe/pdf2ofd',
    'jsofd': 'https://github.com/Hufe921/jsOFD',
}
OFD_SAMPLES = ('hello', 'ano', 'intro', '1000-pages', '999', 'zsbk')
PDF_SAMPLES = (*OFD_SAMPLES, 'GBT_33190-2016')
REPORTS = ('benchmark-results.yaml', 'ofd-to-pdf.md', 'pdf-to-ofd.md')


def sha256(path):
    with path.open('rb') as stream:
        return hashlib.file_digest(stream, 'sha256').hexdigest()


def package(root, output, tag, commit='', run_url=''):
    root, output = Path(root), Path(output)
    if not tag.startswith('v'):
        raise ValueError('Release tags must start with v')
    if output.exists() and any(output.iterdir()):
        raise ValueError(f'Refusing to overwrite nonempty output directory: {output}')
    for name in REPORTS:
        if not (root / 'docs' / name).is_file():
            raise ValueError(f'Missing report: {name}')

    records = {library: {'archive': None, 'files': [], 'missing': []}
               for library in LIBRARIES}
    sources = {}
    for directory, extensions, index in (
        ('output', ('.pdf',), 0),
        ('output_ofd', ('.ofd', '.pdf'), 1),
    ):
        source_dir = root / 'bin' / directory
        if not source_dir.is_dir():
            raise ValueError(f'Missing benchmark output directory: {source_dir}')
        for path in sorted(source_dir.iterdir()):
            if path.suffix not in extensions or not path.is_file():
                continue
            for library, converters in LIBRARIES.items():
                converter = converters[index]
                suffix = f'_{converter}'
                if converter and path.stem.endswith(suffix):
                    sample = path.stem[:-len(suffix)]
                    break
            else:
                # Never silently drop documents when a new converter is added.
                raise ValueError(f'Unknown converter output: {path.name}')
            if path.stat().st_size == 0:
                continue
            if index == 1 and path.suffix == '.pdf':
                ofd = path.with_suffix('.ofd')
                if not ofd.is_file() or ofd.stat().st_size == 0:
                    raise ValueError(f'Round-trip PDF has no source OFD: {path.name}')
                direction = 'roundtrip'
            else:
                direction = 'ofd-to-pdf' if index == 0 else 'pdf-to-ofd'
            destination = f'{library}/{direction}/{sample}{path.suffix}'
            records[library]['files'].append({
                'source': path.relative_to(root).as_posix(),
                'path': destination,
                'size': path.stat().st_size,
            })
            sources[destination] = path

    if not sources:
        raise ValueError('No nonempty PDF/OFD documents to release')

    for library, converters in LIBRARIES.items():
        present = {item['path'] for item in records[library]['files']}
        for converter, samples, direction, extension in (
            (converters[0], OFD_SAMPLES, 'ofd-to-pdf', '.pdf'),
            (converters[1], PDF_SAMPLES, 'pdf-to-ofd', '.ofd'),
            (converters[1], PDF_SAMPLES, 'roundtrip', '.pdf'),
        ):
            if converter:
                records[library]['missing'].extend(
                    f'{direction}/{sample}{extension}' for sample in samples
                    if f'{library}/{direction}/{sample}{extension}' not in present
                )

    output.mkdir(parents=True, exist_ok=True)
    for library, record in records.items():
        if not record['files']:
            continue  # Entirely failed libraries remain visible in the manifest.
        archive = f'{library}-{quote(tag, safe="")}.zip'
        with zipfile.ZipFile(output / archive, 'x', compression=zipfile.ZIP_DEFLATED,
                             compresslevel=6) as bundle:
            for item in record['files']:
                bundle.write(sources[item['path']], item['path'])
        record['archive'] = archive

    for name in REPORTS:
        shutil.copyfile(root / 'docs' / name, output / name)
    manifest = {'tag': tag, 'commit': commit, 'run_url': run_url, 'libraries': records}
    (output / 'manifest.json').write_text(
        json.dumps(manifest, ensure_ascii=False, indent=2) + '\n', encoding='utf-8')
    notes = [f'# Benchmark {tag}', '', f'Commit: `{commit}`', '']
    if run_url:
        notes.extend([f'[Workflow run]({run_url})', ''])
    notes.extend([
        'Each ZIP contains one library’s successful outputs: `ofd-to-pdf/`, '
        '`pdf-to-ofd/`, and `roundtrip/` (PDF rendered by zc310).', '',
        'Missing or empty outputs are omitted, not replaced by previous results. '
        'See `manifest.json` for file lists and missing samples. '
        'Libraries with no documents have no ZIP.', '',
        'Dependencies are resolved at run time; actual tracked versions and runner '
        'details are recorded in `benchmark-results.yaml` and the attached reports.', '',
        '| Library | Documents | Missing outputs | Archive |',
        '|---|---:|---:|---|',
    ])
    for library, record in records.items():
        notes.append(f'| [{library}]({LIBRARY_URLS[library]}) | {len(record["files"])} | '
                     f'{len(record["missing"])} | {record["archive"] or "—"} |')
    (output / 'RELEASE_NOTES.md').write_text('\n'.join(notes) + '\n', encoding='utf-8')
    checksums = [f'{sha256(path)}  {path.name}\n' for path in sorted(output.iterdir())]
    (output / 'SHA256SUMS').write_text(''.join(checksums), encoding='utf-8')
    return manifest


def main():
    parser = argparse.ArgumentParser(description=__doc__)
    parser.add_argument('--root', type=Path, default=Path(__file__).resolve().parent.parent)
    parser.add_argument('--output', type=Path, required=True)
    parser.add_argument('--tag', required=True)
    args = parser.parse_args()
    try:
        package(args.root, args.output, args.tag, os.environ.get('GITHUB_SHA', ''),
                os.environ.get('RELEASE_RUN_URL', ''))
    except (OSError, ValueError) as error:
        parser.exit(1, f'{error}\n')


if __name__ == '__main__':
    main()
