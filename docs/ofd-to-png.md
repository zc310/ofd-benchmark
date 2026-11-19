# OFD → PNG 基准测试（200 DPI）

仅对比可原生指定 **200 DPI** 的 OFD→PNG 渲染，不使用 PDF 中转，不放大低分辨率图片。参考图为数科版式阅读器从 OFD **直接导出的 200 DPI PNG**，位于 `testdata/suwell/<name>/<name>_<页码>.png`。

当前样本：`999`（5 页）、`intro`（42 页）、`zsbk`（2 页）。测试全部页面，包含 `intro` 第 14～19 页；不继承 PDF 评测的缺背景页排除规则。

## 转换器

| 转换器 | 版本 | 路径 |
|---|---|---|
| [go-zc310](https://github.com/zc310/ofd) | v0.2.1 | `converter.Image`，PNG / 200 DPI / 白底 |
| [go-ofdgo](https://github.com/xiaoqidun/ofdgo) | v0.0.0-20261118174012-133e0b6cbf8d | `RenderToPNG`，200 DPI / 包含注释 |
| [rust-ofd-utility](https://github.com/ofd-utility/ofd-utility) | d603791e933a81936b23c91dcab90fdc4c2c0edf | `ofd-cli render --dpi 200 --format png`（200 DPI） |
| [go-ofdrw-rs](https://github.com/lihongjie0209/ofdrw-rs) | v0.5.0 | 原生 `PagePNG(index, 200)`（200 DPI） |
| [rust-ofdbox-folio](https://github.com/ofdbox/ofdbox-folio) | 36e6877bf5d0b989cd824c3c74241533ae1f90ea | 原生 CPU 渲染，200 DPI |
| [rust-sofd](https://github.com/emfox/sOFD) | 65cff468ecc2af676180f540899fdf74da849c3a | 原生 CLI `convert --format png --dpi 200` |
| [rust-easyofd](https://github.com/easy-4-rust/easyofd-rust) | 0.1.4 | 原生 `AWTMaker.with_dpi(200)` |
| [java-ofdrw](https://github.com/ofdrw/ofdrw) | 2.5.1 | 原生 AWT `ImageMaker(reader, 200.0 / 25.4)`（200 DPI） |
| [rust-ofdmanager](https://github.com/feuvan/ofdmanager) | a9b1d29ef90141a369ad47cc32fd6c62dd8b8450 | 原生 CLI `render --dpi 200 --page n`（200 DPI） |

选择键为 `zc310`、`ofdgo`、`ofd-utility`、`ofdrw-rs`、`ofdbox-folio`、`sofd`、`easyofd-rust`、`java`、`ofdmanager`，不指定时全部运行。

`ofd-utility` 的 Python 适配器仅传递固定参数并统一页码/文件名，不解码或重绘图片。`ofdmanager` 上游 CLI 一次只渲染一页，适配器按 OFD 中的页数循环调用，重复解析/字体加载计入耗时；多文档 OFD 会明确拒绝，不静默只测第一个文档。原生 Java `ImageMaker` 的构造参数是**像素/毫米，不是 DPI**。

sOFD 与现有 PDF 包装器保持一致：仅在返回码 1 且 stderr 完整匹配缺字警告时，保留完整 PNG 继续质量比较，让缺字反映在画质得分中；其他非零退出、缺页及损坏文件仍记失败。

Python `ofd2img 0.1.2` 与 Python easyofd 的图片接口均先转 PDF，再固定以 144 DPI 渲染，故不参加本项评测；Node `@miconvert/ofd-to-pdf` 当前发布包仅有 PDF 输出接口，也不参与 PNG 比较。已有 OFD→PDF 评测保持不变。

## 评测规则

- 耗时包含进程启动、解析、渲染、PNG 编码及写盘；不含后续质量比较。毫秒/页为总耗时除以完整输出页数。PNG 总大小是所有页面文件的字节数之和，受各库默认无损压缩策略影响。
- 输出页数、连续页码必须与参考一致；图片尺寸只允许 1 像素的栅格取整差异，统一在右侧/底部补白，不缩放、不裁切、不平移。
- 比较时透明度统一合成白底并转 RGB。参考图的 DPI 元数据若存在，须在 200±1 范围内（数科导出的约 200.03 属正常取整）。
- **内容相似度**：沿用 PDF 评测的参考前景归一化色差，惩罚丢失内容和额外墨迹；优先参考此项与前景 F1，而不是仅看大片空白影响下的全页得分。
- **前景 F1**：允许约 0.265 mm 的物理位置误差，在 200 DPI 对应 2 像素；仅 F1 有容差，颜色比较不移动图像。
- **全页相似度**：`1 − 平均归一化 RGB 色差`。各页等权平均；全白页不抬高内容/F1 得分。丢失几乎全部前景、损坏 PNG、缺页或尺寸不符均记为 `FAILED`。
- 单次转换默认限时 15 分钟，可通过 `BENCHMARK_CONVERTER_TIMEOUT=30m` 调整；超时终止整个进程组并清理部分输出。重跑时清除该转换器旧输出，质量得分重置，避免沿用旧结果。

## 原生 PNG 结果（9 个转换器）

表格由 `docs/benchmark-results.yaml` 生成，勿直接编辑。`FAILED` 表示尚无本轮有效测量或转换/比较失败；首次正式结果由 CI 统一环境测量，本地验证不覆盖已有 PDF 报告的运行环境。

### 总耗时

<!-- benchmark-table:png-speed:start -->

| 转换器 | 999.ofd | intro.ofd | zsbk.ofd | 胜出次数 |
|---|---:|---:|---:|---:|
| [go-zc310](https://github.com/zc310/ofd) | FAILED | FAILED | FAILED | 0 |
| [go-ofdgo](https://github.com/xiaoqidun/ofdgo) | FAILED | FAILED | FAILED | 0 |
| [rust-ofd-utility](https://github.com/ofd-utility/ofd-utility) | FAILED | FAILED | FAILED | 0 |
| [go-ofdrw-rs](https://github.com/lihongjie0209/ofdrw-rs) | FAILED | FAILED | FAILED | 0 |
| [rust-ofdbox-folio](https://github.com/ofdbox/ofdbox-folio) | FAILED | FAILED | FAILED | 0 |
| [rust-sofd](https://github.com/emfox/sOFD) | FAILED | FAILED | FAILED | 0 |
| [rust-easyofd](https://github.com/easy-4-rust/easyofd-rust) | FAILED | FAILED | FAILED | 0 |
| [java-ofdrw](https://github.com/ofdrw/ofdrw) | FAILED | FAILED | FAILED | 0 |
| [rust-ofdmanager](https://github.com/feuvan/ofdmanager) | FAILED | FAILED | FAILED | 0 |
<!-- benchmark-table:png-speed:end -->

### 平均每页耗时

<!-- benchmark-table:png-ms-per-page:start -->

| 转换器 | 999.ofd | intro.ofd | zsbk.ofd | 胜出次数 |
|---|---:|---:|---:|---:|
| [go-zc310](https://github.com/zc310/ofd) | FAILED | FAILED | FAILED | 0 |
| [go-ofdgo](https://github.com/xiaoqidun/ofdgo) | FAILED | FAILED | FAILED | 0 |
| [rust-ofd-utility](https://github.com/ofd-utility/ofd-utility) | FAILED | FAILED | FAILED | 0 |
| [go-ofdrw-rs](https://github.com/lihongjie0209/ofdrw-rs) | FAILED | FAILED | FAILED | 0 |
| [rust-ofdbox-folio](https://github.com/ofdbox/ofdbox-folio) | FAILED | FAILED | FAILED | 0 |
| [rust-sofd](https://github.com/emfox/sOFD) | FAILED | FAILED | FAILED | 0 |
| [rust-easyofd](https://github.com/easy-4-rust/easyofd-rust) | FAILED | FAILED | FAILED | 0 |
| [java-ofdrw](https://github.com/ofdrw/ofdrw) | FAILED | FAILED | FAILED | 0 |
| [rust-ofdmanager](https://github.com/feuvan/ofdmanager) | FAILED | FAILED | FAILED | 0 |
<!-- benchmark-table:png-ms-per-page:end -->

### PNG 总大小

<!-- benchmark-table:png-size:start -->

| 转换器 | 999.ofd | intro.ofd | zsbk.ofd | 胜出次数 |
|---|---:|---:|---:|---:|
| [go-zc310](https://github.com/zc310/ofd) | FAILED | FAILED | FAILED | 0 |
| [go-ofdgo](https://github.com/xiaoqidun/ofdgo) | FAILED | FAILED | FAILED | 0 |
| [rust-ofd-utility](https://github.com/ofd-utility/ofd-utility) | FAILED | FAILED | FAILED | 0 |
| [go-ofdrw-rs](https://github.com/lihongjie0209/ofdrw-rs) | FAILED | FAILED | FAILED | 0 |
| [rust-ofdbox-folio](https://github.com/ofdbox/ofdbox-folio) | FAILED | FAILED | FAILED | 0 |
| [rust-sofd](https://github.com/emfox/sOFD) | FAILED | FAILED | FAILED | 0 |
| [rust-easyofd](https://github.com/easy-4-rust/easyofd-rust) | FAILED | FAILED | FAILED | 0 |
| [java-ofdrw](https://github.com/ofdrw/ofdrw) | FAILED | FAILED | FAILED | 0 |
| [rust-ofdmanager](https://github.com/feuvan/ofdmanager) | FAILED | FAILED | FAILED | 0 |
<!-- benchmark-table:png-size:end -->

### 内容相似度

<!-- benchmark-table:png-content-similarity:start -->

| 转换器 | 999.ofd | intro.ofd | zsbk.ofd | 胜出次数 |
|---|---:|---:|---:|---:|
| [go-zc310](https://github.com/zc310/ofd) | FAILED | FAILED | FAILED | 0 |
| [go-ofdgo](https://github.com/xiaoqidun/ofdgo) | FAILED | FAILED | FAILED | 0 |
| [rust-ofd-utility](https://github.com/ofd-utility/ofd-utility) | FAILED | FAILED | FAILED | 0 |
| [go-ofdrw-rs](https://github.com/lihongjie0209/ofdrw-rs) | FAILED | FAILED | FAILED | 0 |
| [rust-ofdbox-folio](https://github.com/ofdbox/ofdbox-folio) | FAILED | FAILED | FAILED | 0 |
| [rust-sofd](https://github.com/emfox/sOFD) | FAILED | FAILED | FAILED | 0 |
| [rust-easyofd](https://github.com/easy-4-rust/easyofd-rust) | FAILED | FAILED | FAILED | 0 |
| [java-ofdrw](https://github.com/ofdrw/ofdrw) | FAILED | FAILED | FAILED | 0 |
| [rust-ofdmanager](https://github.com/feuvan/ofdmanager) | FAILED | FAILED | FAILED | 0 |
<!-- benchmark-table:png-content-similarity:end -->

### 前景 F1

<!-- benchmark-table:png-foreground-f1:start -->

| 转换器 | 999.ofd | intro.ofd | zsbk.ofd | 胜出次数 |
|---|---:|---:|---:|---:|
| [go-zc310](https://github.com/zc310/ofd) | FAILED | FAILED | FAILED | 0 |
| [go-ofdgo](https://github.com/xiaoqidun/ofdgo) | FAILED | FAILED | FAILED | 0 |
| [rust-ofd-utility](https://github.com/ofd-utility/ofd-utility) | FAILED | FAILED | FAILED | 0 |
| [go-ofdrw-rs](https://github.com/lihongjie0209/ofdrw-rs) | FAILED | FAILED | FAILED | 0 |
| [rust-ofdbox-folio](https://github.com/ofdbox/ofdbox-folio) | FAILED | FAILED | FAILED | 0 |
| [rust-sofd](https://github.com/emfox/sOFD) | FAILED | FAILED | FAILED | 0 |
| [rust-easyofd](https://github.com/easy-4-rust/easyofd-rust) | FAILED | FAILED | FAILED | 0 |
| [java-ofdrw](https://github.com/ofdrw/ofdrw) | FAILED | FAILED | FAILED | 0 |
| [rust-ofdmanager](https://github.com/feuvan/ofdmanager) | FAILED | FAILED | FAILED | 0 |
<!-- benchmark-table:png-foreground-f1:end -->

### 全页相似度

<!-- benchmark-table:png-pixel-similarity:start -->

| 转换器 | 999.ofd | intro.ofd | zsbk.ofd | 胜出次数 |
|---|---:|---:|---:|---:|
| [go-zc310](https://github.com/zc310/ofd) | FAILED | FAILED | FAILED | 0 |
| [go-ofdgo](https://github.com/xiaoqidun/ofdgo) | FAILED | FAILED | FAILED | 0 |
| [rust-ofd-utility](https://github.com/ofd-utility/ofd-utility) | FAILED | FAILED | FAILED | 0 |
| [go-ofdrw-rs](https://github.com/lihongjie0209/ofdrw-rs) | FAILED | FAILED | FAILED | 0 |
| [rust-ofdbox-folio](https://github.com/ofdbox/ofdbox-folio) | FAILED | FAILED | FAILED | 0 |
| [rust-sofd](https://github.com/emfox/sOFD) | FAILED | FAILED | FAILED | 0 |
| [rust-easyofd](https://github.com/easy-4-rust/easyofd-rust) | FAILED | FAILED | FAILED | 0 |
| [java-ofdrw](https://github.com/ofdrw/ofdrw) | FAILED | FAILED | FAILED | 0 |
| [rust-ofdmanager](https://github.com/feuvan/ofdmanager) | FAILED | FAILED | FAILED | 0 |
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
cargo build --release --locked --manifest-path rust/easyofd-convert/Cargo.toml --target-dir rust/target
cp rust/target/release/easyofd-ofd2png bin/
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

# 全部 9 个原生转换器，固定 200 DPI。
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
- CPU：INTEL(R) XEON(R) PLATINUM 8573C（4 逻辑核心）
- 内存：15.6 GiB
<!-- benchmark-environment:end -->
