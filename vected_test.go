package vected

import (
	"context"
	"testing"
)

var _ Component = (*A)(nil)

type A struct {
	Core
	cb func()
}

func (*A) Template() string {
	return ``
}
func (a *A) Render(context.Context, Props, State) *Node {
	return nil
}

func (a *A) ComponentDidMount() {
	if a.cb != nil {
		a.cb()
	}
}

func TestVected_createComponent(t *testing.T) {
	t.Run("must assign ctx ,props and id", func(ts *testing.T) {
		v := New()
		cmp := v.createComponent(context.Background(), &A{}, Props{
			"key": "value",
		})
		core := cmp.core()
		if core.context == nil {
			t.Error("expected context to be set")
		}
		if core.props == nil {
			t.Error("expected props to be set")
		}
		if core.id == 0 {
			t.Error("expected id to be set")
		}
	})
	t.Run("must set default constructor", func(ts *testing.T) {
		v := New()
		cmp := v.createComponent(context.Background(), &A{}, Props{
			"key": "value",
		})
		core := cmp.core()
		expect := "a"
		if core.constructor != expect {
			t.Errorf("expected %s got %s", expect, core.constructor)
		}
	})
	t.Run("must use provided constructor", func(ts *testing.T) {
		v := New()
		a := &A{}
		constructor := "A"
		a.core().constructor = constructor
		cmp := v.createComponent(context.Background(), a, Props{
			"key": "value",
		})
		core := cmp.core()
		if core.constructor != constructor {
			t.Errorf("expected %s got %s", constructor, core.constructor)
		}
	})
}
