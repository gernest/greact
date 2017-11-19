package main

import (
	"bytes"
	"encoding/json"
	"go/format"
	"io/ioutil"
	"log"
	"text/template"
)

func main() {
	b, err := ioutil.ReadFile("html-tags/html-tags.json")
	if err != nil {
		log.Fatal(err)
	}
	var v []string
	err = json.Unmarshal(b, &v)
	if err != nil {
		log.Fatal(err)
	}

	b, err = ioutil.ReadFile("html-tags/html-tags-void.json")
	if err != nil {
		log.Fatal(err)
	}
	var v2 []string
	err = json.Unmarshal(b, &v2)
	if err != nil {
		log.Fatal(err)
	}
	c := make(map[string][]string)
	c["tags"] = v
	c["void"] = v2
	s := `
package goss

var htmlTags=map[string]bool{
	{{range $_,$v:= .tags -}}
	"{{$v}}": true,
	{{end}}
}

var voidTags=map[string]bool{
	{{range $_,$v:= .void -}}
	"{{$v}}": true,
	{{end}}
}
`
	tpl, err := template.New("tags").Parse(s)
	if err != nil {
		log.Fatal(err)
	}
	var buf bytes.Buffer
	err = tpl.Execute(&buf, c)
	if err != nil {
		log.Fatal(err)
	}
	f, err := format.Source(buf.Bytes())
	if err != nil {
		log.Fatal(err)
	}
	ioutil.WriteFile("html_tags.go", f, 0600)
}
