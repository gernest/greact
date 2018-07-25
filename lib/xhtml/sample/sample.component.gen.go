package sample

import (
	"github.com/gernest/vected/lib/xhtml"
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

func (Component) Render() *xhtml.Node {
	return &xhtml.Node{Type: html.ErrorNode, DataAtom: atom.Atom(0), Data: "", Namespace: "", Attr: []html.Attribute{}, Children: []*xhtml.Node{&xhtml.Node{Type: html.ElementNode, DataAtom: atom.Html, Data: "html", Namespace: "", Attr: []html.Attribute{}, Children: []*xhtml.Node{&xhtml.Node{Type: html.ElementNode, DataAtom: atom.Head, Data: "head", Namespace: "", Attr: []html.Attribute{}, Children: []*xhtml.Node{}}, &xhtml.Node{Type: html.ElementNode, DataAtom: atom.Body, Data: "body", Namespace: "", Attr: []html.Attribute{}, Children: []*xhtml.Node{&xhtml.Node{Type: html.ElementNode, DataAtom: atom.Div, Data: "div", Namespace: "", Attr: []html.Attribute{}, Children: []*xhtml.Node{&xhtml.Node{Type: html.TextNode, DataAtom: atom.Atom(0), Data: " hello, world", Namespace: "", Attr: []html.Attribute{}, Children: []*xhtml.Node{}}}}}}}}}}
}
