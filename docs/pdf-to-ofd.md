# PDF → OFD 对比测试

> [OFD 库汇总](ofd-libraries.md)：各 OFD 库能力 / 许可证对比表
>
> [OFD → PDF 对比测试](ofd-to-pdf.md)：反方向 OFD→PDF 转换对比



## 测试库

| 库 | 版本 | 说明 |
|---|---|---|
| [go-zc310-ofd](https://github.com/zc310/ofd) | v0.2.3-0.20261124113956-09f1815cd2bb | 基于 pdfcpu 的 PDF→OFD 导入器（路径/文本/图片/注解/渐变/图案/网格/大纲/元数据），文本保留为可提取 TextObject |
| [go-xiaoqidun-ofdgo](https://github.com/xiaoqidun/ofdgo) | v0.0.0-20261124073143-9d9bd2ba2cde | 使用 ConvertPDF + pdfgo，支持路径/文本/图片/注解 |
| [java-ofdrw](https://github.com/ofdrw/ofdrw) | 2.5.2 | 使用 ofdrw-converter PDFConverter，文本转为矢量轮廓 |
| [python-pdf2ofd](https://github.com/wanglrebe/pdf2ofd) | 0.1.0 | 基于 PyMuPDF 渲染，文本栅格化 |
| [node-jsofd](https://github.com/Hufe921/jsOFD) | 1.0.1 | 使用 pdfToOfd + pdfjs-dist 重放绘制指令流（路径/文本/图片），不转换注解/渐变/图案/Type3/剪裁 |

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
- CPU：Intel(R) Xeon(R) 6975P-C（4 逻辑核心）
- 内存：15.3 GiB
<!-- benchmark-environment:end -->
- 系统已安装 `apt-get install fonts-wqy-zenhei fonts-noto-cjk`
- 测试前运行 `bash tools/install-fonts.sh`，下载 `doc-foundertype.zip`、`doc-zhongyi-etc.zip`、`latin-fonts.zip` 并安装到用户目录 `~/.local/share/fonts/hanzi-fonts/`，随后刷新字体缓存。字体下载链接与依赖安装命令见 [README：快速开始](../README.md#快速开始)；GitHub Actions 已加入此准备步骤。

> 所有转换器均由 GitHub Actions 在上述环境统一评测。回环统一使用 go-zc310-ofd，不改用各库自身回环；转换耗时、文件大小、回环像素及文本比较分别记录。

## 转换速度

图表按库累加全部文件的耗时和输出大小；耗时为单次测量，多页文件可能主导总耗时。

<!-- benchmark-charts:pdf:start -->

![PDF → OFD 各库总耗时及输出总大小](charts/pdf-to-ofd-combined.svg)

每项仅在全部文件有有效结果时显示；缺失项标为“无完整结果”。大小按表格中的取整数值累加，MB 为 100 万字节。文件较小不一定表示质量更好，请结合还原度等结果查看。
<!-- benchmark-charts:pdf:end -->

<!-- benchmark-table:pdf-speed:start -->

| 转换器 | hello.pdf | ano.pdf | 999.pdf | intro.pdf | zsbk.pdf | 1000-pages.pdf | GBT_33190-2016.pdf | 胜出次数 |
|---|---:|---:|---:|---:|---:|---:|---:|---:|
| go-zc310-ofd | **16ms** | **71ms** | **262ms** | **369ms** | 1.13s | **795ms** | **1.42s** | 6 |
| go-xiaoqidun-ofdgo | 76ms | 8.92s | 311ms | 431ms | 1.69s | 5.76s | 1.87s | 0 |
| java-ofdrw | 863ms | 1.53s | 1.60s | 7.58s | 4.64s | 4.72s | 80.93s | 0 |
| python-pdf2ofd | 125ms | 860ms | 551ms | 2.54s | **975ms** | 54.72s | 18.99s | 1 |
| node-jsofd | 229ms | 451ms | 3.29s | 9.60s | 4.99s | 2.55s | 7.60s | 0 |
<!-- benchmark-table:pdf-speed:end -->

### CPU 总时间

CPU 总时间为进程用户态与内核态时间之和，累加所有线程及已等待回收的子进程；包含启动和包装脚本开销。多核并行时可超过墙钟时间。历史数据未采集此指标，FAILED 表示尚未采集或转换失败，需重新运行基准。

<!-- benchmark-table:pdf-cpu:start -->

| 转换器 | hello.pdf | ano.pdf | 999.pdf | intro.pdf | zsbk.pdf | 1000-pages.pdf | GBT_33190-2016.pdf | 胜出次数 |
|---|---:|---:|---:|---:|---:|---:|---:|---:|
| go-zc310-ofd | **22ms** | **87ms** | **302ms** | **403ms** | 1.22s | **996ms** | **1.67s** | 6 |
| go-xiaoqidun-ofdgo | 93ms | 9.15s | 558ms | 1.17s | 2.12s | 11.66s | 5.11s | 0 |
| java-ofdrw | 2.30s | 4.75s | 4.85s | 14.10s | 9.59s | 13.72s | 99.22s | 0 |
| python-pdf2ofd | 125ms | 860ms | 551ms | 2.54s | **975ms** | 54.71s | 18.99s | 1 |
| node-jsofd | 294ms | 796ms | 704ms | 5.64s | 5.13s | 4.13s | 8.61s | 0 |
<!-- benchmark-table:pdf-cpu:end -->

### 进程树内存峰值

Linux 下每 10ms 采样转换进程及其存活子孙进程的 RSS（驻留物理内存）之和，取采样峰值，包含包装脚本和运行时；MB 为 100 万字节。共享内存会重复计数，不等于实际独占物理内存；短于采样间隔的峰值、采样期间退出或脱离进程树的进程可能遗漏，读取受限也可能低估。采样带来少量额外开销。历史数据或未采到有效 RSS 显示为 FAILED，需重新运行基准。

<!-- benchmark-table:pdf-memory:start -->

| 转换器 | hello.pdf | ano.pdf | 999.pdf | intro.pdf | zsbk.pdf | 1000-pages.pdf | GBT_33190-2016.pdf | 胜出次数 |
|---|---:|---:|---:|---:|---:|---:|---:|---:|
| go-zc310-ofd | **34.31MB** | **51.90MB** | 207.15MB | **120.80MB** | 310.50MB | 191.97MB | **143.99MB** | 4 |
| go-xiaoqidun-ofdgo | 56.95MB | 239.12MB | 192.06MB | 148.45MB | 294.12MB | 1865.28MB | 408.26MB | 0 |
| java-ofdrw | 206.80MB | 465.45MB | 489.18MB | 583.66MB | 748.25MB | 832.20MB | 1701.54MB | 0 |
| python-pdf2ofd | 69.19MB | 99.30MB | **81.95MB** | 171.10MB | **148.71MB** | **94.75MB** | 233.11MB | 3 |
| node-jsofd | 144.02MB | 200.67MB | 184.43MB | 539.21MB | 463.25MB | 641.60MB | 1091.01MB | 0 |
<!-- benchmark-table:pdf-memory:end -->

## OFD 文件大小

<!-- benchmark-table:pdf-ofd-size:start -->

| 转换器 | hello.pdf | ano.pdf | 999.pdf | intro.pdf | zsbk.pdf | 1000-pages.pdf | GBT_33190-2016.pdf | 胜出次数 |
|---|---:|---:|---:|---:|---:|---:|---:|---:|
| go-zc310-ofd | **1.68K** | 405.41K | **62.27K** | **7.33M** | **1.60M** | 4.83M | **7.07M** | 5 |
| go-xiaoqidun-ofdgo | 5.62K | **393.52K** | 107.44K | 8.80M | 2.68M | 58.63M | 12.99M | 1 |
| java-ofdrw | 5.55K | 601.33K | 817.57K | 34.24M | 20.45M | **3.48M** | 30.54M | 1 |
| python-pdf2ofd | 4.64K | 1.26M | 733.68K | 8.92M | 1.64M | 86.51M | 36.83M | 0 |
| node-jsofd | 2.95K | 2.18M | 590.22K | 34.76M | 12.87M | 31.75M | 34.32M | 0 |
<!-- benchmark-table:pdf-ofd-size:end -->

## 文件大小对比（输出 OFD / 原 PDF）

比值为输出 OFD 字节数除以源 PDF 字节数，`1.00` 表示与源 PDF 等大；**加粗**为最小值。

<!-- benchmark-table:pdf-size-ratio:start -->

| 转换器 | hello.pdf | ano.pdf | 999.pdf | intro.pdf | zsbk.pdf | 1000-pages.pdf | GBT_33190-2016.pdf |
|---|---:|---:|---:|---:|---:|---:|---:|
| go-zc310-ofd | **0.48** | 0.45 | **0.78** | **1.09** | **1.02** | 2.54 | **1.46** |
| go-xiaoqidun-ofdgo | 1.60 | **0.44** | 1.34 | 1.31 | 1.71 | 30.87 | 2.69 |
| java-ofdrw | 1.58 | 0.67 | 10.23 | 5.10 | 13.00 | **1.83** | 6.32 |
| python-pdf2ofd | 1.32 | 1.43 | 9.18 | 1.33 | 1.04 | 45.55 | 7.62 |
| node-jsofd | 0.84 | 2.47 | 7.38 | 5.18 | 8.18 | 16.72 | 7.10 |
<!-- benchmark-table:pdf-size-ratio:end -->

## 还原度对比

将每个转换器输出的 OFD 再经 go-zc310-ofd 转回 PDF（回环），与源 PDF 做像素对比（96 DPI）与文本对比。

像素脚本对 `1000-pages` 仅取前 3 页（仍校验总页数一致）；文本比较仍处理完整文档。所有转换器均使用相同的采样与评测规则。

**像素相似度**（100% 为完全相同）：

<!-- benchmark-table:pdf-pixel-similarity:start -->

| 转换器             |    hello.pdf |      ano.pdf |      999.pdf |    intro.pdf |     zsbk.pdf | 1000-pages.pdf | GBT_33190-2016.pdf | 胜出次数 |
| ------------------ | -----------: | -----------: | -----------: | -----------: | -----------: | -------------: | -----------------: | -------: |
| go-zc310-ofd       |     99.6301% | **99.7741%** |     98.5711% | **99.9959%** | **99.5571%** |   **99.8448%** |       **99.9783%** |        5 |
| go-xiaoqidun-ofdgo |     99.6305% |     98.6155% |     98.4810% |     97.8163% |     99.2469% |       99.6369% |           99.8210% |        0 |
| java-ofdrw         | **99.8475%** |     98.7176% |     98.7505% |     99.5542% |     99.0587% |       97.9930% |           99.6620% |        1 |
| python-pdf2ofd     |     99.7182% |     98.3471% | **98.9380%** |     99.3640% |     97.6966% |       98.3999% |             FAILED |        1 |
| node-jsofd         |     99.3738% |     96.7730% |     98.1504% |       FAILED |     98.9112% |       99.5017% |           99.9439% |        0 |
<!-- benchmark-table:pdf-pixel-similarity:end -->

**SSIM 结构相似度**（96 DPI）：与 OFD→PDF 使用相同的 RGB 7×7 均匀窗口 SSIM（样本协方差，`K1=0.01`、`K2=0.03`、`data_range=255`），窗口和通道平均后负分截为 0 并乘 100%，各页等权平均。沿用相同的页面尺寸、缺页和近乎空白输出检查，以及 `intro` 排除页和 `1000-pages` 前 3 页采样规则。新增指标尚无统一重测结果时记 `FAILED`。

<!-- benchmark-table:pdf-ssim:start -->

| 转换器             |    hello.pdf |      ano.pdf |      999.pdf |    intro.pdf |     zsbk.pdf | 1000-pages.pdf | GBT_33190-2016.pdf | 胜出次数 |
| ------------------ | -----------: | -----------: | -----------: | -----------: | -----------: | -------------: | -----------------: | -------: |
| go-zc310-ofd       |     98.4914% | **96.6414%** |     94.5940% | **99.9957%** | **98.3616%** |   **99.6360%** |       **99.9912%** |        5 |
| go-xiaoqidun-ofdgo |     98.4935% |     81.9039% |     93.8338% |     95.1711% |     96.0521% |       99.6337% |           99.6990% |        0 |
| java-ofdrw         |     99.3623% |     90.1064% |     96.0226% |     98.3068% |     96.5756% |       95.4830% |           98.9850% |        0 |
| python-pdf2ofd     | **99.3975%** |     90.1858% | **97.9718%** |     97.0948% |     80.0304% |       96.2072% |             FAILED |        2 |
| node-jsofd         |     97.9284% |     89.7617% |     93.2531% |       FAILED |     96.4609% |       99.0147% |           99.9209% |        0 |
<!-- benchmark-table:pdf-ssim:end -->

**提取文本字符数**（从转换后 OFD 提取，参考 PDF 提取字符数见括号）：

<!-- benchmark-table:pdf-text-extraction:start -->

| 转换器 | hello.pdf(22) | ano.pdf(7091) | 999.pdf(3934) | intro.pdf(5543) | zsbk.pdf(1150) | 1000-pages.pdf(263999) | GBT_33190-2016.pdf(107779) |
|---|---:|---:|---:|---:|---:|---:|---:|
| go-zc310-ofd | 22 | 9222 | 9737 | 9314 | 2740 | 326022 | 170697 |
| go-xiaoqidun-ofdgo | 8 | 12861 | 9781 | 8793 | 3290 | 151998 | 201415 |
| java-ofdrw | 0 | 7049 | 0 | 1598 | 3 | 321998 | 714 |
| python-pdf2ofd | 0 | 0 | 0 | 0 | 0 | 0 | 0 |
| node-jsofd | 25 | 10812 | 9941 | 10222 | 2796 | 326448 | 225807 |
<!-- benchmark-table:pdf-text-extraction:end -->

**文本相似度**（4-gram Jaccard，100% 为与源 PDF 提取文本完全一致）：

<!-- benchmark-table:pdf-text-similarity:start -->

| 转换器             | hello.pdf |  ano.pdf |  999.pdf | intro.pdf | zsbk.pdf | 1000-pages.pdf | GBT_33190-2016.pdf |
| ------------------ | --------: | -------: | -------: | --------: | -------: | -------------: | -----------------: |
| go-zc310-ofd       | 100.0000% | 33.3333% | 64.4068% |  15.4102% | 10.4287% |       98.5396% |           41.9911% |
| go-xiaoqidun-ofdgo |  26.3158% | 51.8482% | 39.0131% |  55.7696% |  8.8622% |        3.9886% |           33.5273% |
| java-ofdrw         |   0.0000% | 64.5949% |  0.0000% |   4.5805% |  0.0000% |       98.8604% |            0.0119% |
| python-pdf2ofd     |   0.0000% |  0.0000% |  0.0000% |   0.0000% |  0.0000% |        0.0000% |            0.0000% |
| node-jsofd         | 100.0000% | 36.0781% | 18.2731% |  15.1217% |  6.9421% |       30.2548% |            1.0847% |
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

# 还原度/文本对比（输出位于 /tmp/ofd-benchmark/output_ofd/）
python3 tools/compare_pdf2ofd.py 96
```
