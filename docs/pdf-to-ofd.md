# PDF → OFD 对比测试

> [OFD 库汇总](ofd-libraries.md)：各 OFD 库能力 / 许可证对比表
>
> [OFD → PDF 对比测试](ofd-to-pdf.md)：反方向 OFD→PDF 转换对比

能把 PDF 转回 OFD 的开源库很少。本测试对比 6 个转换器（Go zc310、Go ofdgo、Rust easyofd、Java ofdrw、Python pdf2ofd、JS jsOFD）在 7 个参考 PDF 上的转换速度、输出 OFD 文件大小与还原质量。

## 测试库

| 转换器         | 语言   | 库                                                                                         | 版本                   | 说明                                                                                                         |
|----------------|--------|--------------------------------------------------------------------------------------------|------------------------|--------------------------------------------------------------------------------------------------------------|
| go-zc310       | Go     | [zc310/ofd](https://github.com/zc310/ofd)                                                  | v0.1.3 起（pdfimport） | 基于 pdfcpu 的 PDF→OFD 导入器（路径/文本/图片/注解/渐变/图案/网格/大纲/元数据），文本保留为可提取 TextObject |
| go-ofdgo       | Go     | [xiaoqidun/ofdgo](https://github.com/xiaoqidun/ofdgo)（ConvertPDF + pdfgo）                | 2026-09-26 版          | 基于 pdfgo 的 PDF→OFD 导入器（路径/文本/图片/注解） |
| rust-easyofd   | Rust   | [easyofd-rust](https://github.com/easy-4-rust/easyofd-rust)（easyofd-convert PdfImporter） | 0.1.3                  | 基于 lopdf 提取文本/图片；不转换矢量路径                                                                     |
| java-ofdrw     | Java   | [ofdrw](https://github.com/ofdrw/ofdrw)（ofdrw-converter PDFConverter）                    | 2.4.0                  | 文本转为矢量轮廓，保留原始外观                                                                               |
| python-pdf2ofd | Python | [pdf2ofd](https://github.com/wanglrebe/pdf2ofd)                                            | 0.1.0                  | 基于 PyMuPDF 渲染，文本栅格化                                                                                |
| node-jsofd     | JS     | [Hufe921/jsOFD](https://github.com/Hufe921/jsOFD)（pdfToOfd + pdfjs-dist）                | 1.0.1                  | 基于 pdfjs 重放绘制指令流（路径/文本/图片），不转换注解/渐变/图案/Type3/剪裁 |

> `python-pdf2ofd` 安装：`pip install pdf2ofd`（PEP 668 环境需 `--user --break-system-packages` 或使用 venv）
>
> `node-jsofd` 安装：`npm install @hufe921/jsofd pdfjs-dist`（Node <22 需补 `process.getBuiltinModule`，见 `node/pdf2ofd_jsofd.mjs`）

## 测试文件

直接使用数科版式阅读器导出的参考 PDF（即 OFD→PDF 测试的对比基准），保证与 OFD→PDF 测试同一组文档。

| 文件               | 大小 | 页数    | 说明                                           |
|--------------------|------|---------|------------------------------------------------|
| hello.pdf          | 3.5K | 1 页    | 简单文本                                       |
| ano.pdf            | 904K | 3 页    | 发票                                           |
| 999.pdf            | 80K  | 5 页    | 简单文档                                       |
| intro.pdf          | 6.7M | 42 页   | 介绍文档                                       |
| zsbk.pdf           | 1.6M | 2 页    | 签章文件                                       |
| 1000-pages.pdf     | 2.1M | 1000 页 | 多页文档                                       |
| GBT_33190-2016.pdf | 4.8M | 132 页  | GB/T 33190-2016 标准正文（字体全内嵌、图表多） |

## 测试环境

- 系统：Ubuntu 26.04.1 LTS (64-bit)
- 字体目录：`~/.local/share/fonts`（已安装各种中文字体）
- 系统已安装 `apt-get install fonts-wqy-zenhei fonts-noto-cjk`

## 转换速度

| 转换器         | hello.pdf |   ano.pdf |   999.pdf | intro.pdf | zsbk.pdf | 1000-pages.pdf | GBT_33190-2016.pdf | 胜出次数 |
|----------------|----------:|----------:|----------:|----------:|---------:|---------------:|-------------------:|---------:|
| go-zc310       |      44ms |      88ms |      68ms |     1.13s |    344ms |          2.06s |              3.70s |        0 |
| go-ofdgo      |  15.8ms |  197ms |  325ms |  2.37s |  6.46s |  3.69s |  14.81s | 0 |
| rust-easyofd   | **2.3ms** | **3.4ms** | **3.2ms** | **106ms** | **22ms** |      **276ms** |           **79ms** |        7 |
| java-ofdrw     |     472ms |     1.27s |     1.02s |     4.37s |    2.90s |         25.90s |             12.78s |        0 |
| python-pdf2ofd |      99ms |     800ms |     513ms |     2.39s |    928ms |         34.24s |             18.47s |        0 |
| node-jsofd     |     416ms |     712ms |     3.50s |    10.01s |     5.25s |          4.77s |              9.02s |        0 |

> rust-easyofd 的「最快」源于几乎未提取内容（hello/ano 输出空页），速度优势无实际意义

## OFD 文件大小

| 转换器         | hello.pdf |  ano.pdf |  999.pdf | intro.pdf | zsbk.pdf | 1000-pages.pdf | GBT_33190-2016.pdf | 胜出次数 |
|----------------|----------:|---------:|---------:|----------:|---------:|---------------:|-------------------:|---------:|
| go-zc310       |      1.6K |     381K |    47.2K |      7.4M |     1.5M |           3.1M |               6.6M |        0 |
| go-ofdgo       |      1.9K |   585.9K |    86.6K |      8.5M |     3.6M |           5.6M |              12.2M |        0 |
| rust-easyofd   |  **1.2K** | **1.8K** | **3.6K** |  **5.4M** | **919K** |       **432K** |           **2.0M** |        7 |
| java-ofdrw     |      5.2K |     905K |     818K |     34.4M |    19.4M |          57.4M |              30.5M |        0 |
| python-pdf2ofd |      4.3K |     1.3M |     734K |      9.1M |     1.7M |          44.8M |              37.7M |        0 |
| node-jsofd     |      7.4K |     2.2M |   591.6K |     35.6M |    13.2M |          63.4M |              35.1M |        0 |

> rust-easyofd「体积最小」同样是内容丢失的结果（hello 1.2K/ano 1.8K 实为空页），不具可比性

## 文件大小对比（输出 OFD / 原 PDF）

越小表示压缩率越高，`1.00` 表示与源 PDF 等大；**加粗**为最小。

| 转换器         | hello.pdf |  ano.pdf |  999.pdf | intro.pdf | zsbk.pdf | 1000-pages.pdf | GBT_33190-2016.pdf |
|----------------|----------:|---------:|---------:|----------:|---------:|---------------:|-------------------:|
| go-zc310       |      0.46 |     0.42 |     0.59 |      1.08 |     0.93 |           1.43 |               1.34 |
| go-ofdgo       |      0.53 |     0.65 |     1.08 |      1.24 |     2.21 |           2.58 |               2.47 |
| rust-easyofd   |  **0.33** | **0.00** | **0.04** |  **0.78** | **0.57** |       **0.20** |           **0.41** |
| java-ofdrw     |      1.48 |     1.00 |    10.24 |      5.01 |    12.07 |          26.48 |               6.17 |
| python-pdf2ofd |      1.22 |     1.43 |     9.18 |      1.33 |     1.04 |          20.63 |               7.62 |
| node-jsofd     |      2.08 |     2.47 |     7.40 |      5.18 |     8.18 |          29.25 |               7.09 |

> 两种文本转图形的方案输出普遍膨胀：ofdrw 在 999、zsbk、1000-pages 上是源 PDF 的 10～26 倍，pdf2ofd 在 999、1000-pages 上是 9～21 倍（zc310 全部 ≤1.5x）；jsOFD 全面膨胀（2.1～29 倍，1000-pages 达 29 倍）；rust-easyofd 的 0.00～0.78 是内容丢失而非压缩

## 还原度对比

将每个转换器输出的 OFD 再经 go-zc310 转回 PDF（回环），与源 PDF 做像素对比（96 DPI）与文本对比。

**像素相似度**（100% 为完全相同）：

| 转换器         |  hello.pdf |    ano.pdf |    999.pdf |  intro.pdf |   zsbk.pdf | 1000-pages.pdf | GBT_33190-2016.pdf | 胜出次数 |
|----------------|-----------:|-----------:|-----------:|-----------:|-----------:|---------------:|-------------------:|---------:|
| go-zc310       |     99.95% | **99.75%** |     98.19% |     99.60% | **99.09%** |         99.54% |             99.69% |        2 |
| go-ofdgo      |  99.95% |  99.41% |  98.10% |  **99.80%** |  97.23% |  99.45% |  **99.95%** | 2 |
| rust-easyofd  |  99.84% |  98.56% |  95.77% |  43.84% |  95.62% |  73.67% |  96.28% | 0 |
| java-ofdrw     | **99.98%** |     98.37% |     98.73% | 96.90% |     32.97% |     **99.62%** | 91.85% |        2 |
| python-pdf2ofd |     99.94% |     98.34% | **98.94%** |     99.35% |     97.37% |         99.31% |             99.12% |        1 |
| node-jsofd     |     99.89% |     89.95% |     98.13% |     98.94% |     98.54% |         99.59% |             99.26% |        0 |

**文本提取能力**（从转换后 OFD 提取字符数，参考 PDF 提取字符数见括号）：

| 转换器         | hello.pdf(25) | ano.pdf(7139) | 999.pdf(3934) | intro.pdf(5546) | zsbk.pdf(1150) | 1000-pages.pdf(323700) | GBT_33190-2016.pdf(107779) |
|----------------|---------------|---------------|---------------|-----------------|----------------|------------------------|----------------------------|
| go-zc310       | 37            | 9443          | 9746          | 7612            | 2698           | 420998                 | 170740                     |
| go-ofdgo       | 41            | 9882          | 10163         | 10927           | 3375           | 534998                 | 234779                     |
| rust-easyofd   | 0             | 0             | 242*          | 0               | 10*            | 10996*                 | 0                          |
| java-ofdrw     | 0             | 0             | 0             | 0               | 0              | 0                      | 0                          |
| python-pdf2ofd | 0             | 0             | 0             | 0               | 0              | 0                      | 0                          |
| node-jsofd     | 41            | 11882         | 10039         | 10478           | 2899           | 510996                 | 225836                     |

**文本相似度**（4-gram Jaccard，100% 为与源 PDF 提取文本完全一致）：

| 转换器         | hello.pdf | ano.pdf | 999.pdf | intro.pdf | zsbk.pdf | 1000-pages.pdf | GBT_33190-2016.pdf |
|----------------|----------:|--------:|--------:|----------:|---------:|---------------:|-------------------:|
| go-zc310       |    17.02% |  34.55% |  52.11% |    45.20% |    9.11% |          8.25% |             42.01% |
| go-ofdgo       |    17.02% |  27.13% |  29.63% |    13.71% |    4.63% |          7.34% |              1.08% |
| rust-easyofd   |     0.00% |   0.00% |   0.00% |     0.00% |    0.00% |          0.00% |              0.00% |
| java-ofdrw     |     0.00% |   0.00% |   0.00% |     0.00% |    0.00% |          0.00% |              0.00% |
| python-pdf2ofd |     0.00% |   0.00% |   0.00% |     0.00% |    0.00% |          0.00% |              0.00% |
| node-jsofd     |   10.00% |  19.08% |  11.27% |   13.58% |    4.50% |          5.49% |              1.08% |

> - ofdrw / pdf2ofd 的输出 OFD 不含可提取文本（Content.xml 无 TextObject）——ofdrw 转成矢量轮廓、pdf2ofd 栅格成位图，故相似度为 0
> - go-zc310 保留文本对象（全部 6 个文件均可提取），文本顺序正确（如「欢迎使用」顺序无误），但字形间插入空格（如 `H e l l o`）使 4-gram 相似度偏低，内容本身完整
> - rust-easyofd 实测质量最差：hello/ano 输出空页（Layer 为空）；999、1000-pages 提取的文本为 CID 字体乱码（标 \* 的 242/10996/10 字符）且版面固定在 (10,20)；像素高分（hello 99.84%、ano 98.56%、GBT 96.28%）是因为空白/仅图表占主导，不代表还原正确（GBT 文本 0 字符）
> - go-ofdgo 像素 2 胜（intro 99.80%、GBT 99.95%）；文本提取量普遍多于 zc310（hello 41、1000-pages 534998、GBT 234779 字符）但 4-gram 相似度明显偏低（顺序/字距差异，GBT 仅 1.08%）
> - node-jsofd 逐条重放绘制指令流（路径/文本/图片），不转换注解/渐变/图案/Type3/剪裁：ano 因 Stamp 丢失像素仅 89.95%；其余像素普遍高（98.13～99.89%）但逐字间隔重、相似度低（≤19.08%），输出体积全场最大（1000-pages 63.4M）
> - 回环像素随 zc310 渲染版本波动：渐变渲染修复后 GBT 各家普遍提升（zc310 94.71%→99.69%），但 ofdrw 的 zsbk/GBT 大幅下降（32.97%、91.85%）
> - 像素相似度基于整页平均色差，对内容稀疏的页面（白底小字）区分度有限，需结合文本提取与人工核对判断
> - intro 的像素对比跳过 14～19 页（数科参考 PDF 这几页背景图丢失，不计入对比）
> - 回环统一使用 go-zc310 的 OFD→PDF，度量的是「输出 OFD 的互操作正确性」而非端到端体验：若改用各库自回环，写错坐标/格式可被自家读取器按同错方式读回而掩盖（如 ofdgo 的 16 位色值只有换读取器才暴露）；代价是引入 zc310 渲染质量偏差，数据仅供参考

## 结论

- **go-zc310**：唯一保留可复制文本（顺序正确）且可用的转换器，像素还原 2 胜（ano、zsbk），在可用转换器中速度/体积最优（44ms～3.70s、1.6K～6.6M）
- **go-ofdgo**：像素 2 胜（intro 99.80%、GBT 99.95%）、hello 速度反超 zc310（15.8ms）；文本提取量最大（1000-pages 534998、GBT 234779 字符）但相似度偏低（≤29.63%）
- **rust-easyofd**：速度与体积数值均第一（各 7 胜）但为虚假优势——hello/ano 输出空页、999/1000-pages 文本为 CID 乱码、intro 像素仅 43.84%，GBT 像素 96.28% 亦仅因图表保留（文本 0 字符），实际质量最差
- **java-ofdrw**：像素还原 2 胜（hello、1000-pages），文本转为轮廓不可提取；输出体积大（1000-pages 57MB、GBT 30.5MB）、多页转换慢（GBT 12.78s），zsbk/GBT 回环像素受新版渐变渲染影响大幅下降（32.97%、91.85%）
- **python-pdf2ofd**：简单易用，单页/小文件还原不错（999 像素最高），但纯位图渲染不可提取文本、大文件最慢（1000-pages 34.2s、GBT 18.5s）
- **node-jsofd**：像素普遍接近源 PDF（6 文件 ≥98.1%），但体积最大（1000-pages 63.4M、GBT 35.1M）且大文件偏慢；文本相似度低（≤19.08%），ano 因 Stamp 未转换像素仅 89.95%

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

# 编译 Rust PDF→OFD 转换器（easyofd）
cargo build --release --manifest-path rust/pdf2ofd/Cargo.toml
cp rust/pdf2ofd/target/release/pdf2ofd bin/easyofd-pdf2ofd

# 运行全部转换器（单个文件）
./bin/pdf2ofd-benchmark test/testdata/hello.pdf

# 批量运行所有测试文件
for f in hello ano 999 intro zsbk 1000-pages GBT_33190-2016; do
  ./bin/pdf2ofd-benchmark "test/testdata/${f}.pdf"
done

# 还原度/文本对比（输出位于 bin/output_ofd/）
python3 tools/compare_pdf2ofd.py 96
```