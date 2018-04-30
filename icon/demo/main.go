package main

import (
	"github.com/gernest/gs"
	"github.com/gernest/vected/icon"
	"github.com/gernest/vected/ui"
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
)

func main() {
	vecty.RenderBody(&app{})
}

type app struct {
	vecty.Core

	sheet *gs.Sheet
}

func style() gs.CSSRule {
	return gs.CSS(
		gs.S(".IconList",
			gs.P("margin", "40px 0"),
			gs.P("list-style", "none"),
			gs.P("overflow", "hidden"),
			gs.S("& li",
				gs.P("float", "left"),
				gs.P("width", "16.66%"),
				gs.P("text-align", "center"),
				gs.P(" list-style", "none"),
				gs.P(" cursor", "pointer"),
				gs.P(" height", "100px"),
				gs.P(" transition", " color .3s ease-in-out, background-color .3s ease-in-out"),
				gs.P("position", "relative"),
				gs.P("border-radius", "4px"),
				gs.P("background-color", "#fff"),
				gs.P("overflow", "hidden"),
				gs.P("padding", "10px 0 0"),
				gs.S("&:hover",
					gs.P("background-color", "rgba(0,160,233,.7)"),
					gs.P("color", "#fff"),
				),
			),
		),
	)
}

func (app) Render() vecty.ComponentOrHTML {
	return elem.Body(&IconDemo{})
}

type IconDemo struct {
	vecty.Core
	sheet *gs.Sheet
}

func (c *IconDemo) Render() vecty.ComponentOrHTML {
	if c.sheet == nil {
		c.sheet = ui.NewSheet()
		c.sheet.AddRule(style())
	}
	a := icon.All()
	var list vecty.List
	for _, kind := range a {
		list = append(list, elem.ListItem(
			&icon.Icon{
				Kind: kind,
				CSS: gs.S(".DAnticon",
					gs.P("font-size", "24px"),
					gs.P("margin", "12px 0 16px;"),
					gs.P("transition", "transform .3s ease-in-out"),
					gs.P("will-change", "will-change"),
				),
			},
			&iconbadge{text: string(kind)},
		))
	}
	class := vecty.ClassMap(c.sheet.CLasses.Classes())
	return elem.Div(
		elem.UnorderedList(
			vecty.Markup(
				class,
			),
			list,
		),
	)
}

func (c *IconDemo) Mount() {
	c.sheet.Attach()
}
func (c *IconDemo) Unmount() {
	c.sheet.Detach()
}

type iconbadge struct {
	vecty.Core

	text  string
	sheet *gs.Sheet
}

func (c *iconbadge) Render() vecty.ComponentOrHTML {
	if c.sheet == nil {
		c.sheet = ui.NewSheet()
		c.sheet.AddRule(gs.S(".Badge",
			gs.P("display", "block"),
			gs.P("text-align", "center"),
			gs.P("transform", "scale(0.83)"),
			gs.P("ont-family", `"Lucida Console", Consolas`),
			gs.P(" white-space", "nowrap"),
		))
	}
	class := vecty.ClassMap(c.sheet.CLasses.Classes())
	return elem.Span(
		vecty.Markup(class),
		vecty.Text(c.text),
	)
}

func (c *iconbadge) Mount() {
	c.sheet.Attach()
}

func (c *iconbadge) Unmount() {
	c.sheet.Detach()
}
