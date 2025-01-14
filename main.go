package main

import (
	"flag"

	"github.com/musaubrian/logr"
)

var log = logr.New().WithColor()

func main() {
	var startRepl bool
	var file string

	flag.BoolVar(&startRepl, "r", false, "Open the REPL")
	flag.StringVar(&file, "f", "", "path to file")

	flag.Parse()
	if startRepl {
		StartRepl()
	} else if len(file) > 0 {
		log.Info("File parsing is unimplemented")
	} else {
		log.Warn("No action specified")
		flag.Usage()
	}
}
