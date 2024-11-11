package token

// Token struct to represent a token with its type and literal.
type Token struct {
	Type, Literal string
}

// Constants representing the types of tokens.
const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	IDENTIFIER = "IDENTIFIER"
	INT        = "INT"
	STRING     = "STRING"

	// Operators token
	ASSIGN  = "="
	PLUS    = "+"
	MINUS   = "-"
	ASTERIK = "*"
	SLASH   = "/"

	// Comparisons token
	LESSTHAN    = "<"
	GREATERTHAN = ">"
	EQUAL       = "=="
	NOT         = "!"
	NOTEQUAL    = "!="

	COMMA     = ","
	SEMICOLON = ";"
	COLON     = ":"

	LEFTPARENTHESIS  = "("
	RIGHTPARENTHESIS = ")"
	LEFTBRACE        = "{"
	RIGHTBRACE       = "}"
	LEFTBRACKET      = "["
	RIGHTBRACKET     = "]"

	FUNCTION = "FUNCTION"
	VAR      = "VAR"
	TRUE     = "TRUE"
	FALSE    = "FALSE"
	IF       = "IF"
	ELSE     = "ELSE"
	RETURN   = "RETURN"
)

// Map keywords in the language to their respective token types.
var keywords = map[string]string{
	"func":   FUNCTION,
	"var":    VAR,
	"true":   TRUE,
	"false":  FALSE,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
}

// Check if a given identifier is a keyword and return the corresponding type. Returns the IDENTIFIER type if not a keyword.
func LookupIdentifier(identifier string) string {
	tok, ok := keywords[identifier]
	if ok {
		return tok
	}
	return IDENTIFIER
}
