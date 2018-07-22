package style

import (
	"fmt"
	"strconv"

	"github.com/gernest/vected/lib/gs"
	"github.com/gernest/vected/web/style/color"
	"github.com/gernest/vected/web/style/core/themes"
	"github.com/gernest/vected/web/style/mixins"
)

var klass = themes.Default.AntPrefix + "-alert"
var msgColor = themes.Default.HeadingColor
var textColor = themes.Default.TextColor
var closeColor = themes.Default.TextColor

// Alert returns antd styles for alerts.
func Alert() gs.CSSRule {
	a := themes.Default.FontSizeBase.Value()
	b := themes.Default.LineHeightBase.Value()
	c := 8 + a*b/2 - a/2 + 1
	successColor := color.Generate(themes.Default.SuccessColor)
	successBorder := fmt.Sprintf("%v %v %v", themes.Default.BorderWithBase,
		themes.Default.BorderStyleBase, successColor[2].Hex())
	infoColor := color.Generate(themes.Default.InfoColor)
	infoBorder := fmt.Sprintf("%v %v %v", themes.Default.BorderWithBase,
		themes.Default.BorderStyleBase, infoColor[2].Hex())
	warningColor := color.Generate(themes.Default.WarningColor)
	warningBorder := fmt.Sprintf("%v %v %v", themes.Default.BorderWithBase,
		themes.Default.BorderStyleBase, warningColor[2].Hex())
	errorColor := color.Generate(themes.Default.ErrorColor)
	errorBorder := fmt.Sprintf("%v %v %v", themes.Default.BorderWithBase,
		themes.Default.BorderStyleBase, errorColor[2].Hex())
	return gs.CSS(
		gs.S(klass,
			mixins.ResetComponent(),
			gs.P("position", "relative"),
			gs.P("padding", "8px 15px 8px 37px"),
			gs.P("border-radius", themes.Default.BorderRadiusBase.String()),
			gs.S("&&-no-icon",
				gs.P("padding", "8px 15px"),
			),
			gs.S("&-icon",
				gs.P("top", formatFloat(c)+"px"),
				gs.P("left", "16px"),
				gs.P("position", "absolute"),
			),
			gs.S("&-description",
				gs.P("font-size", themes.Default.FontSizeBase.String()),
				gs.P("line-height", "22px"),
				gs.P("display", "none"),
			),
			gs.S("&-success",
				gs.P("border", successBorder),
				gs.P("background-color", successColor[0].Hex()),
				gs.S("& "+klass+"-icon",
					gs.P("color", themes.Default.SuccessColor.Hex()),
				),
			),
			gs.S("&-info",
				gs.P("border", infoBorder),
				gs.P("background-color", infoColor[0].Hex()),
				gs.S("& "+klass+"-icon",
					// BUG: For some reason that I have no clue. This gives a different value from
					// the one on antd.
					//
					// Here we are getting #178fff; while antd has #1890ff
					gs.P("color", themes.Default.InfoColor.Hex()),
				),
			),
			gs.S("&-warning",
				gs.P("border", warningBorder),
				gs.P("background-color", warningColor[0].Hex()),
				gs.S("& "+klass+"-icon",
					gs.P("color", themes.Default.WarningColor.Hex()),
				),
			),
			gs.S("&-error",
				gs.P("border", errorBorder),
				gs.P("background-color", errorColor[0].Hex()),
				gs.S("& "+klass+"-icon",
					gs.P("color", themes.Default.ErrorColor.Hex()),
				),
			),
		),
	)
}

func formatFloat(v float64) string {
	return strconv.FormatFloat(v, 'f', 1, 64)
}
