package value

type Type int

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

type Value interface {
	Bool() bool
	Call(m string, args ...interface{}) Value
	Float() float64
	Get(string) Value
	Index(int) Value
	Int() int
	Invoke(args ...interface{}) Value
	Set(p string, x interface{})
	String() string
	Type() Type
}

type Callback interface {
	Release()
}

// Keys is like Object.keys, this returns nil if v is not an object.
func Keys(v Value) []string {
	if v.Type() != TypeObject {
		return nil
	}
	k := v.Call("keys")
	var o []string
	size := k.Get("length").Int()
	for i := 0; i < size; i++ {
		o = append(o, k.Index(i).String())
	}
	return o
}
