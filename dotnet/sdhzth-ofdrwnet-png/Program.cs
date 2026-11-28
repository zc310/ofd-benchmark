using NetOfd.Converter;
using SkiaSharp;

if (args.Length != 2)
{
    Console.Error.WriteLine("Usage: OfdrwNetPng <input.ofd> <output-directory>");
    return 1;
}

try
{
    Directory.CreateDirectory(args[1]);
    using var maker = new OfdImageMaker(args[0]);
    var options = new ConverterOptions { PPM = 200.0 / 25.4 };
    for (int page = 0; page < maker.PageCount; page++)
    {
        using var bitmap = maker.MakePage(page, options)
            ?? throw new InvalidOperationException($"Page {page + 1} could not be rendered");
        using var image = SKImage.FromBitmap(bitmap);
        using var data = image.Encode(SKEncodedImageFormat.Png, 100)
            ?? throw new InvalidOperationException($"Page {page + 1} could not be encoded");
        using var output = File.Create(Path.Combine(args[1], $"page_{page + 1}.png"));
        data.SaveTo(output);
    }
    return 0;
}
catch (Exception error)
{
    Console.Error.WriteLine(error);
    return 1;
}
