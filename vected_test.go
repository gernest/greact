package vected

import (
	"context"
	"testing"

	"github.com/gernest/vected/prop"
	"github.com/gernest/vected/state"
	"github.com/gernest/vected/vdom"
	"github.com/gernest/vected/vdom/testutil"
)

var _ Component = (*A)(nil)

type A struct {
	Core
	cb func()
}

func (*A) Template() string {
	return ``
}
func (a *A) Render(context.Context, prop.Props, state.State) *vdom.Node {
	return nil
}
func (a *A) ComponentDidMount() {
	if a.cb != nil {
		a.cb()
	}
}

func TestDOM(t *testing.T) {
	t.Run("removeChildren", func(ts *testing.T) {
		e := testutil.NewObject()
		for i := 0; i < 5; i++ {
			e.Call("appendChild", testutil.NewObject())
		}
		removeChildren(e)
		ch := e.Get("childNodes").Get("length").Int()
		expect := 0
		if ch != expect {
			ts.Errorf("expected %d got %d", expect, ch)
		}
	})
}
