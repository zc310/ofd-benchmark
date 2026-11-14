# PDF → OFD 对比测试

> [OFD 库汇总](ofd-libraries.md)：各 OFD 库能力 / 许可证对比表
>
> [OFD → PDF 对比测试](ofd-to-pdf.md)：反方向 OFD→PDF 转换对比



## 测试库

| 转换器         | 语言   | 库                                                                          | 版本                    | 说明                                                                                                         |
|----------------|--------|-----------------------------------------------------------------------------|-------------------------|--------------------------------------------------------------------------------------------------------------|
| go-zc310       | Go     | [zc310/ofd](https://github.com/zc310/ofd)                                   | v0.1.9 | 基于 pdfcpu 的 PDF→OFD 导入器（路径/文本/图片/注解/渐变/图案/网格/大纲/元数据），文本保留为可提取 TextObject |
| go-ofdgo       | Go     | [xiaoqidun/ofdgo](https://github.com/xiaoqidun/ofdgo)（ConvertPDF + pdfgo） | 2026-11-11（3edf35199ec7）           | 基于 pdfgo 的 PDF→OFD 导入器（路径/文本/图片/注解）                                                          |
| java-ofdrw     | Java   | [ofdrw](https://github.com/ofdrw/ofdrw)（ofdrw-converter PDFConverter）     | 2.5.1                   | 文本转为矢量轮廓，保留原始外观                                                                               |
| python-pdf2ofd | Python | [pdf2ofd](https://github.com/wanglrebe/pdf2ofd)                             | 0.1.0                   | 基于 PyMuPDF 渲染，文本栅格化                                                                                |
| node-jsofd     | JS     | [Hufe921/jsOFD](https://github.com/Hufe921/jsOFD)（pdfToOfd + pdfjs-dist）  | 1.0.1                   | 基于 pdfjs 重放绘制指令流（路径/文本/图片），不转换注解/渐变/图案/Type3/剪裁                                 |

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
| go-zc310       |  **61ms** | **111ms** | **281ms** |  **411ms** |     959ms |  **934ms** |      **1.41s** |        6 |
| go-ofdgo       |      68ms |     328ms |     431ms |     817ms |     1.48s |      12.89s |          4.15s |        0 |
| java-ofdrw     |     566ms |     1.00s |     1.31s |     5.87s |    3.63s |          2.97s |             56.32s |        0 |
| python-pdf2ofd |     154ms |     800ms |     513ms |     2.39s |  **928ms** |         38.11s |             18.47s |        1 |
| node-jsofd     |     494ms |     712ms |     3.50s |    10.01s |    5.25s |          4.78s |              9.02s |        0 |

## OFD 文件大小

| 转换器         | hello.pdf |  ano.pdf |  999.pdf | intro.pdf | zsbk.pdf | 1000-pages.pdf | GBT_33190-2016.pdf | 胜出次数 |
|----------------|----------:|---------:|---------:|----------:|---------:|---------------:|-------------------:|---------:|
| go-zc310       |  **1.7K** | **415.8K** | **62.6K** |  **7.4M** |  **1.6M** |       **4.8M** |           **7.1M** |        7 |
| go-ofdgo       |      6.2K |   709.5K |    117.5K |      9.2M |     3.0M |           60.8M |            17.1M |        0 |
| java-ofdrw     |      5.6K |     601K |     818K |     34.2M |    20.4M |           3.5M |              30.5M |        0 |
| python-pdf2ofd |      4.6K |     1.3M |     734K |      9.1M |     1.7M |          55.4M |              37.7M |        0 |
| node-jsofd     |      3.0K |     2.2M |   591.6K |     35.6M |    13.2M |          61.8M |              35.1M |        0 |

## 文件大小对比（输出 OFD / 原 PDF）

越小表示压缩率越高，`1.00` 表示与源 PDF 等大；**加粗**为最小。

| 转换器         | hello.pdf |  ano.pdf |  999.pdf | intro.pdf | zsbk.pdf | 1000-pages.pdf | GBT_33190-2016.pdf |
|----------------|----------:|---------:|---------:|----------:|---------:|---------------:|-------------------:|
| go-zc310       |      0.47 |     0.46 | **0.78** | **1.10** | **1.02** |          2.54 | **1.47** |
| go-ofdgo       |      1.76 | **0.78** |     1.47 |     1.38 |     1.91 |          31.98 |            3.54 |
| java-ofdrw     |      1.58 |     0.67 |    10.23 |      5.10 |    13.00 |       **1.83** |               6.32 |
| python-pdf2ofd |      1.32 |     1.43 |     9.18 |      1.33 |     1.04 |          22.47 |               7.62 |
| node-jsofd     |  **0.84** |     2.47 |     7.40 |      5.18 |     8.18 |          25.05 |               7.09 |

## 还原度对比

将每个转换器输出的 OFD 再经 go-zc310 转回 PDF（回环），与源 PDF 做像素对比（96 DPI）与文本对比。

**像素相似度**（100% 为完全相同）：

| 转换器         |  hello.pdf |    ano.pdf |    999.pdf |  intro.pdf |   zsbk.pdf | 1000-pages.pdf | GBT_33190-2016.pdf | 胜出次数 |
|----------------|-----------:|-----------:|-----------:|-----------:|-----------:|---------------:|-------------------:|---------:|
| go-zc310       |     99.64% | **99.81%** |     98.74% |     **99.71%** |     **99.51%** |         99.83% |              99.92% |        4 |
| go-ofdgo       |     99.64% |     99.80% |     98.82% |          98.19% |     99.38% |         99.10% |             99.79% |        1 |
| java-ofdrw     | **99.85%** |     98.60% |     98.69% |     99.50% |     99.08% |         96.81% |             99.60% |        1 |
| python-pdf2ofd |     99.72% |     98.34% | **98.94%** |     99.36% |     97.62% |         98.38% |             99.12% |        1 |
| node-jsofd     |     99.37% |     90.23% |     98.13% |     99.23% |     98.82% |     **98.60%** |         **99.94%** |        1 |

**文本提取能力**（从转换后 OFD 提取字符数，参考 PDF 提取字符数见括号）：

| 转换器         | hello.pdf(22) | ano.pdf(7091) | 999.pdf(3934) | intro.pdf(5543) | zsbk.pdf(1150) | 1000-pages.pdf(263999) | GBT_33190-2016.pdf(107779) |
|----------------|---------------|---------------|---------------|-----------------|----------------|------------------------|----------------------------|
| go-zc310       | 25            | 10188         | 9833          | 9446            | 2814           | 375387                 | 170697                     |
| go-ofdgo       | 11            | 14127         | 10163         | 10921           | 3375           | 180998                 | 234775                     |
| java-ofdrw     | 0             | 0             | 0             | 0               | 0              | 0                      | 0                          |
| python-pdf2ofd | 0             | 0             | 0             | 0               | 0              | 0                      | 0                          |
| node-jsofd     | 25            | 11882         | 10039         | 10472           | 2899           | 518998                 | 225836                     |

**文本相似度**（4-gram Jaccard，100% 为与源 PDF 提取文本完全一致）：

| 转换器         | hello.pdf | ano.pdf | 999.pdf | intro.pdf | zsbk.pdf | 1000-pages.pdf | GBT_33190-2016.pdf |
|----------------|----------:|--------:|--------:|----------:|---------:|---------------:|-------------------:|
| go-zc310       |   100.00% |  22.49% |  47.15% |    14.29% |    7.58% |          32.68% |             41.99% |
| go-ofdgo       |    26.32% |  19.19% |  29.63% |    13.68% |    4.63% |           2.59% |              1.07% |
| java-ofdrw     |     0.00% |   0.00% |   0.00% |     0.00% |    0.00% |           0.00% |              0.00% |
| python-pdf2ofd |     0.00% |   0.00% |   0.00% |     0.00% |    0.00% |           0.00% |              0.00% |
| node-jsofd     |   100.00% |  19.08% |  11.27% |    13.58% |    4.50% |          82.15% |              1.08% |


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
