package grid

import (
	"strconv"

	"github.com/gernest/gs"
	"github.com/gernest/vected/style/grid"
	"github.com/gernest/vected/ui"
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
)

// FlexStyle is flexbox layout justification
type FlexStyle = grid.FlexStyle

const (
	// Start sets justify-content:flex-start
	Start FlexStyle = iota

	// End sets justify-content:flex-end
	End

	// Center sets justify-content:center
	Center

	// SpaceAround sets justify-content:space-around
	SpaceAround

	// SpaceBetween sets justify-content:space-between
	SpaceBetween
)

// FlexAlign is flexbox layout  alignment
type FlexAlign = grid.FlexAlign

const (
	// Top sets align-items: lex-start
	Top FlexAlign = iota

	// Middle sets align-items: center
	Middle

	// Bottom sets align-items: lex-end
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

// Render implements vecty.Component interface.
//
// If Gutter >0 margins are computed based on the gutter size and styles applied
// directly on the row's div. In case childerns are of type *Column then the new
// gutter size is applied before rendering of the children's.
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

func (r *Row) style() gs.CSSRule {
	return grid.Row(r.Gutter, r.Flex, r.Justify, r.Align)
}

// Mount attaches component's stylesheets.
func (r *Row) Mount() {
	r.sheet.Attach()
}

// Unmount detach component's stylesheets
func (r *Row) Unmount() {
	r.sheet.Detach()
}
