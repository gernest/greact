package other

import (
	"github.com/gernest/vected/lib/gs"
)

func LoadingCircle() gs.CSSRule {
	return gs.KeyFrame("loadingCircle",
		gs.Cond("0%",
			gs.P("transform-origin", "50% 50%"),
			gs.P("transform", "rotate(0deg)"),
		),
		gs.Cond("100%",
			gs.P("transform-origin", "50% 50%"),
			gs.P("transform", "rotate(360deg)"),
		),
	)
}
