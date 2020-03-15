package parser_test

import (
	"testing"

	"monkey/ast"
	"monkey/lexer"
	"monkey/parser"
)

func TestParseLetStatement(t *testing.T) {

	input := `
    let x = 123;
    let y = 123.256;
    let foobar = 49384023984;
    `

	lx := lexer.New(input)
	p := parser.New(lx)

	prog, err := p.Parse()
	if prog == nil {
		t.Fatalf("Parse() returned nil")
	}

	if err != nil {
		t.Fatalf("Parse() returned error [%v]", err)
	}

	if len(prog.Statements) != 3 {
		t.Fatalf("Program statements lenghth not 3 . got %d", len(prog.Statements))
	}

	tests := []struct {
		exIdentifer string
	}{
		{"x"},
		{"y"},
		{"foobar"},
	}

	for i, tt := range tests {
		stmt := prog.Statements[i]
		if !testParseLetStatement(t, stmt, tt.exIdentifer) {
			return
		}
	}

}

func testParseLetStatement(t *testing.T, stmt ast.Statement, name string) bool {

	if stmt.TokenLiteral() != "let" {
		t.Errorf("statement TokenLiteral not 'let'.got=%q", stmt.TokenLiteral())
		return false
	}

	letStmt, ok := stmt.(*ast.LetStatement)
	if !ok {
		t.Errorf("s not *ast.LetStatement.got = %T", stmt)
		return false
	}

	nv := letStmt.Name.Value
	if nv != name {
		t.Errorf("name value not %s. got = %s", name, nv)
		return false
	}

	nt := letStmt.Name.TokenLiteral()
	if nt != name {
		t.Errorf("name token literal value not %s. got = %s", name, nt)
		return false
	}

	return true
}

func TestParseReturnStatement(t *testing.T) {

	input := `
    return 5;
    return 10;
    return 993322;
    `

	lx := lexer.New(input)
	p := parser.New(lx)

	pg, err := p.Parse()
	if err != nil {
		t.Fatalf("parser.Parse() is error. got=%v", err)
	}

	if len(pg.Statements) != 3 {
		t.Fatalf("pg.Statements does not contain 3 statements. got=%d", len(pg.Statements))
	}

	for _, stmt := range pg.Statements {
		retStmt, ok := stmt.(*ast.ReturnStatement)
		if !ok {
			t.Errorf("stmt not *ast.ReturnStatement. got=%T", stmt)
			continue
		}
		if retStmt.TokenLiteral() != "return" {
			t.Errorf("returnStmt.TokenLiteral not 'return'. got %q", retStmt.TokenLiteral())
		}
	}

}
