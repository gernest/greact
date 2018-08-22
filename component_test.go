package vected

import (
	"context"
	"testing"

	"github.com/gernest/vected/prop"
	"github.com/gernest/vected/state"
	"github.com/gernest/vected/vdom"
)

type Hello struct {
	Core
}

func (*Hello) New() Component {
	return &Hello{}
}

var H = vdom.New
var HA = vdom.Attr
var HAT = vdom.Attrs

// Render implements vected.Renderer interface.
func (h *Hello) Render(ctx context.Context, props prop.Props, state state.State) *vdom.Node {
	return H(3, "", "div", nil,
		H(1, "", "hello,world", nil))
}

func TestCreateComponent(t *testing.T) {
	t.Run("must set initial context and props", func(ts *testing.T) {
		n := createComponent(context.Background(), &Hello{}, make(prop.Props))
		core := n.core()
		if core.context == nil {
			t.Error("expected context to be set")
		}
		if core.props == nil {
			t.Error("expected initial props to be set")
		}
	})
}
