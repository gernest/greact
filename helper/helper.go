package helper

import (
	"net/url"
	"regexp"

	"github.com/gernest/naaz/prop"
	"github.com/gernest/prom"
	"github.com/gernest/xhr"
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
}

// Render implements vecty.Component interface.
func (c *ComponentRunner) Render() vecty.ComponentOrHTML {
	n := c.Next()
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
	cmp := c.cmp()
	return elem.Body(cmp)
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

type App struct {
	vecty.Core

	Units []*prom.T

	// A regular expression to specify which unit tests to run.Tests whose top level
	// description matches will be run.Not macthing tests are simply ignored.
	UnitMatch string

	// A regular expression to specify which integration tests to run.Tests whose
	// top level description matches will be run.Not macthing tests are simply
	// ignored.
	IntegrationMatch string

	// Integrations names of the functions to run integration tests.
	Integrations []string
}

// Mount implements vecty.Mounter. Unit tests are executed here. Integration
// tests arerun in iframes so the handling is done in the Render method.
func (a *App) Mount() {
	x := a.UnitMatch
	if x == "" {
		x = ".*"
	}
	re := regexp.MustCompile(x)
	var list []*prom.T
	for _, v := range a.Units {
		if re.MatchString(v.Suite.Desc) {
			list = append(list, v)
		}
	}
	rs := prom.Exec(list...)
	a.handleUnitsResult(rs)
}

func (a *App) handleUnitsResult(rs *prom.ResultCtx) {
}

func (a *App) handleIntegrationResult(rs *js.Object) {
}

// Render implements vecty.Component interface.
func (a *App) Render() vecty.ComponentOrHTML {
	if len(a.Integrations) > 0 {
		x := a.UnitMatch
		if x == "" {
			x = ".*"
		}
		re := regexp.MustCompile(x)
		var list []string
		for _, v := range a.Integrations {
			if re.MatchString(v) {
				list = append(list, v)
			}
		}
		var ls vecty.List
		for _, v := range list {
			src := a.iframeSrc(v)
			ls = append(ls, elem.InlineFrame(
				vecty.Markup(
					prop.Src(src),
				),
			))
		}
		return elem.Body(ls)
	}
	return elem.Body()
}

func (a *App) iframeSrc(fn string) string {
	return ""
}

func Run() {
	h := js.Global.Get("location").Get("href").String()
	println(h)
	u, err := url.Parse(h)
	if err != nil {
		panic(err)
	}
	go handle(u)
}

func handle(u *url.URL) {
	u.Path = "info"
	req := xhr.NewRequest("GET", u.String())
	if err := req.Send(nil); err != nil {
		panic(err)
	}
	println(req.ResponseText)
}
