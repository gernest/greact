package text

import (
	"bytes"
	"testing"

	"github.com/gernest/vected/lib/vdom"
)

func TestRender(t *testing.T) {
	t.Run("simple components", func(ts *testing.T) {
		a := `<h1>hello,world</h1>`
		ctx := make(map[string]*vdom.Node)
		h, err := vdom.ParseString(a)
		if err != nil {
			ts.Fatal(err)
		}
		ctx["hello"] = h
		world, err := vdom.ParseString(`<Hello/>`)
		if err != nil {
			ts.Fatal(err)
		}
		var buf bytes.Buffer
		err = Render(&buf, world, ctx)
		if err != nil {
			ts.Fatal(err)
		}
		expect := `<h1>hello,world</h1>`
		got := buf.String()
		if got != expect {
			t.Errorf("expected %s got %s", expect, got)
		}
	})
}
