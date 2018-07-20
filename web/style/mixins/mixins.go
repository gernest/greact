package mixins

import (
	"github.com/gernest/vected/lib/gs"
	"github.com/gernest/vected/web/style/core/themes"
)

const (
	appear = "-appear"
	enter  = "-enter"
	leave  = "-leave"
)

// CommonMotion  ==> .motion-common
func CommonMotion(duration string) gs.CSSRule {
	return gs.CSS(
		gs.P("animation-duration", duration),
		gs.P("animation-fill-mode:", "both"),
	)
}

func MakeMotion(klass, keyframe string, duration string) gs.CSSRule {
	return gs.CSS(
		gs.S(join(klass, enter),
			gs.S(join("&,", klass, appear),
				CommonMotion(duration),
				gs.P("animation-play-state", "paused"),
			)),
		gs.S(join(klass, leave),
			CommonMotion(duration),
			gs.P("animation-play-state", "paused"),
		),
		gs.S(join(
			join(klass, enter), join(klass, enter, "-active"),
		),
			gs.S(
				join(join("&,", klass, enter), join(klass, appear, "-active")),
				gs.P("animation-name", join("~", keyframe, "In")),
				gs.P("animation-play-state", "running"),
			),
		),
		gs.S(join(
			join(klass, leave), join(klass, leave, "active"),
		),
			gs.P("animation-name", join("~", keyframe, "Out")),
			gs.P("animation-play-state", "running"),
			gs.P("pointer-events", "none"),
		),
	)
}

// join joins v and  returns the result with prefix .

func join(s ...string) string {
	o := ""
	for _, v := range s {
		o += v
	}
	return o
}

func ClearFix() gs.CSSRule {
	return gs.CSS(
		gs.P("zoom", "1"),
		gs.S("&:before",
			gs.P("content", `""`),
			gs.P("display", "table"),
		),
		gs.S("&:after",
			gs.P("content", `""`),
			gs.P("display", "table"),
			gs.P("clear", "both"),
			gs.P("visibility", "hidden"),
			gs.P("font-size", "0"),
			gs.P("height", "0"),
		),
	)
}

func ResetComponent() gs.CSSRule {
	return gs.CSS(
		gs.P("font-family", themes.Default.FontFamily),
		gs.P("font-size", themes.Default.FontSizeBase.String()),
		gs.P("line-height", themes.Default.LineHeightBase.String()),
		gs.P("color", themes.Default.TextColor.String()),
		gs.P("box-sizing", "border-box"),
		gs.P("margin", "0"),
		gs.P("padding", "0"),
		gs.P("list-style", "none"),
	)
}

//IconFontMixin .iconfont-mixin
func IconFontMixin() gs.CSSRule {
	return gs.CSS(
		gs.P("display", "inline-block"),
		gs.P("font-style", "normal"),
		gs.P("vertical-align", "baseline"),
		gs.P("text-align", "center"),
		gs.P("text-transform:", "none"),
		gs.P("line-height", "1"),
		gs.P("text-rendering", "optimizeLegibility"),
		gs.P("-webkit-font-smoothing", "antialiased"),
		gs.P("-moz-osx-font-smoothing", "grayscale"),
		gs.S("&:before",
			gs.P("display", "block"),
			gs.P("font-family", `"anticon" !important`),
		),
	)
}

// IconFontFont => .iconfont-font
func IconFontFont(content string) gs.CSSRule {
	return gs.CSS(
		gs.P("font-family", "'anticon'"),
		gs.P("text-rendering", "optimizeLegibility"),
		gs.P("-webkit-font-smoothing", "antialiased"),
		gs.P("-moz-osx-font-smoothing", "grayscale"),
		gs.P("content", content),
	)
}
