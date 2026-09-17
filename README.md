# OFD to PDF 速度对比测试

## 测试库

### Go 库

| 库                                                    | 版本                  | OFD→PDF |  更新时间  | 备注   |
|-------------------------------------------------------|-----------------------|:-------:|:----------:|--------|
| [xiaoqidun-ofdgo](https://github.com/xiaoqidun/ofdgo) | v0.0.0-20260915002242 |   ✅    | 2026-09-17 | 已测试 |
| [zc310-ofd](https://github.com/zc310/ofd)             | v0.1.1-20260915002405 |   ✅    | 2026-09-17 | 已测试 |

### Rust 库

| 库                                                          | 版本   | OFD→PDF |  更新时间  | 备注                         |
|-------------------------------------------------------------|--------|:-------:|:----------:|------------------------------|
| [easyofd-rust](https://github.com/easy-4-rust/easyofd-rust) | v0.1.2 |   ✅    | 2026-09-10 | 已测试                       |
| [ofdmanager](https://github.com/feuvan/ofdmanager)          | v0.2.0 |   ❌    | 2026-08-11 | 仅支持 OFD→PNG               |
| [rs_ofd](https://github.com/geniusnut/rs_ofd)               | -      |   ❌    | 2024-11-28 | 仅支持 OFD→PNG               |
| [ofd-utility](https://github.com/ofd-utility/ofd-utility)   | -      |   ❌    | 2026-08-07 | OFD 解析、校验、渲染到图片   |
| [ofdsdk](https://github.com/KaiserY/ofdsdk)                 | v0.2.3 |   ❌    | 2026-07-09 | OFD SDK，XML 解析和包读写    |
| [rofd](https://github.com/linuxdeepin/rofd)                 | -      |   ❌    | 2026-09-05 | deepin，Cairo 渲染，LGPL-2.1 |

### Python 库

| 库                                                                               | 版本     | OFD→PDF |  更新时间  | 备注                               |
|----------------------------------------------------------------------------------|----------|:-------:|:----------:|------------------------------------|
| [easyofd](https://pypi.org/project/easyofd/)                                     | 20260427 |   ⚠️    | 2026-04-27 | 部分文件失败                       |
| [ofd2img](https://pypi.org/project/ofd2img/)                                     | 0.1.2    |   ⚠️    | 2026-05-07 | 部分文件失败                       |
| [ofd2pdf](https://github.com/jsyzdej/ofd2pdf)                                    | 0.0.2    |   ⚠️    | 2026-07-08 | 部分文件失败，基于图片渲染         |
| [ofdreader](https://pypi.org/project/ofdreader/)                                 | 0.1.0    |   ❌    | 2026-06-01 | 空包，无法使用                     |
| [ofdparser](https://pypi.org/project/ofdparser/)                                 | 0.0.8    |   ❌    | 2023-04-11 | 缺 Courier 字体，无法导入          |
| [pdf2ofd](https://github.com/wanglrebe/pdf2ofd)                                  | 0.1.0    |    -    | 2026-05-22 | PDF→OFD                            |
| [ofd-parser](https://github.com/jyh2012/ofd-parser)                              | -        |   ❌    |     -      | OFD 解析工具，GUI 预览，lxml       |
| [OfficeMaster](https://github.com/Chingliu/OfficeMaster_document_convert_system) | -        |   ✅    | 2024-04-27 | 多格式转换系统，需 WPS/Office，MIT |

### Java 库

| 库                                                                          | 版本      | OFD→PDF | 更新时间   | 备注                                             |
|-----------------------------------------------------------------------------|-----------|:-------:|------------|--------------------------------------------------|
| [ofdrw](https://github.com/ofdrw/ofdrw)                                     | 2.4.0     |   ✅    | 2026-08-04 | 需手动下载依赖，全部成功                         |
| [easyofd-java](https://github.com/11627685/easyofd-java)                    | -         |   ❌    | 2026-09-08 | OFD 生成/读取/签章，不支持转 PDF                 |
| [ofdbox](https://gitee.com/bookhhu/ofdbox)                                  | -         |   ⚠️    | 2025-06-08 | OFD→图片可用，PDF 正在开发中                     |
| [ofdpdfsigner](https://github.com/fanzizheng/ofdpdfsigner)                  | -         |   ✅    | 2026-08-15 | C++ 实现，提供 Java JNI 接口                     |
| [ofd-analyze](https://github.com/cooker/ofd-analyze)                        | -         |   ❌    | 2022-05-20 | OFD 解析器，支持 OFD→图片                        |
| [ofdToPdf](https://github.com/thebigboy/ofdToPdf)                           | -         |   ✅    | 2025-09-17 | Spring Boot REST 服务，基于 OFDRW                |
| [JIMU-ConvertPreview](https://github.com/zhangzhen1979/JIMU-ConvertPreview) | 3.0.alpha |   ✅    | 2024-04-09 | 转换+预览服务，水印，需商业 license              |
| [ofdbox-viewer](https://github.com/jiayao-zhang/ofdbox-viewer)              | -         |   ❌    | 2023-07-26 | 后端渲染 OFD→图片/预览，Spring Boot              |
| [kkFileView](https://github.com/kekingcn/kkFileView)                        | 5.0.2     |   ⚠️    | 2026-08-14 | 通用文件预览，OFD→图片预览，14.7k star           |
| [kkfileviewxg](https://github.com/gaoxingzaq/kkfileviewxg)                  | 4.7.9.2   |   ✅    | 2022-10-22 | kkFileView 高雄修改版，OFD 转 SVG/PDF，已停维护  |
| [ofd-server](https://github.com/maczh/ofd-server)                           | -         |   ❌    | 2021-09-28 | Spring Boot 微服务，doc/pdf 生成 OFD、附件、签单 |
| [ofd-parser](https://github.com/ryecrow/ofd-parser)                         | -         |   ❌    | 2021-04-02 | Apache Tika 的 OFD 解析器                        |

### .NET 库

| 库                                                                                                                   | 版本            | OFD→PDF |  更新时间  | 备注                                        |
|----------------------------------------------------------------------------------------------------------------------|-----------------|:-------:|:----------:|---------------------------------------------|
| [github.com/zhuovi/XiaoFeng.Ofd](https://github.com/zhuovi/XiaoFeng.Ofd)                                             | 1.0.0           |   ❌    | 2025-06-03 | README 声称支持但代码未实现                 |
| [github.com/wangyi160/ofdparser](https://github.com/wangyi160/ofdparser)                                             | -               |   ❌    | 2022-02-11 | OFD 解析器                                  |
| [github.com/LvYueMing/OfdViewer](https://github.com/LvYueMing/OfdViewer)                                             | -               |   ❌    | 2026-08-01 | net8.0，SkiaSharp 渲染，MIT                 |
| [github.com/whynpc9/ofdrw.net](https://github.com/whynpc9/ofdrw.net)                                                 | 0.1.0-preview.7 |   ✅    | 2026-09-05 | netstandard2.0/2.1，OFD SDK，转换/编辑/签章 |
| [github.com/lllooollpp/ofdrw-net](https://github.com/lllooollpp/ofdrw-net)                                           | v2.0            |   ✅    | 2025-10-08 | PDF/OFD 互转、CLI、SM2/SM3 签章，无许可证   |
| [github.com/wukonggo/OFDConverter](https://github.com/wukonggo/OFDConverter)                                         | -               |   ❌    | 2024-10-01 | 仅 PDF→OFD，依赖 XiaoFeng.Ofd               |
| [github.com/BootstrapBlazor/BootstrapBlazor.OfdReader](https://github.com/BootstrapBlazor/BootstrapBlazor.OfdReader) | -               |   ❌    | 2024-09-15 | Blazor OFD 阅读器组件，已归档               |
| [github.com/taurusxin/Ofd2Pdf](https://github.com/taurusxin/Ofd2Pdf)                                                 | 1.2.0           |   ✅    | 2023-10-13 | WPF GUI，拖拽转换 OFD→PDF                   |
| [github.com/lanbo0829/ofd2pdf](https://github.com/lanbo0829/ofd2pdf)                                                 | -               |   ✅    | 2023-08-04 | CLI，基于 Spire.PDF，许可证不明             |

### JavaScript 库

| 库                                                                               | 版本  | OFD→PDF |  更新时间  | 备注                                  |
|----------------------------------------------------------------------------------|-------|:-------:|:----------:|---------------------------------------|
| [@miconvert/ofd-to-pdf](https://www.npmjs.com/package/@miconvert/ofd-to-pdf)     | 0.2.0 |   ✅    | 2026-08-03 | 全部成功                              |
| [github.com/isee15/ofdjs](https://github.com/isee15/ofdjs)                       | -     |   ❌    | 2026-07-14 | OFD 解析渲染到 Canvas                 |
| [github.com/Gary-zy/imageConversion](https://github.com/Gary-zy/imageConversion) | -     |   ✅    | 2026-01-06 | Vue/TS 在线工具，OFD→图/PDF、SM2 验签 |
| [github.com/DLTech21/ofd.js](https://github.com/DLTech21/ofd.js)                 | -     |   ❌    | 2024-06-05 | 纯前端 SVG/Canvas 渲染，342 star      |
| [github.com/SignitDoc/liteofd](https://github.com/SignitDoc/liteofd)             | 0.2.6 |   ❌    | 2025-07-03 | 轻量前端 SDK，解析/文本/签章          |
| [github.com/besthqs/bestofdview](https://github.com/besthqs/bestofdview)         | 0.4.1 |   ❌    | 2025-03-24 | Vue3 组件，渲染/验签，npm 包          |
| [github.com/guinanlin/OFDView](https://github.com/guinanlin/OFDView)             | -     |   ❌    | 2024-10-31 | HTML5/SVG 浏览器阅读器                |
| [github.com/betgo/ofd-online](https://github.com/betgo/ofd-online)               | 1.0.0 |   ❌    | 2024-07-25 | 纯前端解析，返回 SVG DOM              |
| [github.com/xxss0903/ofdviewer](https://github.com/xxss0903/ofdviewer)           | 0.0.4 |   ❌    | 2024-02-29 | VSCode 插件，预览 OFD                 |
| [github.com/zsc347/webofd](https://github.com/zsc347/webofd)                     | -     |   ❌    | 2022-04-21 | 纯 Web Canvas 解析渲染                |

### C++ 库

| 库                                                                         | 版本 | OFD→PDF |  更新时间  | 备注                                          |
|----------------------------------------------------------------------------|------|:-------:|:----------:|-----------------------------------------------|
| [github.com/Micats/ofdReader](https://github.com/Micats/ofdReader)         | -    |   ❌    | 2025-08-07 | Qt 跨平台阅读器，渲染签章                     |
| [github.com/Chingliu/XilouReader](https://github.com/Chingliu/XilouReader) | -    |   ✅    | 2025-01-03 | Qt 阅读器，pdfium 双引擎                      |
| [github.com/Chingliu/xilou_core](https://github.com/Chingliu/xilou_core)   | -    |   ❌    | 2024-10-15 | pdfium 双引擎解析渲染内核                     |
| [github.com/isee15/docwriter](https://github.com/isee15/docwriter)         | -    |   ❌    | 2024-08-19 | 生成 docx/OFD（MIT）                          |
| [github.com/uukuguy/libofd](https://github.com/uukuguy/libofd)             | -    |   ✅    | 2023-11-18 | 首个开源 C++ OFD 库，含查看器，已归档         |
| [github.com/KikyoShaw/OFDEditor](https://github.com/KikyoShaw/OFDEditor)   | -    |   ❌    | 2022-02-24 | Qt OFD 编辑器，可发布 OFD（Apache-2.0）       |
| [github.com/roy19831015/OfdiumEx](https://github.com/roy19831015/OfdiumEx) | -    |   ❌    | 2021-04-14 | cairo 渲染 OFD，接口类似 pdfium（Apache-2.0） |
| [github.com/mcoder2014/ofdEditor](https://github.com/mcoder2014/ofdEditor) | -    |   ❌    | 2019-09-04 | Qt OFD 编辑器，2017 中软杯作品（MIT）         |
| [github.com/CryFeiFei/ZipViewer](https://github.com/CryFeiFei/ZipViewer)   | -    |   ❌    | 2017-12-10 | Qt 查看器，支持 docx/pptx/xlsx/ofd（MIT）     |

## OFD 库汇总

> **预览**：库是否内置 OFD 预览/显示能力（查看器、浏览器渲染或可视化组件）。

| 库                                                                                        |  语言  |    许可证    | 预览 | 解析 | → 图片 | → PDF | → TXT | → SVG | → Md | PDF → OFD | 生成 | 签章 | 修改 |
|-------------------------------------------------------------------------------------------|:------:|:------------:|:----:|:----:|:------:|:-----:|:-----:|:-----:|:----:|:---------:|:----:|:----:|:----:|
| [go-ofdgo](https://github.com/xiaoqidun/ofdgo)                                            |   Go   |  Apache-2.0  |  ✅  |  ✅  |   ✅   |  ✅   |  ✅   |  ✅   |  ❌  |    ❌     |  ❌  |  ❌  |  ❌  |
| [go-zc310](https://github.com/zc310/ofd)                                                  |   Go   |  Unlicense   |  ✅  |  ✅  |   ✅   |  ✅   |  ✅   |  ✅   |  ✅  |    ❌     |  ✅  |  ❌  |  ❌  |
| [easyofd-rust](https://github.com/easy-4-rust/easyofd-rust)                               |  Rust  |  Apache-2.0  |  ❌  |  ✅  |   ✅   |  ✅   |  ❌   |  ❌   |  ✅  |    ❌     |  ✅  |  ✅  |  ✅  |
| [ofdmanager](https://github.com/feuvan/ofdmanager)                                        |  Rust  |     MIT      |  ✅  |  ✅  |   ✅   |  ❌   |  ❌   |  ❌   |  ❌  |    ❌     |  ❌  |  ❌  |  ❌  |
| [rs_ofd](https://github.com/geniusnut/rs_ofd)                                             |  Rust  |     MIT      |  ❌  |  ✅  |   ✅   |  ❌   |  ❌   |  ❌   |  ❌  |    ❌     |  ❌  |  ❌  |  ❌  |
| [ofd-utility](https://github.com/ofd-utility/ofd-utility)                                 |  Rust  |     MIT      |  ❌  |  ✅  |   ✅   |  ❌   |  ❌   |  ❌   |  ❌  |    ❌     |  ❌  |  ✅  |  ❌  |
| [ofdsdk](https://github.com/KaiserY/ofdsdk)                                               |  Rust  |  Apache-2.0  |  ❌  |  ✅  |   ❌   |  ❌   |  ❌   |  ❌   |  ❌  |    ❌     |  ❌  |  ❌  |  ❌  |
| [rofd](https://github.com/linuxdeepin/rofd)                                               |  Rust  |   LGPL-2.1   |  ✅  |  ✅  |   ✅   |  ❌   |  ❌   |  ❌   |  ❌  |    ❌     |  ❌  |  ❌  |  ❌  |
| [easyofd](https://pypi.org/project/easyofd/)                                              | Python |  Apache-2.0  |  ✅  |  ✅  |   ✅   |  ⚠️   |  ❌   |  ❌   |  ❌  |    ✅     |  ✅  |  ❌  |  ❌  |
| [ofd2img](https://pypi.org/project/ofd2img/)                                              | Python |  Apache-2.0  |  ❌  |  ✅  |   ✅   |  ⚠️   |  ❌   |  ❌   |  ❌  |    ❌     |  ❌  |  ❌  |  ❌  |
| [ofd2pdf](https://github.com/jsyzdej/ofd2pdf)                                             | Python |      无      |  ❌  |  ✅  |   ✅   |  ⚠️   |  ❌   |  ❌   |  ❌  |    ❌     |  ❌  |  ❌  |  ❌  |
| [ofdreader](https://pypi.org/project/ofdreader/)                                          | Python |     MIT      |  ❌  |  ✅  |   ❌   |  ✅   |  ✅   |  ❌   |  ❌  |    ❌     |  ✅  |  ❌  |  ❌  |
| [pdf2ofd](https://github.com/wanglrebe/pdf2ofd)                                           | Python |     MIT      |  ❌  |  ❌  |   ❌   |  ❌   |  ❌   |  ❌   |  ❌  |    ✅     |  ❌  |  ❌  |  ❌  |
| [ofd-parser](https://github.com/jyh2012/ofd-parser)                                       | Python |      无      |  ✅  |  ✅  |   ❌   |  ❌   |  ❌   |  ❌   |  ❌  |    ❌     |  ❌  |  ❌  |  ❌  |
| [OfficeMaster](https://github.com/Chingliu/OfficeMaster_document_convert_system)          | Python |     MIT      |  ❌  |  ✅  |   ❌   |  ✅   |  ❌   |  ❌   |  ❌  |    ✅     |  ✅  |  ❌  |  ❌  |
| [ofdrw](https://github.com/ofdrw/ofdrw)                                                   |  Java  |  Apache-2.0  |  ❌  |  ✅  |   ✅   |  ✅   |  ✅   |  ✅   |  ❌  |    ✅     |  ✅  |  ✅  |  ✅  |
| [easyofd-java](https://github.com/11627685/easyofd-java)                                  |  Java  |  Apache-2.0  |  ❌  |  ✅  |   ❌   |  ❌   |  ❌   |  ❌   |  ❌  |    ❌     |  ✅  |  ✅  |  ❌  |
| [ofdbox](https://gitee.com/bookhhu/ofdbox)                                                |  Java  |  Apache-2.0  |  ❌  |  ✅  |   ✅   |  ⚠️   |  ❌   |  ❌   |  ❌  |    ❌     |  ❌  |  ❌  |  ❌  |
| [ofd-analyze](https://github.com/cooker/ofd-analyze)                                      |  Java  |      无      |  ⚠️  |  ✅  |   ✅   |  ❌   |  ❌   |  ❌   |  ❌  |    ❌     |  ❌  |  ❌  |  ❌  |
| [ofdToPdf](https://github.com/thebigboy/ofdToPdf)                                         |  Java  |      无      |  ❌  |  ✅  |   ❌   |  ✅   |  ❌   |  ❌   |  ❌  |    ❌     |  ❌  |  ❌  |  ❌  |
| [JIMU-ConvertPreview](https://github.com/zhangzhen1979/JIMU-ConvertPreview)               |  Java  |  Apache-2.0  |  ✅  |  ✅  |   ✅   |  ✅   |  ❌   |  ❌   |  ❌  |    ✅     |  ✅  |  ❌  |  ⚠️  |
| [ofdbox-viewer](https://github.com/jiayao-zhang/ofdbox-viewer)                            |  Java  |      无      |  ✅  |  ✅  |   ✅   |  ❌   |  ❌   |  ❌   |  ❌  |    ❌     |  ❌  |  ❌  |  ❌  |
| [kkFileView](https://github.com/kekingcn/kkFileView)                                      |  Java  |      无      |  ✅  |  ✅  |   ✅   |  ⚠️   |  ❌   |  ❌   |  ❌  |    ❌     |  ❌  |  ❌  |  ⚠️  |
| [kkfileviewxg](https://github.com/gaoxingzaq/kkfileviewxg)                                |  Java  |      无      |  ✅  |  ✅  |   ✅   |  ✅   |  ❌   |  ✅   |  ❌  |    ❌     |  ❌  |  ❌  |  ⚠️  |
| [ofd-server](https://github.com/maczh/ofd-server)                                         |  Java  |      无      |  ❌  |  ❌  |   ❌   |  ❌   |  ❌   |  ❌   |  ❌  |    ✅     |  ✅  |  ✅  |  ✅  |
| [ofd-parser-tika](https://github.com/ryecrow/ofd-parser)                                  |  Java  |  Apache-2.0  |  ❌  |  ✅  |   ❌   |  ❌   |  ✅   |  ❌   |  ❌  |    ❌     |  ❌  |  ❌  |  ❌  |
| [ofdpdfsigner](https://github.com/fanzizheng/ofdpdfsigner)                                |  C++   |   BSL-1.1    |  ❌  |  ✅  |   ✅   |  ✅   |  ❌   |  ✅   |  ❌  |    ✅     |  ❌  |  ✅  |  ❌  |
| [ofdReader](https://github.com/Micats/ofdReader)                                          |  C++   |      无      |  ✅  |  ✅  |   ✅   |  ❌   |  ❌   |  ❌   |  ❌  |    ❌     |  ❌  |  ⚠️  |  ❌  |
| [XilouReader](https://github.com/Chingliu/XilouReader)                                    |  C++   | BSD-3-Clause |  ✅  |  ✅  |   ✅   |  ✅   |  ❌   |  ❌   |  ❌  |    ✅     |  ✅  |  ✅  |  ✅  |
| [xilou_core](https://github.com/Chingliu/xilou_core)                                      |  C++   | BSD-3-Clause |  ✅  |  ✅  |   ✅   |  ❌   |  ❌   |  ❌   |  ❌  |    ❌     |  ❌  |  ❌  |  ❌  |
| [docwriter](https://github.com/isee15/docwriter)                                          |  C++   |     MIT      |  ❌  |  ❌  |   ❌   |  ❌   |  ❌   |  ❌   |  ❌  |    ❌     |  ✅  |  ❌  |  ❌  |
| [libofd](https://github.com/uukuguy/libofd)                                               |  C++   |     MIT      |  ✅  |  ✅  |   ✅   |  ❌   |  ❌   |  ❌   |  ❌  |    ✅     |  ✅  |  ❌  |  ✅  |
| [OFDEditor](https://github.com/KikyoShaw/OFDEditor)                                       |  C++   |  Apache-2.0  |  ✅  |  ✅  |   ❌   |  ❌   |  ❌   |  ❌   |  ❌  |    ❌     |  ✅  |  ❌  |  ✅  |
| [OfdiumEx](https://github.com/roy19831015/OfdiumEx)                                       |  C++   |  Apache-2.0  |  ✅  |  ✅  |   ✅   |  ❌   |  ❌   |  ❌   |  ❌  |    ❌     |  ❌  |  ❌  |  ❌  |
| [ofdEditor](https://github.com/mcoder2014/ofdEditor)                                      |  C++   |     MIT      |  ✅  |  ✅  |   ✅   |  ❌   |  ❌   |  ❌   |  ❌  |    ❌     |  ✅  |  ❌  |  ✅  |
| [ZipViewer](https://github.com/CryFeiFei/ZipViewer)                                       |  C++   |     MIT      |  ✅  |  ✅  |   ❌   |  ❌   |  ❌   |  ❌   |  ❌  |    ❌     |  ❌  |  ❌  |  ❌  |
| [XiaoFeng.Ofd](https://github.com/zhuovi/XiaoFeng.Ofd)                                    |  .NET  |     MIT      |  ❌  |  ✅  |   ❌   |  ❌   |  ❌   |  ❌   |  ❌  |    ❌     |  ✅  |  ✅  |  ❌  |
| [OfdViewer](https://github.com/LvYueMing/OfdViewer)                                       |  .NET  |     MIT      |  ✅  |  ✅  |   ✅   |  ❌   |  ❌   |  ❌   |  ❌  |    ❌     |  ✅  |  ⚠️  |  ❌  |
| [ofdrw.net](https://github.com/whynpc9/ofdrw.net)                                         |  .NET  |     MIT      |  ❌  |  ✅  |   ✅   |  ✅   |  ✅   |  ✅   |  ❌  |    ✅     |  ✅  |  ⚠️  |  ✅  |
| [ofdrw-net](https://github.com/lllooollpp/ofdrw-net)                                      |  .NET  |      无      |  ❌  |  ✅  |   ✅   |  ✅   |  ❌   |  ❌   |  ❌  |    ✅     |  ✅  |  ✅  |  ❌  |
| [OFDConverter](https://github.com/wukonggo/OFDConverter)                                  |  .NET  |  Apache-2.0  |  ❌  |  ❌  |   ❌   |  ❌   |  ❌   |  ❌   |  ❌  |    ✅     |  ❌  |  ❌  |  ❌  |
| [ofdparser](https://github.com/wangyi160/ofdparser)                                       |  .NET  |      无      |  ❌  |  ✅  |   ❌   |  ❌   |  ❌   |  ❌   |  ❌  |    ❌     |  ❌  |  ❌  |  ❌  |
| [BootstrapBlazor.OfdReader](https://github.com/BootstrapBlazor/BootstrapBlazor.OfdReader) |  .NET  |     MIT      |  ✅  |  ✅  |   ❌   |  ❌   |  ❌   |  ❌   |  ❌  |    ❌     |  ❌  |  ❌  |  ❌  |
| [Ofd2Pdf](https://github.com/taurusxin/Ofd2Pdf)                                           |  .NET  |     MIT      |  ❌  |  ✅  |   ❌   |  ✅   |  ❌   |  ❌   |  ❌  |    ❌     |  ❌  |  ❌  |  ❌  |
| [ofd2pdf](https://github.com/lanbo0829/ofd2pdf)                                           |  .NET  |      无      |  ❌  |  ✅  |   ❌   |  ✅   |  ❌   |  ❌   |  ❌  |    ❌     |  ❌  |  ❌  |  ❌  |
| [ofd-to-pdf](https://www.npmjs.com/package/@miconvert/ofd-to-pdf)                         |   JS   |  Apache-2.0  |  ❌  |  ✅  |   ❌   |  ✅   |  ❌   |  ❌   |  ❌  |    ❌     |  ❌  |  ❌  |  ❌  |
| [ofdjs](https://github.com/isee15/ofdjs)                                                  |   JS   |  Apache-2.0  |  ✅  |  ✅  |   ✅   |  ❌   |  ❌   |  ❌   |  ❌  |    ❌     |  ❌  |  ❌  |  ❌  |
| [imageConversion](https://github.com/Gary-zy/imageConversion)      |   JS   |    MIT     |  ✅  |  ✅  |   ✅   |  ✅   |  ❌   |  ❌   |  ❌  |    ❌     |  ❌  |  ⚠️  |  ❌  |
| [liteofd](https://github.com/SignitDoc/liteofd)                                           |   JS   |  Apache-2.0  |  ✅  |  ✅  |   ✅   |  ❌   |  ✅   |  ❌   |  ❌  |    ❌     |  ❌  |  ⚠️  |  ❌  |
| [bestofdview](https://github.com/besthqs/bestofdview)                                     |   JS   |      无      |  ✅  |  ✅  |   ✅   |  ❌   |  ❌   |  ❌   |  ❌  |    ❌     |  ❌  |  ✅  |  ❌  |
| [OFDView](https://github.com/guinanlin/OFDView)                                           |   JS   |  Apache-2.0  |  ✅  |  ✅  |   ✅   |  ❌   |  ❌   |  ❌   |  ❌  |    ❌     |  ❌  |  ❌  |  ❌  |
| [ofd.js](https://github.com/DLTech21/ofd.js)                                              |   JS   |  Apache-2.0  |  ✅  |  ✅  |   ✅   |  ❌   |  ❌   |  ❌   |  ❌  |    ❌     |  ❌  |  ⚠️  |  ❌  |
| [ofd-online](https://github.com/betgo/ofd-online)                                         |   JS   |      无      |  ✅  |  ✅  |   ✅   |  ❌   |  ❌   |  ✅   |  ❌  |    ❌     |  ❌  |  ❌  |  ❌  |
| [ofdviewer](https://github.com/xxss0903/ofdviewer)                                        |   JS   |     MIT      |  ✅  |  ✅  |   ❌   |  ❌   |  ❌   |  ❌   |  ❌  |    ❌     |  ❌  |  ❌  |  ❌  |
| [webofd](https://github.com/zsc347/webofd)                                                |   JS   |      无      |  ✅  |  ✅  |   ✅   |  ❌   |  ❌   |  ❌   |  ❌  |    ❌     |  ❌  |  ❌  |  ❌  |

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
| [python-ofdreader](https://pypi.org/project/ofdreader/)             |    FAILED |   FAILED |    FAILED |         FAILED |    FAILED |   FAILED |        0 |
| [java-ofdrw](https://github.com/ofdrw/ofdrw)                        |     702ms |    736ms |     5.12s |          2.41s |     811ms |    5.51s |        0 |
| [node-ofd2pdf](https://www.npmjs.com/package/@miconvert/ofd-to-pdf) |     276ms |    703ms | **1.20s** |          8.52s |     506ms |    810ms |        1 |

> - `python-easyofd`: `pip install easyofd ofd2img`，自动切换库，部分文件失败
> - `python-ofd2pdf`: `pip install ofd2pdf`，基于图片渲染，部分文件失败
> - `python-ofdreader`: `pip install ofdreader`，空包（无 Python 代码），全部失败
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
