package token

type TokenType string
//tokentype would be a string minimise complexity , u can add int bool etc but it will add the complexity

// TOKEN IS A CUSTOM DATA STRUCTURE WITH 2 DATA TYPES WITH THEIR NAMES
type Token struct {
	Type    TokenType
	Literal string
}
//literal are here string ....example value can be = or ==
//giving char to literal will make language scope way les because we will not be able identify == or strings
//in nexttoken function in lexer we switch cases for a char then  we  goinside that case so that we can find next suitable char, we can do it repeadtly so that we can combine these char to make a string 

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	IDENT  = "IDENT"
	INT    = "INT"
	STRING = "STRING"

	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	BANG     = "!"
	ASTERISK = "*"
	SLASH    = "/"

	LT = "<"
	GT = ">"

	EQ     = "=="
	NOT_EQ = "!="


	COMMA     = ","
	SEMICOLON = ";"

	LPAREN   = "("
	RPAREN   = ")"
	LBRACE   = "{"
	RBRACE   = "}"
	LBRACKET = "["
	RBRACKET = "]"
	COLON    = ":"

	FUNCTION = "FUNCTION"
	LET      = "LET"
	TRUE     = "TRUE"
	FALSE    = "FALSE"
	IF       = "IF"
	ELSE     = "ELSE"
	RETURN   = "RETURN"
	MACRO    = "MACRO"
)

var keywords = map[string]TokenType{
	"fn":        FUNCTION,
	"definekar": LET,
	"true":      TRUE,
	"false":     FALSE,
	"agar":      IF,
	"warna":     ELSE,
	"return":    RETURN,
	"macro":     MACRO,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
