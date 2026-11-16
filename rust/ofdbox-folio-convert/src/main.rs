use std::{collections::HashMap, env, fs, path::Path, process::Command, sync::Arc};

use ofdbox_folio::{
    ir::{DocumentSource, FixedDocument, FontError, FontProvider, FontRequest},
    ofd::{OfdDocument, OfdReadOptions},
    pdf::{PdfOptions, write_fixed},
};

/// Resolves font families through fontconfig, respecting the runner's configured font dirs.
struct FontconfigFonts {
    cache: HashMap<String, Option<Arc<[u8]>>>,
}

impl FontconfigFonts {
    fn new() -> Result<Self, String> {
        let output = Command::new("fc-match")
            .args(["-f", "%{file}", "sans-serif"])
            .output()
            .map_err(|error| format!("fc-match is required to locate fonts: {error}"))?;
        if !output.status.success() {
            return Err("fc-match could not resolve the sans-serif fallback font".to_owned());
        }
        let path = String::from_utf8_lossy(&output.stdout).trim().to_owned();
        if path.is_empty() || !Path::new(&path).is_file() {
            return Err(format!("fontconfig returned a missing font file: {path}"));
        }
        Ok(Self {
            cache: HashMap::new(),
        })
    }
}

impl FontProvider for FontconfigFonts {
    fn load(&mut self, request: &FontRequest<'_>) -> Result<Option<Arc<[u8]>>, FontError> {
        let family = request
            .family_name
            .or(request.font_name)
            .filter(|name| !name.trim().is_empty())
            .unwrap_or("sans-serif")
            .to_owned();
        let style = match (request.bold, request.italic) {
            (true, true) => "Bold Italic",
            (true, false) => "Bold",
            (false, true) => "Italic",
            (false, false) => "Regular",
        };
        let key = format!("{family}|{style}");
        if let Some(font) = self.cache.get(&key) {
            return Ok(font.clone());
        }

        let query = format!("{family}:style={style}");
        let output = Command::new("fc-match")
            .args(["-f", "%{file}", &query])
            .output()
            .map_err(|error| FontError::new(format!("fc-match {query:?}: {error}")))?;
        if !output.status.success() {
            self.cache.insert(key, None);
            return Ok(None);
        }
        let path = String::from_utf8_lossy(&output.stdout).trim().to_owned();
        let font = fs::read(&path).ok().map(Arc::<[u8]>::from);
        self.cache.insert(key, font.clone());
        Ok(font)
    }
}

fn convert(input: &Path, output: &Path) -> Result<usize, Box<dyn std::error::Error>> {
    let mut source = OfdDocument::open_path_with_fonts(
        input,
        OfdReadOptions::default(),
        Some(Box::new(FontconfigFonts::new()?)),
    )?;
    let mut fixed = FixedDocument::default();
    fixed.metadata = source.info().metadata.clone();
    fixed.outline = source.info().outline.clone();
    let page_count = source.info().page_count;
    for index in 0..page_count {
        fixed.pages.push(source.load_page(index)?);
    }
    fixed.resources = source.resources().clone();
    let pdf = write_fixed(&fixed, &PdfOptions::default())?;
    fs::write(output, pdf)?;
    Ok(page_count)
}

fn main() {
    let args: Vec<_> = env::args_os().collect();
    if args.len() != 3 {
        eprintln!("Usage: ofdbox-folio-convert <input.ofd> <output.pdf>");
        std::process::exit(2);
    }
    match convert(Path::new(&args[1]), Path::new(&args[2])) {
        Ok(pages) => eprintln!("ofdbox-folio: converted {pages} pages"),
        Err(error) => {
            eprintln!("ofdbox-folio: {error}");
            std::process::exit(1);
        }
    }
}
