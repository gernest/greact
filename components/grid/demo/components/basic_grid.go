package components

import (
	"github.com/gernest/gs"
	"github.com/gernest/vected/components/grid"
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
)

type BasicGrid struct {
	vecty.Core
}

func BasicGridRow(span grid.Number, fn func() gs.CSSRule) vecty.Component {
	return &grid.Column{
		Span: span,
		CSS:  fn(),
		Children: func() vecty.MarkupOrChild {
			return vecty.Text(span.String())
		},
	}
}
func (BasicGrid) Render() vecty.ComponentOrHTML {
	style := styleBasic()
	return elem.Div(
		&grid.Row{
			CSS: codeBoxDemo(),
			Children: func() vecty.MarkupOrChild {
				return vecty.List{
					BasicGridRow(grid.G12, style),
					BasicGridRow(grid.G12, style),
				}
			},
		},
		&grid.Row{
			CSS: codeBoxDemo(),
			Children: func() vecty.MarkupOrChild {
				return vecty.List{
					BasicGridRow(grid.G8, style),
					BasicGridRow(grid.G8, style),
					BasicGridRow(grid.G8, style),
				}
			},
		},
		&grid.Row{
			CSS: codeBoxDemo(),
			Children: func() vecty.MarkupOrChild {
				return vecty.List{
					BasicGridRow(grid.G6, style),
					BasicGridRow(grid.G6, style),
					BasicGridRow(grid.G6, style),
					BasicGridRow(grid.G6, style),
				}
			},
		},
	)
}

func styleBasic() func() gs.CSSRule {
	on := false
	bg1 := gs.P("background", "rgba(0,160,233,.7)")
	bg2 := gs.P("background", "#00a0e9")
	return func() gs.CSSRule {
		bg := bg1
		if on {
			bg = bg2
		}
		on = !on
		return gs.S(".BasicStyle",
			gs.P("color", "#fff"),
			bg,
			gs.P("text-align", "center"),
			gs.P("padding", "30px 0"),
			gs.P("font-size", "18px"),
			gs.P("border", "none"),
			gs.P("margin-top", "8px"),
			gs.P("margin-bottom", "8px"),
			gs.P("height", "15px"),
		)
	}

}
