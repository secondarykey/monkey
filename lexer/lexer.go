package lexer

import (
	"strings"

	"interpreter/token"
)

type Lexer struct {
	input        string
	position     int
	readPosition int
	ch           byte
}

func New(i string) *Lexer {
	l := Lexer{
		input: i,
	}
	l.readChar()
	return &l
}

func (lx *Lexer) readChar() {

	if lx.readPosition >= len(lx.input) {
		lx.ch = 0
	} else {
		lx.ch = lx.input[lx.readPosition]
	}

	lx.position = lx.readPosition
	lx.readPosition += 1
}

func (lx *Lexer) NextToken() *token.Token {

	lx.skip()

	var tok *token.Token

	val := ""
	if lx.ch != 0 {
		val = string(lx.ch)
	}

	switch lx.ch {
	case 0:
		tok = token.New(token.EOF, val)
	case '+':
		tok = token.New(token.Plus, val)
	case '-':
		tok = token.New(token.Minus, val)
	case '*':
		tok = token.New(token.Asterisk, val)
	case '/':
		tok = token.New(token.Slash, val)
	case ',':
		tok = token.New(token.Comma, val)
	case ';':
		tok = token.New(token.SemiColon, val)
	case '(':
		tok = token.New(token.LParen, val)
	case ')':
		tok = token.New(token.RParen, val)
	case '{':
		tok = token.New(token.LBrace, val)
	case '}':
		tok = token.New(token.RBrace, val)
	case '<', '>', '=', '!':

		nx := lx.peek()

		newVal := val
		if nx != 0 {
			newVal += string(nx)
		}

		if lx.ch == '<' {
			if nx == '=' {
				tok = token.New(token.LE, newVal)
				lx.readChar()
			} else {
				tok = token.New(token.LT, val)
			}
		} else if lx.ch == '>' {
			if nx == '=' {
				tok = token.New(token.GE, newVal)
				lx.readChar()
			} else {
				tok = token.New(token.GT, val)
			}
		} else if lx.ch == '=' {
			if nx == '=' {
				tok = token.New(token.Eq, newVal)
				lx.readChar()
			} else {
				tok = token.New(token.Assign, val)
			}
		} else if lx.ch == '!' {
			if nx == '=' {
				tok = token.New(token.NotEq, newVal)
				lx.readChar()
			} else {
				tok = token.New(token.Bang, val)
			}
		}
	case '"':

		val := lx.readString()
		tok = token.New(token.String, val)
		return tok

	default:
		//文字列の場合
		if isLetter(lx.ch) {
			ident := lx.readIdentifer()
			tok = token.NewKeywordToken(ident)
			return tok
		} else if isDigit(lx.ch) {
			//数値の場合
			num := lx.readNumber()

			typ := token.Int
			if strings.Index(num, ".") != -1 {
				typ = token.Real
			}

			tok = token.New(token.TokenType(typ), num)
			return tok
		} else {
			//不明
			tok = token.New(token.Illegal, val)
		}
	}

	lx.readChar()
	return tok
}

func (lx *Lexer) peek() byte {

	nextPos := lx.position + 1

	if nextPos >= len(lx.input) {
		return 0
	}
	return lx.input[nextPos]
}

func (lx *Lexer) readIdentifer() string {
	start := lx.position
	for isLetter(lx.ch) {
		lx.readChar()
	}
	return lx.input[start:lx.position]
}

func (lx *Lexer) readString() string {

	//" のはずなので1つ進める
	lx.readChar()

	start := lx.position
	for isString(lx.ch) {
		lx.readChar()
	}

	end := lx.position
	//1つすすめる
	lx.readChar()
	return lx.input[start:end]
}

func (lx *Lexer) readNumber() string {
	start := lx.position
	for isDigit(lx.ch) {
		lx.readChar()
	}
	return lx.input[start:lx.position]
}

func (lx *Lexer) skip() {
	for lx.ch == ' ' || lx.ch == '\t' ||
		lx.ch == '\n' || lx.ch == '\r' {
		lx.readChar()
	}
}

//TODO 途中からの数値をOKにする
func isLetter(ch byte) bool {
	return 'a' <= ch && 'z' >= ch ||
		'A' <= ch && 'Z' >= ch || ch == '_'
}

func isString(ch byte) bool {
	return '"' != ch
}

func isDigit(ch byte) bool {
	return '0' <= ch && '9' >= ch || '.' == ch
}
