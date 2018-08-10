package vdom

import (
	"bytes"
	"io/ioutil"
	"strings"
	"testing"

	"github.com/kr/pretty"
	"golang.org/x/net/html"
)

func TestGenerateRenderMethod(t *testing.T) {
	src := `<div name="value"> hello, world</div>`
	o, err := Parse(strings.NewReader(src))
	if err != nil {
		t.Fatal(err)
	}
	v, err := GenerateRenderMethod(o, &Context{
		Package:    "sample",
		StructName: "Component",
	})
	if err != nil {
		t.Fatal(err)
	}
	o = Clear(o)
	t.Error(pretty.Sprint(o))
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

func TestClear(t *testing.T) {
	t.Run("should return  element", func(ts *testing.T) {
		e := `<div></div>`
		n, err := ParseString(e)
		if err != nil {
			ts.Fatal(err)
		}
		if n.Data != "div" {
			t.Errorf("expected div got %s", n.Data)
		}
	})
	t.Run("should return  container element", func(ts *testing.T) {
		e := `
		<div>
		</div>
		<div>
		</div>
		`
		n, err := ParseString(e)
		if err != nil {
			ts.Fatal(err)
		}
		if n.Data != ContainerNode {
			t.Errorf("expected %s got %s", ContainerNode, n.Data)
		}
	})
}
