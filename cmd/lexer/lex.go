package lexer

import (
	"bytes"
	"unicode"
)

const (  
	none     = iota
	word     = iota
	delimetr = iota
)

type worker func(*Lex)

func addBuffer(l *Lex) {
	l.buffer.WriteByte(l.current_byte)
}

func addLexeme(l *Lex) {
	l.lexeme_list = append(l.lexeme_list, l.buffer.String())
	l.buffer = bytes.Buffer{}
}

var (
	stateTable [3][3]worker
)

func init() {
	stateTable[none][word]         = addBuffer
	stateTable[none][delimetr]     = nil
	stateTable[word][word]         = addBuffer
	stateTable[word][delimetr]     = addLexeme
	stateTable[delimetr][word]     = addBuffer
	stateTable[delimetr][delimetr] = nil
}


type Lex struct {
	buffer bytes.Buffer
	current_byte byte
	current_state int
	new_state int
	lexeme_list []string
}

func NewLex() *Lex {
	return &Lex{
		buffer: bytes.Buffer{},
		current_state: none,
		new_state: none,
		lexeme_list: make([]string, 0),
	}
}

func (l *Lex) isAllowedChar() bool {
	return unicode.IsLetter(rune(l.current_byte)) ||
		   unicode.IsPunct(rune(l.current_byte))  ||
		   unicode.IsDigit(rune(l.current_byte))
}

func (l *Lex) DefineState() {
	if l.isAllowedChar() {
		l.new_state = word
	} else if unicode.IsSpace(rune(l.current_byte)) {
		l.new_state = delimetr
    }
}

func (l *Lex) Analyze(line []byte) { 
	for _, c := range line {
		l.current_byte = c
		l.DefineState()
		stateMaker := l.getStateMaker()
		if stateMaker != nil {
			stateMaker(l)
		}
		l.updateState()
	}
	if l.buffer.Len() != 0 {
		addLexeme(l)
	}
}

func (l *Lex) updateState() {
	l.current_state = l.new_state
	l.new_state = none
}

func (l *Lex) GetLexemes() []string {
	return l.lexeme_list
}

func (l *Lex) getStateMaker() worker {
	return stateTable[l.current_state][l.new_state]
}

func (l *Lex) ClearLexer() {
	l.buffer.Reset()
	l.current_state = none
	l.new_state = none
	l.lexeme_list = l.lexeme_list[:0]
}