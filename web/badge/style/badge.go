package style

import (
	"github.com/gernest/vected/lib/gs"
	"github.com/gernest/vected/web/style/core/themes"
	"github.com/gernest/vected/web/style/mixins"
)

var prefix = themes.Default.AntPrefix + "-badge"
var numberPrefix = themes.Default.AntPrefix + "-scroll-number"
var height = themes.Default.BadgeHeight

// Badge returns stylesheet for antd badge.
func Badge() gs.CSSRule {
	return gs.CSS(
		gs.S(prefix,
			mixins.ResetComponent(),
			gs.P("position", "relative"),
			gs.P("display", "inline-block"),
			gs.P("line-height", "1"),
			gs.P("vertical-align", "middle"),
			gs.S("&-count ",
				gs.P("position", "absolute"),
				gs.P("transform", "translateX(-50%)"),
				gs.P("top", height.Div(gs.U("2")).String()),
				gs.P("height", height.String()),
				gs.P("border-radius", height.Div(gs.U("2")).String()),
				gs.P("min-width", height.String()),
				gs.P("background", themes.Default.HighlightColor.Hex()),
				gs.P("color", "#fff"),
				gs.P("line-height", height.String()),
				gs.P("text-align", "center"),
				gs.P("padding", "0 6px"),
				gs.P("font-size", themes.Default.BadgeFontSize.String()),
				gs.P("font-weight", themes.Default.BadgeFontWeight),
				gs.P("white-space", "nowrap"),
				gs.P("transform-origin", "-10% center"),
				gs.P("box-shadow", "0 0 0 1px #fff"),
			),
		),
	)
}
