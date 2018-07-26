package vdom

import (
	"bytes"
	"io/ioutil"
	"strings"
	"testing"

	"golang.org/x/net/html"
)

func TestGenerateRenderMethod(t *testing.T) {
	src := `<div name="value"> hello, world</div>`
	doc, err := html.Parse(strings.NewReader(src))
	if err != nil {
		t.Fatal(err)
	}
	o := &Node{
		DataAtom: doc.DataAtom,
		Data:     doc.Data,
	}
	Clone(doc, o)
	v, err := GenerateRenderMethod(o, &Context{
		Package:    "sample",
		StructName: "Component",
	})
	if err != nil {
		t.Fatal(err)
	}
	ioutil.WriteFile("sample/sample.component.gen.go", v, 0600)
}

func TestRender(t *testing.T) {
	src := `<div name="value"> hello, world</div>`
	doc, err := html.Parse(strings.NewReader(src))
	if err != nil {
		t.Fatal(err)
	}
	o := &Node{
		Type:     doc.Type,
		DataAtom: doc.DataAtom,
		Data:     doc.Data,
	}
	Clone(doc, o)
	var buf bytes.Buffer
	err = Render(&buf, o)
	if err != nil {
		t.Fatal(err)
	}
	t.Error(buf.String())
}
