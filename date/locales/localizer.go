package locales

import (
	"strings"

	"golang.org/x/text/message/catalog"
)

// Width
type Width uint8

const (
	Abbreviated Width = iota // Len = 3
	AbbreviatedStandalone
	Wide // Len = 4
	WideStandalone
	Narrow // Len = 5
	NarrowStandalone
	Short // Len = 6
	ShortStandalone
)

// FieldType describes a type of field
type FieldType uint8

const (
	Era FieldType = iota
	Quarter
	Month
	Day
	DayPeriod
)

// Field describe a date field in cldr
// https://github.com/unicode-cldr/cldr-dates-full/tree/master/main
type Field struct {
	CldrNode string
	Widths   map[Width]string
	Keys     []string
}

func (f Field) Key(width Width, value int) (string, bool) {
	if value >= 0 && value < len(f.Keys) {
		key := []string{f.CldrNode}
		switch width {
		case Abbreviated:
			key = append(key, "abbr")
		case AbbreviatedStandalone:
			key = append(key, "abbr_sa")
		case Narrow:
			key = append(key, "narrow")
		case NarrowStandalone:
			key = append(key, "narrow_sa")
		case Wide:
			key = append(key, "wide")
		case WideStandalone:
			key = append(key, "wide_sa")
		case Short:
			key = append(key, "short")
		case ShortStandalone:
			key = append(key, "short_sa")
		default:
			return "", false
		}
		key = append(key, f.Keys[value])
		return strings.Join(key, "."), true
	}
	return "", false
}

// Fields gets all fields with infos
var Fields = map[FieldType]Field{
	Era: Field{
		CldrNode: "eras",
		Widths: map[Width]string{
			Wide:        "eraNames",
			Narrow:      "eraNarrow",
			Abbreviated: "eraAbbr",
		},
		Keys: []string{"0", "1"},
	},
	Quarter: Field{
		CldrNode: "quarters",
		Widths: map[Width]string{
			Wide:                  "format.wide",
			Narrow:                "format.narrow",
			Abbreviated:           "format.abbreviated",
			WideStandalone:        "stand-alone.wide",
			NarrowStandalone:      "stand-alone.narrow",
			AbbreviatedStandalone: "stand-alone.abbreviated",
		},
		Keys: []string{"1", "2", "3", "4"},
	},
	Month: Field{
		CldrNode: "months",
		Widths: map[Width]string{
			Wide:                  "format.wide",
			Narrow:                "format.narrow",
			Abbreviated:           "format.abbreviated",
			WideStandalone:        "stand-alone.wide",
			NarrowStandalone:      "stand-alone.narrow",
			AbbreviatedStandalone: "stand-alone.abbreviated",
		},
		Keys: []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12"},
	},
	Day: Field{
		CldrNode: "days",
		Widths: map[Width]string{
			Wide:                  "format.wide",
			Narrow:                "format.narrow",
			Abbreviated:           "format.abbreviated",
			Short:                 "format.short",
			WideStandalone:        "stand-alone.wide",
			NarrowStandalone:      "stand-alone.narrow",
			AbbreviatedStandalone: "stand-alone.abbreviated",
			ShortStandalone:       "stand-alone.short",
		},
		Keys: []string{"sun", "mon", "tue", "wed", "thu", "fri", "sat"},
	},
	DayPeriod: Field{
		CldrNode: "dayPeriods",
		Widths: map[Width]string{
			Wide:                  "format.wide",
			Narrow:                "format.narrow",
			Abbreviated:           "format.abbreviated",
			WideStandalone:        "stand-alone.wide",
			NarrowStandalone:      "stand-alone.narrow",
			AbbreviatedStandalone: "stand-alone.abbreviated",
		},
		Keys: []string{"am", "pm", "morning1", "noon", "afternoon1", "evening1", "midnight", "night1"},
	},
}

// Dictionary to stock our locale
// keep x/text compatibility
var locales map[string]catalog.Dictionary = make(map[string]catalog.Dictionary)

// Localize a field value
func Localize(lang string, field FieldType, width Width, value int) string {
	fi, ok := Fields[field]
	if !ok {
		return ""
	}

	key, ok := fi.Key(width, value)
	if !ok {
		return ""
	}

	loc, ok := locales[lang]
	if !ok {
		return ""
	}

	data, _ := loc.Lookup(key)
	return data
}
