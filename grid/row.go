package grid

import (
	"strconv"

	"github.com/gernest/gs"
	"github.com/gernest/vected/style/grid"
	"github.com/gernest/vected/ui"
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
)

// Row is a vecty component using ant design to render a flex row grid layout.
// This component uses gs library to bundle the styles with the component so no
// need for external css.
type Row struct {
	vecty.Core
	Style    vecty.Applyer
	CSS      gs.CSSRule
	Children func() vecty.MarkupOrChild
	Gutter   int64
	Flex     bool
	sheet    *gs.Sheet
}

func (r *Row) Render() vecty.ComponentOrHTML {
	if r.sheet == nil {
		r.sheet = ui.NewSheet()
		r.sheet.AddRule(r.style())
		if r.CSS != nil {
			r.sheet.AddRule(r.CSS)
			println(r.sheet.ListRules())
		}
	}
	ch := r.getChildren()
	style := r.Style
	if r.Gutter > 0 {
		style = vecty.Markup(
			vecty.Style("margin-left", format(r.Gutter/-2)+"px"),
			vecty.Style("margin-right", format(r.Gutter/-2)+"px"),
		)
	}
	classes := vecty.ClassMap(r.sheet.CLasses.Classes())
	return elem.Div(vecty.Markup(classes, style), ch)
}

func (r *Row) getChildren() vecty.MarkupOrChild {
	if r.Children == nil {
		return nil
	}
	ch := r.Children()
	if r.Gutter > 0 {
		if ls, ok := ch.(vecty.List); ok {
			var o vecty.List
			for _, v := range ls {
				if col, ok := v.(*Column); ok {
					col.Gutter = r.Gutter
					v = col
				}
				o = append(o, v)
			}
			return o
		}
	}
	return ch
}

func format(v int64) string {
	return strconv.FormatInt(v, 10)
}

func (c *Row) style() gs.CSSRule {
	return grid.Row(c.Gutter, c.Flex)
}

func (r *Row) Mount() {
	r.sheet.Attach()
}

func (r *Row) Unmount() {
	r.sheet.Detach()
}
