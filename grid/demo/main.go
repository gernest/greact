package main

import (
	"github.com/gernest/naaz/style"
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
	return elem.Body(
		elem.Div(
			&grid.Row{
				Children: vecty.List{
					col(grid.G12, true),
					col(grid.G12, false),
				},
			},
			&grid.Row{
				Children: vecty.List{
					col(grid.G8, true),
					col(grid.G8, false),
					col(grid.G8, true),
				},
			},
			&grid.Row{
				Children: vecty.List{
					col(grid.G6, true),
					col(grid.G6, false),
					col(grid.G6, true),
					col(grid.G6, false),
				},
			},
		),
	)
}

func col(span grid.Number, light bool) vecty.Component {
	blue := vecty.Markup(
		vecty.Style("color", "#fff"),
		vecty.Style("background", "rgba(0,160,233,.7)"),
	)
	lightBlue := vecty.Markup(
		vecty.Style("background", "#00a0e9"),
		vecty.Style("color", "#fff"),
	)
	st := vecty.Markup(
		vecty.Style("text-align", "center"),
		vecty.Style("padding", "30px 0"),
		vecty.Style("font-size", "18px"),
		vecty.Style("border", "none"),
		vecty.Style("margin-top", "8px"),
		vecty.Style("margin-bottom", "8px"),
		style.Height(style.Px(15)),
	)
	var m vecty.Applyer
	if light {
		m = vecty.Markup(blue, st)
	} else {
		m = vecty.Markup(lightBlue, st)
	}
	c := &grid.Column{
		Children: vecty.Text(span.String()),
		Style:    m,
	}
	c.Span = span
	return c
}
