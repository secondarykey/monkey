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

var keywords = map[string]TokenType{
	"fn":     Function,
	"let":    Let,
	"true":   True,
	"false":  False,
	"if":     If,
	"else":   Else,
	"return": Return,
}

// keywordになかったらIdentを生成
func NewKeywordToken(val string) *Token {
	if typ, ok := keywords[val]; ok {
		return New(typ, val)
	}
	return New(Ident, val)
}
