package swing

import (
	"github.com/gernest/gs"
	"github.com/gernest/vected/web/mixins"
	"github.com/gernest/vected/web/themes"
)

func Motion(klass, keyframe string) gs.CSSRule {
	return gs.CSS(
		gs.S(klass+"-enter",
			gs.S("&,"+klass+"-appear",
				mixins.CommonMotion(themes.Default.AnimationDurationBase),
				gs.P("animation-play-state", "paused"),
			),
		),
		gs.S(klass+"-enter"+klass+"-enter-active",
			gs.S("&,"+klass+"-appear"+klass+"-appear-active",
				gs.P("animation-name", "~"+keyframe+"In"),
				gs.P("animation-play-state:", "running"),
			),
		),
	)
}

//key frame name
const (
	Swing = "swing"
)

func KeyFrames() gs.CSSRule {
	return gs.KeyFrame(Swing+"In",
		gs.S("0%",
			gs.S("&,100%",
				gs.P("transform", "translateX(0)"),
			),
		),
		gs.S("20%",
			gs.P("transform", "translateX(-10px)"),
		),
		gs.S("40%",
			gs.P("transform", "translateX(10px)"),
		),
		gs.S("60%",
			gs.P("transform", "translateX(-5px)"),
		),
		gs.S("80%",
			gs.P("transform", "translateX(5px)"),
		),
	)
}
