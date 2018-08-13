package parser

import (
	"fmt"
	"strings"

	"github.com/gernest/vected/vdom"
)

func printAttr(a vdom.Attribute) string {
	return fmt.Sprintf("ha(%q,%q,%v)", a.Namespace, a.Key, interpret(a.Val))
}

func printNode(n *vdom.Node, level int) string {
	v := indent((fmt.Sprintf("h(%d,%q,%q",
		n.Type, n.Namespace, n.Data)), level)
	if len(n.Attr) > 0 {
		v += indent(",hat(", level+2)
		for _, attr := range n.Attr {
			if attr.Key == "" {
				continue
			}
			v += "\n"
			v += indent(printAttr(attr), level+4)
			v += ","
		}
		v += indent(")", level+2)
	} else {
		v += ",nil"
	}
	if len(n.Children) > 0 {
		v += ","
		for _, ch := range n.Children {
			switch ch.Type {
			case vdom.TextNode:
				x := strings.TrimSpace(ch.Data)
				if x == "" {
					continue
				}
			case vdom.CommentNode:
				continue
			}
			v += "\n"
			v += printNode(ch, level+4)
			v += ","
		}

	}
	v += indent(")", level)
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
			parts := strings.Split(x, ".")
			if len(parts) > 1 {
				for k, v := range parts {
					if k == 0 {
						continue
					}
					parts[k] = fmt.Sprintf("[%q]", v)
				}
				return strings.Join(parts, "")
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
