package vected

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"html"
	"io"
	"strings"

	"github.com/gernest/vected/attribute"
)

type object struct {
	id        int
	name      string
	namespace string
	text      bool
	parent    *object
	props     map[string]*object
	value     interface{}
	typ       Type
	nodeValue string
	cache     map[string]Value
	children  []*object
	journal   [][]interface{}
	level     int
}

func newObject() *object {
	return &object{
		id:    idPool.Get().(int),
		props: defaultProps(),
		typ:   TypeObject,
	}
}

func (o *object) Bool() bool {
	return o.value.(bool)
}

func (o *object) Float() float64 {
	return o.value.(float64)
}

func (o *object) Int() int {
	return o.value.(int)
}
func (o *object) String() string {
	return o.value.(string)
}

func (o *object) Type() Type {
	return o.typ
}

func (o *object) Set(k string, v interface{}) {
	o.journal = append(o.journal, []interface{}{
		"set", k, v,
	})
	if o.props == nil {
		o.props = make(map[string]*object)
	}
	switch e := v.(type) {
	case bool:
		o.props[k] = &object{typ: TypeBoolean, value: e}
	case string:
		o.props[k] = &object{typ: TypeString, value: e}
	case float64:
		o.props[k] = &object{typ: TypeNumber, value: e}
	case nil:
		o.props[k] = &object{typ: TypeNull, value: e}
	case Value:
		o.props[k] = &object{typ: TypeObject, value: e}
	}
}

func (o *object) Get(k string) Value {
	o.journal = append(o.journal, []interface{}{
		"get", k,
	})
	switch k {
	case "parentNode":
		if o.parent != nil {
			return o.parent
		}
		return null()
	case "nextSibling":
		if o.parent != nil {
			for k, v := range o.parent.children {
				if v.id == o.id {
					x := k + 1
					if x < len(o.parent.children) {
						return o.parent.children[x]
					}
				}
			}
		}
		return undefined()
	case "previousSibling":
		if o.parent != nil {
			for k, v := range o.parent.children {
				if v.id == o.id {
					x := k - 1
					if x >= 0 {
						return o.parent.children[x]
					}
				}
			}
		}
		return undefined()
	case "lastChild":
		if len(o.children) > 0 {
			return o.children[len(o.children)-1]
		}
		return undefined()
	case "childNodes":
		var children []Value
		for _, ch := range o.children {
			children = append(children, ch)
		}
		return &object{
			value: children,
			typ:   TypeObject,
		}
	case "length":
		if o.typ != TypeObject {
			return undefined()
		}
		switch e := o.value.(type) {
		case []Value:
			return &object{typ: TypeNumber, value: len(e)}
		}
		return undefined()
	case "splitText":
		if o.text {
			return &object{typ: TypeFunction}
		}
		return undefined()
	case "nodeValue":
		return &object{typ: TypeString, value: o.nodeValue}
	}
	if m, ok := o.props[k]; ok {
		return m
	}
	return &object{typ: TypeUndefined}
}

func (o *object) Call(k string, args ...interface{}) Value {
	v := []interface{}{"call", k}
	for _, k := range args {
		if o, ok := k.(*object); ok {
			v = append(v, o.typ)
			s := strings.TrimSpace(o.Steps())
			if s != "" {
				v = append(v, s)
			}
		} else {
			v = append(v, k)
		}
	}
	o.journal = append(o.journal, v)
	switch k {
	case "hasOwnProperty":
		if len(args) > 0 {
			a := args[0]
			if av, ok := a.(string); ok {
				_, ok = o.props[av]
				return &object{typ: TypeBoolean, value: ok}
			}
		}
		return &object{typ: TypeBoolean, value: false}
	case "createElement":
		// element name must be provided.
		name := args[0].(string)
		b := newObject()
		b.name = name
		return b
	case "createElementNS":
		ns := args[0].(string)
		name := args[1].(string)
		b := newObject()
		b.namespace = ns
		b.name = name
		return b
	case "createTextNode":
		text := args[0].(string)
		b := newObject()
		b.text = true
		b.nodeValue = text
		return b
	case "replaceChild":
		if len(args) == 2 {
			a, ok := args[0].(*object)
			if !ok {
				return undefined()
			}
			b, ok := args[0].(*object)
			if !ok {
				return undefined()
			}
			return o.replaceChild(a, b)
		}
	case "removeChild":
		if len(args) == 1 {
			a, ok := args[0].(*object)
			if !ok {
				return undefined()
			}
			if len(o.children) > 0 {
				var sv []*object
				for _, v := range o.children {
					if v.id != a.id {
						sv = append(sv, v)
					}
				}
				o.children = sv
			}
		}
	case "appendChild":
		if len(args) == 1 {
			a, ok := args[0].(*object)
			if !ok {
				return undefined()
			}
			a.parent = o
			a.level = o.level + 2
			o.children = append(o.children, a)
			return undefined()
		}
	case "insertBefore":
		if len(args) == 2 {
			a, ok := args[0].(*object)
			if !ok {
				return undefined()
			}
			b, ok := args[0].(*object)
			if !ok {
				return undefined()
			}
			return o.insertBefore(a, b)
		}
	case "isEqualNode":
		if len(args) == 1 {
			a, ok := args[0].(*object)
			if !ok {
				return undefined()
			}
			return &object{typ: TypeBoolean, value: o.id == a.id}
		}
		return &object{typ: TypeBoolean, value: false}
	}
	return undefined()
}

func (o object) Steps() string {
	var buf bytes.Buffer
	for _, v := range o.journal {
		fmt.Fprintf(&buf, "%s%v\n", indent(o.level), v)
		if len(o.children) > 0 {
			for _, ch := range o.children {
				buf.WriteString(ch.Steps())
			}
		}
	}
	return buf.String()
}

func validAttribute(v string) bool {
	_, ok := attribute.Map[v]
	return ok
}

func indent(n int) (out string) {
	for i := 0; i < n; i++ {
		out += " "
	}
	return
}

func (o *object) replaceChild(a, b *object) *object {
	if len(o.children) > 0 {
		var rst []*object
		for _, v := range o.children {
			if v.id == a.id {
				rst = append(rst, b)
			} else {
				rst = append(rst, v)
			}
		}
		o.children = rst
	}
	return undefined()
}
func (o *object) insertBefore(a, b *object) *object {
	if len(o.children) > 0 {
		var rst []*object
		for _, v := range o.children {
			if v.id == a.id {
				rst = append(rst, b, a)
			} else {
				rst = append(rst, v)
			}
		}
		o.children = rst
	}
	return undefined()
}

func undefined() *object {
	return &object{typ: TypeUndefined}
}
func null() *object {
	return &object{typ: TypeNull}
}

func (o *object) Index(n int) Value {
	if v, ok := o.value.([]Value); ok {
		if n < len(v) {
			return v[n]
		}
	}
	return &object{typ: TypeNull}
}

func (o *object) Invoke(args ...interface{}) Value {
	return &object{typ: TypeUndefined}
}

func defaultProps() map[string]*object {
	return map[string]*object{
		"style": &object{typ: TypeObject},
	}
}

type writer interface {
	io.Writer
	io.ByteWriter
	WriteString(string) (int, error)
}

// Render renders the parse tree n to the given writer.
//
// Rendering is done on a 'best effort' basis: calling Parse on the output of
// Render will always result in something similar to the original tree, but it
// is not necessarily an exact clone unless the original tree was 'well-formed'.
// 'Well-formed' is not easily specified; the HTML5 specification is
// complicated.
//
// Calling Parse on arbitrary input typically results in a 'well-formed' parse
// tree. However, it is possible for Parse to yield a 'badly-formed' parse tree.
// For example, in a 'well-formed' parse tree, no <a> element is a child of
// another <a> element: parsing "<a><a>" results in two sibling elements.
// Similarly, in a 'well-formed' parse tree, no <a> element is a child of a
// <table> element: parsing "<p><table><a>" results in a <p> with two sibling
// children; the <a> is reparented to the <table>'s parent. However, calling
// Parse on "<a><table><a>" does not return an error, but the result has an <a>
// element with an <a> child, and is therefore not 'well-formed'.
//
// Programmatically constructed trees are typically also 'well-formed', but it
// is possible to construct a tree that looks innocuous but, when rendered and
// re-parsed, results in a different tree. A simple example is that a solitary
// text node would become a tree containing <html>, <head> and <body> elements.
// Another example is that the programmatic equivalent of "a<head>b</head>c"
// becomes "<html><head><head/><body>abc</body></html>".
func renderObject(w io.Writer, n *object) error {
	if x, ok := w.(writer); ok {
		return renderNode(x, n)
	}
	buf := bufio.NewWriter(w)
	if err := renderNode(buf, n); err != nil {
		return err
	}
	return buf.Flush()
}

// plaintextAbort is returned from render1 when a <plaintext> element
// has been rendered. No more end tags should be rendered after that.
var plaintextAbort = errors.New("html: internal error (plaintext abort)")

func renderNode(w writer, n *object) error {
	err := render1(w, n)
	if err == plaintextAbort {
		err = nil
	}
	return err
}

func render1(w writer, n *object) error {
	if n.text {
		e := html.EscapeString(n.nodeValue)
		_, err := w.WriteString(e)
		return err
	}

	// Render the <xxx> opening tag.
	if err := w.WriteByte('<'); err != nil {
		return err
	}
	if _, err := w.WriteString(n.name); err != nil {
		return err
	}
	for key, a := range n.props {
		if !validAttribute(key) {
			continue
		}
		switch key {
		case "style":
			if len(a.props) > 0 {
				if err := w.WriteByte(' '); err != nil {
					return err
				}
				var value string
				for propKey, propValue := range a.props {
					switch propValue.typ {
					case TypeBoolean, TypeNumber, TypeString:
						value = fmt.Sprint(a.value)
						s := fmt.Sprintf("%s:%v", propKey, propValue.value)
						if value != "" {
							value += ";" + s
						} else {
							value = s
						}
					default:
						continue
					}
				}
				if _, err := w.WriteString(key); err != nil {
					return err
				}
				if _, err := w.WriteString(`="`); err != nil {
					return err
				}
				_, err := w.WriteString(html.EscapeString(value))
				if err != nil {
					return err
				}
				if err := w.WriteByte('"'); err != nil {
					return err
				}
			}
		default:
			if err := w.WriteByte(' '); err != nil {
				return err
			}
			// if a.Namespace != "" {
			// 	if _, err := w.WriteString(a.Namespace); err != nil {
			// 		return err
			// 	}
			// 	if err := w.WriteByte(':'); err != nil {
			// 		return err
			// 	}
			// }
			var value string
			switch a.typ {
			case TypeBoolean, TypeNumber, TypeString:
				value = fmt.Sprint(a.value)
			default:
				continue
			}
			if _, err := w.WriteString(key); err != nil {
				return err
			}
			if _, err := w.WriteString(`="`); err != nil {
				return err
			}
			_, err := w.WriteString(html.EscapeString(value))
			if err != nil {
				return err
			}
			if err := w.WriteByte('"'); err != nil {
				return err
			}
		}
	}
	if voidElements[n.name] {
		if len(n.children) > 0 {
			return fmt.Errorf("html: void element <%s> has child nodes", n.name)
		}
		_, err := w.WriteString("/>")
		return err
	}
	if err := w.WriteByte('>'); err != nil {
		return err
	}

	if len(n.children) > 0 {
		// Add initial newline where there is danger of a newline beging ignored.
		if c := n.children[0]; c != nil && c.text && strings.HasPrefix(c.nodeValue, "\n") {
			switch n.name {
			case "pre", "listing", "textarea":
				if err := w.WriteByte('\n'); err != nil {
					return err
				}
			}
		}
	}

	// Render any child nodes.
	switch n.name {
	case "iframe", "noembed", "noframes", "noscript", "plaintext", "script", "style", "xmp":
		for _, c := range n.children {
			if c.text {
				if _, err := w.WriteString(c.nodeValue); err != nil {
					return err
				}
			} else {
				if err := render1(w, c); err != nil {
					return err
				}
			}
		}
		if n.name == "plaintext" {
			// Don't render anything else. <plaintext> must be the
			// last element in the file, with no closing tag.
			return plaintextAbort
		}
	default:
		for _, c := range n.children {
			if err := render1(w, c); err != nil {
				return err
			}
		}
	}

	// Render the </xxx> closing tag.
	if _, err := w.WriteString("</"); err != nil {
		return err
	}
	if _, err := w.WriteString(n.name); err != nil {
		return err
	}
	return w.WriteByte('>')
}

// writeQuoted writes s to w surrounded by quotes. Normally it will use double
// quotes, but if s contains a double quote, it will use single quotes.
// It is used for writing the identifiers in a doctype declaration.
// In valid HTML, they can't contain both types of quotes.
func writeQuoted(w writer, s string) error {
	var q byte = '"'
	if strings.Contains(s, `"`) {
		q = '\''
	}
	if err := w.WriteByte(q); err != nil {
		return err
	}
	if _, err := w.WriteString(s); err != nil {
		return err
	}
	if err := w.WriteByte(q); err != nil {
		return err
	}
	return nil
}

// Section 12.1.2, "Elements", gives this list of void elements. Void elements
// are those that can't have any contents.
var voidElements = map[string]bool{
	"area":    true,
	"base":    true,
	"br":      true,
	"col":     true,
	"command": true,
	"embed":   true,
	"hr":      true,
	"img":     true,
	"input":   true,
	"keygen":  true,
	"link":    true,
	"meta":    true,
	"param":   true,
	"source":  true,
	"track":   true,
	"wbr":     true,
}
