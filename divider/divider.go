package divider

import (
	"github.com/gernest/gs"
	"github.com/gopherjs/vecty"
)

type Kind int64

const (
	Horizontal Kind = iota
	Vertical
)

type Orientation int64

const (
	Left Orientation = iota
	Right
)

type Divider struct {
	vecty.Core
	Type        Kind
	Orientation Orientation
	Dashed      bool
	Style       vecty.Applyer
	CSS         gs.CSSRule
	Text        string
	sheet       *gs.Sheet
}

func style() gs.CSSRule {

}
