package themes

import (
	"github.com/gernest/vected/style/tinycolor"
)

var Default = &Theme{}

type Theme struct {
	FontFamily     string
	FontSizeBase   string
	LineHeightBase string
	TextColor      tinycolor.Color
}
