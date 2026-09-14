# OFD to PDF 速度对比测试

## 测试库

### Go 库

| 库                                                               | 版本                  | 支持 OFD→PDF | 更新时间   | 备注   |
|------------------------------------------------------------------|-----------------------|--------------|------------|--------|
| [github.com/xiaoqidun/ofdgo](https://github.com/xiaoqidun/ofdgo) | v0.0.0-20260912163201 | ✅           | 2026-09-14 | 已测试 |
| [github.com/zc310/ofd](https://github.com/zc310/ofd)             | v0.1.1-20260914000557 | ✅           | 2026-09-14 | 已测试 |

### Rust 库

| 库                                                                                 | 版本   | 支持 OFD→PDF | 更新时间   | 备注                       |
|------------------------------------------------------------------------------------|--------|--------------|------------|----------------------------|
| [github.com/easy-4-rust/easyofd-rust](https://github.com/easy-4-rust/easyofd-rust) | v0.1.2 | ✅           | 2026-09-10 | 已测试                     |
| [github.com/feuvan/ofdmanager](https://github.com/feuvan/ofdmanager)               | v0.2.0 | ❌           | 2026-08-11 | 仅支持 OFD→PNG             |
| [github.com/geniusnut/rs_ofd](https://github.com/geniusnut/rs_ofd)                 | -      | ❌           | 2024-11-28 | 仅支持 OFD→PNG             |
| [github.com/ofd-utility/ofd-utility](https://github.com/ofd-utility/ofd-utility)   | -      | ❌           | 2026-08-07 | OFD 解析、校验、渲染到图片 |

### Python 库

| 库                                               | 版本     | 支持 OFD→PDF | 更新时间   | 备注                       |
|--------------------------------------------------|----------|--------------|------------|----------------------------|
| [easyofd](https://pypi.org/project/easyofd/)     | 20260427 | ⚠️           | 2026-04-27 | 部分文件失败               |
| [ofd2img](https://pypi.org/project/ofd2img/)     | 0.1.2    | ⚠️           | 2026-05-07 | 部分文件失败               |
| [ofd2pdf](https://github.com/jsyzdej/ofd2pdf)    | 0.0.2    | ⚠️           | 2026-07-08 | 部分文件失败，基于图片渲染 |
| [ofdreader](https://pypi.org/project/ofdreader/) | 0.1.0    | ❌           | 2026-06-01 | 空包，无法使用             |
| [ofdparser](https://pypi.org/project/ofdparser/) | 0.0.8    | ❌           | 2023-04-11 | 缺 Courier 字体，无法导入  |

### Java 库

| 库                                                   | 版本  | 支持 OFD→PDF | 更新时间   | 备注                         |
|------------------------------------------------------|-------|--------------|------------|------------------------------|
| [ofdrw](https://github.com/ofdrw/ofdrw)              | 2.4.0 | ✅           | 2026-08-04 | 需手动下载依赖，全部成功     |
| [ofdbox](https://gitee.com/bookhhu/ofdbox)           | -     | ⚠️           | 2025-06-08 | OFD→图片可用，PDF 正在开发中 |
| [ofd-analyze](https://github.com/cooker/ofd-analyze) | -     | ❌           | 2022-05-20 | OFD 解析器，支持 OFD→图片    |

### .NET 库

| 库                                                                                 | 版本     | 支持 OFD→PDF | 更新时间   | 备注                               |
|------------------------------------------------------------------------------------|----------|--------------|------------|------------------------------------|
| [github.com/zhuovi/XiaoFeng.Ofd](https://github.com/zhuovi/XiaoFeng.Ofd)           | 1.0.0    | ❌           | 2025-06-03 | README 声称支持但代码未实现        |
| [github.com/wangyi160/ofdparser](https://github.com/wangyi160/ofdparser)           | -        | ❌           | 2022-02-11 | OFD 解析器                         |

### JavaScript 库

| 库                                                                           | 版本  | 支持 OFD→PDF | 更新时间   | 备注                  |
|------------------------------------------------------------------------------|-------|--------------|------------|-----------------------|
| [@miconvert/ofd-to-pdf](https://www.npmjs.com/package/@miconvert/ofd-to-pdf) | 0.2.0 | ✅           | 2026-08-03 | 全部成功              |
| [github.com/isee15/ofdjs](https://github.com/isee15/ofdjs)                   | -     | ❌           | 2026-07-14 | OFD 解析渲染到 Canvas |

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

## 测试结果

### 转换速度

| 转换器           | hello.ofd |  ano.ofd | intro.ofd | 1000-pages.ofd |   999.ofd | zsbk.ofd | 胜出次数 |
|------------------|----------:|---------:|----------:|---------------:|----------:|---------:|---------:|
| go-ofdgo         |      86ms |    213ms |     2.40s |          8.25s |     308ms |    324ms |        0 |
| go-zc310         |     129ms |    341ms |     1.26s |          8.70s |     363ms |    429ms |        0 |
| rust-easyofd     |  **42ms** | **72ms** |     1.82s |      **0.97s** | **143ms** | **70ms** |        5 |
| python-easyofd   |    FAILED |   1084ms |     8.41s |         FAILED |    1020ms |   5122ms |        0 |
| python-ofd2pdf   |     137ms |   FAILED |    FAILED |          40.8s |     401ms |    198ms |        0 |
| python-ofdreader |    FAILED |   FAILED |    FAILED |         FAILED |    FAILED |   FAILED |        0 |
| java-ofdrw       |     497ms |    737ms |     5.13s |          2.13s |     781ms |   5553ms |        0 |
| node-ofd2pdf     |     240ms |    708ms | **1.18s** |          4.66s |     503ms |    795ms |        1 |

> - `python-easyofd`: `pip install easyofd ofd2img`，自动切换库，部分文件失败
> - `python-ofd2pdf`: `pip install ofd2pdf`，基于图片渲染，部分文件失败
> - `python-ofdreader`: `pip install ofdreader`，空包（无 Python 代码），全部失败
> - `java-ofdrw`: `org.ofdrw:ofdrw-converter:2.4.0`，全部成功
> - `node-ofd2pdf`: `npm install @miconvert/ofd-to-pdf`，全部成功

### PDF 文件大小

| 转换器           | hello.ofd | ano.ofd | intro.ofd | 1000-pages.ofd | 999.ofd |   zsbk.ofd | 胜出次数 |
|------------------|----------:|--------:|----------:|---------------:|--------:|-----------:|---------:|
| go-ofdgo         |    14.1KB |   1.1MB |      35MB |           66MB |   1.2MB |      1.9MB |        0 |
| go-zc310         |     9.9KB |  96.1KB |      14MB |      **1.5MB** |  88.7KB |      1.4MB |        1 |
| rust-easyofd     |    37.3KB |  77.1KB |      29MB |      **614KB** |  83.3KB | **77.8KB** |        2 |
| python-easyofd   |    FAILED |  36.1KB |      38MB |         FAILED |  73.1KB |     13.6MB |        1 |
| python-ofd2pdf   |    36.8KB |  FAILED |    FAILED |         76.1MB | 751.7KB |      129KB |        0 |
| python-ofdreader |    FAILED |  FAILED |    FAILED |         FAILED |  FAILED |     FAILED |        0 |
| java-ofdrw       |     4.8KB |  65.6KB |      23MB |          2.0MB |  76.9KB |     15.4MB |        1 |
| node-ofd2pdf     |     1.3KB | 196.8KB | **1.1MB** |          4.2MB |  82.9KB |      559KB |        2 |

### 文本提取能力

| 转换器           | hello.ofd |  ano.ofd |  999.ofd | zsbk.ofd | 胜出次数 |
|------------------|----------:|---------:|---------:|---------:|---------:|
| go-ofdgo         |         8 |        0 |       76 |       19 |        0 |
| go-zc310         |        22 | **7540** |     4595 |      985 |        1 |
| rust-easyofd     |        23 |     6357 |     3869 |      369 |        0 |
| python-easyofd   |         0 |     1469 |     4579 |     2128 |        1 |
| python-ofd2pdf   |         0 |   FAILED |   FAILED |        0 |        0 |
| python-ofdreader |    FAILED |   FAILED |   FAILED |   FAILED |        0 |
| java-ofdrw       |        20 |     6307 |     3581 |      194 |        0 |
| node-ofd2pdf     |        25 |     7049 | **4905** |     1697 |        1 |

> - `python-ofd2pdf` 基于图片渲染，无法提取文本
> - 单位为字符数（chars）

### 质量总结

| 转换器         | 页面尺寸 | 矢量图形 | 文本可提取 | 图片嵌入 | 文件压缩 | 中文支持 | 兼容性   |
|----------------|----------|----------|------------|----------|----------|----------|----------|
| go-ofdgo       | A4       | 支持     | 部分       | 有       | 较差     | 良好     | 全部     |
| go-zc310       | A4       | 支持     | 最佳       | 有       | 良佳     | 最佳     | 全部     |
| rust-easyofd   | A4       | 支持     | 良好       | 有       | 最佳     | 良好     | 全部     |
| python-easyofd | A4       | 支持     | 部分       | 有       | 良佳     | 良好     | 部分失败 |
| python-ofd2pdf | A4       | 不支持   | 无         | 有       | 一般     | 良好     | 部分失败 |
| java-ofdrw     | A4       | 支持     | 良好       | 有       | 良佳     | 良好     | 全部     |
| node-ofd2pdf   | A4       | 支持     | 良好       | 有       | 良佳     | 良好     | 全部     |

**结论：**
- **rust-easyofd**: 速度最快（5胜），压缩率最好（2胜），推荐首选
- **go-zc310**: 文本提取最完整（ano.ofd 7540字符），适合需要搜索/复制文本的场景
- **node-ofd2pdf**: intro.ofd 最快（1.18s），兼容性好
- **go-ofdgo**: 文件体积较大，文本提取较差
- **python-easyofd**: 部分文件转换失败
- **python-ofd2pdf**: 基于图片渲染，无法提取文本
- **java-ofdrw**: 兼容性好，但速度较慢

> **建议**：手工核对输出 PDF 效果，选择合适的库。不同 OFD 文件结构差异较大，实际效果可能与基准测试结果不同。
> 
> 测试输出 PDF 文件可在 [Releases](https://github.com/zc310/ofd-benchmark/releases) 页面下载。

> 以上结论由 AI 根据测试数据自动汇总生成

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
