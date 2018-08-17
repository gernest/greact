package dom

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
}
