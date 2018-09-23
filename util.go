package vected

import (
	"bytes"
	"fmt"
)

type Object struct {
	id        int
	name      string
	namespace string
	text      bool
	parent    *Object
	props     map[string]*Object
	value     interface{}
	typ       Type
	nodeValue string
	cache     map[string]Value
	children  []*Object
	journal   [][]interface{}
}

func NewObject() *Object {
	return &Object{
		id:    idPool.Get().(int),
		props: defaultProps(),
		typ:   TypeObject,
	}
}

func (o *Object) Bool() bool {
	return o.value.(bool)
}

func (o *Object) Float() float64 {
	return o.value.(float64)
}

func (o *Object) Int() int {
	return o.value.(int)
}
func (o *Object) String() string {
	return o.value.(string)
}

func (o *Object) Type() Type {
	return o.typ
}

func (o *Object) Set(k string, v interface{}) {
	o.journal = append(o.journal, []interface{}{
		"set", k, v,
	})
	if o.props == nil {
		o.props = make(map[string]*Object)
	}
	switch e := v.(type) {
	case bool:
		o.props[k] = &Object{typ: TypeBoolean, value: e}
	case string:
		o.props[k] = &Object{typ: TypeString, value: e}
	case float64:
		o.props[k] = &Object{typ: TypeNumber, value: e}
	case nil:
		o.props[k] = &Object{typ: TypeNull, value: e}
	case Value:
		o.props[k] = &Object{typ: TypeObject, value: e}
	}
}

func (o *Object) Get(k string) Value {
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
		return &Object{
			value: children,
			typ:   TypeObject,
		}
	case "length":
		if o.typ != TypeObject {
			return undefined()
		}
		switch e := o.value.(type) {
		case []Value:
			return &Object{typ: TypeNumber, value: len(e)}
		}
		return undefined()
	case "splitText":
		if o.text {
			return &Object{typ: TypeFunction}
		}
		return undefined()
	case "nodeValue":
		return &Object{typ: TypeString, value: o.nodeValue}
	}
	if m, ok := o.props[k]; ok {
		return m
	}
	return &Object{typ: TypeUndefined}
}

func (o *Object) Call(k string, args ...interface{}) Value {
	o.journal = append(o.journal, []interface{}{
		"call", k, args,
	})
	switch k {
	case "hasOwnProperty":
		if len(args) > 0 {
			a := args[0]
			if av, ok := a.(string); ok {
				_, ok = o.props[av]
				return &Object{typ: TypeBoolean, value: ok}
			}
		}
		return &Object{typ: TypeBoolean, value: false}
	case "createElement":
		// element name must be provided.
		name := args[0].(string)
		b := NewObject()
		b.name = name
		return b
	case "createElementNS":
		ns := args[0].(string)
		name := args[1].(string)
		b := NewObject()
		b.namespace = ns
		b.name = name
		return b
	case "createTextNode":
		text := args[0].(string)
		b := NewObject()
		b.text = true
		b.nodeValue = text
		return b
	case "replaceChild":
		if len(args) == 2 {
			a, ok := args[0].(*Object)
			if !ok {
				return undefined()
			}
			b, ok := args[0].(*Object)
			if !ok {
				return undefined()
			}
			return o.replaceChild(a, b)
		}
	case "removeChild":
		if len(args) == 1 {
			a, ok := args[0].(*Object)
			if !ok {
				return undefined()
			}
			if len(o.children) > 0 {
				var sv []*Object
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
			a, ok := args[0].(*Object)
			if !ok {
				return undefined()
			}
			a.parent = o
			o.children = append(o.children, a)
			return undefined()
		}
	case "insertBefore":
		if len(args) == 2 {
			a, ok := args[0].(*Object)
			if !ok {
				return undefined()
			}
			b, ok := args[0].(*Object)
			if !ok {
				return undefined()
			}
			return o.insertBefore(a, b)
		}
	case "isEqualNode":
		if len(args) == 1 {
			a, ok := args[0].(*Object)
			if !ok {
				return undefined()
			}
			return &Object{typ: TypeBoolean, value: o.id == a.id}
		}
		return &Object{typ: TypeBoolean, value: false}
	}
	return undefined()
}

func (o *Object) Steps() string {
	var buf bytes.Buffer
	for _, v := range o.journal {
		fmt.Fprintf(&buf, "%s : %v\n", v[0], v[1:])
	}
	return buf.String()
}
func (o *Object) replaceChild(a, b *Object) *Object {
	if len(o.children) > 0 {
		var rst []*Object
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
func (o *Object) insertBefore(a, b *Object) *Object {
	if len(o.children) > 0 {
		var rst []*Object
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

func undefined() *Object {
	return &Object{typ: TypeUndefined}
}
func null() *Object {
	return &Object{typ: TypeNull}
}

func (o *Object) Index(n int) Value {
	if v, ok := o.value.([]Value); ok {
		if n < len(v) {
			return v[n]
		}
	}
	return &Object{typ: TypeUndefined}
}

func (o *Object) Invoke(args ...interface{}) Value {
	return &Object{typ: TypeUndefined}
}

func defaultProps() map[string]*Object {
	return map[string]*Object{
		"style": &Object{typ: TypeObject},
	}
}
