use std::{env, fs, path::Path};

use easyofd_convert::converter::AWTMaker;

fn convert(input: &Path, output: &Path) -> Result<(), Box<dyn std::error::Error>> {
    fs::create_dir_all(output)?;
    AWTMaker::from_file(input)?
        .with_dpi(200.0)
        .with_background_color(0xffffff)
        .convert(&output.join("render.png"))?;
    if output.join("render.png").is_file() {
        fs::rename(output.join("render.png"), output.join("page_1.png"))?;
    } else {
        let count = fs::read_dir(output)?.count();
        if count == 0 {
            return Err("no PNG pages produced".into());
        }
        for index in 0..count {
            fs::rename(
                output.join(format!("render_{index}.png")),
                output.join(format!("page_{}.png", index + 1)),
            )?;
        }
    }
    Ok(())
}

fn main() {
    let args: Vec<_> = env::args_os().collect();
    if args.len() != 3 {
        eprintln!("Usage: easyofd-ofd2png <input.ofd> <output-directory> (200 DPI)");
        std::process::exit(2);
    }
    if let Err(error) = convert(Path::new(&args[1]), Path::new(&args[2])) {
        eprintln!("easyofd-rust PNG: {error}");
        std::process::exit(1);
    }
}
