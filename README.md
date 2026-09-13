# OFD to PDF 速度对比测试

## 测试库

### 已测试（支持 OFD→PDF）

- [github.com/xiaoqidun/ofdgo](https://github.com/xiaoqidun/ofdgo) - 纯 Go OFD 渲染库
- [github.com/zc310/ofd](https://github.com/zc310/ofd) - OFD 文件转换、创建、浏览工具
- [github.com/easy-4-rust/easyofd-rust](https://github.com/easy-4-rust/easyofd-rust) - Rust OFD 转换库

### 未测试（不支持 OFD→PDF）

| 库 | 原因 |
|---|------|
| [github.com/feuvan/ofdmanager](https://github.com/feuvan/ofdmanager) | 仅支持 OFD→PNG，桌面应用有基于图片的 PDF 导出但 CLI 不支持 |
| [github.com/geniusnut/rs_ofd](https://github.com/geniusnut/rs_ofd) | 仅支持 OFD→PNG |

## 测试结果

| 测试文件 | OFD大小 | ofdgo | ofdgo PDF | zc310 | zc310 PDF | easyofd-rust | easyofd-rust PDF | 胜出 |
|---------|--------:|------:|----------:|------:|----------:|-----:|---------:|------|
| hello.ofd | 1.5K | 6ms | 0.7KB | 16ms | 4.9KB | 7ms | 1.5KB | ofdgo 1.2x |
| ano.ofd | 702K | 212ms | 1099KB | 207ms | 94KB | **14ms** | 5KB | **rust 15x** |
| intro.ofd | 7.2M | 2.39s | 36MB | **1.26s** | 14MB | 1.50s | 29MB | zc310 1.2x |
| 1000-pages.ofd | 456K | 517ms | 1056KB | 2.42s | 1459KB | **54ms** | 571KB | **rust 9.6x** |
| 999.ofd | 30K | 19ms | 6KB | 102ms | 25KB | **12ms** | 8KB | rust 1.6x |
| zsbk.ofd | 1.5M | 305ms | 1886KB | 303ms | 1459KB | **19ms** | 4KB | **rust 16x** |

## 总结

- **easyofd-rust**: 多页/大文件场景下最快（1000页快 9.6x，zsbk 快 16x），输出 PDF 体积最小
- **ofdgo**: 简单文件表现良好
- **zc310/ofd**: 在 7.2M 复杂大文件上最快
- Rust 版本有字体缺失警告，不影响速度对比

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

## 项目结构

```
ofd-benchmark/
├── bin/
│   ├── ofd-benchmark
│   ├── ofdgo-convert
│   ├── zc310-convert
│   └── easyofd          # Rust 二进制
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
