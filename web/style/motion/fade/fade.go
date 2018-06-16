package fade

import (
	"github.com/gernest/gs"
	"github.com/gernest/vected/web/mixins"
	"github.com/gernest/vected/web/themes"
)

func Fade(klass, keyframe string) gs.CSSRule {
	return gs.CSS(
		mixins.MakeMotion(klass, keyframe, themes.Default.AnimationDurationBase),
	)
}

const Klass = "fade"

func Motion() gs.CSSRule {
	return Fade(Klass, "antFade")
}
