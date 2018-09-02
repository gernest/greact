package vected

import (
	"fmt"
	"html/template"
	"reflect"
	"sort"
	"strings"
)

// Props is a map of properties. These are used to pass values to components.
type Props map[interface{}]interface{}

// Merge returns new Props with values from b added to a.
func MergeProps(a, b Props) Props {
	m := make(Props)
	for k, v := range a {
		m[k] = v
	}
	for k, v := range b {
		m[k] = v
	}
	return m
}

// Filter returns a new Props with only values that the filter function fn
// evaluates to true.
func (p Props) Filter(fn func(k, v interface{}) bool) Props {
	m := make(Props)
	for k, v := range p {
		if fn(k, v) {
			m[k] = v
		}
	}
	return m
}

// Attr returns a string for prop attributes. This only collect string keys. For
// boolean values the attribute will not contain the value part. Other types of
// kys/values are simply ignored.
//
// 	p["checked"]=true => checked
// 	p["onClick"]="handleClick" => onClick="handleClick"
func (p Props) Attr() template.HTMLAttr {
	var keys []string
	for k := range p {
		if s, ok := k.(string); ok {
			keys = append(keys, s)
		}
	}
	sort.Strings(keys)
	var o []string
	for _, k := range keys {
		v := p[k]
		switch e := v.(type) {
		case string:
			o = append(o, fmt.Sprintf(`%v="%s"`, k, e))
		case bool:
			o = append(o, fmt.Sprint(k))
		}
	}
	return template.HTMLAttr(strings.Join(o, " "))
}

// Int calls Int with p as first argument.
func (p Props) Int(key interface{}) NullInt {
	return GetInt(p, key)
}

// StringV calls StringV with p as first argument.
func (p Props) StringV(key interface{}, value string) NullString {
	return GetStringV(p, key, value)
}

// Int looks for property value with key, and tries to cast it to an int. This
// will set NullInt.IsNull to true if the key is missing or the value is not of
// type it.
func GetInt(p Props, key interface{}) NullInt {
	if v, ok := p[key]; ok {
		if vi, ok := v.(int); ok {
			return NullInt{Value: vi}
		}
	}
	return NullInt{IsNull: true}
}

// StringV looks for key's value in ctx. This will return value if the key is
// missing, this function will try to cast the value to string.When the value is
// not of type string the NullString.IsNull field will be set to true.
func GetStringV(ctx Props, key interface{}, value string) NullString {
	if v, ok := ctx[key]; ok {
		if vk, ok := v.(string); ok {
			return NullString{Value: vk}
		}
		return NullString{IsNull: true}
	}
	return NullString{Value: value}
}

// String tries to return the value stored by key as a string.
func (p Props) String(key interface{}) NullString {
	return GetString(p, key)
}

// String finds key's value in p and casts it as a string.
func GetString(p Props, key interface{}) NullString {
	if v, ok := p[key]; ok {
		if vi, ok := v.(string); ok {
			return NullString{Value: vi}
		}
	}
	return NullString{IsNull: true}
}

// Bool tries to find key's value in p and casts it to bool if found.
func (p Props) Bool(key interface{}) NullBool {
	return GetBool(p, key)
}

// Bool finds key's value in p and casts it as a bool.
func GetBool(p Props, key interface{}) NullBool {
	if v, ok := p[key]; ok {
		if vi, ok := v.(bool); ok {
			return NullBool{Value: vi}
		}
	}
	return NullBool{IsNull: true}
}

// Children returns child components stored in props.
func (p Props) Children() []*Node {
	if c, ok := p["children"]; ok {
		return c.([]*Node)
	}
	return nil
}

// PropType defines supported prop kinds to offer basic type safety. Props are
// passed on interface{} value which can be anything, this allows components to
// state what kind of prop they expect.
type PropType = reflect.Kind

// supported prop types
const (
	Bool       = reflect.Bool
	Int        = reflect.Int
	Int8       = reflect.Int8
	Int16      = reflect.Int16
	Int32      = reflect.Int32
	Int64      = reflect.Int64
	Uint       = reflect.Uint
	Uint8      = reflect.Uint8
	Uint16     = reflect.Uint16
	Uint32     = reflect.Uint32
	Uint64     = reflect.Uint64
	Float32    = reflect.Float32
	Float64    = reflect.Float64
	Complex64  = reflect.Complex64
	Complex128 = reflect.Complex128
	Array      = reflect.Array
	Chan       = reflect.Chan
	Func       = reflect.Func
	Interface  = reflect.Interface
	Map        = reflect.Map
	Ptr        = reflect.Ptr
	Slice      = reflect.Slice
	String     = reflect.String
	Struct     = reflect.Struct
)

type NullInt struct {
	Value  int
	IsNull bool
}

type NullString struct {
	Value  string
	IsNull bool
}

type NullBool struct {
	Value  bool
	IsNull bool
}

const verifyErr = "key: %s, has wrong prop value expected %s got %s instead"

// Verify verifies that the prop value for key is of kind typ. This offers basic
// type safety of the passedro props.
func Verify(key string, typ PropType, props Props) error {
	if v, ok := props[key]; ok {
		kind := reflect.TypeOf(v).Kind()
		if kind != typ {
			return fmt.Errorf(verifyErr, key, typ, kind)
		}
	}
	return nil
}
