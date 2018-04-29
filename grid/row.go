package grid

import (
	"github.com/gernest/gs"
	"github.com/gernest/vected/style/grid"
	"github.com/gernest/vected/ui"
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
)

type Row struct {
	vecty.Core
	Style    vecty.Applyer
	Children vecty.MarkupOrChild
	sheet    *gs.Sheet
	attached bool
}

func (r *Row) Render() vecty.ComponentOrHTML {
	if r.sheet == nil {
		r.sheet = ui.NewSheet()
		r.sheet.AddRule(r.style())
	}
	classes := vecty.ClassMap(r.sheet.CLasses.Classes())
	return elem.Div(vecty.Markup(classes, r.Style), r.Children)
}

func (c *Row) style() gs.CSSRule {
	return grid.Row()
}

func (r *Row) Mount() {
	r.sheet.Attach()
}

func (r *Row) Unmount() {
	r.sheet.Detach()
}
