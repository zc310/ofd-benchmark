use std::{env, process};

fn main() {
    let args: Vec<_> = env::args_os().collect();
    if args.len() != 3 {
        eprintln!("Usage: easyofd-ofd2pdf <input.ofd> <output.pdf>");
        process::exit(2);
    }
    if let Err(err) = easyofd::ofd_to_pdf(&args[1], &args[2], &easyofd::ConvertOptions::default()) {
        eprintln!("OFD to PDF failed: {err}");
        process::exit(1);
    }
}
