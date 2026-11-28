#!/usr/bin/env python3
"""Prepare Markdown and SVG resources for the GitHub Pages Jekyll build."""

import json
import os
from pathlib import Path
import re
import shutil
import sys


def main():
    source, output = map(Path, sys.argv[1:])
    documents = [(source.parent / "README.md", Path("project.md"))]
    documents += [(path, path.relative_to(source)) for path in sorted(source.rglob("*.md"))]
    output.mkdir(parents=True, exist_ok=True)
    (output / "_config.yml").write_text(
        "title: OFD 转换与渲染对比测试\n"
        "description: OFD 库能力与转换性能对比\n"
        "theme: jekyll-theme-cayman\n",
        encoding="utf-8",
    )

    for svg in source.rglob("*.svg"):
        destination = output / svg.relative_to(source)
        destination.parent.mkdir(parents=True, exist_ok=True)
        shutil.copy2(svg, destination)

    links = []
    for document, relative in documents:
        destination = output / relative
        destination.parent.mkdir(parents=True, exist_ok=True)
        text = document.read_text(encoding="utf-8")
        heading = re.search(r"^#\s+(.+)$", text, re.MULTILINE)
        title = heading.group(1).strip() if heading else document.stem
        # Generator markers can interrupt Jekyll's table parsing.
        text = re.sub(r"^<!-- benchmark-[^\n]* -->[ \t]*$", "", text, flags=re.MULTILINE)
        text = re.sub(r"(\]\((?![a-zA-Z][a-zA-Z0-9+.-]*://)[^)\s]*?)\.md(?=[#)])", r"\1.html", text)
        if document == source.parent / "README.md":
            text = text.replace("](docs/", "](")
        else:
            readme = os.path.relpath(output / "project.html", destination.parent)
            text = text.replace("../README.html", readme)
        destination.write_text(
            f"---\nlayout: default\ntitle: {json.dumps(title, ensure_ascii=False)}\n---\n{text}",
            encoding="utf-8",
        )
        links.append(f"- [{title}]({relative.with_suffix('.html').as_posix()})")

    (output / "index.md").write_text(
        "---\nlayout: default\ntitle: 文档首页\n---\n"
        "# OFD Benchmark 文档\n\n" + "\n".join(links) + "\n",
        encoding="utf-8",
    )


if __name__ == "__main__":
    main()
