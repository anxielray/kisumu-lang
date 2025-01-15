/*
The REPL provides an environment / playground to interact with the language.

Although the language is not currently ready, the REPL will be used to visualize the lexer's functionality in the language of breaking source code to tokens.
*/

package repl

import (
	"bufio"
	"fmt"
	"os"

	"github.com/Zone01-Kisumu-Open-Source-Projects/kisumu-lang/internal/lexer"
	"github.com/Zone01-Kisumu-Open-Source-Projects/kisumu-lang/internal/token"
)

const PROMPT = ">> "

// Start launches the REPL for lexer testing.
func Start() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Welcome to the Kisumu Lang Lexer REPL!")
	fmt.Println("Type your code below. Press Ctrl+C to exit.")

	for {
		fmt.Print(PROMPT)
		if !scanner.Scan() {
			break
		}
		line := scanner.Text()
		l := lexer.New(line, lexer.DefaultConfig)

		// Print tokens until EOF
		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Printf("%+v\n", tok)
		}
	}
}
