# easyofd-rust 集成说明

## 目录结构

```
rust/
├── README.md              # 本文件
├── ofd2pdf/               # OFD → PDF 转换器（独立 Rust 项目）
│   ├── Cargo.toml
│   └── src/
│       └── main.rs
└── pdf2ofd/               # PDF → OFD 转换器（独立 Rust 项目）
    ├── Cargo.toml
    └── src/
        └── main.rs
```

## 方式1：使用独立 Rust 项目（推荐）

```bash
# OFD → PDF
cd rust/ofd2pdf
cargo build --release
./target/release/ofd2pdf input.ofd output.pdf

# PDF → OFD
cd rust/pdf2ofd
cargo build --release
./target/release/pdf2ofd input.pdf output.ofd
```

## 方式2：编译 easyofd-tool

```bash
# 克隆源码
cd /tmp
git clone https://github.com/easy-4-rust/easyofd-rust.git
cd easyofd-rust

# 编译
cargo build --release -p easyofd-tool

# 使用
./target/release/easyofd to-pdf input.ofd output.pdf
```

## 注意事项

- 需要安装 Rust 工具链：`curl --proto '=https' --tlsv1.2 -sSf https://sh.rustup.rs | sh`
- 系统需要安装 CJK 字体才能正确渲染中文
- 安装中文字体：`apt-get install fonts-wqy-zenhei`
