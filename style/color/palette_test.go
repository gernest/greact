package color

import (
	"bytes"
	"html/template"
	"io/ioutil"
	"testing"
)

func TestPalette(t *testing.T) {
	tpl, err := template.New("t").Funcs(
		template.FuncMap{
			"hsv": func(v *Color) template.HTML {
				return template.HTML(PrintColor(v, "hsv"))
			},
		},
	).ParseFiles("index.html")
	if err != nil {
		t.Fatal(err)
	}

	var o bytes.Buffer
	err = tpl.ExecuteTemplate(&o, "index.html", NewPaletter())
	if err != nil {
		t.Fatal(err)
	}
	ioutil.WriteFile("colors.html", o.Bytes(), 0600)
}
