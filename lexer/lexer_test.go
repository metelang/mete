package lexer

import (
	"metelang/token"
	"testing"
)

func TestNextToken(t *testing.T) {
	input := `=;(){},+`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.ASSIGN, "="},
		{token.SEMICOLON, ";"},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.COMMA, ","},
		{token.PLUS, "+"},
		{token.EOF, ""},
	}

	l := New(input);
	for i, tt := range tests {
		tk := l.NextToken()

		if tk.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q", i, tt.expectedLiteral, tk.Literal)
		}

		if tk.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q", i, tt.expectedLiteral, tk.Literal)
		}
	}
}

func TestNextTokenForIdentifierAssignments(t *testing.T) {
	input := `
      let one=1;
      let two=2;

      let add = function(x, y) {
         x+y;
      };

      let result = add(one, two);
    `

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.LET, "let"},
		{token.IDENTITY, "one"},
		{token.ASSIGN, "="},
		{token.INT, "1"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENTITY, "two"},
		{token.ASSIGN, "="},
		{token.INT, "2"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENTITY, "add"},
		{token.ASSIGN, "="},
		{token.FUNCTION, "function"},
		{token.LPAREN, "("},
		{token.IDENTITY, "x"},
		{token.COMMA, ","},
		{token.IDENTITY, "y"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.IDENTITY, "x"},
		{token.PLUS, "+"},
		{token.IDENTITY, "y"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENTITY, "result"},
		{token.ASSIGN, "="},
		{token.IDENTITY, "add"},
		{token.LPAREN, "("},
		{token.IDENTITY, "one"},
		{token.COMMA, ","},
		{token.IDENTITY, "two"},
		{token.RPAREN, ")"},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}

	l := New(input);
	for i, tt := range tests {
		tk := l.NextToken()

		if tk.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q", i, tt.expectedType, tk.Type)
		}

		if tk.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q", i, tt.expectedLiteral, tk.Literal)
		}
	}
}

func TestNextTokenIfElseReturnTrueFalseOperands(t *testing.T) {
	input := `
      if (1 < 2) {
        return true;
      } else {
         return false; 
      }
    `

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.IF, "if"},
		{token.LPAREN, "("},
		{token.INT, "1"},
		{token.LT, "<"},
		{token.INT, "2"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RETURN, "return"},
		{token.TRUE, "true"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.ELSE, "else"},
		{token.LBRACE, "{"},
		{token.RETURN, "return"},
		{token.FALSE, "false"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.EOF, ""},
	}

	l := New(input);
	for i, tt := range tests {
		tk := l.NextToken()

		if tk.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q", i, tt.expectedType, tk.Type)
		}

		if tk.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q", i, tt.expectedLiteral, tk.Literal)
		}
	}
}
