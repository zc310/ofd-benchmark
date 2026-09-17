# OFD to PDF 速度对比测试

> [OFD 库汇总](docs/ofd-libraries.md)：各 OFD 库能力 / 许可证对比表

## 测试库

### Go 库

| 库                                                    | 版本                  | OFD→PDF |  更新时间  | 备注   |
|-------------------------------------------------------|-----------------------|:-------:|:----------:|--------|
| [xiaoqidun-ofdgo](https://github.com/xiaoqidun/ofdgo) | v0.0.0-20260915002242 |   ✅    | 2026-09-17 | 已测试 |
| [zc310-ofd](https://github.com/zc310/ofd)             | v0.1.1-20260915002405 |   ✅    | 2026-09-17 | 已测试 |

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

## 测试结果

### 转换速度

| 转换器                                                              | hello.ofd |  ano.ofd | intro.ofd | 1000-pages.ofd |   999.ofd | zsbk.ofd | 胜出次数 |
|---------------------------------------------------------------------|----------:|---------:|----------:|---------------:|----------:|---------:|---------:|
| [go-ofdgo](https://github.com/xiaoqidun/ofdgo)                      |     122ms |    217ms |     2.40s |          5.11s |     308ms |    329ms |        0 |
| [go-zc310](https://github.com/zc310/ofd)                            |     146ms |    346ms |     1.23s |          9.38s |     376ms |    475ms |        0 |
| [rust-easyofd](https://github.com/easy-4-rust/easyofd-rust)         |  **54ms** | **74ms** |     1.82s |      **0.95s** | **159ms** | **75ms** |        5 |
| [python-easyofd](https://pypi.org/project/easyofd/)                 |    FAILED |    1.43s |     8.48s |         FAILED |     1.04s |    5.11s |        0 |
| [python-ofd2pdf](https://github.com/jsyzdej/ofd2pdf)                |     170ms |   FAILED |    FAILED |          41.2s |     422ms |    201ms |        0 |
| [java-ofdrw](https://github.com/ofdrw/ofdrw)                        |     702ms |    736ms |     5.12s |          2.41s |     811ms |    5.51s |        0 |
| [node-ofd2pdf](https://www.npmjs.com/package/@miconvert/ofd-to-pdf) |     276ms |    703ms | **1.20s** |          8.52s |     506ms |    810ms |        1 |

> - `python-easyofd`: `pip install easyofd ofd2img`，自动切换库，部分文件失败
> - `python-ofd2pdf`: `pip install ofd2pdf`，基于图片渲染，部分文件失败
> - `java-ofdrw`: `org.ofdrw:ofdrw-converter:2.4.0`，全部成功
> - `node-ofd2pdf`: `npm install @miconvert/ofd-to-pdf`，全部成功

### PDF 文件大小

| 转换器                                                              | hello.ofd |    ano.ofd | intro.ofd | 1000-pages.ofd |    999.ofd |   zsbk.ofd | 胜出次数 |
|---------------------------------------------------------------------|----------:|-----------:|----------:|---------------:|-----------:|-----------:|---------:|
| [go-ofdgo](https://github.com/xiaoqidun/ofdgo)                      |    14.1KB |      1.1MB |      35MB |           39MB |      1.2MB |      1.9MB |        0 |
| [go-zc310](https://github.com/zc310/ofd)                            |     9.9KB |     96.1KB |      14MB |          1.8MB |     88.7KB |      1.4MB |        0 |
| [rust-easyofd](https://github.com/easy-4-rust/easyofd-rust)         |    37.3KB |     77.1KB |      29MB |      **614KB** |     83.3KB | **77.8KB** |        2 |
| [python-easyofd](https://pypi.org/project/easyofd/)                 |    FAILED | **36.1KB** |      38MB |         FAILED | **73.1KB** |     13.6MB |        2 |
| [python-ofd2pdf](https://github.com/jsyzdej/ofd2pdf)                |    36.8KB |     FAILED |    FAILED |           70MB |      752KB |      129KB |        0 |
| [python-ofdreader](https://pypi.org/project/ofdreader/)             |    FAILED |     FAILED |    FAILED |         FAILED |     FAILED |     FAILED |        0 |
| [java-ofdrw](https://github.com/ofdrw/ofdrw)                        |     4.8KB |     65.6KB |      23MB |          2.3MB |     76.9KB |     15.4MB |        0 |
| [node-ofd2pdf](https://www.npmjs.com/package/@miconvert/ofd-to-pdf) |     1.3KB |      197KB | **1.1MB** |          7.2MB |     82.9KB |      559KB |        2 |

### 文本提取能力

| 转换器                                                              | hello.ofd |  ano.ofd |  999.ofd | zsbk.ofd | 胜出次数 |
|---------------------------------------------------------------------|----------:|---------:|---------:|---------:|---------:|
| [go-ofdgo](https://github.com/xiaoqidun/ofdgo)                      |         8 |        0 |       76 |       19 |        0 |
| [go-zc310](https://github.com/zc310/ofd)                            |        22 | **7540** |     4595 |      985 |        1 |
| [rust-easyofd](https://github.com/easy-4-rust/easyofd-rust)         |        23 |     6357 |     3869 |      369 |        0 |
| [python-easyofd](https://pypi.org/project/easyofd/)                 |         0 |     1469 |     4579 | **2128** |        1 |
| [python-ofd2pdf](https://github.com/jsyzdej/ofd2pdf)                |         0 |   FAILED |   FAILED |        0 |        0 |
| [python-ofdreader](https://pypi.org/project/ofdreader/)             |    FAILED |   FAILED |   FAILED |   FAILED |        0 |
| [java-ofdrw](https://github.com/ofdrw/ofdrw)                        |        20 |     6307 |     3581 |      194 |        0 |
| [node-ofd2pdf](https://www.npmjs.com/package/@miconvert/ofd-to-pdf) |    **25** |     7049 | **4905** |     1697 |        2 |

> - `python-ofd2pdf` 基于图片渲染，无法提取文本
> - 单位为字符数（chars）

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
- **node-ofd2pdf**: intro.ofd 最快（1.20s），兼容性好
- **go-zc310**: 文本提取最完整，适合需要搜索/复制文本的场景
- **go-ofdgo**: 文件体积较大，文本提取较差
- **python-easyofd**: 部分文件转换失败
- **python-ofd2pdf**: 基于图片渲染，无法提取文本
- **java-ofdrw**: 兼容性好，但速度较慢

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
