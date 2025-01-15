package main

import (
	"flag"
	"fmt"

	"github.com/Zone01-Kisumu-Open-Source-Projects/kisumu-lang/internal/repl"
)

func main() {
	file := flag.String("file", "", "File to read input from")
	flag.Parse()

	if *file != "" {
		repl.ReadFile(*file)
	} else {
		fmt.Println("No file provided. Starting REPL...")
		repl.Start()
	}
}
