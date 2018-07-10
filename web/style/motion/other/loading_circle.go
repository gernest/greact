package other

import (
	"github.com/gernest/gs"
)

func LoadingCircle() gs.CSSRule {
	return gs.KeyFrame("loadingCircle",
		gs.S("0%",
			gs.P("transform-origin", "50% 50%"),
			gs.P("transform", "rotate(0deg)"),
		),
		gs.S("100%",
			gs.P("transform-origin", "50% 50%"),
			gs.P("transform", "rotate(360deg)"),
		),
	)
}
