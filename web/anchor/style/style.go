package anchor

import (
	"github.com/gernest/vected/lib/gs"
	"github.com/gernest/vected/web/style/core/themes"
	"github.com/gernest/vected/web/style/mixins"
)

var prefix = themes.Default.AntPrefix

const borderWidth = "2px"

// Anchor returns css style for antd anchor.
func Anchor() gs.CSSRule {
	return gs.CSS(
		gs.S(prefix+"-anchor",
			mixins.ResetComponent(),
			gs.P("position", "relative"),
			gs.P(" padding-left", borderWidth),
			gs.S("&-wrapper",
				gs.P("background-color", themes.Default.ComponentBackground.Hex()),
				gs.P("overflow", "auto"),
				gs.P("padding-left", "4px"),
				gs.P("margin-left", "-4px"),
			),
			gs.S("&-ink",
				gs.P("position", "absolute"),
				gs.P("height", "100%"),
				gs.P("left", "0"),
				gs.P("top", "0"),
				gs.S("&:before",
					gs.P("content", "' "),
					gs.P("position", "relative"),
					gs.P("width", borderWidth),
					gs.P("height", "100%"),
					gs.P("display", "block"),
					gs.P("background-color", themes.Default.BorderColorSplit.Hex()),
					gs.P("margin", "auto"),
				),
			),
			gs.S("&-ball",
				gs.P("display", "none"),
				gs.P("position", "absolute"),
				gs.P("height", "8px"),
				gs.P("border-radius", "8px"),
				gs.P("border", "2px solid "+themes.Default.PrimaryColor.Hex()),
				gs.P("background-color", themes.Default.ComponentBackground.Hex()),
				gs.P("left", "50%"),
				gs.P("transition", "top .3s ease-in-out"),
				gs.P("transform", "translateX(-50%)"),
				gs.S("&.visible",
					gs.P("display", "inline-block"),
				),
			),
			gs.S("&.fixed &-ink &-ink-ball",
				gs.P("display", "none"),
			),
			gs.S("&-link",
				gs.P("padding", "8px 0 8px 16px"),
				gs.P("line-height", "1"),
				gs.S("&-title",
					gs.P("display", "block"),
					gs.P("position", "relative"),
					gs.P("transition", "all .3s"),
					gs.P("color", themes.Default.TextColor.Hex()),
					gs.P("white-space", "nowrap"),
					gs.P("overflow", "hidden"),
					gs.P("text-overflow", "ellipsis"),
					gs.P("margin-bottom", "8px"),
					gs.S("&:only-child",
						gs.P("margin-bottom", "0"),
					),
				),
				gs.S("&-active > &-title",
					gs.P("color", themes.Default.PrimaryColor.Hex()),
				),
			),
			gs.S("&-link &-link",
				gs.P("padding-top", "6px"),
				gs.P("padding-bottom", "6px"),
			),
		),
	)
}
