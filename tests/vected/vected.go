package vected

import (
	"github.com/gernest/mad"
	"github.com/gernest/vected"
	"github.com/gernest/vected/props"
)

type hello struct {
}

func (hello) ID() string {
	return "Hello"
}
func (hello) Template() string {
	return `<div>hello,world {.name}</div>`
}
func (hello) Render(ctx props.Props) props.Props {
	return props.Props{"name": "gernest"}
}

func TestRenderHTML() mad.Test {
	return mad.It("must render template ", func(t mad.T) {
		vected.Register(hello{})
		v, err := vected.RenderHTML(`{Hello .}`, nil)
		if err != nil {
			t.Fatal(err)
		}
		expect := `<div>hello,world gernest</div>`
		got := string(v)
		if got != expect {
			t.Errorf("expected %s got %s")
		}
	})

}
