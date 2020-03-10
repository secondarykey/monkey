package lexer_test

import (
	"testing"

	"monkey/lexer"
	"monkey/token"
)

func TestNextToken(t *testing.T) {

	input := "=+(){},;"
	tests := []struct {
		exType    token.TokenType
		exLiteral string
	}{
		{token.Assign, "="},
		{token.Plus, "+"},
		{token.LParen, "("},
		{token.RParen, ")"},
		{token.LBrace, "{"},
		{token.RBrace, "}"},
		{token.Comma, ","},
		{token.SemiColon, ";"},
		{token.EOF, ""},
	}

	lx := lexer.New(input)

	for i, tt := range tests {
		tok := lx.NextToken()
		if tok.Type != tt.exType {
			t.Errorf("tests[%d] - token type wrong.expected=%q, got=%q", i, tt.exType, tok.Type)
		}
		if tok.Literal != tt.exLiteral {
			t.Errorf("tests[%d] - token literal wrong.expected=%q, got=%q", i, tt.exLiteral, tok.Literal)
		}
	}

}

func TestNextToken2(t *testing.T) {

	input := `
    let five = 5;
    let ten = 10;

    let add = fn(x,y) {
        x + y;
    };

    let result = add(five,ten);
    `
	tests := []struct {
		exType    token.TokenType
		exLiteral string
	}{
		{token.Let, "let"},
		{token.Ident, "five"},
		{token.Assign, "="},
		{token.Int, "5"},
		{token.SemiColon, ";"},
		{token.Let, "let"},
		{token.Ident, "ten"},
		{token.Assign, "="},
		{token.Int, "10"},
		{token.SemiColon, ";"},
		{token.Let, "let"},
		{token.Ident, "add"},
		{token.Assign, "="},
		{token.Function, "fn"},
		{token.LParen, "("},
		{token.Ident, "x"},
		{token.Comma, ","},
		{token.Ident, "y"},
		{token.RParen, ")"},
		{token.LBrace, "{"},
		{token.Ident, "x"},
		{token.Plus, "+"},
		{token.Ident, "y"},
		{token.SemiColon, ";"},
		{token.RBrace, "}"},
		{token.SemiColon, ";"},
		{token.Let, "let"},
		{token.Ident, "result"},
		{token.Assign, "="},
		{token.Ident, "add"},
		{token.LParen, "("},
		{token.Ident, "five"},
		{token.Comma, ","},
		{token.Ident, "ten"},
		{token.RParen, ")"},
		{token.SemiColon, ";"},
		{token.EOF, ""},
	}

	lx := lexer.New(input)

	for i, tt := range tests {
		tok := lx.NextToken()
		if tok.Type != tt.exType {
			t.Errorf("tests[%d] - token type wrong.expected=%q, got=%q", i, tt.exType, tok.Type)
		}
		if tok.Literal != tt.exLiteral {
			t.Errorf("tests[%d] - token literal wrong.expected=%q, got=%q", i, tt.exLiteral, tok.Literal)
		}
	}
}
