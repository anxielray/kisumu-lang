package lexer_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/Zone01-Kisumu-Open-Source-Projects/kisumu-lang/internal/lexer"
	"github.com/Zone01-Kisumu-Open-Source-Projects/kisumu-lang/internal/token"
)

func TestNextToken(t *testing.T) {
	input := `
	int age = 20;
	if (age > 18) {
		age = age + 1;
		// This is a comment
		/* Multi-line
		   comment */
	}
	`

	expectedTokens := []token.Token{
		{Type: token.TYPE, Literal: "int"},
		{Type: token.IDENTIFIER, Literal: "age"},
		{Type: token.ASSIGN, Literal: "="},
		{Type: token.INT, Literal: "20"},
		{Type: token.SEMICOLON, Literal: ";"},
		{Type: token.IF, Literal: "if"},
		{Type: token.LEFTPARENTHESIS, Literal: "("},
		{Type: token.IDENTIFIER, Literal: "age"},
		{Type: token.GREATERTHAN, Literal: ">"},
		{Type: token.INT, Literal: "18"},
		{Type: token.RIGHTPARENTHESIS, Literal: ")"},
		{Type: token.LEFTBRACE, Literal: "{"},
		{Type: token.IDENTIFIER, Literal: "age"},
		{Type: token.ASSIGN, Literal: "="},
		{Type: token.IDENTIFIER, Literal: "age"},
		{Type: token.PLUS, Literal: "+"},
		{Type: token.INT, Literal: "1"},
		{Type: token.SEMICOLON, Literal: ";"},
		{Type: token.RIGHTBRACE, Literal: "}"},
		{Type: token.EOF, Literal: ""},
	}

	l := lexer.New(input, lexer.DefaultConfig)

	for i, expected := range expectedTokens {
		tok := l.NextToken()

		if tok.Type != expected.Type || tok.Literal != expected.Literal {
			t.Fatalf("tests[%d] - token mismatch: expected=%+v, got=%+v", i, expected, tok)
		}
	}
}

func TestLexerWithSampleFiles(t *testing.T) {
	testsDir := "../../tests/lexer/"

	testFiles := []struct {
		fileName       string
		expectedTokens []token.Token
	}{
		{
			"token_tests.ksm",
			[]token.Token{
				{Type: token.TYPE, Literal: "int"},
				{Type: token.IDENTIFIER, Literal: "x"},
				{Type: token.ASSIGN, Literal: "="},
				{Type: token.INT, Literal: "10"},
				{Type: token.SEMICOLON, Literal: ";"},
				{Type: token.IF, Literal: "if"},
				{Type: token.LEFTPARENTHESIS, Literal: "("},
				{Type: token.IDENTIFIER, Literal: "x"},
				{Type: token.GREATERTHAN, Literal: ">"},
				{Type: token.INT, Literal: "5"},
				{Type: token.RIGHTPARENTHESIS, Literal: ")"},
				{Type: token.LEFTBRACE, Literal: "{"},
				{Type: token.IDENTIFIER, Literal: "x"},
				{Type: token.ASSIGN, Literal: "="},
				{Type: token.IDENTIFIER, Literal: "x"},
				{Type: token.PLUS, Literal: "+"},
				{Type: token.INT, Literal: "1"},
				{Type: token.SEMICOLON, Literal: ";"},
				{Type: token.RIGHTBRACE, Literal: "}"},
				{Type: token.EOF, Literal: ""},
			},
		},
	}

	for _, test := range testFiles {
		input, err := os.ReadFile(filepath.Join(testsDir, test.fileName))
		if err != nil {
			t.Fatalf("failed to read test file %s: %v", test.fileName, err)
		}

		// Init lexer
		l := lexer.New(string(input), lexer.DefaultConfig)

		// Compare tokens
		for i, expected := range test.expectedTokens {
			tok := l.NextToken()
			if tok.Type != expected.Type || tok.Literal != expected.Literal {
				t.Fatalf(
					"File %s tests[%d] - token mismatch: expected=%+v, got=%+v",
					test.fileName, i, expected, tok,
				)
			}
		}
	}
}
