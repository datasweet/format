package lexer

import (
	"unicode"
)

// IsSpace checks if a rune is a space character
func IsSpace(r rune) bool {
	return r == ' ' || r == '\t' || r == '\r' || r == '\n'
}

// IsAlphaNumeric checks if a rune is a digit or a letter
func IsAlphaNumeric(r rune) bool {
	return unicode.IsDigit(r) || unicode.IsLetter(r)
}
