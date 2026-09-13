# OFD to PDF 速度对比测试

## 测试库

### 已测试（支持 OFD→PDF）

| 库                                                                                 | 版本                  |
|------------------------------------------------------------------------------------|-----------------------|
| [github.com/xiaoqidun/ofdgo](https://github.com/xiaoqidun/ofdgo)                   | v0.0.0-20260912163201 |
| [github.com/zc310/ofd](https://github.com/zc310/ofd)                               | v0.1.0                |
| [github.com/easy-4-rust/easyofd-rust](https://github.com/easy-4-rust/easyofd-rust) | v0.1.2                |

### 未测试（不支持 OFD→PDF）

| 库                                                                   | 原因                                                       |
|----------------------------------------------------------------------|------------------------------------------------------------|
| [github.com/feuvan/ofdmanager](https://github.com/feuvan/ofdmanager) | 仅支持 OFD→PNG，桌面应用有基于图片的 PDF 导出但 CLI 不支持 |
| [github.com/geniusnut/rs_ofd](https://github.com/geniusnut/rs_ofd)   | 仅支持 OFD→PNG                                             |

## 测试环境

- 字体目录：`~/.local/share/fonts`（已安装各种中文字体）
- 系统已安装 `apt-get install fonts-wqy-zenhei`

## 测试结果

| 测试文件       | OFD大小 | ofdgo |     zc310 | easyofd-rust | 胜出          |
|----------------|--------:|------:|----------:|-------------:|---------------|
| hello.ofd      |    1.5K |  86ms |     134ms |     **53ms** | rust 1.6x     |
| ano.ofd        |    702K | 218ms |     353ms |     **96ms** | **rust 2.3x** |
| intro.ofd      |    7.2M | 2.41s | **1.24s** |        1.97s | zc310 1.6x    |
| 1000-pages.ofd |    456K | 8.35s |     8.70s |    **1.11s** | **rust 7.5x** |
| 999.ofd        |     30K | 318ms |     369ms |    **191ms** | rust 1.7x     |
| zsbk.ofd       |    1.5M | 320ms |     449ms |     **87ms** | **rust 3.7x** |

## PDF 质量对比

### 文件大小对比

| 文件           |  ofdgo |    zc310 |      rust | 最优  |
|----------------|-------:|---------:|----------:|-------|
| hello.ofd      |   15KB | **10KB** |      13KB | zc310 |
| ano.ofd        | 1119KB |     97KB |  **44KB** | rust  |
| intro.ofd      |   36MB | **14MB** |      30MB | zc310 |
| 1000-pages.ofd |   66MB |   1548KB | **591KB** | rust  |
| 999.ofd        | 1234KB |     89KB |  **50KB** | rust  |
| zsbk.ofd       | 1941KB |   1470KB |  **43KB** | rust  |

### 文本提取能力

| 文件      | ofdgo |      zc310 |   rust |
|-----------|------:|-----------:|-------:|
| hello.ofd |  19字 |       33字 |   34字 |
| ano.ofd   |   3字 | **9141字** | 6630字 |
| zsbk.ofd  |  23字 | **2947字** |  714字 |

### 质量总结

| 特性       | ofdgo | zc310 | easyofd-rust |
|------------|-------|-------|--------------|
| 页面尺寸   | A4    | A4    | A4           |
| 矢量图形   | 支持  | 支持  | 支持         |
| 文本可提取 | 部分  | 最佳  | 良好         |
| 图片嵌入   | 有    | 有    | 有           |
| 文件压缩   | 较差  | 良佳  | 最佳         |
| 中文支持   | 良好  | 最佳  | 良好         |

**结论：**
- **zc310**: 文本提取最完整，适合需要搜索/复制文本的场景
- **rust**: 压缩率最好，文件体积最小，速度最快
- **ofdgo**: 文件体积较大，文本提取较差

## 使用方法

```bash
# 编译
go build -o bin/ofd-benchmark .
go build -o bin/ofdgo-convert ./cmd/ofdgo-convert
go build -o bin/zc310-convert ./cmd/zc310-convert
cp <rust-binary>/easyofd bin/easyofd

# 运行测试（默认测试所有转换器）
./bin/ofd-benchmark <file.ofd>

# 只测试某个转换器
./bin/ofd-benchmark <file.ofd> ofdgo
./bin/ofd-benchmark <file.ofd> zc310
./bin/ofd-benchmark <file.ofd> rust
```

转换后的 PDF 文件保存在 `bin/output/` 目录，命名格式：`{原文件名}_{转换器}.pdf`

## 项目结构

```
ofd-benchmark/
├── bin/
│   ├── ofd-benchmark
│   ├── ofdgo-convert
│   ├── zc310-convert
│   ├── easyofd          # Rust 二进制
│   └── output/          # 转换后的 PDF 文件
├── cmd/
│   ├── ofdgo-convert/
│   └── zc310-convert/
├── rust/                # Rust 集成说明和示例
│   ├── README.md
│   └── ofd2pdf/
├── main.go
├── go.mod
└── README.md
```
