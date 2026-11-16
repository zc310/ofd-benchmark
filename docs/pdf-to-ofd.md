# PDF → OFD 对比测试

> [OFD 库汇总](ofd-libraries.md)：各 OFD 库能力 / 许可证对比表
>
> [OFD → PDF 对比测试](ofd-to-pdf.md)：反方向 OFD→PDF 转换对比



## 测试库

| 转换器         | 语言   | 库                                                                          | 版本                    | 说明                                                                                                         |
|----------------|--------|-----------------------------------------------------------------------------|-------------------------|--------------------------------------------------------------------------------------------------------------|
| go-zc310       | Go     | [zc310/ofd](https://github.com/zc310/ofd)                                   | v0.1.10-0.20261114154015-087ae6006559 | 基于 pdfcpu 的 PDF→OFD 导入器（路径/文本/图片/注解/渐变/图案/网格/大纲/元数据），文本保留为可提取 TextObject |
| go-ofdgo       | Go     | [xiaoqidun/ofdgo](https://github.com/xiaoqidun/ofdgo)（ConvertPDF + pdfgo） | v0.0.0-20261111160922-3edf35199ec7 | 基于 pdfgo 的 PDF→OFD 导入器（路径/文本/图片/注解）                                                          |
| rust-easyofd | Rust | [easyofd-rust](https://github.com/easy-4-rust/easyofd-rust) | 0.1.4 | 默认 PDF→OFD 转换选项；成功生成文件不代表回环外观或文本通过，见下方实测 |
| java-ofdrw     | Java   | [ofdrw](https://github.com/ofdrw/ofdrw)（ofdrw-converter PDFConverter）     | 2.5.1 | 文本转为矢量轮廓，保留原始外观                                                                               |
| python-pdf2ofd | Python | [pdf2ofd](https://github.com/wanglrebe/pdf2ofd)                             | 0.1.0                   | 基于 PyMuPDF 渲染，文本栅格化                                                                                |
| node-jsofd     | JS     | [Hufe921/jsOFD](https://github.com/Hufe921/jsOFD)（pdfToOfd + pdfjs-dist）  | 1.0.1                   | 基于 pdfjs 重放绘制指令流（路径/文本/图片），不转换注解/渐变/图案/Type3/剪裁                                 |

> `python-pdf2ofd` 安装：在独立 venv 安装 `python/requirements-pdf2ofd.txt`，通过 `PDF2OFD_PYTHON` 指定该环境的 Python（完整命令见 README）。其 PyMuPDF 1.24.14 与像素评测环境的 1.26.7 隔离。
>
> `node-jsofd` 安装：`npm install @hufe921/jsofd pdfjs-dist`（Node <22 需补 `process.getBuiltinModule`，见 `node/pdf2ofd_jsofd.mjs`）

## 测试文件

直接使用数科版式阅读器导出的参考 PDF（即 OFD→PDF 测试的对比基准），保证与 OFD→PDF 测试同一组文档。

| 文件               | 大小 | 页数    | 说明                                           |
|--------------------|------|---------|------------------------------------------------|
| hello.pdf          | 3.5K | 1 页    | 简单文本                                       |
| ano.pdf            | 904K | 3 页    | 发票                                           |
| 999.pdf            | 80K  | 5 页    | 简单文档                                       |
| intro.pdf          | 6.7M | 42 页   | 介绍文档                                       |
| zsbk.pdf           | 1.6M | 2 页    | 签章文件                                       |
| 1000-pages.pdf     | 2.4M | 1000 页 | 多页文档                                       |
| GBT_33190-2016.pdf | 4.8M | 132 页  | GB/T 33190-2016 标准正文（字体全内嵌、图表多） |

## 测试环境

<!-- benchmark-environment:start -->

- 系统：Ubuntu 24.04.5 LTS (x86_64)
- CPU：AMD EPYC 9V74 80-Core Processor（4 逻辑核心）
- 内存：15.6 GiB
<!-- benchmark-environment:end -->
- 系统已安装 `apt-get install fonts-wqy-zenhei fonts-noto-cjk`
- 测试前运行 `bash tools/install-fonts.sh`，下载 `doc-foundertype.zip`、`doc-zhongyi-etc.zip`、`latin-fonts.zip` 并安装到用户目录 `~/.local/share/fonts/hanzi-fonts/`，随后刷新字体缓存。字体下载链接与依赖安装命令见 [README：快速开始](../README.md#快速开始)；GitHub Actions 已加入此准备步骤。

> `rust-easyofd` 为 2026-11-16 本地补测（Ubuntu 26.04 LTS、AMD Ryzen 9 5950X，PyMuPDF 1.26.7 / NumPy 2.3.5）；其他行保留原结果。其耗时与上述 CI 环境不是同机数据，速度排名及胜出次数暂不作为性能结论；下次完整 CI 将统一重测。回环统一使用 go-zc310，不改用 easyofd 自身回环。

## 转换速度

<!-- benchmark-table:pdf-speed:start -->

| 转换器 | hello.pdf | ano.pdf | 999.pdf | intro.pdf | zsbk.pdf | 1000-pages.pdf | GBT_33190-2016.pdf | 胜出次数 |
|---|---:|---:|---:|---:|---:|---:|---:|---:|
| go-zc310 | 30ms | 93ms | 300ms | 440ms | 1.23s | 1.01s | 1.57s | 0 |
| go-ofdgo | 53ms | 356ms | 479ms | 978ms | 1.71s | 16.57s | 4.83s | 0 |
| rust-easyofd | **2ms** | **4ms** | **3ms** | **105ms** | **22ms** | **250ms** | **75ms** | 7 |
| java-ofdrw | 1.06s | 1.86s | 1.95s | 8.23s | 4.97s | 4.69s | 71.25s | 0 |
| python-pdf2ofd | 120ms | 965ms | 621ms | 2.85s | 1.12s | 61.18s | 21.59s | 0 |
| node-jsofd | 241ms | 452ms | 3.30s | 9.74s | 5.17s | 2.65s | 7.92s | 0 |
<!-- benchmark-table:pdf-speed:end -->

## OFD 文件大小

<!-- benchmark-table:pdf-ofd-size:start -->

| 转换器 | hello.pdf | ano.pdf | 999.pdf | intro.pdf | zsbk.pdf | 1000-pages.pdf | GBT_33190-2016.pdf | 胜出次数 |
|---|---:|---:|---:|---:|---:|---:|---:|---:|
| go-zc310 | 1.7K | 405.4K | 62.3K | 7.3M | 1.6M | 4.8M | 7.1M | 0 |
| go-ofdgo | 6.2K | 709.5K | 117.5K | 9.2M | 3.0M | 60.7M | 17.1M | 0 |
| rust-easyofd | **1.2K** | **1.8K** | **3.6K** | **5.2M** | **918.8K** | **304.6K** | **2.0M** | 7 |
| java-ofdrw | 5.6K | 601.3K | 817.6K | 34.2M | 20.4M | 3.5M | 30.5M | 0 |
| python-pdf2ofd | 4.6K | 1.3M | 733.7K | 8.9M | 1.6M | 86.5M | 36.8M | 0 |
| node-jsofd | 3.0K | 2.2M | 590.6K | 34.8M | 12.9M | 31.8M | 34.3M | 0 |
<!-- benchmark-table:pdf-ofd-size:end -->

## 文件大小对比（输出 OFD / 原 PDF）

越小表示压缩率越高，`1.00` 表示与源 PDF 等大；**加粗**为最小。

<!-- benchmark-table:pdf-size-ratio:start -->

| 转换器 | hello.pdf | ano.pdf | 999.pdf | intro.pdf | zsbk.pdf | 1000-pages.pdf | GBT_33190-2016.pdf |
|---|---:|---:|---:|---:|---:|---:|---:|
| go-zc310 | 0.48 | 0.45 | 0.78 | 1.09 | 1.02 | 2.54 | 1.46 |
| go-ofdgo | 1.76 | 0.78 | 1.47 | 1.38 | 1.91 | 31.98 | 3.53 |
| rust-easyofd | **0.33** | **0.00** | **0.04** | **0.78** | **0.57** | **0.16** | **0.41** |
| java-ofdrw | 1.58 | 0.67 | 10.23 | 5.10 | 13.00 | 1.83 | 6.32 |
| python-pdf2ofd | 1.32 | 1.43 | 9.18 | 1.33 | 1.04 | 45.55 | 7.62 |
| node-jsofd | 0.84 | 2.47 | 7.39 | 5.18 | 8.18 | 16.72 | 7.10 |
<!-- benchmark-table:pdf-size-ratio:end -->

## 还原度对比

将每个转换器输出的 OFD 再经 go-zc310 转回 PDF（回环），与源 PDF 做像素对比（96 DPI）与文本对比。

像素脚本对 `1000-pages` 仅取前 3 页（仍校验总页数一致）；文本比较仍处理完整文档。除本次补测的 `rust-easyofd` 外，现有回环结果尚未按采样规则重算，下次运行脚本更新。

`rust-easyofd 0.1.4` 本次 7/7 成功生成 OFD，但 hello、ano、intro、1000-pages、GBT_33190-2016 回环触发近乎空白页检查，像素记 `FAILED`；999、zsbk 分别为 96.26%、95.85%，不代表内容完整。7 个样本的文本相似度均为 0.00%。回环日志还出现非数字 StID（如 `t_0_0`、`SimSun`）解析及部分 BMP 图像解码错误；这些是当前输出/回环链路的限制，未通过修补输出规避。

**像素相似度**（100% 为完全相同）：

<!-- benchmark-table:pdf-pixel-similarity:start -->

| 转换器 | hello.pdf | ano.pdf | 999.pdf | intro.pdf | zsbk.pdf | 1000-pages.pdf | GBT_33190-2016.pdf | 胜出次数 |
|---|---:|---:|---:|---:|---:|---:|---:|---:|
| go-zc310 | 99.63% | **99.77%** | 98.43% | **99.91%** | **99.56%** | **99.84%** | **99.98%** | 5 |
| go-ofdgo | 99.63% | 99.38% | 98.34% | 98.01% | 99.18% | 99.63% | 99.82% | 0 |
| rust-easyofd | FAILED | FAILED | 96.26% | FAILED | 95.85% | FAILED | FAILED | 0 |
| java-ofdrw | **99.85%** | 98.72% | 98.75% | 99.49% | 99.07% | 97.99% | 99.66% | 1 |
| python-pdf2ofd | 99.72% | 98.34% | **98.94%** | 99.33% | 97.62% | 98.39% | 99.12% | 1 |
| node-jsofd | 99.37% | 96.77% | 98.14% | 98.88% | 98.82% | 99.50% | 99.94% | 0 |
<!-- benchmark-table:pdf-pixel-similarity:end -->

**文本提取能力**（从转换后 OFD 提取字符数，参考 PDF 提取字符数见括号）：

<!-- benchmark-table:pdf-text-extraction:start -->

| 转换器 | hello.pdf(22) | ano.pdf(7091) | 999.pdf(3934) | intro.pdf(5543) | zsbk.pdf(1150) | 1000-pages.pdf(263999) | GBT_33190-2016.pdf(107779) |
|---|---:|---:|---:|---:|---:|---:|---:|
| go-zc310 | 22 | 9222 | 9737 | 9314 | 2740 | 326022 | 170697 |
| go-ofdgo | 8 | 14127 | 10009 | 10920 | 3373 | 179998 | 234709 |
| rust-easyofd | 0 | 0 | 242 | 0 | 10 | 0 | 0 |
| java-ofdrw | 0 | 7049 | 0 | 1598 | 3 | 321998 | 167 |
| python-pdf2ofd | 0 | 0 | 0 | 0 | 0 | 0 | 0 |
| node-jsofd | 25 | 10812 | 9941 | 10222 | 2796 | 326448 | 225807 |
<!-- benchmark-table:pdf-text-extraction:end -->

**文本相似度**（4-gram Jaccard，100% 为与源 PDF 提取文本完全一致）：

<!-- benchmark-table:pdf-text-similarity:start -->

| 转换器 | hello.pdf | ano.pdf | 999.pdf | intro.pdf | zsbk.pdf | 1000-pages.pdf | GBT_33190-2016.pdf |
|---|---:|---:|---:|---:|---:|---:|---:|
| go-zc310 | 100.00% | 39.16% | 64.41% | 15.47% | 10.43% | 98.54% | 41.99% |
| go-ofdgo | 26.32% | 23.16% | 44.56% | 13.71% | 4.63% | 2.63% | 1.08% |
| rust-easyofd | 0.00% | 0.00% | 0.00% | 0.00% | 0.00% | 0.00% | 0.00% |
| java-ofdrw | 0.00% | 89.55% | 0.00% | 4.60% | 0.00% | 98.86% | 0.00% |
| python-pdf2ofd | 0.00% | 0.00% | 0.00% | 0.00% | 0.00% | 0.00% | 0.00% |
| node-jsofd | 100.00% | 42.55% | 18.27% | 15.19% | 6.94% | 30.25% | 1.08% |
<!-- benchmark-table:pdf-text-similarity:end -->


## 复现

```bash
# 编译 zc310 ofd-converter（PDF→OFD / OFD→PDF / OFD→TXT）
cd /home/go/workspace/ofd
go build -o ../ofd-benchmark/bin/ofd-converter ./cmd/ofd-converter

cd ../ofd-benchmark
# 编译基准测试驱动（上级存在 go.work 时需 GOWORK=off）
GOWORK=off go build -o bin/pdf2ofd-benchmark ./cmd/pdf2ofd-benchmark

# 编译 ofdgo PDF→OFD 转换器
GOWORK=off go build -o bin/ofdgo-pdf2ofd ./cmd/ofdgo-pdf2ofd

# 编译 easyofd-rust 双向转换器（固定 0.1.4）
cargo build --release --locked --manifest-path rust/easyofd-convert/Cargo.toml --target-dir rust/target
cp rust/target/release/easyofd-ofd2pdf rust/target/release/easyofd-pdf2ofd bin/

# 运行全部转换器（单个文件）
./bin/pdf2ofd-benchmark testdata/suwell/hello.pdf

# 单独测试 easyofd-rust
./bin/pdf2ofd-benchmark testdata/suwell/hello.pdf easyofd-rust

# 批量运行所有测试文件
for f in testdata/suwell/*.pdf testdata/standards/*.pdf; do
  ./bin/pdf2ofd-benchmark "$f"
done

# 还原度/文本对比（输出位于 bin/output_ofd/）
python3 tools/compare_pdf2ofd.py 96
```
