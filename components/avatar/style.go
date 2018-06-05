package avatar

import (
	"github.com/gernest/gs"
	"github.com/gernest/vected/themes"
)

const avatarClass = "vected-avatar"

var t = themes.Default

func style() gs.CSSRule {
	return gs.CSS(
		gs.S(avatarClass,
			gs.P("font-family", themes.Default.FontFamily),
			gs.P("font-size", themes.Default.FontSizeBase),
			gs.P("line-height", themes.Default.LineHeightBase),
			gs.P("color", themes.Default.TextColor.Hex()),
			gs.P("ox-sizing", "border-box"),
			gs.P("margin", "0"),
			gs.P("padding", "0"),
			gs.P("list-style", "none"),
			gs.P("display", "inline-block"),
			gs.P("text-align", "center"),
			gs.P("color", t.AvatarColor.String()),
			gs.P("white-space", "nowrap"),
			gs.P("vertical-align", "middle"),
			gs.S("&-image",
				gs.P("background", "transparent"),
			),
		),
	)
}
