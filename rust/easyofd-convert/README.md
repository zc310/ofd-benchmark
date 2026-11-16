# easyofd-rust 双向评测

使用 crates.io `easyofd = "=0.1.4"` 和提交的 Cargo.lock，保持默认转换选项。
与 Python easyofd 是不同的实现；不修补上游输出或放宽质量检查。

在项目根目录执行：

```bash
cargo build --release --locked --manifest-path rust/easyofd-convert/Cargo.toml --target-dir rust/target
cp rust/target/release/easyofd-ofd2pdf rust/target/release/easyofd-pdf2ofd bin/
GOWORK=off go build -o bin/ofd-benchmark .
GOWORK=off go build -o bin/pdf2ofd-benchmark ./cmd/pdf2ofd-benchmark
bin/ofd-benchmark testdata/ofd-creator/hello.ofd easyofd-rust
bin/pdf2ofd-benchmark testdata/suwell/hello.pdf easyofd-rust
```

输出名称为 `bin/output/<name>_rust-easyofd.pdf` 和
`bin/output_ofd/<name>_pdf2ofd-easyofd-rust.ofd`。
像素/文本脚本自动识别这两个名称，并将结果写回 `rust-easyofd` 行。
PDF→OFD 使用项目统一的 go-zc310 回环，而不是 easyofd 自身回环：

```bash
bash tools/roundtrip-pdf2ofd.sh bin/ofd-converter bin/output_ofd
python3 tools/compare_pdf.py 96
python3 tools/compare_text.py
python3 tools/compare_pdf2ofd.py 96
GOWORK=off go run ./tools/gen-benchmarks
```

成功生成文件仅表示转换进程完成；近乎空白页、页数/尺寸不符及渲染警告等仍会使视觉指标记为 `FAILED`。
