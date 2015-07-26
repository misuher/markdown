package markdown_test

import (
	"strings"
	"testing"

	. "github.com/misuher/markdown/markparser"
)

func TestScan(t *testing.T) {
	var tests = []struct {
		input    string
		expected Item
	}{
		{"", Item{EOF, "\x00"}},
		{"#", Item{HASH, "#"}},
		{"##", Item{DOUBLEHASH, "##"}},
		{"###", Item{TRIPLEHASH, "###"}},
		{"####", Item{QUADHASH, "####"}},
		{"*", Item{ASTERISK, "*"}},
		{"**", Item{DOUBLEASTERISK, "**"}},
		{" ", Item{WS, " "}},
		{"\t", Item{TAB, "\t"}},
		{"\n", Item{NL, "\n"}},
		{"[", Item{SQUAREOPEN, "["}},
		{"]", Item{SQUARECLOSE, "]"}},
		{"(", Item{PARANOPEN, "("}},
		{")", Item{PARANCLOSE, ")"}},
		{"!", Item{EXCLAMATION, "!"}},
		{"aaaa", Item{LITERAL, "aaaa"}},
		{"AAAA", Item{LITERAL, "AAAA"}},
		{"1234", Item{LITERAL, "1234"}},
		{"a1b2", Item{LITERAL, "a1b2"}},
	}

	for pos, elem := range tests {
		s := NewScanner(strings.NewReader(elem.input))
		item := s.Scan()

		if elem.expected.Tok != item.Tok {
			t.Errorf("Test index %d: %q token mismatch, expected %q but got %q", pos, elem.input, elem.expected.Tok, item.Tok)
		} else if elem.expected.Lit != item.Lit {
			t.Errorf("Test index %d:  %q literal mismatch, expected %q but got %q", pos, elem.input, elem.expected.Lit, item.Lit)
		}
	}
}
