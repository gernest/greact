package main

import (
	"github.com/gernest/vected/grid"
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
	firstCol := &grid.Column{
		Children: vecty.Text("col-12"),
	}
	firstCol.Span = grid.G12
	secondCol := &grid.Column{
		Children: vecty.Text("col-12"),
	}
	secondCol.Span = grid.G12
	return elem.Body(
		elem.Div(
			&grid.Row{
				Children: vecty.List{
					firstCol,
					secondCol,
				},
			},
		),
	)
}
