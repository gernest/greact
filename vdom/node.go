package vdom

import (
	"fmt"
	"strings"

	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

const (
	// ContainerNode this is tha name of a node which acts as a container to a
	// slice of other nodes.
	ContainerNode = "__internal_container___"
)

// Attribute represent html attribute. Val can be any valid go expression.
type Attribute struct {
	Namespace, Key string
	Val            interface{}
}

func (a Attribute) Print() string {
	if a.Namespace != "" {
		return fmt.Sprintf("{Namespace:%q,Key:%q,Val:%v},\n",
			a.Namespace, a.Key, interpret(a.Val))
	}
	return fmt.Sprintf("{Key:%q,Val:%v},\n", a.Key, interpret(a.Val))
}

type Node struct {
	Type      html.NodeType
	DataAtom  atom.Atom
	Data      string
	Namespace string
	Attr      []Attribute
	Children  []*Node
}

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
		v += indent("Attr:[]vdom.Attribute{\n", level+2)
		for _, attr := range n.Attr {
			if attr.Key != "" {
				v += indent(attr.Print(), level+4)
			}
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

func interpret(v interface{}) string {
	switch e := v.(type) {
	case nil:
		return "nil"
	case string:
		e = strings.TrimSpace(e)
		if strings.HasPrefix(e, "{") {
			x := strings.TrimPrefix(e, "{")
			x = strings.TrimSuffix(x, "}")
			x = strings.TrimSpace(x)
			if strings.HasPrefix(x, "\"") {
				return fmt.Sprintf("%q", x)
			}
			return x
		}
		return fmt.Sprintf("%q", e)
	default:
		return fmt.Sprint(v)
	}
}

func indent(v string, n int) string {
	s := " "
	for i := 0; i < n; i++ {
		s += " "
	}
	return s + v
}
func formatNodeType(n html.NodeType) string {
	switch n {
	case html.ErrorNode:
		return "ErrorNode"
	case html.TextNode:
		return "TextNode"
	case html.DocumentNode:
		return "DocumentNode"
	case html.ElementNode:
		return "ElementNode"
	case html.CommentNode:
		return "CommentNode"
	case html.DoctypeNode:
		return "DoctypeNode"
	default:
		return "ErrorNode"
	}
}
