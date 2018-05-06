package divider

import (
	"github.com/gernest/gs"
	"github.com/gernest/vected/style/divider"
	"github.com/gernest/vected/ui"
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
)

type Kind int64

const (
	Horizontal Kind = iota
	Vertical
)

type Orientation int64

const (
	Left Orientation = iota + 1
	Right
)

func (o Orientation) String() string {
	switch o {
	case 1:
		return divider.WithTextLeft
	case 2:
		return divider.WithTextRight
	default:
		return divider.WithText
	}
}

func (o Kind) String() string {
	switch o {
	case 0:
		return divider.Horizontal
	case 1:
		return divider.Vertical
	default:
		return ""
	}
}

type Divider struct {
	vecty.Core
	Type        Kind
	Orientation Orientation
	Dashed      bool
	// These are additional style rules to apply. If provided the rule will be
	// compiled and the resulting classes will be applied to the <div></div> html
	// element rendered.
	//
	// Please define only classes/selectors here with gs.S, for trivial css
	// properties change use the Style field.
	CSS gs.CSSRule

	// Markup style applied to the divider's root  <div></div> element.
	Style vecty.Applyer

	Children func() vecty.MarkupOrChild
	sheet    *gs.Sheet
}

func (d *Divider) Render() vecty.ComponentOrHTML {
	if d.sheet == nil {
		d.sheet = ui.NewSheet()
		d.sheet.AddRule(divider.Style())
	}
	var commonClass vecty.ClassMap
	if d.CSS != nil {
		commonClass = vecty.ClassMap(d.sheet.AddRule(d.CSS).Classes())
	}
	children := d.getChildren()
	cn := d.sheet.CLasses[divider.BaseClass]
	cn = cn[1:]
	cls := vecty.ClassMap{
		cn: true,
		join(cn, d.Type.String()):        true,
		join(cn, d.Orientation.String()): children != nil,
		join(cn, divider.Dashed):         !!d.Dashed,
	}
	return elem.Div(
		vecty.Markup(cls, commonClass),
		vecty.If(children != nil,
			elem.Span(
				vecty.Markup(
					vecty.Class(join(cn, divider.InnerText)),
				),
				children,
			),
		),
	)
}

func (d *Divider) getChildren() vecty.MarkupOrChild {
	if d.Children != nil {
		return d.Children()
	}
	return nil
}

func join(v ...string) string {
	o := ""
	for _, s := range v {
		o += s
	}
	return o
}

func (d *Divider) Mount() {
	d.sheet.Attach()
}
func (d *Divider) Unmount() {
	d.sheet.Detach()
}
