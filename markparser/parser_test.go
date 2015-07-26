package markdown

import (
	"strings"
	"testing"
)

func TestMarkdown(t *testing.T) {
	var tests = []struct {
		input    string
		expected string
	}{
		{"# header", "<h1>header</h1>"},
		{"## header", "<h2>header</h2>"},
		{"### header", "<h3>header</h3>"},
		{"#### header", "<h4>header</h4>"},
		{"\x00", ""},
	}

	for pos, elem := range tests {
		p := NewParser(strings.NewReader(elem.input))
		result := p.Markdown()

		if result != elem.expected {
			t.Errorf("expected %s but got %s", elem.expected, result)
		}
	}
}
