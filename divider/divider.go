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
		return ""
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
	Style       vecty.Applyer
	CSS         gs.CSSRule
	Children    func() vecty.MarkupOrChild
	sheet       *gs.Sheet
	classMap    gs.ClassMap
}

func (d *Divider) Render() vecty.ComponentOrHTML {
	if d.sheet == nil {
		d.sheet = ui.NewSheet()
		d.classMap = d.sheet.AddRule(divider.Style())
	}
	children := d.getChildren()
	cn := d.classMap[divider.BaseClass]
	cn = cn[1:]
	cls := vecty.ClassMap{
		cn: true,
		join(cn, d.Type.String()):                          true,
		join(cn, divider.WithText, d.Orientation.String()): children != nil,
		join(cn, divider.Dashed):                           !!d.Dashed,
	}
	return elem.Div(
		vecty.Markup(cls),
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
