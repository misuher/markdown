package markdown

import (
	"bytes"
	"fmt"
	"io"
)

//Parser wraps a Scanner to get the tokens and a buffer to write the formatted output
type Parser struct {
	s         *Scanner
	lastItem  Item
	unscanned bool
	output    bytes.Buffer
}

//NewParser initializes and return a new parser
func NewParser(r io.Reader) *Parser {
	return &Parser{s: NewScanner(r)}
}

//Markdown run the parser and when it finished return the result
func (p *Parser) Markdown() string {
	err := p.nextState(p.stateParse())
	if err != nil {
		panic(err)
	}
	return p.output.String()
}

//stateFn is a type returned by all the parser functions as we want it to be a state machine so one state lead to
//the next just returning a funtion representing the next state
type stateFn func(*Parser) stateFn

//nextState execute the state machine calling the next state recursevely
func (p *Parser) nextState(state stateFn) stateFn {
	newstate := state
	if newstate != nil {
		p.nextState(newstate)
	}
	return nil
}

//stateParse get he next token with the scanner and choose the next state based on it
func (p *Parser) stateParse() stateFn {
	item := p.getNextItem()

	switch item.Tok {
	case HASH, DOUBLEHASH, TRIPLEHASH, QUADHASH:
		return p.stateHash(item.Tok)
	case ASTERISK, DOUBLEASTERISK, TRIPLEASTERISK:
		return p.stateAsterisk(item.Tok)
	case GT:
		return p.stateGT()
	case TAB:
		return p.stateTab()
	case WS:
		return p.stateParse()
	case NL:
		return p.stateNewLine()
	case EXCLAMATION:
		return p.stateExclamation()
	case SQUAREOPEN:
		return p.stateSquareOpen()
	case EOF:
		return nil
	default:
		return p.stateLiteral(item)
	}
}

func (p *Parser) unscan() {
	p.unscanned = true
}

func (p *Parser) getNextItem() Item {
	if p.unscanned {
		p.unscanned = false
		return p.lastItem
	}
	return p.s.Scan()
}

func (p *Parser) stateHash(tok Token) stateFn {
	//after a hash chain we must have a withespace and a literal
	p.lastItem = p.s.Scan()
	if p.lastItem.Tok != WS {
		//if there is no WS treat it all like a literal
		p.output.WriteString(string(tok) + p.lastItem.Lit)
		return p.stateParse()
	}
	p.lastItem = p.s.Scan()
	if p.lastItem.Tok != LITERAL {
		return nil
	}

	switch tok {
	case HASH:
		p.output.WriteString("<h1>" + p.lastItem.Lit + "</h1>")
		break
	case DOUBLEHASH:
		p.output.WriteString("<h2>" + p.lastItem.Lit + "</h2>")
		break
	case TRIPLEHASH:
		p.output.WriteString("<h3>" + p.lastItem.Lit + "</h3>")
		break
	case QUADHASH:
		p.output.WriteString("<h4>" + p.lastItem.Lit + "</h4>")
		break
	}
	return p.stateParse()
}

func (p *Parser) stateAsterisk(tok Token) stateFn {

	switch tok {
	case ASTERISK:
		return p.stateSingleAsterisk()

	case DOUBLEASTERISK:
		p.lastItem = p.s.Scan()
		p.output.WriteString("<strong>" + p.lastItem.Lit + "</strong>")
		p.lastItem = p.s.Scan()
		if p.lastItem.Tok != DOUBLEASTERISK {
			return nil
		}
		break

	case TRIPLEASTERISK:
		p.output.WriteString("<hr />")
	}
	return p.stateParse()
}

func (p *Parser) stateSingleAsterisk() stateFn {
	p.lastItem = p.s.Scan()

	switch p.lastItem.Tok {
	case WS:
		p.output.WriteString("<ul>")
		p.lastItem = p.s.Scan()
		for {
			if p.lastItem.Tok != LITERAL {
				break
			}
			p.output.WriteString("<li>" + p.lastItem.Lit + "</li>")

			p.lastItem = p.s.Scan()
			if p.lastItem.Tok != NL {
				fmt.Println("making list: not a NL")
				break
			}

			p.lastItem = p.s.Scan()
			if p.lastItem.Tok != ASTERISK {
				fmt.Println("making list: not an ASTERISK")
				break
			}

			p.lastItem = p.s.Scan()
			if p.lastItem.Tok != WS {
				fmt.Println("making list: not a WS")
				break
			}
		}
		p.output.WriteString("</ul>")
		break

	case LITERAL:
		p.output.WriteString("<em>" + p.lastItem.Lit + "</em>")
		p.lastItem = p.s.Scan()
		if p.lastItem.Tok != ASTERISK {
			return nil
		}
		break
	}
	return p.stateParse()
}

func (p *Parser) stateGT() stateFn {
	p.output.WriteString("<blockquote><p>")
	for {
		p.lastItem = p.s.Scan()
		if p.lastItem.Tok != WS {
			return nil
		}

		p.lastItem = p.s.Scan()
		if p.lastItem.Tok != LITERAL {
			return nil
		}

		p.lastItem = p.s.Scan()
		if p.lastItem.Tok != NL {
			return nil
		}

		p.lastItem = p.s.Scan()
		if p.lastItem.Tok != GT {
			break
		}
	}
	p.output.WriteString("</p></blockquote>")
	return p.stateParse()
}

func (p *Parser) stateTab() stateFn {
	p.lastItem = p.s.Scan()

	switch p.lastItem.Tok {
	case WS:
		return p.stateTab()

	case LITERAL:
		p.output.WriteString("<pre><code>" + p.lastItem.Lit + "</code></pre>")
	}
	return p.stateParse()
}

func (p *Parser) stateLiteral(it Item) stateFn {
	p.output.WriteString("<p>" + it.Lit + "</p>")
	return p.stateParse()
}

func (p *Parser) stateNewLine() stateFn {
	p.output.WriteString("<br>")
	return p.stateParse()
}

func (p *Parser) stateExclamation() stateFn {
	p.lastItem = p.s.Scan()
	if p.lastItem.Tok != SQUAREOPEN {
		return nil
	}

	itemText := p.s.Scan()
	if itemText.Tok != LITERAL {
		return nil
	}

	p.lastItem = p.s.Scan()
	if p.lastItem.Tok != SQUARECLOSE {
		return nil
	}

	p.lastItem = p.s.Scan()
	if p.lastItem.Tok != PARANOPEN {
		return nil
	}

	itemURL := p.s.Scan()
	if itemURL.Tok != LITERAL {
		return nil
	}
	p.output.WriteString("<img scr=\"" + itemURL.Lit + "\" alt=\"" + itemText.Lit + "\"/>")

	p.lastItem = p.s.Scan()
	if p.lastItem.Tok != PARANCLOSE {
		return nil
	}
	return p.stateParse()
}

func (p *Parser) stateSquareOpen() stateFn {
	itemText := p.s.Scan()
	if itemText.Tok != LITERAL {
		return nil
	}

	p.lastItem = p.s.Scan()
	if p.lastItem.Tok != SQUARECLOSE {
		return nil
	}

	p.lastItem = p.s.Scan()
	if p.lastItem.Tok != PARANOPEN {
		return nil
	}

	itemURL := p.s.Scan()
	if itemURL.Tok != LITERAL {
		return nil
	}
	p.output.WriteString("<a href=\"" + itemURL.Lit + "\">" + itemText.Lit + "</a>")

	p.lastItem = p.s.Scan()
	if p.lastItem.Tok != PARANCLOSE {
		return nil
	}
	return p.stateParse()
}
