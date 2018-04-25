package grid

import (
	"github.com/gopherjs/vecty"
)

type Number uint

const (
	Empty Number = iota
)

func (Number) String() string {
	return ""
}

type ColOptions struct {
	Span, Order, Offset, Push, Pull Number
}

type Column struct {
	vecty.Core
	ColOptions

	Style                   vecty.Applyer
	Children                vecty.MarkupOrChild
	XS, SM, MD, LG, XL, XXL *ColOptions
}

func (c *Column) Render() vecty.ComponentOrHTML {
	return nil
}

func join(s ...string) string {
	o := ""
	for _, v := range s {
		o += v
	}
	return o
}
