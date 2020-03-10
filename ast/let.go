package ast

import (
	"interpreter/token"
)

type LetStatement struct {
	Token *token.Token
	Name  *Identifer
	Value Expression
}

func (ls *LetStatement) statementNode() {}

func (ls *LetStatement) TokenLiteral() string {
	return ls.Token.Literal
}
