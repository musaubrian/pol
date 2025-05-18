use crate::rpn;
use std::io::Write;

#[derive(Debug)]
enum ColorCode {
    BLUE,
    RED,
    DIM,
}

fn colorize(code: ColorCode, content: String) -> String {
    match code {
        ColorCode::BLUE => format!("\x1b[94m{content}\x1b[0m"),
        ColorCode::RED => format!("\x1b[91m{content}\x1b[0m"),
        ColorCode::DIM => format!("\x1b[2m{content}\x1b[0m"),
    }
}

pub fn start() {
    let mut input_buf = String::new();
    let stdin = std::io::stdin();
    let mut stdout = std::io::stdout();
    println!("\nRPN REPL\n---");
    loop {
        print!("{}", colorize(ColorCode::BLUE, ">> ".to_string()));
        stdout.flush().unwrap();
        input_buf.clear(); // clear the buff for the next iteration

        if stdin.read_line(&mut input_buf).is_err() {
            eprintln!(
                "{}",
                colorize(
                    ColorCode::RED,
                    "[REPL ERR]: failed to get input".to_string()
                )
            );
        }

        input_buf = input_buf.trim().to_string();
        if input_buf == ".exit" {
            println!("{}", colorize(ColorCode::DIM, "\\. Bye".to_string()));
            break;
        }
        if input_buf.is_empty() {
            continue;
        }

        match rpn::eval(&input_buf) {
            Ok(result) => {
                println!("{} {result}", colorize(ColorCode::DIM, ":".to_string()));
            }
            Err(e) => match e {
                rpn::PolErr::ParseError(e) => {
                    eprintln!("{}", colorize(ColorCode::RED, e.to_string()))
                }
                rpn::PolErr::InvalidOperator(op) => eprintln!(
                    "{}",
                    colorize(ColorCode::RED, format!("Operator <{op}> not supported"))
                ),
                rpn::PolErr::NotEnoughValues(p) => eprintln!(
                    "{p}\n {}",
                    colorize(
                        ColorCode::RED,
                        "Expected at least 2 values got 1".to_string()
                    )
                ),
                rpn::PolErr::DivisionByZero => eprintln!(
                    "{}",
                    colorize(ColorCode::RED, "Division By Zero".to_string())
                ),
                rpn::PolErr::FileParseErr(_) => eprintln!(
                    "{}",
                    colorize(ColorCode::RED, "repl can't handle files yet".to_string())
                ),
            },
        }
    }
}
