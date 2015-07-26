package markdown_test

import (
	"strings"
	"testing"

	"github.com/misuher/markdown/markparser"
)

func TestScan(t *testing.T) {
	var tests = []struct {
		input    string
		expected item
	}{
		{input: "", expected: {markdown.EOF, "\x00"}},
		{input: "#", expected: {markdown.HASH, "#"}},
		{input: "##", expected: {markdown.DOUBLEHASH, "##"}},
		{input: "###", expected: {markdown.TRIPLEHASH, "###"}},
		{input: "####", expected: {markdown.QUADHASH, "####"}},
		{input: "*", expected: {markdown.ASTERISK, "*"}},
		{input: "**", expected: {markdown.DOUBLEASTERISK, "**"}},
		{input: " ", expected: {markdown.WS, ""}},
		{input: "\t", expected: {markdown.TAB, ""}},
		{input: "\n", expected: {markdown.NL, ""}},
	}

	for pos, elem := range tests {
		s := markdown.NewScanner(strings.NewReader(elem.input))
		item := s.Scan()

		if elem.expected.tok != item.tok {
			t.Errorf("%d. %q token mismatch, expected %q but got %q", pos, elem.input, elem.expected.tok, item.tok)
		} else if elem.expected.lit != item.lit {
			t.Errorf("%d. %q literal mismatch, expected %q but got %q", pos, elem.input, elem.expected.lit, item.lit)
		}
	}
}
