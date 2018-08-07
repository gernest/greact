package style

import (
	"fmt"

	"github.com/gernest/vected/lib/gs"
	"github.com/gernest/vected/web/style/color"
	"github.com/gernest/vected/web/style/core/themes"
)

var (
	ghostColor  = themes.Default.TextColor
	ghostBG     = "transparent"
	ghostBorder = themes.Default.BorderColorBase
)

func size(height, padding, fontSize, borderRadius string) gs.CSSRule {
	return gs.CSS(
		gs.P("padding", padding),
		gs.P("font-size", fontSize),
		gs.P("border-radius", borderRadius),
		gs.P("height", height),
	)
}

func disabled() gs.CSSRule {
	return gs.CSS(
		gs.S("&.disabled",
			gs.S("&, &:hover, &:focus, &:active, &.active",
				btnColor(themes.Default.BtnDisabledColor.String(),
					themes.Default.BtnDisabledBG.Hex(),
					themes.Default.BtnDisabledBorder.Hex()),
			),
		),
		gs.S("&[disabled]",
			gs.S("&, &:hover, &:focus, &:active, &.active",
				btnColor(themes.Default.BtnDisabledColor.String(),
					themes.Default.BtnDisabledBG.Hex(),
					themes.Default.BtnDisabledBorder.Hex()),
			),
		),
	)
}

func btnColor(color, background, border string) gs.CSSRule {
	return gs.CSS(
		gs.P("color", color),
		gs.P("background-colo", border),
		gs.S("> a:only-child",
			gs.P("color", "currentColor"),
			gs.S("&:after",
				gs.P("content", "''"),
				gs.P("position", "absolute"),
				gs.P("top", "0"),
				gs.P("left", "0"),
				gs.P("bottom", "0"),
				gs.P("right", "0"),
				gs.P("background", "transparent"),
			),
		),
	)
}

func variantPrimary(c string, background *color.Color) gs.CSSRule {
	bgHex := background.Hex()
	bg := color.Generate(background)
	a := bg[4].Hex()
	b := bg[6].Hex()
	return gs.CSS(
		btnColor(c, bgHex, bgHex),
		gs.S("&:hover, &:focus", btnColor(c, a, a)),
		gs.S("&:active, &.active", btnColor(c, b, b)),
		disabled(),
	)
}

func variantOther(color, background, border string) gs.CSSRule {
	a := themes.Default.Primary[4]
	b := themes.Default.Primary[6]
	return gs.CSS(
		btnColor(color, background, border),
		gs.S("&:hover, &:focus", btnColor(a.String(), background, a.Hex())),
		gs.S("&:active, &.active", btnColor(b.String(), background, b.Hex())),
		disabled(),
	)
}
func variantDanger(c, background *color.Color, border string) gs.CSSRule {
	g := color.Generate(c)
	a := g[4]
	b := g[6]
	return gs.CSS(
		btnColor(c.String(), background.Hex(), border),
		gs.S("&:hover",
			btnColor(themes.Default.PrimaryColor.String(),
				a.Hex(), a.Hex())),
		gs.S("&:focus", btnColor(b.String(), "#fff", a.Hex())),
		gs.S("&:active, &.active", btnColor(themes.Default.PrimaryColor.String(),
			b.Hex(), b.Hex())),
		disabled(),
	)
}

func variantGhost(c *color.Color) gs.CSSRule {
	g := color.Generate(c)
	a := g[4]
	b := g[6]
	return gs.CSS(
		btnColor(c.String(), "transparent", c.Hex()),
		gs.S("&:hover, &:focus", btnColor(a.String(), "transparent", a.Hex())),
		gs.S("&:active, &.active", btnColor(b.String(), "transparent", b.Hex())),
		disabled(),
	)
}

func groupBase(klass string) gs.CSSRule {
	return gs.CSS(
		gs.P("position", "relative"),
		gs.P("display", "inline-block"),
		gs.S("> "+klass,
			gs.P("position", "relative"),
			gs.P("line-height", themes.Default.BtnHeighBase.Sub(gs.U("2px")).String()),
			gs.S("&:hover,&:focus,&:active,&.active",
				gs.P("z-index", "2"),
			),
			gs.S("&:disabled",
				gs.P("z-index", "0"),
			),
		),
		gs.S("> span > "+klass,
			gs.P("position", "relative"),
			gs.P("line-height", themes.Default.BtnHeighBase.Sub(gs.U("2px")).String()),
			gs.S("&:hover,&:focus,&:active,&.active",
				gs.P("z-index", "2"),
			),
			gs.S("&:disabled",
				gs.P("z-index", "0"),
			),
		),
		gs.S("&-lg > "+klass,
			size(themes.Default.BtnHeighLG.String(), themes.Default.BtnPaddingLG,
				themes.Default.BtnFontSizeLG, "0"),
			gs.P("line-height", themes.Default.BtnHeighLG.Sub(gs.U("2px")).String()),
		),
		gs.S("&-lg > span > "+klass,
			size(themes.Default.BtnHeighLG.String(), themes.Default.BtnPaddingLG,
				themes.Default.BtnFontSizeLG, "0"),
			gs.P("line-height", themes.Default.BtnHeighLG.Sub(gs.U("2px")).String()),
		),
		gs.S("&-sm > "+klass,
			size(themes.Default.BtnHeighSM.String(), themes.Default.BtnPaddingSM,
				themes.Default.FontSizeBase.String(), "0"),
			gs.P("line-height", themes.Default.BtnHeighSM.Sub(gs.U("2px")).String()),
		),
		gs.S("&-sm > span > "+klass,
			size(themes.Default.BtnHeighSM.String(), themes.Default.BtnPaddingSM,
				themes.Default.FontSizeBase.String(), "0"),
			gs.P("line-height", themes.Default.BtnHeighSM.Sub(gs.U("2px")).String()),
		),
	)
}

func btn() gs.CSSRule {
	return gs.CSS(
		gs.P("display", "inline-block"),
		gs.P("font-weight", themes.Default.BtnFontWeight),
		gs.P("text-align", "center"),
		gs.P("touch-action", "manipulation"),
		gs.P("cursor", "pointer"),
		gs.P("background", "none"),
		gs.P("background-image", "none"),
		gs.P("border",
			fmt.Sprintf("%s %s transparent", themes.Default.BorderWithBase,
				themes.Default.BorderStyleBase)),
		gs.P("white-space", "nowrap"),
		size(themes.Default.BtnHeighBase.String(),
			themes.Default.BtnPaddingBase, themes.Default.FontSizeBase.String(),
			themes.Default.BtnBorderRadiusBase),
		gs.P("user-select", "none"),
		gs.P("transition", "all .3s "+themes.Default.EaseInOut),
		gs.P("position", "relative"),
		gs.S("> "+themes.Default.IconFontPrefix,
			gs.P("line-height", "1"),
		),
		gs.S("&, &:action,&:focus",
			gs.P("outline", "0"),
		),
		gs.S("&:not([disabled]):hove",
			gs.P("text-decoration", "none"),
		),
		gs.S("&:not([disabled]):active",
			gs.P("outline", "0"),
			gs.P("transition", "none"),
		),
		gs.S("&.disabled,&[disabled]",
			gs.P("cursor", "not-allowed"),
			gs.S("> *",
				gs.P("pointer-events", "none"),
			),
		),
		gs.S("&-lg",
			size(themes.Default.BtnHeighLG.String(),
				themes.Default.BtnPaddingLG, themes.Default.BtnFontSizeLG,
				themes.Default.BtnBorderRadiusBase),
		),
		gs.S("&-sm",
			size(themes.Default.BtnHeighSM.String(),
				themes.Default.BtnPaddingSM, themes.Default.BtnFontSizeSM,
				themes.Default.BtnBorderRadiusSM),
		),
	)
}

func primary() gs.CSSRule {
	return variantPrimary(themes.Default.BtnPrimaryColor.String(),
		themes.Default.BtnPrimaryBG)
}

func btnDefault() gs.CSSRule {
	return gs.CSS(
		variantOther(themes.Default.BtnDefaultColor.String(),
			themes.Default.BtnDefaultBG.Hex(), themes.Default.BtnDefaultBorder.Hex()),
		gs.S("&:hover,&:focus,&:active,&.active",
			gs.P("background", themes.Default.BtnDefaultBG.Hex()),
			gs.P("text-decoration", "none"),
		),
	)
}

func ghost() gs.CSSRule {
	return variantOther(ghostColor.String(), ghostBG, ghostBorder.Hex())
}

func dashed() gs.CSSRule {
	return gs.CSS(
		variantOther(themes.Default.BtnDefaultColor.String(),
			themes.Default.BtnDefaultBG.Hex(), themes.Default.BtnDefaultBorder.Hex()),
		gs.P("border-style", "dashed"),
	)
}

func danger() gs.CSSRule {
	return variantDanger(themes.Default.BtnDangerColor,
		themes.Default.BtnDangerBG, themes.Default.BtnDangerBorder.Hex())
}
