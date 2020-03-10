package token

import ()

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

func New(typ TokenType, lit string) *Token {
	return &Token{Type: typ, Literal: lit}
}

const (
	Illegal = "ILLEGAL"
	EOF     = "EOF"

	Ident = "IDENT"
	Int   = "INT"

	Function = "FUNCTION"
	Let      = "LET"
	If       = "IF"
	Else     = "ELSE"
	True     = "TRUE"
	False    = "FALSE"
	Return   = "RETURN"

	Assign   = "="
	Bang     = "!"
	Plus     = "+"
	Minus    = "-"
	Asterisk = "*"
	Slash    = "/"
	LT       = "<"
	GT       = ">"

	Eq    = "=="
	NotEq = "!="
	LE    = "<="
	GE    = ">="

	Comma     = ","
	SemiColon = ";"

	LParen = "("
	RParen = ")"
	LBrace = "{"
	RBrace = "}"
)

var keywords = map[string]TokenType{
	"fn":     Function,
	"let":    Let,
	"true":   True,
	"false":  False,
	"if":     If,
	"else":   Else,
	"return": Return,
}

func NewKeywordToken(val string) *Token {
	if typ, ok := keywords[val]; ok {
		return New(typ, val)
	}
	return New(Ident, val)
}
