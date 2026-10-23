<!-- 本文件由 tools/gen-libraries 生成，勿直接编辑；数据源 docs/libraries.yaml -->
# OFD 库汇总
> **预览**：库是否内置 OFD 预览/显示能力（查看器、浏览器渲染或可视化组件）。
> 
> **发票**：库是否支持电子发票类 OFD 处理（发票解析/信息提取、验签验章、发票渲染等）。⚠️ 为可兼容发票文档的通用 OFD 库，但无发票专项功能。
> 
> **→ HTML**：库是否支持 OFD → HTML（网页/浏览器渲染）输出。⚠️ 为在浏览器中基于 Canvas 渲染显示但不直接生成 HTML DOM/SVG 结构的库。
> 
> **PDF → OFD**：库是否支持 PDF 转 OFD 导入。⚠️ 为部分支持（仅文本/图片、或对非嵌入字体等常见 PDF 直接失败）。
> 
> **加密文档**：库是否支持 OFD 文档级加密/解密（GM/T 0099-2020 密码应用：口令/证书加密、解密打开）。⚠️ 为部分支持（如仅加密不支持解密打开）。据公开资料核查。
> 
> **打印**：库是否内置 OFD 打印能力（系统打印管线/打印对话框，或按页渲染后走浏览器/系统打印）。
> 
> **检索**：库是否支持文档内全文检索/关键词定位（含跨对象文字搜索、命中跳页）。
> 
> **注释**：库是否支持 OFD 注释对象（Annots）的读取（解析/渲染）或写入。
> 
> **归档**：库是否支持长期归档能力（归档预检/清单/封装准备，如 GB/T 33190、GB/T 42133）。
> 
> **编辑器**：库是否附带可视化编辑界面（GUI/WYSIWYG，可交互增删改文档内容），区别于「修改」的 API 程序化修改。

## Go

|                            库                            |   许可证   | 预览 | 解析 | 发票 | → 图片 | → PDF | → TXT | → SVG | → Md | → HTML | PDF → OFD | 生成 | 签章 | 加密文档 | 打印 | 检索 | 注释 | 归档 | 修改 | 编辑器 |
|:--------------------------------------------------------:|:----------:|:----:|:----:|:----:|:------:|:-----:|:-----:|:-----:|:----:|:------:|:---------:|:----:|:----:|:--------:|:----:|:----:|:----:|:----:|:----:|:------:|
|         [go-zc310](https://github.com/zc310/ofd)         | Apache-2.0 |  ✅  |  ✅  |  ✅  |   ✅   |  ✅   |  ✅   |  ✅   |  ✅  |   ✅   |    ✅     |  ✅  |  ✅  |    ✅    |  ❌  |  ✅  |  ❌  |  ✅  |  ✅  |   ❌   |
|      [go-ofdgo](https://github.com/xiaoqidun/ofdgo)      | Apache-2.0 |  ✅  |  ✅  |  ✅  |   ✅   |  ✅   |  ✅   |  ✅   |  ❌  |   ❌   |    ✅     |  ✅  |  ✅  |    ❌    |  ❌  |  ✅  |  ❌  |  ❌  |  ✅  |   ✅   |
|       [ofd-go](https://github.com/itlabers/ofd-go)       | Apache-2.0 |  ❌  |  ✅  |  ⚠️  |   ❌   |  ❌   |  ❌   |  ❌   |  ❌  |   ❌   |    ❌     |  ❌  |  ✅  |    ❌    |  ❌  |  ❌  |  ❌  |  ❌  |  ❌  |   ❌   |
|       [go-ofd](https://github.com/ppxz2014/go-ofd)       | Apache-2.0 |  ❌  |  ✅  |  ❌  |   ❌   |  ❌   |  ❌   |  ❌   |  ❌  |   ❌   |    ❌     |  ❌  |  ✅  |    ❌    |  ❌  |  ❌  |  ❌  |  ❌  |  ❌  |   ❌   |
|      [leijacob-ofd](https://gitee.com/leijacob/ofd)      |     无     |  ❌  |  ✅  |  ❌  |   ❌   |  ❌   |  ❌   |  ❌   |  ❌  |   ❌   |    ✅     |  ✅  |  ✅  |    ❌    |  ❌  |  ❌  |  ⚠️  |  ❌  |  ✅  |   ❌   |
| [signer-tools](https://gitee.com/leijacob/signer-tools)  |     无     |  ❌  |  ✅  |  ❌  |   ❌   |  ❌   |  ❌   |  ❌   |  ❌  |   ❌   |    ❌     |  ❌  |  ✅  |    ❌    |  ❌  |  ❌  |  ❌  |  ❌  |  ❌  |   ❌   |
| [pdf-to-ofd](https://github.com/huanxiang666/pdf-to-ofd) |    MIT     |  ❌  |  ❌  |  ❌  |   ❌   |  ❌   |  ❌   |  ❌   |  ❌  |   ❌   |    ⚠️     |  ❌  |  ❌  |    ❌    |  ❌  |  ❌  |  ❌  |  ❌  |  ❌  |   ❌   |

## Rust

|                             库                              |   许可证   | 预览 | 解析 | 发票 | → 图片 | → PDF | → TXT | → SVG | → Md | → HTML | PDF → OFD | 生成 | 签章 | 加密文档 | 打印 | 检索 | 注释 | 归档 | 修改 | 编辑器 |
|:-----------------------------------------------------------:|:----------:|:----:|:----:|:----:|:------:|:-----:|:-----:|:-----:|:----:|:------:|:---------:|:----:|:----:|:--------:|:----:|:----:|:----:|:----:|:----:|:------:|
| [easyofd-rust](https://github.com/easy-4-rust/easyofd-rust) | Apache-2.0 |  ❌  |  ✅  |  ⚠️  |   ✅   |  ✅   |  ❌   |  ❌   |  ✅  |   ❌   |    ⚠️     |  ✅  |  ✅  |    ✅    |  ❌  |  ✅  |  ❌  |  ✅  |  ✅  |   ❌   |
|     [ofdmanager](https://github.com/feuvan/ofdmanager)      |    MIT     |  ✅  |  ✅  |  ❌  |   ✅   |  ❌   |  ❌   |  ❌   |  ❌  |   ❌   |    ❌     |  ❌  |  ❌  |    ❌    |  ❌  |  ❌  |  ⚠️  |  ❌  |  ❌  |   ❌   |
|        [rs_ofd](https://github.com/geniusnut/rs_ofd)        |    MIT     |  ❌  |  ✅  |  ❌  |   ✅   |  ❌   |  ❌   |  ❌   |  ❌  |   ❌   |    ❌     |  ❌  |  ❌  |    ❌    |  ❌  |  ❌  |  ❌  |  ❌  |  ❌  |   ❌   |
|  [ofd-utility](https://github.com/ofd-utility/ofd-utility)  |    MIT     |  ❌  |  ✅  |  ❌  |   ✅   |  ❌   |  ❌   |  ❌   |  ❌  |   ❌   |    ❌     |  ❌  |  ✅  |    ❌    |  ❌  |  ❌  |  ✅  |  ❌  |  ❌  |   ❌   |
|       [ofdbox-rust](https://github.com/ofdbox/ofdbox)       | Apache-2.0 |  ❌  |  ✅  |  ❌  |   ✅   |  ✅   |  ❌   |  ❌   |  ❌  |   ❌   |    ⚠️     |  ❌  |  ✅  |    ✅    |  ❌  |  ❌  |  ✅  |  ❌  |  ✅  |   ❌   |
|         [ofdsdk](https://github.com/KaiserY/ofdsdk)         | Apache-2.0 |  ❌  |  ✅  |  ❌  |   ❌   |  ❌   |  ❌   |  ❌   |  ❌  |   ❌   |    ❌     |  ❌  |  ❌  |    ❌    |  ❌  |  ❌  |  ❌  |  ❌  |  ❌  |   ❌   |
|          [rofd](https://github.com/office-rs/rofd)          | Apache-2.0 |  ✅  |  ✅  |  ❌  |   ⚠️   |  ❌   |  ❌   |  ❌   |  ❌  |   ❌   |    ❌     |  ❌  |  ❌  |    ❌    |  ❌  |  ✅  |  ✅  |  ❌  |  ✅  |   ❌   |
|   [ofd-viewer](https://github.com/shaoyi1998/ofd-viewer)    |    MIT     |  ✅  |  ✅  |  ❌  |   ❌   |  ✅   |  ✅   |  ❌   |  ❌  |   ❌   |    ❌     |  ❌  |  ❌  |    ❌    |  ✅  |  ✅  |  ✅  |  ❌  |  ❌  |   ❌   |
|    [fapiao-print](https://github.com/erma0/fapiao-print)    |    MIT     |  ✅  |  ✅  |  ✅  |   ⚠️   |  ✅   |  ❌   |  ✅   |  ❌  |   ❌   |    ❌     |  ❌  |  ❌  |    ❌    |  ❌  |  ❌  |  ❌  |  ❌  |  ❌  |   ❌   |
|      [jian-yue](https://github.com/lapin2024/jian-yue)      |    MIT     |  ✅  |  ✅  |  ⚠️  |   ✅   |  ⚠️   |  ✅   |  ❌   |  ❌  |   ❌   |    ⚠️     |  ❌  |  ⚠️  |    ❌    |  ✅  |  ✅  |  ✅  |  ❌  |  ⚠️  |   ✅   |
|   [linuxdeepin-rofd](https://github.com/linuxdeepin/rofd)   |  LGPL-2.1  |  ✅  |  ✅  |  ⚠️  |   ✅   |  ❌   |  ❌   |  ❌   |  ❌  |   ❌   |    ❌     |  ❌  |  ❌  |    ❌    |  ❌  |  ❌  |  ❌  |  ❌  |  ❌  |   ❌   |
|      [easy_print](https://github.com/qzue/easy_print)       |    MIT     |  ✅  |  ✅  |  ✅  |   ❌   |  ✅   |  ❌   |  ❌   |  ❌  |   ❌   |    ❌     |  ❌  |  ❌  |    ❌    |  ✅  |  ❌  |  ❌  |  ❌  |  ❌  |   ❌   |
|        [ofd-rs](https://github.com/YeaLooooo/ofd-rs)        | Apache-2.0 |  ❌  |  ❌  |  ❌  |   ❌   |  ❌   |  ❌   |  ❌   |  ❌  |   ❌   |    ❌     |  ✅  |  ❌  |    ❌    |  ❌  |  ❌  |  ❌  |  ❌  |  ❌  |   ❌   |
|        [SteleOFD](https://github.com/emfox/SteleOFD)        | Apache-2.0 |  ✅  |  ✅  |  ⚠️  |   ✅   |  ✅   |  ❌   |  ✅   |  ❌  |   ⚠️   |    ❌     |  ❌  |  ✅  |    ❌    |  ✅  |  ✅  |  ✅  |  ✅  |  ✅  |   ❌   |

## Python

|                                        库                                        |   许可证   | 预览 | 解析 | 发票 | → 图片 | → PDF | → TXT | → SVG | → Md | → HTML | PDF → OFD | 生成 | 签章 | 加密文档 | 打印 | 检索 | 注释 | 归档 | 修改 | 编辑器 |
|:--------------------------------------------------------------------------------:|:----------:|:----:|:----:|:----:|:------:|:-----:|:-----:|:-----:|:----:|:------:|:---------:|:----:|:----:|:--------:|:----:|:----:|:----:|:----:|:----:|:------:|
|                   [easyofd](https://pypi.org/project/easyofd/)                   | Apache-2.0 |  ✅  |  ✅  |  ✅  |   ✅   |  ⚠️   |  ❌   |  ❌   |  ❌  |   ❌   |    ✅     |  ✅  |  ❌  |    ❌    |  ❌  |  ❌  |  ❌  |  ❌  |  ❌  |   ❌   |
|          [fapiao-ofd2pdf](https://github.com/seawander/fapiao-ofd2pdf)           |    MIT     |  ❌  |  ✅  |  ⚠️  |   ❌   |  ✅   |  ❌   |  ❌   |  ❌  |   ❌   |    ❌     |  ❌  |  ❌  |    ❌    |  ❌  |  ❌  |  ✅  |  ❌  |  ❌  |   ❌   |
|                   [ofd2img](https://pypi.org/project/ofd2img/)                   | Apache-2.0 |  ❌  |  ✅  |  ⚠️  |   ✅   |  ⚠️   |  ❌   |  ❌   |  ❌  |   ❌   |    ❌     |  ❌  |  ❌  |    ❌    |  ❌  |  ❌  |  ❌  |  ❌  |  ❌  |   ❌   |
|                  [ofd2pdf](https://github.com/jsyzdej/ofd2pdf)                   |     无     |  ❌  |  ✅  |  ❌  |   ✅   |  ⚠️   |  ❌   |  ❌   |  ❌  |   ❌   |    ❌     |  ❌  |  ❌  |    ❌    |  ❌  |  ❌  |  ✅  |  ❌  |  ❌  |   ❌   |
|                [OFDtoPDF](https://github.com/njuzzy1979/OFDtoPDF)                |     无     |  ❌  |  ✅  |  ❌  |   ❌   |  ✅   |  ❌   |  ❌   |  ❌  |   ❌   |    ❌     |  ❌  |  ⚠️  |    ❌    |  ❌  |  ❌  |  ❌  |  ❌  |  ❌  |   ❌   |
|                 [ofdreader](https://pypi.org/project/ofdreader/)                 |    MIT     |  ❌  |  ✅  |  ❌  |   ❌   |  ✅   |  ✅   |  ❌   |  ❌  |   ❌   |    ❌     |  ✅  |  ❌  |    ❌    |  ❌  |  ❌  |  ❌  |  ❌  |  ❌  |   ❌   |
|                 [pdf2ofd](https://github.com/wanglrebe/pdf2ofd)                  |    MIT     |  ❌  |  ❌  |  ❌  |   ❌   |  ❌   |  ❌   |  ❌   |  ❌  |   ❌   |    ✅     |  ❌  |  ❌  |    ❌    |  ❌  |  ❌  |  ❌  |  ❌  |  ❌  |   ❌   |
|               [ofd-parser](https://github.com/jyh2012/ofd-parser)                |     无     |  ✅  |  ✅  |  ❌  |   ❌   |  ❌   |  ❌   |  ❌   |  ❌  |   ❌   |    ❌     |  ❌  |  ❌  |    ❌    |  ❌  |  ❌  |  ❌  |  ❌  |  ❌  |   ❌   |
| [OfficeMaster](https://github.com/Chingliu/OfficeMaster_document_convert_system) |    MIT     |  ❌  |  ✅  |  ❌  |   ❌   |  ✅   |  ❌   |  ❌   |  ❌  |   ❌   |    ✅     |  ✅  |  ❌  |    ❌    |  ❌  |  ❌  |  ❌  |  ❌  |  ❌  |   ❌   |
|          [swingmonkey-ofd2pdf](https://github.com/swingmonkey/ofd2pdf)           |    MIT     |  ❌  |  ✅  |  ⚠️  |   ❌   |  ✅   |  ❌   |  ❌   |  ❌  |   ❌   |    ❌     |  ❌  |  ❌  |    ❌    |  ❌  |  ❌  |  ❌  |  ❌  |  ❌  |   ❌   |
|          [win-doc-converter](https://github.com/yefd/win-doc-converter)          |     无     |  ❌  |  ✅  |  ⚠️  |   ❌   |  ✅   |  ❌   |  ❌   |  ❌  |   ❌   |    ❌     |  ❌  |  ❌  |    ❌    |  ❌  |  ❌  |  ❌  |  ❌  |  ❌  |   ❌   |
|    [ce-conversion-engine](https://github.com/ThePandaQ/ce-conversion-engine)     |  AGPL-3.0  |  ❌  |  ✅  |  ⚠️  |   ✅   |  ⚠️   |  ✅   |  ❌   |  ❌  |   ❌   |    ✅     |  ✅  |  ❌  |    ❌    |  ❌  |  ❌  |  ❌  |  ❌  |  ❌  |   ❌   |

## Java

|                                     库                                      |   许可证   | 预览 | 解析 | 发票 | → 图片 | → PDF | → TXT | → SVG | → Md | → HTML | PDF → OFD | 生成 | 签章 | 加密文档 | 打印 | 检索 | 注释 | 归档 | 修改 | 编辑器 |
|:---------------------------------------------------------------------------:|:----------:|:----:|:----:|:----:|:------:|:-----:|:-----:|:-----:|:----:|:------:|:---------:|:----:|:----:|:--------:|:----:|:----:|:----:|:----:|:----:|:------:|
|                   [ofdrw](https://github.com/ofdrw/ofdrw)                   | Apache-2.0 |  ❌  |  ✅  |  ✅  |   ✅   |  ✅   |  ✅   |  ✅   |  ❌  |   ✅   |    ✅     |  ✅  |  ✅  |    ✅    |  ❌  |  ❌  |  ✅  |  ✅  |  ✅  |   ❌   |
|          [easyofd-java](https://github.com/11627685/easyofd-java)           | Apache-2.0 |  ❌  |  ✅  |  ❌  |   ❌   |  ❌   |  ❌   |  ❌   |  ❌  |   ❌   |    ❌     |  ✅  |  ✅  |    ❌    |  ❌  |  ❌  |  ❌  |  ❌  |  ❌  |   ❌   |
|                 [ofdbox](https://gitee.com/bookhhu/ofdbox)                  | Apache-2.0 |  ❌  |  ✅  |  ❌  |   ✅   |  ⚠️   |  ❌   |  ❌   |  ❌  |   ❌   |    ❌     |  ❌  |  ❌  |    ❌    |  ❌  |  ❌  |  ✅  |  ❌  |  ❌  |   ❌   |
|            [ofd-analyze](https://github.com/cooker/ofd-analyze)             |     无     |  ⚠️  |  ✅  |  ✅  |   ✅   |  ❌   |  ❌   |  ❌   |  ❌  |   ❌   |    ❌     |  ❌  |  ❌  |    ❌    |  ❌  |  ❌  |  ❌  |  ❌  |  ❌  |   ❌   |
|              [ofdToPdf](https://github.com/thebigboy/ofdToPdf)              |     无     |  ❌  |  ✅  |  ❌  |   ❌   |  ✅   |  ❌   |  ❌   |  ❌  |   ❌   |    ❌     |  ❌  |  ❌  |    ❌    |  ❌  |  ❌  |  ❌  |  ❌  |  ❌  |   ❌   |
| [JIMU-ConvertPreview](https://github.com/zhangzhen1979/JIMU-ConvertPreview) | Apache-2.0 |  ✅  |  ✅  |  ❌  |   ✅   |  ✅   |  ❌   |  ❌   |  ❌  |   ❌   |    ✅     |  ✅  |  ❌  |    ❌    |  ❌  |  ❌  |  ❌  |  ❌  |  ⚠️  |   ❌   |
|       [ofdbox-viewer](https://github.com/jiayao-zhang/ofdbox-viewer)        |     无     |  ✅  |  ✅  |  ❌  |   ✅   |  ❌   |  ❌   |  ❌   |  ❌  |   ⚠️   |    ❌     |  ❌  |  ❌  |    ❌    |  ❌  |  ❌  |  ❌  |  ❌  |  ❌  |   ❌   |
|            [kkFileView](https://github.com/kekingcn/kkFileView)             |     无     |  ✅  |  ✅  |  ✅  |   ✅   |  ⚠️   |  ❌   |  ❌   |  ❌  |   ⚠️   |    ❌     |  ❌  |  ❌  |    ❌    |  ⚠️  |  ❌  |  ❌  |  ❌  |  ⚠️  |   ❌   |
|         [kkfileviewxg](https://github.com/gaoxingzaq/kkfileviewxg)          |     无     |  ✅  |  ✅  |  ❌  |   ✅   |  ✅   |  ❌   |  ✅   |  ❌  |   ⚠️   |    ❌     |  ❌  |  ❌  |    ❌    |  ⚠️  |  ❌  |  ❌  |  ❌  |  ⚠️  |   ❌   |
|              [ofd-server](https://github.com/maczh/ofd-server)              |     无     |  ❌  |  ❌  |  ❌  |   ❌   |  ❌   |  ❌   |  ❌   |  ❌  |   ❌   |    ✅     |  ✅  |  ✅  |    ❌    |  ❌  |  ❌  |  ❌  |  ❌  |  ✅  |   ❌   |
|          [ofd-parser-tika](https://github.com/ryecrow/ofd-parser)           | Apache-2.0 |  ❌  |  ✅  |  ❌  |   ❌   |  ❌   |  ✅   |  ❌   |  ❌  |   ❌   |    ❌     |  ❌  |  ❌  |    ❌    |  ❌  |  ❌  |  ❌  |  ❌  |  ❌  |   ❌   |
|         [fuyue-convert](https://github.com/wmforever/fuyue-convert)         | Apache-2.0 |  ❌  |  ✅  |  ⚠️  |   ✅   |  ✅   |  ✅   |  ❌   |  ❌  |   ❌   |    ⚠️     |  ❌  |  ⚠️  |    ❌    |  ❌  |  ❌  |  ❌  |  ❌  |  ❌  |   ❌   |

## C++

|                             库                             |    许可证    | 预览 | 解析 | 发票 | → 图片 | → PDF | → TXT | → SVG | → Md | → HTML | PDF → OFD | 生成 | 签章 | 加密文档 | 打印 | 检索 | 注释 | 归档 | 修改 | 编辑器 |
|:----------------------------------------------------------:|:------------:|:----:|:----:|:----:|:------:|:-----:|:-----:|:-----:|:----:|:------:|:---------:|:----:|:----:|:--------:|:----:|:----:|:----:|:----:|:----:|:------:|
| [ofdpdfsigner](https://github.com/fanzizheng/ofdpdfsigner) |   BSL-1.1    |  ✅  |  ✅  |  ❌  |   ✅   |  ✅   |  ❌   |  ✅   |  ❌  |   ❌   |    ✅     |  ❌  |  ✅  |    ❌    |  ❌  |  ❌  |  ❌  |  ❌  |  ❌  |   ❌   |
|      [ofdReader](https://github.com/Micats/ofdReader)      |      无      |  ✅  |  ✅  |  ❌  |   ✅   |  ❌   |  ❌   |  ❌   |  ❌  |   ❌   |    ❌     |  ❌  |  ⚠️  |    ❌    |  ❌  |  ❌  |  ✅  |  ❌  |  ❌  |   ❌   |
|   [XilouReader](https://github.com/Chingliu/XilouReader)   | BSD-3-Clause |  ✅  |  ✅  |  ✅  |   ✅   |  ✅   |  ❌   |  ❌   |  ❌  |   ❌   |    ✅     |  ✅  |  ✅  |    ❌    |  ❌  |  ❌  |  ✅  |  ❌  |  ✅  |   ❌   |
|    [xilou_core](https://github.com/Chingliu/xilou_core)    | BSD-3-Clause |  ✅  |  ✅  |  ⚠️  |   ✅   |  ❌   |  ❌   |  ❌   |  ❌  |   ❌   |    ❌     |  ❌  |  ❌  |    ❌    |  ❌  |  ❌  |  ❌  |  ❌  |  ❌  |   ❌   |
|      [docwriter](https://github.com/isee15/docwriter)      |     MIT      |  ❌  |  ❌  |  ❌  |   ❌   |  ❌   |  ❌   |  ❌   |  ❌  |   ❌   |    ❌     |  ✅  |  ❌  |    ❌    |  ❌  |  ❌  |  ❌  |  ❌  |  ❌  |   ❌   |
|        [libofd](https://github.com/uukuguy/libofd)         |     MIT      |  ✅  |  ✅  |  ❌  |   ✅   |  ❌   |  ❌   |  ❌   |  ❌  |   ❌   |    ✅     |  ✅  |  ❌  |    ❌    |  ❌  |  ❌  |  ❌  |  ❌  |  ✅  |   ❌   |
|    [OFDEditor](https://github.com/KikyoShaw/OFDEditor)     |  Apache-2.0  |  ✅  |  ✅  |  ❌  |   ❌   |  ❌   |  ❌   |  ❌   |  ❌  |   ❌   |    ❌     |  ✅  |  ❌  |    ❌    |  ✅  |  ❌  |  ❌  |  ❌  |  ✅  |   ✅   |
|    [OfdiumEx](https://github.com/roy19831015/OfdiumEx)     |  Apache-2.0  |  ✅  |  ✅  |  ❌  |   ✅   |  ❌   |  ❌   |  ❌   |  ❌  |   ❌   |    ❌     |  ❌  |  ❌  |    ❌    |  ❌  |  ❌  |  ❌  |  ❌  |  ❌  |   ❌   |
|    [ofdEditor](https://github.com/mcoder2014/ofdEditor)    |     MIT      |  ✅  |  ✅  |  ❌  |   ✅   |  ❌   |  ❌   |  ❌   |  ❌  |   ❌   |    ❌     |  ✅  |  ❌  |    ❌    |  ❌  |  ❌  |  ❌  |  ❌  |  ✅  |   ✅   |
|    [ZipViewer](https://github.com/CryFeiFei/ZipViewer)     |     MIT      |  ✅  |  ✅  |  ❌  |   ❌   |  ❌   |  ❌   |  ❌   |  ❌  |   ❌   |    ❌     |  ❌  |  ❌  |    ❌    |  ❌  |  ❌  |  ❌  |  ❌  |  ❌  |   ❌   |
|   [papers-ofd](https://github.com/YangtseSu/papers-ofd)    |   GPL-2.0    |  ✅  |  ✅  |  ⚠️  |   ✅   |  ✅   |  ✅   |  ❌   |  ❌  |   ❌   |    ❌     |  ❌  |  ❌  |    ❌    |  ✅  |  ✅  |  ❌  |  ❌  |  ❌  |   ❌   |

## .NET

|                                            库                                             |   许可证   | 预览 | 解析 | 发票 | → 图片 | → PDF | → TXT | → SVG | → Md | → HTML | PDF → OFD | 生成 | 签章 | 加密文档 | 打印 | 检索 | 注释 | 归档 | 修改 | 编辑器 |
|:-----------------------------------------------------------------------------------------:|:----------:|:----:|:----:|:----:|:------:|:-----:|:-----:|:-----:|:----:|:------:|:---------:|:----:|:----:|:--------:|:----:|:----:|:----:|:----:|:----:|:------:|
|                  [XiaoFeng.Ofd](https://github.com/zhuovi/XiaoFeng.Ofd)                   |    MIT     |  ❌  |  ✅  |  ❌  |   ❌   |  ❌   |  ❌   |  ❌   |  ❌  |   ❌   |    ❌     |  ✅  |  ✅  |    ❌    |  ❌  |  ❌  |  ✅  |  ❌  |  ❌  |   ❌   |
|                    [OfdViewer](https://github.com/LvYueMing/OfdViewer)                    |    MIT     |  ✅  |  ✅  |  ❌  |   ✅   |  ❌   |  ❌   |  ❌   |  ❌  |   ❌   |    ❌     |  ✅  |  ⚠️  |    ❌    |  ❌  |  ❌  |  ✅  |  ❌  |  ❌  |   ❌   |
|                     [ofdrw.net](https://github.com/whynpc9/ofdrw.net)                     |    MIT     |  ❌  |  ✅  |  ❌  |   ✅   |  ✅   |  ✅   |  ✅   |  ❌  |   ✅   |    ✅     |  ✅  |  ⚠️  |    ✅    |  ❌  |  ❌  |  ✅  |  ✅  |  ✅  |   ❌   |
|                   [ofdrw-net](https://github.com/lllooollpp/ofdrw-net)                    |     无     |  ❌  |  ✅  |  ❌  |   ✅   |  ✅   |  ❌   |  ❌   |  ❌  |   ❌   |    ✅     |  ✅  |  ✅  |    ✅    |  ❌  |  ❌  |  ✅  |  ❌  |  ❌  |   ❌   |
|                 [OFDConverter](https://github.com/wukonggo/OFDConverter)                  | Apache-2.0 |  ❌  |  ❌  |  ❌  |   ❌   |  ❌   |  ❌   |  ❌   |  ❌  |   ❌   |    ✅     |  ❌  |  ❌  |    ❌    |  ❌  |  ❌  |  ❌  |  ❌  |  ❌  |   ❌   |
|                    [ofdparser](https://github.com/wangyi160/ofdparser)                    |     无     |  ❌  |  ✅  |  ❌  |   ❌   |  ❌   |  ❌   |  ❌   |  ❌  |   ❌   |    ❌     |  ❌  |  ❌  |    ❌    |  ❌  |  ❌  |  ❌  |  ❌  |  ❌  |   ❌   |
| [BootstrapBlazor.OfdReader](https://github.com/BootstrapBlazor/BootstrapBlazor.OfdReader) |    MIT     |  ✅  |  ✅  |  ❌  |   ❌   |  ❌   |  ❌   |  ❌   |  ❌  |   ❌   |    ❌     |  ❌  |  ❌  |    ❌    |  ❌  |  ❌  |  ❌  |  ❌  |  ❌  |   ❌   |
|                      [Ofd2Pdf](https://github.com/taurusxin/Ofd2Pdf)                      |    MIT     |  ❌  |  ✅  |  ❌  |   ❌   |  ✅   |  ❌   |  ❌   |  ❌  |   ❌   |    ❌     |  ❌  |  ❌  |    ❌    |  ❌  |  ❌  |  ❌  |  ❌  |  ❌  |   ❌   |
|                      [ofd2pdf](https://github.com/lanbo0829/ofd2pdf)                      |     无     |  ❌  |  ✅  |  ❌  |   ❌   |  ✅   |  ❌   |  ❌   |  ❌  |   ❌   |    ❌     |  ❌  |  ❌  |    ❌    |  ❌  |  ❌  |  ❌  |  ❌  |  ❌  |   ❌   |

## JavaScript

|                                      库                                       |   许可证   | 预览 | 解析 | 发票 | → 图片 | → PDF | → TXT | → SVG | → Md | → HTML | PDF → OFD | 生成 | 签章 | 加密文档 | 打印 | 检索 | 注释 | 归档 | 修改 | 编辑器 |
|:-----------------------------------------------------------------------------:|:----------:|:----:|:----:|:----:|:------:|:-----:|:-----:|:-----:|:----:|:------:|:---------:|:----:|:----:|:--------:|:----:|:----:|:----:|:----:|:----:|:------:|
|       [ofd-to-pdf](https://www.npmjs.com/package/@miconvert/ofd-to-pdf)       | Apache-2.0 |  ❌  |  ✅  |  ✅  |   ❌   |  ✅   |  ❌   |  ❌   |  ❌  |   ❌   |    ❌     |  ❌  |  ❌  |    ❌    |  ❌  |  ❌  |  ❌  |  ❌  |  ❌  |   ❌   |
|                   [ofdjs](https://github.com/isee15/ofdjs)                    | Apache-2.0 |  ✅  |  ✅  |  ⚠️  |   ✅   |  ❌   |  ❌   |  ❌   |  ❌  |   ⚠️   |    ❌     |  ❌  |  ❌  |    ❌    |  ❌  |  ❌  |  ✅  |  ❌  |  ❌  |   ❌   |
|         [imageConversion](https://github.com/Gary-zy/imageConversion)         |    MIT     |  ✅  |  ✅  |  ❌  |   ✅   |  ✅   |  ❌   |  ❌   |  ❌  |   ❌   |    ❌     |  ❌  |  ⚠️  |    ❌    |  ❌  |  ❌  |  ❌  |  ❌  |  ❌  |   ❌   |
|                [liteofd](https://github.com/SignitDoc/liteofd)                | Apache-2.0 |  ✅  |  ✅  |  ⚠️  |   ✅   |  ❌   |  ✅   |  ❌   |  ❌  |   ✅   |    ❌     |  ❌  |  ⚠️  |    ❌    |  ❌  |  ✅  |  ✅  |  ❌  |  ❌  |   ❌   |
|             [bestofdview](https://github.com/besthqs/bestofdview)             |     无     |  ✅  |  ✅  |  ❌  |   ✅   |  ❌   |  ❌   |  ❌   |  ❌  |   ⚠️   |    ❌     |  ❌  |  ✅  |    ❌    |  ✅  |  ❌  |  ❌  |  ❌  |  ❌  |   ❌   |
|                [OFDView](https://github.com/guinanlin/OFDView)                | Apache-2.0 |  ✅  |  ✅  |  ❌  |   ✅   |  ❌   |  ❌   |  ❌   |  ❌  |   ✅   |    ❌     |  ❌  |  ❌  |    ❌    |  ✅  |  ❌  |  ❌  |  ❌  |  ❌  |   ❌   |
|                 [ofd.js](https://github.com/DLTech21/ofd.js)                  | Apache-2.0 |  ✅  |  ✅  |  ⚠️  |   ✅   |  ❌   |  ❌   |  ❌   |  ❌  |   ✅   |    ❌     |  ❌  |  ⚠️  |    ❌    |  ❌  |  ❌  |  ❌  |  ❌  |  ❌  |   ❌   |
|               [ofd-online](https://github.com/betgo/ofd-online)               |     无     |  ✅  |  ✅  |  ❌  |   ✅   |  ❌   |  ❌   |  ✅   |  ❌  |   ⚠️   |    ❌     |  ❌  |  ❌  |    ❌    |  ❌  |  ❌  |  ❌  |  ❌  |  ❌  |   ❌   |
|              [ofdviewer](https://github.com/xxss0903/ofdviewer)               |    MIT     |  ✅  |  ✅  |  ❌  |   ❌   |  ❌   |  ❌   |  ❌   |  ❌  |   ⚠️   |    ❌     |  ❌  |  ❌  |    ❌    |  ❌  |  ❌  |  ❌  |  ❌  |  ❌  |   ❌   |
|                  [webofd](https://github.com/zsc347/webofd)                   |     无     |  ✅  |  ✅  |  ❌  |   ✅   |  ❌   |  ❌   |  ❌   |  ❌  |   ⚠️   |    ❌     |  ❌  |  ❌  |    ❌    |  ❌  |  ❌  |  ❌  |  ❌  |  ❌  |   ❌   |
| [ofd_editor_frontend](https://github.com/LamplightShadow/ofd_editor_frontend) |     无     |  ✅  |  ✅  |  ❌  |   ✅   |  ✅   |  ❌   |  ❌   |  ❌  |   ⚠️   |    ❌     |  ✅  |  ✅  |    ⚠️    |  ❌  |  ❌  |  ❌  |  ❌  |  ✅  |   ✅   |
|            [ofdjs-viewer](https://github.com/Atw-Lee/ofdjs-viewer)            |    MIT     |  ✅  |  ✅  |  ❌  |   ✅   |  ❌   |  ✅   |  ❌   |  ❌  |   ⚠️   |    ❌     |  ❌  |  ❌  |    ❌    |  ❌  |  ❌  |  ✅  |  ❌  |  ❌  |   ❌   |
|            [piaopinpin](https://github.com/qq87636108/piaopinpin)             |    MIT     |  ✅  |  ✅  |  ✅  |   ✅   |  ❌   |  ❌   |  ❌   |  ❌  |   ⚠️   |    ❌     |  ❌  |  ❌  |    ❌    |  ✅  |  ❌  |  ❌  |  ❌  |  ❌  |   ❌   |
|         [jit-viewer-sdk](https://github.com/Drexr9558/jit-viewer-sdk)         | Apache-2.0 |  ✅  |  ✅  |  ⚠️  |   ✅   |  ❌   |  ✅   |  ❌   |  ❌  |   ⚠️   |    ❌     |  ❌  |  ⚠️  |    ❌    |  ❌  |  ❌  |  ❌  |  ❌  |  ❌  |   ❌   |
|             [ofd-reader](https://github.com/Maple-pro/ofd-reader)             |     无     |  ✅  |  ✅  |  ⚠️  |   ✅   |  ❌   |  ❌   |  ❌   |  ❌  |   ⚠️   |    ❌     |  ❌  |  ❌  |    ❌    |  ✅  |  ❌  |  ❌  |  ❌  |  ❌  |   ❌   |
|                   [jsOFD](https://github.com/Hufe921/jsOFD)                   |    MIT     |  ❌  |  ❌  |  ❌  |   ❌   |  ❌   |  ❌   |  ❌   |  ❌  |   ❌   |    ✅     |  ✅  |  ❌  |    ❌    |  ❌  |  ❌  |  ❌  |  ❌  |  ❌  |   ❌   |
|                [ofd-lzy](https://github.com/Yusuf-dz/ofd-lzy)                 |  BUSL-1.1  |  ✅  |  ✅  |  ✅  |   ✅   |  ✅   |  ✅   |  ✅   |  ❌  |   ⚠️   |    ❌     |  ❌  |  ✅  |    ❌    |  ✅  |  ✅  |  ✅  |  ❌  |  ❌  |   ✅   |
|    [local-doc-viewer](https://github.com/LocalDoc-Viewer/local-doc-viewer)    |  AGPL-3.0  |  ✅  |  ✅  |  ⚠️  |   ❌   |  ❌   |  ❌   |  ❌   |  ❌  |   ❌   |    ❌     |  ❌  |  ❌  |    ❌    |  ❌  |  ❌  |  ❌  |  ❌  |  ❌  |   ❌   |
|        [ofd.js-js-nuxt](https://github.com/astonmario/ofd.js-js-nuxt)         | Apache-2.0 |  ✅  |  ✅  |  ⚠️  |   ❌   |  ❌   |  ❌   |  ❌   |  ❌  |   ✅   |    ❌     |  ❌  |  ⚠️  |    ❌    |  ❌  |  ❌  |  ❌  |  ❌  |  ❌  |   ❌   |
|                [ofdview](https://github.com/moxudocs/ofdview)                 | Apache-2.0 |  ✅  |  ✅  |  ✅  |   ❌   |  ❌   |  ❌   |  ❌   |  ❌  |   ⚠️   |    ❌     |  ❌  |  ✅  |    ❌    |  ❌  |  ✅  |  ✅  |  ❌  |  ❌  |   ❌   |
|           [file-viewer](https://github.com/flyfish-dev/file-viewer)           | Apache-2.0 |  ✅  |  ✅  |  ⚠️  |   ❌   |  ❌   |  ❌   |  ❌   |  ❌  |   ✅   |    ❌     |  ❌  |  ⚠️  |    ❌    |  ❌  |  ❌  |  ❌  |  ❌  |  ❌  |   ❌   |
|                   [ofdh5](https://github.com/iwlking/ofdh5)                   |    MIT     |  ✅  |  ✅  |  ❌  |   ❌   |  ❌   |  ❌   |  ❌   |  ❌  |   ❌   |    ❌     |  ❌  |  ❌  |    ❌    |  ❌  |  ❌  |  ❌  |  ❌  |  ❌  |   ❌   |
|            [ofd.js (LaanChae)](https://github.com/LaanChae/ofd.js)            |     无     |  ✅  |  ✅  |  ❌  |   ❌   |  ✅   |  ❌   |  ❌   |  ❌  |   ⚠️   |    ✅     |  ❌  |  ✅  |    ✅    |  ✅  |  ✅  |  ✅  |  ❌  |  ✅  |   ⚠️   |
|                [OFDee](https://github.com/gityongzhong/OFDee)                 |    MIT     |  ✅  |  ✅  |  ❌  |   ✅   |  ✅   |  ❌   |  ❌   |  ❌  |   ❌   |    ❌     |  ❌  |  ❌  |    ❌    |  ❌  |  ❌  |  ❌  |  ❌  |  ❌  |   ❌   |
|     [ofd-to-pdf (campfire)](https://github.com/campfire-code/ofd-to-pdf)      |    MIT     |  ❌  |  ✅  |  ⚠️  |   ❌   |  ✅   |  ❌   |  ❌   |  ❌  |   ❌   |    ❌     |  ❌  |  ❌  |    ❌    |  ❌  |  ❌  |  ❌  |  ❌  |  ❌  |   ❌   |
|             [ofdrw-viewer](https://github.com/ofdrw/ofdrw-viewer)             |    MIT     |  ✅  |  ✅  |  ⚠️  |   ✅   |  ❌   |  ✅   |  ❌   |  ❌  |   ⚠️   |    ❌     |  ❌  |  ✅  |    ✅    |  ✅  |  ✅  |  ❌  |  ❌  |  ❌  |   ❌   |

## Pascal

|                         库                         |   许可证    | 预览 | 解析 | 发票 | → 图片 | → PDF | → TXT | → SVG | → Md | → HTML | PDF → OFD | 生成 | 签章 | 加密文档 | 打印 | 检索 | 注释 | 归档 | 修改 | 编辑器 |
|:--------------------------------------------------:|:-----------:|:----:|:----:|:----:|:------:|:-----:|:-----:|:-----:|:----:|:------:|:---------:|:----:|:----:|:--------:|:----:|:----:|:----:|:----:|:----:|:------:|
| [tinyofd](https://github.com/miemiekurisu/tinyofd) | PolyForm-NC |  ✅  |  ✅  |  ❌  |   ❌   |  ❌   |  ✅   |  ❌   |  ❌  |   ❌   |    ❌     |  ❌  |  ❌  |    ❌    |  ❌  |  ❌  |  ❌  |  ❌  |  ❌  |   ❌   |

