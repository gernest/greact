package grid

import (
	"github.com/gernest/vected/style/grid"
	"github.com/gopherjs/vecty"
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

type Row struct {
	vecty.Core
}
