package main

import (
	"bytes"
	"encoding/json"
	"go/format"
	"io/ioutil"
	"log"
	"strings"
	"text/template"

	"github.com/gernest/inflect"
)

func main() {
	b, err := ioutil.ReadFile("css-properties/w3c-css-properties.json")
	if err != nil {
		log.Fatal(err)
	}
	var v []string
	err = json.Unmarshal(b, &v)
	if err != nil {
		log.Fatal(err)
	}
	s := `
package goss

var cssprops=map[string]bool{
	{{range $_,$v:= . -}}
	"{{$v}}": true,
	{{end}}
}

// useful for avoiding typing strings all the time.
const(
	{{range $_,$v:= . -}}
	{{$v|camel}} ="{{$v}}"
	{{end}}
)
`
	fu := template.FuncMap{
		"camel": func(a string) string {
			if strings.HasPrefix(a, "@") {
				return inflect.Camelize(a[1:])
			}
			return inflect.Camelize(a)
		},
	}
	tpl, err := template.New("props").Funcs(fu).Parse(s)
	if err != nil {
		log.Fatal(err)
	}
	var buf bytes.Buffer
	err = tpl.Execute(&buf, v)
	if err != nil {
		log.Fatal(err)
	}
	f, err := format.Source(buf.Bytes())
	if err != nil {
		log.Fatal(err)
	}
	ioutil.WriteFile("css_properties.go", f, 0600)
}
