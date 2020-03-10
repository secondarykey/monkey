package token

const (
	Illegal TokenType = "ILLEGAL"
	EOF     TokenType = "EOF"

	Ident  TokenType = "IDENT"
	Int    TokenType = "INT"
	Real   TokenType = "REAL"
	String TokenType = "STRING"

	Function TokenType = "FUNCTION"
	Let      TokenType = "LET"
	If       TokenType = "IF"
	Else     TokenType = "ELSE"
	True     TokenType = "TRUE"
	False    TokenType = "FALSE"
	Return   TokenType = "RETURN"

	Assign   TokenType = "ASSIGN"
	Bang     TokenType = "BANG"
	Plus     TokenType = "PLUS"
	Minus    TokenType = "MINUS"
	Asterisk TokenType = "ASTERISK"
	Slash    TokenType = "SLASH"
	LT       TokenType = "LESS_THAN"
	GT       TokenType = "GREATER_THAN"

	Equal    TokenType = "EQUAL"
	NotEqual TokenType = "NOT_EQUAL"
	LE       TokenType = "LESS_THAN_OR_EQUAL"
	GE       TokenType = "GREATER_THAN_OR_EQUAL"

	Comma     TokenType = "COMMA"
	SemiColon TokenType = "SEMI_COLON"

	LParen TokenType = "L_PAREN"
	RParen TokenType = "R_PAREN"
	LBrace TokenType = "L_BRACE"
	RBrace TokenType = "R_BRACE"
)
