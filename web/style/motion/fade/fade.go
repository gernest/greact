package fade

import (
	"github.com/gernest/gs"
	"github.com/gernest/vected/web/mixins"
	"github.com/gernest/vected/web/themes"
)

func Fade(klass, keyframe string) gs.CSSRule {
	return gs.CSS(
		mixins.MakeMotion(klass, keyframe, themes.Default.AnimationDurationBase),
		gs.S(klass+"-enter",
			gs.S("&,"+klass+"-appear",
				gs.P("opacity", "0"),
				gs.P("animation-timing-function", "linear"),
			),
		),
		gs.S(klass+"-leave",
			gs.P(" animation-timing-function", "linear"),
		),
	)
}

const Klass = ".fade"
const frameName = "antFade"

func Motion() gs.CSSRule {
	return Fade(Klass, frameName)
}

func KeyFrame() gs.CSSRule {
	return gs.CSS(
		gs.KeyFrame(frameName+"In",
			gs.Cond("0%",
				gs.P("opacity", "0"),
			),
			gs.Cond("100%",
				gs.P("opacity", "1"),
			),
		),
		gs.KeyFrame(frameName+"Out",
			gs.Cond("0%",
				gs.P("opacity", "1"),
			),
			gs.Cond("100%",
				gs.P("opacity", "0"),
			),
		),
	)
}
