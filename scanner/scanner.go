package scanner

import (
	"fmt"
	"strconv"

	"github.com/regmicmahesh/crafting-interpreters/token"
)

type Scanner struct {
	Source string
	Tokens []*token.Token

	start   int
	current int
	line    int
}

func NewScanner(source string) *Scanner {

	return &Scanner{
		Source: source,
		Tokens: make([]*token.Token, 0),

		start:   0,
		current: 0,
		line:    1,
	}
}

func (s Scanner) isAtEnd() bool {
	return s.current >= len(s.Source)
}

func (s *Scanner) ScanTokens() []*token.Token {

	for !s.isAtEnd() {
		s.start = s.current
		s.scanToken()
	}

	s.Tokens = append(s.Tokens, token.NewToken(token.EOF, "", nil, s.line))

	return s.Tokens

}

func (s *Scanner) advance() byte {

	currChar := s.Source[s.current]
	s.current++

	return currChar
}

func (s *Scanner) addToken(type_ token.TokenType) {
	s.addTokenWithLiteral(type_, nil)
}

func (s *Scanner) addTokenWithLiteral(type_ token.TokenType, literal any) {

	text := s.Source[s.start:s.current]
	s.Tokens = append(s.Tokens, token.NewToken(type_, text, literal, s.line))

}

func (s Scanner) Peek() byte {
	if s.isAtEnd() {
		return byte(0)
	}
	return s.Source[s.current]
}

func (s *Scanner) Match(expected byte) bool {

	if s.isAtEnd() {
		return false
	}

	nextByte := s.Peek()

	if nextByte == expected {
		s.current++
		return true
	}

	return false

}

func (s *Scanner) scanToken() {

	c := s.advance()

	switch c {
	case '(':
		s.addToken(token.LEFT_PAREN)
	case ')':
		s.addToken(token.RIGHT_PAREN)
	case '{':
		s.addToken(token.LEFT_BRACE)
	case '}':
		s.addToken(token.RIGHT_BRACE)
	case ',':
		s.addToken(token.COMMA)
	case '.':
		s.addToken(token.DOT)
	case '-':
		s.addToken(token.MINUS)
	case '+':
		s.addToken(token.PLUS)
	case ';':
		s.addToken(token.SEMICOLON)
	case '*':
		s.addToken(token.STAR)

	case '!':
		if s.Match('=') {
			s.addToken(token.BANG_EQUAL)
		} else {
			s.addToken(token.BANG)
		}

	case '=':
		if s.Match('=') {
			s.addToken(token.EQUAL_EQUAL)
		} else {
			s.addToken(token.EQUAL)
		}

	case '>':
		if s.Match('=') {
			s.addToken(token.GREATER_EQUAL)
		} else {
			s.addToken(token.GREATER)
		}

	case '<':
		if s.Match('=') {
			s.addToken(token.LESS_EQUAL)
		} else {
			s.addToken(token.LESS)
		}

	case '/':
		if s.Match('/') {

			for (s.Peek() != '\n') && !s.isAtEnd() {
				s.advance()
			}
		} else {
			s.addToken(token.SLASH)
		}

	case ' ', '\r', '\t':
		fmt.Print("")

	case '\n':
		s.line++

	case '"':
		s.parseString()

	default:
		if s.isDigit(c) {
			s.parseNumber()
		} else if s.isAlpha(c) {
			s.parseIdentifier()
		}

	}

}

func (s *Scanner) isAlpha(c byte) bool {
	return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || c == '_'
}

func (s *Scanner) isAlphaNumeric(c byte) bool {

	return (s.isDigit(c) || s.isAlpha(c))

}

func (s *Scanner) parseIdentifier() {
	for s.isAlphaNumeric(s.Peek()) {
		s.advance()
	}

	text := s.Source[s.start:s.current]

	if keyword, found := token.GetKeyword(text); found {
		s.addToken(keyword)
	} else {
		s.addToken(token.IDENTIFER)
	}

}

func (s *Scanner) parseNumber() {

	for s.isDigit(s.Peek()) {
		s.advance()
	}

	if s.Peek() == '.' && s.isDigit(s.peekNext()) {

		// consume the '.'
		s.advance()

		for s.isDigit(s.Peek()) {
			s.advance()
		}

	}

	rawNumber := s.Source[s.start:s.current]

	num, err := strconv.ParseFloat(rawNumber, 10)

	if err != nil {
		panic(err)
	}

	s.addTokenWithLiteral(token.NUMBER, num)

}

func (s Scanner) peekNext() byte {

	if s.current+1 >= len(s.Source) {
		return byte(0)
	}

	return s.Source[s.current+1]

}

func (s Scanner) isDigit(c byte) bool {

	return c >= '0' && c <= '9'

}

func (s *Scanner) parseString() {

	for (s.Peek() != '"') && !s.isAtEnd() {

		if s.Peek() == '\n' {
			s.line++
		}

		s.advance()

	}

	if s.isAtEnd() {
		panic("Unterminated String")
	}

	// consume the closing "
	s.advance()

	val := s.Source[s.start+1 : s.current-1]
	s.addTokenWithLiteral(token.STRING, val)

}
