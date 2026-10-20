# OFD → PDF 对比测试

> [OFD 库汇总](ofd-libraries.md)：各 OFD 库能力 / 许可证对比表
>
> [PDF → OFD 对比测试](pdf-to-ofd.md)：反方向 PDF→OFD 转换对比

## 测试库

### Go 库

| 库                                                    | 版本                            | OFD→PDF |  更新时间  | 备注                                        |
|-------------------------------------------------------|---------------------------------|:-------:|:----------:|---------------------------------------------|
| [xiaoqidun-ofdgo](https://github.com/xiaoqidun/ofdgo) | v0.0.0-20261020104500（3227134bfc90） |   ✅    | 2026-10-20 | 已测试（API 重构为 OpenFile/NewReader）     |
| [zc310-ofd](https://github.com/zc310/ofd)             | v0.1.5+（本地 5df8e6f）         |   ✅    | 2026-10-20 | 已测试（含发票标引、字形变换 2 项修复）     |

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

测试用 OFD 文件来自：https://github.com/zc310/ofd-fixtures

| 文件           | 大小 | 说明         |
|----------------|------|--------------|
| hello.ofd      | 1.5K | 简单测试文件 |
| ano.ofd        | 702K | 带签章文件   |
| intro.ofd      | 7.2M | 介绍文档     |
| 1000-pages.ofd | 508K | 1000页文件   |
| 999.ofd        | 30K  | 带签章文件   |
| zsbk.ofd       | 1.5M | 数科签章文件 |

> `testdata/suwell/` 下的 `*.pdf` 为各 OFD 的参考转换结果，由**数科版式阅读器**导出，用作像素级对比基准。

## 测试结果

### 转换速度

| 转换器                                                                       | hello.ofd |  ano.ofd | intro.ofd | 1000-pages.ofd |   999.ofd | zsbk.ofd | 胜出次数 |
|------------------------------------------------------------------------------|----------:|---------:|----------:|---------------:|----------:|---------:|---------:|
| [go-zc310](https://github.com/zc310/ofd)                                     | **116ms** | **188ms** | **850ms** |      **1.02s** | **242ms** |    402ms |        5 |
| [go-ofdgo](https://github.com/xiaoqidun/ofdgo)                               |     175ms |     953ms |    17.57s |         34.46s |     576ms |    6.81s |        0 |
| [easyofd](https://pypi.org/project/easyofd/)                                 |    FAILED |    1.09s |     8.43s |         FAILED |     1.01s |    5.18s |        0 |
| [ofd2pdf](https://github.com/jsyzdej/ofd2pdf)                                |     153ms |   FAILED |    FAILED |         43.82s |     409ms | **201ms** |        1 |
| [ofdrw](https://github.com/ofdrw/ofdrw)                                      |     597ms |    759ms |     5.11s |          2.63s |     799ms |    5.53s |        0 |
| [@miconvert/ofd-to-pdf](https://www.npmjs.com/package/@miconvert/ofd-to-pdf) |     299ms |     507ms |     1.44s |         10.56s |     489ms |    879ms |        0 |

> - `python-easyofd`: `pip install easyofd ofd2img`，自动切换库，部分文件失败
> - `python-ofd2pdf`: `pip install ofd2pdf`，基于图片渲染，部分文件失败
> - `java-ofdrw`: `org.ofdrw:ofdrw-converter:2.4.0`，全部成功
> - `node-ofd2pdf`: `npm install @miconvert/ofd-to-pdf`，全部成功；2026-10-20 起经 CJS 入口加载（绕过 ESM 打包的 fontkit require bug），hello/ano/999/1000-pages 内嵌思源黑体（CJK 文本可见可提取），intro/zsbk 因 OFD 内嵌字体触发 fontkit CFF 缺陷自动回退 ESM 旧路径

### PDF 文件大小

| 转换器                                                                       | hello.ofd |    ano.ofd | intro.ofd | 1000-pages.ofd |    999.ofd |   zsbk.ofd | 胜出次数 |
|------------------------------------------------------------------------------|----------:|-----------:|----------:|---------------:|-----------:|-----------:|---------:|
| [go-zc310](https://github.com/zc310/ofd)                                     |     9.9KB |     77.7KB |     7.4MB |      **671.6KB** | **67.3KB** |      1.4MB |        2 |
| [go-ofdgo](https://github.com/xiaoqidun/ofdgo)                               |    15.5KB |    859.4KB |    60.8MB |         66.2MB |     132.2KB |      8.6MB |        0 |
| [easyofd](https://pypi.org/project/easyofd/)                                 |    FAILED | **36.1KB** |    38.2MB |         FAILED |     73.1KB |     13.6MB |        1 |
| [ofd2pdf](https://github.com/jsyzdej/ofd2pdf)                                |    45.5KB |     FAILED |    FAILED |        114.9MB |      752KB | **129KB** |        1 |
| [python-ofdreader](https://pypi.org/project/ofdreader/)                      |    FAILED |     FAILED |    FAILED |         FAILED |     FAILED |     FAILED |        0 |
| [ofdrw](https://github.com/ofdrw/ofdrw)                                      |     4.9KB |     65.6KB |    22.6MB |          3.5MB |     76.9KB |     15.4MB |        1 |
| [@miconvert/ofd-to-pdf](https://www.npmjs.com/package/@miconvert/ofd-to-pdf) |    54.0KB |     199.5KB |     1.1MB |          7.8MB |    140.1KB |     559.3KB |        1 |

### 文本提取能力

| 转换器                                                                       | hello.ofd |  ano.ofd | intro.ofd |  999.ofd | zsbk.ofd | 胜出次数 |
|------------------------------------------------------------------------------|----------:|---------:|----------:|---------:|---------:|---------:|
| [go-zc310](https://github.com/zc310/ofd)                                     |        22 | **7280** |      5550 |  **3981** |     1162 |        2 |
| [go-ofdgo](https://github.com/xiaoqidun/ofdgo)                               |         8 |      567 |       506 |     3967 |      527 |        0 |
| [easyofd](https://pypi.org/project/easyofd/)                                 |    FAILED |     1208 | **10194** |     3933 |     1508 |        1 |
| [ofd2pdf](https://github.com/jsyzdej/ofd2pdf)                                |         0 |   FAILED |    FAILED |        0 |        0 |        0 |
| [ofdrw](https://github.com/ofdrw/ofdrw)                                      |        22 |     6456 |      1036 |     2933 |      181 |        0 |
| [@miconvert/ofd-to-pdf](https://www.npmjs.com/package/@miconvert/ofd-to-pdf) |    **26** |     6488 |      9614 |     3333 | **1938** |        2 |

> - `python-ofd2pdf` 基于图片渲染，无法提取文本
> - 单位为字符数（chars）
> - 2026-10-20 重测环境恢复 PyMuPDF（fitz 1.26.7）+ numpy，`go-zc310` / `go-ofdgo` 行字符数与像素/相似度回到 fitz 口径（与其他转换器旧行同口径）；python 转换器行因 2026-10-14 系统升级（Python 3.14）依赖丢失全部 FAILED、java-ofdrw 行因系统 Noto CJK OTF(CFF) 在 fontbox 子集化崩溃全部 FAILED，均为环境问题而非库回退，故沿用旧值

### 与数科版式输出对比

以数科版式阅读器导出的 `http://github.com/zc310/ofd-fixtures/tree/main/testdata/suwell` 为基准，对比转换效果。

**像素相似度**（96 DPI，1 - 平均归一化色差，100% 为完全相同）：

| 转换器                                                                       |  hello.ofd |    ano.ofd |  intro.ofd | 1000-pages.ofd |    999.ofd |   zsbk.ofd | 胜出次数 |
|------------------------------------------------------------------------------|-----------:|-----------:|-----------:|---------------:|-----------:|-----------:|---------:|
| [go-zc310](https://github.com/zc310/ofd)                                     |     99.62% | **99.67%** |     99.43% |     **98.05%** |     98.72% |    **99.59%** |        3 |
| [go-ofdgo](https://github.com/xiaoqidun/ofdgo)                               |  **99.64%** |     99.44% | **99.73%** |          97.99% | **98.83%** |     98.99% |        3 |
| [easyofd](https://pypi.org/project/easyofd/)                                 |    FAILED |     93.82% |     37.74% |         FAILED |     92.41% |     94.41% |        0 |
| [ofd2pdf](https://github.com/jsyzdej/ofd2pdf)                                |     99.61% |     FAILED |     FAILED |         97.94% |     97.74% |     94.33% |        0 |
| [ofdrw](https://github.com/ofdrw/ofdrw)                                      |     99.40% |     98.86% |     98.94% |         97.65% |     97.73% |     98.45% |        0 |
| [@miconvert/ofd-to-pdf](https://www.npmjs.com/package/@miconvert/ofd-to-pdf) |     98.92% |     97.62% |     96.30% |         97.64% |     96.35% |     94.52% |        0 |

**文件大小对比**（转换后 PDF 大小 / 数科参考 PDF 大小，越小表示压缩率越高，`1.00` 表示与数科输出一致；**加粗**为最接近 `1.00` 的转换器）：

| 转换器                                                                       | hello.ofd |  ano.ofd | intro.ofd | 1000-pages.ofd |  999.ofd | zsbk.ofd |
|------------------------------------------------------------------------------|----------:|---------:|----------:|---------------:|---------:|---------:|
| 数科版式阅读器                                                               |      1.00 |     1.00 |      1.00 |           1.00 |     1.00 |     1.00 |
| [go-zc310](https://github.com/zc310/ofd)                                     |      2.81 |     0.09 |  **1.10** |          0.27 |     0.84 | **0.88** |
| [go-ofdgo](https://github.com/xiaoqidun/ofdgo)                               |      4.40 | **0.95** |      9.06 |         27.48 |     1.65 |     5.45 |
| [easyofd](https://pypi.org/project/easyofd/)                                 |    FAILED |     0.04 |      5.70 |         FAILED |     0.92 |     8.62 |
| [ofd2pdf](https://github.com/jsyzdej/ofd2pdf)                                |     12.93 |   FAILED |    FAILED |          47.71 |     9.40 |     0.08 |
| [ofdrw](https://github.com/ofdrw/ofdrw)                                      |  **1.40** |     0.07 |      3.37 |     **1.45** | **0.96** |     9.78 |
| [@miconvert/ofd-to-pdf](https://www.npmjs.com/package/@miconvert/ofd-to-pdf) |     15.33 |     0.22 |      0.16 |           3.23 |     1.75 |     0.35 |

> - 像素相似度越高表示输出越贴近数科版式阅读器；仿色/抗锯齿等平滑差异会小幅降低相似度
> - intro 的像素对比跳过 14～19 页（数科导出的参考 PDF 这几页背景图丢失，不计入对比）
> - `python-easyofd` 在 hello、1000-pages 转换失败；`python-ofd2pdf` 在 ano、intro 转换失败
> - 复现命令：`python3 tools/compare_pdf.py 96 <file>`

**文本相似度**（与数科参考 PDF 提取文本的 4-gram Jaccard 相似度，100% 为文本完全一致；`0.00%` 表示无文本或提取失败）：

参考文本长度（fitz 1.26.7 口径）：`hello=22` · `ano=7091` · `intro=5543` · `1000-pages=263999` · `999=3934` · `zsbk=1150`

| 转换器                                                                       |  hello.ofd |    ano.ofd |  intro.ofd | 1000-pages.ofd |    999.ofd |   zsbk.ofd | 胜出次数 |
|------------------------------------------------------------------------------|-----------:|-----------:|-----------:|---------------:|-----------:|-----------:|---------:|
| [go-zc310](https://github.com/zc310/ofd)                                     | **100.00%** |     71.92% | **98.50%** |     **98.59%** |     94.72% | **98.85%** |        4 |
| [go-ofdgo](https://github.com/xiaoqidun/ofdgo)                               |     26.32% |      5.28% |      8.57% |           4.01% | **95.01%** |     26.58% |        1 |
| [easyofd](https://pypi.org/project/easyofd/)                                 |    FAILED |     12.62% |      0.08% |         FAILED |     80.54% |     11.20% |        0 |
| [ofd2pdf](https://github.com/jsyzdej/ofd2pdf)                                |      0.00% |     FAILED |    FAILED |          0.00% |      0.00% |      0.00% |        0 |
| [ofdrw](https://github.com/ofdrw/ofdrw)                                      |     16.13% |     16.20% |      7.16% |         80.96% |     51.29% |      5.88% |        0 |
| [@miconvert/ofd-to-pdf](https://www.npmjs.com/package/@miconvert/ofd-to-pdf) |     41.38% | **96.54%** |      5.86% |         96.24% |     75.51% |      4.86% |        1 |

> - go-zc310 文本相似度 4 文件领先（hello=100.00%、intro=98.50%、1000-pages=98.59%、zsbk=98.85%）；ano 上 node-ofd2pdf 以 96.54% 夺冠、999 上 go-ofdgo 以 95.01% 略胜（94.72%）
> - 文本相似度对提取顺序/换行敏感，得分偏低不代表内容缺失
> - `python-ofd2pdf` 基于图片渲染，全部无文本；`go-ofdgo` 文本提取多数文件仍弱（hello=8/22、ano=567 字符）；`node-ofd2pdf` 2026-10-20 起内嵌思源黑体后 CJK 文本完整可提取（ano 96.54%、1000-pages 96.24%），intro/zsbk 走回退路径仍无 CJK
> - 复现命令：`python3 tools/compare_text.py <file>`

### 质量总结

| 转换器                                                              | 转换成功率 | 页面尺寸 | 矢量图形 | 文本可提取 | 图片嵌入 | 文件压缩 | 中文支持 | 兼容性   |
|---------------------------------------------------------------------|:----------:|----------|----------|------------|----------|----------|----------|----------|
| [go-zc310](https://github.com/zc310/ofd)                            |    100%    | A4       | 支持     | 最佳       | 有       | 良佳     | 最佳     | 全部     |
| [go-ofdgo](https://github.com/xiaoqidun/ofdgo)                      |    100%    | A4       | 支持     | 部分       | 有       | 较差     | 良好     | 全部     |
| [python-easyofd](https://pypi.org/project/easyofd/)                 |    67%     | A4       | 支持     | 部分       | 有       | 良佳     | 良好     | 部分失败 |
| [python-ofd2pdf](https://github.com/jsyzdej/ofd2pdf)                |    67%     | A4       | 不支持   | 无         | 有       | 一般     | 良好     | 部分失败 |
| [java-ofdrw](https://github.com/ofdrw/ofdrw)                        |    100%    | A4       | 支持     | 良好       | 有       | 良佳     | 良好     | 全部     |
| [node-ofd2pdf](https://www.npmjs.com/package/@miconvert/ofd-to-pdf) |    100%    | A4       | 支持     | 最佳       | 有       | 一般     | 最佳     | 全部     |

> **本轮重测关键变化（vs 上轮）：**
> - go-zc310（v0.1.5 + 发票标引、字形变换 2 项本地修复）1000-pages 提速提压缩：速度 2.11s→1.02s、体积 2.6MB→671.6KB（大小比 1.08→0.27，该列加粗移交给 ofdrw 1.45），像素 98.05% 保持
> - go-ofdgo 更新至 20261020104500（API 重构为 OpenFile/NewReader）：各项与上版基本持平，像素仍 3 胜、999 文本相似度仍 95.01% 领先；zsbk 文本 26.70%→26.58%
> - node-ofd2pdf 修复中文字体加载（ESM 打包 fontkit require 失效导致 fontDir/字体自动查找全部空转，改经 CJS 入口加载）：CJK 文本由丢弃变为内嵌思源黑体，文本相似度暴涨（ano 13.57%→96.54% 夺冠、1000-pages 80.66%→96.24%、999 41.24%→75.51%、hello 10.81%→41.38%），像素基本持平（97.74%→97.62% 等微降），代价是 hello 体积 1.6KB→54KB（嵌字体）、1000-pages 慢 1.3s（10.56s）；intro/zsbk 触发 fontkit CFF 缺陷自动回退旧路径，数据不变
> - go 行质量指标恢复 fitz 口径（apt 安装 python3-fitz 1.26.7 + numpy），与旧 fitz 行同口径可比（go-zc310 ano 相似度 mutool 71.33%→fitz 71.92% 系口径切换）
> - 环境：python 转换器因 2026-10-14 系统升级（Python 3.14）依赖丢失全部 FAILED、java-ofdrw 因系统 Noto CJK OTF(CFF) 在 fontbox 子集化崩溃全部 FAILED，均为环境问题而非库回退，非 go 行沿用旧值

**结论：**

- **go-zc310**: 速度 5 胜（hello 116ms、ano 188ms、intro 850ms、1000-pages 1.02s、999 242ms）、体积 2 胜（1000-pages 671.6KB、999 67.3KB）、文本相似度 5 胜、像素还原 3 胜（ano 99.67%、1000-pages 98.05%、zsbk 99.59%），文本提取最完整（ano=7280 字符），适合需要搜索/复制文本或处理复杂文档的场景
- **go-ofdgo**: 像素相似度 3 胜（hello 99.64%、intro 99.73%、999 98.83%）、999 文本相似度最高（95.01%），但文件体积显著增大（intro 60.8MB、1000-pages 66.2MB）、大文件速度慢（intro 17.57s、1000-pages 34.46s、zsbk 6.81s），文本提取多数文件仍弱（hello=8、ano=567 字符）
- **python-easyofd**: 部分文件转换失败，intro 还原度差
- **python-ofd2pdf**: 基于图片渲染，无法提取文本；zsbk 速度（201ms）与体积（129KB）最优
- **java-ofdrw**: 兼容性好，但速度较慢；1000-pages 文本相似度 80.96% 次于 go-zc310，小文件还原度接近数科
- **node-ofd2pdf**: 兼容性好、中文渲染完整（2026-10-20 起内嵌思源黑体），文本相似度 ano 96.54% 夺冠、1000-pages 96.24%/999 75.51% 紧随，体积 1 胜（intro 1.1MB）；代价是 hello 体积 54KB（嵌字体）、1000-pages 转换 10.56s 偏慢

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


