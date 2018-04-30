package components

import (
	"github.com/gopherjs/vecty"
)

type BasicGrid struct {
	vecty.Core
}

func (BasicGrid) Render() vecty.ComponentOrHTML {
	return nil
}
