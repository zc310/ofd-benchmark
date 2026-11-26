# OFD 转换与渲染对比测试

> [OFD 库汇总](docs/ofd-libraries.md) —— 各 OFD 库能力 / 许可证对比表（预览、解析、发票、→ 图片/PDF/TXT/SVG/Md/HTML、PDF→OFD、生成、签章、修改）

本项目对比多种开源 OFD 库的转换与渲染能力：

| 方向      | 文档                                     | 对比内容                                                                                     |
|-----------|------------------------------------------|----------------------------------------------------------------------------------------------|
| OFD → PDF | [docs/ofd-to-pdf.md](docs/ofd-to-pdf.md) | 8 种转换器（Go×3 / Rust×2 / Python / Java / Node）的速度、文件大小、文本提取与数科版式还原度 |
| PDF → OFD | [docs/pdf-to-ofd.md](docs/pdf-to-ofd.md) | 5 种转换器（zc310 / ofdgo / ofdrw / pdf2ofd / jsOFD）的速度、文件大小与回环还原度 |
| OFD → PNG | [docs/ofd-to-png.md](docs/ofd-to-png.md) | 8 个原生 200 DPI 转换器的速度、体积与数科直接导出 PNG 的还原度，不使用 PDF 中转 |

基准参考输出来自**数科版式阅读器**（`testdata`），保证同一组文档双向可测。测试用的 PDF/OFD 可在 [zc310/ofd-fixtures](https://github.com/zc310/ofd-fixtures) 下载。

## 快速开始

测试前先准备字体环境（Linux，需 `curl`、`unzip`、`fontconfig`）：

```bash
sudo apt-get install -y fonts-wqy-zenhei fonts-noto-cjk fontconfig libfontconfig1-dev pkg-config curl unzip
bash tools/install-fonts.sh
```

脚本从 [zispace/hanzi-fonts v1.0](https://github.com/zispace/hanzi-fonts/releases/tag/v1.0) 下载以下字体包，解压到用户目录 `~/.local/share/fonts/hanzi-fonts/`，并执行 `fc-cache` 刷新字体缓存；字体安装本身无需 sudo。GitHub Actions 也会在运行测试前执行此脚本。

- [doc-foundertype.zip](https://github.com/zispace/hanzi-fonts/releases/download/v1.0/doc-foundertype.zip)
- [doc-zhongyi-etc.zip](https://github.com/zispace/hanzi-fonts/releases/download/v1.0/doc-zhongyi-etc.zip)
- [latin-fonts.zip](https://github.com/zispace/hanzi-fonts/releases/download/v1.0/latin-fonts.zip)

```bash
# OFD → PDF 基准（详见 docs/ofd-to-pdf.md）
GOWORK=off go build -o bin/ofdrw-rs-convert ./cmd/ofdrw-rs-convert
GOWORK=off cargo build --release --locked --manifest-path rust/ofdbox-folio-convert/Cargo.toml --target-dir rust/target
cp rust/target/release/ofdbox-folio-convert bin/ofdbox-folio-convert
cargo install sofd-cli --git https://github.com/emfox/sOFD --rev 65cff468ecc2af676180f540899fdf74da849c3a --locked --bin sofd --root /tmp/sofd-install
cp /tmp/sofd-install/bin/sofd bin/sofd-cli
go build -o bin/ofd-benchmark .
./bin/ofd-benchmark testdata/ofd-creator/hello.ofd

# PDF → OFD 基准（详见 docs/pdf-to-ofd.md）
cd /home/go/workspace/ofd && go build -o ../ofd-benchmark/bin/ofd-converter ./cmd/ofd-converter
cd ../ofd-benchmark && go build -o bin/pdf2ofd-benchmark ./cmd/pdf2ofd-benchmark
./bin/pdf2ofd-benchmark testdata/suwell/hello.pdf
```

转换结果分别保存在 `bin/output/`（OFD→PDF）与 `bin/output_ofd/`（PDF→OFD）。

OFD→PNG 独立评测见 [运行说明](docs/ofd-to-png.md#运行)，输出到 `bin/output_png/`。当前参考图为 `testdata/suwell/{999,intro,zsbk}/` 下数科直接导出的 200 DPI PNG，逐页比较，`intro` 第 14～19 页正常参与；参考图通过 `ofd-fixtures` 同步，不提交到本仓库。

`ofdrw-rs` 使用 Go bindings v0.5.0 和同版本 Linux x64 原生库。首次运行前下载原生库：

```bash
curl -fL -o /tmp/ofdrw-native.zip https://github.com/lihongjie0209/ofdrw-rs/releases/download/v0.5.0/ofdrw-native-linux-x64-v0.5.0.zip
unzip -o /tmp/ofdrw-native.zip -d /tmp/ofdrw-native
mkdir -p bin/ofdrw-rs
cp /tmp/ofdrw-native/ofdrw-native-linux-x64-v0.5.0/libofdrw_ffi.so bin/ofdrw-rs/
```

也可通过 `OFDRW_LIB_PATH` 指定原生库路径。基准会将该转换器标记为 `go-ofdrw-rs`。

三个基准驱动的单次转换默认限时 15 分钟，可通过 `BENCHMARK_CONVERTER_TIMEOUT=30m` 调整。超时会终止转换器及其子进程，将结果记为 `FAILED`，并继续测试其他转换器。像素比较遇到损坏文件、页数不一致或无可比较页时，同样将该项记为 `FAILED`，不会中断其他输出的比较。

回环 PDF/TXT 生成使用 `bash tools/roundtrip-pdf2ofd.sh bin/ofd-converter bin/output_ofd`，每个格式单独限时 15 分钟；可通过 `BENCHMARK_ROUNDTRIP_TIMEOUT=30m` 调整（GNU timeout 时长格式）。失败或超时会删除该格式的部分输出并继续。

CI 的 Python 依赖（含间接依赖）由 `python/requirements.txt` 锁定版本及哈希，Node 由 `node/package-lock.json` 锁定并用 `npm ci` 安装。pdf2ofd 依赖旧版 PyMuPDF，须安装到独立环境，避免将像素评测渲染器降级：

```bash
python3.12 -m venv .venv-benchmark
.venv-benchmark/bin/python -m pip install --only-binary=:all: --require-hashes -r python/requirements.txt
python3.12 -m venv .venv-pdf2ofd
.venv-pdf2ofd/bin/python -m pip install --only-binary=:all: --require-hashes -r python/requirements-pdf2ofd.txt
export PDF2OFD_PYTHON="$PWD/.venv-pdf2ofd/bin/python"
export PATH="$PWD/.venv-benchmark/bin:$PATH"
npm --prefix node ci
```

更新 Python 依赖时修改对应 `.in` 文件，再按 `.txt` 文件头记录的 `uv pip compile` 命令重新生成锁文件。CI 在生成报告后及 rebase 后检查报告与 YAML 一致，检查失败不推送。

## 基准结果数据维护

`docs/benchmark-results.yaml` 是三份基准报告中量化结果表的数据源。修改 YAML 后运行以下命令重新生成表格；最优值加粗和胜出次数由生成器计算：

```bash
GOWORK=off go run ./tools/gen-benchmarks
GOWORK=off go run ./tools/gen-benchmarks -check  # 校验文档是否最新
```

报告中的说明文字不由生成器管理；带 `benchmark-table` 和 `benchmark-charts` 标记的表格及图表，以及 `docs/charts/` 下的图表，由 YAML 生成，勿直接编辑。

基准驱动会将转换耗时、输出大小和大小比写回 YAML；对比脚本会补充像素相似度、文本提取和文本相似度。GitHub Actions 每周检查 `zc310/ofd`、`xiaoqidun/ofdgo` 的最新提交及 Maven Central 上 `org.ofdrw:ofdrw-converter` 的最新发布版，`ofdrw-rs` 当前固定使用 v0.5.0；只有受跟踪版本变化时才运行完整基准并提交更新，也可在 Actions 页面手动触发（手动运行不要求有更新）。Workflow 会下载匹配的 Linux x64 原生库。本轮测试的库版本和 runner 的 OS、CPU 型号/逻辑核心数、内存均记入 YAML 并显示在报告中；Go 模块升级只在 workflow runner 临时生效，不会提交 `go.mod` 变更。完整运行会拉取 `zc310/ofd-fixtures` 样例，并在开始时使用 `-reset` 将待测单元重置为 `FAILED`，避免失败或缺失结果沿用旧数据。

## 标签发布评测产物

推送 `v*` 标签会强制运行完整基准，成功后自动创建同名 GitHub Release：

```bash
git tag v2026.11.16
git push origin v2026.11.16
```

- 每个库一个 ZIP（如 `zc310-v2026.11.16.zip`），包含 `ofd-to-pdf/`、`pdf-to-ofd/` 和 `roundtrip/` 下的 PDF/OFD；回环 PDF 统一由 zc310 渲染。不同语言实现分别打包，仅支持单向转换的库只包含对应产物。
- 附件还包含本次三份 Markdown 报告、`benchmark-results.yaml`、`manifest.json`、发布说明和 `SHA256SUMS`。清单记录文件及缺失样例；失败或空输出不打包，全部失败的库不生成空 ZIP。构建或评测流程失败则不发布正式版本。PNG 页面暂通过 Actions 的 `ofd-to-png-200dpi` 附件保存 7 天，不加入按库的 PDF/OFD Release ZIP。
- 标签运行不提交报告、不移动标签。依赖仍按现有逻辑在运行时解析，实际受跟踪版本及运行环境见附件报告；标签本身不保证所有外部依赖和样例可复现。
- 先创建草稿并上传全部附件，再正式发布；已有 Release（含草稿）不会覆盖。上传失败留下的草稿需手动删除后重试；已发布结果需要重测时请打新标签。Actions 中间附件保留 7 天，Release 附件不受此期限影响。

本地打包已有评测输出（目标目录必须为空或不存在）：

```bash
python3 tools/package_benchmark_release.py --tag v2026.11.16 --output /tmp/benchmark-release
```

## 项目结构

```
ofd-benchmark/
├── bin/
│   ├── ofd-benchmark       # OFD→PDF 基准测试主程序
│   ├── ofdgo-convert       # Go xiaoqidun/ofdgo OFD→PDF 转换器
│   ├── ofdrw-rs-convert    # ofdrw-rs Go binding OFD→PDF 转换器
│   ├── ofdgo-pdf2ofd       # Go xiaoqidun/ofdgo PDF→OFD 转换器
│   ├── zc310-convert       # Go zc310/ofd OFD→PDF 转换器
│   ├── ofd-converter       # Go zc310/ofd 通用转换器（PDF→OFD、OFD→PDF/TXT）
│   ├── pdf2ofd-benchmark   # PDF→OFD 基准测试主程序
│   ├── ofd2png-benchmark   # OFD→PNG 原生 200 DPI 基准驱动
│   ├── zc310-png          # zc310 原生 PNG 适配器
│   ├── ofdgo-png          # ofdgo 原生 PNG 适配器
│   ├── ofd-utility-cli    # 固定 revision 的 Rust ofd-cli
│   ├── output_png/        # OFD→PNG 按样本/转换器分目录输出
│   ├── output/             # OFD→PDF 转换结果
│   └── output_ofd/         # PDF→OFD 转换结果（含回环 PDF / 提取文本）
├── cmd/
│   ├── ofdgo-convert/
│   ├── ofdrw-rs-convert/
│   ├── ofdgo-pdf2ofd/
│   ├── ofd2png-benchmark/
│   ├── zc310-png/
│   ├── ofdgo-png/
│   ├── zc310-convert/
│   └── pdf2ofd-benchmark/  # PDF→OFD 基准驱动
├── rust/
│   ├── ofdbox-folio-convert/ # Rust ofdbox-folio OFD→PDF 转换器
│   └── sofd-convert.sh       # sOFD CLI 结果适配器
├── python/
│   ├── ofd2pdf.py          # Python 转换器（仅 easyofd）
│   ├── ofd2pdf_ofd2pdf.py  # 旧独立脚本，仅供参考，不参与评测
│   └── pdf2ofd_convert.py  # Python pdf2ofd 转换器
├── java/
│   ├── Ofd2Pdf.java        # Java ofdrw OFD→PDF 转换器
│   ├── Pdf2Ofd.java        # Java ofdrw PDF→OFD 转换器（PDFConverter）
│   ├── ofd2pdf_java.sh
│   ├── pdf2ofd_java.sh
│   ├── pom.xml             # ofdrw converter + zip4j 依赖
│   └── lib/                # JAR 依赖
├── node/
│   ├── ofd2pdf.mjs         # Node.js 转换器（@miconvert/ofd-to-pdf）
│   └── ofd2pdf_node.sh
├── tools/
│   ├── install-fonts.sh    # 测试前下载字体并安装到用户目录
│   ├── compare_pdf.py      # OFD→PDF 与数科参考 PDF 像素对比
│   ├── compare_png.py      # 原生 200 DPI PNG 与数科直接导出的 PNG 比较
│   ├── ofd_utility_png.py  # ofd-utility CLI 页码/文件名适配
│   ├── compare_text.py     # OFD→PDF 文本相似度对比
│   ├── compare_pdf2ofd.py  # PDF→OFD 回环还原度/文本对比
│   ├── benchmark_results.py # 对比脚本写入结果 YAML 的桥接
│   ├── gen-benchmarks/     # 基准 Markdown 表格生成器
│   ├── benchmark-versions/ # 检查/记录 Go 库版本
│   └── update-benchmark-results/ # YAML 原子更新命令
├── docs/
│   ├── benchmark-results.yaml # OFD↔PDF 基准结果数据源
│   ├── ofd-libraries.md    # OFD 库能力 / 许可证对比表
│   ├── ofd-to-pdf.md       # OFD→PDF 基准测试详细结果
│   ├── ofd-to-png.md       # 原生 OFD→PNG 基准测试（200 DPI）
│   └── pdf-to-ofd.md       # PDF→OFD 基准测试详细结果
├── testdata/               # OFD/PDF 测试文件（ofd-creator/ ofdrw/ standards/ suwell/）
├── main.go                 # OFD→PDF 基准主程序（单文件输入）
└── go.mod                  # 含 replace 到本地 zc310/ofd 与 zc310/canvas
```
