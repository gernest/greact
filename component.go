package vected

import (
	"container/list"
	"context"
	"reflect"

	"github.com/gernest/vected/prop"
)

var recyclerComponents = list.New()

func createComponent(ctx context.Context, cmp Component, props prop.Props) Component {
	var ncmp Component
	if in, ok := cmp.(Constructor); ok {
		ncmp = in.New()
	} else {
		// we use reflection to create a new component
		v := reflect.ValueOf(cmp)
		if v.Kind() != reflect.Ptr || v.Elem().Kind() != reflect.Struct {
			panic("component must be pointer to struct")
		}
		ncmp = reflect.New(v.Elem().Type()).Interface().(Component)
	}
	core := ncmp.core()
	core.context = ctx
	core.props = props
	return ncmp
}
