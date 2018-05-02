package main

import (
	"github.com/gernest/vected/divider"
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
)

func main() {
	vecty.RenderBody(&demo{})
}

type demo struct {
	vecty.Core
}

func (demo) Render() vecty.ComponentOrHTML {
	return elem.Body(
		loren(),
		&divider.Divider{},
		loren(),
		&divider.Divider{
			Children: func() vecty.MarkupOrChild {
				return vecty.Text("With text")
			},
		},
		loren(),
		&divider.Divider{
			Dashed: true,
		},
		loren(),
		&divider.Divider{},
	)
}

func loren() *vecty.HTML {
	return elem.Paragraph(
		vecty.Text("Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed nonne merninisti licere mihi ista probare, quae sunt a te dicta? Refert tamen, quo modo"),
	)
}
