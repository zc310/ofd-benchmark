# OFD → PNG 基准测试（200 DPI）

仅对比可原生指定 **200 DPI** 的 OFD→PNG 渲染，不使用 PDF 中转，不放大低分辨率图片。参考图为数科版式阅读器从 OFD **直接导出的 200 DPI PNG**，位于 `testdata/suwell/<name>/<name>_<页码>.png`。

当前样本：`999`（5 页）、`intro`（42 页）、`zsbk`（2 页）。测试全部页面，包含 `intro` 第 14～19 页；不继承 PDF 评测的缺背景页排除规则。

## 转换器

| 转换器 | 版本 | 路径 |
|---|---|---|
| [go-zc310](https://github.com/zc310/ofd) | v0.2.3-0.20261124113956-09f1815cd2bb | `converter.Image`，PNG / 200 DPI / 白底 |
| [go-ofdgo](https://github.com/xiaoqidun/ofdgo) | v0.0.0-20261124073143-9d9bd2ba2cde | `RenderToPNG`，200 DPI / 包含注释 |
| [rust-ofd-utility](https://github.com/ofd-utility/ofd-utility) | d603791e933a81936b23c91dcab90fdc4c2c0edf | `ofd-cli render --dpi 200 --format png`（200 DPI） |
| [go-ofdrw-rs](https://github.com/lihongjie0209/ofdrw-rs) | v0.5.0 | 原生 `PagePNG(index, 200)`（200 DPI） |
| [rust-ofdbox-folio](https://github.com/ofdbox/ofdbox-folio) | 36e6877bf5d0b989cd824c3c74241533ae1f90ea | 原生 CPU 渲染，200 DPI |
| [rust-sofd](https://github.com/emfox/sOFD) | 65cff468ecc2af676180f540899fdf74da849c3a | 原生 CLI `convert --format png --dpi 200` |
| [java-ofdrw](https://github.com/ofdrw/ofdrw) | 2.5.2 | 原生 AWT `ImageMaker(reader, 200.0 / 25.4)`（200 DPI） |
| [rust-ofdmanager](https://github.com/feuvan/ofdmanager) | a9b1d29ef90141a369ad47cc32fd6c62dd8b8450 | 原生 CLI `render --dpi 200 --page n`（200 DPI） |

选择键为 `zc310`、`ofdgo`、`ofd-utility`、`ofdrw-rs`、`ofdbox-folio`、`sofd`、`java`、`ofdmanager`，不指定时全部运行。

`ofd-utility` 的 Python 适配器仅传递固定参数并统一页码/文件名，不解码或重绘图片。`ofdmanager` 上游 CLI 一次只渲染一页，适配器按 OFD 中的页数循环调用，重复解析/字体加载计入耗时；多文档 OFD 会明确拒绝，不静默只测第一个文档。原生 Java `ImageMaker` 的构造参数是**像素/毫米，不是 DPI**。

sOFD 与现有 PDF 包装器保持一致：仅在返回码 1 且 stderr 完整匹配缺字警告时，保留完整 PNG 继续质量比较，让缺字反映在画质得分中；其他非零退出、缺页及损坏文件仍记失败。

Python `ofd2img 0.1.2` 与 Python easyofd 的图片接口均先转 PDF，再固定以 144 DPI 渲染，故不参加本项评测；Node `@miconvert/ofd-to-pdf` 当前发布包仅有 PDF 输出接口，也不参与 PNG 比较。已有 OFD→PDF 评测保持不变。

## 评测规则

- 耗时包含进程启动、解析、渲染、PNG 编码及写盘；不含后续质量比较。毫秒/页为总耗时除以完整输出页数。PNG 总大小是所有页面文件的字节数之和，受各库默认无损压缩策略影响。
- 输出页数、连续页码必须与参考一致；图片尺寸只允许 1 像素的栅格取整差异，统一在右侧/底部补白，不缩放、不裁切、不平移。
- 比较时透明度统一合成白底并转 RGB。参考图的 DPI 元数据若存在，须在 200±1 范围内（数科导出的约 200.03 属正常取整）。
- **SSIM 结构相似度**：原尺寸 RGB 三通道分别计算 7×7 均匀窗口 SSIM（样本协方差，`K1=0.01`、`K2=0.03`、`data_range=255`），对有效窗口及通道取平均，负分截为 0 后乘 100%。各页等权平均，不缩放、不平移。SSIM 仍可能被大面积白底抬高，应结合内容相似度与前景 F1 阅读。
- **内容相似度**：沿用 PDF 评测的参考前景归一化色差，计入丢失内容和额外墨迹的误差，以参考非白像素数归一化。
- **前景 F1**：允许约 0.265 mm 的物理位置误差，在 200 DPI 对应 2 像素；仅 F1 有容差，颜色比较不移动图像。
- **全页相似度**：`1 − 平均归一化 RGB 色差`。各页等权平均；全白页不抬高内容/F1 得分。丢失几乎全部前景、损坏 PNG、缺页或尺寸不符均记为 `FAILED`。
- 单次转换默认限时 15 分钟，可通过 `BENCHMARK_CONVERTER_TIMEOUT=30m` 调整；超时终止整个进程组并清理部分输出。重跑时清除该转换器旧输出，质量得分重置，避免沿用旧结果。

## 原生 PNG 结果（8 个转换器）

表格由 `docs/benchmark-results.yaml` 生成，勿直接编辑。`FAILED` 表示尚无本轮有效测量或转换/比较失败；首次正式结果由 CI 统一环境测量，本地验证不覆盖已有 PDF 报告的运行环境。

### 总耗时

<!-- benchmark-table:png-speed:start -->

| 转换器 | 999.ofd | intro.ofd | zsbk.ofd | 胜出次数 |
|---|---:|---:|---:|---:|
| [go-zc310](https://github.com/zc310/ofd) | 523ms | 8.70s | 2.40s | 0 |
| [go-ofdgo](https://github.com/xiaoqidun/ofdgo) | 753ms | 12.65s | 2.17s | 0 |
| [rust-ofd-utility](https://github.com/ofd-utility/ofd-utility) | 274ms | FAILED | **103ms** | 1 |
| [go-ofdrw-rs](https://github.com/lihongjie0209/ofdrw-rs) | 6.92s | 30.27s | 5.89s | 0 |
| [rust-ofdbox-folio](https://github.com/ofdbox/ofdbox-folio) | 348ms | 5.60s | 892ms | 0 |
| [rust-sofd](https://github.com/emfox/sOFD) | **199ms** | **2.18s** | 466ms | 2 |
| [java-ofdrw](https://github.com/ofdrw/ofdrw) | 1.58s | 12.30s | 4.08s | 0 |
| [rust-ofdmanager](https://github.com/feuvan/ofdmanager) | 715ms | 9.34s | 1.21s | 0 |
<!-- benchmark-table:png-speed:end -->

### CPU 总时间

CPU 总时间为进程用户态与内核态时间之和，累加所有线程及已等待回收的子进程；包含启动和包装脚本开销。多核并行时可超过墙钟时间。历史数据未采集此指标，FAILED 表示尚未采集或转换失败，需重新运行基准。

<!-- benchmark-table:png-cpu:start -->

| 转换器 | 999.ofd | intro.ofd | zsbk.ofd | 胜出次数 |
|---|---:|---:|---:|---:|
| [go-zc310](https://github.com/zc310/ofd) | 550ms | 8.82s | 2.49s | 0 |
| [go-ofdgo](https://github.com/xiaoqidun/ofdgo) | 787ms | 15.19s | 3.42s | 0 |
| [rust-ofd-utility](https://github.com/ofd-utility/ofd-utility) | 274ms | FAILED | **102ms** | 1 |
| [go-ofdrw-rs](https://github.com/lihongjie0209/ofdrw-rs) | 6.92s | 30.30s | 5.89s | 0 |
| [rust-ofdbox-folio](https://github.com/ofdbox/ofdbox-folio) | 348ms | 5.64s | 930ms | 0 |
| [rust-sofd](https://github.com/emfox/sOFD) | **199ms** | **2.18s** | 466ms | 2 |
| [java-ofdrw](https://github.com/ofdrw/ofdrw) | 4.16s | 17.86s | 7.52s | 0 |
| [rust-ofdmanager](https://github.com/feuvan/ofdmanager) | 714ms | 9.34s | 1.21s | 0 |
<!-- benchmark-table:png-cpu:end -->

### 进程树内存峰值

Linux 下每 10ms 采样转换进程及其存活子孙进程的 RSS（驻留物理内存）之和，取采样峰值，包含包装脚本和运行时；MB 为 100 万字节。共享内存会重复计数，不等于实际独占物理内存；短于采样间隔的峰值、采样期间退出或脱离进程树的进程可能遗漏，读取受限也可能低估。采样带来少量额外开销。历史数据或未采到有效 RSS 显示为 FAILED，需重新运行基准。

<!-- benchmark-table:png-memory:start -->

| 转换器 | 999.ofd | intro.ofd | zsbk.ofd | 胜出次数 |
|---|---:|---:|---:|---:|
| [go-zc310](https://github.com/zc310/ofd) | 228.01MB | 223.04MB | 366.24MB | 0 |
| [go-ofdgo](https://github.com/xiaoqidun/ofdgo) | 318.97MB | 249.76MB | 465.25MB | 0 |
| [rust-ofd-utility](https://github.com/ofd-utility/ofd-utility) | **76.61MB** | FAILED | **57.23MB** | 2 |
| [go-ofdrw-rs](https://github.com/lihongjie0209/ofdrw-rs) | 80.64MB | 214.70MB | 149.81MB | 0 |
| [rust-ofdbox-folio](https://github.com/ofdbox/ofdbox-folio) | 117.01MB | 152.17MB | 218.64MB | 0 |
| [rust-sofd](https://github.com/emfox/sOFD) | 97.39MB | **106.46MB** | 160.04MB | 1 |
| [java-ofdrw](https://github.com/ofdrw/ofdrw) | 399.71MB | 684.61MB | 838.08MB | 0 |
| [rust-ofdmanager](https://github.com/feuvan/ofdmanager) | 163.41MB | 198.24MB | 300.68MB | 0 |
<!-- benchmark-table:png-memory:end -->

### 平均每页耗时

<!-- benchmark-table:png-ms-per-page:start -->

| 转换器 | 999.ofd | intro.ofd | zsbk.ofd | 胜出次数 |
|---|---:|---:|---:|---:|
| [go-zc310](https://github.com/zc310/ofd) | 104.58ms | 207.19ms | 1200.69ms | 0 |
| [go-ofdgo](https://github.com/xiaoqidun/ofdgo) | 150.70ms | 301.12ms | 1083.34ms | 0 |
| [rust-ofd-utility](https://github.com/ofd-utility/ofd-utility) | 54.76ms | FAILED | **51.30ms** | 1 |
| [go-ofdrw-rs](https://github.com/lihongjie0209/ofdrw-rs) | 1383.33ms | 720.71ms | 2942.71ms | 0 |
| [rust-ofdbox-folio](https://github.com/ofdbox/ofdbox-folio) | 69.69ms | 133.35ms | 446.25ms | 0 |
| [rust-sofd](https://github.com/emfox/sOFD) | **39.78ms** | **51.90ms** | 233.11ms | 2 |
| [java-ofdrw](https://github.com/ofdrw/ofdrw) | 315.24ms | 292.89ms | 2040.09ms | 0 |
| [rust-ofdmanager](https://github.com/feuvan/ofdmanager) | 143.03ms | 222.42ms | 606.01ms | 0 |
<!-- benchmark-table:png-ms-per-page:end -->

### PNG 总大小

<!-- benchmark-table:png-size:start -->

| 转换器 | 999.ofd | intro.ofd | zsbk.ofd | 胜出次数 |
|---|---:|---:|---:|---:|
| [go-zc310](https://github.com/zc310/ofd) | 1.64MB | 39.37MB | 6.77MB | 0 |
| [go-ofdgo](https://github.com/xiaoqidun/ofdgo) | 1.69MB | 50.81MB | 8.55MB | 0 |
| [rust-ofd-utility](https://github.com/ofd-utility/ofd-utility) | 2.36MB | FAILED | **335.15KB** | 1 |
| [go-ofdrw-rs](https://github.com/lihongjie0209/ofdrw-rs) | 972.55KB | 40.02MB | 7.35MB | 0 |
| [rust-ofdbox-folio](https://github.com/ofdbox/ofdbox-folio) | 1.41MB | 43.16MB | 3.92MB | 0 |
| [rust-sofd](https://github.com/emfox/sOFD) | 2.40MB | 42.10MB | 9.13MB | 0 |
| [java-ofdrw](https://github.com/ofdrw/ofdrw) | **567.95KB** | **25.65MB** | 6.99MB | 2 |
| [rust-ofdmanager](https://github.com/feuvan/ofdmanager) | 2.39MB | 31.92MB | 8.95MB | 0 |
<!-- benchmark-table:png-size:end -->

### SSIM 结构相似度

新增指标，尚未统一重测的单元格记为 `FAILED`；原有像素分数不改名、不混用。

<!-- benchmark-table:png-ssim:start -->

| 转换器 | 999.ofd | intro.ofd | zsbk.ofd | 胜出次数 |
|---|---:|---:|---:|---:|
| [go-zc310](https://github.com/zc310/ofd) | 96.35% | **98.60%** | **92.27%** | 2 |
| [go-ofdgo](https://github.com/xiaoqidun/ofdgo) | 96.25% | 98.50% | 91.66% | 0 |
| [rust-ofd-utility](https://github.com/ofd-utility/ofd-utility) | 96.69% | FAILED | FAILED | 0 |
| [go-ofdrw-rs](https://github.com/lihongjie0209/ofdrw-rs) | 96.88% | 97.18% | 89.54% | 0 |
| [rust-ofdbox-folio](https://github.com/ofdbox/ofdbox-folio) | 96.71% | 95.73% | 84.31% | 0 |
| [rust-sofd](https://github.com/emfox/sOFD) | 91.43% | 94.71% | 90.07% | 0 |
| [java-ofdrw](https://github.com/ofdrw/ofdrw) | 95.36% | 85.64% | 91.07% | 0 |
| [rust-ofdmanager](https://github.com/feuvan/ofdmanager) | **96.91%** | FAILED | 90.44% | 1 |
<!-- benchmark-table:png-ssim:end -->

### 内容相似度

<!-- benchmark-table:png-content-similarity:start -->

| 转换器 | 999.ofd | intro.ofd | zsbk.ofd | 胜出次数 |
|---|---:|---:|---:|---:|
| [go-zc310](https://github.com/zc310/ofd) | 66.48% | **99.13%** | 97.86% | 1 |
| [go-ofdgo](https://github.com/xiaoqidun/ofdgo) | 67.44% | 99.09% | 97.88% | 0 |
| [rust-ofd-utility](https://github.com/ofd-utility/ofd-utility) | 68.58% | FAILED | FAILED | 0 |
| [go-ofdrw-rs](https://github.com/lihongjie0209/ofdrw-rs) | 69.65% | 98.36% | 97.74% | 0 |
| [rust-ofdbox-folio](https://github.com/ofdbox/ofdbox-folio) | 68.87% | 95.01% | 95.57% | 0 |
| [rust-sofd](https://github.com/emfox/sOFD) | 26.75% | 96.09% | 97.45% | 0 |
| [java-ofdrw](https://github.com/ofdrw/ofdrw) | 58.56% | 94.56% | **97.93%** | 1 |
| [rust-ofdmanager](https://github.com/feuvan/ofdmanager) | **69.80%** | FAILED | 97.81% | 1 |
<!-- benchmark-table:png-content-similarity:end -->

### 前景 F1

<!-- benchmark-table:png-foreground-f1:start -->

| 转换器 | 999.ofd | intro.ofd | zsbk.ofd | 胜出次数 |
|---|---:|---:|---:|---:|
| [go-zc310](https://github.com/zc310/ofd) | **100.00%** | **100.00%** | **100.00%** | 3 |
| [go-ofdgo](https://github.com/xiaoqidun/ofdgo) | **100.00%** | **100.00%** | **100.00%** | 0 |
| [rust-ofd-utility](https://github.com/ofd-utility/ofd-utility) | 99.20% | FAILED | FAILED | 0 |
| [go-ofdrw-rs](https://github.com/lihongjie0209/ofdrw-rs) | **100.00%** | 99.62% | **100.00%** | 0 |
| [rust-ofdbox-folio](https://github.com/ofdbox/ofdbox-folio) | 99.07% | 99.67% | 63.37% | 0 |
| [rust-sofd](https://github.com/emfox/sOFD) | 95.92% | 99.71% | 99.99% | 0 |
| [java-ofdrw](https://github.com/ofdrw/ofdrw) | 99.73% | 99.97% | **100.00%** | 0 |
| [rust-ofdmanager](https://github.com/feuvan/ofdmanager) | **100.00%** | FAILED | **100.00%** | 0 |
<!-- benchmark-table:png-foreground-f1:end -->

### 全页相似度

<!-- benchmark-table:png-pixel-similarity:start -->

| 转换器 | 999.ofd | intro.ofd | zsbk.ofd | 胜出次数 |
|---|---:|---:|---:|---:|
| [go-zc310](https://github.com/zc310/ofd) | 98.82% | **99.71%** | 98.21% | 1 |
| [go-ofdgo](https://github.com/xiaoqidun/ofdgo) | 98.80% | 99.69% | 98.23% | 0 |
| [rust-ofd-utility](https://github.com/ofd-utility/ofd-utility) | 98.86% | FAILED | FAILED | 0 |
| [go-ofdrw-rs](https://github.com/lihongjie0209/ofdrw-rs) | 98.92% | 99.36% | 98.10% | 0 |
| [rust-ofdbox-folio](https://github.com/ofdbox/ofdbox-folio) | 98.88% | 97.15% | 96.30% | 0 |
| [rust-sofd](https://github.com/emfox/sOFD) | 97.44% | 96.97% | 97.87% | 0 |
| [java-ofdrw](https://github.com/ofdrw/ofdrw) | 98.55% | 95.74% | **98.27%** | 1 |
| [rust-ofdmanager](https://github.com/feuvan/ofdmanager) | **98.93%** | FAILED | 98.16% | 1 |
<!-- benchmark-table:png-pixel-similarity:end -->

## 运行

先准备项目字体环境及 `python/requirements.txt` 中的 Python 依赖（见 README）。

```bash
GOWORK=off go build -o bin/ofd2png-benchmark ./cmd/ofd2png-benchmark
GOWORK=off go build -o bin/zc310-png ./cmd/zc310-png
GOWORK=off go build -o bin/ofdgo-png ./cmd/ofdgo-png
GOWORK=off go build -o bin/ofdrw-rs-png ./cmd/ofdrw-rs-png
cargo build --release --locked --manifest-path rust/ofdbox-folio-convert/Cargo.toml --target-dir rust/target
cp rust/target/release/ofdbox-folio-convert bin/
# Java 依赖按 java/pom.xml 安装。排除本地历史损坏占位 JAR。
CP=$(find java/lib -name '*.jar' ! -name 'jaxen.jar' ! -name 'pdfbox-sandbox.jar' ! -name 'sqlite.jar' -printf '%p:')
javac -cp "$CP" java/Ofd2Png.java

# Rust 1.93+；固定源码 revision，使用上游 Cargo.lock。
cargo install --git https://github.com/ofd-utility/ofd-utility \
  --rev d603791e933a81936b23c91dcab90fdc4c2c0edf \
  --locked --bin ofd-cli --root /tmp/opencode/ofd-utility-install ofd-cli
cp /tmp/opencode/ofd-utility-install/bin/ofd-cli bin/ofd-utility-cli
cargo install --git https://github.com/feuvan/ofdmanager \
  --rev a9b1d29ef90141a369ad47cc32fd6c62dd8b8450 \
  --locked --bin ofd-cli --root /tmp/opencode/ofdmanager-install ofd-cli
cp /tmp/opencode/ofdmanager-install/bin/ofd-cli bin/ofdmanager-cli
# sOFD CLI、ofdrw-rs 原生库按 README 的现有构建说明准备。

# 全部 8 个原生转换器，固定 200 DPI。
for name in 999 intro zsbk; do
  ./bin/ofd2png-benchmark "testdata/ofdrw/$name.ofd"
done
python3 tools/compare_png.py
GOWORK=off go run ./tools/gen-benchmarks
GOWORK=off go run ./tools/gen-benchmarks -check

# 单独选择转换器；质量比较选择键使用完整输出名称。
./bin/ofd2png-benchmark testdata/ofdrw/999.ofd ofd-utility
python3 tools/compare_png.py 999 --converters rust-ofd-utility
```

输出保存为 `bin/output_png/<name>_<converter>/page_<页码>.png`。仅查看质量、不修改结果 YAML 可加 `--no-write`。

## 运行环境

下列环境与现有报告共用，由 CI 统一记录。

<!-- benchmark-environment:start -->

- 系统：Ubuntu 24.04.5 LTS (x86_64)
- CPU：Intel(R) Xeon(R) 6975P-C（4 逻辑核心）
- 内存：15.3 GiB
<!-- benchmark-environment:end -->
