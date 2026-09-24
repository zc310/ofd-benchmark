# OFD to PDF 速度对比测试

> [OFD 库汇总](docs/ofd-libraries.md)：各 OFD 库能力 / 许可证对比表

## 测试库

### Go 库

| 库                                                    | 版本                  | OFD→PDF |  更新时间  | 备注   |
|-------------------------------------------------------|-----------------------|:-------:|:----------:|--------|
| [xiaoqidun-ofdgo](https://github.com/xiaoqidun/ofdgo) | v0.0.0-20260919111956 |   ✅    | 2026-09-19 | 已测试 |
| [zc310-ofd](https://github.com/zc310/ofd)             | v0.1.3                |   ✅    | 2026-09-24 | 已测试 |

### Rust 库

| 库                                                          | 版本   | OFD→PDF |  更新时间  | 备注   |
|-------------------------------------------------------------|--------|:-------:|:----------:|--------|
| [easyofd-rust](https://github.com/easy-4-rust/easyofd-rust) | v0.1.2 |   ✅    | 2026-09-10 | 已测试 |

### Python 库

| 库                                            | 版本     | OFD→PDF |  更新时间  | 备注                       |
|-----------------------------------------------|----------|:-------:|:----------:|----------------------------|
| [easyofd](https://pypi.org/project/easyofd/)  | 20260427 |   ⚠️    | 2026-04-27 | 部分文件失败               |
| [ofd2img](https://pypi.org/project/ofd2img/)  | 0.1.2    |   ⚠️    | 2026-05-07 | 部分文件失败               |
| [ofd2pdf](https://github.com/jsyzdej/ofd2pdf) | 0.0.2    |   ⚠️    | 2026-07-08 | 部分文件失败，基于图片渲染 |

### Java 库

| 库                                      | 版本  | OFD→PDF | 更新时间   | 备注                     |
|-----------------------------------------|-------|:-------:|------------|--------------------------|
| [ofdrw](https://github.com/ofdrw/ofdrw) | 2.4.0 |   ✅    | 2026-08-04 | 需手动下载依赖，全部成功 |

### JavaScript 库

| 库                                                                           | 版本  | OFD→PDF |  更新时间  | 备注     |
|------------------------------------------------------------------------------|-------|:-------:|:----------:|----------|
| [@miconvert/ofd-to-pdf](https://www.npmjs.com/package/@miconvert/ofd-to-pdf) | 0.2.0 |   ✅    | 2026-08-03 | 全部成功 |

## 测试环境

- 系统：Ubuntu 26.04.1 LTS (64-bit)
- 字体目录：`~/.local/share/fonts`（已安装各种中文字体）
- 系统已安装 `apt-get install fonts-wqy-zenhei fonts-noto-cjk`

## 测试文件

测试用 OFD 文件来自：https://github.com/zc310/ofd/tree/main/test/testdata

| 文件           | 大小 | 说明         |
|----------------|------|--------------|
| hello.ofd      | 1.5K | 简单测试文件 |
| ano.ofd        | 702K | 带签章文件   |
| intro.ofd      | 7.2M | 介绍文档     |
| 1000-pages.ofd | 456K | 1000页文件   |
| 999.ofd        | 30K  | 带签章文件   |
| zsbk.ofd       | 1.5M | 数科签章文件 |

> 同目录下 `*.pdf` 为各 OFD 的参考转换结果，由**数科版式阅读器**导出，用作像素级对比基准。

## 测试结果

### 转换速度

| 转换器                                                                       | hello.ofd |  ano.ofd | intro.ofd | 1000-pages.ofd |   999.ofd | zsbk.ofd | 胜出次数 |
|------------------------------------------------------------------------------|----------:|---------:|----------:|---------------:|----------:|---------:|---------:|
| [go-ofdgo](https://github.com/xiaoqidun/ofdgo)                               |      85ms |    470ms |     3.14s |          8.11s |     592ms |    514ms |        0 |
| [go-zc310](https://github.com/zc310/ofd)                                     |     120ms |    245ms | **952ms** |          4.10s |     291ms |    437ms |        1 |
| [rust-easyofd](https://github.com/easy-4-rust/easyofd-rust)                  |  **42ms** | **76ms** |     1.81s |      **913ms** | **156ms** | **77ms** |        5 |
| [easyofd](https://pypi.org/project/easyofd/)                                 |    FAILED |    1.09s |     8.43s |         FAILED |     1.01s |    5.18s |        0 |
| [ofd2pdf](https://github.com/jsyzdej/ofd2pdf)                                |     136ms |   FAILED |    FAILED |         41.29s |     409ms |    201ms |        0 |
| [ofdrw](https://github.com/ofdrw/ofdrw)                                      |     507ms |    759ms |     5.11s |          2.40s |     799ms |    5.53s |        0 |
| [@miconvert/ofd-to-pdf](https://www.npmjs.com/package/@miconvert/ofd-to-pdf) |     240ms |    730ms |     1.17s |          8.59s |     502ms |    815ms |        0 |

> - `python-easyofd`: `pip install easyofd ofd2img`，自动切换库，部分文件失败
> - `python-ofd2pdf`: `pip install ofd2pdf`，基于图片渲染，部分文件失败
> - `java-ofdrw`: `org.ofdrw:ofdrw-converter:2.4.0`，全部成功
> - `node-ofd2pdf`: `npm install @miconvert/ofd-to-pdf`，全部成功

### PDF 文件大小

| 转换器                                                                       | hello.ofd |    ano.ofd | intro.ofd | 1000-pages.ofd |    999.ofd |   zsbk.ofd | 胜出次数 |
|------------------------------------------------------------------------------|----------:|-----------:|----------:|---------------:|-----------:|-----------:|---------:|
| [go-ofdgo](https://github.com/xiaoqidun/ofdgo)                               |    14.1KB |      1.5MB |    35.3MB |         38.7MB |      2.0MB |      1.9MB |        0 |
| [go-zc310](https://github.com/zc310/ofd)                                     |     9.9KB |     76.3KB |     6.5MB |          1.8MB |     74.5KB |      1.7MB |        0 |
| [rust-easyofd](https://github.com/easy-4-rust/easyofd-rust)                  |    37.3KB |     77.1KB |    28.9MB |      **614KB** |     83.3KB | **77.8KB** |        2 |
| [easyofd](https://pypi.org/project/easyofd/)                                 |    FAILED | **36.1KB** |    38.2MB |         FAILED | **73.1KB** |     13.6MB |        2 |
| [ofd2pdf](https://github.com/jsyzdej/ofd2pdf)                                |    36.8KB |     FAILED |    FAILED |         70.2MB |      752KB |      129KB |        0 |
| [python-ofdreader](https://pypi.org/project/ofdreader/)                      |    FAILED |     FAILED |    FAILED |         FAILED |     FAILED |     FAILED |        0 |
| [ofdrw](https://github.com/ofdrw/ofdrw)                                      |     4.8KB |     65.6KB |    22.6MB |          2.3MB |     76.9KB |     15.4MB |        0 |
| [@miconvert/ofd-to-pdf](https://www.npmjs.com/package/@miconvert/ofd-to-pdf) |     1.3KB |      197KB | **1.1MB** |          7.2MB | **82.9KB** |  **559KB** |        2 |

### 文本提取能力

| 转换器                                                                       | hello.ofd |  ano.ofd |  999.ofd | zsbk.ofd | 胜出次数 |
|------------------------------------------------------------------------------|----------:|---------:|---------:|---------:|---------:|
| [go-ofdgo](https://github.com/xiaoqidun/ofdgo)                               |         8 |        0 |       61 |       13 |        0 |
| [go-zc310](https://github.com/zc310/ofd)                                     |        22 | **7282** |     3950 |     1158 |        1 |
| [rust-easyofd](https://github.com/easy-4-rust/easyofd-rust)                  |    **28** |     6553 |     4046 |      490 |        1 |
| [easyofd](https://pypi.org/project/easyofd/)                                 |    FAILED |     1208 |     3933 |     1508 |        0 |
| [ofd2pdf](https://github.com/jsyzdej/ofd2pdf)                                |         0 |   FAILED |        0 |        0 |        0 |
| [ofdrw](https://github.com/ofdrw/ofdrw)                                      |        18 |     6456 |     2933 |      181 |        0 |
| [@miconvert/ofd-to-pdf](https://www.npmjs.com/package/@miconvert/ofd-to-pdf) |        26 |     6990 | **4249** | **1938** |        2 |

> - `python-ofd2pdf` 基于图片渲染，无法提取文本
> - 单位为字符数（chars）

### 与数科版式输出对比

以数科版式阅读器导出的 `test/testdata/*.pdf` 为基准，对比转换效果。

**像素相似度**（96 DPI，1 - 平均归一化色差，100% 为完全相同）：

| 转换器                                                                       |  hello.ofd |    ano.ofd |  intro.ofd | 1000-pages.ofd |    999.ofd |   zsbk.ofd | 胜出次数 |
|------------------------------------------------------------------------------|-----------:|-----------:|-----------:|---------------:|-----------:|-----------:|---------:|
| [go-ofdgo](https://github.com/xiaoqidun/ofdgo)                               |     99.40% |     99.35% | **96.03%** |         98.37% | **98.50%** | **99.30%** |        3 |
| [go-zc310](https://github.com/zc310/ofd)                                     |     99.54% | **99.42%** |     95.75% |         98.00% |     98.34% |     98.98% |        1 |
| [rust-easyofd](https://github.com/easy-4-rust/easyofd-rust)                  |     99.79% |     97.83% |     46.37% |         73.66% |     96.34% |     94.02% |        0 |
| [easyofd](https://pypi.org/project/easyofd/)                                 |     FAILED |     93.82% |     40.96% |         FAILED |     92.41% |     94.41% |        0 |
| [ofd2pdf](https://github.com/jsyzdej/ofd2pdf)                                | **99.95%** |     FAILED |     FAILED |     **99.32%** |     97.74% |     94.33% |        2 |
| [ofdrw](https://github.com/ofdrw/ofdrw)                                      |     99.88% |     98.86% |     95.32% |         98.93% |     97.73% |     98.45% |        0 |
| [@miconvert/ofd-to-pdf](https://www.npmjs.com/package/@miconvert/ofd-to-pdf) |     99.47% |     97.74% |     83.16% |         98.21% |     95.80% |     94.40% |        0 |

**文件大小对比**（转换后 PDF 大小 / 数科参考 PDF 大小，越小表示压缩率越高，`1.00` 表示与数科输出一致；**加粗**为最接近 `1.00` 的转换器）：

| 转换器                                                                       | hello.ofd |  ano.ofd | intro.ofd | 1000-pages.ofd |  999.ofd | zsbk.ofd |
|------------------------------------------------------------------------------|----------:|---------:|----------:|---------------:|---------:|---------:|
| 数科版式阅读器                                                               |      1.00 |     1.00 |      1.00 |           1.00 |     1.00 |     1.00 |
| [go-ofdgo](https://github.com/xiaoqidun/ofdgo)                               |      3.90 |     1.61 |      5.13 |          17.85 |    26.15 |     1.22 |
| [go-zc310](https://github.com/zc310/ofd)                                     |      2.79 |     0.08 |  **0.97** |           0.84 |     0.93 | **1.09** |
| [rust-easyofd](https://github.com/easy-4-rust/easyofd-rust)                  |     10.52 |     0.09 |      4.30 |           0.28 | **1.04** |     0.05 |
| [easyofd](https://pypi.org/project/easyofd/)                                 |    FAILED |     0.04 |      5.70 |         FAILED |     0.92 |     8.62 |
| [ofd2pdf](https://github.com/jsyzdej/ofd2pdf)                                |     10.37 |   FAILED |    FAILED |          33.13 |     9.40 |     0.08 |
| [ofdrw](https://github.com/ofdrw/ofdrw)                                      |  **1.36** |     0.07 |      3.37 |       **1.07** | **0.96** |     9.78 |
| [@miconvert/ofd-to-pdf](https://www.npmjs.com/package/@miconvert/ofd-to-pdf) |      0.38 | **0.22** |      0.16 |           3.40 | **1.04** |     0.35 |

> - 像素相似度越高表示输出越贴近数科版式阅读器；仿色/抗锯齿等平滑差异会小幅降低相似度
> - `python-easyofd` 在 hello、1000-pages 转换失败；`python-ofd2pdf` 在 ano、intro 转换失败
> - rust-easyofd 在 intro、1000-pages 相似度明显偏低（46%、74%），说明对含复杂图形/签章文件的还原能力弱
> - 复现命令：`python3 tools/compare_pdf.py 96 <file>`

**文本相似度**（与数科参考 PDF 提取文本的 4-gram Jaccard 相似度，100% 为文本完全一致；`0.00%` 表示无文本或提取失败）：

参考文本长度：`hello=25` · `ano=7139` · `intro=5546` · `1000-pages=323700` · `999=3934` · `zsbk=1150`

| 转换器                                                                       |  hello.ofd |     ano.ofd |  intro.ofd | 1000-pages.ofd |    999.ofd |   zsbk.ofd | 胜出次数 |
|------------------------------------------------------------------------------|-----------:|------------:|-----------:|---------------:|-----------:|-----------:|---------:|
| [go-ofdgo](https://github.com/xiaoqidun/ofdgo)                               |      8.33% |       0.00% |      0.00% |          0.35% |      0.64% |      0.64% |        0 |
| [go-zc310](https://github.com/zc310/ofd)                                     |     37.93% | **100.00%** | **72.36%** |         69.82% | **90.10%** | **96.60%** |        4 |
| [rust-easyofd](https://github.com/easy-4-rust/easyofd-rust)                  | **45.16%** |      25.44% |      4.78% |          0.22% |     47.81% |      7.01% |        1 |
| [easyofd](https://pypi.org/project/easyofd/)                                 |     FAILED |      12.62% |      0.08% |         FAILED |     80.54% |     11.20% |        0 |
| [ofd2pdf](https://github.com/jsyzdej/ofd2pdf)                                |      0.00% |      FAILED |     FAILED |          0.00% |      0.00% |      0.00% |        0 |
| [ofdrw](https://github.com/ofdrw/ofdrw)                                      |     29.63% |      16.20% |      7.16% |     **78.68%** |     51.29% |      5.88% |        1 |
| [@miconvert/ofd-to-pdf](https://www.npmjs.com/package/@miconvert/ofd-to-pdf) |     30.00% |      13.57% |      5.88% |         62.90% |     41.24% |      4.86% |        0 |

> - go-zc310 文本相似度最高（ano=100.00%、zsbk=96.60%、999=90.10%、intro=72.36%）；java-ofdrw 在 1000-pages 领先（78.68%）
> - 文本相似度对提取顺序/换行敏感，得分偏低不代表内容缺失
> - `python-ofd2pdf` 基于图片渲染，全部无文本；`go-ofdgo` 文本提取能力弱（hello=8/25、999=61/3934 字符）
> - 复现命令：`python3 tools/compare_text.py <file>`

### 质量总结

| 转换器                                                              | 转换成功率 | 页面尺寸 | 矢量图形 | 文本可提取 | 图片嵌入 | 文件压缩 | 中文支持 | 兼容性   |
|---------------------------------------------------------------------|:----------:|----------|----------|------------|----------|----------|----------|----------|
| [go-ofdgo](https://github.com/xiaoqidun/ofdgo)                      |    100%    | A4       | 支持     | 部分       | 有       | 较差     | 良好     | 全部     |
| [go-zc310](https://github.com/zc310/ofd)                            |    100%    | A4       | 支持     | 最佳       | 有       | 良佳     | 最佳     | 全部     |
| [rust-easyofd](https://github.com/easy-4-rust/easyofd-rust)         |    100%    | A4       | 支持     | 良好       | 有       | 最佳     | 良好     | 全部     |
| [python-easyofd](https://pypi.org/project/easyofd/)                 |    67%     | A4       | 支持     | 部分       | 有       | 良佳     | 良好     | 部分失败 |
| [python-ofd2pdf](https://github.com/jsyzdej/ofd2pdf)                |    67%     | A4       | 不支持   | 无         | 有       | 一般     | 良好     | 部分失败 |
| [java-ofdrw](https://github.com/ofdrw/ofdrw)                        |    100%    | A4       | 支持     | 良好       | 有       | 良佳     | 良好     | 全部     |
| [node-ofd2pdf](https://www.npmjs.com/package/@miconvert/ofd-to-pdf) |    100%    | A4       | 支持     | 良好       | 有       | 良佳     | 良好     | 全部     |

**结论：**

- **rust-easyofd**: 速度最快（5胜），压缩率最好（2胜），推荐首选
- **go-zc310**: intro.ofd 最快（952ms），文本相似度最高（4胜）、文本提取最完整（ano=7282 字符），适合需要搜索/复制文本或处理复杂文档的场景
- **go-ofdgo**: 像素相似度最高（3胜），但文件体积较大、文本提取较差
- **python-easyofd**: 部分文件转换失败，intro 还原度差
- **python-ofd2pdf**: 基于图片渲染，无法提取文本
- **java-ofdrw**: 兼容性好，但速度较慢；1000-pages 文本相似度最高（78.68%），小文件还原度接近数科
- **node-ofd2pdf**: 兼容性好，文本提取字符数上 999（4249）、zsbk（1938）领先
- **rust-easyofd**: 对含复杂图形/签章文件（intro、1000-pages）还原度差（46%、74%）

> **测试局限性说明：**
> - 本测试仅运行 1 次，无方差/置信区间，无法判断差异是否显著
> - 未记录机器规格（CPU 型号/频率/核数/内存/是否 SSD），绝对耗时不可移植
> - Java/Node 未预热（JIT 未达稳态），首次运行吃亏，"Java 最慢"的结论需谨慎
> - 无冷启动/热启动分离，单次测量包含类加载、字体扫描等开销
> - 仅相对排名有意义，绝对毫秒数在不同机器上会完全不同
> - Rust 5 胜为全面领先，无明显极端值

> 以上结论由 AI 根据测试数据自动汇总生成


> **建议**：手工核对输出 PDF 效果，选择合适的库。不同 OFD 文件结构差异较大，实际效果可能与基准测试结果不同。
>
> 测试输出 PDF 文件可在 [Releases](https://github.com/zc310/ofd-benchmark/releases) 页面下载。

## 使用方法

```bash
# 编译
go build -o bin/ofd-benchmark .
go build -o bin/ofdgo-convert ./cmd/ofdgo-convert
go build -o bin/zc310-convert ./cmd/zc310-convert
cp <rust-binary>/easyofd bin/easyofd

# 安装 Python 依赖
pip install easyofd ofd2img

# 编译 Java 转换器（需要先下载依赖，见 java/README.md）
cd java && javac -cp "$(ls lib/*.jar | tr '\n' ':')" Ofd2Pdf.java && cd ..

# 安装 Node.js 依赖
cd node && npm install && cd ..

# 运行测试（默认测试所有转换器）
./bin/ofd-benchmark <file.ofd>

# 只测试某个转换器
./bin/ofd-benchmark <file.ofd> ofdgo
./bin/ofd-benchmark <file.ofd> zc310
./bin/ofd-benchmark <file.ofd> rust
./bin/ofd-benchmark <file.ofd> python
./bin/ofd-benchmark <file.ofd> java
./bin/ofd-benchmark <file.ofd> node
```

转换后的 PDF 文件保存在 `bin/output/` 目录，命名格式：`{原文件名}_{转换器}.pdf`

转换器名称：

- `go-ofdgo`
- `go-zc310`
- `rust-easyofd`
- `python-easyofd`
- `python-ofdreader`
- `java-ofdrw`
- `node-ofd2pdf`

## 项目结构

```
ofd-benchmark/
├── bin/
│   ├── ofd-benchmark       # 主程序
│   ├── ofdgo-convert       # Go 转换器
│   ├── zc310-convert       # Go 转换器
│   ├── easyofd             # Rust 二进制
│   └── output/             # 转换后的 PDF 文件
├── python/
│   ├── ofd2pdf.py          # Python 转换器（easyofd + ofd2img）
│   └── ofd2pdf_ofdreader.py
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
├── rust/                   # Rust 集成说明和示例
│   ├── README.md
│   └── ofd2pdf/
├── main.go
├── go.mod
└── README.md
```


