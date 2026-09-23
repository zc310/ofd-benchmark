# OFD 库汇总

> **预览**：库是否内置 OFD 预览/显示能力（查看器、浏览器渲染或可视化组件）。

## Go

| 库                                                      |   许可证   | 预览 | 解析 | → 图片 | → PDF | → TXT | → SVG | → Md | PDF → OFD | 生成 | 签章 | 修改 |
|---------------------------------------------------------|:----------:|:----:|:----:|:------:|:-----:|:-----:|:-----:|:----:|:---------:|:----:|:----:|:----:|
| [go-ofdgo](https://github.com/xiaoqidun/ofdgo)          | Apache-2.0 |  ✅  |  ✅  |   ✅   |  ✅   |  ✅   |  ✅   |  ❌  |    ❌     |  ✅  |  ❌  |  ✅  |
| [go-zc310](https://github.com/zc310/ofd)                | Apache-2.0 |  ✅  |  ✅  |   ✅   |  ✅   |  ✅   |  ✅   |  ✅  |    ❌     |  ✅  |  ✅  |  ✅  |
| [ofd-go](https://github.com/itlabers/ofd-go)            | Apache-2.0 |  ❌  |  ✅  |   ❌   |  ❌   |  ❌   |  ❌   |  ❌  |    ❌     |  ❌  |  ✅  |  ❌  |
| [go-ofd](https://github.com/ppxz2014/go-ofd)            | Apache-2.0 |  ❌  |  ✅  |   ❌   |  ❌   |  ❌   |  ❌   |  ❌  |    ❌     |  ❌  |  ✅  |  ❌  |
| [leijacob-ofd](https://gitee.com/leijacob/ofd)          |     无     |  ❌  |  ✅  |   ❌   |  ❌   |  ❌   |  ❌   |  ❌  |    ✅     |  ✅  |  ✅  |  ✅  |
| [signer-tools](https://gitee.com/leijacob/signer-tools) |     无     |  ❌  |  ✅  |   ❌   |  ❌   |  ❌   |  ❌   |  ❌  |    ❌     |  ❌  |  ✅  |  ❌  |

## Rust

| 库                                                          |   许可证   | 预览 | 解析 | → 图片 | → PDF | → TXT | → SVG | → Md | PDF → OFD | 生成 | 签章 | 修改 |
|-------------------------------------------------------------|:----------:|:----:|:----:|:------:|:-----:|:-----:|:-----:|:----:|:---------:|:----:|:----:|:----:|
| [easyofd-rust](https://github.com/easy-4-rust/easyofd-rust) | Apache-2.0 |  ❌  |  ✅  |   ✅   |  ✅   |  ❌   |  ❌   |  ✅  |    ⚠️     |  ✅  |  ✅  |  ✅  |
| [ofdmanager](https://github.com/feuvan/ofdmanager)          |    MIT     |  ✅  |  ✅  |   ✅   |  ❌   |  ❌   |  ❌   |  ❌  |    ❌     |  ❌  |  ❌  |  ❌  |
| [rs_ofd](https://github.com/geniusnut/rs_ofd)               |    MIT     |  ❌  |  ✅  |   ✅   |  ❌   |  ❌   |  ❌   |  ❌  |    ❌     |  ❌  |  ❌  |  ❌  |
| [ofd-utility](https://github.com/ofd-utility/ofd-utility)   |    MIT     |  ❌  |  ✅  |   ✅   |  ❌   |  ❌   |  ❌   |  ❌  |    ❌     |  ❌  |  ✅  |  ❌  |
| [ofdsdk](https://github.com/KaiserY/ofdsdk)                 | Apache-2.0 |  ❌  |  ✅  |   ❌   |  ❌   |  ❌   |  ❌   |  ❌  |    ❌     |  ❌  |  ❌  |  ❌  |
| [rofd](https://github.com/linuxdeepin/rofd)                 |  LGPL-2.1  |  ✅  |  ✅  |   ✅   |  ❌   |  ❌   |  ❌   |  ❌  |    ❌     |  ❌  |  ❌  |  ❌  |
| [ofd-viewer](https://github.com/shaoyi1998/ofd-viewer)      |    MIT     |  ✅  |  ✅  |   ❌   |  ✅   |  ✅   |  ❌   |  ❌  |    ❌     |  ❌  |  ❌  |  ❌  |
| [invoice-engine](https://github.com/erma0/fapiao-print)     |    MIT     |  ✅  |  ✅  |   ⚠️   |  ❌   |  ❌   |  ✅   |  ❌  |    ❌     |  ❌  |  ❌  |  ❌  |

## Python

| 库                                                                               |   许可证   | 预览 | 解析 | → 图片 | → PDF | → TXT | → SVG | → Md | PDF → OFD | 生成 | 签章 | 修改 |
|----------------------------------------------------------------------------------|:----------:|:----:|:----:|:------:|:-----:|:-----:|:-----:|:----:|:---------:|:----:|:----:|:----:|
| [easyofd](https://pypi.org/project/easyofd/)                                     | Apache-2.0 |  ✅  |  ✅  |   ✅   |  ⚠️   |  ❌   |  ❌   |  ❌  |    ✅     |  ✅  |  ❌  |  ❌  |
| [ofd2img](https://pypi.org/project/ofd2img/)                                     | Apache-2.0 |  ❌  |  ✅  |   ✅   |  ⚠️   |  ❌   |  ❌   |  ❌  |    ❌     |  ❌  |  ❌  |  ❌  |
| [ofd2pdf](https://github.com/jsyzdej/ofd2pdf)                                    |     无     |  ❌  |  ✅  |   ✅   |  ⚠️   |  ❌   |  ❌   |  ❌  |    ❌     |  ❌  |  ❌  |  ❌  |
| [OFDtoPDF](https://github.com/njuzzy1979/OFDtoPDF)                               |     无     |  ❌  |  ✅  |   ❌   |  ✅   |  ❌   |  ❌   |  ❌  |    ❌     |  ❌  |  ⚠️  |  ❌  |
| [ofdreader](https://pypi.org/project/ofdreader/)                                 |    MIT     |  ❌  |  ✅  |   ❌   |  ✅   |  ✅   |  ❌   |  ❌  |    ❌     |  ✅  |  ❌  |  ❌  |
| [pdf2ofd](https://github.com/wanglrebe/pdf2ofd)                                  |    MIT     |  ❌  |  ❌  |   ❌   |  ❌   |  ❌   |  ❌   |  ❌  |    ✅     |  ❌  |  ❌  |  ❌  |
| [ofd-parser](https://github.com/jyh2012/ofd-parser)                              |     无     |  ✅  |  ✅  |   ❌   |  ❌   |  ❌   |  ❌   |  ❌  |    ❌     |  ❌  |  ❌  |  ❌  |
| [OfficeMaster](https://github.com/Chingliu/OfficeMaster_document_convert_system) |    MIT     |  ❌  |  ✅  |   ❌   |  ✅   |  ❌   |  ❌   |  ❌  |    ✅     |  ✅  |  ❌  |  ❌  |

## Java

| 库                                                                                        |    许可证    | 预览 | 解析 | → 图片 | → PDF | → TXT | → SVG | → Md | PDF → OFD | 生成 | 签章 | 修改 |
|-------------------------------------------------------------------------------------------|:------------:|:----:|:----:|:------:|:-----:|:-----:|:-----:|:----:|:---------:|:----:|:----:|:----:|
| [ofdrw](https://github.com/ofdrw/ofdrw)                                                   |  Apache-2.0  |  ❌  |  ✅  |   ✅   |  ✅   |  ✅   |  ✅   |  ❌  |    ✅     |  ✅  |  ✅  |  ✅  |
| [easyofd-java](https://github.com/11627685/easyofd-java)                                  |  Apache-2.0  |  ❌  |  ✅  |   ❌   |  ❌   |  ❌   |  ❌   |  ❌  |    ❌     |  ✅  |  ✅  |  ❌  |
| [ofdbox](https://gitee.com/bookhhu/ofdbox)                                                |  Apache-2.0  |  ❌  |  ✅  |   ✅   |  ⚠️   |  ❌   |  ❌   |  ❌  |    ❌     |  ❌  |  ❌  |  ❌  |
| [ofd-analyze](https://github.com/cooker/ofd-analyze)                                      |      无      |  ⚠️  |  ✅  |   ✅   |  ❌   |  ❌   |  ❌   |  ❌  |    ❌     |  ❌  |  ❌  |  ❌  |
| [ofdToPdf](https://github.com/thebigboy/ofdToPdf)                                         |      无      |  ❌  |  ✅  |   ❌   |  ✅   |  ❌   |  ❌   |  ❌  |    ❌     |  ❌  |  ❌  |  ❌  |
| [JIMU-ConvertPreview](https://github.com/zhangzhen1979/JIMU-ConvertPreview)               |  Apache-2.0  |  ✅  |  ✅  |   ✅   |  ✅   |  ❌   |  ❌   |  ❌  |    ✅     |  ✅  |  ❌  |  ⚠️  |
| [ofdbox-viewer](https://github.com/jiayao-zhang/ofdbox-viewer)                            |      无      |  ✅  |  ✅  |   ✅   |  ❌   |  ❌   |  ❌   |  ❌  |    ❌     |  ❌  |  ❌  |  ❌  |
| [kkFileView](https://github.com/kekingcn/kkFileView)                                      |      无      |  ✅  |  ✅  |   ✅   |  ⚠️   |  ❌   |  ❌   |  ❌  |    ❌     |  ❌  |  ❌  |  ⚠️  |
| [kkfileviewxg](https://github.com/gaoxingzaq/kkfileviewxg)                                |      无      |  ✅  |  ✅  |   ✅   |  ✅   |  ❌   |  ✅   |  ❌  |    ❌     |  ❌  |  ❌  |  ⚠️  |
| [ofd-server](https://github.com/maczh/ofd-server)                                         |      无      |  ❌  |  ❌  |   ❌   |  ❌   |  ❌   |  ❌   |  ❌  |    ✅     |  ✅  |  ✅  |  ✅  |
| [ofd-parser-tika](https://github.com/ryecrow/ofd-parser)                                  |  Apache-2.0  |  ❌  |  ✅  |   ❌   |  ❌   |  ✅   |  ❌   |  ❌  |    ❌     |  ❌  |  ❌  |  ❌  |

## C++

| 库                                                                                        |    许可证    | 预览 | 解析 | → 图片 | → PDF | → TXT | → SVG | → Md | PDF → OFD | 生成 | 签章 | 修改 |
|-------------------------------------------------------------------------------------------|:------------:|:----:|:----:|:------:|:-----:|:-----:|:-----:|:----:|:---------:|:----:|:----:|:----:|
| [ofdpdfsigner](https://github.com/fanzizheng/ofdpdfsigner)                                |   BSL-1.1    |  ❌  |  ✅  |   ✅   |  ✅   |  ❌   |  ✅   |  ❌  |    ✅     |  ❌  |  ✅  |  ❌  |
| [ofdReader](https://github.com/Micats/ofdReader)                                          |      无      |  ✅  |  ✅  |   ✅   |  ❌   |  ❌   |  ❌   |  ❌  |    ❌     |  ❌  |  ⚠️  |  ❌  |
| [XilouReader](https://github.com/Chingliu/XilouReader)                                    | BSD-3-Clause |  ✅  |  ✅  |   ✅   |  ✅   |  ❌   |  ❌   |  ❌  |    ✅     |  ✅  |  ✅  |  ✅  |
| [xilou_core](https://github.com/Chingliu/xilou_core)                                      | BSD-3-Clause |  ✅  |  ✅  |   ✅   |  ❌   |  ❌   |  ❌   |  ❌  |    ❌     |  ❌  |  ❌  |  ❌  |
| [docwriter](https://github.com/isee15/docwriter)                                          |     MIT      |  ❌  |  ❌  |   ❌   |  ❌   |  ❌   |  ❌   |  ❌  |    ❌     |  ✅  |  ❌  |  ❌  |
| [libofd](https://github.com/uukuguy/libofd)                                               |     MIT      |  ✅  |  ✅  |   ✅   |  ❌   |  ❌   |  ❌   |  ❌  |    ✅     |  ✅  |  ❌  |  ✅  |
| [OFDEditor](https://github.com/KikyoShaw/OFDEditor)                                       |  Apache-2.0  |  ✅  |  ✅  |   ❌   |  ❌   |  ❌   |  ❌   |  ❌  |    ❌     |  ✅  |  ❌  |  ✅  |
| [OfdiumEx](https://github.com/roy19831015/OfdiumEx)                                       |  Apache-2.0  |  ✅  |  ✅  |   ✅   |  ❌   |  ❌   |  ❌   |  ❌  |    ❌     |  ❌  |  ❌  |  ❌  |
| [ofdEditor](https://github.com/mcoder2014/ofdEditor)                                      |     MIT      |  ✅  |  ✅  |   ✅   |  ❌   |  ❌   |  ❌   |  ❌  |    ❌     |  ✅  |  ❌  |  ✅  |
| [ZipViewer](https://github.com/CryFeiFei/ZipViewer)                                       |     MIT      |  ✅  |  ✅  |   ❌   |  ❌   |  ❌   |  ❌   |  ❌  |    ❌     |  ❌  |  ❌  |  ❌  |

## .NET

| 库                                                                                        |    许可证    | 预览 | 解析 | → 图片 | → PDF | → TXT | → SVG | → Md | PDF → OFD | 生成 | 签章 | 修改 |
|-------------------------------------------------------------------------------------------|:------------:|:----:|:----:|:------:|:-----:|:-----:|:-----:|:----:|:---------:|:----:|:----:|:----:|
| [XiaoFeng.Ofd](https://github.com/zhuovi/XiaoFeng.Ofd)                                    |     MIT      |  ❌  |  ✅  |   ❌   |  ❌   |  ❌   |  ❌   |  ❌  |    ❌     |  ✅  |  ✅  |  ❌  |
| [OfdViewer](https://github.com/LvYueMing/OfdViewer)                                       |     MIT      |  ✅  |  ✅  |   ✅   |  ❌   |  ❌   |  ❌   |  ❌  |    ❌     |  ✅  |  ⚠️  |  ❌  |
| [ofdrw.net](https://github.com/whynpc9/ofdrw.net)                                         |     MIT      |  ❌  |  ✅  |   ✅   |  ✅   |  ✅   |  ✅   |  ❌  |    ✅     |  ✅  |  ⚠️  |  ✅  |
| [ofdrw-net](https://github.com/lllooollpp/ofdrw-net)                                      |      无      |  ❌  |  ✅  |   ✅   |  ✅   |  ❌   |  ❌   |  ❌  |    ✅     |  ✅  |  ✅  |  ❌  |
| [OFDConverter](https://github.com/wukonggo/OFDConverter)                                  |  Apache-2.0  |  ❌  |  ❌  |   ❌   |  ❌   |  ❌   |  ❌   |  ❌  |    ✅     |  ❌  |  ❌  |  ❌  |
| [ofdparser](https://github.com/wangyi160/ofdparser)                                       |      无      |  ❌  |  ✅  |   ❌   |  ❌   |  ❌   |  ❌   |  ❌  |    ❌     |  ❌  |  ❌  |  ❌  |
| [BootstrapBlazor.OfdReader](https://github.com/BootstrapBlazor/BootstrapBlazor.OfdReader) |     MIT      |  ✅  |  ✅  |   ❌   |  ❌   |  ❌   |  ❌   |  ❌  |    ❌     |  ❌  |  ❌  |  ❌  |
| [Ofd2Pdf](https://github.com/taurusxin/Ofd2Pdf)                                           |     MIT      |  ❌  |  ✅  |   ❌   |  ✅   |  ❌   |  ❌   |  ❌  |    ❌     |  ❌  |  ❌  |  ❌  |
| [ofd2pdf](https://github.com/lanbo0829/ofd2pdf)                                           |      无      |  ❌  |  ✅  |   ❌   |  ✅   |  ❌   |  ❌   |  ❌  |    ❌     |  ❌  |  ❌  |  ❌  |

## JavaScript

| 库                                                                                        |    许可证    | 预览 | 解析 | → 图片 | → PDF | → TXT | → SVG | → Md | PDF → OFD | 生成 | 签章 | 修改 |
|-------------------------------------------------------------------------------------------|:------------:|:----:|:----:|:------:|:-----:|:-----:|:-----:|:----:|:---------:|:----:|:----:|:----:|
| [ofd-to-pdf](https://www.npmjs.com/package/@miconvert/ofd-to-pdf)                         |  Apache-2.0  |  ❌  |  ✅  |   ❌   |  ✅   |  ❌   |  ❌   |  ❌  |    ❌     |  ❌  |  ❌  |  ❌  |
| [ofdjs](https://github.com/isee15/ofdjs)                                                  |  Apache-2.0  |  ✅  |  ✅  |   ✅   |  ❌   |  ❌   |  ❌   |  ❌  |    ❌     |  ❌  |  ❌  |  ❌  |
| [imageConversion](https://github.com/Gary-zy/imageConversion)                             |     MIT      |  ✅  |  ✅  |   ✅   |  ✅   |  ❌   |  ❌   |  ❌  |    ❌     |  ❌  |  ⚠️  |  ❌  |
| [liteofd](https://github.com/SignitDoc/liteofd)                                           |  Apache-2.0  |  ✅  |  ✅  |   ✅   |  ❌   |  ✅   |  ❌   |  ❌  |    ❌     |  ❌  |  ⚠️  |  ❌  |
| [bestofdview](https://github.com/besthqs/bestofdview)                                     |      无      |  ✅  |  ✅  |   ✅   |  ❌   |  ❌   |  ❌   |  ❌  |    ❌     |  ❌  |  ✅  |  ❌  |
| [OFDView](https://github.com/guinanlin/OFDView)                                           |  Apache-2.0  |  ✅  |  ✅  |   ✅   |  ❌   |  ❌   |  ❌   |  ❌  |    ❌     |  ❌  |  ❌  |  ❌  |
| [ofd.js](https://github.com/DLTech21/ofd.js)                                              |  Apache-2.0  |  ✅  |  ✅  |   ✅   |  ❌   |  ❌   |  ❌   |  ❌  |    ❌     |  ❌  |  ⚠️  |  ❌  |
| [ofd-online](https://github.com/betgo/ofd-online)                                         |      无      |  ✅  |  ✅  |   ✅   |  ❌   |  ❌   |  ✅   |  ❌  |    ❌     |  ❌  |  ❌  |  ❌  |
| [ofdviewer](https://github.com/xxss0903/ofdviewer)                                        |     MIT      |  ✅  |  ✅  |   ❌   |  ❌   |  ❌   |  ❌   |  ❌  |    ❌     |  ❌  |  ❌  |  ❌  |
| [webofd](https://github.com/zsc347/webofd)                                                |      无      |  ✅  |  ✅  |   ✅   |  ❌   |  ❌   |  ❌   |  ❌  |    ❌     |  ❌  |  ❌  |  ❌  |
| [ofd_editor_frontend](https://github.com/LamplightShadow/ofd_editor_frontend)             |      无      |  ✅  |  ✅  |   ✅   |  ✅   |  ❌   |  ❌   |  ❌  |    ❌     |  ✅  |  ✅  |  ✅  |
| [ofdjs-viewer](https://github.com/Atw-Lee/ofdjs-viewer)                                   |     MIT      |  ✅  |  ✅  |   ✅   |  ❌   |  ✅   |  ❌   |  ❌  |    ❌     |  ❌  |  ❌  |  ❌  |
| [jsOFD](https://github.com/Hufe921/jsOFD)                                                 |     MIT      |  ❌  |  ❌  |   ❌   |  ❌   |  ❌   |  ❌   |  ❌  |    ✅     |  ✅  |  ❌  |  ❌  |

## Pascal

| 库                                                                                        |    许可证    | 预览 | 解析 | → 图片 | → PDF | → TXT | → SVG | → Md | PDF → OFD | 生成 | 签章 | 修改 |
|-------------------------------------------------------------------------------------------|:------------:|:----:|:----:|:------:|:-----:|:-----:|:-----:|:----:|:---------:|:----:|:----:|:----:|
| [tinyofd](https://github.com/miemiekurisu/tinyofd)                                        | PolyForm-NC  |  ✅  |  ✅  |   ❌   |  ❌   |  ✅   |  ❌   |  ❌  |    ❌     |  ❌  |  ❌  |  ❌  |
