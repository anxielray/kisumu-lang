package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	// OPERATORS
	ASSIGN             = "="
	PLUS               = "+"
	MULTIPLY           = "*"
	SUBTRACT           = "-"
	DIVIDE             = "/"
	MODULO             = "%"

	// COMPARISON OPERATORS
	//Using a prefix like CMP_  to makes it easy to identify 
	// these as comparison operators when using them, 
	// keeping THE code organized and easy to maintain.
	CMP_EQUAL              = "=="
	CMP_NOTEQUAL           = "!="
	CMP_GREATERTHAN        = ">"
	CMP_LESSTHAN           = "<"
	CMP_GREATERTHANOREQUAL = ">="
	CMP_LESSTHANOREQUAL    = "<="

	// LOGICAL OPERATORS
	CMP_AND = "&&"
	CMP_OR = "||"
	CMP_NOT = "!"

	// Assignment Operators
	ASSIGN_ADD = "+="
	ASSIGN_SUBTRACT = "-="
	ASSIGN_MULTIPLY = "*="
	ASSIGN_DIVIDE = "/="
	ASSIGN_MODULO = "%="

	// Increment/Decrement Operators
	INCREMENT = "++"
	DECREMENT = "--"

	// PUNCTUATION/SPECIAL SYMBOLS
	LPAREN    = "("
	RPAREN    = ")"
	LBRACE    = "{"
	RBRACE    = "}"
	LBRACKET  = "["
	RBRACKET  = "]"
	COMMA     = ","
	SEMICOLON = ";"
	COLON     = ":"

	// KEYWORDS
	KW_IF     = "if"
	KW_ELSE   = "else"
	KW_FOR    = "for"
	KW_WHILE  = "while"
	KW_RETURN = "return"
	KW_FUNC   = "func"
	KW_VAR    = "var"
	KW_CONST  = "const"
	KW_TRUE   = "true"
	KW_FALSE  = "false"

	// LITERALS
	LIT_INT   = "int"
	LIT_FLOAT = "float"
	LIT_STRING = "string"
	LIT_BOOL   = "bool"
)
