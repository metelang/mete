package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// Identifiers + literals
	IDENTITY = "IDENTITY"
	INT      = "INT"

	// Operators
	ASSIGN = "="
	PLUS   = "+"
	MINUS  = "-"
	EXCL   = "!"
	MLTP   = "*"
	DIVIDE = "/"

	LT = "<"
	GT = ">"

	//DELIMETERS
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "(";
	RPAREN = ")";
	LBRACE = "{"
	RBRACE = "}"

	// Keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
	TRUE     = "TRUE"
	FALSE    = "FALSE"
	IF       = "IF"
	ELSE     = "ELSE"
	RETURN   = "RETURN"
)

var reservedKeywords = map[string]TokenType{
	"function": FUNCTION,
	"let":      LET,
	"true":     TRUE,
	"false":    FALSE,
	"if":       IF,
	"else":     ELSE,
	"return":   RETURN,
}

func LookupKeywords(word string) TokenType {
	if tk, ok := reservedKeywords[word]; ok {
		return tk;
	}
	return IDENTITY;
}
