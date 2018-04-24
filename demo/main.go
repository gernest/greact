package main

import (
	"github.com/gernest/naaz"
	"github.com/gernest/naaz/elem"
)

func main() {
	naaz.RenderBody(&Hello{})
}

type Hello struct {
	naaz.Core
}

func (Hello) Render() naaz.ComponentOrHTML {
	return elem.Body(
		naaz.Text("hello, world"),
	)
}
