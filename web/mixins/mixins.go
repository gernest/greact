package mixins

import (
	"github.com/gernest/gs"
	"github.com/gernest/vected/web/themes"
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

func MakeMotion(class, keyframe string, duration string) gs.CSSRule {
	return gs.CSS(
		gs.S(ToClass(class, enter),
			gs.S("&,"+ToClass(class, appear),
				CommonMotion(duration),
				gs.P("animation-play-state", "paused"),
			)),
		gs.S(ToClass(class, leave),
			CommonMotion(duration),
			gs.P("animation-play-state", "paused"),
		),
		gs.S(join(
			ToClass(class, enter), ToClass(class, enter, "-active"),
		),
			gs.S(
				join("&,", ToClass(class, enter), ToClass(class, appear, "-active")),
				gs.P("animation-name", join("~", keyframe, "In")),
				gs.P("animation-play-state", "running"),
			),
		),
		gs.S(join(
			ToClass(class, leave), ToClass(class, leave, "active"),
		),
			gs.P("animation-name", join("~", keyframe, "Out")),
			gs.P("animation-play-state", "running"),
			gs.P("pointer-events", "none"),
		),
	)
}

// ToClass joins v and  returns the result with prefix .
func ToClass(s ...string) string {
	v := []string{"."}
	v = append(v, s...)
	return join(v...)
}

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
			gs.P("content", " "),
			gs.P("display", "table"),
		),
		gs.S("&:after",
			gs.P("content", " "),
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
		gs.P("font-size", themes.Default.FontSizeBase),
		gs.P("line-height", themes.Default.LineHeightBase),
		gs.P("color", themes.Default.TextColor.Hex()),
		gs.P("ox-sizing", "border-box"),
		gs.P("margin", "0"),
		gs.P("padding", "0"),
		gs.P("list-style", "none"),
	)
}

func Size(height, width string) gs.CSSRule {
	return gs.CSS(
		gs.P("width", width),
		gs.P("height", height),
	)
}

func Square(size string) gs.CSSRule {
	return Size(size, size)
}

// .iconfont-mixin

func IconFontFont(content string) gs.CSSRule {
	return gs.CSS(
		gs.P("font-family", "'anticon'"),
		gs.P("text-rendering", "optimizeLegibility"),
		gs.P("-webkit-font-smoothing", "antialiased"),
		gs.P("-moz-osx-font-smoothing", "grayscale"),
		gs.P("content", content),
	)
}
