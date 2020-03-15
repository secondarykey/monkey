package parser

import (
	"fmt"

	"monkey/ast"
	"monkey/lexer"
	"monkey/token"
)

type Parser struct {
	lx           *lexer.Lexer
	currentToken *token.Token
	peekToken    *token.Token
}

func New(lx *lexer.Lexer) *Parser {
	p := Parser{}
	p.lx = lx

	p.nextToken()
	p.nextToken()

	return &p
}

func (p *Parser) nextToken() {
	p.currentToken = p.peekToken
	p.peekToken = p.lx.NextToken()
}

func (p *Parser) Parse() (*ast.Program, error) {

	prog := ast.Program{}

	for p.currentToken.Type != token.EOF {
		stmt, err := p.parseStatement()
		if err != nil {
			return nil, err
		} else {
			prog.Statements = append(prog.Statements, stmt)
		}

		p.nextToken()
	}

	return &prog, nil
}

var (
	NotImplementedError = fmt.Errorf("not implemented.")
)

func (p *Parser) parseStatement() (ast.Statement, error) {

	switch p.currentToken.Type {
	case token.Let:
		return p.parseLetStatement()
	case token.Return:
		return p.parseReturnStatement()
	default:
	}

	return nil, NotImplementedError
}

func (p *Parser) parseLetStatement() (*ast.LetStatement, error) {

	let := ast.LetStatement{}
	let.Token = p.currentToken

	if !p.expectPeek(token.Ident) {
		return nil, fmt.Errorf("let peek ident error")
	}

	ident := ast.Identifer{}

	ident.Token = p.currentToken
	ident.Value = p.currentToken.Literal

	let.Name = &ident
	if !p.expectPeek(token.Assign) {
		return nil, fmt.Errorf("let peek assign error")
	}

	//TODO 最後まで飛ばしているだけ
	for !p.currentTokenIs(token.SemiColon) {
		p.nextToken()
	}

	return &let, nil
}

func (p *Parser) expectPeek(t token.TokenType) bool {
	if p.peekTokenIs(t) {
		p.nextToken()
		return true
	}
	return false
}

func (p *Parser) currentTokenIs(t token.TokenType) bool {
	if p.currentToken.Type == t {
		return true
	}
	return false
}

func (p *Parser) peekTokenIs(t token.TokenType) bool {
	if p.peekToken.Type == t {
		return true
	}
	return false
}

func (p *Parser) parseReturnStatement() (*ast.ReturnStatement, error) {
	stmt := &ast.ReturnStatement{Token: p.currentToken}

	p.nextToken()

	//TODO 読み飛ばしているだけ
	for !p.currentTokenIs(token.SemiColon) {
		p.nextToken()
	}

	return stmt, nil
}
