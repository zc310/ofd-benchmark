import org.ofdrw.converter.ofdconverter.PDFConverter;

import java.nio.file.Path;
import java.nio.file.Paths;

/**
 * PDF -> OFD converter using ofdrw (org.ofdrw:ofdrw-converter).
 *
 * Usage:
 *     java -cp ".:lib/*" Pdf2Ofd <in.pdf> <out.ofd>
 */
public class Pdf2Ofd {
    public static void main(String[] args) throws Exception {
        if (args.length != 2) {
            System.err.println("Usage: Pdf2Ofd <in.pdf> <out.ofd>");
            System.exit(1);
        }
        Path src = Paths.get(args[0]);
        Path dst = Paths.get(args[1]);
        try (PDFConverter converter = new PDFConverter(dst)) {
            converter.setEnableCopyAttachFiles(true);
            converter.setEnableCopyBookmarks(true);
            converter.convert(src);
        }
    }
}