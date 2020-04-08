package date

import (
	"strings"

	"github.com/datasweet/format/lexer"
)

const (
	eof lexer.TokenType = iota
	text

	// https://www.unicode.org/reports/tr35/tr35-dates.html#Date_Field_Symbol_Table
	eraG
	yeary
	yearY
	yearu
	yearU
	yearr
	quarterQ
	quarterq
	monthM
	monthL
	weekw
	weekW
	dayd
	dayD
	dayF
	dayg
	weekdayE
	weekdaye
	weekdayc
	perioda
	periodb
	periodB
	hourh
	hourH
	hourK
	hourk
	hourj
	hourJ
	hourC
	minutem
	seconds
	secondS
	secondA
	zonez
	zoneZ
	zoneO
	zonev
	zoneV
	zoneX
	zonex

	// Specific
	timestampt // timestamp in millisecond, compliant with json 't'
)

// Lex to lex a moment format
func lex(input string) ([]lexer.Token, error) {
	return lexer.Lex(input, lexRoot)
}

func lexRoot(l lexer.Lexer) lexer.StateFn {
	switch r := l.Next(); r {
	case -1:
		l.Emit(eof)
		return nil
	case '\'', '"':
		l.Rewind()
		return lexQuote
	case 'a':
		// AM, PM
		l.Take("a")
		l.Emit(perioda)
	case 'b':
		// AM, PM, noon, midnight
		l.Take("b")
		l.Emit(periodb)
	case 'c':
		// Stand-alone local day of week
		l.Take("c")
		l.Emit(weekdayc)
	case 'd':
		// Day of month
		l.Take("d")
		l.Emit(dayd)
	case 'e':
		// Local day of week
		l.Take("e")
		l.Emit(weekdaye)
	case 'g':
		// Modified Julian day
		l.Take("g")
		l.Emit(dayg)
	case 'h':
		// Hour [1-12]
		l.Take("h")
		l.Emit(hourh)
	case 'j':
		// Localized hour w/ day period
		l.Take("j")
		l.Emit(hourj)
	case 'k':
		// Hour [1-24]
		l.Take("k")
		l.Emit(hourk)
	case 'm':
		// Minute
		l.Take("m")
		l.Emit(minutem)
	case 'q':
		// Stand-alone quarter
		l.Take("q")
		l.Emit(quarterq)
	case 'r':
		// Related Gregorian year
		l.Take("r")
		l.Emit(yearr)
	case 's':
		// Second
		l.Take("s")
		l.Emit(seconds)
	case 't':
		// timestamp in millisecond
		// custom
		l.Take("t")
		l.Emit(timestampt)
	case 'u':
		// Extended year
		l.Take("u")
		l.Emit(yearu)
	case 'v':
		// Timezone (generic non-locat.)
		l.Take("v")
		l.Emit(zonev)
	case 'w':
		// Local week of year
		l.Accept("w")
		l.Emit(weekw)
	case 'x':
		// Timezone (ISO-8601 w/o Z)
		l.Take("x")
		l.Emit(zonex)
	case 'y':
		// Year (abs)
		l.Take("y")
		l.Emit(yeary)
	case 'z':
		// Timezone (specific non-locat.)
		l.Take("z")
		l.Emit(zonez)
	case 'A':
		// Milliseconds in day
		l.Take("A")
		l.Emit(secondA)
	case 'B':
		// Flexible day period
		l.Take("B")
		l.Emit(periodB)
	case 'C':
		// Localized hour w/ day period
		l.Take("C")
		l.Emit(hourC)
	case 'D':
		// Day of year
		l.Take("D")
		l.Emit(dayD)
	case 'E':
		// Day of week
		l.Take("E")
		l.Emit(weekdayE)
	case 'F':
		// Day of week in month
		l.Take("F")
		l.Emit(dayF)
	case 'G':
		// Era
		l.Take("G")
		l.Emit(eraG)
	case 'H':
		// Hour [0-23]
		l.Take("H")
		l.Emit(hourH)
	case 'J':
		// Localized hour w/o day period
		l.Take("J")
		l.Emit(hourJ)
	case 'K':
		// Hour [0-11]
		l.Take("K")
		l.Emit(hourK)
	case 'L':
		// Stand-alone month
		l.Take("L")
		l.Emit(monthL)
	case 'M':
		// Month
		l.Take("M")
		l.Emit(monthM)
	case 'O':
		// Timezone (GMT)
		l.Take("O")
		l.Emit(zoneO)
	case 'Q':
		// Quarter
		l.Take("Q")
		l.Emit(quarterQ)
	case 'S':
		// Fraction of second
		l.Take("S")
		l.Emit(secondS)
	case 'U':
		// Cyclic year
		l.Take("U")
		l.Emit(yearU)
	case 'V':
		// Timezone (location)
		l.Take("V")
		l.Emit(zoneV)
	case 'W':
		// Week of month
		l.Emit(weekW)
	case 'X':
		// Timezone (ISO-8601)
		l.Take("X")
		l.Emit(zoneX)
	case 'Y':
		// Local week-numbering year
		l.Take("Y")
		l.Emit(yearY)
	case 'Z':
		// Timezone (aliases)
		l.Take("Z")
		l.Emit(zoneZ)
	default:
		// return l.Errorf("unrecognized character: %#U", r)
		l.Ignore()
	}
	return lexRoot
}

func lexQuote(l lexer.Lexer) lexer.StateFn {
	quote := l.Next()
Loop:
	for {
		switch l.Next() {
		case '\\':
			if r := l.Next(); r != -1 && r != '\n' {
				break
			}
			fallthrough
		case -1:
			return l.Errorf("unterminated string")
		case quote:
			break Loop
		}
	}
	q := string(quote)
	value := strings.Trim(l.Current(), q)
	value = strings.Replace(value, "\\"+q, q, -1)
	value = strings.Replace(value, "\\n", "\n", -1)
	value = strings.Replace(value, "\\t", "\t", -1)
	l.EmitValue(text, value)
	return lexRoot
}

func emitDateToken(l lexer.Lexer, token lexer.TokenType) {
	r := l.Next()
	if !lexer.IsAlphaNumeric(r) {
		l.Rewind()
		l.Emit(token)
		return
	}
	l.Emit(text)
}
