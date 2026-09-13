import java.nio.file.Path;
import java.nio.file.Paths;
import org.ofdrw.converter.export.OFDExporter;
import org.ofdrw.converter.export.PDFExporterPDFBox;

public class Ofd2Pdf {
    public static void main(String[] args) throws Exception {
        if (args.length < 2) {
            System.out.println("Usage: Ofd2Pdf <input.ofd> <output.pdf>");
            System.exit(1);
        }

        Path input = Paths.get(args[0]);
        Path output = Paths.get(args[1]);

        try (OFDExporter exporter = new PDFExporterPDFBox(input, output)) {
            exporter.export();
        }

        System.out.println("Converted: " + input + " -> " + output);
    }
}
