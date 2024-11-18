package lexer

import "github.com/Zone01-Kisumu-Open-Source-Projects/kisumu-lang/internal/token"

type Lexer struct {
	input        string
	position     int
	readPosition int
	char         byte
}

// Will be used to create a new Lexer instance
func New(input string) *Lexer {
	l := &Lexer{
		input: input,
	}
	l.readCurrentCharacter()
	return l
}

// Will read the current character in the input
func (l *Lexer) readCurrentCharacter() {
	if l.readPosition >= len(l.input) {
		l.char = 0
	} else {
		l.char = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition++
}

// Will read characters from the input until it encounters a double quote or a null character
func (l *Lexer) readString() string {
	position := l.position + 1

	for {
		l.readCurrentCharacter()

		if l.char == '"' || l.char == 0 {
			break
		}
	}
	return l.input[position:l.position]
}

// Will create and return a new Token object using a token type and the character
func newToken(tokenType string, char byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(char)}
}

// Will read characters from an input that form an identifier starting from the current position and returns the identifier as a string
func (l *Lexer) readIdentifier() string {
	position := l.position

	for isLetter(l.char) {
		l.readCurrentCharacter()
	}
	return l.input[position:l.position]
}

func isLetter(char byte) bool {
	return 'a' <= char && char <= 'z' || 'A' <= char && char <= 'Z' || char == '_'
}
