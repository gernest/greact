package anchor

import (
	"github.com/gernest/vected/lib/gs"
	"github.com/gernest/vected/web/style/core/themes"
	"github.com/gernest/vected/web/style/mixins"
)

var klass = themes.Default.AntPrefix
var msgColor = themes.Default.HeadingColor
var textColor = themes.Default.TextColor
var closeColor = themes.Default.TextColor

func Style() gs.CSSRule {
	return gs.CSS(
		gs.S(klass,
			mixins.ResetComponent(),
			gs.P("position", "relative"),
			gs.P("padding", "8px 15px 8px 37px"),
			gs.P("border-radius", themes.Default.BorderRadiusBase),
			gs.S("&&-no-icon",
				gs.P("padding", "8px 15px"),
			),
		),
	)
}
