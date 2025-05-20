use std::{fmt, num::ParseFloatError, time::SystemTimeError};

pub mod file;
pub mod repl;
pub mod rpn;

#[derive(Debug, PartialEq)]
pub enum PolErr {
    ParseError(ParseFloatError),
    InvalidOperator(String),
    NotEnoughValues(String),
    DivisionByZero,

    FileParseErr(String),
    SystemTimeErr(String),
}

impl From<ParseFloatError> for PolErr {
    fn from(err: ParseFloatError) -> Self {
        PolErr::ParseError(err)
    }
}
impl From<SystemTimeError> for PolErr {
    fn from(err: SystemTimeError) -> Self {
        PolErr::SystemTimeErr(err.to_string())
    }
}

impl fmt::Display for PolErr {
    fn fmt(&self, f: &mut fmt::Formatter<'_>) -> fmt::Result {
        match self {
            PolErr::ParseError(e) => write!(f, "[PARSE ERR]: {e}"),
            PolErr::InvalidOperator(op) => write!(f, "[OPERATOR ERR]: invalid operator: {op}"),
            PolErr::NotEnoughValues(p) => write!(f, "[STACK ERR]: {p}"),
            PolErr::DivisionByZero => write!(f, "[DIVISION ERR]: division by zero not allowed"),
            PolErr::FileParseErr(err) => write!(f, "[FILE ERR]: {err}"),
            PolErr::SystemTimeErr(e) => write!(f, "[SYSTEM TIME ERR]: {e}"),
        }
    }
}

impl std::error::Error for PolErr {}
