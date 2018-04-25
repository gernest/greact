package icon

import (
	"github.com/gernest/gs"
	"github.com/gernest/vected/style/themes"
)

func Style(className, content string, spin bool) gs.CSSRule {
	var before gs.CSSRule
	if spin {
		before = gs.S("&:before",
			gs.P("display", "inline-block;"),
			gs.P("font-family", `"anticon" !important`),
			gs.P("-webkit-animation", "loadingCircle 1s infinite linear"),
			gs.P("animation", "loadingCircle 1s infinite linear"),
		)
	} else {
		before = gs.S("&:before",
			gs.P("display", "block"),
			gs.P("font-family", `"anticon" !important`),
			gs.P("content", content),
		)
	}
	return gs.CSS(
		gs.S(className,
			gs.P("display", "inline-block"),
			gs.P("font-style", "normal"),
			gs.P("vertical-align", "baseline"),
			gs.P("text-align", "center"),
			gs.P("text-transform", "none"),
			gs.P("line-height", "1"),
			gs.P("text-rendering", "optimizeLegibility"),
			gs.P("-webkit-font-smoothing", "antialiased"),
			gs.P("-moz-osx-font-smoothing", "grayscale"),
			before,
		),
	)
}

// FontFace is to be applied at global level.
func FontFace() gs.CSSRule {
	u := src(".woff") + `format('woff'),` +
		src(".ttf") + `format('truetype'),` +
		src(".svg#iconfont'") + `format('svg')`
	return gs.FontFace(
		gs.P("font-family", "'anticon'"),
		gs.P("src", src(".eot")),
		gs.P("src", u),
	)
}

func src(ext string) string {
	return "url('" + themes.Default.IconURL + "')" + ext
}
