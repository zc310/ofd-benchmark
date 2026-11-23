# PDF → OFD 对比测试

> [OFD 库汇总](ofd-libraries.md)：各 OFD 库能力 / 许可证对比表
>
> [OFD → PDF 对比测试](ofd-to-pdf.md)：反方向 OFD→PDF 转换对比



## 测试库

| 转换器         | 语言   | 库                                                                          | 版本                    | 说明                                                                                                         |
|----------------|--------|-----------------------------------------------------------------------------|-------------------------|--------------------------------------------------------------------------------------------------------------|
| go-zc310       | Go     | [zc310/ofd](https://github.com/zc310/ofd)                                   | v0.2.2 | 基于 pdfcpu 的 PDF→OFD 导入器（路径/文本/图片/注解/渐变/图案/网格/大纲/元数据），文本保留为可提取 TextObject |
| go-ofdgo       | Go     | [xiaoqidun/ofdgo](https://github.com/xiaoqidun/ofdgo)（ConvertPDF + pdfgo） | v0.0.0-20261122015454-e70fdfb4451c | 基于 pdfgo 的 PDF→OFD 导入器（路径/文本/图片/注解）                                                          |
| java-ofdrw     | Java   | [ofdrw](https://github.com/ofdrw/ofdrw)（ofdrw-converter PDFConverter）     | 2.5.2 | 文本转为矢量轮廓，保留原始外观                                                                               |
| python-pdf2ofd | Python | [pdf2ofd](https://github.com/wanglrebe/pdf2ofd)                             | 0.1.0                   | 基于 PyMuPDF 渲染，文本栅格化                                                                                |
| node-jsofd     | JS     | [Hufe921/jsOFD](https://github.com/Hufe921/jsOFD)（pdfToOfd + pdfjs-dist）  | 1.0.1                   | 基于 pdfjs 重放绘制指令流（路径/文本/图片），不转换注解/渐变/图案/Type3/剪裁                                 |

> `python-pdf2ofd` 安装：在独立 venv 安装 `python/requirements-pdf2ofd.txt`，通过 `PDF2OFD_PYTHON` 指定该环境的 Python（完整命令见 README）。其 PyMuPDF 1.24.14 与像素评测环境的 1.26.7 隔离。
>
> `node-jsofd` 安装：`npm install @hufe921/jsofd pdfjs-dist`（Node <22 需补 `process.getBuiltinModule`，见 `node/pdf2ofd_jsofd.mjs`）

## 测试文件

直接使用数科版式阅读器导出的参考 PDF（即 OFD→PDF 测试的对比基准），保证与 OFD→PDF 测试同一组文档。

| 文件               | 大小 | 页数    | 说明                                           |
|--------------------|------|---------|------------------------------------------------|
| hello.pdf          | 3.50K | 1 页    | 简单文本                                       |
| ano.pdf            | 904.00K | 3 页    | 发票                                           |
| 999.pdf            | 80.00K  | 5 页    | 简单文档                                       |
| intro.pdf          | 6.70M | 42 页   | 介绍文档                                       |
| zsbk.pdf           | 1.60M | 2 页    | 签章文件                                       |
| 1000-pages.pdf     | 2.40M | 1000 页 | 多页文档                                       |
| GBT_33190-2016.pdf | 4.80M | 132 页  | GB/T 33190-2016 标准正文（字体全内嵌、图表多） |

## 测试环境

<!-- benchmark-environment:start -->

- 系统：Ubuntu 24.04.5 LTS (x86_64)
- CPU：INTEL(R) XEON(R) PLATINUM 8573C（4 逻辑核心）
- 内存：15.6 GiB
<!-- benchmark-environment:end -->
- 系统已安装 `apt-get install fonts-wqy-zenhei fonts-noto-cjk`
- 测试前运行 `bash tools/install-fonts.sh`，下载 `doc-foundertype.zip`、`doc-zhongyi-etc.zip`、`latin-fonts.zip` 并安装到用户目录 `~/.local/share/fonts/hanzi-fonts/`，随后刷新字体缓存。字体下载链接与依赖安装命令见 [README：快速开始](../README.md#快速开始)；GitHub Actions 已加入此准备步骤。

> 所有转换器均由 GitHub Actions 在上述环境统一评测。回环统一使用 go-zc310，不改用各库自身回环；速度排名及胜出次数应结合回环外观与文本质量解读，成功生成文件不代表通过质量检查。

## 转换速度

<!-- benchmark-table:pdf-speed:start -->

| 转换器 | hello.pdf | ano.pdf | 999.pdf | intro.pdf | zsbk.pdf | 1000-pages.pdf | GBT_33190-2016.pdf | 胜出次数 |
|---|---:|---:|---:|---:|---:|---:|---:|---:|
| go-zc310 | **22ms** | **97ms** | **298ms** | **497ms** | 1.29s | **1.04s** | **1.84s** | 6 |
| go-ofdgo | 56ms | FAILED | FAILED | 566ms | FAILED | 8.83s | 2.69s | 0 |
| java-ofdrw | 1.08s | 1.92s | 2.32s | 9.98s | 5.96s | 5.89s | 104.59s | 0 |
| python-pdf2ofd | 149ms | 1.10s | 705ms | 3.28s | **1.26s** | 70.91s | 24.49s | 1 |
| node-jsofd | 291ms | 604ms | 3.36s | 10.69s | 5.54s | 3.31s | 8.94s | 0 |
<!-- benchmark-table:pdf-speed:end -->

### CPU 总时间

CPU 总时间为进程用户态与内核态时间之和，累加所有线程及已等待回收的子进程；包含启动和包装脚本开销。多核并行时可超过墙钟时间，CPU 时间越小通常表示消耗的计算资源越少。历史数据未采集此指标，FAILED 表示尚未采集或转换失败，需重新运行基准。

<!-- benchmark-table:pdf-cpu:start -->

| 转换器 | hello.pdf | ano.pdf | 999.pdf | intro.pdf | zsbk.pdf | 1000-pages.pdf | GBT_33190-2016.pdf | 胜出次数 |
|---|---:|---:|---:|---:|---:|---:|---:|---:|
| go-zc310 | FAILED | FAILED | FAILED | FAILED | FAILED | FAILED | FAILED | 0 |
| go-ofdgo | FAILED | FAILED | FAILED | FAILED | FAILED | FAILED | FAILED | 0 |
| java-ofdrw | FAILED | FAILED | FAILED | FAILED | FAILED | FAILED | FAILED | 0 |
| python-pdf2ofd | FAILED | FAILED | FAILED | FAILED | FAILED | FAILED | FAILED | 0 |
| node-jsofd | FAILED | FAILED | FAILED | FAILED | FAILED | FAILED | FAILED | 0 |
<!-- benchmark-table:pdf-cpu:end -->

### 进程树内存峰值

Linux 下每 10ms 采样转换进程及其存活子孙进程的 RSS（驻留物理内存）之和，取采样峰值，包含包装脚本和运行时；MB 为 100 万字节。共享内存会重复计数，不等于实际独占物理内存；短于采样间隔的峰值、采样期间退出或脱离进程树的进程可能遗漏，读取受限也可能低估。采样带来少量额外开销。历史数据或未采到有效 RSS 显示为 FAILED，需重新运行基准。

<!-- benchmark-table:pdf-memory:start -->

| 转换器 | hello.pdf | ano.pdf | 999.pdf | intro.pdf | zsbk.pdf | 1000-pages.pdf | GBT_33190-2016.pdf | 胜出次数 |
|---|---:|---:|---:|---:|---:|---:|---:|---:|
| go-zc310 | FAILED | FAILED | FAILED | FAILED | FAILED | FAILED | FAILED | 0 |
| go-ofdgo | FAILED | FAILED | FAILED | FAILED | FAILED | FAILED | FAILED | 0 |
| java-ofdrw | FAILED | FAILED | FAILED | FAILED | FAILED | FAILED | FAILED | 0 |
| python-pdf2ofd | FAILED | FAILED | FAILED | FAILED | FAILED | FAILED | FAILED | 0 |
| node-jsofd | FAILED | FAILED | FAILED | FAILED | FAILED | FAILED | FAILED | 0 |
<!-- benchmark-table:pdf-memory:end -->

## OFD 文件大小

<!-- benchmark-table:pdf-ofd-size:start -->

| 转换器 | hello.pdf | ano.pdf | 999.pdf | intro.pdf | zsbk.pdf | 1000-pages.pdf | GBT_33190-2016.pdf | 胜出次数 |
|---|---:|---:|---:|---:|---:|---:|---:|---:|
| go-zc310 | **1.70K** | **405.40K** | **62.30K** | **7.30M** | **1.60M** | 4.80M | **7.10M** | 6 |
| go-ofdgo | 6.20K | FAILED | FAILED | 8.90M | FAILED | 64.70M | 13.80M | 0 |
| java-ofdrw | 5.60K | 601.30K | 817.60K | 34.20M | 20.40M | **3.50M** | 30.50M | 1 |
| python-pdf2ofd | 4.60K | 1.30M | 733.70K | 8.90M | **1.60M** | 86.50M | 36.80M | 0 |
| node-jsofd | 3.00K | 2.20M | 590.20K | 34.80M | 12.90M | 31.80M | 34.30M | 0 |
<!-- benchmark-table:pdf-ofd-size:end -->

## 文件大小对比（输出 OFD / 原 PDF）

越小表示压缩率越高，`1.00` 表示与源 PDF 等大；**加粗**为最小。

<!-- benchmark-table:pdf-size-ratio:start -->

| 转换器 | hello.pdf | ano.pdf | 999.pdf | intro.pdf | zsbk.pdf | 1000-pages.pdf | GBT_33190-2016.pdf |
|---|---:|---:|---:|---:|---:|---:|---:|
| go-zc310 | **0.48** | **0.45** | **0.78** | **1.09** | **1.02** | 2.54 | **1.46** |
| go-ofdgo | 1.77 | FAILED | FAILED | 1.33 | FAILED | 34.05 | 2.85 |
| java-ofdrw | 1.58 | 0.67 | 10.23 | 5.10 | 13.00 | **1.83** | 6.32 |
| python-pdf2ofd | 1.32 | 1.43 | 9.18 | 1.33 | 1.04 | 45.55 | 7.62 |
| node-jsofd | 0.84 | 2.47 | 7.38 | 5.18 | 8.18 | 16.72 | 7.10 |
<!-- benchmark-table:pdf-size-ratio:end -->

## 还原度对比

将每个转换器输出的 OFD 再经 go-zc310 转回 PDF（回环），与源 PDF 做像素对比（96 DPI）与文本对比。

像素脚本对 `1000-pages` 仅取前 3 页（仍校验总页数一致）；文本比较仍处理完整文档。所有转换器均使用相同的采样与评测规则。

**像素相似度**（100% 为完全相同）：

<!-- benchmark-table:pdf-pixel-similarity:start -->

| 转换器 | hello.pdf | ano.pdf | 999.pdf | intro.pdf | zsbk.pdf | 1000-pages.pdf | GBT_33190-2016.pdf | 胜出次数 |
|---|---:|---:|---:|---:|---:|---:|---:|---:|
| go-zc310 | 99.63% | **99.77%** | 98.57% | **99.93%** | **99.56%** | **99.84%** | **99.98%** | 5 |
| go-ofdgo | 99.63% | FAILED | FAILED | 98.02% | FAILED | 99.64% | 99.82% | 0 |
| java-ofdrw | **99.85%** | 98.72% | 98.75% | 99.49% | 99.06% | 97.99% | 99.66% | 1 |
| python-pdf2ofd | 99.72% | 98.35% | **98.94%** | 99.34% | 97.70% | 98.40% | FAILED | 1 |
| node-jsofd | 99.37% | 96.77% | 98.14% | FAILED | 98.82% | 99.50% | 99.94% | 0 |
<!-- benchmark-table:pdf-pixel-similarity:end -->

**文本提取能力**（从转换后 OFD 提取字符数，参考 PDF 提取字符数见括号）：

<!-- benchmark-table:pdf-text-extraction:start -->

| 转换器 | hello.pdf(22) | ano.pdf(7091) | 999.pdf(3934) | intro.pdf(5543) | zsbk.pdf(1150) | 1000-pages.pdf(263999) | GBT_33190-2016.pdf(107779) |
|---|---:|---:|---:|---:|---:|---:|---:|
| go-zc310 | 22 | 9222 | 9737 | 9314 | 2740 | 326022 | 170697 |
| go-ofdgo | 8 | FAILED | FAILED | 10920 | FAILED | 179998 | 234709 |
| java-ofdrw | 0 | 7049 | 0 | 1598 | 3 | 321998 | 714 |
| python-pdf2ofd | 0 | 0 | 0 | 0 | 0 | 0 | 0 |
| node-jsofd | 25 | 10812 | 9941 | 10222 | 2796 | 326448 | 225807 |
<!-- benchmark-table:pdf-text-extraction:end -->

**文本相似度**（4-gram Jaccard，100% 为与源 PDF 提取文本完全一致）：

<!-- benchmark-table:pdf-text-similarity:start -->

| 转换器 | hello.pdf | ano.pdf | 999.pdf | intro.pdf | zsbk.pdf | 1000-pages.pdf | GBT_33190-2016.pdf |
|---|---:|---:|---:|---:|---:|---:|---:|
| go-zc310 | 100.00% | 33.33% | 64.41% | 15.41% | 10.43% | 98.54% | 41.99% |
| go-ofdgo | 26.32% | FAILED | FAILED | 13.67% | FAILED | 2.63% | 1.08% |
| java-ofdrw | 0.00% | 64.59% | 0.00% | 4.58% | 0.00% | 98.86% | 0.01% |
| python-pdf2ofd | 0.00% | 0.00% | 0.00% | 0.00% | 0.00% | 0.00% | 0.00% |
| node-jsofd | 100.00% | 36.08% | 18.27% | 15.12% | 6.94% | 30.25% | 1.08% |
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

# 运行全部转换器（单个文件）
./bin/pdf2ofd-benchmark testdata/suwell/hello.pdf

# 批量运行所有测试文件
for f in testdata/suwell/*.pdf testdata/standards/*.pdf; do
  ./bin/pdf2ofd-benchmark "$f"
done

# 还原度/文本对比（输出位于 bin/output_ofd/）
python3 tools/compare_pdf2ofd.py 96
```
