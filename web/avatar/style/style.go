package Style

import (
	"github.com/gernest/vected/lib/gs"
	"github.com/gernest/vected/web/style/core/themes"
	"github.com/gernest/vected/web/style/mixins"
)

var avatarPrefix = themes.Default.AntPrefix + "-avatar"

func size(asize, font string) gs.CSSRule {
	return gs.CSS(
		gs.P("width", asize),
		gs.P("height", asize),
		gs.P("line-height", asize),
		gs.S("& > *",
			gs.P("line-height", asize),
		),
		gs.S(avatarPrefix+"-icon",
			gs.P("font-size", font),
		),
	)
}

// Avatar returns css for antd Avatar component.
func Avatar() gs.CSSRule {
	return gs.S(avatarPrefix,
		mixins.ResetComponent(),
		gs.P("display", "inline-block"),
		gs.P("text-align", "center"),
		gs.P("background", themes.Default.AvatarBG.Hex()),
		gs.P("color", themes.Default.AvatarColor.Hex()),
		gs.P("white-space", "nowrap"),
		gs.P("position", "relative"),
		gs.P("overflow", "hidden"),
		gs.P("vertical-align", "middle"),
		gs.S("&-image",
			gs.P("background", "transparent"),
		),
		size(themes.Default.AvatarSizeBase, themes.Default.AvatarFontSizeBase),
		gs.S("&-lg",
			size(themes.Default.AvatarSizeLG, themes.Default.AvatarFontSizeLG),
		),
		gs.S("&-sm",
			size(themes.Default.AvatarSizeSM, themes.Default.AvatarFontSizeSM),
		),
		gs.S("&-square",
			gs.P("border-radius", themes.Default.AvatarBorderRadius.String()),
		),
		gs.S("& > img",
			gs.P("width", "100%"),
			gs.P("height", "100%"),
			gs.P("display", "block"),
		),
	)
}
