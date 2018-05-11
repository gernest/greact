package test

import (
	"github.com/gernest/prom"
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
)

func TestRenderBody() prom.Integration {
	return prom.RenderBody("Body", helloFunc,
		prom.It("says hello world", func(rs prom.T) {
			node := rs.(prom.Node).Node()
			txt := "hello, world"
			g := node.Get("textContent").String()
			if g != txt {
				rs.Errorf("expected %s got %s", txt, g)
			}
		}),
	)
}

func helloFunc() interface{} {
	return &hello{}
}

type hello struct {
	vecty.Core
}

func (hello) Render() vecty.ComponentOrHTML {
	return elem.Body(
		elem.Div(
			vecty.Text("hello, world"),
		),
	)
}

func TestRender() prom.Integration {
	txt := "hello, world 2"
	return prom.Render("HTML node", renderFunc(txt),
		prom.It("says hello world", func(rs prom.T) {
			node := rs.(prom.Node).Node()
			g := node.Get("textContent").String()
			if g != txt {
				rs.Errorf("expected %s got %s", txt, g)
			}
		}),
	)
}

func renderFunc(txt string) func() interface{} {
	return func() interface{} {
		return elem.Div(vecty.Text(txt))
	}
}
