package components

import (
	"github.com/gernest/gs"
	"github.com/gernest/vected/grid"
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
)

type BasicGrid struct {
	vecty.Core
}

func BasicGridRow(span grid.Number) vecty.Component {
	return &grid.Column{
		Span: span,
		Children: func() vecty.MarkupOrChild {
			return vecty.Text(span.String())
		},
	}
}
func (BasicGrid) Render() vecty.ComponentOrHTML {
	return elem.Div(
		&grid.Row{
			CSS: codeBoxDemo(),
			Children: func() vecty.MarkupOrChild {
				return vecty.List{
					BasicGridRow(grid.G12),
					BasicGridRow(grid.G12),
				}
			},
		},
		&grid.Row{
			CSS: codeBoxDemo(),
			Children: func() vecty.MarkupOrChild {
				return vecty.List{
					BasicGridRow(grid.G8),
					BasicGridRow(grid.G8),
					BasicGridRow(grid.G8),
				}
			},
		},
		&grid.Row{
			CSS: codeBoxDemo(),
			Children: func() vecty.MarkupOrChild {
				return vecty.List{
					BasicGridRow(grid.G6),
					BasicGridRow(grid.G6),
					BasicGridRow(grid.G6),
					BasicGridRow(grid.G6),
				}
			},
		},
	)
}

func styleBasic() gs.CSSRule {
	return gs.S(".BasicStyle",
		gs.P("color", "#fff"),
		gs.P("background", "#00a0e9"),
		gs.P("text-align", "center"),
		gs.P("padding", "30px 0"),
		gs.P("font-size", "18px"),
		gs.P("border", "none"),
		gs.P("margin-top", "8px"),
		gs.P("margin-bottom", "8px"),
		gs.P("height", "15px"),
	)
}
