package grid

import (
	"github.com/gernest/gs"
	"github.com/gernest/vected/style/grid"
	"github.com/gernest/vected/ui"
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
)

type Number = grid.Number

const (
	Empty Number = iota
	G1
	G2
	G3
	G4
	G5
	G6
	G7
	G8
	G9
	G10
	G11
	G12
	G13
	G14
	G15
	G16
	G17
	G18
	G19
	G20
	G21
	G22
	G23
	G24
)

type ColOptions = grid.ColOptions

type Column struct {
	vecty.Core
	ColOptions

	Style                   vecty.Applyer
	Children                vecty.MarkupOrChild
	XS, SM, MD, LG, XL, XXL *ColOptions

	sheet *gs.Sheet
}

func (c *Column) Mount() {
	c.sheet.Attach()
}

func (c *Column) Render() vecty.ComponentOrHTML {
	if c.sheet == nil {
		c.sheet = ui.NewSheet()
		c.sheet.AddRule(c.style())
	}
	classes := vecty.ClassMap(c.sheet.CLasses.Classes())
	return elem.Div(vecty.Markup(classes, c.Style), c.Children)
}

func (c *Column) style() gs.CSSRule {
	var media []grid.MediaOption
	if c.XS != nil {
		media = append(media, grid.MediaOption{
			Type: grid.XS,
			Opts: c.XS,
		})
	}
	if c.SM != nil {
		media = append(media, grid.MediaOption{
			Type: grid.SM,
			Opts: c.SM,
		})
	}
	if c.MD != nil {
		media = append(media, grid.MediaOption{
			Type: grid.MD,
			Opts: c.MD,
		})
	}
	if c.LG != nil {
		media = append(media, grid.MediaOption{
			Type: grid.LG,
			Opts: c.LG,
		})
	}

	if c.XL != nil {
		media = append(media, grid.MediaOption{
			Type: grid.XL,
			Opts: c.XL,
		})
	}
	if c.XXL != nil {
		media = append(media, grid.MediaOption{
			Type: grid.XXL,
			Opts: c.XXL,
		})
	}
	return grid.Column(&c.ColOptions, media...)
}

func join(s ...string) string {
	o := ""
	for _, v := range s {
		o += v
	}
	return o
}

func (c *Column) Unmount() {
	c.sheet.Detach()
}
