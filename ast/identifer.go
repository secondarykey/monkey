package ast

import (
	"monkey/token"
)

type Identifer struct {
	Token *token.Token
	Value string
}

func (id *Identifer) expressionNode() {}

func (id *Identifer) TokenLiteral() string {
	return id.Token.Literal
}
