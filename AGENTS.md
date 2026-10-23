# AGENTS.md — OFD Benchmark

## 项目简介

OFD（开放版式文档，GB/T 33190-2016）转 PDF 速度对比测试项目。比较多种语言的 OFD→PDF 转换库的性能、输出质量和文本提取能力。

## 构建

```bash
# 编译所有 Go 二进制
cd /home/go/workspace/ofd-benchmark
export GOWORK=off  # 上级存在 go.work 时需关闭 workspace，否则报 not one of the workspace modules
go build -o bin/ofd-benchmark .
go build -o bin/ofdgo-convert ./cmd/ofdgo-convert
go build -o bin/zc310-convert ./cmd/zc310-convert
go build -o bin/ofdgo-pdf2ofd ./cmd/ofdgo-pdf2ofd
go build -o bin/pdf2ofd-benchmark ./cmd/pdf2ofd-benchmark

# 编译 zc310 通用转换器（PDF→OFD / OFD→PDF / OFD→TXT），来自本地 zc310/ofd 仓库
cd /home/go/workspace/ofd
go build -o ../ofd-benchmark/bin/ofd-converter ./cmd/ofd-converter

# Java Pdf2Ofd（ofdrw PDF→OFD）——需排除损坏的占位 jar 才能通过 javac
cd /home/go/workspace/ofd-benchmark/java
CP=$(ls lib/*.jar | rg -v 'jaxen|pdfbox-sandbox|sqlite' | tr '\n' ':')
javac -cp "$CP" Pdf2Ofd.java
cp Pdf2Ofd.class ../bin/java/

# Java 运行还需 zip4j.jar（ofdrw 打包 OFD 时用到），Maven Central 下载后放 java/lib/ 与 bin/java/lib/
curl -fsSL -o lib/zip4j.jar https://repo1.maven.org/maven2/net/lingala/zip4j/zip4j/2.11.5/zip4j-2.11.5.jar
```

Java 依赖位于 `java/lib/`，无需额外构建。

Node 依赖位于 `node/node_modules/`，如需更新：

```bash
cd node && npm install
```

## 运行基准测试

```bash
# 运行全部转换器（单个文件）
./bin/ofd-benchmark testdata/ofd-creator/hello.ofd

# 运行指定转换器
./bin/ofd-benchmark testdata/ofd-creator/hello.ofd ofdgo zc310

# 批量运行所有测试文件
for f in testdata/ofd-creator/*.ofd testdata/ofdrw/*.ofd; do
  ./bin/ofd-benchmark "$f"
done
```

## 运行 PDF→OFD 基准测试

```bash
# 运行全部转换器（zc310 / ofdgo / java-ofdrw / python-pdf2ofd / node-jsofd）
./bin/pdf2ofd-benchmark testdata/suwell/hello.pdf

# 运行指定转换器
./bin/pdf2ofd-benchmark testdata/suwell/hello.pdf zc310 ofdgo jsofd

# 批量运行所有测试文件
for f in testdata/suwell/*.pdf testdata/standards/*.pdf; do
  ./bin/pdf2ofd-benchmark "$f"
done
```

结果保存到 `bin/output_ofd/<name>_<converter>.ofd`。回环还原度/文本对比：

```bash
python3 tools/compare_pdf2ofd.py 96
```

## 测试文件

位于 `testdata`，来源：https://github.com/zc310/ofd-fixtures

| 文件 | 大小 | 说明 |
|------|------|------|
| hello.ofd | 1.4KB | 最小示例 |
| ano.ofd | 702KB | 发票 |
| 999.ofd | 29KB | 简单文档 |
| 1000-pages.ofd | 456KB | 多页文档 |
| intro.ofd | 7.1MB | 复杂文档 |
| zsbk.ofd | 1.5MB | 签章文件 |

## 项目结构

```
ofd-benchmark/
├── main.go                    # 基准测试主程序（OFD→PDF 单文件输入）
├── cmd/
│   ├── ofdgo-convert/         # Go xiaoqidun/ofdgo 转换器封装
│   ├── ofdgo-pdf2ofd/         # Go xiaoqidun/ofdgo PDF→OFD 转换器封装
│   ├── zc310-convert/         # Go zc310/ofd 转换器封装
│   └── pdf2ofd-benchmark/     # PDF→OFD 基准测试驱动
├── python/
│   ├── ofd2pdf.py             # Python easyofd 转换脚本
│   ├── ofd2pdf_ofd2pdf.py     # Python ofd2pdf 转换脚本
│   ├── ofd2pdf_ofdreader.py   # Python ofdreader 转换脚本
│   └── pdf2ofd_convert.py     # Python pdf2ofd 转换脚本（PDF→OFD）
├── java/
│   ├── Ofd2Pdf.java           # Java ofdrw OFD→PDF 转换器
│   ├── Pdf2Ofd.java           # Java ofdrw PDF→OFD 转换器
│   ├── ofd2pdf_java.sh        # Java 转换启动脚本
│   ├── pdf2ofd_java.sh        # Java PDF→OFD 启动脚本
│   └── lib/                   # Java 依赖 JAR（含 zip4j）
├── node/
│   ├── ofd2pdf.mjs            # Node.js ofd-to-pdf 转换脚本
│   ├── ofd2pdf_node.sh        # Node 转换启动脚本
│   ├── pdf2ofd_jsofd.mjs      # Node jsOFD PDF→OFD 转换脚本
│   └── node_modules/          # Node 依赖（@miconvert、@hufe921/jsofd、pdfjs-dist）
├── bin/                       # 编译输出目录
│   ├── ofd-benchmark          # OFD→PDF 主测试程序
│   ├── pdf2ofd-benchmark      # PDF→OFD 主测试程序
│   ├── ofd-converter          # zc310 通用转换器（PDF→OFD / OFD→PDF/TXT）
│   ├── ofdgo-convert          # Go ofdgo OFD→PDF 转换器
│   ├── ofdgo-pdf2ofd          # Go ofdgo PDF→OFD 转换器
│   ├── zc310-convert          # Go 转换器
│   ├── output/                # OFD→PDF 生成文件
│   ├── output_ofd/            # PDF→OFD 生成文件（.ofd + 回环 .pdf + .txt）
│   ├── python/                # Python 脚本副本
│   ├── java/                  # Java 脚本副本
│   └── node/                  # Node 脚本副本
├── testdata/                  # OFD/PDF 测试文件（ofd-creator/ ofdrw/ standards/ suwell/）
├── tools/
│   ├── compare_pdf.py         # 与数科参考 PDF 的像素级对比脚本
│   ├── compare_text.py        # 与数科参考 PDF 的文本相似度对比脚本
│   ├── compare_pdf2ofd.py     # PDF→OFD 回环还原度/文本对比脚本
│   └── gen-libraries/         # 库汇总表生成器（Go，读 yaml 生成 md）
├── docs/
│   ├── libraries.yaml         # OFD 库汇总数据源（列定义+图例+库条目）
│   ├── ofd-libraries.md       # OFD 库汇总对比表（由 gen-libraries 生成，勿手改）
│   ├── ofd-to-pdf.md          # OFD→PDF 基准测试结果
│   └── pdf-to-ofd.md          # PDF→OFD 基准测试结果
└── README.md                  # 项目索引
```

## 基准测试流程

`main.go` 对每个输入文件依次执行所有转换器：

1. 每个转换器是一个独立进程，接收 `<input.ofd> <output.pdf>` 参数
2. 转换器在 `os.TempDir()` 中执行
3. 生成的 PDF 移动到 `bin/output/<input>_<converter>.pdf`
4. 输出转换耗时和文件大小

## 文本提取测试

从生成的 PDF 中提取文本，统计字符数：

```python
import fitz  # pip install pymupdf

doc = fitz.open("bin/output/hello_go-zc310.pdf")
text = ""
for page in doc:
    text += page.get_text()
print(len(text.strip()))
```

## 与数科版式输出对比

`testdata` 为各 OFD 的参考转换结果（数科版式阅读器导出）。对比转换后 PDF 与参考 PDF：

```bash
# 像素级对比（默认 96 DPI），可传多个文件名
python3 tools/compare_pdf.py 96 hello ano 999 1000-pages intro zsbk

# 文本相似度对比（4-gram Jaccard，与数科参考 PDF 的提取文本比较）
python3 tools/compare_text.py
```

像素对比输出：参考 PDF 大小、转换后 PDF 大小、大小比、像素相似度（1 - 平均归一化色差）。文本对比输出：参考/转换后字符数、字符数比、文本相似度。

> **注意事项**：
> - `testdata` 来自数科版式阅读器，勿删除（像素对比基准）
> - 像素对比适用于所有页数相同的文件；1000-pages 用 96 DPI 需约 3.5 分钟

## 添加新转换器

1. 在对应语言目录创建转换脚本，接收 `<input.ofd> <output.pdf>` 参数
2. 在 `main.go` 的 `converters` map 中添加条目
3. 在 `bin/` 下放置可执行脚本（或二进制）
4. 运行测试验证

## 库汇总表维护

`docs/ofd-libraries.md` 为生成物，**勿直接编辑**。数据源是 `docs/libraries.yaml`（列定义+图例+库条目，值仅限 ✅/⚠️/❌）：

```bash
GOWORK=off go run ./tools/gen-libraries          # 改完 yaml 重新生成 md
GOWORK=off go run ./tools/gen-libraries -check   # 校验 md 与数据源一致（CI 可用）
```

- 加库/改能力值：只改 `libraries.yaml` 对应条目后重新生成
- 加新列：`columns` 里插入新条目（建议带 `legend` 图例文案），并为**每个**库补齐 `caps` 中同名键，否则校验报错
- 新增列在表中的位置 = `columns` 里的顺序（name/license 为固定结构列）

## 注意事项

- `python-ofdreader` 模块未安装，该转换器始终 FAILED
- `python-easyofd` 对部分文件（hello、1000-pages）会 crash（list index out of range）
- `python-ofd2pdf` 基于图片渲染，无法提取文本
- PDF→OFD：`python-pdf2ofd` 需 `pip install pdf2ofd`（PEP 668 环境加 `--user --break-system-packages`）
- `java/lib/` 中有 3 个损坏占位 jar（jaxen/pdfbox-sandbox/sqlite，554 字节）：`javac` 需排除，运行时 `java` 会自动忽略
- ofdrw PDF→OFD（Pdf2Ofd.java）运行需要 `zip4j.jar`（已下载到 java/lib 与 bin/java/lib）
- ofdrw / pdf2ofd 的 PDF→OFD 输出不含可提取文本（转矢量轮廓 / 位图），仅 go-zc310 保留 TextObject
- ofdgo 的 PDF→OFD（ConvertPDF + pdfgo）7 个测试文件 7/7 全部成功（2026-09-25 15:21 版修复非 Identity-H 复合字体编码、Stamp 注解、nil panic）；输出 OFD 写 16 位色值，新版 zc310 ofd-converter 已兼容读取（回环可测）
- node-jsofd（@hufe921/jsofd + pdfjs-dist）：`npm install` 在 node/ 下；Node <22 需补 `process.getBuiltinModule`（见 `node/pdf2ofd_jsofd.mjs`）否则 CMap 加载失败丢 CJK；不转换注解/渐变/图案
- 构建 ofd-benchmark 时若存在上级 `go.work`（workspace 模式），需 `GOWORK=off` 否则报 “not one of the workspace modules”
- 测试文件需从本地 `/home/go/workspace/ofd/test/testdata/` 复制（网络下载超时）
- `testdata` 是数科版式阅读器的参考输出，作像素对比基准，同时用作 PDF→OFD 的输入
- intro.pdf 的像素对比自动跳过 14～19 页（数科导出这几页背景图丢失）：`tools/compare_pdf.py` 与 `compare_pdf2ofd.py` 内置 `EXCLUDE_PAGES`
- 运行环境：Ubuntu 26.04.1 LTS，已安装中文字体（wqy-zenhei, noto-cjk）
