package token

// TokenType represents type of token for lexer
type TokenType string

// Token for lexer
type Token struct {
	Type    TokenType
	Literal string
}

var keywords = map[string]TokenType{
	"fn":     FUNCTION,
	"let":    LET,
	"true":   TRUE,
	"false":  FALSE,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
}

// LookupIdent checks to see if a given identifier is a builtin keyword
func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}

const (
	//ILLEGAL illegal type
	ILLEGAL = "ILLEGAL"
	//EOF end of file
	EOF = "EOF"

	//IDENT Identifiers
	IDENT = "IDENT" // add, foodbar, x, y ...
	//INT integers
	INT = "INT" // 123456

	//ASSIGN is assignation type
	ASSIGN = "="
	//PLUS for addition
	PLUS     = "+"
	MINUS    = "-"
	BANG     = "!"
	ASTERISK = "*"
	SLASH    = "/"
	LT       = "<"
	GT       = ">"
	//COMMA comma delim
	COMMA = ","
	//SEMICOLON semicolo delim
	SEMICOLON = ";"
	//LPAREN left parens
	LPAREN = "("
	//RPAREN right parens
	RPAREN = ")"
	//LBRACE left curly
	LBRACE = "{"
	//RBRACE right curly
	RBRACE = "}"

	//FUNCTION keyword for func
	FUNCTION = "FUNCTION"
	//LET for assignation
	LET    = "LET"
	TRUE   = "TRUE"
	FALSE  = "FALSE"
	IF     = "IF"
	ELSE   = "ELSE"
	RETURN = "RETURN"
	EQ     = "=="
	NOT_EQ = "!="
)
