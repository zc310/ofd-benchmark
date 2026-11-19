import java.awt.image.BufferedImage;
import java.nio.file.Files;
import java.nio.file.Path;
import javax.imageio.ImageIO;
import org.ofdrw.converter.ImageMaker;
import org.ofdrw.reader.OFDReader;

/** Native AWT rendering, not PDF rasterization. */
public class Ofd2Png {
    public static void main(String[] args) throws Exception {
        if (args.length != 2) {
            throw new IllegalArgumentException("Usage: Ofd2Png <input.ofd> <output-directory> (200 DPI)");
        }
        Path output = Path.of(args[1]);
        Files.createDirectories(output);
        try (OFDReader reader = new OFDReader(Path.of(args[0]))) {
            // Both ImageMaker constructors take pixels/mm, NOT DPI.
            ImageMaker maker = new ImageMaker(reader, 200.0 / 25.4);
            for (int i = 0; i < reader.getNumberOfPages(); i++) {
                BufferedImage image = maker.makePage(i);
                if (!ImageIO.write(image, "png", output.resolve("page_" + (i + 1) + ".png").toFile())) {
                    throw new IllegalStateException("No PNG encoder available");
                }
            }
        }
    }
}
