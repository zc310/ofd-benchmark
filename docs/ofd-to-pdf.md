# OFD → PDF 对比测试

> [OFD 库汇总](ofd-libraries.md)：各 OFD 库能力 / 许可证对比表
>
> [PDF → OFD 对比测试](pdf-to-ofd.md)：反方向 PDF→OFD 转换对比

## 测试库

### Go 库

| 库                                                    | 版本                                  | OFD→PDF |  更新时间  | 备注                                    |
|-------------------------------------------------------|---------------------------------------|:-------:|:----------:|-----------------------------------------|
| [zc310-ofd](https://github.com/zc310/ofd)             | v0.2.2-0.20261119021541-2f1efcb69926 |   ✅    | 2026-11-14 | 已测试                                  |
| [xiaoqidun-ofdgo](https://github.com/xiaoqidun/ofdgo) | v0.0.0-20261119013312-38748f24bfef |   ✅    | 2026-11-14 | 已测试                                  |
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
- CPU：Intel(R) Xeon(R) 6975P-C（4 逻辑核心）
- 内存：15.3 GiB
<!-- benchmark-environment:end -->
- 系统已安装 `apt-get install fonts-wqy-zenhei fonts-noto-cjk`
- 测试前运行 `bash tools/install-fonts.sh`，下载 `doc-foundertype.zip`、`doc-zhongyi-etc.zip`、`latin-fonts.zip` 并安装到用户目录 `~/.local/share/fonts/hanzi-fonts/`，随后刷新字体缓存。字体下载链接与依赖安装命令见 [README：快速开始](../README.md#快速开始)；GitHub Actions 已加入此准备步骤。

> 所有转换器均由 GitHub Actions 在上述环境统一评测。像素、内容区域及前景 F1 均使用当前算法计算；速度排名及胜出次数应结合输出质量解读，成功生成文件不代表通过质量检查。

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
| [go-zc310](https://github.com/zc310/ofd) | 76ms | 157ms | 833ms | **587ms** | **176ms** | 589ms | 2 |
| [go-ofdgo](https://github.com/xiaoqidun/ofdgo) | 2.06s | 1.23s | 2.52s | 31.33s | 528ms | 1.69s | 0 |
| [go-ofdrw-rs](https://github.com/lihongjie0209/ofdrw-rs) | **29ms** | 500ms | 2.73s | 261.24s | 443ms | 551ms | 1 |
| [rust-ofdbox-folio](https://github.com/ofdbox/ofdbox-folio) | 111ms | **135ms** | **445ms** | 5.32s | 331ms | **269ms** | 3 |
| [rust-sofd](https://github.com/emfox/sOFD) | 35ms | 341ms | 2.66s | 21.83s | 289ms | 2.88s | 0 |
| [easyofd](https://pypi.org/project/easyofd/) | 1.92s | 951ms | 9.40s | 1.47s | 926ms | 5.52s | 0 |
| [ofdrw](https://github.com/ofdrw/ofdrw) | 1.39s | 1.12s | 2.79s | 3.14s | 1.75s | 1.77s | 0 |
| [@miconvert/ofd-to-pdf](https://www.npmjs.com/package/@miconvert/ofd-to-pdf) | 226ms | 574ms | 1.46s | 8.83s | 422ms | 903ms | 0 |
<!-- benchmark-table:ofd-speed:end -->

### CPU 总时间

CPU 总时间为进程用户态与内核态时间之和，累加所有线程及已等待回收的子进程；包含启动和包装脚本开销。多核并行时可超过墙钟时间，CPU 时间越小通常表示消耗的计算资源越少。历史数据未采集此指标，FAILED 表示尚未采集或转换失败，需重新运行基准。

<!-- benchmark-table:ofd-cpu:start -->

| 转换器 | hello.ofd | ano.ofd | intro.ofd | 1000-pages.ofd | 999.ofd | zsbk.ofd | 胜出次数 |
|---|---:|---:|---:|---:|---:|---:|---:|
| [go-zc310](https://github.com/zc310/ofd) | FAILED | FAILED | FAILED | FAILED | FAILED | FAILED | 0 |
| [go-ofdgo](https://github.com/xiaoqidun/ofdgo) | FAILED | FAILED | FAILED | FAILED | FAILED | FAILED | 0 |
| [go-ofdrw-rs](https://github.com/lihongjie0209/ofdrw-rs) | FAILED | FAILED | FAILED | FAILED | FAILED | FAILED | 0 |
| [rust-ofdbox-folio](https://github.com/ofdbox/ofdbox-folio) | FAILED | FAILED | FAILED | FAILED | FAILED | FAILED | 0 |
| [rust-sofd](https://github.com/emfox/sOFD) | FAILED | FAILED | FAILED | FAILED | FAILED | FAILED | 0 |
| [easyofd](https://pypi.org/project/easyofd/) | FAILED | FAILED | FAILED | FAILED | FAILED | FAILED | 0 |
| [ofdrw](https://github.com/ofdrw/ofdrw) | FAILED | FAILED | FAILED | FAILED | FAILED | FAILED | 0 |
| [@miconvert/ofd-to-pdf](https://www.npmjs.com/package/@miconvert/ofd-to-pdf) | FAILED | FAILED | FAILED | FAILED | FAILED | FAILED | 0 |
<!-- benchmark-table:ofd-cpu:end -->

### 进程树内存峰值

Linux 下每 10ms 采样转换进程及其存活子孙进程的 RSS（驻留物理内存）之和，取采样峰值，包含包装脚本和运行时；MB 为 100 万字节。共享内存会重复计数，不等于实际独占物理内存；短于采样间隔的峰值、采样期间退出或脱离进程树的进程可能遗漏，读取受限也可能低估。采样带来少量额外开销。历史数据或未采到有效 RSS 显示为 FAILED，需重新运行基准。

<!-- benchmark-table:ofd-memory:start -->

| 转换器 | hello.ofd | ano.ofd | intro.ofd | 1000-pages.ofd | 999.ofd | zsbk.ofd | 胜出次数 |
|---|---:|---:|---:|---:|---:|---:|---:|
| [go-zc310](https://github.com/zc310/ofd) | FAILED | FAILED | FAILED | FAILED | FAILED | FAILED | 0 |
| [go-ofdgo](https://github.com/xiaoqidun/ofdgo) | FAILED | FAILED | FAILED | FAILED | FAILED | FAILED | 0 |
| [go-ofdrw-rs](https://github.com/lihongjie0209/ofdrw-rs) | FAILED | FAILED | FAILED | FAILED | FAILED | FAILED | 0 |
| [rust-ofdbox-folio](https://github.com/ofdbox/ofdbox-folio) | FAILED | FAILED | FAILED | FAILED | FAILED | FAILED | 0 |
| [rust-sofd](https://github.com/emfox/sOFD) | FAILED | FAILED | FAILED | FAILED | FAILED | FAILED | 0 |
| [easyofd](https://pypi.org/project/easyofd/) | FAILED | FAILED | FAILED | FAILED | FAILED | FAILED | 0 |
| [ofdrw](https://github.com/ofdrw/ofdrw) | FAILED | FAILED | FAILED | FAILED | FAILED | FAILED | 0 |
| [@miconvert/ofd-to-pdf](https://www.npmjs.com/package/@miconvert/ofd-to-pdf) | FAILED | FAILED | FAILED | FAILED | FAILED | FAILED | 0 |
<!-- benchmark-table:ofd-memory:end -->

> - `python-easyofd`: `pip install easyofd`，仅使用 easyofd，不再回退到 ofd2img
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
| [go-zc310](https://github.com/zc310/ofd) | 10.0KB | 89.2KB | 7.5MB | **686.8KB** | 80.3KB | 1.4MB | 1 |
| [go-ofdgo](https://github.com/xiaoqidun/ofdgo) | 15.3KB | 830.6KB | 9.9MB | 57.7MB | 123.2KB | 2.2MB | 0 |
| [go-ofdrw-rs](https://github.com/lihongjie0209/ofdrw-rs) | 5.4KB | 682.8KB | 9.5MB | 25.8MB | 208.7KB | 2.7MB | 0 |
| [rust-ofdbox-folio](https://github.com/ofdbox/ofdbox-folio) | 81.6KB | 258.4KB | 9.0MB | 32.8MB | 602.9KB | 1.9MB | 0 |
| [rust-sofd](https://github.com/emfox/sOFD) | 4.4KB | 971.2KB | 31.4MB | 88.4MB | 834.7KB | 15.9MB | 0 |
| [easyofd](https://pypi.org/project/easyofd/) | 23.6KB | **36.1KB** | 38.2MB | 714.1KB | **73.1KB** | 13.6MB | 2 |
| [ofdrw](https://github.com/ofdrw/ofdrw) | 767.3KB | 2.4MB | 10.9MB | 2.0MB | 1.9MB | 5.0MB | 0 |
| [@miconvert/ofd-to-pdf](https://www.npmjs.com/package/@miconvert/ofd-to-pdf) | **1.6KB** | 196.5KB | **1.1MB** | 9.2MB | 82.9KB | **559.3KB** | 3 |
<!-- benchmark-table:ofd-pdf-size:end -->

### 文本提取能力

<!-- benchmark-table:ofd-text-extraction:start -->

| 转换器 | hello.ofd | ano.ofd | intro.ofd | 999.ofd | zsbk.ofd | 胜出次数 |
|---|---:|---:|---:|---:|---:|---:|
| [go-zc310](https://github.com/zc310/ofd) | 22 | 7232 | 5544 | 3971 | 1162 | 0 |
| [go-ofdgo](https://github.com/xiaoqidun/ofdgo) | 8 | 567 | 506 | 3967 | 527 | 0 |
| [go-ofdrw-rs](https://github.com/lihongjie0209/ofdrw-rs) | 22 | 7093 | 5494 | 3934 | 1150 | 0 |
| [rust-ofdbox-folio](https://github.com/ofdbox/ofdbox-folio) | 8 | 6497 | 1959 | 773 | 717 | 0 |
| [rust-sofd](https://github.com/emfox/sOFD) | 0 | 0 | 0 | 0 | 0 | 0 |
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
| [easyofd](https://pypi.org/project/easyofd/) | FAILED | FAILED | FAILED | FAILED | FAILED | FAILED | 0 |
| [ofdrw](https://github.com/ofdrw/ofdrw) | 99.30% | 99.53% | 99.27% | **99.76%** | 97.33% | 99.06% | 1 |
| [@miconvert/ofd-to-pdf](https://www.npmjs.com/package/@miconvert/ofd-to-pdf) | 99.01% | 97.55% | 96.30% | 97.56% | 95.82% | 94.39% | 0 |
<!-- benchmark-table:ofd-pixel-similarity:end -->

**内容区域像素相似度**（96 DPI）：RGB 任一通道低于 250 视为非白内容。参考内容处的颜色误差与参考内容外的额外输出误差相加，仅除以参考非白像素数；分数为 `max(0, 1 - 归一化误差) × 100%`。不扩边、不裁剪，不允许输出扩大掩码分母稀释误差。双方空白页不参与内容分数平均，全空白文档无内容分数（记 `FAILED`）。此指标有方向性，始终以数科输出为参考；与旧并集扩边分数不可混用。

<!-- benchmark-table:ofd-content-similarity:start -->

| 转换器 | hello.ofd | ano.ofd | intro.ofd | 1000-pages.ofd | 999.ofd | zsbk.ofd | 胜出次数 |
|---|---:|---:|---:|---:|---:|---:|---:|
| [go-zc310](https://github.com/zc310/ofd) | 67.08% | **96.54%** | **99.73%** | 99.39% | 80.52% | **99.41%** | 3 |
| [go-ofdgo](https://github.com/xiaoqidun/ofdgo) | **67.34%** | 91.58% | 99.52% | 99.38% | 80.64% | 98.69% | 1 |
| [go-ofdrw-rs](https://github.com/lihongjie0209/ofdrw-rs) | 65.19% | 91.20% | 99.05% | 98.05% | 80.58% | 99.26% | 0 |
| [rust-ofdbox-folio](https://github.com/ofdbox/ofdbox-folio) | 66.08% | 83.10% | 97.86% | 99.45% | 79.65% | 96.96% | 0 |
| [rust-sofd](https://github.com/emfox/sOFD) | 45.27% | FAILED | FAILED | 99.37% | **81.51%** | 97.42% | 1 |
| [easyofd](https://pypi.org/project/easyofd/) | FAILED | FAILED | FAILED | FAILED | FAILED | FAILED | 0 |
| [ofdrw](https://github.com/ofdrw/ofdrw) | 38.55% | 94.92% | 98.77% | **99.72%** | 64.49% | 98.88% | 1 |
| [@miconvert/ofd-to-pdf](https://www.npmjs.com/package/@miconvert/ofd-to-pdf) | 12.94% | 62.67% | 91.96% | 97.21% | 43.51% | 93.31% | 0 |
<!-- benchmark-table:ofd-content-similarity:end -->

**前景覆盖 F1**：独立衡量形状覆盖，允许约 0.265 mm 的邻域距离（96 DPI 时为 1 像素），兼顾漏画召回率与多画精确率，不替代颜色误差。双方空白页不参与平均。

<!-- benchmark-table:ofd-foreground-f1:start -->

| 转换器 | hello.ofd | ano.ofd | intro.ofd | 1000-pages.ofd | 999.ofd | zsbk.ofd | 胜出次数 |
|---|---:|---:|---:|---:|---:|---:|---:|
| [go-zc310](https://github.com/zc310/ofd) | **86.46%** | 98.30% | 99.96% | 99.98% | 98.09% | **100.00%** | 2 |
| [go-ofdgo](https://github.com/xiaoqidun/ofdgo) | 86.44% | **98.37%** | **100.00%** | 99.97% | 98.10% | **100.00%** | 2 |
| [go-ofdrw-rs](https://github.com/lihongjie0209/ofdrw-rs) | 82.35% | 98.34% | 99.50% | 99.62% | 98.10% | **100.00%** | 0 |
| [rust-ofdbox-folio](https://github.com/ofdbox/ofdbox-folio) | 82.63% | 92.06% | 99.91% | **100.00%** | 97.50% | 65.34% | 1 |
| [rust-sofd](https://github.com/emfox/sOFD) | 67.14% | FAILED | FAILED | **100.00%** | **99.27%** | 99.98% | 1 |
| [easyofd](https://pypi.org/project/easyofd/) | FAILED | FAILED | FAILED | FAILED | FAILED | FAILED | 0 |
| [ofdrw](https://github.com/ofdrw/ofdrw) | 62.07% | 73.37% | 99.76% | **100.00%** | 87.67% | **100.00%** | 0 |
| [@miconvert/ofd-to-pdf](https://www.npmjs.com/package/@miconvert/ofd-to-pdf) | 6.08% | 4.83% | 92.90% | 99.21% | 11.47% | 20.41% | 0 |
<!-- benchmark-table:ofd-foreground-f1:end -->

> 页面显示尺寸相差超过 0.5 pt 时判失败；像素尺寸仅允许 1 像素舍入差异，差异部分用白色补齐。MuPDF 渲染警告或错误判为 `FAILED`，不对不完整渲染给分。
>
> `1000-pages` 因内容重复，全页像素相似度、内容区域像素相似度及前景 F1 仅比较前 3 页，仍校验总页数一致。转换、输出大小及文本比较仍针对完整文档。

**文件大小对比**（转换后 PDF 大小 / 数科参考 PDF 大小）：

<!-- benchmark-table:ofd-size-ratio:start -->

| 转换器 | hello.ofd | ano.ofd | intro.ofd | 1000-pages.ofd | 999.ofd | zsbk.ofd |
|---|---:|---:|---:|---:|---:|---:|
| 数科版式阅读器 | 1.00 | 1.00 | 1.00 | 1.00 | 1.00 | 1.00 |
| [go-zc310](https://github.com/zc310/ofd) | 2.84 | 0.10 | 1.12 | **0.35** | 1.00 | 0.91 |
| [go-ofdgo](https://github.com/xiaoqidun/ofdgo) | 4.36 | 0.92 | 1.48 | 30.36 | 1.54 | 1.43 |
| [go-ofdrw-rs](https://github.com/lihongjie0209/ofdrw-rs) | 1.54 | 0.76 | 1.42 | 13.59 | 2.61 | 1.69 |
| [rust-ofdbox-folio](https://github.com/ofdbox/ofdbox-folio) | 23.19 | 0.29 | 1.35 | 17.26 | 7.54 | 1.21 |
| [rust-sofd](https://github.com/emfox/sOFD) | 1.26 | 1.07 | 4.69 | 46.54 | 10.44 | 10.12 |
| [easyofd](https://pypi.org/project/easyofd/) | 6.70 | **0.04** | 5.70 | 0.37 | **0.92** | 8.62 |
| [ofdrw](https://github.com/ofdrw/ofdrw) | 217.96 | 2.71 | 1.62 | 1.05 | 23.71 | 3.17 |
| [@miconvert/ofd-to-pdf](https://www.npmjs.com/package/@miconvert/ofd-to-pdf) | **0.47** | 0.22 | **0.16** | 4.83 | 1.04 | **0.35** |
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
| [go-zc310](https://github.com/zc310/ofd) | **100.00%** | 99.06% | **99.71%** | **100.00%** | 95.85% | 98.85% | 2 |
| [go-ofdgo](https://github.com/xiaoqidun/ofdgo) | 26.32% | 5.28% | 8.57% | 3.99% | 95.01% | 26.58% | 0 |
| [go-ofdrw-rs](https://github.com/lihongjie0209/ofdrw-rs) | **100.00%** | **99.73%** | 94.88% | **100.00%** | **100.00%** | **100.00%** | 3 |
| [rust-ofdbox-folio](https://github.com/ofdbox/ofdbox-folio) | 26.32% | 18.04% | 3.67% | 3.99% | 39.65% | 0.00% | 0 |
| [rust-sofd](https://github.com/emfox/sOFD) | 0.00% | 0.00% | 0.00% | 0.00% | 0.00% | 0.00% | 0 |
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
