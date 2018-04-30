package grid

import (
	"strconv"

	"github.com/gernest/gs"
	"github.com/gernest/vected/style/grid"
	"github.com/gernest/vected/ui"
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
)

type FlexStyle = grid.FlexStyle

const (
	Start FlexStyle = iota
	End
	Center
	SpaceAround
	SpaceBetween
)

type FlexAlign = grid.FlexAlign

const (
	Top FlexAlign = iota
	Middle
	Bottom
)

// Row is a vecty component using ant design to render a flex row grid layout.
// This component uses gs library to bundle the styles with the component so no
// need for external css.
type Row struct {
	vecty.Core

	// This will be  appkied to the row's <div>
	Style vecty.Applyer

	// Optional styles to be attached to this component's style sheet. Forbetter
	// results supply class selectors gs.S()
	CSS gs.CSSRule

	// Children is a function which returns components to be rendered inside to
	// row.
	// For consistency and better results this should return vecty.List of *Column
	// component. It is fine to mix Columns with other components and they will be
	// rendered correctly.
	Children func() vecty.MarkupOrChild

	// This is spacing between grids
	Gutter int64

	// Flex uses flex layout when this field is set to true. You can use Justify
	// and Align to control the layout.
	//
	// Default is false.
	Flex bool

	// Justify is horizontal arrangement of the flex layout: tart end center
	// space-around space-between
	//
	// Default is Start
	Justify FlexStyle

	// Align is the vertical alignment of the flex layout: top middle bottom
	//
	// Default is top
	Align FlexAlign
	sheet *gs.Sheet
}

func (r *Row) Render() vecty.ComponentOrHTML {
	if r.sheet == nil {
		r.sheet = ui.NewSheet()
		r.sheet.AddRule(r.style())
		if r.CSS != nil {
			r.sheet.AddRule(r.CSS)
		}
	}
	ch := r.getChildren()
	style := r.Style
	if r.Gutter > 0 {
		style = vecty.Markup(
			vecty.Style("margin-left", format(r.Gutter/-2)+"px"),
			vecty.Style("margin-right", format(r.Gutter/-2)+"px"),
		)
		if ls, ok := ch.(vecty.List); ok {
			var o vecty.List
			for _, v := range ls {
				if col, ok := v.(*Column); ok {
					col.Gutter = r.Gutter
					v = col
				}
				o = append(o, v)
			}
			ch = o
		}
	}
	classes := vecty.ClassMap(r.sheet.CLasses.Classes())
	return elem.Div(vecty.Markup(classes, style), ch)
}

func (r *Row) getChildren() vecty.MarkupOrChild {
	if r.Children != nil {
		return r.Children()
	}
	return nil
}

func format(v int64) string {
	return strconv.FormatInt(v, 10)
}

func (c *Row) style() gs.CSSRule {
	return grid.Row(c.Gutter, c.Flex, c.Justify, c.Align)
}

func (r *Row) Mount() {
	r.sheet.Attach()
}

func (r *Row) Unmount() {
	r.sheet.Detach()
}
