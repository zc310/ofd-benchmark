# Java OFD 转换器

基于 [ofdrw](https://github.com/ofdrw/ofdrw) 的 OFD→PDF 和 PDF→OFD 转换工具。

## 环境要求

- JDK（提供 `java` 和 `javac`）
- Maven 3

以下命令均在 `java/` 目录下执行。

## 下载依赖并编译

`pom.xml` 管理 ofdrw、zip4j 及其传递依赖。使用 Maven 将依赖复制到脚本使用的 `lib/` 目录，再编译两个转换器：

```bash
mvn -B dependency:copy-dependencies -DoutputDirectory="$PWD/lib"
javac -cp 'lib/*' Ofd2Pdf.java Pdf2Ofd.java
```

建议使用干净的 `lib/` 目录：Maven 不会自动清理旧版本或手动下载的 JAR，残留文件可能导致编译或运行出错。已有文件请在确认后另行移走。

默认 ofdrw 版本由 `pom.xml` 中的 `ofdrw.version` 指定；需要测试其他版本时可通过 `-Dofdrw.version=版本号` 覆盖。

## 使用

```bash
# OFD → PDF
./ofd2pdf_java.sh <input.ofd> <output.pdf>

# PDF → OFD
./pdf2ofd_java.sh <input.pdf> <output.ofd>
```

也可以直接运行（Linux/macOS）：

```bash
java -cp '.:lib/*' Ofd2Pdf <input.ofd> <output.pdf>
java -cp '.:lib/*' Pdf2Ofd <input.pdf> <output.ofd>
```

脚本从自身所在目录读取 `.class` 文件和 `lib/` 下的依赖。若用于基准测试，请将编译后的 `.class`、脚本和 `lib/` 部署到 `bin/java/`。
