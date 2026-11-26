use std::{
    env,
    error::Error,
    fs::{self, File},
    path::Path,
};

use cairo::{Context, Format, ImageSurface};
use rofd_core::{Document, LoadOptions};
use rofd_render::{CairoRenderer, ImageDecoder, RenderOptions, SystemFontResolver};

fn convert(input: &Path, output: &Path) -> Result<usize, Box<dyn Error>> {
    let document = Document::open(input, LoadOptions::default())?;
    let count = document.page_count();
    if count == 0 {
        return Err("document has no pages".into());
    }
    let first = document.page(0)?;
    let fonts = SystemFontResolver::with_system_fonts(
        vec!["Noto Sans CJK SC".into(), "Noto Sans Mono".into()],
        first.resource_limits().max_font_bytes,
    );
    let decoder = ImageDecoder::default();
    let options = RenderOptions {
        dpi: 200.0,
        ..Default::default()
    };
    fs::create_dir_all(output)?;
    for index in 0..count {
        let page = document.page(index)?;
        let (width, height) = CairoRenderer::pixel_size(&page, &options)?;
        let surface = ImageSurface::create(Format::ARgb32, width, height)?;
        let context = Context::new(&surface)?;
        let report =
            CairoRenderer.render_page_with_services(&page, &context, &options, &fonts, &decoder)?;
        for diagnostic in report.diagnostics() {
            eprintln!("page {}: {diagnostic:?}", index + 1);
        }
        surface.write_to_png(&mut File::create(
            output.join(format!("page_{}.png", index + 1)),
        )?)?;
    }
    Ok(count)
}

fn main() {
    let args: Vec<_> = env::args_os().collect();
    if args.len() != 3 {
        eprintln!("Usage: linuxdeepin-rofd-png <input.ofd> <output-directory>");
        std::process::exit(2);
    }
    match convert(Path::new(&args[1]), Path::new(&args[2])) {
        Ok(pages) => eprintln!("linuxdeepin-rofd: converted {pages} pages"),
        Err(error) => {
            eprintln!("linuxdeepin-rofd: {error}");
            std::process::exit(1);
        }
    }
}
