package vected

import (
	"fmt"
	"reflect"
)

// Props is a map of properties. These are used to pass values to components.
type Props map[string]interface{}

// MergeProps returns new Props with values from b added to a.
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

func (p Props) String(key string) string {
	if v, ok := p[key]; ok {
		return v.(string)
	}
	return ""
}
