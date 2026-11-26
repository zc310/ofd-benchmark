# PDF → OFD 对比测试

> [OFD 库汇总](ofd-libraries.md)：各 OFD 库能力 / 许可证对比表
>
> [OFD → PDF 对比测试](ofd-to-pdf.md)：反方向 OFD→PDF 转换对比



## 测试库

| 转换器         | 语言   | 库                                                                          | 版本                    | 说明                                                                                                         |
|----------------|--------|-----------------------------------------------------------------------------|-------------------------|--------------------------------------------------------------------------------------------------------------|
| go-zc310       | Go     | [zc310/ofd](https://github.com/zc310/ofd)                                   | v0.2.3-0.20261124113956-09f1815cd2bb | 基于 pdfcpu 的 PDF→OFD 导入器（路径/文本/图片/注解/渐变/图案/网格/大纲/元数据），文本保留为可提取 TextObject |
| go-ofdgo       | Go     | [xiaoqidun/ofdgo](https://github.com/xiaoqidun/ofdgo)（ConvertPDF + pdfgo） | v0.0.0-20261124073143-9d9bd2ba2cde | 基于 pdfgo 的 PDF→OFD 导入器（路径/文本/图片/注解）                                                          |
| java-ofdrw     | Java   | [ofdrw](https://github.com/ofdrw/ofdrw)（ofdrw-converter PDFConverter）     | 2.5.2 | 文本转为矢量轮廓                                                                               |
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
- CPU：AMD EPYC 7763 64-Core Processor（4 逻辑核心）
- 内存：15.6 GiB
<!-- benchmark-environment:end -->
- 系统已安装 `apt-get install fonts-wqy-zenhei fonts-noto-cjk`
- 测试前运行 `bash tools/install-fonts.sh`，下载 `doc-foundertype.zip`、`doc-zhongyi-etc.zip`、`latin-fonts.zip` 并安装到用户目录 `~/.local/share/fonts/hanzi-fonts/`，随后刷新字体缓存。字体下载链接与依赖安装命令见 [README：快速开始](../README.md#快速开始)；GitHub Actions 已加入此准备步骤。

> 所有转换器均由 GitHub Actions 在上述环境统一评测。回环统一使用 go-zc310，不改用各库自身回环；转换耗时、文件大小、回环像素及文本比较分别记录。

## 转换速度

图表按库累加下方速度表中全部文件的耗时，单位为秒；耗时为单次测量，多页文件可能主导总耗时。未完成全部文件的库不绘制柱形。

<!-- benchmark-charts:pdf:start -->

![PDF → OFD 各库全部文件总耗时](charts/pdf-to-ofd-speed.svg)
<!-- benchmark-charts:pdf:end -->

<!-- benchmark-table:pdf-speed:start -->

| 转换器 | hello.pdf | ano.pdf | 999.pdf | intro.pdf | zsbk.pdf | 1000-pages.pdf | GBT_33190-2016.pdf | 胜出次数 |
|---|---:|---:|---:|---:|---:|---:|---:|---:|
| go-zc310 | **23ms** | **97ms** | **315ms** | **511ms** | 1.63s | **1.23s** | **1.91s** | 6 |
| go-ofdgo | 94ms | 9.78s | 392ms | 622ms | 2.48s | 7.81s | 2.84s | 0 |
| java-ofdrw | 1.39s | 2.39s | 2.62s | 10.27s | 6.58s | 7.03s | 90.90s | 0 |
| python-pdf2ofd | 168ms | 1.23s | 795ms | 3.55s | **1.39s** | 79.09s | 27.62s | 1 |
| node-jsofd | 327ms | 643ms | 3.40s | 10.64s | 5.62s | 3.57s | 9.81s | 0 |
<!-- benchmark-table:pdf-speed:end -->

### CPU 总时间

CPU 总时间为进程用户态与内核态时间之和，累加所有线程及已等待回收的子进程；包含启动和包装脚本开销。多核并行时可超过墙钟时间。历史数据未采集此指标，FAILED 表示尚未采集或转换失败，需重新运行基准。

<!-- benchmark-table:pdf-cpu:start -->

| 转换器 | hello.pdf | ano.pdf | 999.pdf | intro.pdf | zsbk.pdf | 1000-pages.pdf | GBT_33190-2016.pdf | 胜出次数 |
|---|---:|---:|---:|---:|---:|---:|---:|---:|
| go-zc310 | **34ms** | **118ms** | **382ms** | **560ms** | 1.77s | **1.51s** | **2.26s** | 6 |
| go-ofdgo | 116ms | 10.09s | 708ms | 1.68s | 3.01s | 15.92s | 7.45s | 0 |
| java-ofdrw | 3.51s | 7.15s | 7.83s | 20.04s | 14.05s | 18.76s | 114.46s | 0 |
| python-pdf2ofd | 168ms | 1.23s | 795ms | 3.55s | **1.39s** | 79.09s | 27.62s | 1 |
| node-jsofd | 405ms | 1.09s | 992ms | 7.40s | 6.51s | 5.65s | 12.41s | 0 |
<!-- benchmark-table:pdf-cpu:end -->

### 进程树内存峰值

Linux 下每 10ms 采样转换进程及其存活子孙进程的 RSS（驻留物理内存）之和，取采样峰值，包含包装脚本和运行时；MB 为 100 万字节。共享内存会重复计数，不等于实际独占物理内存；短于采样间隔的峰值、采样期间退出或脱离进程树的进程可能遗漏，读取受限也可能低估。采样带来少量额外开销。历史数据或未采到有效 RSS 显示为 FAILED，需重新运行基准。

<!-- benchmark-table:pdf-memory:start -->

| 转换器 | hello.pdf | ano.pdf | 999.pdf | intro.pdf | zsbk.pdf | 1000-pages.pdf | GBT_33190-2016.pdf | 胜出次数 |
|---|---:|---:|---:|---:|---:|---:|---:|---:|
| go-zc310 | **37.33MB** | **50.56MB** | 209.39MB | **118.87MB** | 324.11MB | 190.98MB | **140.99MB** | 4 |
| go-ofdgo | 56.26MB | 238.39MB | 190.32MB | 138.87MB | 298.45MB | 1722.63MB | 377.95MB | 0 |
| java-ofdrw | 205.95MB | 461.40MB | 542.09MB | 610.77MB | 528.33MB | 759.55MB | 1715.90MB | 0 |
| python-pdf2ofd | 69.25MB | 99.33MB | **81.80MB** | 171.29MB | **148.03MB** | **94.77MB** | 233.14MB | 3 |
| node-jsofd | 142.72MB | 207.14MB | 180.38MB | 548.72MB | 462.73MB | 652.61MB | 1051.13MB | 0 |
<!-- benchmark-table:pdf-memory:end -->

## OFD 文件大小

<!-- benchmark-table:pdf-ofd-size:start -->

| 转换器 | hello.pdf | ano.pdf | 999.pdf | intro.pdf | zsbk.pdf | 1000-pages.pdf | GBT_33190-2016.pdf | 胜出次数 |
|---|---:|---:|---:|---:|---:|---:|---:|---:|
| go-zc310 | **1.68K** | 405.41K | **62.27K** | **7.33M** | **1.60M** | 4.83M | **7.07M** | 5 |
| go-ofdgo | 5.62K | **393.52K** | 107.44K | 8.80M | 2.68M | 58.63M | 12.99M | 1 |
| java-ofdrw | 5.55K | 601.33K | 817.57K | 34.24M | 20.45M | **3.48M** | 30.54M | 1 |
| python-pdf2ofd | 4.64K | 1.26M | 733.68K | 8.92M | 1.64M | 86.51M | 36.83M | 0 |
| node-jsofd | 2.95K | 2.18M | 590.22K | 34.76M | 12.87M | 31.75M | 34.32M | 0 |
<!-- benchmark-table:pdf-ofd-size:end -->

## 文件大小对比（输出 OFD / 原 PDF）

比值为输出 OFD 字节数除以源 PDF 字节数，`1.00` 表示与源 PDF 等大；**加粗**为最小值。

<!-- benchmark-table:pdf-size-ratio:start -->

| 转换器 | hello.pdf | ano.pdf | 999.pdf | intro.pdf | zsbk.pdf | 1000-pages.pdf | GBT_33190-2016.pdf |
|---|---:|---:|---:|---:|---:|---:|---:|
| go-zc310 | **0.48** | 0.45 | **0.78** | **1.09** | **1.02** | 2.54 | **1.46** |
| go-ofdgo | 1.60 | **0.44** | 1.34 | 1.31 | 1.71 | 30.87 | 2.69 |
| java-ofdrw | 1.58 | 0.67 | 10.23 | 5.10 | 13.00 | **1.83** | 6.32 |
| python-pdf2ofd | 1.32 | 1.43 | 9.18 | 1.33 | 1.04 | 45.55 | 7.62 |
| node-jsofd | 0.84 | 2.47 | 7.38 | 5.18 | 8.18 | 16.72 | 7.10 |
<!-- benchmark-table:pdf-size-ratio:end -->

## 还原度对比

将每个转换器输出的 OFD 再经 go-zc310 转回 PDF（回环），与源 PDF 做像素对比（96 DPI）与文本对比。

像素脚本对 `1000-pages` 仅取前 3 页（仍校验总页数一致）；文本比较仍处理完整文档。所有转换器均使用相同的采样与评测规则。

**像素相似度**（100% 为完全相同）：

<!-- benchmark-table:pdf-pixel-similarity:start -->

| 转换器         |    hello.pdf |      ano.pdf |      999.pdf |    intro.pdf |     zsbk.pdf | 1000-pages.pdf | GBT_33190-2016.pdf | 胜出次数 |
| -------------- | -----------: | -----------: | -----------: | -----------: | -----------: | -------------: | -----------------: | -------: |
| go-zc310       |     99.6301% | **99.7741%** |     98.5711% | **99.9959%** | **99.5571%** |   **99.8448%** |       **99.9783%** |        5 |
| go-ofdgo       |     99.6305% |     98.6155% |     98.4810% |     97.8163% |     99.2469% |       99.6369% |           99.8210% |        0 |
| java-ofdrw     | **99.8475%** |     98.7176% |     98.7505% |     99.5542% |     99.0587% |       97.9930% |           99.6620% |        1 |
| python-pdf2ofd |     99.7182% |     98.3471% | **98.9380%** |     99.3640% |     97.6966% |       98.3999% |             FAILED |        1 |
| node-jsofd     |     99.3738% |     96.7730% |     98.1504% |       FAILED |     98.9112% |       99.5017% |           99.9439% |        0 |
<!-- benchmark-table:pdf-pixel-similarity:end -->

**SSIM 结构相似度**（96 DPI）：与 OFD→PDF 使用相同的 RGB 7×7 均匀窗口 SSIM（样本协方差，`K1=0.01`、`K2=0.03`、`data_range=255`），窗口和通道平均后负分截为 0 并乘 100%，各页等权平均。沿用相同的页面尺寸、缺页和近乎空白输出检查，以及 `intro` 排除页和 `1000-pages` 前 3 页采样规则。新增指标尚无统一重测结果时记 `FAILED`。

<!-- benchmark-table:pdf-ssim:start -->

| 转换器         |    hello.pdf |      ano.pdf |      999.pdf |    intro.pdf |     zsbk.pdf | 1000-pages.pdf | GBT_33190-2016.pdf | 胜出次数 |
| -------------- | -----------: | -----------: | -----------: | -----------: | -----------: | -------------: | -----------------: | -------: |
| go-zc310       |     98.4914% | **96.6414%** |     94.5940% | **99.9957%** | **98.3616%** |   **99.6360%** |       **99.9912%** |        5 |
| go-ofdgo       |     98.4935% |     81.9039% |     93.8338% |     95.1711% |     96.0521% |       99.6337% |           99.6990% |        0 |
| java-ofdrw     |     99.3623% |     90.1064% |     96.0226% |     98.3068% |     96.5756% |       95.4830% |           98.9850% |        0 |
| python-pdf2ofd | **99.3975%** |     90.1858% | **97.9718%** |     97.0948% |     80.0304% |       96.2072% |             FAILED |        2 |
| node-jsofd     |     97.9284% |     89.7617% |     93.2531% |       FAILED |     96.4609% |       99.0147% |           99.9209% |        0 |
<!-- benchmark-table:pdf-ssim:end -->

**提取文本字符数**（从转换后 OFD 提取，参考 PDF 提取字符数见括号）：

<!-- benchmark-table:pdf-text-extraction:start -->

| 转换器 | hello.pdf(22) | ano.pdf(7091) | 999.pdf(3934) | intro.pdf(5543) | zsbk.pdf(1150) | 1000-pages.pdf(263999) | GBT_33190-2016.pdf(107779) |
|---|---:|---:|---:|---:|---:|---:|---:|
| go-zc310 | 22 | 9222 | 9737 | 9314 | 2740 | 326022 | 170697 |
| go-ofdgo | 8 | 12861 | 9781 | 8793 | 3290 | 151998 | 201415 |
| java-ofdrw | 0 | 7049 | 0 | 1598 | 3 | 321998 | 714 |
| python-pdf2ofd | 0 | 0 | 0 | 0 | 0 | 0 | 0 |
| node-jsofd | 25 | 10812 | 9941 | 10222 | 2796 | 326448 | 225807 |
<!-- benchmark-table:pdf-text-extraction:end -->

**文本相似度**（4-gram Jaccard，100% 为与源 PDF 提取文本完全一致）：

<!-- benchmark-table:pdf-text-similarity:start -->

| 转换器         | hello.pdf |  ano.pdf |  999.pdf | intro.pdf | zsbk.pdf | 1000-pages.pdf | GBT_33190-2016.pdf |
| -------------- | --------: | -------: | -------: | --------: | -------: | -------------: | -----------------: |
| go-zc310       | 100.0000% | 33.3333% | 64.4068% |  15.4102% | 10.4287% |       98.5396% |           41.9911% |
| go-ofdgo       |  26.3158% | 51.8482% | 39.0131% |  55.7696% |  8.8622% |        3.9886% |           33.5273% |
| java-ofdrw     |   0.0000% | 64.5949% |  0.0000% |   4.5805% |  0.0000% |       98.8604% |            0.0119% |
| python-pdf2ofd |   0.0000% |  0.0000% |  0.0000% |   0.0000% |  0.0000% |        0.0000% |            0.0000% |
| node-jsofd     | 100.0000% | 36.0781% | 18.2731% |  15.1217% |  6.9421% |       30.2548% |            1.0847% |
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
