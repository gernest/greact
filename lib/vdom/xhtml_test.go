package vdom

import (
	"bytes"
	"go/ast"
	"go/parser"
	"go/token"
	"io/ioutil"
	"strings"
	"testing"

	"github.com/kr/pretty"

	"golang.org/x/net/html"
)

func TestHTML(t *testing.T) {
	src := `<div key="{{value}}"> hello, world</div>`
	doc, err := html.Parse(strings.NewReader(src))
	if err != nil {
		t.Fatal(err)
	}
	o := &Node{}
	Clone(doc, o)

	t.Error(pretty.Sprint(o))
	// t.Error(pretty.Sprint(o))
}

const helloSrc = `package test

func run()*xhtml.Node{
	return &xhtml.Node{
		Type:      0x0,
		DataAtom:  atom.Div,
		Data:      "",
		Namespace: "",
		Children: []*xhtml.Node {
			&xhtml.Node{
				Type:      0x3,
				DataAtom:  0x14704,
				Data:      "html",
				Namespace: "",
				Attr:[]html.Attribute{
					{Namespace:"", Key:"key", Val:"{{value}}"},
					{Namespace:"", Key:"key2", Val:"{{value2}}"},
				},
				// Children: {
				// 	&xhtml.Node{
				// 		Type:      0x3,
				// 		DataAtom:  0x32904,
				// 		Data:      "head",
				// 		Namespace: "",
				// 		Children: nil,
				// 	},
				// 	&xhtml.Node{
				// 		Type:      0x3,
				// 		DataAtom:  0x2804,
				// 		Data:      "body",
				// 		Namespace: "",
				// 		Children: []xhtml.Node{
				// 			&xhtml.Node{
				// 				Type:      0x3,
				// 				DataAtom:  0x16b03,
				// 				Data:      "div",
				// 				Namespace: "",
				// 				Children: []xhtml.Node{
				// 					&xhtml.Node{
				// 						Type:      0x1,
				// 						DataAtom:  0x0,
				// 						Data:      " hello, world",
				// 						Namespace: "",
				// 					},
				// 				},
				// 			},
				// 		},
				// 	},
				// },
			},
			&xhtml.Node{
				Type:      0x3,
				DataAtom:  0x14704,
				Data:      "html",
				Namespace: "",
				Attr:[]html.Attribute{
				},
				// Children: {
				// 	&xhtml.Node{
				// 		Type:      0x3,
				// 		DataAtom:  0x32904,
				// 		Data:      "head",
				// 		Namespace: "",
				// 		Children: nil,
				// 	},
				// 	&xhtml.Node{
				// 		Type:      0x3,
				// 		DataAtom:  0x2804,
				// 		Data:      "body",
				// 		Namespace: "",
				// 		Children: []xhtml.Node{
				// 			&xhtml.Node{
				// 				Type:      0x3,
				// 				DataAtom:  0x16b03,
				// 				Data:      "div",
				// 				Namespace: "",
				// 				Children: []xhtml.Node{
				// 					&xhtml.Node{
				// 						Type:      0x1,
				// 						DataAtom:  0x0,
				// 						Data:      " hello, world",
				// 						Namespace: "",
				// 					},
				// 				},
				// 			},
				// 		},
				// 	},
				// },
			},
		},
	}
}
`

func TestAST(t *testing.T) {
	fset := token.NewFileSet() // positions are relative to fset
	f, err := parser.ParseFile(fset, "", helloSrc, 0)
	if err != nil {
		panic(err)
	}

	// Print the AST.
	ast.Print(fset, f)
	var buf bytes.Buffer
	ast.Fprint(&buf, fset, f, ast.NotNilFilter)
	ioutil.WriteFile("ast.txt", buf.Bytes(), 0600)
}

func TestMkaeNode(t *testing.T) {
	src := `<div> hello, world</div>`
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
