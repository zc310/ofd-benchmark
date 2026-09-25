# OFD <-> PDF 转换对比测试

> [OFD 库汇总](docs/ofd-libraries.md) —— 各 OFD 库能力 / 许可证对比表（预览、解析、发票、→ 图片/PDF/TXT/SVG/Md/HTML、PDF→OFD、生成、签章、修改）

本项目对比多种开源 OFD 库的转换能力，覆盖两个方向：

| 方向 | 文档 | 对比内容 |
|------|------|----------|
| OFD → PDF | [docs/ofd-to-pdf.md](docs/ofd-to-pdf.md) | 8 种转换器（Go×2 / Rust / Python×3 / Java / Node）的速度、文件大小、文本提取与数科版式还原度 |
| PDF → OFD | [docs/pdf-to-ofd.md](docs/pdf-to-ofd.md) | 5 种转换器（zc310 / ofdgo / easyofd-rust / ofdrw / pdf2ofd）的速度、文件大小与回环还原度 |

基准参考输出来自**数科版式阅读器**（`test/testdata/*.pdf`），保证同一组文档双向可测。

## 快速开始

```bash
# OFD → PDF 基准（详见 docs/ofd-to-pdf.md）
go build -o bin/ofd-benchmark .
./bin/ofd-benchmark test/testdata/hello.ofd

# PDF → OFD 基准（详见 docs/pdf-to-ofd.md）
cd /home/go/workspace/ofd && go build -o ../ofd-benchmark/bin/ofd-converter ./cmd/ofd-converter
cd ../ofd-benchmark && go build -o bin/pdf2ofd-benchmark ./cmd/pdf2ofd-benchmark
./bin/pdf2ofd-benchmark test/testdata/hello.pdf
```

转换结果分别保存在 `bin/output/`（OFD→PDF）与 `bin/output_ofd/`（PDF→OFD）。

## 项目结构

```
ofd-benchmark/
├── bin/
│   ├── ofd-benchmark       # OFD→PDF 基准测试主程序
│   ├── ofdgo-convert       # Go xiaoqidun/ofdgo OFD→PDF 转换器
│   ├── ofdgo-pdf2ofd       # Go xiaoqidun/ofdgo PDF→OFD 转换器
│   ├── zc310-convert       # Go zc310/ofd OFD→PDF 转换器
│   ├── ofd-converter       # Go zc310/ofd 通用转换器（PDF→OFD、OFD→PDF/TXT）
│   ├── pdf2ofd-benchmark   # PDF→OFD 基准测试主程序
│   ├── easyofd             # Rust 二进制
│   ├── output/             # OFD→PDF 转换结果
│   └── output_ofd/         # PDF→OFD 转换结果（含回环 PDF / 提取文本）
├── cmd/
│   ├── ofdgo-convert/
│   ├── ofdgo-pdf2ofd/
│   ├── zc310-convert/
│   └── pdf2ofd-benchmark/  # PDF→OFD 基准驱动
├── python/
│   ├── ofd2pdf.py          # Python 转换器（easyofd + ofd2img）
│   ├── ofd2pdf_ofdreader.py
│   └── pdf2ofd_convert.py  # Python pdf2ofd 转换器
├── java/
│   ├── Ofd2Pdf.java        # Java ofdrw OFD→PDF 转换器
│   ├── Pdf2Ofd.java        # Java ofdrw PDF→OFD 转换器（PDFConverter）
│   ├── ofd2pdf_java.sh
│   ├── pdf2ofd_java.sh
│   └── lib/                # JAR 依赖
├── node/
│   ├── ofd2pdf.mjs         # Node.js 转换器（@miconvert/ofd-to-pdf）
│   └── ofd2pdf_node.sh
├── rust/
│   ├── ofd2pdf/            # Rust OFD→PDF 转换器（easyofd）
│   └── pdf2ofd/            # Rust PDF→OFD 转换器（easyofd）
├── tools/
│   ├── compare_pdf.py      # OFD→PDF 与数科参考 PDF 像素对比
│   ├── compare_text.py     # OFD→PDF 文本相似度对比
│   └── compare_pdf2ofd.py  # PDF→OFD 回环还原度/文本对比
├── docs/
│   ├── ofd-libraries.md    # OFD 库能力 / 许可证对比表
│   ├── ofd-to-pdf.md       # OFD→PDF 基准测试详细结果
│   └── pdf-to-ofd.md       # PDF→OFD 基准测试详细结果
├── test/testdata/          # OFD/PDF 测试文件
├── main.go                 # OFD→PDF 基准主程序（单文件输入）
└── go.mod                  # 含 replace 到本地 zc310/ofd 与 zc310/canvas
```