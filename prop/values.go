package prop

import (
	"fmt"
	"reflect"
)

// Type defines supported prop kinds to offer basic type safety. Props are
// passed on interface{} value which can be anything, this allows components to
// state what kind of prop they expect.
type Type = reflect.Kind

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
// type safety of the passed props.
func Verify(key string, typ Type, props Props) error {
	if v, ok := props[key]; ok {
		kind := reflect.TypeOf(v).Kind()
		if kind != typ {
			return fmt.Errorf(verifyErr, key, typ, kind)
		}
	}
	return nil
}
