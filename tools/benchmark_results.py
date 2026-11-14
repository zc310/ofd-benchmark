"""Bridge comparison scripts to the shared Go YAML updater."""

import json
import os
import subprocess


OFD_CONVERTER_ROWS = {
    "go-zc310": "go-zc310",
    "go-ofdgo": "go-ofdgo",
    "python-easyofd": "easyofd",
    "python-ofd2pdf": "ofd2pdf",
    "python-ofdreader": "python-ofdreader",
    "java-ofdrw": "ofdrw",
    "node-ofd2pdf": "@miconvert/ofd-to-pdf",
}

PDF_CONVERTER_ROWS = {
    "pdf2ofd-zc310": "go-zc310",
    "pdf2ofd-ofdgo": "go-ofdgo",
    "pdf2ofd-ofdrw": "java-ofdrw",
    "python-pdf2ofd": "python-pdf2ofd",
    "node-jsofd": "node-jsofd",
}


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
