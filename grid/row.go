package grid

import (
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
	Children vecty.MarkupOrChild
	sheet    *gs.Sheet
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
