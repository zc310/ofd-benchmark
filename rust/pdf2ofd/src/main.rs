use std::env;
use std::process;
use easyofd::ConvertOptions;

fn main() {
    let args: Vec<String> = env::args().collect();

    if args.len() < 3 {
        eprintln!("Usage: pdf2ofd <input.pdf> <output.ofd>");
        process::exit(1);
    }

    let input = &args[1];
    let output = &args[2];

    let options = ConvertOptions::default();

    match easyofd::pdf_to_ofd(input, output, &options) {
        Ok(()) => {
            println!("Converted: {} -> {}", input, output);
        }
        Err(e) => {
            eprintln!("Error: {}", e);
            process::exit(1);
        }
    }
}