package ast

import (
	"monkey/token"
)

type ExpressionStatement struct {
	Token      *token.Token
	Expression Expression
}

func (ex *ExpressionStatement) statementNode() {
}

func (ex *ExpressionStatement) TokenLiteral() string {
	return ex.Token.Literal
}

func (ex *ExpressionStatement) String() string {
	if ex.Expression != nil {
		return ex.Expression.String()
	}
	return ""
}
