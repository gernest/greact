package divider

import (
	"github.com/gernest/gs"
	"github.com/gernest/vected/style/divider"
	"github.com/gernest/vected/ui"
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
)

// Kind is the type of the divider.
type Kind int64

const (
	//Horizontal this draws a horizontal divider.
	Horizontal Kind = iota

	// Vertical draws vertical divider.
	Vertical
)

// Orientation is where the text node within the divider element is placed.
type Orientation int64

const (
	// Left the text is placed to the left.
	Left Orientation = iota + 1

	// Right the text is placed to the right.
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

// Divider is a component that draws a line to separate different content.
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

	// The text node component which is wrapped in a <span> element. You can be
	// creative and render something else.
	Children func() vecty.MarkupOrChild
	sheet    *gs.Sheet
}

// Render implements vecty.Component interface.
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

// Mount implements vecty.Mounter interface.This attaches the component's style
// sheet.
func (d *Divider) Mount() {
	d.sheet.Attach()
}

// Unmount implements vecty.Unmounter interface.This detach the component's style
// sheet.
func (d *Divider) Unmount() {
	d.sheet.Detach()
}
