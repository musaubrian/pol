package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	RESET   = "\033[0m"
	BLUEISH = 94
	REDISH  = 91
	DIM     = 2
	DARK    = 90
)

func StartRepl() {
	fmt.Println(colorize(DIM, "Welcome to POL\n"))
	rd := bufio.NewReader(os.Stdin)
	for {
		fmt.Print(colorize(BLUEISH, ">> "))
		expr, err := rd.ReadString('\n')
		if err != nil {
			log.Error(err.Error())
			return
		}
		expr = strings.TrimSpace(expr)
		if expr == ".exit" {
			fmt.Println(colorize(DIM, "\\. See ya"))
			return
		} else if strings.HasPrefix(expr, ".help") {
			printHelp(expr)
			continue
		} else if strings.HasPrefix(expr, ".") {
			fmt.Println(colorize(REDISH, "Unkown Command"))
			continue
		}
		st := NewStack()
		evaled := Eval(expr)
		st.Push(*evaled)

		for _, expression := range st.Expressions {
			if expression.Err != nil {
				fmt.Println(colorize(REDISH, expression.Err.Error()))
				continue
			}

			fmt.Println(colorize(DARK, fmt.Sprintf("%.3f", expression.Result)))
		}
	}
}

func printHelp(cmd string) {
	expression := `
Pol is a calculator that uses a slightly modified reverse polish notation

Operands: + - / ^ *

expressions are in the form:
	single:	5 4 +           ~ evaluates to 5 + 4
	grouped: (54+) (67-) +  ~ evaluates to (5 + 4) + (6 - 7)

Currently only supports a max of 2 grouped expressions not space separated
`
	commands := `
Available Commands:
	.exit           ~ exits the REPL
	.help <subcmd>  ~ shows docs on the subcommand
	      expr      ~ shows documentation on expressions
	      cmds      ~ shows help on commands
	      <empty>   ~ still brings up this menu
`
	cmds := strings.Split(cmd, " ")
	if len(cmds) == 1 {
		fmt.Println(colorize(DARK, commands))
		return
	}

	if cmds[1] == "expr" {
		fmt.Println(colorize(DARK, expression))
		return
	} else if cmds[1] == "cmd" {
		fmt.Println(colorize(DARK, commands))
		return
	} else {
		fmt.Println(colorize(DARK, commands))
		return
	}
}

func colorize(code int, text string) string {
	return fmt.Sprintf("\033[%sm%s%s", strconv.Itoa(code), text, RESET)
}
