package vected

import "testing"

func TestObject(t *testing.T) {
	t.Run("hasOwnProperty", func(ts *testing.T) {
		o := newObject()
		o.Set("a", true)
		if !o.Call("hasOwnProperty", "a").Bool() {
			t.Errorf("expected to have a property set")
		}
		if o.Call("hasOwnProperty", "b").Bool() {
			t.Errorf("expected to not have b property set")
		}
	})
}
