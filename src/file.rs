use std::time::{SystemTime, UNIX_EPOCH};

use crate::{PolErr, rpn};

pub fn eval_file(path: &str) -> Result<String, PolErr> {
    let file_contents = match std::fs::read_to_string(path) {
        Ok(contents) => contents,
        Err(e) => return Err(PolErr::FileParseErr(e.to_string())),
    };

    let now = SystemTime::now();
    let since_epoch = now.duration_since(UNIX_EPOCH)?.as_millis();

    let solutions_file = format!("{since_epoch}_{path}");

    let mut solutions_buf = String::new();

    for line in file_contents.split('\n') {
        // skip empty lines or comments
        if line.is_empty() || line.contains("#") {
            continue;
        }

        let result = rpn::eval(line)?;
        let string = format!("{line}\n := {result}\n\n");
        solutions_buf.push_str(&string);
    }

    match std::fs::write(&solutions_file, solutions_buf) {
        Ok(_) => Ok(solutions_file),
        Err(err) => return Err(PolErr::FileParseErr(err.to_string())),
    }
}
