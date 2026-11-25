# OFD → PDF 对比测试

> [OFD 库汇总](ofd-libraries.md)：各 OFD 库能力 / 许可证对比表
>
> [PDF → OFD 对比测试](pdf-to-ofd.md)：反方向 PDF→OFD 转换对比

## 测试库

### Go 库

| 库                                                    | 版本                                  | OFD→PDF |  更新时间  | 备注                                    |
|-------------------------------------------------------|---------------------------------------|:-------:|:----------:|-----------------------------------------|
| [zc310-ofd](https://github.com/zc310/ofd)             | v0.2.3-0.20261124113956-09f1815cd2bb |   ✅    | 2026-11-14 | 已测试                                  |
| [xiaoqidun-ofdgo](https://github.com/xiaoqidun/ofdgo) | v0.0.0-20261124073143-9d9bd2ba2cde |   ✅    | 2026-11-14 | 已测试                                  |
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
| [ofdrw](https://github.com/ofdrw/ofdrw) | 2.5.2 |   ✅    | 2026-11-14 | 需手动下载依赖，全部成功 |

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

> 所有转换器均由 GitHub Actions 在上述环境统一评测。转换耗时、文件大小、文本提取及像素比较分别记录。

## 测试文件

测试用 OFD 文件来自：https://github.com/zc310/ofd-fixtures

| 文件           | 大小 | 说明         |
|----------------|------|--------------|
| hello.ofd      | 1.50K | 简单测试文件 |
| ano.ofd        | 702.00K | 带签章文件   |
| intro.ofd      | 7.20M | 介绍文档     |
| 1000-pages.ofd | 508.00K | 1000页文件   |
| 999.ofd        | 30.00K  | 带签章文件   |
| zsbk.ofd       | 1.50M | 数科签章文件 |

> `testdata/suwell/` 下的 `*.pdf` 为各 OFD 的参考转换结果，由**数科版式阅读器**导出，用作像素级对比基准。

## 测试结果

### 转换速度

<!-- benchmark-table:ofd-speed:start -->

| 转换器 | hello.ofd | ano.ofd | intro.ofd | 1000-pages.ofd | 999.ofd | zsbk.ofd | 胜出次数 |
|---|---:|---:|---:|---:|---:|---:|---:|
| [go-zc310](https://github.com/zc310/ofd) | 87ms | 172ms | **343ms** | 164ms | **218ms** | 339ms | 2 |
| [go-ofdgo](https://github.com/xiaoqidun/ofdgo) | 1.99s | 417ms | 360ms | **150ms** | 665ms | 574ms | 1 |
| [go-ofdrw-rs](https://github.com/lihongjie0209/ofdrw-rs) | **31ms** | 535ms | 2.84s | 265.31s | 493ms | 563ms | 1 |
| [rust-ofdbox-folio](https://github.com/ofdbox/ofdbox-folio) | 114ms | **143ms** | 451ms | 5.31s | 338ms | **273ms** | 2 |
| [rust-sofd](https://github.com/emfox/sOFD) | 37ms | 344ms | 2.68s | 21.87s | 294ms | 2.91s | 0 |
| [easyofd](https://pypi.org/project/easyofd/) | 2.65s | 1.08s | 9.65s | 1.59s | 1.02s | 5.63s | 0 |
| [ofdrw](https://github.com/ofdrw/ofdrw) | 1.38s | 1.18s | 3.02s | 3.03s | 1.96s | 1.83s | 0 |
| [@miconvert/ofd-to-pdf](https://www.npmjs.com/package/@miconvert/ofd-to-pdf) | 264ms | 595ms | 1.53s | 9.12s | 443ms | 933ms | 0 |
<!-- benchmark-table:ofd-speed:end -->

### CPU 总时间

CPU 总时间为进程用户态与内核态时间之和，累加所有线程及已等待回收的子进程；包含启动和包装脚本开销。多核并行时可超过墙钟时间。历史数据未采集此指标，FAILED 表示尚未采集或转换失败，需重新运行基准。

<!-- benchmark-table:ofd-cpu:start -->

| 转换器 | hello.ofd | ano.ofd | intro.ofd | 1000-pages.ofd | 999.ofd | zsbk.ofd | 胜出次数 |
|---|---:|---:|---:|---:|---:|---:|---:|
| [go-zc310](https://github.com/zc310/ofd) | 99ms | 183ms | **391ms** | **195ms** | **245ms** | 373ms | 3 |
| [go-ofdgo](https://github.com/xiaoqidun/ofdgo) | 203ms | 518ms | 1.17s | 483ms | 1.42s | 673ms | 0 |
| [go-ofdrw-rs](https://github.com/lihongjie0209/ofdrw-rs) | **15ms** | 530ms | 2.84s | 265.03s | 487ms | 563ms | 1 |
| [rust-ofdbox-folio](https://github.com/ofdbox/ofdbox-folio) | 114ms | **143ms** | 451ms | 5.31s | 338ms | **273ms** | 2 |
| [rust-sofd](https://github.com/emfox/sOFD) | 37ms | 344ms | 2.68s | 21.87s | 294ms | 2.91s | 0 |
| [easyofd](https://pypi.org/project/easyofd/) | 959ms | 1.08s | 9.65s | 1.59s | 1.02s | 5.63s | 0 |
| [ofdrw](https://github.com/ofdrw/ofdrw) | 1.78s | 3.46s | 7.78s | 9.28s | 4.25s | 5.29s | 0 |
| [@miconvert/ofd-to-pdf](https://www.npmjs.com/package/@miconvert/ofd-to-pdf) | 431ms | 1.29s | 2.58s | 14.15s | 897ms | 1.72s | 0 |
<!-- benchmark-table:ofd-cpu:end -->

### 进程树内存峰值

Linux 下每 10ms 采样转换进程及其存活子孙进程的 RSS（驻留物理内存）之和，取采样峰值，包含包装脚本和运行时；MB 为 100 万字节。共享内存会重复计数，不等于实际独占物理内存；短于采样间隔的峰值、采样期间退出或脱离进程树的进程可能遗漏，读取受限也可能低估。采样带来少量额外开销。历史数据或未采到有效 RSS 显示为 FAILED，需重新运行基准。

<!-- benchmark-table:ofd-memory:start -->

| 转换器 | hello.ofd | ano.ofd | intro.ofd | 1000-pages.ofd | 999.ofd | zsbk.ofd | 胜出次数 |
|---|---:|---:|---:|---:|---:|---:|---:|
| [go-zc310](https://github.com/zc310/ofd) | 84.37MB | 144.20MB | 146.05MB | **31.09MB** | 180.45MB | 159.87MB | 1 |
| [go-ofdgo](https://github.com/xiaoqidun/ofdgo) | 132.46MB | 314.14MB | 152.04MB | 49.17MB | 776.68MB | 285.37MB | 0 |
| [go-ofdrw-rs](https://github.com/lihongjie0209/ofdrw-rs) | **18.94MB** | 61.09MB | 99.73MB | 208.07MB | 66.49MB | 124.00MB | 1 |
| [rust-ofdbox-folio](https://github.com/ofdbox/ofdbox-folio) | 35.38MB | 32.33MB | **66.45MB** | 236.19MB | 68.65MB | **46.03MB** | 2 |
| [rust-sofd](https://github.com/emfox/sOFD) | 23.43MB | **16.28MB** | 103.87MB | 346.17MB | **56.51MB** | 143.82MB | 2 |
| [easyofd](https://pypi.org/project/easyofd/) | 214.32MB | 219.17MB | 358.78MB | 237.13MB | 227.54MB | 372.01MB | 0 |
| [ofdrw](https://github.com/ofdrw/ofdrw) | 153.43MB | 388.49MB | 722.91MB | 554.79MB | 480.03MB | 520.87MB | 0 |
| [@miconvert/ofd-to-pdf](https://www.npmjs.com/package/@miconvert/ofd-to-pdf) | 101.49MB | 146.08MB | 263.09MB | 1075.13MB | 132.06MB | 187.20MB | 0 |
<!-- benchmark-table:ofd-memory:end -->

> - `python-easyofd`: `pip install easyofd`，仅使用 easyofd，不再回退到 ofd2img
> - `java-ofdrw`: `org.ofdrw:ofdrw-converter:2.5.2`
> - `node-ofd2pdf`: `npm install @miconvert/ofd-to-pdf`，优先使用 CJS 入口，失败后回退 ESM 入口
> - `rust-ofdbox-folio`: 固定 Git revision `36e6877`；字体由 `fc-match` 按系统 fontconfig 配置解析，不依赖特定用户目录
> - `rust-sofd`: 固定 Git revision `65cff46`；字体由 font-kit/fontconfig 从系统配置解析；缺字警告时 CLI 返回非零，但包装器保留有效 PDF 继续视觉评测
>
> 各转换器的成功或失败情况以本轮结果表为准。

### PDF 文件大小

<!-- benchmark-table:ofd-pdf-size:start -->

| 转换器 | hello.ofd | ano.ofd | intro.ofd | 1000-pages.ofd | 999.ofd | zsbk.ofd | 胜出次数 |
|---|---:|---:|---:|---:|---:|---:|---:|
| [go-zc310](https://github.com/zc310/ofd) | 10.02KB | 90.54KB | 7.46MB | 683.92KB | 81.90KB | 1.44MB | 0 |
| [go-ofdgo](https://github.com/xiaoqidun/ofdgo) | 4.88KB | 56.05KB | 7.56MB | **295.71KB** | **63.12KB** | 1.53MB | 2 |
| [go-ofdrw-rs](https://github.com/lihongjie0209/ofdrw-rs) | 5.41KB | 682.85KB | 9.54MB | 25.81MB | 208.70KB | 2.66MB | 0 |
| [rust-ofdbox-folio](https://github.com/ofdbox/ofdbox-folio) | 81.64KB | 258.43KB | 9.03MB | 32.78MB | 602.89KB | 1.90MB | 0 |
| [rust-sofd](https://github.com/emfox/sOFD) | 4.45KB | 971.23KB | 31.44MB | 88.40MB | 834.72KB | 15.91MB | 0 |
| [easyofd](https://pypi.org/project/easyofd/) | 23.58KB | **36.06KB** | 38.22MB | 714.14KB | 73.14KB | 13.55MB | 1 |
| [ofdrw](https://github.com/ofdrw/ofdrw) | 767.34KB | 2.39MB | 10.87MB | 2.00MB | 1.85MB | 4.98MB | 0 |
| [@miconvert/ofd-to-pdf](https://www.npmjs.com/package/@miconvert/ofd-to-pdf) | **1.64KB** | 196.51KB | **1.08MB** | 9.18MB | 82.92KB | **559.27KB** | 3 |
<!-- benchmark-table:ofd-pdf-size:end -->

### 提取文本字符数

<!-- benchmark-table:ofd-text-extraction:start -->

| 转换器 | hello.ofd | ano.ofd | intro.ofd | 1000-pages.ofd | 999.ofd | zsbk.ofd | 胜出次数 |
|---|---:|---:|---:|---:|---:|---:|---:|
| [go-zc310](https://github.com/zc310/ofd) | 22 | 7232 | 5544 | 242999 | 3971 | 1162 | 0 |
| [go-ofdgo](https://github.com/xiaoqidun/ofdgo) | 22 | 7091 | 5543 | 242999 | 3971 | 1150 | 0 |
| [go-ofdrw-rs](https://github.com/lihongjie0209/ofdrw-rs) | 22 | 7093 | 5494 | 242999 | 3934 | 1150 | 0 |
| [rust-ofdbox-folio](https://github.com/ofdbox/ofdbox-folio) | 8 | 6497 | 1959 | 129999 | 773 | 717 | 0 |
| [rust-sofd](https://github.com/emfox/sOFD) | 0 | 0 | 0 | 0 | 0 | 0 | 0 |
| [easyofd](https://pypi.org/project/easyofd/) | 22 | 1208 | **10194** | 20223 | 3933 | 1508 | 1 |
| [ofdrw](https://github.com/ofdrw/ofdrw) | 26 | **7282** | 5550 | 242999 | 4162 | 1169 | 1 |
| [@miconvert/ofd-to-pdf](https://www.npmjs.com/package/@miconvert/ofd-to-pdf) | **31** | 6489 | 9614 | **325999** | **4249** | **1938** | 4 |
<!-- benchmark-table:ofd-text-extraction:end -->

> - 单位为字符数（chars）
> - 本轮使用 PyMuPDF 1.26.7 与 NumPy 2.3.5；表格数值为本轮测试结果
> - `FAILED` 表示暂无有效结果；`1000-pages` 的字符数由 CI 对完整文档提取后更新，不使用像素比较的前 3 页采样。

### 与数科版式输出对比

以数科版式阅读器导出的 `http://github.com/zc310/ofd-fixtures/tree/main/testdata/suwell` 为基准，对比转换效果。

**全页像素相似度**（96 DPI，1 - 平均归一化色差，100% 为完全相同；白底可能稀释内容差异）：

<!-- benchmark-table:ofd-pixel-similarity:start -->

| 转换器                                                                       |    hello.ofd |      ano.ofd |    intro.ofd | 1000-pages.ofd |      999.ofd |     zsbk.ofd | 胜出次数 |
| ---------------------------------------------------------------------------- | -----------: | -----------: | -----------: | -------------: | -----------: | -----------: | -------: |
| [go-zc310](https://github.com/zc310/ofd)                                     | **99.6300%** |     99.7000% | **99.9300%** |       99.4700% |     98.5500% | **99.5100%** |        3 |
| [go-ofdgo](https://github.com/xiaoqidun/ofdgo)                               |     99.6200% | **99.7100%** |     99.9000% |       99.4700% |     98.5600% |     99.3700% |        1 |
| [go-ofdrw-rs](https://github.com/lihongjie0209/ofdrw-rs)                     |     99.6000% |     99.3800% |     99.5400% |       98.3000% |     98.5500% |     99.3800% |        0 |
| [rust-ofdbox-folio](https://github.com/ofdbox/ofdbox-folio)                  |     99.6100% |     98.8900% |     99.4600% |       99.5200% |     98.4800% |     97.4600% |        0 |
| [rust-sofd](https://github.com/emfox/sOFD)                                   |     99.3800% |       FAILED |       FAILED |       99.4500% | **98.5900%** |     97.8400% |        1 |
| [easyofd](https://pypi.org/project/easyofd/)                                 |       FAILED |       FAILED |       FAILED |         FAILED |       FAILED |       FAILED |        0 |
| [ofdrw](https://github.com/ofdrw/ofdrw)                                      |     99.3000% |     99.5300% |     99.2700% |   **99.7600%** |     97.3300% |     99.0600% |        1 |
| [@miconvert/ofd-to-pdf](https://www.npmjs.com/package/@miconvert/ofd-to-pdf) |     99.0100% |     97.5500% |     96.3000% |       97.5600% |     95.8200% |     94.3900% |        0 |
<!-- benchmark-table:ofd-pixel-similarity:end -->

**SSIM 结构相似度**（96 DPI）：原尺寸 RGB 三通道分别计算 7×7 均匀窗口 SSIM，使用样本协方差、`K1=0.01`、`K2=0.03`、`data_range=255`；对有效窗口及通道取平均，负分截为 0 后乘 100%，再对页面等权平均。不缩放、不平移，图片至少为 7×7 像素。白底仍可能抬高 SSIM，应结合内容相似度及前景 F1 阅读。新增指标，尚未统一重测的单元格记为 `FAILED`，不将旧像素分数当作 SSIM。

<!-- benchmark-table:ofd-ssim:start -->

| 转换器                                                                       |    hello.ofd |      ano.ofd |    intro.ofd | 1000-pages.ofd |      999.ofd |     zsbk.ofd | 胜出次数 |
| ---------------------------------------------------------------------------- | -----------: | -----------: | -----------: | -------------: | -----------: | -----------: | -------: |
| [go-zc310](https://github.com/zc310/ofd)                                     |     98.4100% |     96.4000% | **99.9200%** |       98.9300% |     94.6100% | **98.2800%** |        2 |
| [go-ofdgo](https://github.com/xiaoqidun/ofdgo)                               |     98.4200% | **96.5100%** |     99.7900% |       98.9400% |     94.6300% |     97.9200% |        1 |
| [go-ofdrw-rs](https://github.com/lihongjie0209/ofdrw-rs)                     |     98.4500% |     94.7400% |     98.2900% |       94.0100% |     94.5800% |     97.9600% |        0 |
| [rust-ofdbox-folio](https://github.com/ofdbox/ofdbox-folio)                  | **98.5100%** |     91.9000% |     97.6100% |       99.2600% |     94.1500% |     92.1500% |        1 |
| [rust-sofd](https://github.com/emfox/sOFD)                                   |     98.0400% |       FAILED |       FAILED |       98.9400% | **95.7400%** |     81.2500% |        1 |
| [easyofd](https://pypi.org/project/easyofd/)                                 |       FAILED |       FAILED |       FAILED |         FAILED |       FAILED |       FAILED |        0 |
| [ofdrw](https://github.com/ofdrw/ofdrw)                                      |     97.6800% |     92.2900% |     96.4100% |   **99.7300%** |     88.8900% |     96.0000% |        1 |
| [@miconvert/ofd-to-pdf](https://www.npmjs.com/package/@miconvert/ofd-to-pdf) |     96.6300% |     73.1700% |     89.2100% |       90.1200% |     73.7300% |     79.3200% |        0 |
<!-- benchmark-table:ofd-ssim:end -->

**内容区域像素相似度**（96 DPI）：RGB 任一通道低于 250 视为非白内容。参考内容处的颜色误差与参考内容外的额外输出误差相加，仅除以参考非白像素数；分数为 `max(0, 1 - 归一化误差) × 100%`。不扩边、不裁剪，不允许输出扩大掩码分母稀释误差。双方空白页不参与内容分数平均，全空白文档无内容分数（记 `FAILED`）。此指标有方向性，始终以数科输出为参考；与旧并集扩边分数不可混用。

<!-- benchmark-table:ofd-content-similarity:start -->

| 转换器                                                                       |    hello.ofd |      ano.ofd |    intro.ofd | 1000-pages.ofd |      999.ofd |     zsbk.ofd | 胜出次数 |
| ---------------------------------------------------------------------------- | -----------: | -----------: | -----------: | -------------: | -----------: | -----------: | -------: |
| [go-zc310](https://github.com/zc310/ofd)                                     | **67.0800%** |     96.5400% | **99.8800%** |       99.3900% |     80.5200% | **99.4100%** |        3 |
| [go-ofdgo](https://github.com/xiaoqidun/ofdgo)                               |     67.0200% | **96.7400%** |     99.8300% |       99.4000% |     80.6300% |     99.2400% |        1 |
| [go-ofdrw-rs](https://github.com/lihongjie0209/ofdrw-rs)                     |     65.1900% |     91.2000% |     99.0500% |       98.0500% |     80.5800% |     99.2600% |        0 |
| [rust-ofdbox-folio](https://github.com/ofdbox/ofdbox-folio)                  |     66.0800% |     83.1000% |     97.8600% |       99.4500% |     79.6500% |     96.9600% |        0 |
| [rust-sofd](https://github.com/emfox/sOFD)                                   |     45.2700% |       FAILED |       FAILED |       99.3700% | **81.5100%** |     97.4200% |        1 |
| [easyofd](https://pypi.org/project/easyofd/)                                 |       FAILED |       FAILED |       FAILED |         FAILED |       FAILED |       FAILED |        0 |
| [ofdrw](https://github.com/ofdrw/ofdrw)                                      |     38.5500% |     94.9200% |     98.7700% |   **99.7200%** |     64.4900% |     98.8800% |        1 |
| [@miconvert/ofd-to-pdf](https://www.npmjs.com/package/@miconvert/ofd-to-pdf) |     12.9400% |     62.6700% |     91.9600% |       97.2100% |     43.5100% |     93.3100% |        0 |
<!-- benchmark-table:ofd-content-similarity:end -->

**前景覆盖 F1**：独立衡量形状覆盖，允许约 0.265 mm 的邻域距离（96 DPI 时为 1 像素），兼顾漏画召回率与多画精确率，不替代颜色误差。双方空白页不参与平均。

<!-- benchmark-table:ofd-foreground-f1:start -->

| 转换器                                                                       |    hello.ofd |      ano.ofd |     intro.ofd | 1000-pages.ofd |      999.ofd |      zsbk.ofd | 胜出次数 |
| ---------------------------------------------------------------------------- | -----------: | -----------: | ------------: | -------------: | -----------: | ------------: | -------: |
| [go-zc310](https://github.com/zc310/ofd)                                     |     86.4600% |     98.3000% | **100.0000%** |       99.9800% |     98.0900% | **100.0000%** |        2 |
| [go-ofdgo](https://github.com/xiaoqidun/ofdgo)                               | **86.5100%** | **98.3900%** | **100.0000%** |       99.9800% |     98.1000% | **100.0000%** |        2 |
| [go-ofdrw-rs](https://github.com/lihongjie0209/ofdrw-rs)                     |     82.3500% |     98.3400% |      99.5000% |       99.6200% |     98.1000% | **100.0000%** |        0 |
| [rust-ofdbox-folio](https://github.com/ofdbox/ofdbox-folio)                  |     82.6300% |     92.0600% |      99.9100% |  **100.0000%** |     97.5000% |      65.3400% |        1 |
| [rust-sofd](https://github.com/emfox/sOFD)                                   |     67.1400% |       FAILED |        FAILED |  **100.0000%** | **99.2700%** |      99.9800% |        1 |
| [easyofd](https://pypi.org/project/easyofd/)                                 |       FAILED |       FAILED |        FAILED |         FAILED |       FAILED |        FAILED |        0 |
| [ofdrw](https://github.com/ofdrw/ofdrw)                                      |     62.0700% |     73.3700% |      99.7600% |  **100.0000%** |     87.6700% | **100.0000%** |        0 |
| [@miconvert/ofd-to-pdf](https://www.npmjs.com/package/@miconvert/ofd-to-pdf) |      6.0800% |      4.8300% |      92.9000% |       99.2100% |     11.4700% |      20.4100% |        0 |
<!-- benchmark-table:ofd-foreground-f1:end -->

> 页面显示尺寸相差超过 0.5 pt 时判失败；像素尺寸仅允许 1 像素舍入差异，差异部分用白色补齐。MuPDF 渲染警告或错误判为 `FAILED`，不对不完整渲染给分。
>
> `1000-pages` 因内容重复，全页像素相似度、SSIM、内容区域像素相似度及前景 F1 仅比较前 3 页，仍校验总页数一致。转换、输出大小及文本比较仍针对完整文档。

**文件大小对比**（转换后 PDF 大小 / 数科参考 PDF 大小）：

<!-- benchmark-table:ofd-size-ratio:start -->

| 转换器 | hello.ofd | ano.ofd | intro.ofd | 1000-pages.ofd | 999.ofd | zsbk.ofd |
|---|---:|---:|---:|---:|---:|---:|
| 数科版式阅读器 | 1.00 | 1.00 | 1.00 | 1.00 | 1.00 | 1.00 |
| [go-zc310](https://github.com/zc310/ofd) | 2.85 | 0.10 | 1.11 | 0.35 | 1.02 | 0.92 |
| [go-ofdgo](https://github.com/xiaoqidun/ofdgo) | 1.39 | 0.06 | 1.13 | **0.15** | **0.79** | 0.97 |
| [go-ofdrw-rs](https://github.com/lihongjie0209/ofdrw-rs) | 1.54 | 0.76 | 1.42 | 13.59 | 2.61 | 1.69 |
| [rust-ofdbox-folio](https://github.com/ofdbox/ofdbox-folio) | 23.19 | 0.29 | 1.35 | 17.26 | 7.54 | 1.21 |
| [rust-sofd](https://github.com/emfox/sOFD) | 1.26 | 1.07 | 4.69 | 46.54 | 10.44 | 10.12 |
| [easyofd](https://pypi.org/project/easyofd/) | 6.70 | **0.04** | 5.70 | 0.37 | 0.92 | 8.62 |
| [ofdrw](https://github.com/ofdrw/ofdrw) | 217.96 | 2.71 | 1.62 | 1.05 | 23.71 | 3.17 |
| [@miconvert/ofd-to-pdf](https://www.npmjs.com/package/@miconvert/ofd-to-pdf) | **0.47** | 0.22 | **0.16** | 4.83 | 1.04 | **0.35** |
<!-- benchmark-table:ofd-size-ratio:end -->

> - 仿色、抗锯齿等渲染差异会影响像素相似度
> - 若参考页有可见内容而输出页几乎全白，像素对比会标记为 `FAILED`，避免大面积白底虚高
> - intro 的像素对比跳过 14～19 页（数科导出的参考 PDF 这几页背景图丢失，不计入对比）
> - 各转换器失败情况以结果表为准；ofd2img、python-ofd2pdf 已移出评测
> - 复现命令：`python3 tools/compare_pdf.py 96 <file>`

**文本相似度**（与数科参考 PDF 提取文本的 4-gram Jaccard 相似度，100% 为文本完全一致；`0.0000%` 表示无文本或提取失败）：

参考文本长度（fitz 1.26.7 口径）：`hello=22` · `ano=7091` · `intro=5543` · `1000-pages=263999` · `999=3934` · `zsbk=1150`

<!-- benchmark-table:ofd-text-similarity:start -->

| 转换器                                                                       |     hello.ofd |       ano.ofd |    intro.ofd | 1000-pages.ofd |       999.ofd |      zsbk.ofd | 胜出次数 |
| ---------------------------------------------------------------------------- | ------------: | ------------: | -----------: | -------------: | ------------: | ------------: | -------: |
| [go-zc310](https://github.com/zc310/ofd)                                     | **100.0000%** |      99.0600% | **99.7100%** |  **100.0000%** |      95.8500% |      98.8500% |        2 |
| [go-ofdgo](https://github.com/xiaoqidun/ofdgo)                               | **100.0000%** | **100.0000%** |     99.6000% |  **100.0000%** |      95.8500% | **100.0000%** |        2 |
| [go-ofdrw-rs](https://github.com/lihongjie0209/ofdrw-rs)                     | **100.0000%** |      99.7300% |     94.8800% |  **100.0000%** | **100.0000%** | **100.0000%** |        1 |
| [rust-ofdbox-folio](https://github.com/ofdbox/ofdbox-folio)                  |      26.3200% |      18.0400% |      3.6700% |        3.9900% |      39.6500% |       0.0000% |        0 |
| [rust-sofd](https://github.com/emfox/sOFD)                                   |       0.0000% |       0.0000% |      0.0000% |        0.0000% |       0.0000% |       0.0000% |        0 |
| [easyofd](https://pypi.org/project/easyofd/)                                 | **100.0000%** |      12.5700% |      0.0600% |       98.3800% |      80.5400% |      11.2000% |        0 |
| [ofdrw](https://github.com/ofdrw/ofdrw)                                      |      17.1400% |      70.5800% |     38.2800% |  **100.0000%** |      56.7800% |      64.7200% |        1 |
| [@miconvert/ofd-to-pdf](https://www.npmjs.com/package/@miconvert/ofd-to-pdf) |      10.8100% |      97.3300% |      5.8600% |       83.3100% |      41.2400% |       4.8600% |        0 |
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
