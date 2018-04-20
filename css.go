package goss

import (
	"fmt"
	"sort"
	"strings"
)

//go:generate go run bin/cssprops/main.go
//go:generate go run bin/tags/main.go

type Object struct {
	parent   *Object
	selector string
	key      string
	value    interface{}
	children []*Object
}

type Styler interface {
	Set(*Object)
}

func C(s ...interface{}) *Object {
	o := &Object{}
	for _, v := range s {
		switch e := v.(type) {
		case *Object:
			e.parent = o
			o.children = append(o.children, e)
		case Styler:
			e.Set(o)
		default:
			panic("unknown type type")
		}
	}
	return o
}

type StyleFunc func(*Object)

func (s StyleFunc) Set(o *Object) {
	s(o)
}

func Prop(key string, value interface{}) Styler {
	return StyleFunc(func(o *Object) {
		o.children = append(o.children, &Object{
			parent: o,
			key:    key,
			value:  value,
		})
	})
}

func Selector(key string) Styler {
	return StyleFunc(func(o *Object) {
		o.selector = key
	})
}

func All(c ...interface{}) *Object {
	return WrapArgs(Selector("*"), c...)
}

func WrapArgs(first interface{}, opts ...interface{}) *Object {
	args := []interface{}{first}
	if len(opts) > 0 {
		args = append(args, opts...)
	}
	return C(args...)
}

type List []Styler

func (l List) Set(o *Object) {
	for _, v := range l {
		v.Set(o)
	}
}

func renderObject(selectors []string, out *Sheet, o *Object) error {
	sel := make([]string, cap(selectors))
	copy(sel, selectors)
	if o.selector != "" {
		sel = append(sel, o.selector)
	}
	if len(o.children) > 0 {
		for _, v := range o.children {
			err := renderObject(sel, out, v)
			if err != nil {
				return err
			}
		}
		return nil
	}
	if o.key != "" {
		if ob, ok := o.value.(*Object); ok {
			err := renderObject(append(sel, o.key), out, ob)
			if err != nil {
				return err
			}
		} else {
			out.set(sel, o.key, printValue(o.value))
		}
	}
	return nil
}

func printSelectors(sel []string) string {
	o := ""
	for k, v := range sel {
		v = strings.TrimSpace(v)
		if htmlTags[v] {

			// standard html tag selectors. We add space incase it is the only selector.
			if k == 0 {

				// [div] ==> div
				o += v
			} else {
				//[div,p]  ==> div p
				o += " " + v
			}
			continue
		}
		switch v[0] {
		case '>', '+', '~':
			o += " " + v
		default:
			o += v
		}
	}
	return o
}

func printValue(v interface{}) string {
	return fmt.Sprint(v)
}

type Sheet struct {
	classes map[string]string
	idGen   func(string) string
	styles  map[string][]style
}

func NewSheet() *Sheet {
	return &Sheet{
		classes: make(map[string]string),
		styles:  make(map[string][]style),
		idGen:   genId(),
	}
}

func genId() func(string) string {
	id := 0
	return func(s string) string {
		id++
		return fmt.Sprintf("%s-%d", s, id)
	}
}

func (s *Sheet) set(selectors []string, key, value string) {
	var ns []string
	for _, v := range selectors {
		if isClass(v) {
			if cn, ok := s.classes[v]; ok {
				ns = append(ns, cn)
			} else {
				g := s.idGen(v)
				s.classes[v] = g
				ns = append(ns, g)
			}
		} else {
			ns = append(ns, v)
		}
	}
	sk := printSelectors(ns)
	if v, ok := s.styles[sk]; ok {
		v = append(v, style{key: key, value: value})
		s.styles[sk] = v
	} else {
		s.styles[sk] = []style{{key: key, value: value}}
	}
}

func isClass(s string) bool {
	return s[0] == '.' && len(s) > 1
}

type style struct {
	key, value string
}

func (s *Sheet) ClassName(c string) string {
	if n, ok := s.classes[c]; ok {
		return n
	}
	return c
}

func (s *Sheet) ClassNames() []string {
	var n []string
	for _, v := range s.classes {
		n = append(n, v)
	}
	sort.Strings(n)
	return n
}

var htmlTags = map[string]bool{
	"a":          true,
	"abbr":       true,
	"address":    true,
	"area":       true,
	"article":    true,
	"aside":      true,
	"audio":      true,
	"b":          true,
	"base":       true,
	"bdi":        true,
	"bdo":        true,
	"blockquote": true,
	"body":       true,
	"br":         true,
	"button":     true,
	"canvas":     true,
	"caption":    true,
	"cite":       true,
	"code":       true,
	"col":        true,
	"colgroup":   true,
	"data":       true,
	"datalist":   true,
	"dd":         true,
	"del":        true,
	"details":    true,
	"dfn":        true,
	"dialog":     true,
	"div":        true,
	"dl":         true,
	"dt":         true,
	"em":         true,
	"embed":      true,
	"fieldset":   true,
	"figcaption": true,
	"figure":     true,
	"footer":     true,
	"form":       true,
	"h1":         true,
	"h2":         true,
	"h3":         true,
	"h4":         true,
	"h5":         true,
	"h6":         true,
	"head":       true,
	"header":     true,
	"hgroup":     true,
	"hr":         true,
	"html":       true,
	"i":          true,
	"iframe":     true,
	"img":        true,
	"input":      true,
	"ins":        true,
	"kbd":        true,
	"keygen":     true,
	"label":      true,
	"legend":     true,
	"li":         true,
	"link":       true,
	"main":       true,
	"map":        true,
	"mark":       true,
	"math":       true,
	"menu":       true,
	"menuitem":   true,
	"meta":       true,
	"meter":      true,
	"nav":        true,
	"noscript":   true,
	"object":     true,
	"ol":         true,
	"optgroup":   true,
	"option":     true,
	"output":     true,
	"p":          true,
	"param":      true,
	"picture":    true,
	"pre":        true,
	"progress":   true,
	"q":          true,
	"rb":         true,
	"rp":         true,
	"rt":         true,
	"rtc":        true,
	"ruby":       true,
	"s":          true,
	"samp":       true,
	"script":     true,
	"section":    true,
	"select":     true,
	"slot":       true,
	"small":      true,
	"source":     true,
	"span":       true,
	"strong":     true,
	"style":      true,
	"sub":        true,
	"summary":    true,
	"sup":        true,
	"svg":        true,
	"table":      true,
	"tbody":      true,
	"td":         true,
	"template":   true,
	"textarea":   true,
	"tfoot":      true,
	"th":         true,
	"thead":      true,
	"time":       true,
	"title":      true,
	"tr":         true,
	"track":      true,
	"u":          true,
	"ul":         true,
	"var":        true,
	"video":      true,
	"wbr":        true,
}
