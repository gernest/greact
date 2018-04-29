package grid

import (
	"strconv"

	"github.com/gernest/vected/style/themes"

	"github.com/gernest/gs"
	"github.com/gernest/vected/style/mixins"
)

type Number int64

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
	RowClass    = ".Row"
	ColClass    = ".Col"
	PushClass   = ".Push"
	PullClass   = ".Pull"
	OffsetClass = ".Offset"
	OrderClass  = ".Order"
)

func (m MediaType) Screen() string {
	return "@media (min-width:" + string(m) + ")"
}

type ColOptions struct {
	Span, Order, Offset, Push, Pull Number
}

func MakeRow(gutter int64) gs.CSSRule {
	return gs.CSS(
		gs.P("position", "relative"),
		gs.P("margin-left", format(gutter/-2)),
		gs.P("margin-right", format(gutter/-2)),
		gs.P("height", "auto"),
		mixins.ClearFix(),
	)
}

func format(v int64) string {
	return strconv.FormatInt(v, 10)
}
func formatFloat(v float64) string {
	return strconv.FormatFloat(v, 'f', -1, 64)
}

func Row() gs.CSSRule {
	return gs.CSS(
		gs.S(RowClass,
			gs.P("display", "flex"),
			gs.P("flex-flow", "row wrap"),
			gs.S("&:before", gs.P("display", "flex")),
			gs.S("&:after", gs.P("display", "flex")),
		),
	)
}

func MakeColumn(index, gutter, numCols int64) gs.CSSRule {
	return gs.CSS(
		gs.S(ColClass,
			gs.P("position", "relative"),
			gs.P("display", "block"),
			gs.P("background", "blue"),
			gs.P("min-height", "1px"),
			gs.P("padding-left", format(gutter/2)),
			gs.P("padding-right", format(gutter/2)),
			gs.P("float", "left"),
			gs.P("flex", "0 0 auto"),
			gs.P("box-sizing", "border-box"),
			gs.P("width", precent(float64(index)/float64(numCols))),
		),
	)
}

func Push(index int64, numCols int64) gs.CSSRule {
	return gs.S(PushClass,
		gs.P("left", precent(float64(index)/float64(numCols))),
	)
}

func Pull(index int64, numCols int64) gs.CSSRule {
	return gs.S(PullClass,
		gs.P("right", precent(float64(index)/float64(numCols))),
	)
}

func Offset(index int64, numCols int64) gs.CSSRule {
	return gs.S(OffsetClass,
		gs.P("margin-left", precent(float64(index)/float64(numCols))),
	)
}

func Order(index int64) gs.CSSRule {
	return gs.S(OrderClass,
		gs.P("order", format(index)),
	)
}

func precent(v float64) string {
	return formatFloat(v*100) + "%"
}

func Column(opts *ColOptions, mediaQuery ...MediaOption) gs.CSSRule {
	var rules gs.RuleList
	index := int64(opts.Span)
	cols := themes.Default.GridColumns
	rules = append(rules, MakeColumn(index, themes.Default.GridGutterWidth,
		cols))
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
		rules = append(rules, Column(v))
	}
	return gs.Cond(cond.Screen(), rules...)
}
