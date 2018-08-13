package parser

import (
	"fmt"
	"strings"

	"github.com/gernest/vected/vdom"
)

func printAttr(a vdom.Attribute) string {
	if a.Namespace != "" {
		return fmt.Sprintf("{Namespace:%q,Key:%q,Val:%v},\n",
			a.Namespace, a.Key, interpret(a.Val))
	}
	return fmt.Sprintf("{Key:%q,Val:%v},\n", a.Key, interpret(a.Val))
}

func printNode(n *vdom.Node, level int, child bool) string {
	v := indent("&vdom.Node{\n", level)
	if n.Type != 0 {
		v += indent("Type: vdom."+n.Type.String()+",\n", level+2)
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
				v += indent(printAttr(attr), level+4)
			}
		}
		v += indent("},\n", level+2)
	}
	if len(n.Children) > 0 {
		v += indent("Children:[]*vdom.Node{\n", level+2)
		for _, ch := range n.Children {
			v += printNode(ch, level+4, true)
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
