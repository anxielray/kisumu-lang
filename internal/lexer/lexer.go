package lexer

import (
	"fmt"
	"unicode"

	"github.com/Zone01-Kisumu-Open-Source-Projects/kisumu-lang/internal/token"
)

/*
The Lexer's purpose in the language comes down to breaking down source code written as input to smaller understandable and easier to work with units referred as tokens through a process known as lexical analysis.

Any source code written will be stripped down to tokens by the lexer and compared against the tokens defined on the language. Incase of a mismatch, a syntax error will likely be flagged.
*/

// LexerConfig provides configuration options for the lexer
type LexerConfig struct {
	AllowComments    bool   // Toggle comment handling
	StringDelimiters []byte // Allowed string delimiters, e.g., `"`, `'`
}

// DefaultLexerConfig defines default settings for the lexer
var DefaultLexerConfig = LexerConfig{
	AllowComments:    true,
	StringDelimiters: []byte{'"'},
}

// Lexer performs lexical analysis on the input source code
type Lexer struct {
	input        string
	position     int  // Current position in the input
	readPosition int  // Next position in the input
	char         byte // Current character being processed
	line         int  // Current line number
	column       int  // Current column number
	config       LexerConfig
}

// New creates a new Lexer instance with the provided configuration
func New(input string, config LexerConfig) *Lexer {
	l := &Lexer{
		input:  input,
		config: config,
		line:   1,
		column: 0,
	}
	l.readCurrentCharacter()
	return l
}

// readCurrentCharacter advances the lexer to the next character in the input
func (l *Lexer) readCurrentCharacter() {
	if l.readPosition >= len(l.input) {
		l.char = 0 // Null character to indicate EOF
	} else {
		l.char = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition++
	if l.char == '\n' {
		l.line++
		l.column = 0
	} else {
		l.column++
	}
}

// peekNextCharacter returns the next character without advancing the position
func (l *Lexer) peekNextCharacter() byte {
	if l.readPosition >= len(l.input) {
		return 0
	}
	return l.input[l.readPosition]
}

// skipWhiteSpace skips over whitespace and optionally comments
func (l *Lexer) skipWhiteSpace() {
	for unicode.IsSpace(rune(l.char)) {
		l.readCurrentCharacter()
	}

	if l.config.AllowComments {
		// Handle single-line comments
		if l.char == '/' && l.peekNextCharacter() == '/' {
			for l.char != '\n' && l.char != 0 {
				l.readCurrentCharacter()
			}
			l.skipWhiteSpace() // Continue skipping whitespace
		}

		// Handle multi-line comments
		if l.char == '/' && l.peekNextCharacter() == '*' {
			l.readCurrentCharacter() // Consume '/'
			l.readCurrentCharacter() // Consume '*'
			for !(l.char == '*' && l.peekNextCharacter() == '/') && l.char != 0 {
				l.readCurrentCharacter()
			}
			if l.char == '*' {
				l.readCurrentCharacter() // Consume '*'
				l.readCurrentCharacter() // Consume '/'
			} else {
				fmt.Printf("Unterminated comment at line %d, column %d\n", l.line, l.column)
			}
			l.skipWhiteSpace()
		}
	}
}

// readString reads a string literal, supporting multiple delimiters
func (l *Lexer) readString() string {
	startDelimiter := l.char
	position := l.position + 1
	for {
		l.readCurrentCharacter()
		if l.char == startDelimiter {
			break
		}
		if l.char == 0 {
			return "unterminated string"
		}
	}
	return l.input[position:l.position]
}

// readIdentifier reads an identifier or keyword
func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.char) {
		l.readCurrentCharacter()
	}
	return l.input[position:l.position]
}

// readNumber reads a number literal, supporting integers and floats
func (l *Lexer) readNumber() string {
	position := l.position
	isFloat := false
	for isDigit(l.char) || (l.char == '.' && !isFloat) {
		if l.char == '.' {
			isFloat = true
		}
		l.readCurrentCharacter()
	}
	return l.input[position:l.position]
}

// newToken creates a new token with the given type and literal
func newToken(tokenType string, literal string) token.Token {
	return token.Token{Type: tokenType, Literal: literal}
}

// NextToken retrieves the next token from the input
func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhiteSpace()

	// Handle single-character tokens
	if tokType, ok := singleCharTokens[l.char]; ok {
		tok = newToken(tokType, string(l.char))
		l.readCurrentCharacter()
		return tok
	}

	switch l.char {
	case '=':
		tok = l.matchTwoCharToken('=', token.ASSIGN, token.EQUAL)
	case '!':
		tok = l.matchTwoCharToken('=', token.NOT, token.NOTEQUAL)
	case '"', '\'':
		if contains(l.config.StringDelimiters, l.char) {
			tok.Type = token.STRING
			tok.Literal = l.readString()

			// Check for unterminated string
			if tok.Literal == "unterminated string" {
				tok.Type = token.ILLEGAL
				tok.Literal = "unterminated string"
			}
			return tok
		}
	case 0:
		tok.Type = token.EOF
		tok.Literal = ""
	default:
		if isLetter(l.char) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdentifier(tok.Literal)
			return tok
		} else if isDigit(l.char) {
			tok.Type = token.INT
			tok.Literal = l.readNumber()
			return tok
		} else {
			tok = newToken(token.ILLEGAL, string(l.char))
		}
	}

	l.readCurrentCharacter()
	return tok
}

// matchTwoCharToken matches either a single-character or two-character token
func (l *Lexer) matchTwoCharToken(nextChar byte, singleType, doubleType string) token.Token {
	if l.peekNextCharacter() == nextChar {
		current := l.char
		l.readCurrentCharacter()
		literal := string(current) + string(l.char)
		return token.Token{Type: doubleType, Literal: literal}
	}
	return newToken(singleType, string(l.char))
}

// Utility functions
func isLetter(char byte) bool {
	return unicode.IsLetter(rune(char)) || char == '_'
}

func isDigit(char byte) bool {
	return '0' <= char && char <= '9'
}

func contains(slice []byte, item byte) bool {
	for _, b := range slice {
		if b == item {
			return true
		}
	}
	return false
}

// singleCharTokens maps single-character tokens to their types
var singleCharTokens = map[byte]string{
	'+': token.PLUS,
	'-': token.MINUS,
	'*': token.ASTERISK,
	'/': token.SLASH,
	'<': token.LESSTHAN,
	'>': token.GREATERTHAN,
	',': token.COMMA,
	';': token.SEMICOLON,
	'(': token.LEFTPARENTHESIS,
	')': token.RIGHTPARENTHESIS,
	'{': token.LEFTBRACE,
	'}': token.RIGHTBRACE,
	'[': token.LEFTBRACKET,
	']': token.RIGHTBRACKET,
	':': token.COLON,
}
