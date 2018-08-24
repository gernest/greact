package testutil

import (
	"sync"

	"github.com/gernest/vected/vdom/value"
)

var idx int64

var idPool = &sync.Pool{
	New: func() interface{} {
		idx++
		return idx
	},
}

type CallbackHandle func([]value.Value)

func (CallbackHandle) Release() {}

type Object struct {
	id        int64
	name      string
	namespace string
	parent    *Object
	props     map[string]*Object
	value     interface{}
	typ       value.Type
	cache     map[string]value.Value
	children  []*Object
}

func NewObject() *Object {
	return &Object{id: idPool.Get().(int64), props: defaultProps()}
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

func (o *Object) Type() value.Type {
	return o.typ
}

func (o *Object) Set(k string, v interface{}) {
	if o.props == nil {
		o.props = make(map[string]*Object)
	}
	switch e := v.(type) {
	case bool:
		o.props[k] = &Object{typ: value.TypeBoolean, value: e}
	case string:
		o.props[k] = &Object{typ: value.TypeString, value: e}
	case float64:
		o.props[k] = &Object{typ: value.TypeNumber, value: e}
	case nil:
		o.props[k] = &Object{typ: value.TypeNull, value: e}
	case value.Value:
		o.props[k] = &Object{typ: value.TypeObject, value: e}
	}
}

func (o *Object) Get(k string) value.Value {
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
	}
	if m, ok := o.props[k]; ok {
		return m
	}
	return &Object{typ: value.TypeUndefined}
}

func (o *Object) Call(k string, args ...interface{}) value.Value {
	switch k {
	case "hasOwnProperty":
		if len(args) > 0 {
			a := args[0]
			if av, ok := a.(string); ok {
				_, ok = o.props[av]
				return &Object{typ: value.TypeBoolean, value: ok}
			}
		}
		return &Object{typ: value.TypeBoolean, value: false}
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
	}
	return undefined()
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
	return &Object{typ: value.TypeUndefined}
}
func null() *Object {
	return &Object{typ: value.TypeNull}
}

func (o *Object) Index(n int) value.Value {
	if v, ok := o.value.([]interface{}); ok {
		return &Object{value: v[n]}
	}
	return &Object{typ: value.TypeUndefined}
}

func (o *Object) Invoke(args ...interface{}) value.Value {
	return &Object{typ: value.TypeUndefined}
}

func defaultProps() map[string]*Object {
	return map[string]*Object{
		"style": &Object{typ: value.TypeObject},
	}
}
