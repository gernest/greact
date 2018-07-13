package anchor

import (
	"strconv"

	"github.com/gernest/vected/lib/gs"
	"github.com/gernest/vected/web/style/core/themes"
	"github.com/gernest/vected/web/style/mixins"
)

var klass = themes.Default.AntPrefix + "-alert"
var msgColor = themes.Default.HeadingColor
var textColor = themes.Default.TextColor
var closeColor = themes.Default.TextColor

func Style() gs.CSSRule {
	a := themes.Default.FontSizeBase.Value()
	b := themes.Default.LineHeightBase.Value()
	c := (8 + a) * ((b / 2) - (a / 2) + 1)
	return gs.CSS(
		gs.S(klass,
			mixins.ResetComponent(),
			gs.P("position", "relative"),
			gs.P("padding", "8px 15px 8px 37px"),
			gs.P("border-radius", themes.Default.BorderRadiusBase.String()),
			gs.S("&&-no-icon",
				gs.P("padding", "8px 15px"),
			),
			gs.S("&&-icon",
				gs.P("top", formatFloat(c)+"px"),
				gs.P("left", "16px"),
				gs.P("position", "absolute"),
			),
			gs.S("&-description",
				gs.P("font-size", themes.Default.FontSizeBase.String()),
				gs.P("line-height", "22px"),
				gs.P("display", "none"),
			),
		),
	)
}

func formatFloat(v float64) string {
	return strconv.FormatFloat(v, 'f', 8, 64)
}
