package zoom

import (
	"github.com/gernest/gs"
	"github.com/gernest/vected/web/mixins"
	"github.com/gernest/vected/web/themes"
)

func Motion(klass, keyframe, duration string) gs.CSSRule {
	return gs.CSS(
		mixins.MakeMotion(klass, keyframe, duration),
		gs.S(klass+"-enter",
			gs.S("&,"+klass+"-appear",
				gs.P("transform", "scale(0)"),
				gs.P("animation-timing-function", themes.Default.EaseOutCirc),
			),
		),
		gs.S(klass+"-leave",
			gs.P("animation-timing-function", themes.Default.EaseInOutCirc),
		),
	)
}
