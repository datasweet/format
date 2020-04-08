package date

import (
	"strconv"
	"strings"
	"time"

	"github.com/datasweet/format/date/locales"
	"golang.org/x/text/language"
)

// Format the date with a unicode format
// https://www.unicode.org/reports/tr35/tr35-dates.html#Date_Field_Symbol_Table
func Format(lang language.Tag, t time.Time, format string) string {
	formatter := &formatter{
		lang: lang.String(),
		buff: new(strings.Builder),
	}
	return formatter.Format(t, format)
}

type formatter struct {
	lang string
	buff *strings.Builder
}

// Morning is from sunrise to 11:59 AM. Sunrise typically occurs around 6 AM.
// Noon is at 12:00 PM.
// Afternoon is from 12:01 PM to around 5:00 PM.
// Evening is from 5:01 PM to 8 PM, or around sunset.
// Night is from sunset to sunrise, so from 8:01 PM until 5:59 AM.
const (
	midnight = 0
	sunrise  = 6 * 60
	noon     = 12 * 60
	tea      = 17 * 60
	sunset   = 20 * 60
)

func (f *formatter) Format(t time.Time, format string) string {
	tokens, err := lex(format)
	if err != nil || len(tokens) < 2 {
		return format
	}

	pos := 0

	var (
		// cache
		year, mo, day  = t.Date()
		hour, min, sec = t.Clock()
		name, offset   = t.Zone()
	)

	// time.Month Jan = 1
	month := int(mo) - 1

	for i, tk := range tokens {
		if pos < tk.Pos && (i > 0 && tokens[i-1].Type != text) {
			f.buff.WriteString(format[pos:tk.Pos])
		}

		switch tk.Type {
		case text:
			f.buff.WriteString(tk.Value)

		case eraG:
			// Era
			era := 0
			if year > 0 {
				era = 1
			}
			f.appendLocalize(locales.Era, era, tk.Value, none)

		case yeary:
			// Year (abs)
			y := year
			if tk.Value == "yy" {
				y = y % 100
			}
			f.appendInt(y, len(tk.Value))

		case yearY:
			// Local week-numbering year
			wy, _ := t.ISOWeek()
			if tk.Value == "YY" {
				wy = wy % 100
			}
			f.appendInt(wy, len(tk.Value))

		case yearu:
			// Extended year
			f.appendInt(year, len(tk.Value))

		case yearU:
			// Cyclic year
			// Not implemented
			f.buff.WriteString(tk.Value)

		case yearr:
			// Related Gregorian year
			// Not implemented
			f.buff.WriteString(tk.Value)

		case quarterQ:
			// Quarter
			f.appendLocalize(locales.Quarter, month/3, tk.Value, candigit)

		case quarterq:
			// Stand-alone quarter
			f.appendLocalize(locales.Quarter, month/3, tk.Value, candigit|standalone)

		case monthM:
			// Month
			f.appendLocalize(locales.Month, month, tk.Value, candigit)

		case monthL:
			// Stand-alone month
			f.appendLocalize(locales.Month, month, tk.Value, candigit|standalone)

		case weekw:
			// Local week of year
			// Not implemented
			f.buff.WriteString(tk.Value)

		case weekW:
			// Week of month
			// Not implemented
			f.buff.WriteString(tk.Value)

		case dayd:
			// Day of month
			f.appendInt(day, len(tk.Value))

		case dayD:
			// Day of year
			f.appendInt(t.YearDay(), len(tk.Value))

		case dayF:
			// Day of week in month
			// Not implemented
			f.buff.WriteString(tk.Value)

		case dayg:
			// Modified Julian day
			// Not implemented
			f.buff.WriteString(tk.Value)

		case weekdayE:
			// Day of week
			f.appendLocalize(locales.Day, int(t.Weekday()), tk.Value, canshort)

		case weekdaye:
			// Local day of week
			// TODO: config to get the starting day of week
			// Here we starting at Monday
			if l := len(tk.Value); l <= 2 {
				localDayOfWeek := (int(t.Weekday()) - int(time.Monday) + 8) % 7
				if localDayOfWeek == 0 {
					localDayOfWeek = 7
				}
				f.appendInt(localDayOfWeek, l)
			} else {
				f.appendLocalize(locales.Day, int(t.Weekday()), tk.Value, candigit|canshort)
			}

		case weekdayc:
			// Stand-alone local day of week
			// Not implemented
			f.appendLocalize(locales.Day, int(t.Weekday()), tk.Value, candigit|canshort|standalone)

		case perioda:
			// AM, PM
			period := 0
			if hour >= 12 {
				period = 1
			}
			f.appendLocalize(locales.DayPeriod, period, tk.Value, none)

		case periodb:
			// AM, PM, noon, midnight
			// Keys: []string{"am", "pm", "morning1", "noon", "afternoon1", "evening1", "midnight", "night1"},
			period := 0
			if hour > 12 {
				period = 1
			} else if hour == 12 {
				period = 1
				if min == 0 {
					period = 3
				}
			} else if hour == 0 && min == 0 {
				period = 6 // midnight
			}
			f.appendLocalize(locales.DayPeriod, period, tk.Value, none)

		case periodB:
			// Flexible day period
			// Morning is from sunrise to 11:59 AM. Sunrise typically occurs around 6 AM.
			// Noon is at 12:00 PM.
			// Afternoon is from 12:01 PM to around 5:00 PM.
			// Evening is from 5:01 PM to 8 PM, or around sunset.
			// Night is from sunset to sunrise, so from 8:01 PM until 5:59 AM.
			// Keys: []string{"am", "pm", "morning1", "noon", "afternoon1", "evening1", "midnight", "night1"},
			period := 7 // night
			moment := hour*60 + min

			if moment == 0 {
				period = 6 // midnight
			} else if moment >= sunrise && moment < noon {
				period = 2 // morning
			} else if moment == noon {
				period = 3 // noon
			} else if moment > noon && moment <= tea {
				period = 4 // afternoon
			} else if moment > tea && moment <= sunset {
				period = 5 // evening
			}
			f.appendLocalize(locales.DayPeriod, period, tk.Value, none)

		case hourh:
			// Hour [1-12]
			hr := hour % 12
			if hr == 0 {
				hr = 12
			}
			f.appendInt(hr, len(tk.Value))

		case hourH:
			// Hour [0-23]
			f.appendInt(hour, len(tk.Value))

		case hourK:
			// Hour [0-11]
			f.appendInt(hour%12, len(tk.Value))

		case hourk:
			// Hour [1-24]
			hr := hour
			if hr == 0 {
				hr = 24
			}
			f.appendInt(hr, len(tk.Value))

		case hourj:
			// Localized hour w/ day period
			// Not implemented
			f.buff.WriteString(tk.Value)

		case hourJ:
			// Localized hour w/o day period
			// Not implemented
			f.buff.WriteString(tk.Value)

		case hourC:
			// Localized hour w/ day period
			// Not implemented
			f.buff.WriteString(tk.Value)

		case minutem:
			// Minute
			f.appendInt(min, len(tk.Value))

		case seconds:
			// Second
			f.appendInt(sec, len(tk.Value))

		case secondS:
			// Fraction of second
			f.appendNano(uint(t.Nanosecond()), len(tk.Value), false)

		case secondA:
			// Milliseconds in day
			// Not implemented
			f.buff.WriteString(tk.Value)

		case zonez:
			// Timezone (specific non-locat.)
			switch len(tk.Value) {
			case 1, 2, 3:
				if name != "" {
					f.buff.WriteString(name)
				} else {
					f.timezoneOffset(offset, tzGMT|tzOneDigit) // "O" equivalent
				}
			case 4:
				// "OOOO" equivalent
				f.timezoneOffset(offset, tzGMT|tzMin|tzSep)
			default:
				f.buff.WriteString(tk.Value)
			}

		case zoneZ:
			// Timezone (aliases)
			switch len(tk.Value) {
			case 1, 2, 3:
				// "xxxx" equivalent
				f.timezoneOffset(offset, tzMin|tzSec)
			case 4:
				// "OOOO" equivalent
				f.timezoneOffset(offset, tzGMT|tzMin|tzSep)
			case 5:
				// "XXXXX" equivalent
				f.timezoneOffset(offset, tzZ|tzMin|tzSec|tzSep)
			default:
				f.buff.WriteString(tk.Value)
			}

		case zoneO:
			// Timezone (GMT)
			switch len(tk.Value) {
			case 1:
				f.timezoneOffset(offset, tzGMT|tzOneDigit)
			case 4:
				f.timezoneOffset(offset, tzGMT|tzMin|tzSep)
			default:
				f.buff.WriteString(tk.Value)
			}

		case zonev:
			// Timezone (generic non-locat.)
			// Not implemented

		case zoneV:
			// Timezone (location)
			// Not implemented

		case zoneX:
			// Timezone (ISO-8601)
			switch len(tk.Value) {
			case 1:
				f.timezoneOffset(offset, tzZ)
			case 2:
				f.timezoneOffset(offset, tzZ|tzMin)
			case 3:
				f.timezoneOffset(offset, tzZ|tzMin|tzSep)
			case 4:
				f.timezoneOffset(offset, tzZ|tzMin|tzSec)
			case 5:
				f.timezoneOffset(offset, tzZ|tzMin|tzSec|tzSep)
			default:
				f.buff.WriteString(tk.Value)
			}

		case zonex:
			// Timezone (ISO-8601 w/o Z)
			switch len(tk.Value) {
			case 1:
				f.timezoneOffset(offset, tznone)
			case 2:
				f.timezoneOffset(offset, tzMin)
			case 3:
				f.timezoneOffset(offset, tzMin|tzSep)
			case 4:
				f.timezoneOffset(offset, tzMin|tzSec)
			case 5:
				f.timezoneOffset(offset, tzMin|tzSec|tzSep)
			default:
				f.buff.WriteString(tk.Value)
			}

		case timestampt:
			f.buff.WriteString(strconv.FormatInt(t.UnixNano()/int64(time.Millisecond), 10))
		}

		pos = tk.Pos + len(tk.Value)
	}

	return f.buff.String()

}

// Copyright 2010 The Go Authors. All rights reserved.
// time.go appendInt
func (f *formatter) appendInt(x int, width int) {
	u := uint(x)
	if x < 0 {
		f.buff.WriteByte('-')
		u = uint(-x)
	}

	// Assemble decimal in reverse order.
	var buf [20]byte
	i := len(buf)
	for u >= 10 {
		i--
		q := u / 10
		buf[i] = byte('0' + u - q*10)
		u = q
	}
	i--
	buf[i] = byte('0' + u)

	// Add 0-padding.
	// Add 0-padding.
	for w := len(buf) - i; w < width; w++ {
		f.buff.WriteByte('0')
	}

	f.buff.Write(buf[i:])
}

// Copyright 2010 The Go Authors. All rights reserved.
// time.go formatNano appends a fractional second, as nanoseconds, to b
// and returns the result.
func (f *formatter) appendNano(nanosec uint, n int, trim bool) {
	u := nanosec
	var buf [9]byte
	for start := len(buf); start > 0; {
		start--
		buf[start] = byte(u%10 + '0')
		u /= 10
	}

	if n > 9 {
		n = 9
	}
	if trim {
		for n > 0 && buf[n-1] == '0' {
			n--
		}
		if n == 0 {
			return
		}
	}
	//sb.WriteByte('.')
	f.buff.Write(buf[:n])
}

type localizeMode uint8

const (
	none       localizeMode = 0
	standalone localizeMode = 1 << iota
	candigit
	canshort
)

func (f *formatter) appendLocalize(field locales.FieldType, x int, pattern string, mode localizeMode) {
	width := locales.Abbreviated
	switch l := len(pattern); l {
	case 1, 2:
		if mode&candigit == candigit {
			f.appendInt(x+1, l)
			return
		}
		width = locales.Abbreviated
		if mode&standalone == standalone {
			width = locales.AbbreviatedStandalone
		}
	case 3:
		width = locales.Abbreviated
		if mode&standalone == standalone {
			width = locales.AbbreviatedStandalone
		}
	case 4:
		width = locales.Wide
		if mode&standalone == standalone {
			width = locales.WideStandalone
		}
	case 5:
		width = locales.Narrow
		if mode&standalone == standalone {
			width = locales.NarrowStandalone
		}
	case 6:
		if mode&canshort != canshort {
			f.buff.WriteString(pattern)
			return
		}
		width = locales.Short
		if mode&standalone == standalone {
			width = locales.ShortStandalone
		}
	default:
		f.buff.WriteString(pattern)
		return
	}

	if loc := locales.Localize(f.lang, field, locales.Width(width), x); len(loc) > 0 {
		f.buff.WriteString(loc)
		return
	}

	f.buff.WriteString(pattern)
}

type tzMode uint8

const (
	tznone     tzMode = 0
	tzGMT      tzMode = 1 << iota // Write GMT timezone name
	tzZ                           // The ISO8601 UTC indicator "Z" is used when local time offset is 0
	tzSep                         // Separate offset with ':
	tzMin                         // Include minutes
	tzSec                         // IncludeSeconde
	tzOneDigit                    // use one digit
)

func (f *formatter) timezoneOffset(offset int, mode tzMode) {
	if offset == 0 && mode&tzZ == tzZ {
		f.buff.WriteByte('Z')
		return
	}

	if mode&tzGMT == tzGMT {
		f.buff.WriteString("GMT")
	}

	zone := offset / 60
	absoffset := offset
	if zone < 0 {
		f.buff.WriteByte('-')
		zone = -zone
		absoffset = -absoffset
	} else {
		f.buff.WriteByte('+')
	}

	digit := 2
	if mode&tzOneDigit == tzOneDigit {
		digit = 1
	}
	f.appendInt(zone/60, digit)

	if mode&tzMin == tzMin {
		if mode&tzSep == tzSep {
			f.buff.WriteByte(':')
		}
		f.appendInt(zone%60, digit)

		if mode&tzSec == tzSec {
			if sec := absoffset % 60; sec > 0 {
				if mode&tzSep == tzSep {
					f.buff.WriteByte(':')
				}
				f.appendInt(sec, digit)
			}
		}
	}
}
