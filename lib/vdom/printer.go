package vdom

import (
	"fmt"
	"strings"
)

func (n *Node) Print(level int, child bool) string {
	v := indent("&vdom.Node{\n", level)
	if n.Type != 0 {
		v += indent("Type: html."+formatNodeType(n.Type)+",\n", level+2)
	}
	if n.DataAtom != 0 {
		v += indent("DataAtom: atom."+strings.Title(n.DataAtom.String())+",\n", level+2)
	}
	if n.Data != "" {
		v += indent("Data:"+fmt.Sprintf("%q", n.Data)+",\n", level+2)
	}
	if n.Namespace != "" {
		v += indent("Namespace:"+fmt.Sprintf("%q", n.Namespace)+",\n", level+2)
	}
	if len(n.Attr) > 0 {
		v += indent("Attr:[]html.Attribute{\n", level+2)
		for _, attr := range n.Attr {
			v += indent(fmt.Sprintf("{Namespace:%q,Key:%q,Val:%q},\n", attr.Namespace, attr.Key, attr.Val), level+4)
		}
		v += indent("},\n", level+2)
	}
	if len(n.Children) > 0 {
		v += indent("Children:[]*vdom.Node{\n", level+2)
		for _, ch := range n.Children {
			v += ch.Print(level+4, true)
		}
		v += indent("},\n", level+2)
	}
	if child {
		v += indent("},\n", level)
	} else {
		v += indent("}\n", level)
	}
	return v
}

func indent(v string, n int) string {
	s := " "
	for i := 0; i < n; i++ {
		s += " "
	}
	return s + v
}
