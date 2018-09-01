package vected

import (
	"bytes"
	"fmt"
	"go/parser"
	"go/printer"
	"go/token"
	"strings"
)

func printAttr(a Attribute) string {
	return fmt.Sprintf("HA(%q,%q,%v)", a.Namespace, a.Key, interpret(a.Val))
}

func printNode(n *Node, level int) string {
	v := indent((fmt.Sprintf("H(%d,%q,%q",
		n.Type, n.Namespace, n.Data)), level)
	if len(n.Attr) > 0 {
		v += indent(",HAT(", level+2)
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
			case TextNode:
				x := strings.TrimSpace(ch.Data)
				if x == "" {
					continue
				}
			case CommentNode:
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
			// We remove prefix { and suffix }, pick what is left and evaluate if it is a
			// valid go expression.
			//
			// TODO handle errors when given wrong expressions.
			x := strings.TrimPrefix(e, "{")
			x = strings.TrimSuffix(x, "}")
			x = strings.TrimSpace(x)
			v, err := parser.ParseExpr(x)
			if err != nil {
				return "nil"
			}
			var buf bytes.Buffer
			fset := token.NewFileSet()
			printer.Fprint(&buf, fset, v)
			return buf.String()
		}
		return fmt.Sprintf("%q", e)
	default:
		return "nil"
	}
}

func indent(v string, n int) string {
	s := " "
	for i := 0; i < n; i++ {
		s += " "
	}
	return s + v
}
