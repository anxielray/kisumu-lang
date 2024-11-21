package lexer

import "github.com/Zone01-Kisumu-Open-Source-Projects/kisumu-lang/internal/token"

/*
The Lexer's purpose in the language comes down to breaking down source code written as input to smaller understandable and easier to work with units referred as tokens through a process known as lexical analysis.

Any source code written will be stripped down to tokens by the lexer and compared against the tokens defined on the language. Incase of a mismatch, a syntax error will likely be flagged.
*/

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

// Will check if a character is a letter by either being an alphabet letter or an underscore
func isLetter(char byte) bool {
	return 'a' <= char && char <= 'z' || 'A' <= char && char <= 'Z' || char == '_'
}

// Will read a sequence of digit characters from the input and returns it a string
func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.char) {
		l.readCurrentCharacter()
	}
	return l.input[position:l.position]
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

// Will skip and ignore any occurrence of a tab, newline, carriage return and space character
func (l *Lexer) skipWhiteSpace() {
	for l.char == ' ' || l.char == '\t' || l.char == '\n' || l.char == '\r' {
		l.readCurrentCharacter()
	}
}

// Will check the next character in the input string without advancing the read position.
func (l *Lexer) peekNextCharacter() byte {
	if l.readPosition >= len(l.input) {
		return 0
	}
	return l.input[l.readPosition]
}

// Will read the next token from the input source
func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhiteSpace()

	switch l.char {
	// Checks whether the next token is either EQUAL (==) or ASSIGN (=)
	case '=':
		if l.peekNextCharacter() == '=' {
			ch := l.char
			l.readCurrentCharacter()
			literal := string(ch) + string(l.char)
			tok = token.Token{Type: token.EQUAL, Literal: literal}
		} else {
			tok = newToken(token.ASSIGN, l.char)
		}
	case '+':
		tok = newToken(token.PLUS, l.char)
	case '-':
		tok = newToken(token.MINUS, l.char)
		// Checks whether the next token is either NOTEQUAL (!=) or NOT (!)
	case '!':
		if l.peekNextCharacter() == '=' {
			ch := l.char
			l.readCurrentCharacter()
			literal := string(ch) + string(l.char)
			tok = token.Token{Type: token.NOTEQUAL, Literal: literal}
		} else {
			tok = newToken(token.NOT, l.char)
		}
	case '*':
		tok = newToken(token.ASTERISK, l.char)
	case '/':
		tok = newToken(token.SLASH, l.char)
	case '<':
		tok = newToken(token.LESSTHAN, l.char)
	case '>':
		tok = newToken(token.GREATERTHAN, l.char)
	case ',':
		tok = newToken(token.COMMA, l.char)
	case ';':
		tok = newToken(token.SEMICOLON, l.char)
	case '(':
		tok = newToken(token.LEFTPARENTHESIS, l.char)
	case ')':
		tok = newToken(token.RIGHTPARENTHESIS, l.char)
	case '{':
		tok = newToken(token.LEFTBRACE, l.char)
	case '}':
		tok = newToken(token.RIGHTBRACE, l.char)
	case '"':
		tok.Type = token.STRING
		tok.Literal = l.readString()
	case '[':
		tok = newToken(token.LEFTBRACKET, l.char)
	case ']':
		tok = newToken(token.RIGHTBRACKET, l.char)
	case ':':
		tok = newToken(token.COLON, l.char)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
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
			tok = newToken(token.ILLEGAL, l.char)
		}
	}
	l.readCurrentCharacter()
	return tok
}
