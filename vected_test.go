package vected

import (
	"context"
	"fmt"
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
	return NewNode(ElementNode, "", "div", nil,
		NewNode(TextNode, "", "Hello,World", nil),
	)
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

func TestVected_Render(t *testing.T) {
	v := New()
	v.Document = NewObject()
	hello := NewNode(ElementNode, "", "div", nil,
		NewNode(TextNode, "", "Hello,World", nil),
	)
	el := NewObject()
	// err := wrapPanic(func() {
	v.Render(hello, el)
	// })
	// if err != nil {
	// 	t.Error(err)
	// }
}

func wrapPanic(fn func()) (err error) {
	defer func() {
		v := recover()
		if v != nil {
			err = fmt.Errorf("%v", v)
		}
	}()
	fn()
	return
}
func TestSetAccessor(t *testing.T) {
	t.Run("should set classname", func(ts *testing.T) {
		e := NewObject()
		SetAccessor(nil, e, "className", nil, "classa", false)
		v := e.Get("className").String()
		if v != "classa" {
			t.Error("expected className to be set")
		}
		SetAccessor(nil, e, "class", nil, "classb", false)
		v = e.Get("className").String()
		if v != "classb" {
			t.Error("expected className to be set")
		}
	})
	t.Run("should set style", func(ts *testing.T) {
		text := "color:blue;"
		e := NewObject()
		SetAccessor(nil, e, "style", nil, text, false)
		v := e.Get("style").Get("cssText").String()
		if v != text {
			t.Error("expected style.cssText to be set")
		}
	})
}
