package date_test

import (
	"strings"
	"testing"
	"time"

	"github.com/datasweet/format/date"
	"github.com/stretchr/testify/assert"
	"golang.org/x/text/language"
)

type FormatTest struct {
	name   string
	format string
	result string
}

func notImplemented(pattern string) string {
	var r []string
	for i := 1; i <= 7; i++ {
		r = append(r, strings.Repeat(pattern, i))
	}
	return strings.Join(r, " ")
}

func TestPattern(t *testing.T) {
	patternTests := []FormatTest{
		// https://www.unicode.org/reports/tr35/tr35-dates.html#Date_Field_Symbol_Table
		{"Era", "G", "AD AD AD Anno Domini A GGGGGG GGGGGGG"},
		{"Year", "y", "2009 09 2009 2009 02009 002009 0002009"},
		{"Year", "Y", "2009 09 2009 2009 02009 002009 0002009"},
		{"Year", "u", "2009 2009 2009 2009 02009 002009 0002009"},
		{"Year", "U", notImplemented("U")},
		{"Year", "r", notImplemented("r")},
		{"Quarter", "Q", "1 01 Q1 1st quarter 1 QQQQQQ QQQQQQQ"},
		{"Quarter", "q", "1 01 Q1 1st quarter 1 qqqqqq qqqqqqq"},
		{"Month", "M", "2 02 Feb February F MMMMMM MMMMMMM"},
		{"Month", "L", "2 02 Feb February F LLLLLL LLLLLLL"},
		{"Week", "w", notImplemented("w")},
		{"Week", "W", notImplemented("W")},
		{"Day", "d", "4 04 004 0004 00004 000004 0000004"},
		{"Day", "D", "35 35 035 0035 00035 000035 0000035"},
		{"Day", "F", notImplemented("F")},
		{"Day", "g", notImplemented("g")},
		{"Week day", "E", "Wed Wed Wed Wednesday W We EEEEEEE"},
		{"Week day", "e", "3 03 Wed Wednesday W We eeeeeee"},
		{"Week day", "c", "4 04 Wed Wednesday W We ccccccc"},
		{"Period", "a", "PM PM PM PM p aaaaaa aaaaaaa"},
		{"Period", "b", "PM PM PM PM p bbbbbb bbbbbbb"},
		{"Period", "B", "at night at night at night at night at night BBBBBB BBBBBBB"},
		{"Hour", "h", "9 09 009 0009 00009 000009 0000009"},
		{"Hour", "H", "21 21 021 0021 00021 000021 0000021"},
		{"Hour", "K", "9 09 009 0009 00009 000009 0000009"},
		{"Hour", "k", "21 21 021 0021 00021 000021 0000021"},
		{"Hour", "j", notImplemented("j")},
		{"Hour", "J", notImplemented("J")},
		{"Hour", "C", notImplemented("C")},
		{"Minute", "m", "0 00 000 0000 00000 000000 0000000"},
		{"Second", "s", "57 57 057 0057 00057 000057 0000057"},
		{"Second", "S", "0 01 012 0123 01234 012345 0123456"},
		{"Second", "A", notImplemented("A")},
	}

	// from time_test.go / format_test.go
	// Force location
	loc, err := time.LoadLocation("America/Los_Angeles")
	assert.NoError(t, err)
	assert.NotNil(t, loc)

	// The numeric time represents Thu Feb  4 21:00:57.012345600 PST 2009
	tm := time.Unix(0, 1233810057012345600).In(loc)

	for _, test := range patternTests {
		var r []string
		for i := 1; i <= 7; i++ {
			r = append(r, date.Format(language.English, tm, strings.Repeat(test.format, i)))
		}
		result := strings.Join(r, " ")
		if result != test.result {
			t.Errorf("%s %s expected %q got %q", test.name, test.format, test.result, result)
		}
	}
}

func TestFormat(t *testing.T) {
	formatTests := []FormatTest{
		{"ANSIC", "EEE MMM  d HH:mm:ss yyyy", "Wed Feb  4 21:00:57 2009"},
		{"UnixDate", "EEE MMM  d HH:mm:ss z yyyy", "Wed Feb  4 21:00:57 PST 2009"},
		{"RubyDate", "EEE MMM dd HH:mm:ss XX yyyy", "Wed Feb 04 21:00:57 -0800 2009"},
		{"RFC822", "dd MMM yy HH:mm z", "04 Feb 09 21:00 PST"},
		{"RFC850", "EEEE, dd-MMM-yy HH:mm:ss z", "Wednesday, 04-Feb-09 21:00:57 PST"},
		{"RFC1123", "EEE, dd MMM yyyy HH:mm:ss z", "Wed, 04 Feb 2009 21:00:57 PST"},
		{"RFC1123Z", "EEE, dd MMM yyyy HH:mm:ss XXXX", "Wed, 04 Feb 2009 21:00:57 -0800"},
		{"RFC3339", "yyyy-MM-ddTHH:mm:ssXXX", "2009-02-04T21:00:57-08:00"},
		{"RFC3339Nano", "yyyy-MM-ddTHH:mm:ss.SSSSSSSSSXXX", "2009-02-04T21:00:57.012345600-08:00"},
		{"Kitchen", "h:mma", "9:00PM"},
		{"AM/PM", "ha", "9PM"},
		{"Flexible days periods", "h B", "9 at night"},
		{"two-digit year", "yy MM dd", "09 02 04"},
		// Three-letter months and days must not be followed by lower-case letter.
		//{"Janet", "Hi Janet, the Month is MMMM", "Hi Janet, the Month is February"},
		// Time stamps, Fractional seconds.
		{"Stamp", "MMM  d HH:mm:ss", "Feb  4 21:00:57"},
		{"StampMilli", "MMM  d HH:mm:ss.SSS", "Feb  4 21:00:57.012"},
		{"StampMicro", "MMM  d HH:mm:ss.SSSSSS", "Feb  4 21:00:57.012345"},
		{"StampNano", "MMM  d HH:mm:ss.SSSSSSSSS", "Feb  4 21:00:57.012345600"},
		{"YearDay", "MMM  d DDD  DD d", "Feb  4 035  35 4"},
		{"Timestamp", "t", "1233810057012"},
	}

	// from time_test.go / format_test.go
	// Force location
	loc, err := time.LoadLocation("America/Los_Angeles")
	assert.NoError(t, err)
	assert.NotNil(t, loc)

	// The numeric time represents Thu Feb  4 21:00:57.012345600 PST 2009
	tm := time.Unix(0, 1233810057012345600).In(loc)

	for _, test := range formatTests {
		result := date.Format(language.English, tm, test.format)
		if result != test.result {
			t.Errorf("%s - lang en expected %q got %q", test.name, test.result, result)
		}
	}
}
