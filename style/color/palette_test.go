package color

import (
	"bytes"
	"html/template"
	"io/ioutil"
	"testing"
)

func TestPalette(t *testing.T) {
	tpl, err := template.ParseFiles("index.html")
	if err != nil {
		t.Fatal(err)
	}

	var o bytes.Buffer
	err = tpl.ExecuteTemplate(&o, "index.html", New())
	if err != nil {
		t.Fatal(err)
	}
	ioutil.WriteFile("colors.html", o.Bytes(), 0600)
}
