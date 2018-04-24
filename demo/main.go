package main

import (
	"github.com/gernest/gs"
	"github.com/gernest/gs/normalize"
	"github.com/gernest/naaz"
	"github.com/gernest/naaz/elem"
)

func main() {
	naaz.RenderBody(&Hello{})
}

type Hello struct {
	naaz.Core
	sheet *gs.Sheet
}

func (h *Hello) Mount() {
	s := naaz.NewSheet()
	s.AddRule(normalize.New())
	s.Attach()
	h.sheet = s
}

func (Hello) Render() naaz.ComponentOrHTML {
	return elem.Body(
		naaz.Text("hello, world"),
	)
}

func (h *Hello) Unmount() {
	h.sheet.Detach()
}
