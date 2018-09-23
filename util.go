package vected

import (
	"bytes"
	"fmt"
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
	o.journal = append(o.journal, []interface{}{
		"call", k, args,
	})
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

func (o *object) Steps() string {
	var buf bytes.Buffer
	for _, v := range o.journal {
		fmt.Fprintf(&buf, "%s : %v\n", v[0], v[1:])
	}
	return buf.String()
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
	return &object{typ: TypeUndefined}
}

func (o *object) Invoke(args ...interface{}) Value {
	return &object{typ: TypeUndefined}
}

func defaultProps() map[string]*object {
	return map[string]*object{
		"style": &object{typ: TypeObject},
	}
}
