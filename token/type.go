package token

const (
	Illegal = "ILLEGAL"
	EOF     = "EOF"

	Ident  = "IDENT"
	Int    = "INT"
	Real   = "REAL"
	String = "STRING"

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
