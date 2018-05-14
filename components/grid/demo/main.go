package main

import (
	"github.com/gernest/vected/grid/demo/components"
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
)

func main() {
	vecty.RenderBody(&base{})
}

type base struct {
	vecty.Core
}

func (base) Render() vecty.ComponentOrHTML {
	return elem.Body(
		&components.BasicGrid{},
		// &components.Grid2{},
		// &components.Grid3{},
	)
}
