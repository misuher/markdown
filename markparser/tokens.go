package markdown

type tokens int

const (
	TOKEN_ERROR = iota
	TOKEN_EOF

	TOKEN_HASH
	TOKEN_ASTERIX
	TOKEN_GT
	TOKEN_SQUARE_OPEN
	TOKEN_SQUARE_CLOSE
	TOKEN_PARAN_OPEN
	TOKEN_PARAN_CLOSE
	TOKEN_EXCLAMATION

	TOKEN_STRING

	TOKEN_WITHESPACE
	TOKEN_NEWLINE
)

const (
	EOF rune = 0

	HASH         string = "#"
	ASTERIX      string = "*"
	GT           string = ">"
	SQUARE_OPEN  string = "["
	SQUARE_CLOSE string = "]"
	PARAN_OPEN   string = "("
	PARAN_CLOSE  string = ")"
	EXCLAMATION  string = "!"

	WHITESPACE string = " "
	NEWLINE    string = "\n"
)
