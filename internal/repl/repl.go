/*
The REPL provides an environment / playground to interact with the language.

Although the language is not currently ready, the REPL will be used to visualize the lexer's functionality in the language of breaking source code to tokens.
*/

package repl

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"github.com/Zone01-Kisumu-Open-Source-Projects/kisumu-lang/internal/lexer"
	"github.com/Zone01-Kisumu-Open-Source-Projects/kisumu-lang/internal/token"
)

// Start function reads a .ksm file, tokenizes its content using a lexer, and prints each token until the end of the file is reached.
func Start(out io.Writer) {
	ksmFile, err := os.Open("./main.ksm")
	if err != nil {
		fmt.Printf("Failed to open source file: %s\n", err)
		os.Exit(1)
	}
	defer ksmFile.Close()

	scanner := bufio.NewScanner(ksmFile)

	for {
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line, lexer.DefaultConfig)

		// Prints all the tokens the lexer gives out until we encounter EOF.
		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Printf("%+v\n", tok)
		}
	}
}
