package date

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/datasweet/format/lexer"
)

type lexTest struct {
	input  string
	tokens []lexer.Token
}

var lexTests = []lexTest{
	{
		"a aa aaa aaaa",
		[]lexer.Token{
			{Type: perioda, Value: "a"},
			{Type: perioda, Value: "aa"},
			{Type: perioda, Value: "aaa"},
			{Type: perioda, Value: "aaaa"},
			{Type: eof, Value: ""},
		},
	},
	{
		"b bb bbb bbbb",
		[]lexer.Token{
			{Type: periodb, Value: "b"},
			{Type: periodb, Value: "bb"},
			{Type: periodb, Value: "bbb"},
			{Type: periodb, Value: "bbbb"},
			{Type: eof, Value: ""},
		},
	},
	{
		"c cc ccc cccc",
		[]lexer.Token{
			{Type: weekdayc, Value: "c"},
			{Type: weekdayc, Value: "cc"},
			{Type: weekdayc, Value: "ccc"},
			{Type: weekdayc, Value: "cccc"},
			{Type: eof, Value: ""},
		},
	},
	{
		"d dd ddd dddd",
		[]lexer.Token{
			{Type: dayd, Value: "d"},
			{Type: dayd, Value: "dd"},
			{Type: dayd, Value: "ddd"},
			{Type: dayd, Value: "dddd"},
			{Type: eof, Value: ""},
		},
	},
	{
		"e ee eee eeee",
		[]lexer.Token{
			{Type: weekdaye, Value: "e"},
			{Type: weekdaye, Value: "ee"},
			{Type: weekdaye, Value: "eee"},
			{Type: weekdaye, Value: "eeee"},
			{Type: eof, Value: ""},
		},
	},
	{
		"g gg ggg gggg",
		[]lexer.Token{
			{Type: dayg, Value: "g"},
			{Type: dayg, Value: "gg"},
			{Type: dayg, Value: "ggg"},
			{Type: dayg, Value: "gggg"},
			{Type: eof, Value: ""},
		},
	},
	{
		"ddd 'ddd' d 'du' dd:",
		[]lexer.Token{
			{Type: dayd, Value: "ddd"},
			{Type: text, Value: "ddd"},
			{Type: dayd, Value: "d"},
			{Type: text, Value: "du"},
			{Type: dayd, Value: "dd"},
			{Type: eof, Value: ""},
		},
	},
	{
		"yyyy-MM-ddTHH:mm:ss.SSSSSSSSS'Z'+08:00",
		[]lexer.Token{
			{Type: yeary, Value: "yyyy"},
			{Type: monthM, Value: "MM"},
			{Type: dayd, Value: "dd"},
			{Type: hourH, Value: "HH"},
			{Type: minutem, Value: "mm"},
			{Type: seconds, Value: "ss"},
			{Type: secondS, Value: "SSSSSSSSS"},
			{Type: text, Value: "Z"},
			{Type: eof, Value: ""},
		},
	},
}

func TestLex(t *testing.T) {
	for _, test := range lexTests {
		tokens, err := lex(test.input)
		assert.NoError(t, err)
		assert.Equal(t, len(test.tokens), len(tokens))
		for i, tk := range test.tokens {
			assert.Equalf(t, tk.Type, tokens[i].Type, "%s wrong type at %d", test.input, i)
			assert.Equalf(t, tk.Value, tokens[i].Value, "%s wrong value at %d", test.input, i)
		}
	}
}
