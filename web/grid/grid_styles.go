package grid

import (
	"strconv"

	"github.com/gernest/vected/web/style/core/themes"

	"github.com/gernest/gs"
	"github.com/gernest/vected/web/style/mixins"
)

func makeRow(gutter int64) gs.CSSRule {
	return gs.CSS(
		gs.P("position", "relative"),
		gs.P("margin-left", formatInt(gutter/-2)),
		gs.P("margin-right", formatInt(gutter/-2)),
		gs.P("height", "auto"),
		mixins.ClearFix(),
	)
}

func formatInt(v int64) string {
	return strconv.FormatInt(v, 10)
}

var prefix = themes.Default.AntPrefix

// Styles complete styles for antd grids
func Styles() gs.CSSRule {
	return gs.CSS(
		gs.S(prefix+"-row",
			makeRow(themes.Default.GridGutterWidth),
			gs.P("display", "block"),
			gs.P("box-sizing", "border-box"),
		),
		gs.S(prefix+"-row-flex",
			gs.P("display", "flex"),
			gs.P("flex-flow", "row wrap"),
			gs.S("&:before", gs.P("display", "flex")),
			gs.S("&:after", gs.P("display", "flex")),
		),
		gs.S(prefix+"-row-flex-start", gs.P("justify-content", "flex-start")),
		gs.S(prefix+"-row-flex-center", gs.P("justify-content", "center")),
		gs.S(prefix+"-row-flex-end", gs.P("justify-content", "flex-end")),
		gs.S(prefix+"-row-flex-space-between", gs.P("justify-content", "space-between")),
		gs.S(prefix+"-row-flex-space-around", gs.P("justify-content", "space-around")),
		gs.S(prefix+"-row-flex-top", gs.P("justify-content", "flex-start")),
		gs.S(prefix+"-row-flex-middle", gs.P("justify-content", "center")),
		gs.S(prefix+"-row-flex-bottom", gs.P("justify-content", "flex-bottom")),
	)
}
