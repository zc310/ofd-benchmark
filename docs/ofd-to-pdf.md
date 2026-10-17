# OFD → PDF 对比测试

> [OFD 库汇总](ofd-libraries.md)：各 OFD 库能力 / 许可证对比表
>
> [PDF → OFD 对比测试](pdf-to-ofd.md)：反方向 PDF→OFD 转换对比

## 测试库

### Go 库

| 库                                                    | 版本                  | OFD→PDF |  更新时间  | 备注   |
|-------------------------------------------------------|-----------------------|:-------:|:----------:|--------|
| [xiaoqidun-ofdgo](https://github.com/xiaoqidun/ofdgo) | v0.0.0-20261017123659 |   ✅    | 2026-10-17 | 已测试 |
| [zc310-ofd](https://github.com/zc310/ofd)             | v0.1.5                |   ✅    | 2026-10-17 | 已测试 |

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
| [@miconvert/ofd-to-pdf](https://www.npmjs.com/package/@miconvert/ofd-to-pdf) | 0.2.3 |   ✅    | 2026-08-03 | 全部成功 |

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
| 1000-pages.ofd | 508K | 1000页文件   |
| 999.ofd        | 30K  | 带签章文件   |
| zsbk.ofd       | 1.5M | 数科签章文件 |

> 同目录下 `*.pdf` 为各 OFD 的参考转换结果，由**数科版式阅读器**导出，用作像素级对比基准。

## 测试结果

### 转换速度

| 转换器                                                                       | hello.ofd |  ano.ofd | intro.ofd | 1000-pages.ofd |   999.ofd | zsbk.ofd | 胜出次数 |
|------------------------------------------------------------------------------|----------:|---------:|----------:|---------------:|----------:|---------:|---------:|
| [go-zc310](https://github.com/zc310/ofd)                                     | **120ms** | **202ms** | **900ms** |      **2.11s** | **267ms** |    454ms |        5 |
| [go-ofdgo](https://github.com/xiaoqidun/ofdgo)                               |     179ms |     945ms |    17.67s |         34.92s |     591ms |    6.94s |        0 |
| [easyofd](https://pypi.org/project/easyofd/)                                 |    FAILED |    1.09s |     8.43s |         FAILED |     1.01s |    5.18s |        0 |
| [ofd2pdf](https://github.com/jsyzdej/ofd2pdf)                                |     153ms |   FAILED |    FAILED |         43.82s |     409ms | **201ms** |        1 |
| [ofdrw](https://github.com/ofdrw/ofdrw)                                      |     597ms |    759ms |     5.11s |          2.63s |     799ms |    5.53s |        0 |
| [@miconvert/ofd-to-pdf](https://www.npmjs.com/package/@miconvert/ofd-to-pdf) |     261ms |    711ms |     1.21s |          9.56s |     507ms |    801ms |        0 |

> - `python-easyofd`: `pip install easyofd ofd2img`，自动切换库，部分文件失败
> - `python-ofd2pdf`: `pip install ofd2pdf`，基于图片渲染，部分文件失败
> - `java-ofdrw`: `org.ofdrw:ofdrw-converter:2.4.0`，全部成功
> - `node-ofd2pdf`: `npm install @miconvert/ofd-to-pdf`，全部成功

### PDF 文件大小

| 转换器                                                                       | hello.ofd |    ano.ofd | intro.ofd | 1000-pages.ofd |    999.ofd |   zsbk.ofd | 胜出次数 |
|------------------------------------------------------------------------------|----------:|-----------:|----------:|---------------:|-----------:|-----------:|---------:|
| [go-zc310](https://github.com/zc310/ofd)                                     |     9.9KB |     77.7KB |     7.4MB |      **2.6MB** | **65.7KB** |      1.4MB |        2 |
| [go-ofdgo](https://github.com/xiaoqidun/ofdgo)                               |    15.5KB |    859.4KB |    60.8MB |         66.2MB |     132.8KB |      8.6MB |        0 |
| [easyofd](https://pypi.org/project/easyofd/)                                 |    FAILED | **36.1KB** |    38.2MB |         FAILED |     73.1KB |     13.6MB |        1 |
| [ofd2pdf](https://github.com/jsyzdej/ofd2pdf)                                |    45.5KB |     FAILED |    FAILED |        114.9MB |      752KB | **129KB** |        1 |
| [python-ofdreader](https://pypi.org/project/ofdreader/)                      |    FAILED |     FAILED |    FAILED |         FAILED |     FAILED |     FAILED |        0 |
| [ofdrw](https://github.com/ofdrw/ofdrw)                                      |     4.9KB |     65.6KB |    22.6MB |          3.5MB |     76.9KB |     15.4MB |        0 |
| [@miconvert/ofd-to-pdf](https://www.npmjs.com/package/@miconvert/ofd-to-pdf) |  **1.6KB** |      197KB | **1.1MB** |          9.1MB |     82.9KB |      559KB |        2 |

### 文本提取能力

| 转换器                                                                       | hello.ofd |  ano.ofd | intro.ofd |  999.ofd | zsbk.ofd | 胜出次数 |
|------------------------------------------------------------------------------|----------:|---------:|----------:|---------:|---------:|---------:|
| [go-zc310](https://github.com/zc310/ofd)                                     |        23 | **7474** |      5836 |     4201 |     1051 |        1 |
| [go-ofdgo](https://github.com/xiaoqidun/ofdgo)                               |         8 |      680 |       613 |     4186 |      509 |        0 |
| [easyofd](https://pypi.org/project/easyofd/)                                 |    FAILED |     1208 | **10194** |     3933 |     1508 |        1 |
| [ofd2pdf](https://github.com/jsyzdej/ofd2pdf)                                |         0 |   FAILED |    FAILED |        0 |        0 |        0 |
| [ofdrw](https://github.com/ofdrw/ofdrw)                                      |        22 |     6456 |      1036 |     2933 |      181 |        0 |
| [@miconvert/ofd-to-pdf](https://www.npmjs.com/package/@miconvert/ofd-to-pdf) |    **31** |     6990 |      9614 | **4249** | **1938** |        3 |

> - `python-ofd2pdf` 基于图片渲染，无法提取文本
> - 单位为字符数（chars）
> - 2026-10-17 重测环境无 PyMuPDF（fitz），`go-zc310` / `go-ofdgo` 行的字符数与像素/相似度改用 `mutool` 1.27（同为 MuPDF 内核，96 DPI）提取，与旧 fitz 口径相差约 ±3%；其他转换器行为旧 fitz 数值

### 与数科版式输出对比

以数科版式阅读器导出的 `test/testdata/*.pdf` 为基准，对比转换效果。

**像素相似度**（96 DPI，1 - 平均归一化色差，100% 为完全相同）：

| 转换器                                                                       |  hello.ofd |    ano.ofd |  intro.ofd | 1000-pages.ofd |    999.ofd |   zsbk.ofd | 胜出次数 |
|------------------------------------------------------------------------------|-----------:|-----------:|-----------:|---------------:|-----------:|-----------:|---------:|
| [go-zc310](https://github.com/zc310/ofd)                                     |     99.62% | **99.67%** |     99.43% |     **98.05%** |     98.72% |    **99.59%** |        3 |
| [go-ofdgo](https://github.com/xiaoqidun/ofdgo)                               |  **99.64%** |     99.44% | **99.73%** |          97.99% | **98.83%** |     98.99% |        3 |
| [easyofd](https://pypi.org/project/easyofd/)                                 |    FAILED |     93.82% |     37.74% |         FAILED |     92.41% |     94.41% |        0 |
| [ofd2pdf](https://github.com/jsyzdej/ofd2pdf)                                |     99.61% |     FAILED |     FAILED |         97.94% |     97.74% |     94.33% |        0 |
| [ofdrw](https://github.com/ofdrw/ofdrw)                                      |     99.40% |     98.86% |     98.94% |         97.65% |     97.73% |     98.45% |        0 |
| [@miconvert/ofd-to-pdf](https://www.npmjs.com/package/@miconvert/ofd-to-pdf) |     99.01% |     97.74% |     96.29% |         97.04% |     95.80% |     94.40% |        0 |

**文件大小对比**（转换后 PDF 大小 / 数科参考 PDF 大小，越小表示压缩率越高，`1.00` 表示与数科输出一致；**加粗**为最接近 `1.00` 的转换器）：

| 转换器                                                                       | hello.ofd |  ano.ofd | intro.ofd | 1000-pages.ofd |  999.ofd | zsbk.ofd |
|------------------------------------------------------------------------------|----------:|---------:|----------:|---------------:|---------:|---------:|
| 数科版式阅读器                                                               |      1.00 |     1.00 |      1.00 |           1.00 |     1.00 |     1.00 |
| [go-zc310](https://github.com/zc310/ofd)                                     |      2.81 |     0.09 |  **1.10** |     **1.08** |     0.82 | **0.91** |
| [go-ofdgo](https://github.com/xiaoqidun/ofdgo)                               |      4.40 | **0.95** |      9.06 |         27.48 |     1.66 |     5.45 |
| [easyofd](https://pypi.org/project/easyofd/)                                 |    FAILED |     0.04 |      5.70 |         FAILED |     0.92 |     8.62 |
| [ofd2pdf](https://github.com/jsyzdej/ofd2pdf)                                |     12.93 |   FAILED |    FAILED |          47.71 |     9.40 |     0.08 |
| [ofdrw](https://github.com/ofdrw/ofdrw)                                      |  **1.40** |     0.07 |      3.37 |           1.45 | **0.96** |     9.78 |
| [@miconvert/ofd-to-pdf](https://www.npmjs.com/package/@miconvert/ofd-to-pdf) |      0.47 |     0.22 |      0.16 |           3.79 | **1.04** |     0.35 |

> - 像素相似度越高表示输出越贴近数科版式阅读器；仿色/抗锯齿等平滑差异会小幅降低相似度
> - intro 的像素对比跳过 14～19 页（数科导出的参考 PDF 这几页背景图丢失，不计入对比）
> - `python-easyofd` 在 hello、1000-pages 转换失败；`python-ofd2pdf` 在 ano、intro 转换失败
> - 复现命令：`python3 tools/compare_pdf.py 96 <file>`

**文本相似度**（与数科参考 PDF 提取文本的 4-gram Jaccard 相似度，100% 为文本完全一致；`0.00%` 表示无文本或提取失败）：

参考文本长度（fitz 旧口径）：`hello=22` · `ano=7139` · `intro=5546` · `1000-pages=263999` · `999=3934` · `zsbk=1150`（2026-10-17 mutool 重测口径：`hello=23` · `ano=7285` · `intro=5827` · `1000-pages=283996` · `999=4138` · `zsbk=1040`）

| 转换器                                                                       |  hello.ofd |    ano.ofd |  intro.ofd | 1000-pages.ofd |    999.ofd |   zsbk.ofd | 胜出次数 |
|------------------------------------------------------------------------------|-----------:|-----------:|-----------:|---------------:|-----------:|-----------:|---------:|
| [go-zc310](https://github.com/zc310/ofd)                                     | **100.00%** | **71.33%** | **98.50%** |     **98.59%** |     94.72% | **98.85%** |        5 |
| [go-ofdgo](https://github.com/xiaoqidun/ofdgo)                               |     26.32% |      5.28% |      8.57% |           4.01% | **95.01%** |     26.70% |        1 |
| [easyofd](https://pypi.org/project/easyofd/)                                 |    FAILED |     12.62% |      0.08% |         FAILED |     80.54% |     11.20% |        0 |
| [ofd2pdf](https://github.com/jsyzdej/ofd2pdf)                                |      0.00% |     FAILED |    FAILED |          0.00% |      0.00% |      0.00% |        0 |
| [ofdrw](https://github.com/ofdrw/ofdrw)                                      |     16.13% |     16.20% |      7.16% |         80.96% |     51.29% |      5.88% |        0 |
| [@miconvert/ofd-to-pdf](https://www.npmjs.com/package/@miconvert/ofd-to-pdf) |     10.81% |     13.57% |      5.88% |         80.66% |     41.24% |      4.86% |        0 |

> - go-zc310 文本相似度 5 文件领先（hello=100.00%、intro=98.50%、1000-pages=98.59%、zsbk=98.85%、ano=71.33%）；999 上 go-ofdgo 以 95.01% 略胜（94.72%）
> - 文本相似度对提取顺序/换行敏感，得分偏低不代表内容缺失；2026-10-17 重测的 go-zc310/go-ofdgo 行为 mutool 口径，与旧 fitz 数值不可直接对比（如 ano 旧 99.07% 系 fitz 提取顺序差异）
> - `python-ofd2pdf` 基于图片渲染，全部无文本；`go-ofdgo` 文本提取较上版明显改善（999=4186/4138 字符）但多数文件仍弱（hello=8/22、ano=680 字符）
> - 复现命令：`python3 tools/compare_text.py <file>`

### 质量总结

| 转换器                                                              | 转换成功率 | 页面尺寸 | 矢量图形 | 文本可提取 | 图片嵌入 | 文件压缩 | 中文支持 | 兼容性   |
|---------------------------------------------------------------------|:----------:|----------|----------|------------|----------|----------|----------|----------|
| [go-zc310](https://github.com/zc310/ofd)                            |    100%    | A4       | 支持     | 最佳       | 有       | 良佳     | 最佳     | 全部     |
| [go-ofdgo](https://github.com/xiaoqidun/ofdgo)                      |    100%    | A4       | 支持     | 部分       | 有       | 较差     | 良好     | 全部     |
| [python-easyofd](https://pypi.org/project/easyofd/)                 |    67%     | A4       | 支持     | 部分       | 有       | 良佳     | 良好     | 部分失败 |
| [python-ofd2pdf](https://github.com/jsyzdej/ofd2pdf)                |    67%     | A4       | 不支持   | 无         | 有       | 一般     | 良好     | 部分失败 |
| [java-ofdrw](https://github.com/ofdrw/ofdrw)                        |    100%    | A4       | 支持     | 良好       | 有       | 良佳     | 良好     | 全部     |
| [node-ofd2pdf](https://www.npmjs.com/package/@miconvert/ofd-to-pdf) |    100%    | A4       | 支持     | 良好       | 有       | 良佳     | 良好     | 全部     |

> **本轮重测关键变化（vs 上轮）：**
> - go-ofdgo 大文件速度明显退步：intro 3.12s→17.67s、1000-pages 9.92s→34.92s、zsbk 527ms→6.94s；hello 失去速度冠军（106ms→179ms，被 go-zc310 120ms 反超）
> - go-ofdgo 体积分化：intro 34.8MB→60.8MB、1000-pages 50.6MB→66.2MB、zsbk 2.5MB→8.6MB 膨胀，ano 1.4MB→859.4KB、999 2.1MB→132.8KB 缩小
> - go-ofdgo 文本提取从几乎为零提升到部分可用：999 61→4186 字符（相似度 95.01%，该列首次夺冠），多数文件仍弱
> - go-zc310 像素 ano 99.62%→99.67%、zsbk 99.50%→99.59%（像素胜出 2→3），hello 以 120ms 夺回速度冠军
> - 文本提取/相似度列因 fitz→mutool 口径变更整体位移（如 ano 相似度 99.07%→71.33% 系提取顺序差异，非内容缺失），见上文脚注

**结论：**

- **go-zc310**: 速度 5 胜（hello 120ms、ano 202ms、intro 900ms、1000-pages 2.11s、999 267ms）、体积 2 胜（1000-pages 2.6MB、999 65.7KB）、文本相似度 4 胜、像素还原 3 胜（ano 99.67%、1000-pages 98.05%、zsbk 99.59%），文本提取最完整（ano=7474 字符），适合需要搜索/复制文本或处理复杂文档的场景
- **go-ofdgo**: 像素相似度 3 胜（hello 99.64%、intro 99.73%、999 98.83%）、999 文本相似度最高（95.01%），但文件体积显著增大（intro 60.8MB、1000-pages 66.2MB）、大文件速度退步（intro 17.67s、1000-pages 34.92s、zsbk 6.94s），文本提取多数文件仍弱（hello=8、ano=680 字符）
- **python-easyofd**: 部分文件转换失败，intro 还原度差
- **python-ofd2pdf**: 基于图片渲染，无法提取文本；zsbk 速度（201ms）与体积（129KB）最优
- **java-ofdrw**: 兼容性好，但速度较慢；1000-pages 文本相似度 80.96% 次于 go-zc310，小文件还原度接近数科
- **node-ofd2pdf**: 兼容性好，文件体积 2 胜（hello 1.6KB、intro 1.1MB），文本提取字符数上 hello（31）、999（4249）、zsbk（1938）领先

> **测试局限性说明：**
> - 本测试仅运行 1 次，无方差/置信区间，无法判断差异是否显著
> - 未记录机器规格（CPU 型号/频率/核数/内存/是否 SSD），绝对耗时不可移植
> - Java/Node 未预热（JIT 未达稳态），首次运行吃亏，"Java 最慢"的结论需谨慎
> - 无冷启动/热启动分离，单次测量包含类加载、字体扫描等开销
> - 仅相对排名有意义，绝对毫秒数在不同机器上会完全不同

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
./bin/ofd-benchmark <file.ofd> python
./bin/ofd-benchmark <file.ofd> java
./bin/ofd-benchmark <file.ofd> node
```

转换后的 PDF 文件保存在 `bin/output/` 目录，命名格式：`{原文件名}_{转换器}.pdf`

转换器名称：

- `go-ofdgo`
- `go-zc310`
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
├── main.go
├── go.mod
└── README.md
```


