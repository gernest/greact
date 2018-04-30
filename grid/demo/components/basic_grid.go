package components

import (
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
