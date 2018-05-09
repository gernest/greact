package helper

import (
	"github.com/gernest/prom"
	"github.com/gopherjs/gopherjs/js"
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
)

// ComponentRunner is a vecty component for running integration tests. This
// doesn't handle collection of results. You need to supply AfterFunc as a
// callback, which will be called whenever a component test suite is complete.
type ComponentRunner struct {
	vecty.Core

	Next func() vecty.ComponentOrHTML

	AfterFunc func(*prom.ResultCtx)

	// This when set will be called when all the components retruned by next have
	// been successfully mounted and the testcases executed.
	Done func()
}

// Render implements vecty.Component interface.
func (c *ComponentRunner) Render() vecty.ComponentOrHTML {
	n := c.Next()
	if n == nil {
		if c.Done != nil {
			c.Done()
		}
		return nil
	}

	// safe to do this. Only the component struct implements Integration interface.
	cmp := n.(*component)
	cmp.after = c.after
	return cmp

}

func (c *ComponentRunner) after(rs *prom.ResultCtx) {
	if c.AfterFunc != nil {
		c.AfterFunc(rs)
	}
	// trigger rendering the next component.We are sure now that the test suite for
	// the previous componet was complete.
	vecty.Rerender(c)
}

type component struct {
	vecty.Core
	id     string
	cmp    func() vecty.ComponentOrHTML
	isBody bool
	cases  prom.List
	after  func(*prom.ResultCtx)
}

func (c *component) Mount() {
	node := js.Global.Get("document").Get("body")
	if !c.isBody {
		node = node.Get("firstChild")
	}

	s := &prom.Suite{Desc: c.id, Cases: c.cases, ResultFN: func() prom.Result {
		return &prom.RSWithNode{Object: node}
	}}
	rs := prom.ExecSuite(s)
	if c.after != nil {
		c.after(rs)
	}
}

func (c *component) Render() vecty.ComponentOrHTML {
	if c.isBody {
		return c.cmp()
	}
	return elem.Body(c.cmp())
}

func Wrap(i prom.Integration) *component {
	v := i.(*prom.Component)
	return &component{
		id: v.ID,
		cmp: func() vecty.ComponentOrHTML {
			return v.Component().(vecty.ComponentOrHTML)
		},
		isBody: v.IsBody,
		cases:  v.Cases,
	}
}

func NextFunc(v ...prom.Integration) func() vecty.ComponentOrHTML {
	i := 0
	return func() vecty.ComponentOrHTML {
		if i < len(v) {
			c := v[i]
			i++
			return Wrap(c)
		}
		return nil
	}
}
