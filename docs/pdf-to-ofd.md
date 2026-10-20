# PDF → OFD 对比测试

> [OFD 库汇总](ofd-libraries.md)：各 OFD 库能力 / 许可证对比表
>
> [OFD → PDF 对比测试](ofd-to-pdf.md)：反方向 OFD→PDF 转换对比

能把 PDF 转回 OFD 的开源库很少。本测试对比 5 个转换器（Go zc310、Go ofdgo、Java ofdrw、Python pdf2ofd、JS jsOFD）在 7 个参考 PDF 上的转换速度、输出 OFD 文件大小与还原质量。

## 测试库

| 转换器         | 语言   | 库                                                                                         | 版本                   | 说明                                                                                                         |
|----------------|--------|--------------------------------------------------------------------------------------------|------------------------|--------------------------------------------------------------------------------------------------------------|
| go-zc310       | Go     | [zc310/ofd](https://github.com/zc310/ofd)                                                  | v0.1.5+（本地 5df8e6f） | 基于 pdfcpu 的 PDF→OFD 导入器（路径/文本/图片/注解/渐变/图案/网格/大纲/元数据），文本保留为可提取 TextObject |
| go-ofdgo       | Go     | [xiaoqidun/ofdgo](https://github.com/xiaoqidun/ofdgo)（ConvertPDF + pdfgo）                | 2026-10-20 版          | 基于 pdfgo 的 PDF→OFD 导入器（路径/文本/图片/注解）；新版严格日期校验，GBT 元数据导入失败            |
| java-ofdrw     | Java   | [ofdrw](https://github.com/ofdrw/ofdrw)（ofdrw-converter PDFConverter）                    | 2.4.0                  | 文本转为矢量轮廓，保留原始外观                                                                               |
| python-pdf2ofd | Python | [pdf2ofd](https://github.com/wanglrebe/pdf2ofd)                                            | 0.1.0                  | 基于 PyMuPDF 渲染，文本栅格化                                                                                |
| node-jsofd     | JS     | [Hufe921/jsOFD](https://github.com/Hufe921/jsOFD)（pdfToOfd + pdfjs-dist）                 | 1.0.1                  | 基于 pdfjs 重放绘制指令流（路径/文本/图片），不转换注解/渐变/图案/Type3/剪裁                                 |

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
| 1000-pages.pdf     | 2.4M | 1000 页 | 多页文档                                       |
| GBT_33190-2016.pdf | 4.8M | 132 页  | GB/T 33190-2016 标准正文（字体全内嵌、图表多） |

## 测试环境

- 系统：Ubuntu 26.04.1 LTS (64-bit)
- 字体目录：`~/.local/share/fonts`（已安装各种中文字体）
- 系统已安装 `apt-get install fonts-wqy-zenhei fonts-noto-cjk`

## 转换速度

| 转换器         | hello.pdf |   ano.pdf |   999.pdf | intro.pdf | zsbk.pdf | 1000-pages.pdf | GBT_33190-2016.pdf | 胜出次数 |
|----------------|----------:|----------:|----------:|----------:|---------:|---------------:|-------------------:|---------:|
| go-zc310       |  **57ms** | **112ms** | **275ms** |  **415ms** |     981ms |  **924ms** |      **1.43s** |        6 |
| go-ofdgo       |      68ms |     325ms |     445ms |     935ms |     1.70s |      10.53s |         FAILED |        0 |
| java-ofdrw     |     1.17s |     1.27s |     1.02s |     4.37s |    2.90s |         26.45s |             12.78s |        0 |
| python-pdf2ofd |     154ms |     800ms |     513ms |     2.39s |  **928ms** |         38.11s |             18.47s |        1 |
| node-jsofd     |     494ms |     712ms |     3.50s |    10.01s |    5.25s |          4.78s |              9.02s |        0 |

## OFD 文件大小

| 转换器         | hello.pdf |  ano.pdf |  999.pdf | intro.pdf | zsbk.pdf | 1000-pages.pdf | GBT_33190-2016.pdf | 胜出次数 |
|----------------|----------:|---------:|---------:|----------:|---------:|---------------:|-------------------:|---------:|
| go-zc310       |  **1.7K** | **415.8K** | **62.6K** |  **7.4M** |  **1.6M** |       **2.8M** |           **7.1M** |        7 |
| go-ofdgo       |      6.2K |   689.0K |    114.3K |      9.2M |     3.0M |           44.9M |            FAILED |        0 |
| java-ofdrw     |      5.6K |     905K |     818K |     34.4M |    19.4M |          63.5M |              30.5M |        0 |
| python-pdf2ofd |      4.6K |     1.3M |     734K |      9.1M |     1.7M |          55.4M |              37.7M |        0 |
| node-jsofd     |      3.0K |     2.2M |   591.6K |     35.6M |    13.2M |          61.8M |              35.1M |        0 |

## 文件大小对比（输出 OFD / 原 PDF）

越小表示压缩率越高，`1.00` 表示与源 PDF 等大；**加粗**为最小。

| 转换器         | hello.pdf |  ano.pdf |  999.pdf | intro.pdf | zsbk.pdf | 1000-pages.pdf | GBT_33190-2016.pdf |
|----------------|----------:|---------:|---------:|----------:|---------:|---------------:|-------------------:|
| go-zc310       |      0.47 |     0.46 | **0.78** | **1.10** | **1.02** | **1.15** | **1.47** |
| go-ofdgo       |      1.77 |     0.76 |     1.43 |     1.38 |     1.91 |          18.65 |            FAILED |
| java-ofdrw     |      1.58 | **1.00** |    10.24 |      5.01 |    12.07 |          25.76 |               6.17 |
| python-pdf2ofd |      1.32 |     1.43 |     9.18 |      1.33 |     1.04 |          22.47 |               7.62 |
| node-jsofd     |  **0.84** |     2.47 |     7.40 |      5.18 |     8.18 |          25.05 |               7.09 |

> 两种文本转图形的方案输出普遍膨胀：ofdrw 在 999、zsbk、1000-pages 上是源 PDF 的 10～26 倍，pdf2ofd 在 999、1000-pages 上是 9～22 倍（zc310 全部 ≤1.5x）；jsOFD 除 hello（0.84 倍）外膨胀 2.5～25 倍（1000-pages 达 25 倍）；ofdgo 在 1000-pages 上达 18.7 倍、GBT 转换失败（见下）

## 还原度对比

将每个转换器输出的 OFD 再经 go-zc310 转回 PDF（回环），与源 PDF 做像素对比（96 DPI）与文本对比。

**像素相似度**（100% 为完全相同）：

| 转换器         |  hello.pdf |    ano.pdf |    999.pdf |  intro.pdf |   zsbk.pdf | 1000-pages.pdf | GBT_33190-2016.pdf | 胜出次数 |
|----------------|-----------:|-----------:|-----------:|-----------:|-----------:|---------------:|-------------------:|---------:|
| go-zc310       |     99.64% | **99.81%** |     98.74% |     **99.71%** |     **99.51%** |         98.29% |              99.92% |        3 |
| go-ofdgo       |     99.64% |     99.80% |     98.82% |          98.19% |     99.38% |         98.09% |             FAILED |        0 |
| java-ofdrw     | **99.85%** |     98.37% |     98.73% |     99.42% |     99.11% |         98.29% |             99.64% |        1 |
| python-pdf2ofd |     99.72% |     98.34% | **98.94%** |     99.36% |     97.62% |         98.38% |             99.12% |        1 |
| node-jsofd     |     99.37% |     90.23% |     98.13% |     99.23% |     98.82% |     **98.60%** |         **99.94%** |        2 |

**文本提取能力**（从转换后 OFD 提取字符数，参考 PDF 提取字符数见括号）：

| 转换器         | hello.pdf(22) | ano.pdf(7091) | 999.pdf(3934) | intro.pdf(5543) | zsbk.pdf(1150) | 1000-pages.pdf(263999) | GBT_33190-2016.pdf(107779) |
|----------------|---------------|---------------|---------------|-----------------|----------------|------------------------|----------------------------|
| go-zc310       | 25            | 10188         | 9833          | 9446            | 2814           | 306998                 | 170697                     |
| go-ofdgo       | 11            | 14127         | 10163         | 10921           | 3375           | 236998                 | FAILED                     |
| java-ofdrw     | 0             | 0             | 0             | 0               | 0              | 0                      | 0                          |
| python-pdf2ofd | 0             | 0             | 0             | 0               | 0              | 0                      | 0                          |
| node-jsofd     | 25            | 11882         | 10039         | 10472           | 2899           | 518998                 | 225836                     |

**文本相似度**（4-gram Jaccard，100% 为与源 PDF 提取文本完全一致）：

| 转换器         | hello.pdf | ano.pdf | 999.pdf | intro.pdf | zsbk.pdf | 1000-pages.pdf | GBT_33190-2016.pdf |
|----------------|----------:|--------:|--------:|----------:|---------:|---------------:|-------------------:|
| go-zc310       |   100.00% |  22.49% |  47.15% |    14.29% |    7.58% |          91.05% |             41.99% |
| go-ofdgo       |    26.32% |  19.19% |  29.63% |    13.68% |    4.63% |           0.38% |             FAILED |
| java-ofdrw     |     0.00% |   0.00% |   0.00% |     0.00% |    0.00% |           0.00% |              0.00% |
| python-pdf2ofd |     0.00% |   0.00% |   0.00% |     0.00% |    0.00% |           0.00% |              0.00% |
| node-jsofd     |   100.00% |  19.08% |  11.27% |    13.58% |    4.50% |          82.15% |              1.08% |

> - ofdrw / pdf2ofd 的输出 OFD 不含可提取文本（Content.xml 无 TextObject）——ofdrw 转成矢量轮廓、pdf2ofd 栅格成位图，故相似度为 0
> - go-zc310 保留文本对象（全部 7 个文件均可提取），文本顺序正确（如「欢迎使用」顺序无误），但字形间插入空格（如 `H e l l o`）使 4-gram 相似度偏低，内容本身完整
> - go-ofdgo 2026-10-20 版 GBT_33190-2016 转换失败（`PDF metadata CreationDate: invalid PDF date`——新版「有限兼容日期，保留严格校验」拒绝 GBT 的非法日期元数据，上版 4.08s 可转换）；ano 回环像素 99.34%→98.09%→99.80% 逼近 go-zc310 99.81%；1000-pages 回环文本仍塌陷（0.38%）；hello/1000-pages 文本提取仍低（11、236998 字符）
> - node-jsofd 逐条重放绘制指令流（路径/文本/图片），不转换注解/渐变/图案/Type3/剪裁：ano 因 Stamp 丢失像素仅 90.23%；其余像素普遍高（98.13～99.94%）但逐字间隔重、文本相似度低（hello/1000-pages 除外，≤19.08%），输出体积偏大（1000-pages 61.8M、GBT 35.1M）
> - 2026-10-20 重测环境恢复 PyMuPDF（fitz 1.26.7）+ numpy，go-zc310 / go-ofdgo 行回环像素与文本相似度回到 fitz 口径；python-pdf2ofd 行因 2026-10-14 系统升级（Python 3.14）依赖丢失 FAILED，沿用旧值
> - 像素相似度基于整页平均色差，对内容稀疏的页面（白底小字）区分度有限，需结合文本提取与人工核对判断
> - intro 的像素对比跳过 14～19 页（数科参考 PDF 这几页背景图丢失，不计入对比）
> - 回环统一使用 go-zc310 的 OFD→PDF，度量的是「输出 OFD 的互操作正确性」而非端到端体验：若改用各库自回环，写错坐标/格式可被自家读取器按同错方式读回而掩盖（如 ofdgo 的 16 位色值只有换读取器才暴露）；代价是引入 zc310 渲染质量偏差，数据仅供参考

## 结论

> **本轮重测关键变化（vs 上轮）：**
> - go-ofdgo 更新至 20261020104500（API 重构）后 GBT_33190-2016 转换回退为 FAILED：新严格日期校验拒绝 GBT 元数据非法 CreationDate（上版 4.08s/15.5M 成功），本轮 6/7 成功
> - go-ofdgo ano 回环像素大幅回升：99.34%→99.80%，距 go-zc310 99.81% 仅 0.01pp；其余像素/文本基本持平（1000-pages 回环文本仍 0.38%）
> - go-zc310 各项与上版持平（速度 6 胜、体积 7 胜、像素 3 胜保持；1000-pages 提速 938ms→924ms），GBT 像素 99.92% 仍微逊 node-jsofd 99.94%
> - go 行质量指标恢复 fitz 口径（apt 安装 python3-fitz 1.26.7 + numpy），与旧 fitz 行同口径可比；python-pdf2ofd 因 2026-10-14 系统升级依赖丢失 FAILED，沿用旧值

- **go-zc310**：唯一保留可复制文本（顺序正确）且可用的转换器，速度 6 胜（hello 57ms、ano 112ms、999 275ms、intro 415ms、1000-pages 924ms、GBT 1.43s）、体积 7 胜全部最小（1.7K～7.1M），像素还原 3 胜（ano 99.81%、intro 99.71%、zsbk 99.51%）
- **go-ofdgo**：2026-10-20 版无速度/像素胜出（像素最高 ano 99.80%）；GBT_33190-2016 转换失败（严格日期校验回退）；1000-pages 回环文本相似度仅 0.38%；体积膨胀明显（1000-pages 44.9M、intro 9.2M）、1000-pages 转换 10.53s
- **java-ofdrw**：像素还原 1 胜（hello 99.85%），文本转为轮廓不可提取；输出体积大（1000-pages 63.5M、GBT 30.5M）、多页转换慢（GBT 12.78s）
- **python-pdf2ofd**：简单易用，单页/小文件还原不错（999 像素最高），但纯位图渲染不可提取文本、大文件最慢（1000-pages 38.1s、GBT 18.5s）
- **node-jsofd**：像素普遍接近源 PDF（1000-pages 98.60%、GBT 99.94% 两列最高，除 ano 外 ≥98.1%），但体积大（1000-pages 61.8M、GBT 35.1M）且大文件偏慢；文本相似度低（hello/1000-pages 除外，≤19.08%），ano 因 Stamp 未转换像素仅 90.23%

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

# 运行全部转换器（单个文件）
./bin/pdf2ofd-benchmark testdata/suwell/hello.pdf

# 批量运行所有测试文件
for f in testdata/suwell/*.pdf testdata/standards/*.pdf; do
  ./bin/pdf2ofd-benchmark "$f"
done

# 还原度/文本对比（输出位于 bin/output_ofd/）
python3 tools/compare_pdf2ofd.py 96
```