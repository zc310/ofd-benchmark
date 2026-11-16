use std::{env, process};

fn main() {
    let args: Vec<_> = env::args_os().collect();
    if args.len() != 3 {
        eprintln!("Usage: easyofd-pdf2ofd <input.pdf> <output.ofd>");
        process::exit(2);
    }
    if let Err(err) = easyofd::pdf_to_ofd(&args[1], &args[2], &easyofd::ConvertOptions::default()) {
        eprintln!("PDF to OFD failed: {err}");
        process::exit(1);
    }
}
