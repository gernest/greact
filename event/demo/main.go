package main

import (
	"github.com/gernest/naaz/prop"
	"github.com/gernest/vected/dom"
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
	"github.com/gopherjs/vecty/event"
)

func main() {
	vecty.RenderBody(&Drag{})
}

type Drag struct {
	vecty.Core
}

func (Drag) Render() vecty.ComponentOrHTML {
	return elem.Body(
		elem.UnorderedList(
			elem.ListItem(
				vecty.Markup(
					prop.ID("i1"),
					vecty.Attribute("draggable", "true"),
					event.DragStart(func(e *vecty.Event) {
						d := dom.ToDragEvent(e.Object)
						id := e.Target.Get("id").String()
						d.DataTransfer.SetData("text/plain", id)
						d.DataTransfer.EffectAllowed = "move"
					}),
				),
				vecty.Text("Drag Item 1 to the Drop Zone"),
			),
			elem.ListItem(
				vecty.Markup(
					prop.ID("i2"),
					vecty.Attribute("draggable", "true"),
					event.DragStart(func(e *vecty.Event) {
						d := dom.ToDragEvent(e.Object)
						id := e.Target.Get("id").String()
						d.DataTransfer.SetData("text/plain", id)
						d.DataTransfer.EffectAllowed = "move"
					}),
				),
				vecty.Text("Drag Item 2 to the Drop Zone"),
			),
		),
		elem.Div(
			vecty.Markup(
				prop.ID("target"),
				event.Drop(func(e *vecty.Event) {
					dt := dom.ToDragEvent(e.Object)
					println(dt.DataTransfer)
				}),
			),
			vecty.Text("Drop Zone"),
		),
	)
}
