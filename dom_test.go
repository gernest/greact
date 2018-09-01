package vected

import (
	"testing"

	"github.com/gernest/vected/vdom/testutil"
)

func TestSetAccessor(t *testing.T) {
	t.Run("should set classname", func(ts *testing.T) {
		e := testutil.NewObject()
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
		e := testutil.NewObject()
		SetAccessor(nil, e, "style", nil, text, false)
		v := e.Get("style").Get("cssText").String()
		if v != text {
			t.Error("expected style.cssText to be set")
		}
	})
}
