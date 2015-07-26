package markdown

import (
	"bufio"
	"io"
)

//Scanner is a input buffer wrapper from within we try to obtain the tokens
type Scanner struct {
	r *bufio.Reader
}

//Item is compose of a token and its literal representation
type Item struct {
	tok Token
	lit string
}

//NewScanner initializes the Scanner type
func NewScanner(r io.Reader) *Scanner {
	return &Scanner{r: bufio.NewReader(r)}
}

//Scan returns the next token and value
func (s *Scanner) Scan() Item {
	ch := s.read()

	switch ch {
	case eof:
		return Item{tok: EOF, lit: ""}
	case '#':
		return s.scanHash(ch)
	case '*':
		return s.scanAsterisks(ch)
	case '!':
		return Item{tok: EXCLAMATION, lit: string(ch)}
	default:
		return s.scanLiteral(ch)
	}

}

//scanHash detects a chain of hashes+whitespace and returns its token
func (s *Scanner) scanHash(ch rune) Item {
	return Item{ILLEGAL, ""}
}

//scanWhitespace detecs if it is a single withespace or a four row whitespaces to handle it like a tab
func (s *Scanner) scanWhitespace(ch rune) Item {
	return Item{ILLEGAL, ""}
}

//scanString keeps reading runes to get a literal
func (s *Scanner) scanLiteral(ch rune) Item {
	return Item{ILLEGAL, ""}
}

//scanAsterisks detecs if it is a single or double asterisk token
func (s *Scanner) scanAsterisks(ch rune) Item {
	ch2 := s.read()

	if ch2 == '*' {
		return Item{tok: DOUBLEASTERISK, lit: string(ch)}
	}
	s.unread()
	return Item{tok: ASTERISK, lit: string(ch)}
}

func (s *Scanner) read() rune {
	ch, _, err := s.r.ReadRune()
	if err != nil {
		return eof
	}
	return ch
}

func (s *Scanner) unread() {
	_ = s.r.UnreadRune()
}

var eof = rune(0)
