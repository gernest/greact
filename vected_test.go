package vected

import (
	"container/list"
	"context"
	"testing"

	"github.com/gernest/vected/prop"
	"github.com/gernest/vected/state"
	"github.com/gernest/vected/vdom"
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

func TestFlashMounts(t *testing.T) {
	defer func() {
		mounts = list.New()
	}()
	t.Run("should pop mounted components", func(ts *testing.T) {
		for i := 0; i < 5; i++ {
			mounts.PushBack(&A{})
		}
		flushMounts()
		if mounts.Len() != 0 {
			t.Errorf("expected 0 mounts but got %d", mounts.Len())
		}
	})
	t.Run("should call ComponentDidMount", func(ts *testing.T) {
		n := 0
		size := 5
		for i := 0; i < size; i++ {
			mounts.PushBack(&A{cb: func() {
				n++
			}})
		}
		flushMounts()
		if n != size {
			t.Errorf("expected %d mounts but got %d", size, n)
		}
	})
}
