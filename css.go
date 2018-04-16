package goss

//go:generate go run bin/cssprops/main.go
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
