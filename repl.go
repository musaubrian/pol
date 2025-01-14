package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func StartRepl() {
	fmt.Println("Welcome to Pol\nCommands: .quit")
	rd := bufio.NewReader(os.Stdin)
	for {
		fmt.Print(">> ")
		expr, err := rd.ReadString('\n')
		if err != nil {
			log.Error(err.Error())
			return
		}
		expr = strings.TrimSpace(expr)
		if expr == ".quit" {
			fmt.Println("\\. See ya")
			return
		} else if strings.HasPrefix(expr, ".") {
			log.Warn("Unknown command")
			return
		}
		st, err := Parse(expr)
		if err != nil {
			log.Error(err.Error())
			return
		}
		res := st.Eval()
		fmt.Println(res)
	}
}
