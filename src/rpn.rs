pub(crate) use crate::PolErr;

#[derive(Debug, PartialEq)]
pub enum Operation {
    Multiply,
    Addition,
    Subtraction,
    Divide,
    Power,
    Unknown,
}

pub fn eval(content: &str) -> Result<f64, PolErr> {
    let tokens: Vec<&str> = content.trim().split_whitespace().collect();
    let mut pol_stack: Vec<f64> = Vec::new();
    if tokens.is_empty() {
        return Err(PolErr::NotEnoughValues(content.to_string()));
    }

    for token in tokens {
        if is_operator(token) {
            if pol_stack.len() < 2 {
                return Err(PolErr::NotEnoughValues(content.to_string()));
            }

            let operation = map_operation(token);
            if operation == Operation::Unknown {
                return Err(PolErr::InvalidOperator(token.to_string()));
            }

            let second = pol_stack.pop().unwrap();
            let first = pol_stack.pop().unwrap();

            let res = calc(first, second, operation)?;
            pol_stack.push(res);
            continue;
        }

        let num = token.parse::<f64>()?;
        pol_stack.push(num);
    }

    let res = pol_stack.pop().unwrap();
    Ok(res)
}

fn calc(a: f64, b: f64, op: Operation) -> Result<f64, PolErr> {
    match op {
        Operation::Multiply => Ok(a * b),
        Operation::Addition => Ok(a + b),
        Operation::Subtraction => Ok(a - b),
        Operation::Power => Ok(a.powf(b)),
        Operation::Unknown => panic!("State should not be possible"),
        Operation::Divide => {
            if b == 0.0 {
                return Err(PolErr::DivisionByZero);
            }
            Ok(a / b)
        }
    }
}
fn is_operator(token: &str) -> bool {
    match token {
        "+" | "-" | "*" | "/" | "^" => true,
        _ => false,
    }
}

fn map_operation(v: &str) -> Operation {
    match v {
        "+" => Operation::Addition,
        "-" => Operation::Subtraction,
        "*" => Operation::Multiply,
        "/" => Operation::Divide,
        "^" => Operation::Power,
        _ => Operation::Unknown,
    }
}

#[cfg(test)]
mod test {
    use crate::rpn::{PolErr, eval};

    #[test]
    fn empty_token() {
        let res = eval("");
        assert!(res.is_err());
        assert_eq!(res.unwrap_err(), PolErr::NotEnoughValues(String::from("")));
    }

    #[test]
    fn simple_operations() {
        let test_cases = vec![
            ("3 4 5 + -", -6.0),
            ("3 4 +", 7.0),
            ("10 2 -", 8.0),
            ("5 6 *", 30.0),
            ("8 2 /", 4.0),
            ("2 3 ^", 8.0),
            ("-3 -4 +", -7.0),
            ("-10 2 -", -12.0),
            ("-5 -6 *", 30.0),
            ("-8 -2 /", 4.0),
            ("-2 -3 ^", -0.125),
            ("1000000 2000000 +", 3000000.0),
            ("1000000 2000000 *", 2000000000000.0),
        ];

        for (input, expected) in test_cases {
            let result = eval(input);
            let def = eval(input);
            assert!(&result.is_ok());
            assert!(
                (result.unwrap() - expected).abs() < f64::EPSILON,
                "\nFailed for input: [{}]\n\texpected <{}>, got <{}>",
                input,
                expected,
                def.unwrap()
            );
        }
    }

    #[test]
    fn mixed_operations() {
        let test_cases = vec![
            ("3 4 + 2 *", 14.0),
            ("5 1 + 4 + 3 -", 7.0),
            ("2 4 * 9 +", 17.0),
            ("3 4 + 2 + 1 *", 9.0),
        ];

        for (input, expected) in test_cases {
            let result = eval(input);
            let def = eval(input);
            assert!(result.is_ok());
            assert!(
                (result.unwrap() - expected).abs() < f64::EPSILON,
                "\nFailed for input: [{}]\n\texpected <{}>, got <{}>",
                input,
                expected,
                def.unwrap()
            );
        }
    }

    #[test]
    fn div_by_zero() {
        let test_cases = vec![
            ("1 0 /", PolErr::DivisionByZero),
            ("0 0 /", PolErr::DivisionByZero),
        ];

        for (input, _expected) in test_cases {
            let result = eval(input);
            assert!(result.is_err());
        }
    }

    #[test]
    fn complex_operations() {
        let test_cases = vec![
            ("3 4 + 2 - 5 * 6 -", 19.0),
            (" 2 3 ^ 4  - 5 ^", 1024.0),
            ("15 7 1 1 + - / 3 * 2 1 1 + + - ", 5.0),
            ("3 4 2 * 1 5 - 2 3 ^ * / +", 2.75),
            ("2 3 ^ 3 2 ^ + 4 / 5 1 - * 6 +", 23.0),
        ];

        for (input, expected) in test_cases {
            let result = eval(input);
            let def = eval(input);
            assert!(result.is_ok());
            assert!(
                (result.unwrap() - expected).abs() < f64::EPSILON,
                "\nFailed for input: [{}]\n\texpected <{}>, got <{}>",
                input,
                expected,
                def.unwrap()
            );
        }
    }
}
