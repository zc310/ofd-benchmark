"""Bridge comparison scripts to the shared Go YAML updater."""

import json
import os
import subprocess


OFD_CONVERTER_ROWS = {
    "go-zc310": "go-zc310-ofd",
    "go-ofdgo": "go-xiaoqidun-ofdgo",
    "go-ofdrw-rs": "go-ofdrw-rs",
    "rust-ofdbox-folio": "rust-ofdbox-folio",
    "rust-sofd": "rust-sOFD",
    "python-easyofd": "python-easyofd",
    "java-ofdrw": "java-ofdrw",
    "node-ofd2pdf": "node-ofd-to-pdf",
}

PDF_CONVERTER_ROWS = {
    "pdf2ofd-zc310": "go-zc310-ofd",
    "pdf2ofd-ofdgo": "go-xiaoqidun-ofdgo",
    "pdf2ofd-ofdrw": "java-ofdrw",
    "python-pdf2ofd": "python-pdf2ofd",
    "node-jsofd": "node-jsofd",
}


def format_percentage(value):
    """Scores already use the 0-100 scale; keep four fractional digits."""
    return f'{value:.4f}%' if value is not None else 'FAILED'


def append(updates, report, table, input_name, converter, value):
    updates.append({
        "report": report,
        "table": table,
        "input": input_name,
        "converter": converter,
        "value": value,
    })


def write_updates(base, updates):
    if not updates:
        return
    writer = os.environ.get("BENCHMARK_RESULTS_WRITER")
    if writer:
        command = [writer]
    else:
        command = ["go", "run", "./tools/update-benchmark-results"]
    command += ["-yaml", os.path.join(base, "docs", "benchmark-results.yaml"), "-input", "-"]
    env = os.environ.copy()
    env["GOWORK"] = "off"
    subprocess.run(
        command,
        cwd=base,
        env=env,
        input=json.dumps(updates),
        text=True,
        check=True,
    )
