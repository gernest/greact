package button

import (
	"github.com/gernest/gs"
	"github.com/gernest/vected/style/themes"
)

var theme = themes.Default

func disable() gs.CSSRule {
	return gs.CSS(
		gs.S("&.disabled",
			buttonColor(
				theme.DisabledColor.String(),
				theme.BtnDisabledBG.String(),
				theme.BtnDisabledBorder.String(),
			),
		),
		gs.S("&[disabled]",
			gs.S("&",
				buttonColor(
					theme.DisabledColor.String(),
					theme.BtnDisabledBG.String(),
					theme.BtnDisabledBorder.String(),
				),
			),
			gs.S("&:hover",
				buttonColor(
					theme.DisabledColor.String(),
					theme.BtnDisabledBG.String(),
					theme.BtnDisabledBorder.String(),
				),
			),
			gs.S("&:focus",
				buttonColor(
					theme.DisabledColor.String(),
					theme.BtnDisabledBG.String(),
					theme.BtnDisabledBorder.String(),
				),
			),
			gs.S("&:active",
				buttonColor(
					theme.DisabledColor.String(),
					theme.BtnDisabledBG.String(),
					theme.BtnDisabledBorder.String(),
				),
			),
		),
	)
}

func buttonColor(color, background, border string) gs.CSSRule {
	return gs.CSS(
		gs.P("color", color),
		gs.P("border-color", border),
	)
}
