# OFD → PDF 对比测试

> [OFD 库汇总](ofd-libraries.md)：各 OFD 库能力 / 许可证对比表
>
> [PDF → OFD 对比测试](pdf-to-ofd.md)：反方向 PDF→OFD 转换对比

## 测试库

### Go 库

| 库                                                    | 版本                                  | OFD→PDF |  更新时间  | 备注                                    |
|-------------------------------------------------------|---------------------------------------|:-------:|:----------:|-----------------------------------------|
| [zc310-ofd](https://github.com/zc310/ofd)             | v0.1.10-0.20261115081953-72f2d3bed048 |   ✅    | 2026-11-14 | 已测试                                  |
| [xiaoqidun-ofdgo](https://github.com/xiaoqidun/ofdgo) | v0.0.0-20261111160922-3edf35199ec7 |   ✅    | 2026-11-14 | 已测试                                  |
| [ofdrw-rs Go binding](https://github.com/lihongjie0209/ofdrw-rs) | v0.5.0 | ✅ | 2026-11-15 | Go bindings + Linux 原生库，已测试 |

### Rust 库

| 库 | 版本 | OFD→PDF | 更新时间 | 备注 |
|---|---|:---:|:---:|---|
| [ofdbox-folio](https://github.com/ofdbox/ofdbox-folio) | `36e6877` | ✅ | 2026-11-16 | 通过 fontconfig 动态定位字体，已接入测试 |
| [sOFD](https://github.com/emfox/sOFD) | `65cff46` | ✅ | 2026-11-16 | Rust CLI；font-kit/fontconfig；缺字警告时保留 PDF 供视觉评测 |
| [easyofd-rust](https://github.com/easy-4-rust/easyofd-rust) | 0.1.4 | ✅ | 2026-11-16 | 恢复评测；使用默认转换选项，质量以本轮实测为准 |

### Python 库

| 库                                            | 版本     | OFD→PDF |  更新时间  | 备注                       |
|-----------------------------------------------|----------|:-------:|:----------:|----------------------------|
| [easyofd](https://pypi.org/project/easyofd/)  | 20260427 |   ⚠️    | 2026-04-27 | 部分文件失败               |

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
- CPU：AMD EPYC 9V45 96-Core Processor（4 逻辑核心）
- 内存：15.6 GiB
<!-- benchmark-environment:end -->
- 系统已安装 `apt-get install fonts-wqy-zenhei fonts-noto-cjk`
- 测试前运行 `bash tools/install-fonts.sh`，下载 `doc-foundertype.zip`、`doc-zhongyi-etc.zip`、`latin-fonts.zip` 并安装到用户目录 `~/.local/share/fonts/hanzi-fonts/`，随后刷新字体缓存。字体下载链接与依赖安装命令见 [README：快速开始](../README.md#快速开始)；GitHub Actions 已加入此准备步骤。

> `rust-easyofd` 为 2026-11-16 本地补测（Ubuntu 26.04 LTS、AMD Ryzen 9 5950X，PyMuPDF 1.26.7 / NumPy 2.3.5）；其他行保留原结果。其耗时与上述 CI 环境不是同机数据，速度排名及胜出次数暂不作为性能结论；下次完整 CI 将统一重测。像素、内容区域及前景 F1 均使用当前算法重新计算。

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
| [go-zc310](https://github.com/zc310/ofd) | 61ms | 135ms | 825ms | **586ms** | 164ms | 566ms | 1 |
| [go-ofdgo](https://github.com/xiaoqidun/ofdgo) | 177ms | 1.02s | 2.30s | 27.51s | 565ms | 1.51s | 0 |
| [go-ofdrw-rs](https://github.com/lihongjie0209/ofdrw-rs) | **19ms** | 708ms | 2.59s | 223.77s | 501ms | 563ms | 1 |
| [rust-ofdbox-folio](https://github.com/ofdbox/ofdbox-folio) | 90ms | 117ms | **367ms** | 4.51s | 310ms | 246ms | 1 |
| [rust-sofd](https://github.com/emfox/sOFD) | 36ms | 275ms | 2.18s | 18.67s | 269ms | 2.54s | 0 |
| [rust-easyofd](https://github.com/easy-4-rust/easyofd-rust) | 42ms | **69ms** | 1.85s | 985ms | **163ms** | **74ms** | 3 |
| [easyofd](https://pypi.org/project/easyofd/) | 1.06s | 865ms | 7.63s | 1.41s | 884ms | 4.86s | 0 |
| [ofdrw](https://github.com/ofdrw/ofdrw) | 835ms | 1.01s | 2.95s | 2.82s | 2.03s | 1.66s | 0 |
| [@miconvert/ofd-to-pdf](https://www.npmjs.com/package/@miconvert/ofd-to-pdf) | 194ms | 480ms | 1.18s | 7.31s | 383ms | 798ms | 0 |
<!-- benchmark-table:ofd-speed:end -->

> - `python-easyofd`: `pip install easyofd`，仅使用 easyofd，不再回退到 ofd2img
> - `java-ofdrw`: `org.ofdrw:ofdrw-converter:2.5.1`
> - `node-ofd2pdf`: `npm install @miconvert/ofd-to-pdf`，优先使用 CJS 入口，失败后回退 ESM 入口
> - `rust-ofdbox-folio`: 固定 Git revision `36e6877`；字体由 `fc-match` 按系统 fontconfig 配置解析，不依赖特定用户目录
> - `rust-sofd`: 固定 Git revision `65cff46`；字体由 font-kit/fontconfig 从系统配置解析；缺字警告时 CLI 返回非零，但包装器保留有效 PDF 继续视觉评测
> - `rust-easyofd`: 固定 crates.io `easyofd = "=0.1.4"`；使用 `ConvertOptions::default()`，保留上游字体回退及输出行为，不对生成 PDF 修补或降级评测规则
>
> 各转换器的成功或失败情况以本轮结果表为准。

### PDF 文件大小

<!-- benchmark-table:ofd-pdf-size:start -->

| 转换器 | hello.ofd | ano.ofd | intro.ofd | 1000-pages.ofd | 999.ofd | zsbk.ofd | 胜出次数 |
|---|---:|---:|---:|---:|---:|---:|---:|
| [go-zc310](https://github.com/zc310/ofd) | 9.9KB | 88.0KB | 7.5MB | 687.1KB | 78.4KB | 1.4MB | 0 |
| [go-ofdgo](https://github.com/xiaoqidun/ofdgo) | 15.5KB | 859.4KB | 9.6MB | 59.6MB | 130.4KB | 2.2MB | 0 |
| [go-ofdrw-rs](https://github.com/lihongjie0209/ofdrw-rs) | 5.4KB | 682.8KB | 9.5MB | 25.8MB | 208.7KB | 2.7MB | 0 |
| [rust-ofdbox-folio](https://github.com/ofdbox/ofdbox-folio) | 81.6KB | 258.4KB | 9.0MB | 32.8MB | 602.9KB | 1.9MB | 0 |
| [rust-sofd](https://github.com/emfox/sOFD) | 4.4KB | 971.2KB | 31.4MB | 88.4MB | 834.7KB | 15.9MB | 0 |
| [rust-easyofd](https://github.com/easy-4-rust/easyofd-rust) | 37.3KB | 77.1KB | 28.9MB | **614.5KB** | 83.3KB | **77.8KB** | 2 |
| [easyofd](https://pypi.org/project/easyofd/) | 23.6KB | **36.1KB** | 38.2MB | 714.1KB | **73.1KB** | 13.6MB | 2 |
| [ofdrw](https://github.com/ofdrw/ofdrw) | 767.3KB | 2.4MB | 10.9MB | 2.0MB | 1.9MB | 5.0MB | 0 |
| [@miconvert/ofd-to-pdf](https://www.npmjs.com/package/@miconvert/ofd-to-pdf) | **1.6KB** | 196.5KB | **1.1MB** | 9.2MB | 82.9KB | 559.3KB | 2 |
<!-- benchmark-table:ofd-pdf-size:end -->

### 文本提取能力

<!-- benchmark-table:ofd-text-extraction:start -->

| 转换器 | hello.ofd | ano.ofd | intro.ofd | 999.ofd | zsbk.ofd | 胜出次数 |
|---|---:|---:|---:|---:|---:|---:|
| [go-zc310](https://github.com/zc310/ofd) | 22 | 7232 | 5544 | 3965 | 1162 | 0 |
| [go-ofdgo](https://github.com/xiaoqidun/ofdgo) | 8 | 567 | 506 | 3967 | 527 | 0 |
| [go-ofdrw-rs](https://github.com/lihongjie0209/ofdrw-rs) | 22 | 7093 | 5494 | 3934 | 1150 | 0 |
| [rust-ofdbox-folio](https://github.com/ofdbox/ofdbox-folio) | 8 | 6497 | 1959 | 773 | 717 | 0 |
| [rust-sofd](https://github.com/emfox/sOFD) | 0 | 0 | 0 | 0 | 0 | 0 |
| [rust-easyofd](https://github.com/easy-4-rust/easyofd-rust) | 22 | 6523 | 5622 | 3308 | 332 | 0 |
| [easyofd](https://pypi.org/project/easyofd/) | 22 | 1208 | **10194** | 3933 | 1508 | 1 |
| [ofdrw](https://github.com/ofdrw/ofdrw) | 26 | **7282** | 5550 | 4162 | 1169 | 1 |
| [@miconvert/ofd-to-pdf](https://www.npmjs.com/package/@miconvert/ofd-to-pdf) | **31** | 6489 | 9614 | **4249** | **1938** | 3 |
<!-- benchmark-table:ofd-text-extraction:end -->

> - 单位为字符数（chars）
> - 本轮使用 PyMuPDF 1.26.7 与 NumPy 2.3.5；表格数值为本轮测试结果

### 与数科版式输出对比

以数科版式阅读器导出的 `http://github.com/zc310/ofd-fixtures/tree/main/testdata/suwell` 为基准，对比转换效果。

**全页像素相似度**（96 DPI，1 - 平均归一化色差，100% 为完全相同；白底可能稀释内容差异）：

<!-- benchmark-table:ofd-pixel-similarity:start -->

| 转换器 | hello.ofd | ano.ofd | intro.ofd | 1000-pages.ofd | 999.ofd | zsbk.ofd | 胜出次数 |
|---|---:|---:|---:|---:|---:|---:|---:|
| [go-zc310](https://github.com/zc310/ofd) | **99.63%** | **99.70%** | **99.86%** | 99.47% | 98.55% | **99.51%** | 3 |
| [go-ofdgo](https://github.com/xiaoqidun/ofdgo) | **99.63%** | 99.39% | 99.75% | 99.46% | 98.56% | 98.90% | 1 |
| [go-ofdrw-rs](https://github.com/lihongjie0209/ofdrw-rs) | 99.60% | 99.38% | 99.54% | 98.30% | 98.55% | 99.38% | 0 |
| [rust-ofdbox-folio](https://github.com/ofdbox/ofdbox-folio) | 99.61% | 98.89% | 99.46% | 99.52% | 98.48% | 97.46% | 0 |
| [rust-sofd](https://github.com/emfox/sOFD) | 99.38% | FAILED | FAILED | 99.45% | **98.59%** | 97.84% | 1 |
| [rust-easyofd](https://github.com/easy-4-rust/easyofd-rust) | FAILED | FAILED | FAILED | FAILED | FAILED | FAILED | 0 |
| [easyofd](https://pypi.org/project/easyofd/) | FAILED | FAILED | FAILED | FAILED | FAILED | FAILED | 0 |
| [ofdrw](https://github.com/ofdrw/ofdrw) | 99.30% | 99.53% | 99.27% | **99.76%** | 97.33% | 99.06% | 1 |
| [@miconvert/ofd-to-pdf](https://www.npmjs.com/package/@miconvert/ofd-to-pdf) | 99.01% | 97.55% | 96.30% | 97.56% | 95.82% | 94.39% | 0 |
<!-- benchmark-table:ofd-pixel-similarity:end -->

**内容区域像素相似度**（96 DPI）：RGB 任一通道低于 250 视为非白内容。参考内容处的颜色误差与参考内容外的额外输出误差相加，仅除以参考非白像素数；分数为 `max(0, 1 - 归一化误差) × 100%`。不扩边、不裁剪，不允许输出扩大掩码分母稀释误差。双方空白页不参与内容分数平均，全空白文档无内容分数（记 `FAILED`）。此指标有方向性，始终以数科输出为参考；与旧并集扩边分数不可混用。

<!-- benchmark-table:ofd-content-similarity:start -->

| 转换器 | hello.ofd | ano.ofd | intro.ofd | 1000-pages.ofd | 999.ofd | zsbk.ofd | 胜出次数 |
|---|---:|---:|---:|---:|---:|---:|---:|
| [go-zc310](https://github.com/zc310/ofd) | 67.08% | **96.54%** | **99.73%** | 99.39% | 80.50% | **99.41%** | 3 |
| [go-ofdgo](https://github.com/xiaoqidun/ofdgo) | **67.34%** | 91.58% | 99.52% | 99.38% | 80.64% | 98.69% | 1 |
| [go-ofdrw-rs](https://github.com/lihongjie0209/ofdrw-rs) | 65.19% | 91.20% | 99.05% | 98.05% | 80.58% | 99.26% | 0 |
| [rust-ofdbox-folio](https://github.com/ofdbox/ofdbox-folio) | 66.08% | 83.10% | 97.86% | 99.45% | 79.65% | 96.96% | 0 |
| [rust-sofd](https://github.com/emfox/sOFD) | 45.27% | FAILED | FAILED | 99.37% | **81.51%** | 97.42% | 1 |
| [rust-easyofd](https://github.com/easy-4-rust/easyofd-rust) | FAILED | FAILED | FAILED | FAILED | FAILED | FAILED | 0 |
| [easyofd](https://pypi.org/project/easyofd/) | FAILED | FAILED | FAILED | FAILED | FAILED | FAILED | 0 |
| [ofdrw](https://github.com/ofdrw/ofdrw) | 38.55% | 94.92% | 98.77% | **99.72%** | 64.49% | 98.88% | 1 |
| [@miconvert/ofd-to-pdf](https://www.npmjs.com/package/@miconvert/ofd-to-pdf) | 12.94% | 62.67% | 91.96% | 97.21% | 43.51% | 93.31% | 0 |
<!-- benchmark-table:ofd-content-similarity:end -->

**前景覆盖 F1**：独立衡量形状覆盖，允许约 0.265 mm 的邻域距离（96 DPI 时为 1 像素），兼顾漏画召回率与多画精确率，不替代颜色误差。双方空白页不参与平均。

<!-- benchmark-table:ofd-foreground-f1:start -->

| 转换器 | hello.ofd | ano.ofd | intro.ofd | 1000-pages.ofd | 999.ofd | zsbk.ofd | 胜出次数 |
|---|---:|---:|---:|---:|---:|---:|---:|
| [go-zc310](https://github.com/zc310/ofd) | **86.46%** | 98.30% | 99.96% | 99.98% | 97.96% | **100.00%** | 2 |
| [go-ofdgo](https://github.com/xiaoqidun/ofdgo) | 86.44% | **98.37%** | **100.00%** | 99.97% | 98.10% | **100.00%** | 2 |
| [go-ofdrw-rs](https://github.com/lihongjie0209/ofdrw-rs) | 82.35% | 98.34% | 99.50% | 99.62% | 98.10% | **100.00%** | 0 |
| [rust-ofdbox-folio](https://github.com/ofdbox/ofdbox-folio) | 82.63% | 92.06% | 99.91% | **100.00%** | 97.50% | 65.34% | 1 |
| [rust-sofd](https://github.com/emfox/sOFD) | 67.14% | FAILED | FAILED | **100.00%** | **99.27%** | 99.98% | 1 |
| [rust-easyofd](https://github.com/easy-4-rust/easyofd-rust) | FAILED | FAILED | FAILED | FAILED | FAILED | FAILED | 0 |
| [easyofd](https://pypi.org/project/easyofd/) | FAILED | FAILED | FAILED | FAILED | FAILED | FAILED | 0 |
| [ofdrw](https://github.com/ofdrw/ofdrw) | 62.07% | 73.37% | 99.76% | **100.00%** | 87.67% | **100.00%** | 0 |
| [@miconvert/ofd-to-pdf](https://www.npmjs.com/package/@miconvert/ofd-to-pdf) | 6.08% | 4.83% | 92.90% | 99.21% | 11.47% | 20.41% | 0 |
<!-- benchmark-table:ofd-foreground-f1:end -->

> 页面显示尺寸相差超过 0.5 pt 时判失败；像素尺寸仅允许 1 像素舍入差异，差异部分用白色补齐。MuPDF 渲染警告或错误判为 `FAILED`，不对不完整渲染给分。
>
> `1000-pages` 因内容重复，全页像素相似度、内容区域像素相似度及前景 F1 仅比较前 3 页，仍校验总页数一致。转换、输出大小及文本比较仍针对完整文档。

`rust-easyofd 0.1.4` 本次 6/6 成功生成 PDF，但视觉指标 0/6 通过评测检查：hello、ano、1000-pages、999、zsbk 出现 `FT_Load_Glyph ... invalid argument` / `cannot render glyph`，intro 第 2 页触发近乎空白页检查。因此全页、内容区域与前景 F1 均记 `FAILED`，不是重新使用旧像素分数。文本提取独立评测，仍保留字符数及文本相似度。

**文件大小对比**（转换后 PDF 大小 / 数科参考 PDF 大小）：

<!-- benchmark-table:ofd-size-ratio:start -->

| 转换器 | hello.ofd | ano.ofd | intro.ofd | 1000-pages.ofd | 999.ofd | zsbk.ofd |
|---|---:|---:|---:|---:|---:|---:|
| 数科版式阅读器 | 1.00 | 1.00 | 1.00 | 1.00 | 1.00 | 1.00 |
| [go-zc310](https://github.com/zc310/ofd) | 2.82 | 0.10 | 1.11 | 0.35 | 0.98 | 0.89 |
| [go-ofdgo](https://github.com/xiaoqidun/ofdgo) | 4.40 | 0.95 | 1.43 | 31.37 | 1.63 | 1.43 |
| [go-ofdrw-rs](https://github.com/lihongjie0209/ofdrw-rs) | 1.54 | 0.76 | 1.42 | 13.59 | 2.61 | 1.69 |
| [rust-ofdbox-folio](https://github.com/ofdbox/ofdbox-folio) | 23.19 | 0.29 | 1.35 | 17.26 | 7.54 | 1.21 |
| [rust-sofd](https://github.com/emfox/sOFD) | 1.26 | 1.07 | 4.69 | 46.54 | 10.44 | 10.12 |
| [rust-easyofd](https://github.com/easy-4-rust/easyofd-rust) | 10.60 | 0.09 | 4.30 | **0.32** | 1.04 | **0.05** |
| [easyofd](https://pypi.org/project/easyofd/) | 6.70 | **0.04** | 5.70 | 0.37 | **0.92** | 8.62 |
| [ofdrw](https://github.com/ofdrw/ofdrw) | 217.96 | 2.71 | 1.62 | 1.05 | 23.71 | 3.17 |
| [@miconvert/ofd-to-pdf](https://www.npmjs.com/package/@miconvert/ofd-to-pdf) | **0.47** | 0.22 | **0.16** | 4.83 | 1.04 | 0.35 |
<!-- benchmark-table:ofd-size-ratio:end -->

> - 像素相似度越高表示输出越贴近数科版式阅读器；仿色/抗锯齿等平滑差异会小幅降低相似度
> - 若参考页有可见内容而输出页几乎全白，像素对比会标记为 `FAILED`，避免大面积白底虚高
> - intro 的像素对比跳过 14～19 页（数科导出的参考 PDF 这几页背景图丢失，不计入对比）
> - 各转换器失败情况以结果表为准；ofd2img、python-ofd2pdf 已移出评测
> - 复现命令：`python3 tools/compare_pdf.py 96 <file>`

**文本相似度**（与数科参考 PDF 提取文本的 4-gram Jaccard 相似度，100% 为文本完全一致；`0.00%` 表示无文本或提取失败）：

参考文本长度（fitz 1.26.7 口径）：`hello=22` · `ano=7091` · `intro=5543` · `1000-pages=263999` · `999=3934` · `zsbk=1150`

<!-- benchmark-table:ofd-text-similarity:start -->

| 转换器 | hello.ofd | ano.ofd | intro.ofd | 1000-pages.ofd | 999.ofd | zsbk.ofd | 胜出次数 |
|---|---:|---:|---:|---:|---:|---:|---:|
| [go-zc310](https://github.com/zc310/ofd) | **100.00%** | 99.06% | **99.71%** | **100.00%** | 94.18% | 98.85% | 2 |
| [go-ofdgo](https://github.com/xiaoqidun/ofdgo) | 26.32% | 5.28% | 8.57% | 3.99% | 95.01% | 26.58% | 0 |
| [go-ofdrw-rs](https://github.com/lihongjie0209/ofdrw-rs) | **100.00%** | **99.73%** | 94.88% | **100.00%** | **100.00%** | **100.00%** | 3 |
| [rust-ofdbox-folio](https://github.com/ofdbox/ofdbox-folio) | 26.32% | 18.04% | 3.67% | 3.99% | 39.65% | 0.00% | 0 |
| [rust-sofd](https://github.com/emfox/sOFD) | 0.00% | 0.00% | 0.00% | 0.00% | 0.00% | 0.00% | 0 |
| [rust-easyofd](https://github.com/easy-4-rust/easyofd-rust) | 65.22% | 26.83% | 46.67% | 82.83% | 63.67% | 21.17% | 0 |
| [easyofd](https://pypi.org/project/easyofd/) | **100.00%** | 12.57% | 0.06% | 98.38% | 80.54% | 11.20% | 0 |
| [ofdrw](https://github.com/ofdrw/ofdrw) | 17.14% | 70.58% | 38.28% | **100.00%** | 56.78% | 64.72% | 1 |
| [@miconvert/ofd-to-pdf](https://www.npmjs.com/package/@miconvert/ofd-to-pdf) | 10.81% | 97.33% | 5.86% | 83.31% | 41.24% | 4.86% | 0 |
<!-- benchmark-table:ofd-text-similarity:end -->

> - 复现命令：`python3 tools/compare_text.py <file>`

## 使用方法

```bash
# 编译
go build -o bin/ofd-benchmark .
go build -o bin/ofdgo-convert ./cmd/ofdgo-convert
go build -o bin/zc310-convert ./cmd/zc310-convert
cargo build --release --locked --manifest-path rust/ofdbox-folio-convert/Cargo.toml --target-dir rust/target
cp rust/target/release/ofdbox-folio-convert bin/ofdbox-folio-convert
cargo build --release --locked --manifest-path rust/easyofd-convert/Cargo.toml --target-dir rust/target
cp rust/target/release/easyofd-ofd2pdf rust/target/release/easyofd-pdf2ofd bin/
cargo install sofd-cli --git https://github.com/emfox/sOFD --rev 65cff468ecc2af676180f540899fdf74da849c3a --locked --bin sofd --root /tmp/sofd-install
cp /tmp/sofd-install/bin/sofd bin/sofd-cli

# 安装 Python 依赖
python -m pip install --only-binary=:all: --require-hashes -r python/requirements.txt

# 编译 Java 转换器（需要先下载依赖，见 java/README.md）
cd java && javac -cp "$(ls lib/*.jar | tr '\n' ':')" Ofd2Pdf.java && cd ..

# 安装 Node.js 依赖
cd node && npm ci && cd ..

# 运行测试（默认测试所有转换器）
./bin/ofd-benchmark <file.ofd>

# 只测试某个转换器
./bin/ofd-benchmark <file.ofd> ofdgo
./bin/ofd-benchmark <file.ofd> zc310
./bin/ofd-benchmark <file.ofd> python
./bin/ofd-benchmark <file.ofd> ofdbox-folio
./bin/ofd-benchmark <file.ofd> sofd
./bin/ofd-benchmark <file.ofd> easyofd-rust
./bin/ofd-benchmark <file.ofd> java
./bin/ofd-benchmark <file.ofd> node
```

转换后的 PDF 文件保存在 `bin/output/` 目录，命名格式：`{原文件名}_{转换器}.pdf`

转换器名称：

- `go-ofdgo`
- `go-zc310`
- `rust-ofdbox-folio`
- `rust-sofd`
- `rust-easyofd`
- `python-easyofd`
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
│   ├── ofd2pdf.py          # Python 转换器（仅 easyofd）
│   └── ofd2pdf_ofd2pdf.py  # 旧独立脚本，仅供参考，不参与评测
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
