package components

import (
	"github.com/gernest/gs"
	"github.com/gernest/vected/components/grid"
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
)

func col(span grid.Number, gutter int64, light bool) vecty.Component {
	var m gs.CSSRule
	if light {
		m = style1()
	} else {
		m = style2()
	}
	return &grid.Column{
		Span:   span,
		Gutter: gutter,
		Children: func() vecty.MarkupOrChild {
			return vecty.Text(span.String())
		},
		CSS: m,
	}
}

func col2(span grid.Number, gutter int64, light bool) vecty.Component {
	m := gs.S(".Style1",
		gs.P("color", "#fff"),
		gs.P("text-align", "center"),
		gs.P("padding", "30px 0"),
		gs.P("font-size", "18px"),
		gs.P("border", "none"),
		gs.P("margin-top", "8px"),
		gs.P("margin-bottom", "8px"),
		gs.P("height", "15px"),
	)
	return &grid.Column{
		Span:   span,
		Gutter: gutter,
		Children: func() vecty.MarkupOrChild {
			return gutterBox(span.String())
		},
		CSS: m,
	}
}

func col3(span, offset grid.Number, gutter int64, light bool) vecty.Component {
	var m gs.CSSRule
	if light {
		m = style1()
	} else {
		m = style2()
	}
	return &grid.Column{
		Span:   span,
		Offset: offset,
		Gutter: gutter,
		Children: func() vecty.MarkupOrChild {
			return gutterBox(span.String())
		},
		CSS: m,
	}
}

func style1() gs.CSSRule {
	return gs.S(".Style1",
		gs.P("color", "#fff"),
		gs.P("background", "rgba(0,160,233,.7)"),
		gs.P("text-align", "center"),
		gs.P("padding", "30px 0"),
		gs.P("font-size", "18px"),
		gs.P("border", "none"),
		gs.P("margin-top", "8px"),
		gs.P("margin-bottom", "8px"),
		gs.P("height", "15px"),
	)
}
func style2() gs.CSSRule {
	return gs.S(".Style1",
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

func codeBoxDemo() gs.CSSRule {
	img := `linear-gradient(90deg, #f5f5f5 4.16666667%, transparent 4.16666667%, transparent 8.33333333%, #f5f5f5 8.33333333%, #f5f5f5 12.5%, transparent 12.5%, transparent 16.66666667%, #f5f5f5 16.66666667%, #f5f5f5 20.83333333%, transparent 20.83333333%, transparent 25%, #f5f5f5 25%, #f5f5f5 29.16666667%, transparent 29.16666667%, transparent 33.33333333%, #f5f5f5 33.33333333%, #f5f5f5 37.5%, transparent 37.5%, transparent 41.66666667%, #f5f5f5 41.66666667%, #f5f5f5 45.83333333%, transparent 45.83333333%, transparent 50%, #f5f5f5 50%, #f5f5f5 54.16666667%, transparent 54.16666667%, transparent 58.33333333%, #f5f5f5 58.33333333%, #f5f5f5 62.5%, transparent 62.5%, transparent 66.66666667%, #f5f5f5 66.66666667%, #f5f5f5 70.83333333%, transparent 70.83333333%, transparent 75%, #f5f5f5 75%, #f5f5f5 79.16666667%, transparent 79.16666667%, transparent 83.33333333%, #f5f5f5 83.33333333%, #f5f5f5 87.5%, transparent 87.5%, transparent 91.66666667%, #f5f5f5 91.66666667%, #f5f5f5 95.83333333%, transparent 95.83333333%)`
	return gs.S(".CodeBoxDemo",
		gs.P("background-image", img),
		gs.P("overflow", "hidden"),
		gs.P("margin-bottom", "8px"),
	)
}

type Grid2 struct {
	vecty.Core
}

func (Grid2) Render() vecty.ComponentOrHTML {
	return elem.Div(
		&grid.Row{
			CSS:    codeBoxDemo(),
			Gutter: 16,
			Children: func() vecty.MarkupOrChild {
				return vecty.List{
					col2(grid.G6, 16, true),
					col2(grid.G6, 16, true),
					col2(grid.G6, 16, true),
					col2(grid.G6, 16, true),
				}
			},
		},
	)
}

func gutterBox(text string) *vecty.HTML {
	return elem.Div(
		vecty.Markup(
			vecty.Style("background", "#00A0E9"),
			vecty.Style("padding", "5px 0"),
		),
		vecty.Text(text),
	)
}

type Grid3 struct {
	vecty.Core
}

func (Grid3) Render() vecty.ComponentOrHTML {
	return elem.Div(
		&grid.Row{
			CSS:    codeBoxDemo(),
			Gutter: 16,
			Children: func() vecty.MarkupOrChild {
				return vecty.List{
					col(grid.G8, 0, true),
					col3(grid.G8, grid.G8, 0, false),
				}
			},
		},
		&grid.Row{
			CSS:    codeBoxDemo(),
			Gutter: 16,
			Children: func() vecty.MarkupOrChild {
				return vecty.List{
					col3(grid.G6, grid.G6, 0, false),
					col3(grid.G6, grid.G6, 0, false),
				}
			},
		},
		&grid.Row{
			CSS:    codeBoxDemo(),
			Gutter: 16,
			Children: func() vecty.MarkupOrChild {
				return vecty.List{
					col3(grid.G12, grid.G6, 0, false),
				}
			},
		},
	)
}
