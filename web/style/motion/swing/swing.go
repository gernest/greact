package swing

import (
	"github.com/gernest/vected/lib/gs"
	"github.com/gernest/vected/web/style/core/themes"
	"github.com/gernest/vected/web/style/mixins"
)

// Motion returns styles for swing motion.
func Motion(klass, keyframe string) gs.CSSRule {
	return gs.CSS(
		gs.S(klass+"-enter,",
			gs.S(klass+"-appear",
				mixins.CommonMotion(themes.Default.AnimationDurationBase),
				gs.P("animation-play-state", "paused"),
			),
		),
		gs.S(klass+"-enter"+klass+"-enter-active,",
			gs.S(""+klass+"-appear"+klass+"-appear-active",
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
		gs.Cond("0%,\n100%",
			gs.P("transform", "translateX(0)"),
		),
		gs.Cond("20%",
			gs.P("transform", "translateX(-10px)"),
		),
		gs.Cond("40%",
			gs.P("transform", "translateX(10px)"),
		),
		gs.Cond("60%",
			gs.P("transform", "translateX(-5px)"),
		),
		gs.Cond("80%",
			gs.P("transform", "translateX(5px)"),
		),
	)
}
