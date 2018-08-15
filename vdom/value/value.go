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
