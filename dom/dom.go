// +build !js

package dom

import (
	"fmt"
	"strconv"
	"unsafe"
)

// Type defines a javascript object type This is an abstraction over the
// syscall/js, allowing the library development to happen without GOARCH=wasm to
// speed up things since tooling for wasm is lacking.
type Type int

// supported javascript object types.
const (
	TypeUndefined Type = iota
	TypeNull
	TypeBoolean
	TypeNumber
	TypeString
	TypeSymbol
	TypeObject
	TypeFunction
)

func (t Type) String() string {
	switch t {
	case TypeUndefined:
		return "undefined"
	case TypeNull:
		return "null"
	case TypeBoolean:
		return "bool"
	case TypeNumber:
		return "Number"
	case TypeString:
		return "String"
	case TypeSymbol:
		return "Symbol"
	case TypeObject:
		return "Object"
	case TypeFunction:
		return "Function"
	default:
		return "unknown"
	}
}

type ValueError struct {
	Method string
	Type   Type
}

func (e *ValueError) Error() string {
	return "syscall/js: call of " + e.Method + " on " + e.Type.String()
}

type Value struct {
	parent *Value
	v      interface{}
	typ    Type
}

func (v Value) JSValue() Value {
	return v
}

func (v Value) IsNull() bool {
	return v.typ == TypeUndefined
}

func (v Value) Set(p string, x interface{}) {
	if v.typ != TypeObject {
		panic(&ValueError{"Value.Set", v.typ})
	}
	v.v.(map[string]Value)[p] = ValueOf(x)
}

func (v Value) Get(p string) Value {
	if v.typ != TypeObject {
		panic(&ValueError{"Value.Get", v.typ})
	}
	switch p {
	case "parentNode":
		if v.parent != nil {
			return *v.parent
		}
		return Value{}
	case "nextSibling":
		//TODO
	case "previousSibling":
		//TODO
	case "lastChild":
	case "childNodes":
		//TODO
	case "length":
	case "splitText":
		//TODO
	case "nodeValue":
		//TODO
	}
	return v.v.(map[string]Value)[p]
}

func (v Value) Call(m string, args ...interface{}) Value {
	if v.typ != TypeObject {
		panic(&ValueError{"Value.Get", v.typ})
	}
	switch m {
	case "hasOwnProperty":
		// TODO
	case "createElement":
	case "createElementNS":
	case "replaceChild":
	case "removeChild":
	case "insertBefore":
	case "isEqualNode":
	}
	prop := v.Get(m)
	if prop.typ != TypeFunction {
		panic("syscall/js: Value.Call: property " + m + " is not a function, got " + prop.typ.String())
	}
	var a []Value
	for _, arg := range args {
		a = append(a, ValueOf(arg))
	}
	return ValueOf(v.v.(Func).fn(v, a))
}

func IsNumber(v Value) bool {
	return v.typ == TypeNumber
}

func (v Value) float(method string) float64 {
	if v.typ != TypeNumber {
		panic(&ValueError{method, v.typ})
	}
	return v.v.(float64)
}

func (v Value) Int() int {
	return int(v.float("Value.Int"))
}
func ValueOf(x interface{}) Value {
	switch e := x.(type) {
	case Value:
		return e
	case Func:
		return Value{v: e, typ: TypeFunction}
	case nil:
		return Value{typ: TypeNull}
	case bool:
		return Value{v: e, typ: TypeBoolean}
	case int, int8, int16, int32, int64,
		uint8, uint16, uint32, uint64, uintptr,
		unsafe.Pointer, float32, float64:
		return Value{v: e, typ: TypeNumber}
	case string:
		return Value{v: e, typ: TypeString}
	case []interface{}:
		var a []Value
		for _, v := range e {
			a = append(a, ValueOf(v))
		}
		return Value{v: a, typ: TypeObject}
	case map[string]interface{}:
		a := make(map[string]Value)
		for k, v := range e {
			a[k] = ValueOf(v)
		}
		return Value{v: a, typ: TypeObject}
	default:
		panic("ValueOf: invalid value")
	}
}

func Null() Value {
	return Value{typ: TypeNull}
}

type Func struct {
	fn func(this Value, args []Value) interface{}
}

func FuncOf(fn func(this Value, args []Value) interface{}) Func {
	return Func{fn: fn}
}

func (c Func) Release() {}

// Keys is like Object.keys, this returns nil if v is not an object.
func Keys(v Value) (keys []string) {
	if v.typ != TypeObject {
		panic(&ValueError{"Value.Object", v.typ})
	}
	var o []string
	for v := range v.v.(map[string]Value) {
		o = append(o, v)
	}
	return o
}
func (v Value) String() string {
	switch v.typ {
	case TypeString:
		return v.v.(string)
	case TypeUndefined:
		return "<undefined>"
	case TypeNull:
		return "<null>"
	case TypeBoolean:
		return "<boolean: " + strconv.FormatBool(v.v.(bool)) + ">"
	case TypeNumber:
		return "<number: " + fmt.Sprint(v.v) + ">"
	case TypeSymbol:
		return "<symbol>"
	case TypeObject:
		return "<object>"
	case TypeFunction:
		return "<function>"
	default:
		panic("bad type")
	}
}

func (v Value) Index(i int) Value {
	if v.typ != TypeObject {
		panic(&ValueError{"Value.Index", v.typ})
	}
	return v.v.([]Value)[i]
}

func (v Value) IsUndefined() bool {
	return v.typ == TypeUndefined
}

// IsEqual returns true if the ywo elements are equal
func IsEqual(a, b Value) bool {
	if !Valid(a) || !Valid(b) {
		return false
	}
	return a.Call("isEqualNode", b).Bool()
}

// Valid returns true if value is not null or undefined.
func Valid(v Value) bool {
	if v.IsNull() {
		return false
	}
	return v.typ != TypeUndefined
}

func (v Value) Bool() bool {
	if v.typ != TypeBoolean {
		panic(&ValueError{"Value.Bool", v.typ})
	}
	return v.v.(bool)
}
