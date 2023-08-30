package token

import "fmt"

type Token struct {
	Type    TokenType
	Lexeme  string
	Literal any
	Line    int
}

func NewToken(type_ TokenType, lexeme string, literal any, line int) *Token {

	return &Token{
		Type:    type_,
		Lexeme:  lexeme,
		Literal: literal,
		Line:    line,
	}

}

func (t Token) String() string {
	return fmt.Sprintf("%v %v %v", t.Type.String(), t.Lexeme, t.Literal)
}
