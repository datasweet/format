//+build ignore

package main

//This program generates locales.go. It can be invoked by running
//go:generate

import (
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/datasweet/format/date/locales"

	"github.com/datasweet/format/third_party/gen"
	"github.com/datasweet/jsonmap"
	"github.com/pkg/errors"
	"golang.org/x/text/language"
)

type field struct {
	name   string
	widths map[locales.Width]string
	keys   []string
}

type keyIndex struct {
	pos   int
	key   string
	xpath string
}

type translation struct {
	lang      string
	langData  []string
	langIndex []uint32
}

func main() {
	files, err := filepath.Glob("locales/*.json")
	die(errors.Wrap(err, "can't list locales"))

	var indexes []*keyIndex
	var translations []*translation

	// analyze fields
	for _, field := range locales.Fields {
		for w, p := range field.Widths {
			for i, k := range field.Keys {
				if key, ok := field.Key(w, i); ok {
					indexes = append(indexes, &keyIndex{
						pos:   len(indexes),
						key:   key,
						xpath: strings.Join([]string{field.CldrNode, p, k}, "."),
					})
				}
			}
		}
	}

	for _, filename := range files {
		tag := language.MustParse(filename[8 : len(filename)-5])

		data, err := ioutil.ReadFile(filename)
		die(errors.Wrapf(err, "read file %s", filename))

		j := jsonmap.FromBytes(data)
		lang := j.Get("main").Get(tag.String())
		if jsonmap.IsNil(lang) {
			log.Fatalf("wrong lang : expected '%s' got '%s'", tag)
		}

		dates := lang.Get("dates.calendars.gregorian")
		if jsonmap.IsNil(dates) {
			log.Fatal("not a gregorian calendar")
		}

		translation := new(translation)
		translation.lang = tag.String()
		translation.langIndex = []uint32{0}
		pos := 0

		for _, keyIndex := range indexes {
			s := dates.Get(keyIndex.xpath).AsString()
			// fmt.Println(translation.lang, keyIndex.key, keyIndex.pos, keyIndex.xpath, s)
			pos += len(s)
			translation.langData = append(translation.langData, s)
			translation.langIndex = append(translation.langIndex, uint32(pos))
		}

		translations = append(translations, translation)
	}

	cw := gen.NewCodeWriter()

	langs := make([]string, 0, len(translations))
	for _, tr := range translations {
		langs = append(langs, tr.lang)
	}

	// Generate code file
	x := &struct {
		Languages []string
	}{
		Languages: langs,
	}
	err = lookup.Execute(cw, x)
	die(errors.Wrap(err, "tmpl"))

	fmt.Fprint(cw, "var messageKeyToIndex = map[string]int{\n")
	for _, keyIndex := range indexes {
		fmt.Fprintf(cw, "%q: %d,\n", keyIndex.key, keyIndex.pos)
	}
	fmt.Fprint(cw, "}\n\n")

	for _, tr := range translations {
		cw.WriteVar(fmt.Sprintf("%sIndex", tr.lang), tr.langIndex)
		cw.WriteConst(fmt.Sprintf("%sData", tr.lang), strings.Join(tr.langData, ""))
	}
	cw.WriteGoFile("locales/locales_gen.go", "locales")

	// Generate test file
	cw = gen.NewCodeWriter()
	err = test.Execute(cw, x)
	die(errors.Wrap(err, "tmpl"))
	cw.WriteGoFile("locales/locales_gen_test.go", "locales_test")

	fmt.Println("Done, check file locales/locales_gen.go")
}

func die(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

var lookup = template.Must(template.New("gen").Parse(`

import "golang.org/x/text/message/catalog"

type dictionary struct {
	index []uint32
	data  string
}
func (d *dictionary) Lookup(key string) (data string, ok bool) {
	p, ok := messageKeyToIndex[key]
	if !ok {
		return "", false
	}
	start, end := d.index[p], d.index[p+1]
	if start == end {
		return "", false
	}
	return d.data[start:end], true
}
func init() {
	locales = map[string]catalog.Dictionary{
		{{range .Languages}}"{{.}}": &dictionary{index: {{.}}Index, data: {{.}}Data },
		{{end}}
	}
}
`))

var test = template.Must(template.New("test").Parse(`
import (
	"testing"

	"github.com/datasweet/format/date/locales"
	"github.com/stretchr/testify/assert"
)
type dictionary struct {
	index []uint32
	data  string
}
{{range .Languages}}
func TestLocalize{{.}}(t *testing.T) {
	for ft, fi := range locales.Fields {
		for w := range fi.Widths {
			for i := range fi.Keys {
				key, ok := fi.Key(w, i)
				assert.True(t, ok, key)
				assert.NotEmpty(t, key, key)
				s := locales.Localize("{{.}}", ft, w, i)
				assert.NotEmpty(t, s, key)
			}
		}
	}
}
{{end}}
`))
