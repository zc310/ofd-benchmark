# Java OFD 转换器

基于 [ofdrw](https://github.com/ofdrw/ofdrw) 的 OFD 转 PDF 工具。

## 环境要求

- Java 1.8+

## 依赖下载

在 `lib/` 目录下放置以下 JAR 文件：

### ofdrw 核心库

从 Maven Central 下载：

```bash
# ofdrw 核心模块
curl -sL "https://repo1.maven.org/maven2/org/ofdrw/ofdrw-converter/2.4.0/ofdrw-converter-2.4.0.jar" -o lib/ofdrw-converter.jar
curl -sL "https://repo1.maven.org/maven2/org/ofdrw/ofdrw-reader/2.4.0/ofdrw-reader-2.4.0.jar" -o lib/ofdrw-reader.jar
curl -sL "https://repo1.maven.org/maven2/org/ofdrw/ofdrw-core/2.4.0/ofdrw-core-2.4.0.jar" -o lib/ofdrw-core.jar
curl -sL "https://repo1.maven.org/maven2/org/ofdrw/ofdrw-gv/2.4.0/ofdrw-gv-2.4.0.jar" -o lib/ofdrw-gv.jar
curl -sL "https://repo1.maven.org/maven2/org/ofdrw/ofdrw-layout/2.4.0/ofdrw-layout-2.4.0.jar" -o lib/ofdrw-layout.jar
curl -sL "https://repo1.maven.org/maven2/org/ofdrw/ofdrw-graphics2d/2.4.0/ofdrw-graphics2d-2.4.0.jar" -o lib/ofdrw-graphics2d.jar
curl -sL "https://repo1.maven.org/maven2/org/ofdrw/ofdrw-pkg/2.4.0/ofdrw-pkg-2.4.0.jar" -o lib/ofdrw-pkg.jar
curl -sL "https://repo1.maven.org/maven2/org/ofdrw/ofdrw-font/2.4.0/ofdrw-font-2.4.0.jar" -o lib/ofdrw-font.jar
curl -sL "https://repo1.maven.org/maven2/org/ofdrw/ofdrw-gm/2.4.0/ofdrw-gm-2.4.0.jar" -o lib/ofdrw-gm.jar
```

### 第三方依赖

```bash
# iText PDF
curl -sL "https://repo1.maven.org/maven2/com/itextpdf/commons/7.2.6/commons-7.2.6.jar" -o lib/itext-commons.jar
curl -sL "https://repo1.maven.org/maven2/com/itextpdf/io/7.2.6/io-7.2.6.jar" -o lib/itext-io.jar
curl -sL "https://repo1.maven.org/maven2/com/itextpdf/kernel/7.2.6/kernel-7.2.6.jar" -o lib/itext-kernel.jar
curl -sL "https://repo1.maven.org/maven2/com/itextpdf/layout/7.2.6/layout-7.2.6.jar" -o lib/itext-layout.jar
curl -sL "https://repo1.maven.org/maven2/com/itextpdf/font-asian/7.2.6/font-asian-7.2.6.jar" -o lib/itext-font-asian.jar

# Apache PDFBox
curl -sL "https://repo1.maven.org/maven2/org/apache/pdfbox/pdfbox/2.0.27/pdfbox-2.0.27.jar" -o lib/pdfbox.jar
curl -sL "https://repo1.maven.org/maven2/org/apache/pdfbox/fontbox/2.0.27/fontbox-2.0.27.jar" -o lib/fontbox.jar
curl -sL "https://repo1.maven.org/maven2/org/apache/pdfbox/jbig2-imageio/3.0.0/jbig2-imageio-3.0.0.jar" -o lib/jbig2.jar

# BouncyCastle 加密库
curl -sL "https://repo1.maven.org/maven2/org/bouncycastle/bcprov-jdk18on/1.84/bcprov-jdk18on-1.84.jar" -o lib/bcprov.jar
curl -sL "https://repo1.maven.org/maven2/org/bouncycastle/bcpkix-jdk18on/1.84/bcpkix-jdk18on-1.84.jar" -o lib/bcpkix.jar

# Apache XML Graphics (Batik)
curl -sL "https://repo1.maven.org/maven2/org/apache/xmlgraphics/batik-all/1.19/batik-all-1.19.jar" -o lib/batik.jar

# UJMP 矩阵库
curl -sL "https://repo1.maven.org/maven2/org/ujmp/ujmp-core/0.3.0/ujmp-core-0.3.0.jar" -o lib/ujmp.jar

# XML 解析
curl -sL "https://repo1.maven.org/maven2/org/dom4j/dom4j/2.1.4/dom4j-2.1.4.jar" -o lib/dom4j.jar

# Apache Commons
curl -sL "https://repo1.maven.org/maven2/org/apache/commons/commons-compress/1.26.1/commons-compress-1.26.1.jar" -o lib/compress.jar
curl -sL "https://repo1.maven.org/maven2/commons-io/commons-io/2.16.1/commons-io-2.16.1.jar" -o lib/commons-io.jar
curl -sL "https://repo1.maven.org/maven2/commons-logging/commons-logging/1.3.1/commons-logging-1.3.1.jar" -o lib/logging.jar

# SLF4J 日志
curl -sL "https://repo1.maven.org/maven2/org/slf4j/slf4j-api/2.0.9/slf4j-api-2.0.9.jar" -o lib/slf4j.jar
curl -sL "https://repo1.maven.org/maven2/org/apache/logging/log4j/log4j-slf4j-impl/2.16.0/log4j-slf4j-impl-2.16.0.jar" -o lib/log4j.jar
curl -sL "https://repo1.maven.org/maven2/org/apache/logging/log4j/log4j-core/2.16.0/log4j-core-2.16.0.jar" -o lib/log4j-core.jar
```

## 编译

```bash
# 编译 .java 文件
CP=$(ls lib/*.jar | tr '\n' ':')
javac -cp "$CP" Ofd2Pdf.java
```

## 打包 JAR（可选）

将编译后的 .class 文件打包成可执行 JAR：

```bash
# 创建 MANIFEST.MF
mkdir -p META-INF
echo "Main-Class: Ofd2Pdf" > META-INF/MANIFEST.MF
echo "Class-Path: $(ls lib/*.jar | xargs -I {} echo {} | tr '\n' ' ')" >> META-INF/MANIFEST.MF

# 打包 JAR
jar cfm ofd2pdf.jar META-INF/MANIFEST.MF Ofd2Pdf.class

# 运行 JAR
java -jar ofd2pdf.jar <input.ofd> <output.pdf>
```

## 使用

```bash
# 方式一：通过 shell 脚本
./ofd2pdf_java.sh <input.ofd> <output.pdf>

# 方式二：直接运行（需要指定 classpath）
CP=$(ls lib/*.jar | tr '\n' ':')
java -cp "$CP:." Ofd2Pdf <input.ofd> <output.pdf>

# 方式三：运行打包后的 JAR
java -jar ofd2pdf.jar <input.ofd> <output.pdf>
```
