package grid

import (
	"strconv"

	"github.com/gernest/gs"
	"github.com/gernest/vected/style/mixins"
)

type Number int64

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

func RowStyle() gs.CSSRule {
	return gs.CSS(
		gs.S(".rowFlex",
			gs.P("display", "flex"),
			gs.P("flex-flow", "row wrap"),
			gs.S("&:before", gs.P("display", "flex")),
			gs.S("&:after", gs.P("display", "flex")),
		),
	)
}

func MakeColumn(index, gutter, numCols int64) gs.CSSRule {
	return gs.CSS(
		gs.S(".Col",
			gs.P("position", "relative"),
			gs.P("display", "block"),
			gs.P("min-height", "1px"),
			gs.P(" padding-left", format(gutter/2)),
			gs.P(" padding-right", format(gutter/2)),
			gs.P("float", "left"),
			gs.P("flex", "0 0 auto"),
			gs.P("box-sizing", "border-box"),
			gs.P("width", precent(float64(index)/float64(numCols))),
		),
	)
}

func Push(index int64, numCols int64) gs.CSSRule {
	return gs.S(".Push",
		gs.P("left", precent(float64(index)/float64(numCols))),
	)
}

func Pull(index int64, numCols int64) gs.CSSRule {
	return gs.S(".Pull",
		gs.P("right", precent(float64(index)/float64(numCols))),
	)
}

func Offset(index int64, numCols int64) gs.CSSRule {
	return gs.S(".Offset",
		gs.P("margin-left", precent(float64(index)/float64(numCols))),
	)
}

func Order(index int64) gs.CSSRule {
	return gs.S(".Order",
		gs.P("order", format(index)),
	)
}

func precent(v float64) string {
	return formatFloat(v) + "%"
}
