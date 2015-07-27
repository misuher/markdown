package markdown

import (
	"bufio"
	"bytes"
	"io"
)

//Scanner is a input buffer wrapper from within we try to obtain the tokens
type Scanner struct {
	r *bufio.Reader
}

//Item is compose of a token and its literal representation
type Item struct {
	Tok Token
	Lit string
}

//NewScanner initializes the Scanner type
func NewScanner(r io.Reader) *Scanner {
	return &Scanner{r: bufio.NewReader(r)}
}

//Scan returns the next token and value
func (s *Scanner) Scan() Item {
	ch := s.read()

	switch ch {
	case '#':
		return s.scanHash(ch)
	case '*':
		return s.scanAsterisks()
	case '!':
		return Item{Tok: EXCLAMATION, Lit: string(ch)}
	case '[':
		return Item{Tok: SQUAREOPEN, Lit: string(ch)}
	case ']':
		return Item{Tok: SQUARECLOSE, Lit: string(ch)}
	case '(':
		return Item{Tok: PARANOPEN, Lit: string(ch)}
	case ')':
		return Item{Tok: PARANCLOSE, Lit: string(ch)}
	case '>':
		return Item{Tok: GT, Lit: string(ch)}
	case ' ':
		return s.scanWhitespace(ch)
	case '\t':
		return Item{Tok: TAB, Lit: string(ch)}
	case '\n':
		return Item{Tok: NL, Lit: string(ch)}
	case eof:
		return Item{Tok: EOF, Lit: "\x00"}
	default:
		return s.scanLiteral(ch)
	}
}

//scanHash detects a chain of hashes+whitespace and returns its token
func (s *Scanner) scanHash(ch rune) Item {
	counter := 1
	for {
		ch = s.read()

		if ch == '#' {
			counter++
		} else {
			s.unread()
			break
		}
	}

	switch counter {
	case 1:
		return Item{HASH, "#"}
	case 2:
		return Item{DOUBLEHASH, "##"}
	case 3:
		return Item{TRIPLEHASH, "###"}
	default:
		return Item{QUADHASH, "####"}
	}
}

//scanWhitespace detecs if it is a single withespace or a four row whitespaces to handle it like a tab
func (s *Scanner) scanWhitespace(ch rune) Item {
	counter := 1
	for {
		ch = s.read()
		if ch == ' ' {
			counter++
		} else {
			s.unread()
			break
		}
		if counter == 4 {
			return Item{TAB, "\t"}
		}
	}
	return Item{WS, " "}
}

//scanString keeps reading runes to get a literal
func (s *Scanner) scanLiteral(ch rune) Item {
	var literal bytes.Buffer

	for isAlphanum(ch) || isNotToken(ch) || isWhiteSpace(ch) {
		literal.WriteString(string(ch))
		ch = s.read()
	}
	s.unread()

	return Item{LITERAL, literal.String()}
}

//scanAsterisks detecs if it is a single or double asterisk token
func (s *Scanner) scanAsterisks() Item {
	counter := 1
	ch := s.read()
	for ch == '*' {
		counter++
		ch = s.read()
	}

	s.unread()
	switch counter {
	case 1:
		return Item{ASTERISK, "*"}
	case 2:
		return Item{DOUBLEASTERISK, "**"}
	case 3:
		return Item{TRIPLEASTERISK, "***"}
	}

	//if there is more than 3 asterisk we will treat them like a literal
	var lit bytes.Buffer
	for i := 0; i < counter; i++ {
		lit.WriteString("*")
	}
	return Item{LITERAL, lit.String()}
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

func isAlphanum(ch rune) bool {
	return ch >= 'a' && ch <= 'z' || ch >= 'A' && ch <= 'Z' || ch >= '1' && ch <= '9'
}

func isNotToken(ch rune) bool {
	notTokens := ":/.,<{}^%=@~&-_ñÑ"
	for _, elem := range notTokens {
		if elem == ch {
			return true
		}
	}
	return false
}

func isWhiteSpace(ch rune) bool {
	return string(ch) == " "
}

var eof = rune(0)
