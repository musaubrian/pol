# POL/RS

naive reverse polish notation calculator impl.

## Setup

```sh

git clone github.com/musaubrian/pol

cd pol

git switch rs
cargo build --release

./target/release/pol --help
```

## Usage

```sh
./target/release/pol --help

# repl support
./target/release/pol --repl


# file input support
./target/release/pol --file /path/to//file
```

## TODO:
- [x] Single Expression
- [x] Complex Expressions
- [x] file input expression evaluation
- [x] REPL
