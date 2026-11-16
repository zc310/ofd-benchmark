# OFD → PDF 对比测试

> [OFD 库汇总](ofd-libraries.md)：各 OFD 库能力 / 许可证对比表
>
> [PDF → OFD 对比测试](pdf-to-ofd.md)：反方向 PDF→OFD 转换对比

## 测试库

### Go 库

| 库                                                    | 版本                                  | OFD→PDF |  更新时间  | 备注                                    |
|-------------------------------------------------------|---------------------------------------|:-------:|:----------:|-----------------------------------------|
| [zc310-ofd](https://github.com/zc310/ofd)             | v0.1.10-0.20261114154015-087ae6006559 |   ✅    | 2026-11-14 | 已测试                                  |
| [xiaoqidun-ofdgo](https://github.com/xiaoqidun/ofdgo) | v0.0.0-20261111160922-3edf35199ec7 |   ✅    | 2026-11-14 | 已测试                                  |
| [ofdrw-rs Go binding](https://github.com/lihongjie0209/ofdrw-rs) | v0.5.0 | ✅ | 2026-11-15 | Go bindings + Linux 原生库，已测试 |

### Rust 库

| 库 | 版本 | OFD→PDF | 更新时间 | 备注 |
|---|---|:---:|:---:|---|
| [ofdbox-folio](https://github.com/ofdbox/ofdbox-folio) | `36e6877` | ✅ | 2026-11-16 | 通过 fontconfig 动态定位字体，已接入测试 |
| [sOFD](https://github.com/emfox/sOFD) | `65cff46` | ✅ | 2026-11-16 | Rust CLI；font-kit/fontconfig；缺字警告时保留 PDF 供视觉评测 |

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
- CPU：AMD EPYC 9V74 80-Core Processor（4 逻辑核心）
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
| [go-zc310](https://github.com/zc310/ofd) | 132ms | 234ms | 1.17s | **860ms** | **244ms** | 775ms | 2 |
| [go-ofdgo](https://github.com/xiaoqidun/ofdgo) | 243ms | 1.37s | 3.12s | 37.40s | 657ms | 1.90s | 0 |
| [go-ofdrw-rs](https://github.com/lihongjie0209/ofdrw-rs) | **26ms** | 903ms | 2.90s | 280.74s | 870ms | 635ms | 1 |
| [rust-ofdbox-folio](https://github.com/ofdbox/ofdbox-folio) | 104ms | **140ms** | **415ms** | 4.89s | 309ms | 267ms | 2 |
| [rust-sofd](https://github.com/emfox/sOFD) | 50ms | 301ms | 2.27s | 18.29s | 256ms | 2.39s | 0 |
| [easyofd](https://pypi.org/project/easyofd/) | 1.16s | 1.16s | 10.84s | 1.83s | 1.09s | 6.43s | 0 |
| [ofd2pdf](https://github.com/jsyzdej/ofd2pdf) | 166ms | FAILED | FAILED | 55.41s | 484ms | **241ms** | 1 |
| [ofdrw](https://github.com/ofdrw/ofdrw) | 1.56s | 1.35s | 3.51s | 3.59s | 2.76s | 2.10s | 0 |
| [@miconvert/ofd-to-pdf](https://www.npmjs.com/package/@miconvert/ofd-to-pdf) | 262ms | 652ms | 1.66s | 9.93s | 494ms | 1.07s | 0 |
<!-- benchmark-table:ofd-speed:end -->

> - `python-easyofd`: `pip install easyofd ofd2img`，依次尝试 easyofd、ofd2img
> - `python-ofd2pdf`: `pip install ofd2pdf`（Python ≥3.12），基于图片渲染
> - `java-ofdrw`: `org.ofdrw:ofdrw-converter:2.5.1`
> - `node-ofd2pdf`: `npm install @miconvert/ofd-to-pdf`，优先使用 CJS 入口，失败后回退 ESM 入口
> - `rust-ofdbox-folio`: 固定 Git revision `36e6877`；字体由 `fc-match` 按系统 fontconfig 配置解析，不依赖特定用户目录
> - `rust-sofd`: 固定 Git revision `65cff46`；字体由 font-kit/fontconfig 从系统配置解析；缺字警告时 CLI 返回非零，但包装器保留有效 PDF 继续视觉评测
>
> 各转换器的成功或失败情况以本轮结果表为准。

### PDF 文件大小

<!-- benchmark-table:ofd-pdf-size:start -->

| 转换器 | hello.ofd | ano.ofd | intro.ofd | 1000-pages.ofd | 999.ofd | zsbk.ofd | 胜出次数 |
|---|---:|---:|---:|---:|---:|---:|---:|
| [go-zc310](https://github.com/zc310/ofd) | 9.9KB | 88.0KB | 7.5MB | **687.1KB** | 78.4KB | 1.4MB | 1 |
| [go-ofdgo](https://github.com/xiaoqidun/ofdgo) | 15.5KB | 859.4KB | 9.6MB | 59.6MB | 130.4KB | 2.2MB | 0 |
| [go-ofdrw-rs](https://github.com/lihongjie0209/ofdrw-rs) | 5.4KB | 687.6KB | 9.5MB | 25.8MB | 208.9KB | 2.7MB | 0 |
| [rust-ofdbox-folio](https://github.com/ofdbox/ofdbox-folio) | 81.6KB | 270.8KB | 9.0MB | 32.8MB | 603.7KB | 1.9MB | 0 |
| [rust-sofd](https://github.com/emfox/sOFD) | 4.4KB | 971.2KB | 31.4MB | 88.4MB | 834.7KB | 15.9MB | 0 |
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
| [go-ofdrw-rs](https://github.com/lihongjie0209/ofdrw-rs) | 22 | 7093 | 5494 | 3934 | 1150 | 0 |
| [rust-ofdbox-folio](https://github.com/ofdbox/ofdbox-folio) | 8 | 6497 | 1959 | 773 | 717 | 0 |
| [rust-sofd](https://github.com/emfox/sOFD) | 0 | 0 | 0 | 0 | 0 | 0 |
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

**全页像素相似度**（96 DPI，1 - 平均归一化色差，100% 为完全相同；白底可能稀释内容差异）：

<!-- benchmark-table:ofd-pixel-similarity:start -->

| 转换器 | hello.ofd | ano.ofd | intro.ofd | 1000-pages.ofd | 999.ofd | zsbk.ofd | 胜出次数 |
|---|---:|---:|---:|---:|---:|---:|---:|
| [go-zc310](https://github.com/zc310/ofd) | **99.63%** | **99.70%** | **99.86%** | 99.46% | 98.55% | **99.51%** | 3 |
| [go-ofdgo](https://github.com/xiaoqidun/ofdgo) | **99.63%** | 99.39% | 99.74% | 99.45% | 98.53% | 98.98% | 1 |
| [go-ofdrw-rs](https://github.com/lihongjie0209/ofdrw-rs) | 99.62% | 99.42% | 99.58% | 98.19% | **98.83%** | 99.47% | 1 |
| [rust-ofdbox-folio](https://github.com/ofdbox/ofdbox-folio) | **99.63%** | 98.97% | 99.46% | 99.52% | 98.77% | 97.57% | 0 |
| [rust-sofd](https://github.com/emfox/sOFD) | 99.37% | FAILED | 96.91% | 99.45% | 98.63% | 97.88% | 0 |
| [easyofd](https://pypi.org/project/easyofd/) | 99.35% | 93.82% | 37.74% | 74.44% | 92.41% | 94.41% | 0 |
| [ofd2pdf](https://github.com/jsyzdej/ofd2pdf) | 99.61% | FAILED | FAILED | 99.34% | 97.74% | 94.33% | 0 |
| [ofdrw](https://github.com/ofdrw/ofdrw) | 99.30% | 99.53% | 99.27% | **99.75%** | 97.33% | 99.07% | 1 |
| [@miconvert/ofd-to-pdf](https://www.npmjs.com/package/@miconvert/ofd-to-pdf) | 99.01% | 97.55% | 96.29% | 97.56% | 95.80% | 94.40% | 0 |
<!-- benchmark-table:ofd-pixel-similarity:end -->

**内容区域像素相似度**（96 DPI）：只统计参考图或输出图中至少一方有内容的区域，并向外扩 1 像素；漏画和多画都会计入差异，避免大面积共同白底虚高。

<!-- benchmark-table:ofd-content-similarity:start -->

| 转换器 | hello.ofd | ano.ofd | intro.ofd | 1000-pages.ofd | 999.ofd | zsbk.ofd | 胜出次数 |
|---|---:|---:|---:|---:|---:|---:|---:|
| [go-zc310](https://github.com/zc310/ofd) | FAILED | FAILED | FAILED | FAILED | FAILED | FAILED | 0 |
| [go-ofdgo](https://github.com/xiaoqidun/ofdgo) | FAILED | FAILED | FAILED | FAILED | FAILED | FAILED | 0 |
| [go-ofdrw-rs](https://github.com/lihongjie0209/ofdrw-rs) | 76.39% | **96.38%** | **99.24%** | 97.96% | **91.62%** | **99.37%** | 4 |
| [rust-ofdbox-folio](https://github.com/ofdbox/ofdbox-folio) | **76.84%** | 93.51% | 98.66% | **99.45%** | 91.26% | 97.11% | 2 |
| [rust-sofd](https://github.com/emfox/sOFD) | 67.10% | FAILED | 94.35% | 99.38% | 90.69% | 97.49% | 0 |
| [easyofd](https://pypi.org/project/easyofd/) | FAILED | FAILED | FAILED | FAILED | FAILED | FAILED | 0 |
| [ofd2pdf](https://github.com/jsyzdej/ofd2pdf) | FAILED | FAILED | FAILED | FAILED | FAILED | FAILED | 0 |
| [ofdrw](https://github.com/ofdrw/ofdrw) | FAILED | FAILED | FAILED | FAILED | FAILED | FAILED | 0 |
| [@miconvert/ofd-to-pdf](https://www.npmjs.com/package/@miconvert/ofd-to-pdf) | FAILED | FAILED | FAILED | FAILED | FAILED | FAILED | 0 |
<!-- benchmark-table:ofd-content-similarity:end -->

**文件大小对比**（转换后 PDF 大小 / 数科参考 PDF 大小）：

<!-- benchmark-table:ofd-size-ratio:start -->

| 转换器 | hello.ofd | ano.ofd | intro.ofd | 1000-pages.ofd | 999.ofd | zsbk.ofd |
|---|---:|---:|---:|---:|---:|---:|
| 数科版式阅读器 | 1.00 | 1.00 | 1.00 | 1.00 | 1.00 | 1.00 |
| [go-zc310](https://github.com/zc310/ofd) | 2.82 | 0.10 | 1.11 | **0.35** | 0.98 | 0.89 |
| [go-ofdgo](https://github.com/xiaoqidun/ofdgo) | 4.40 | 0.95 | 1.43 | 31.37 | 1.63 | 1.43 |
| [go-ofdrw-rs](https://github.com/lihongjie0209/ofdrw-rs) | 1.54 | 0.76 | 1.42 | 13.59 | 2.61 | 1.69 |
| [rust-ofdbox-folio](https://github.com/ofdbox/ofdbox-folio) | 23.19 | 0.30 | 1.35 | 17.26 | 7.55 | 1.21 |
| [rust-sofd](https://github.com/emfox/sOFD) | 1.26 | 1.07 | 4.69 | 46.54 | 10.44 | 10.12 |
| [easyofd](https://pypi.org/project/easyofd/) | 6.70 | **0.04** | 5.70 | 0.37 | **0.92** | 8.62 |
| [ofd2pdf](https://github.com/jsyzdej/ofd2pdf) | 12.93 | FAILED | FAILED | 49.62 | 9.40 | **0.08** |
| [ofdrw](https://github.com/ofdrw/ofdrw) | 217.96 | 2.71 | 1.62 | 1.05 | 23.71 | 3.17 |
| [@miconvert/ofd-to-pdf](https://www.npmjs.com/package/@miconvert/ofd-to-pdf) | **0.47** | 0.22 | **0.16** | 4.83 | 1.04 | 0.35 |
<!-- benchmark-table:ofd-size-ratio:end -->

> - 像素相似度越高表示输出越贴近数科版式阅读器；仿色/抗锯齿等平滑差异会小幅降低相似度
> - 若参考页有可见内容而输出页几乎全白，像素对比会标记为 `FAILED`，避免大面积白底虚高
> - intro 的像素对比跳过 14～19 页（数科导出的参考 PDF 这几页背景图丢失，不计入对比）
> - `python-easyofd` 在 hello、1000-pages 转换失败；`python-ofd2pdf` 在 ano、intro 转换失败
> - 复现命令：`python3 tools/compare_pdf.py 96 <file>`

**文本相似度**（与数科参考 PDF 提取文本的 4-gram Jaccard 相似度，100% 为文本完全一致；`0.00%` 表示无文本或提取失败）：

参考文本长度（fitz 1.26.7 口径）：`hello=22` · `ano=7091` · `intro=5543` · `1000-pages=263999` · `999=3934` · `zsbk=1150`

<!-- benchmark-table:ofd-text-similarity:start -->

| 转换器 | hello.ofd | ano.ofd | intro.ofd | 1000-pages.ofd | 999.ofd | zsbk.ofd | 胜出次数 |
|---|---:|---:|---:|---:|---:|---:|---:|
| [go-zc310](https://github.com/zc310/ofd) | **100.00%** | 99.07% | **100.00%** | **100.00%** | 95.85% | 98.85% | 2 |
| [go-ofdgo](https://github.com/xiaoqidun/ofdgo) | 26.32% | 6.19% | 8.57% | 3.99% | 95.01% | 26.58% | 0 |
| [go-ofdrw-rs](https://github.com/lihongjie0209/ofdrw-rs) | **100.00%** | **99.73%** | 94.88% | **100.00%** | **100.00%** | **100.00%** | 3 |
| [rust-ofdbox-folio](https://github.com/ofdbox/ofdbox-folio) | 26.32% | 18.04% | 3.67% | 3.99% | 39.65% | 0.00% | 0 |
| [rust-sofd](https://github.com/emfox/sOFD) | 0.00% | 0.00% | 0.00% | 0.00% | 0.00% | 0.00% | 0 |
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
cargo build --release --manifest-path rust/ofdbox-folio-convert/Cargo.toml --target-dir rust/target
cp rust/target/release/ofdbox-folio-convert bin/ofdbox-folio-convert
cargo install --git https://github.com/emfox/sOFD --rev 65cff468ecc2af676180f540899fdf74da849c3a --locked --bin sofd --root /tmp/sofd-install
cp /tmp/sofd-install/bin/sofd bin/sofd-cli

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
./bin/ofd-benchmark <file.ofd> ofdbox-folio
./bin/ofd-benchmark <file.ofd> sofd
./bin/ofd-benchmark <file.ofd> java
./bin/ofd-benchmark <file.ofd> node
```

转换后的 PDF 文件保存在 `bin/output/` 目录，命名格式：`{原文件名}_{转换器}.pdf`

转换器名称：

- `go-ofdgo`
- `go-zc310`
- `rust-ofdbox-folio`
- `rust-sofd`
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
