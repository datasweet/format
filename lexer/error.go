package lexer

import (
	"fmt"
	"strings"
)

type syntaxError struct {
	message string
	input   string
	pos     int
}

func (e *syntaxError) Error() string {
	snippet := ""
	if len(e.input) > 0 {
		snippet = fmt.Sprintf("\n%s\n%s", e.input, strings.Repeat("-", e.pos)+"^")
	}
	return e.message + snippet
}

func (e *syntaxError) at(t Token) error {
	e.pos = t.Pos
	return e
}
