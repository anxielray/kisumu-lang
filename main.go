package main

import (
	"os"

	"github.com/Zone01-Kisumu-Open-Source-Projects/kisumu-lang/internal/repl"
)

func main() {
	repl.Start(os.Stdout)
}
