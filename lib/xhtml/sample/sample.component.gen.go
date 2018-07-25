package sample

import (
	"github.com/gernest/vected/lib/xhtml"
	"golang.org/x/net/html/atom"
)

func (Component) Render() *xhtml.Node {
	return &xhtml.Node{Type: 0, DataAtom: atom.Atom(0), Data: "", Namespace: "", Attr: []html.Attribute{}, Children: []*xhtml.Node{&xhtml.Node{Type: 3, DataAtom: atom.Html, Data: "html", Namespace: "", Attr: []html.Attribute{}, Children: []*xhtml.Node{&xhtml.Node{Type: 3, DataAtom: atom.Head, Data: "head", Namespace: "", Attr: []html.Attribute{}, Children: []*xhtml.Node{}}, &xhtml.Node{Type: 3, DataAtom: atom.Body, Data: "body", Namespace: "", Attr: []html.Attribute{}, Children: []*xhtml.Node{&xhtml.Node{Type: 3, DataAtom: atom.Div, Data: "div", Namespace: "", Attr: []html.Attribute{}, Children: []*xhtml.Node{&xhtml.Node{Type: 1, DataAtom: atom.Atom(0), Data: " hello, world", Namespace: "", Attr: []html.Attribute{}, Children: []*xhtml.Node{}}}}}}}}}}
}
