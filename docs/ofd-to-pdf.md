# OFD → PDF 对比测试

> [OFD 库汇总](ofd-libraries.md)：各 OFD 库能力 / 许可证对比表
>
> [PDF → OFD 对比测试](pdf-to-ofd.md)：反方向 PDF→OFD 转换对比

## 测试库

### Go 库

| 库                                                    | 版本                                  | OFD→PDF |  更新时间  | 备注                                    |
|-------------------------------------------------------|---------------------------------------|:-------:|:----------:|-----------------------------------------|
| [zc310-ofd](https://github.com/zc310/ofd)             | v0.1.10-0.20261115025254-44bea96b8629 |   ✅    | 2026-11-14 | 已测试                                  |
| [xiaoqidun-ofdgo](https://github.com/xiaoqidun/ofdgo) | v0.0.0-20261111160922-3edf35199ec7 |   ✅    | 2026-11-14 | 已测试                                  |
| [ofdrw-rs Go binding](https://github.com/lihongjie0209/ofdrw-rs) | v0.5.0 | ✅ | 2026-11-15 | Go bindings + Linux 原生库，已测试 |


### Python 库

| 库                                            | 版本     | OFD→PDF |  更新时间  | 备注                       |
|-----------------------------------------------|----------|:-------:|:----------:|----------------------------|
| [easyofd](https://pypi.org/project/easyofd/)  | 20260427 |   ⚠️    | 2026-04-27 | 部分文件失败               |
| [ofd2img](https://pypi.org/project/ofd2img/)  | 0.1.2    |   ⚠️    | 2026-05-07 | 部分文件失败               |
| [ofd2pdf](https://github.com/jsyzdej/ofd2pdf) | 0.0.2    |   ⚠️    | 2026-07-08 | 部分文件失败，基于图片渲染 |

### Java 库

| 库                                      | 版本  | OFD→PDF | 更新时间   | 备注                     |
|-----------------------------------------|-------|:-------:|------------|--------------------------|
| [ofdrw](https://github.com/ofdrw/ofdrw) | 2.5.1 |   ✅    | 2026-11-14 | 需手动下载依赖，全部成功 |

### JavaScript 库

| 库                                                                           | 版本  | OFD→PDF |  更新时间  | 备注     |
|------------------------------------------------------------------------------|-------|:-------:|:----------:|----------|
| [@miconvert/ofd-to-pdf](https://www.npmjs.com/package/@miconvert/ofd-to-pdf) | 0.2.3 |   ✅    | 2026-08-03 | 全部成功 |

## 测试环境

<!-- benchmark-environment:start -->

- 系统：Ubuntu 24.04.5 LTS (x86_64)
- CPU：INTEL(R) XEON(R) PLATINUM 8573C（4 逻辑核心）
- 内存：15.6 GiB
<!-- benchmark-environment:end -->
- 系统已安装 `apt-get install fonts-wqy-zenhei fonts-noto-cjk`
- 测试前运行 `bash tools/install-fonts.sh`，下载 `doc-foundertype.zip`、`doc-zhongyi-etc.zip`、`latin-fonts.zip` 并安装到用户目录 `~/.local/share/fonts/hanzi-fonts/`，随后刷新字体缓存。字体下载链接与依赖安装命令见 [README：快速开始](../README.md#快速开始)；GitHub Actions 已加入此准备步骤。

## 测试文件

测试用 OFD 文件来自：https://github.com/zc310/ofd-fixtures

| 文件           | 大小 | 说明         |
|----------------|------|--------------|
| hello.ofd      | 1.5K | 简单测试文件 |
| ano.ofd        | 702K | 带签章文件   |
| intro.ofd      | 7.2M | 介绍文档     |
| 1000-pages.ofd | 508K | 1000页文件   |
| 999.ofd        | 30K  | 带签章文件   |
| zsbk.ofd       | 1.5M | 数科签章文件 |

> `testdata/suwell/` 下的 `*.pdf` 为各 OFD 的参考转换结果，由**数科版式阅读器**导出，用作像素级对比基准。

## 测试结果

### 转换速度

<!-- benchmark-table:ofd-speed:start -->

| 转换器 | hello.ofd | ano.ofd | intro.ofd | 1000-pages.ofd | 999.ofd | zsbk.ofd | 胜出次数 |
|---|---:|---:|---:|---:|---:|---:|---:|
| [go-zc310](https://github.com/zc310/ofd) | 163ms | **268ms** | **1.39s** | **943ms** | **295ms** | 928ms | 4 |
| [go-ofdgo](https://github.com/xiaoqidun/ofdgo) | 352ms | 1.81s | 3.95s | 48.85s | 806ms | 2.43s | 0 |
| [go-ofdrw-rs](https://github.com/lihongjie0209/ofdrw-rs) | **29ms** | 799ms | 3.92s | 361.75s | 645ms | 789ms | 1 |
| [easyofd](https://pypi.org/project/easyofd/) | 1.58s | 1.41s | 13.06s | 2.05s | 1.28s | 7.54s | 0 |
| [ofd2pdf](https://github.com/jsyzdej/ofd2pdf) | 181ms | FAILED | FAILED | 53.60s | 521ms | **256ms** | 1 |
| [ofdrw](https://github.com/ofdrw/ofdrw) | 1.54s | 1.60s | 4.37s | 4.05s | 2.69s | 2.28s | 0 |
| [@miconvert/ofd-to-pdf](https://www.npmjs.com/package/@miconvert/ofd-to-pdf) | 316ms | 819ms | 2.05s | 12.29s | 588ms | 1.26s | 0 |
<!-- benchmark-table:ofd-speed:end -->

> - `python-easyofd`: `pip install easyofd ofd2img`，依次尝试 easyofd、ofd2img
> - `python-ofd2pdf`: `pip install ofd2pdf`（Python ≥3.12），基于图片渲染
> - `java-ofdrw`: `org.ofdrw:ofdrw-converter:2.5.1`
> - `node-ofd2pdf`: `npm install @miconvert/ofd-to-pdf`，优先使用 CJS 入口，失败后回退 ESM 入口
>
> 各转换器的成功或失败情况以本轮结果表为准。

### PDF 文件大小

<!-- benchmark-table:ofd-pdf-size:start -->

| 转换器 | hello.ofd | ano.ofd | intro.ofd | 1000-pages.ofd | 999.ofd | zsbk.ofd | 胜出次数 |
|---|---:|---:|---:|---:|---:|---:|---:|
| [go-zc310](https://github.com/zc310/ofd) | 9.9KB | 88.0KB | 7.5MB | **687.1KB** | 78.4KB | 1.4MB | 1 |
| [go-ofdgo](https://github.com/xiaoqidun/ofdgo) | 15.5KB | 859.4KB | 9.6MB | 59.6MB | 130.4KB | 2.2MB | 0 |
| [go-ofdrw-rs](https://github.com/lihongjie0209/ofdrw-rs) | 5.4KB | 682.8KB | 9.5MB | 25.8MB | 208.7KB | 2.7MB | 0 |
| [easyofd](https://pypi.org/project/easyofd/) | 23.6KB | **36.1KB** | 38.2MB | 714.1KB | **73.1KB** | 13.6MB | 2 |
| [ofd2pdf](https://github.com/jsyzdej/ofd2pdf) | 45.5KB | FAILED | FAILED | 94.2MB | 751.7KB | **128.9KB** | 1 |
| [ofdrw](https://github.com/ofdrw/ofdrw) | 767.3KB | 2.4MB | 10.9MB | 2.0MB | 1.9MB | 5.0MB | 0 |
| [@miconvert/ofd-to-pdf](https://www.npmjs.com/package/@miconvert/ofd-to-pdf) | **1.6KB** | 196.5KB | **1.1MB** | 9.2MB | 82.9KB | 559.3KB | 2 |
<!-- benchmark-table:ofd-pdf-size:end -->

### 文本提取能力

<!-- benchmark-table:ofd-text-extraction:start -->

| 转换器 | hello.ofd | ano.ofd | intro.ofd | 999.ofd | zsbk.ofd | 胜出次数 |
|---|---:|---:|---:|---:|---:|---:|
| [go-zc310](https://github.com/zc310/ofd) | 22 | 7280 | 5550 | 3971 | 1162 | 0 |
| [go-ofdgo](https://github.com/xiaoqidun/ofdgo) | 8 | 567 | 506 | 3967 | 527 | 0 |
| [go-ofdrw-rs](https://github.com/lihongjie0209/ofdrw-rs) | 22 | 7093 | 5505 | 3934 | 1150 | 0 |
| [easyofd](https://pypi.org/project/easyofd/) | 22 | 1208 | **10194** | 3933 | 1508 | 1 |
| [ofd2pdf](https://github.com/jsyzdej/ofd2pdf) | 0 | FAILED | FAILED | 0 | 0 | 0 |
| [ofdrw](https://github.com/ofdrw/ofdrw) | 26 | **7282** | 5550 | 4162 | 1169 | 1 |
| [@miconvert/ofd-to-pdf](https://www.npmjs.com/package/@miconvert/ofd-to-pdf) | **31** | 6534 | 9614 | **4249** | **1938** | 3 |
<!-- benchmark-table:ofd-text-extraction:end -->

> - `python-ofd2pdf` 基于图片渲染，无法提取文本
> - 单位为字符数（chars）
> - 本轮使用 PyMuPDF 1.26.7 与 NumPy 2.3.5；表格数值为本轮测试结果

### 与数科版式输出对比

以数科版式阅读器导出的 `http://github.com/zc310/ofd-fixtures/tree/main/testdata/suwell` 为基准，对比转换效果。

**像素相似度**（96 DPI，1 - 平均归一化色差，100% 为完全相同）：

<!-- benchmark-table:ofd-pixel-similarity:start -->

| 转换器 | hello.ofd | ano.ofd | intro.ofd | 1000-pages.ofd | 999.ofd | zsbk.ofd | 胜出次数 |
|---|---:|---:|---:|---:|---:|---:|---:|
| [go-zc310](https://github.com/zc310/ofd) | **99.63%** | **99.70%** | **99.86%** | 99.46% | **98.55%** | **99.51%** | 4 |
| [go-ofdgo](https://github.com/xiaoqidun/ofdgo) | **99.63%** | 99.39% | 99.74% | 99.45% | 98.53% | 98.98% | 1 |
| [go-ofdrw-rs](https://github.com/lihongjie0209/ofdrw-rs) | 99.60% | 99.37% | 99.53% | 98.30% | **98.55%** | 99.38% | 0 |
| [easyofd](https://pypi.org/project/easyofd/) | 99.35% | 93.82% | 37.74% | 74.44% | 92.41% | 94.41% | 0 |
| [ofd2pdf](https://github.com/jsyzdej/ofd2pdf) | 99.61% | FAILED | FAILED | 99.34% | 97.74% | 94.33% | 0 |
| [ofdrw](https://github.com/ofdrw/ofdrw) | 99.30% | 99.53% | 99.27% | **99.75%** | 97.33% | 99.07% | 1 |
| [@miconvert/ofd-to-pdf](https://www.npmjs.com/package/@miconvert/ofd-to-pdf) | 99.01% | 97.55% | 96.29% | 97.56% | 95.80% | 94.40% | 0 |
<!-- benchmark-table:ofd-pixel-similarity:end -->

**文件大小对比**（转换后 PDF 大小 / 数科参考 PDF 大小）：

<!-- benchmark-table:ofd-size-ratio:start -->

| 转换器 | hello.ofd | ano.ofd | intro.ofd | 1000-pages.ofd | 999.ofd | zsbk.ofd |
|---|---:|---:|---:|---:|---:|---:|
| 数科版式阅读器 | 1.00 | 1.00 | 1.00 | 1.00 | 1.00 | 1.00 |
| [go-zc310](https://github.com/zc310/ofd) | 2.82 | 0.10 | 1.11 | **0.35** | 0.98 | 0.89 |
| [go-ofdgo](https://github.com/xiaoqidun/ofdgo) | 4.40 | 0.95 | 1.43 | 31.37 | 1.63 | 1.43 |
| [go-ofdrw-rs](https://github.com/lihongjie0209/ofdrw-rs) | 1.54 | 0.76 | 1.42 | 13.59 | 2.61 | 1.69 |
| [easyofd](https://pypi.org/project/easyofd/) | 6.70 | **0.04** | 5.70 | 0.37 | **0.92** | 8.62 |
| [ofd2pdf](https://github.com/jsyzdej/ofd2pdf) | 12.93 | FAILED | FAILED | 49.62 | 9.40 | **0.08** |
| [ofdrw](https://github.com/ofdrw/ofdrw) | 217.96 | 2.71 | 1.62 | 1.05 | 23.71 | 3.17 |
| [@miconvert/ofd-to-pdf](https://www.npmjs.com/package/@miconvert/ofd-to-pdf) | **0.47** | 0.22 | **0.16** | 4.83 | 1.04 | 0.35 |
<!-- benchmark-table:ofd-size-ratio:end -->

> - 像素相似度越高表示输出越贴近数科版式阅读器；仿色/抗锯齿等平滑差异会小幅降低相似度
> - intro 的像素对比跳过 14～19 页（数科导出的参考 PDF 这几页背景图丢失，不计入对比）
> - `python-easyofd` 在 hello、1000-pages 转换失败；`python-ofd2pdf` 在 ano、intro 转换失败
> - 复现命令：`python3 tools/compare_pdf.py 96 <file>`

**文本相似度**（与数科参考 PDF 提取文本的 4-gram Jaccard 相似度，100% 为文本完全一致；`0.00%` 表示无文本或提取失败）：

参考文本长度（fitz 1.26.7 口径）：`hello=22` · `ano=7091` · `intro=5543` · `1000-pages=263999` · `999=3934` · `zsbk=1150`

<!-- benchmark-table:ofd-text-similarity:start -->

| 转换器 | hello.ofd | ano.ofd | intro.ofd | 1000-pages.ofd | 999.ofd | zsbk.ofd | 胜出次数 |
|---|---:|---:|---:|---:|---:|---:|---:|
| [go-zc310](https://github.com/zc310/ofd) | **100.00%** | **99.07%** | **100.00%** | **100.00%** | 95.85% | 98.85% | 3 |
| [go-ofdgo](https://github.com/xiaoqidun/ofdgo) | 26.32% | 6.19% | 8.57% | 3.99% | 95.01% | 26.58% | 0 |
| [go-ofdrw-rs](https://github.com/lihongjie0209/ofdrw-rs) | **100.00%** | 72.50% | 95.77% | **100.00%** | **100.00%** | **100.00%** | 2 |
| [easyofd](https://pypi.org/project/easyofd/) | **100.00%** | 12.62% | 0.08% | 98.38% | 80.54% | 11.20% | 0 |
| [ofd2pdf](https://github.com/jsyzdej/ofd2pdf) | 0.00% | FAILED | FAILED | 0.00% | 0.00% | 0.00% | 0 |
| [ofdrw](https://github.com/ofdrw/ofdrw) | 17.14% | 97.24% | 38.30% | **100.00%** | 56.78% | 64.72% | 1 |
| [@miconvert/ofd-to-pdf](https://www.npmjs.com/package/@miconvert/ofd-to-pdf) | 10.81% | 70.77% | 5.88% | 83.31% | 41.24% | 4.86% | 0 |
<!-- benchmark-table:ofd-text-similarity:end -->

> - 复现命令：`python3 tools/compare_text.py <file>`

## 使用方法

```bash
# 编译
go build -o bin/ofd-benchmark .
go build -o bin/ofdgo-convert ./cmd/ofdgo-convert
go build -o bin/zc310-convert ./cmd/zc310-convert

# 安装 Python 依赖
pip install easyofd ofd2img ofd2pdf

# 编译 Java 转换器（需要先下载依赖，见 java/README.md）
cd java && javac -cp "$(ls lib/*.jar | tr '\n' ':')" Ofd2Pdf.java && cd ..

# 安装 Node.js 依赖
cd node && npm install && cd ..

# 运行测试（默认测试所有转换器）
./bin/ofd-benchmark <file.ofd>

# 只测试某个转换器
./bin/ofd-benchmark <file.ofd> ofdgo
./bin/ofd-benchmark <file.ofd> zc310
./bin/ofd-benchmark <file.ofd> python
./bin/ofd-benchmark <file.ofd> ofd2pdf
./bin/ofd-benchmark <file.ofd> java
./bin/ofd-benchmark <file.ofd> node
```

转换后的 PDF 文件保存在 `bin/output/` 目录，命名格式：`{原文件名}_{转换器}.pdf`

转换器名称：

- `go-ofdgo`
- `go-zc310`
- `python-easyofd`
- `python-ofd2pdf`
- `java-ofdrw`
- `node-ofd2pdf`

## 项目结构

```
ofd-benchmark/
├── bin/
│   ├── ofd-benchmark       # 主程序
│   ├── ofdgo-convert       # Go 转换器
│   ├── zc310-convert       # Go 转换器
│   └── output/             # 转换后的 PDF 文件
├── python/
│   ├── ofd2pdf.py          # Python 转换器（easyofd + ofd2img）
│   └── ofd2pdf_ofd2pdf.py
├── java/
│   ├── Ofd2Pdf.java        # Java 转换器（ofdrw）
│   ├── ofd2pdf_java.sh     # Java 启动脚本
│   ├── lib/                # JAR 依赖
│   └── README.md           # 依赖下载说明
├── node/
│   ├── ofd2pdf.mjs         # Node.js 转换器（@miconvert/ofd-to-pdf）
│   ├── ofd2pdf_node.sh     # Node.js 启动脚本
│   └── package.json
├── cmd/
│   ├── ofdgo-convert/
│   └── zc310-convert/
├── main.go
├── go.mod
└── README.md
```
