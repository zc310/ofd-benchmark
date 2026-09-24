# PDF → OFD 对比测试

> [OFD 库汇总](ofd-libraries.md)：各 OFD 库能力 / 许可证对比表
>
> [OFD → PDF 对比测试](ofd-to-pdf.md)：反方向 OFD→PDF 转换对比

能把 PDF 转回 OFD 的开源库很少。本测试对比 4 个可用转换器（Go zc310、Rust easyofd、Java ofdrw、Python pdf2ofd）在 6 个参考 PDF 上的转换速度、输出 OFD 文件大小与还原质量。

## 测试库

| 转换器 | 语言 | 库 | 版本 | 说明 |
|--------|------|----|------|------|
| go-zc310 | Go | [zc310/ofd](https://github.com/zc310/ofd) | v0.1.3 起（pdfimport） | 基于 pdfcpu 的 PDF→OFD 导入器，文本保留为可提取 TextObject |
| rust-easyofd | Rust | [easyofd-rust](https://github.com/easy-4-rust/easyofd-rust)（easyofd-convert PdfImporter） | 0.1.3 | 基于 lopdf 提取文本/图片；不转换矢量路径 |
| java-ofdrw | Java | [ofdrw](https://github.com/ofdrw/ofdrw)（ofdrw-converter PDFConverter） | 2.4.0 | 文本转为矢量轮廓，保留原始外观 |
| python-pdf2ofd | Python | [pdf2ofd](https://github.com/wanglrebe/pdf2ofd) | 0.1.0 | 基于 PyMuPDF 渲染，文本栅格化 |

> `python-pdf2ofd` 安装：`pip install pdf2ofd`（PEP 668 环境需 `--user --break-system-packages` 或使用 venv）

## 测试文件

直接使用数科版式阅读器导出的参考 PDF（即 OFD→PDF 测试的对比基准），保证与 OFD→PDF 测试同一组文档。

| 文件           | 大小   | 页数 | 说明         |
|----------------|--------|------|--------------|
| hello.pdf      | 3.5K   | 1 页 | 简单文本     |
| ano.pdf        | 904K   | 3 页 | 发票         |
| 999.pdf        | 80K    | 5 页 | 简单文档     |
| intro.pdf      | 6.7M   | 42 页 | 介绍文档     |
| zsbk.pdf       | 1.6M   | 2 页 | 签章文件     |
| 1000-pages.pdf | 2.1M   | 1000 页 | 多页文档   |

## 测试环境

- 系统：Ubuntu 26.04.1 LTS (64-bit)
- 字体目录：`~/.local/share/fonts`（已安装各种中文字体）
- 系统已安装 `apt-get install fonts-wqy-zenhei fonts-noto-cjk`

## 转换速度

| 转换器        | hello.pdf |  ano.pdf |  999.pdf | intro.pdf |  zsbk.pdf | 1000-pages.pdf | 胜出次数 |
|---------------|----------:|---------:|---------:|----------:|----------:|---------------:|---------:|
| go-zc310      |    47ms  |    89ms  |    68ms  |    1.15s  |    342ms  |     1.99s  | 0 |
| rust-easyofd  |  **2.6ms** |  **3.6ms** |  **2.6ms** |   **104ms** |     **23ms** |    **277ms**  | 6 |
| java-ofdrw    |  496ms  |   1.25s  |   1.04s  |    4.40s  |   2.92s   |    26.09s  | 0 |
| python-pdf2ofd|  106ms  |   817ms  |   531ms  |    2.43s  |   943ms   |    34.49s  | 0 |

> rust-easyofd 的「最快」源于几乎未提取内容（hello/ano 输出空页），速度优势无实际意义

## OFD 文件大小

| 转换器        | hello.pdf |   ano.pdf |  999.pdf | intro.pdf |  zsbk.pdf | 1000-pages.pdf | 胜出次数 |
|---------------|----------:|----------:|---------:|----------:|----------:|---------------:|---------:|
| go-zc310      |    1.6K  |    381K   |   48.3K  |    7.4M   |    1.5M   |      3.2M    | 0 |
| rust-easyofd  |  **1.2K** |   **1.8K** |   **3.6K** |    **5.4M** |   **919K** |      **432K**  | 6 |
| java-ofdrw    |    5.2K  |   905K   |   818K   |   34.4M  |   19.4M  |     57.4M  | 0 |
| python-pdf2ofd|    4.3K  |   1.3M   |   734K   |    9.1M   |    1.7M   |     44.8M  | 0 |

> rust-easyofd「体积最小」同样是内容丢失的结果（hello 1.2K/ano 1.8K 实为空页），不具可比性

## 文件大小对比（输出 OFD / 原 PDF）

越小表示压缩率越高，`1.00` 表示与源 PDF 等大；**加粗**为最小。

| 转换器        | hello.pdf |  ano.pdf | 999.pdf | intro.pdf | zsbk.pdf | 1000-pages.pdf |
|---------------|----------:|---------:|--------:|----------:|---------:|---------------:|
| go-zc310      |    0.46  |   0.42  |   0.60  |   1.08  |   0.93  |      1.48    |
| rust-easyofd  |  **0.33** |  **0.00** |  **0.04** |  **0.78** |  **0.57** |     **0.20**  |
| java-ofdrw    |    1.48  |   1.00  |  10.24  |   5.01  |  12.07  |     26.48  |
| python-pdf2ofd|    1.22  |   1.43  |   9.18  |   1.33  |   1.04  |     20.63  |

> 两种文本转图形的方案输出普遍膨胀：ofdrw 在 999、zsbk、1000-pages 上是源 PDF 的 10~26 倍，pdf2ofd 在 999、1000-pages 上是 9~21 倍（zc310 全部 ≤1.5x）；rust-easyofd 的 0.00~0.78 是内容丢失而非压缩

## 还原度对比

将每个转换器输出的 OFD 再经 go-zc310 转回 PDF（回环），与源 PDF 做像素对比（96 DPI）与文本对比。

**像素相似度**（100% 为完全相同）：

| 转换器        | hello.pdf |  ano.pdf |  999.pdf | intro.pdf |  zsbk.pdf | 1000-pages.pdf | 胜出次数 |
|---------------|----------:|---------:|---------:|----------:|----------:|---------------:|---------:|
| go-zc310      |  99.95% | **99.62%** |  98.24% |  95.91% | **99.32%** |  99.59% | 2 |
| rust-easyofd  |  99.84% |  98.56% |  95.77% |  43.71% |  95.62% |  73.67% | 0 |
| java-ofdrw    | **99.98%** |  98.37% |  98.72% | **99.34%** |  99.22% | **99.62%** | 3 |
| python-pdf2ofd|  99.94% |  98.34% | **98.94%** |  93.23% |  97.61% |  99.31% | 1 |

**文本提取能力**（从转换后 OFD 提取字符数，参考 PDF 提取字符数见括号）：

| 转换器        | hello.pdf(25) | ano.pdf(7139) | 999.pdf(3934) | intro.pdf(5546) | zsbk.pdf(1150) | 1000-pages.pdf(323700) |
|---------------|---------------|---------------|---------------|-----------------|----------------|------------------------|
| go-zc310      |      37       |     9839      |     9759      |      7612       |      2712      |       423998          |
| rust-easyofd  |       0       |       0       |     242*      |        0        |       10*      |        10996*         |
| java-ofdrw    |       0       |       0       |       0       |        0        |        0       |          0            |
| python-pdf2ofd|       0       |       0       |       0       |        0        |        0       |          0            |

**文本相似度**（4-gram Jaccard，100% 为与源 PDF 提取文本完全一致）：

| 转换器        | hello.pdf |  ano.pdf |  999.pdf | intro.pdf |  zsbk.pdf | 1000-pages.pdf |
|---------------|----------:|---------:|---------:|----------:|----------:|---------------:|
| go-zc310      |   17.02% |  33.33% |  51.26% |   45.20% |    7.86% |       8.76% |
| rust-easyofd  |    0.00% |   0.00% |   0.00% |    0.00% |    0.00% |       0.00% |
| java-ofdrw    |    0.00% |   0.00% |   0.00% |    0.00% |    0.00% |       0.00% |
| python-pdf2ofd|    0.00% |   0.00% |   0.00% |    0.00% |    0.00% |       0.00% |

> - ofdrw / pdf2ofd 的输出 OFD 不含可提取文本（Content.xml 无 TextObject）——ofdrw 转成矢量轮廓、pdf2ofd 栅格成位图，故相似度为 0
> - go-zc310 保留文本对象（全部 6 个文件均可提取），文本顺序正确（如「欢迎使用」顺序无误），但字形间插入空格（如 `H e l l o`）使 4-gram 相似度偏低，内容本身完整
> - rust-easyofd 实测质量最差：hello/ano 输出空页（Layer 为空）；999、1000-pages 提取的文本为 CID 字体乱码（标 \* 的 242/10996/10 字符）且版面固定在 (10,20)；像素高分（hello 99.84%、ano 98.56%）是因为空白页占主导，不代表还原正确
> - 像素相似度基于整页平均色差，对内容稀疏的页面（白底小字）区分度有限，需结合文本提取与人工核对判断
> - 回环对比引入了 go-zc310 的 OFD→PDF 渲染质量，数据仅供参考

## 结论

- **go-zc310**：唯一保留可复制文本（顺序正确）且可用的转换器，像素还原 2 胜（ano、zsbk），在可用转换器中速度/体积最优（47ms~1.99s、1.6K~3.2M）
- **rust-easyofd**：速度与体积数值均第一（各 6 胜）但为虚假优势——hello/ano 输出空页、999/1000-pages 文本为 CID 乱码、intro 像素仅 43.71%，实际质量最差
- **java-ofdrw**：像素还原度最高（3胜，外观最接近源 PDF），文本转为轮廓不可提取；输出体积大（1000-pages 57MB）、多页转换慢（26s）
- **python-pdf2ofd**：简单易用，单页/小文件还原不错（999 像素最高），但纯位图渲染不可提取文本、大文件最慢（34.5s）

## 复现

```bash
# 编译 zc310 ofd-converter（PDF→OFD / OFD→PDF / OFD→TXT）
cd /home/go/workspace/ofd
go build -o ../ofd-benchmark/bin/ofd-converter ./cmd/ofd-converter

cd ../ofd-benchmark
# 编译基准测试驱动
go build -o bin/pdf2ofd-benchmark ./cmd/pdf2ofd-benchmark

# 编译 Rust PDF→OFD 转换器（easyofd）
cargo build --release --manifest-path rust/pdf2ofd/Cargo.toml
cp rust/pdf2ofd/target/release/pdf2ofd bin/easyofd-pdf2ofd

# 运行全部转换器（单个文件）
./bin/pdf2ofd-benchmark test/testdata/hello.pdf

# 批量运行所有测试文件
for f in hello ano 999 intro zsbk 1000-pages; do
  ./bin/pdf2ofd-benchmark "test/testdata/${f}.pdf"
done

# 还原度/文本对比（输出位于 bin/output_ofd/）
python3 tools/compare_pdf2ofd.py 96
```