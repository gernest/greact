package divider

import (
	"github.com/gernest/gs"
	"github.com/gernest/vected/style/themes"
)

const BaseClass = ".Divider"

func Style() gs.CSSRule {
	return gs.S(BaseClass,
		//reset component
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
