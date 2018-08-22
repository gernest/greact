package vected

import (
	"container/list"
	"context"

	"github.com/gernest/vected/prop"
)

var recyclerComponents = list.New()

func createComponent(ctx context.Context, cmp Component, props prop.Props) Component {
	var ncmp Component
	if in, ok := cmp.(Constructor); ok {
		ncmp = in.New()
	}
	core := ncmp.core()
	core.context = ctx
	core.props = props
	return ncmp
}
