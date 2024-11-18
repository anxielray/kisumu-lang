package lexer

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
