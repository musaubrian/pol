use crate::rpn;

pub fn eval_file(path: &str) -> Result<(), Box<dyn std::error::Error>> {
    let file_contents = std::fs::read_to_string(path)?;
    let solutions_file = format!("{path}.solution");
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

    std::fs::write(&solutions_file, solutions_buf)?;

    Ok(())
}
