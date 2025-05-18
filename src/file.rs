use crate::{PolErr, rpn};

pub fn eval_file(path: &str) -> Result<String, PolErr> {
    let file_contents = match std::fs::read_to_string(path) {
        Ok(contents) => contents,
        Err(e) => return Err(PolErr::FileParseErr(e.to_string())),
    };

    let id = gen_random_id(5);
    let solutions_file = format!("{path}_{id}");

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

fn gen_random_id(word_len: u8) -> String {
    let _ = word_len;
    // wl := 3
    // cs := "BCDFGHJKLMNPQRSTVWXZY"
    // vs := "AEIOU"
    // var result string
    //
    // for i := 0; i < wl; i++ {
    // 	result += string(cs[rand.Intn(len(cs))])
    // 	result += string(vs[rand.Intn(len(vs))])
    // }
    let chars = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789";
    let mut rand_id = String::new();

    rand_id
}
