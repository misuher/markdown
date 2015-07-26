package markdown

//Token type identifies the token as an integer
type Token int

const (
	ILLEGAL = iota
	EOF

	HASH
	DOUBLEHASH
	TRIPLEHASH
	QUADHASH
	ASTERISK
	DOUBLEASTERISK
	GT
	SQUAREOPEN
	SQUARECLOSE
	PARANOPEN
	PARANCLOSE
	EXCLAMATION

	LITERAL

	TAB
	WS
	NL
)
