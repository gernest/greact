package fade

import (
	"github.com/gernest/gs"
	"github.com/gernest/vected/web/mixins"
	"github.com/gernest/vected/web/style/core/themes"
)

const (
	FadeIn    = "fadeIn"
	FadeOut   = "fadeOut"
	FrameName = "fade"
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

func KeyFrame() gs.CSSRule {
	return gs.CSS(
		gs.KeyFrame(FadeIn,
			gs.S("0%",
				gs.P("opacity", "0"),
			),
			gs.S("100%",
				gs.P("opacity", "1"),
			),
		),
		gs.KeyFrame(FadeOut,
			gs.S("0%",
				gs.P("opacity", "1"),
			),
			gs.S("100%",
				gs.P("opacity", "0"),
			),
		),
	)
}
