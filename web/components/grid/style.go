package grid

import (
	"strconv"

	"github.com/gernest/gs"
	"github.com/gernest/vected/web/mixins"
	"github.com/gernest/vected/web/themes"
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

const (
	RowClass     = ".Row"
	RowFLexClass = ".RowFlex"
	ColClass     = ".Col"
	PushClass    = ".Push"
	PullClass    = ".Pull"
	OffsetClass  = ".Offset"
	OrderClass   = ".Order"
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

// RowStyle returns and styles for a grid row.
func RowStyle(gutter int64, flex bool, justify FlexStyle, align FlexAlign) gs.CSSRule {
	if flex {
		var s gs.CSSRule
		switch justify {
		case End:
			s = gs.P("justify-content", "flex-end")
		case Center:
			s = gs.P("justify-content", "center")
		case SpaceAround:
			s = gs.P("justify-content", "space-around")
		case SpaceBetween:
			s = gs.P("justify-content", "space-between")
		case Start:
			s = gs.P("justify-content", "flex-start")
		default:
			s = gs.P("justify-content", "flex-start")
		}
		var a gs.CSSRule
		switch align {
		case Top:
			a = gs.P("align-items", "flex-start")
		case Middle:
			a = gs.P("align-items", "center")
		case Bottom:
			a = gs.P("align-items", "flex-end")
		default:
			a = gs.P("align-items", "flex-start")
		}
		return gs.CSS(
			gs.S(RowFLexClass,
				gs.P("display", "flex"),
				gs.P("flex-flow", "row wrap"),
				gs.S("&:before", gs.P("display", "flex")),
				gs.S("&:after", gs.P("display", "flex")),
				gs.S("&:after", gs.P("display", "flex")),
				s, a,
			),
		)
	}
	return gs.CSS(
		gs.S(RowClass,
			gs.P("position", "relative"),
			gs.P("margin-left", format(gutter/-2)+"px"),
			gs.P("margin-right", format(gutter/-2)+"px"),
			gs.P("box-sizing", "border-box"),
			gs.P("display", "block"),
			gs.P("height", "auto"),
			gs.P("zoom", "1"),
			gs.S("&:before",
				gs.P("content", ""),
				gs.P("display", "table"),
			),
			gs.S("&:after",
				gs.P("content", ""),
				gs.P("display", "table"),
				gs.P("clear", "both"),
				gs.P("visibility", "hidden"),
				gs.P("font-size", "0"),
				gs.P("height", "0"),
			)),
	)
}

func clearFix() gs.CSSRule {
	return gs.CSS(
		gs.P("zoom", "1"),
		gs.S("&:before",
			gs.P("content", ""),
			gs.P("display", "table"),
		),
		gs.S("&:after",
			gs.P("content", ""),
			gs.P("display", "table"),
			gs.P("clear", "both"),
			gs.P("visibility", "hidden"),
			gs.P("font-size", "0"),
			gs.P("height", "0"),
		),
	)

}

func MakeColumn(index, gutter, numCols int64) gs.CSSRule {
	display := gs.P("display", "block")
	if index == 0 {
		display = gs.P("display", "none")
	}
	return gs.CSS(
		gs.S(ColClass,
			gs.P("position", "relative"),
			display,
			gs.P("min-height", "1px"),
			gs.P("box-sizing", "border-box"),
			gs.P("float", "left"),
			gs.P("flex", "0 0 auto"),
			gs.P("width", precent(float64(index)/float64(numCols))),
		),
	)
}

func Push(index int64, numCols int64) gs.CSSRule {
	if index == 0 {
		return gs.S(PushClass,
			gs.P("left", "auto"),
		)
	}
	return gs.S(PushClass,
		gs.P("left", precent(float64(index)/float64(numCols))),
	)
}

func Pull(index int64, numCols int64) gs.CSSRule {
	if index == 0 {
		return gs.S(PushClass,
			gs.P("right", "auto"),
		)
	}
	return gs.S(PullClass,
		gs.P("right", precent(float64(index)/float64(numCols))),
	)
}

func Offset(index int64, numCols int64) gs.CSSRule {
	if index == 0 {
		return gs.S(PushClass,
			gs.P("margin-left", "0"),
		)
	}
	return gs.S(OffsetClass,
		gs.P("margin-left", precent(float64(index)/float64(numCols))),
	)
}

func Order(index int64) gs.CSSRule {
	if index == 0 {
		return gs.S(PushClass,
			gs.P("order", "0"),
		)
	}
	return gs.S(OrderClass,
		gs.P("order", format(index)),
	)
}

func precent(v float64) string {
	return formatFloat(v*100) + "%"
}

func ColumnStyle(opts *ColOptions, mediaQuery ...MediaOption) gs.CSSRule {
	var rules gs.RuleList
	index := int64(opts.Span)
	cols := themes.Default.GridColumns
	rules = append(rules, MakeColumn(index, opts.Gutter, cols))
	if opts.Pull != 0 {
		rules = append(rules, Pull(int64(opts.Pull), cols))
	}
	if opts.Push != 0 {
		rules = append(rules, Push(int64(opts.Push), cols))
	}
	if opts.Offset != 0 {
		rules = append(rules, Offset(int64(opts.Offset), cols))
	}
	if opts.Order != 0 {
		rules = append(rules, Order(int64(opts.Order)))
	}
	for _, v := range mediaQuery {
		rules = append(rules, Media(v.Type, v.Opts))
	}
	return rules
}

type MediaOption struct {
	Type MediaType
	Opts *ColOptions
}

func Media(cond MediaType, opts ...*ColOptions) gs.CSSRule {
	var rules gs.RuleList
	for _, v := range opts {
		rules = append(rules, ColumnStyle(v))
	}
	return gs.Cond(cond.Screen(), rules...)
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
