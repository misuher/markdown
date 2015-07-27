package markdown_test

import (
	"strings"
	"testing"

	. "github.com/misuher/markdown/markparser"
)

func TestMarkdown(t *testing.T) {
	var tests = []struct {
		input    string
		expected string
	}{
		{"normal", "<p>normal</p>"},
		{"text with spaces", "<p>text with spaces</p>"},
		{"# header", "<h1>header</h1>"},
		{"## header", "<h2>header</h2>"},
		{"### header", "<h3>header</h3>"},
		{"#### header", "<h4>header</h4>"},
		{"*Italic*", "<em>Italic</em>"},
		{"**bold**", "<strong>bold</strong>"},
		{"* item1\n* item2\n* item3", "<ul><li>item1</li><li>item2</li><li>item3</li></ul>"},
		{"![Alt text](url)", "<img scr=\"url\" alt=\"Alt text\"/>"},
		{"[text](http://www.g.com)", "<a href=\"http://www.g.com\">text</a>"},
		{"> text", "<blockquote><p>text</p></blockquote>"},
		{"> line1\n> line2", "<blockquote><p>line1 line2</p></blockquote>"},
		{"\tcode", "<pre><code>code</code></pre>"},
		{"\t code", "<pre><code>code</code></pre>"},
		{"    code", "<pre><code>code</code></pre>"},
		{"***", "<hr />"},
		{"\x00", ""},
	}

	for _, elem := range tests {
		p := NewParser(strings.NewReader(elem.input))
		result := p.Markdown()

		if result != elem.expected {
			t.Errorf("expected %s but got %s", elem.expected, result)
		}
	}
}
