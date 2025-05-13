use std::{fmt, num::ParseFloatError};

pub mod file;
pub mod repl;
pub mod rpn;

#[derive(Debug, PartialEq)]
pub enum PolErr {
    ParseError(ParseFloatError),
    InvalidOperator(String),
    NotEnoughValues(String),
    DivisionByZero,
}

impl From<ParseFloatError> for PolErr {
    fn from(err: ParseFloatError) -> Self {
        PolErr::ParseError(err)
    }
}

impl fmt::Display for PolErr {
    fn fmt(&self, f: &mut fmt::Formatter<'_>) -> fmt::Result {
        match self {
            PolErr::ParseError(e) => write!(f, "Parse error: {}", e),
            PolErr::InvalidOperator(op) => write!(f, "Invalid operator: {}", op),
            PolErr::NotEnoughValues(p) => write!(f, "> {p}\n Not enough values on stack"),
            PolErr::DivisionByZero => write!(f, "Division by zero"),
        }
    }
}

impl std::error::Error for PolErr {}
