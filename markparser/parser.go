package markdown

import "io"

//Parser wraps a Scanner to get the tokens and a buffer to write the formatted output
type Parser struct {
	s      *Scanner
	output string
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
	return p.output
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
	item := p.s.Scan()

	switch item.tok {
	case HASH, DOUBLEHASH, TRIPLEHASH, QUADHASH:
		return p.stateHash(item.tok)
	case ASTERISK, DOUBLEASTERISK:
		return p.stateAsterisk(item.tok)
	case TAB:
		return p.stateTab()
	case NL, WS:
		return p.stateParse()
	case LITERAL:
		return p.stateLiteral()
	case EOF:
		return nil
	}
}

func (p *Parser) stateHash(tok Token) stateFn {

}

func (p *Parser) stateAsterisk(tok Token) stateFn {

}

func (p *Parser) stateTab() stateFn {

}

func (p *Parser) stateLiteral() stateFn {

}
