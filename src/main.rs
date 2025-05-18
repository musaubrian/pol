use pol::{file::eval_file, repl};
use std::env::args;

fn main() {
    let mut args = args();
    args.next();

    if let Some(opt) = args.next() {
        match opt.as_str() {
            "-r" | "--repl" => repl::start(),
            "-f" | "--file" => {
                if let Some(file_path) = args.next() {
                    match eval_file(&file_path) {
                        Ok(outpath) => println!("solution generated at {outpath}"),
                        Err(e) => eprintln!("{e}"),
                    };
                } else {
                    eprintln!("[ARG ERR]: Missing filename");
                    usage()
                }
            }
            _ => usage(),
        }
    } else {
        usage();
    }
}

fn usage() {
    let mut args = args();
    if let Some(program) = args.next() {
        eprintln!(
            "
Usage:
    {program} [OPTION]

Options:
    -r, --repl     Start the REPL (Read-Eval-Print Loop)
    -f, --file     Evaluate a file line by line
    -h, --help     Prints this help message

Examples:
    {program} -r
    {program} -f path/to/file

For more information, visit the documentation or run '{program} --help'.
",
        );
    };
}
