package grid

import (
	"strconv"

	"github.com/gernest/gs"
	"github.com/gernest/vected/web/mixins"
	"github.com/gernest/vected/web/style/core/themes"
)

func (n Number) String() string {
	b := "Col-"
	switch int64(n) {
	case 1:
		return b + "1"
	case 2:
		return b + "2"
	case 3:
		return b + "3"
	case 4:
		return b + "4"
	case 5:
		return b + "5"
	case 6:
		return b + "6"
	case 7:
		return b + "7"
	case 8:
		return b + "8"
	case 9:
		return b + "9"
	case 10:
		return b + "10"
	case 11:
		return b + "11"
	case 12:
		return b + "12"
	case 13:
		return b + "13"
	case 14:
		return b + "14"
	case 15:
		return b + "15"
	case 16:
		return b + "16"
	case 17:
		return b + "17"
	case 18:
		return b + "18"
	case 19:
		return b + "19"
	case 20:
		return b + "20"
	case 21:
		return b + "21"
	case 22:
		return b + "22"
	case 23:
		return b + "23"
	case 24:
		return b + "24"
	default:
		return ""
	}
}

type MediaType string

var (
	XS  = MediaType(themes.Default.ScreenXS)
	SM  = MediaType(themes.Default.ScreenSM)
	MD  = MediaType(themes.Default.ScreenMD)
	LG  = MediaType(themes.Default.ScreenLG)
	XL  = MediaType(themes.Default.ScreenXL)
	XXL = MediaType(themes.Default.ScreenXXL)
)

// main class keys for vected grid.
const (
	RowClass     = ".vected-row"
	RowFLexClass = ".vected-row-flex"
	ColClass     = ".vected-col"
	PushClass    = ".vected-col-push"
	PullClass    = ".vected-col-pull"
	OffsetClass  = ".vected-col-offset"
	OrderClass   = ".vected-col-order"
)

func (m MediaType) Screen() string {
	return "@media (min-width:" + string(m) + ")"
}

type ColOptions struct {
	Span, Order, Offset, Push, Pull Number
	Gutter                          int64
}

func formatFloat(v float64) string {
	return strconv.FormatFloat(v, 'f', -1, 64)
}

func precent(v float64) string {
	return formatFloat(v*100) + "%"
}

type MediaOption struct {
	Type MediaType
	Opts *ColOptions
}

// New Direction

func MakeRow(gutter int64) gs.CSSRule {
	if gutter == 0 {
		gutter = themes.Default.GridGutterWidth
	}
	return gs.CSS(
		gs.P("position", "relative"),
		gs.P("margin-left", format(gutter/-2)+"px"),
		gs.P("margin-right", format(gutter/-2)+"px"),
		gs.P("height", "auto"),
		mixins.ClearFix(),
	)
}

func RowStyle(gutter int64) gs.CSSRule {
	return gs.S(RowClass,
		MakeRow(gutter),
		gs.P("display", "block"),
		gs.P("box-sizing", "border-box"),
	)
}

func format(v int64) string {
	return strconv.FormatInt(v, 10)
}
