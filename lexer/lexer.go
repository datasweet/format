package lexer

import (
	"errors"
	"fmt"
	"strings"
	"unicode/utf8"
)

// Lexer is our lexer
type Lexer interface {
	// Next returns the next rune in the input.
	Next() rune

	// Peek returns but does not consume the next rune in the input.
	Peek() rune

	// Rewind steps back one rune. Can only be called once per call of next.
	Rewind()

	// Current returns the current value being analyzed
	Current() string

	// Emit a  new token
	Emit(typ TokenType)

	// Emit a  new token with a specific value
	EmitValue(typ TokenType, value string)

	// Ignore skigs over the pending input before this point.
	Ignore()

	// Accept consumes the next rune if the rune is in valid string
	Accept(valid string) bool

	// Take continues to consume the next rune while rune founded is in valid string
	Take(valid string)

	// Errorf to create a error state
	Errorf(format string, args ...interface{}) StateFn
}

// StateFn describes our state
type StateFn func(Lexer) StateFn

// TokenType to identify the token
type TokenType uint

// Token descrines a token
type Token struct {
	Type  TokenType
	Value string
	Pos   int
}

// Lex main entry point to lex an input string
func Lex(input string, lexRoot StateFn) ([]Token, error) {
	if lexRoot == nil {
		return nil, errors.New("missing main lex function")
	}
	l := &lexer{
		input: input,
	}
	for state := lexRoot; state != nil; {
		state = state(l)
	}
	return l.tokens, l.err
}

// main implementation
type lexer struct {
	input  string  // the string being scanned
	pos    int     // current position in the input
	start  int     // start position of this token
	width  int     // width of last rune read from input
	tokens []Token // slice of scanned tokens
	err    error   // last error
}

func (l *lexer) Next() rune {
	if l.pos >= len(l.input) {
		l.width = 0
		return -1
	}
	r, w := utf8.DecodeRuneInString(l.input[l.pos:])
	l.width = w
	l.pos += l.width
	return r
}

func (l *lexer) Peek() rune {
	r := l.Next()
	l.Rewind()
	return r
}

func (l *lexer) Rewind() {
	l.pos -= l.width
}

func (l *lexer) Current() string {
	return l.input[l.start:l.pos]
}

func (l *lexer) Emit(typ TokenType) {
	l.EmitValue(typ, l.Current())
}

func (l *lexer) EmitValue(typ TokenType, value string) {
	l.tokens = append(l.tokens, Token{
		Type:  typ,
		Value: value,
		Pos:   l.start,
	})
	l.start = l.pos
}

func (l *lexer) Ignore() {
	l.start = l.pos
}

func (l *lexer) Accept(valid string) bool {
	if strings.ContainsRune(valid, l.Next()) {
		return true
	}
	l.Rewind()
	return false
}

func (l *lexer) Take(valid string) {
	for strings.ContainsRune(valid, l.Next()) {
	}
	l.Rewind()
}

func (l *lexer) Errorf(format string, args ...interface{}) StateFn {
	l.err = &syntaxError{
		message: fmt.Sprintf(format, args...),
		input:   l.input,
		pos:     l.start,
	}
	return nil
}
