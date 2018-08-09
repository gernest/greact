package style

import (
	"fmt"

	"github.com/gernest/vected/lib/gs"
	"github.com/gernest/vected/web/style/color"
	"github.com/gernest/vected/web/style/core/themes"
	"github.com/gernest/vected/web/style/mixins"
)

var (
	ghostColor  = themes.Default.TextColor
	ghostBG     = "transparent"
	ghostBorder = themes.Default.BorderColorBase
	prefix      = themes.Default.AntPrefix + "-btn"
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
			gs.S("&",
				btnColor(themes.Default.BtnDisabledColor.String(),
					themes.Default.BtnDisabledBG.Hex(),
					themes.Default.BtnDisabledBorder.Hex()),
			),
			gs.S("&:hover",
				btnColor(themes.Default.BtnDisabledColor.String(),
					themes.Default.BtnDisabledBG.Hex(),
					themes.Default.BtnDisabledBorder.Hex()),
			),
			gs.S("&:focus",
				btnColor(themes.Default.BtnDisabledColor.String(),
					themes.Default.BtnDisabledBG.Hex(),
					themes.Default.BtnDisabledBorder.Hex()),
			),
			gs.S("&:active",
				btnColor(themes.Default.BtnDisabledColor.String(),
					themes.Default.BtnDisabledBG.Hex(),
					themes.Default.BtnDisabledBorder.Hex()),
			),
			gs.S("&.active",
				btnColor(themes.Default.BtnDisabledColor.String(),
					themes.Default.BtnDisabledBG.Hex(),
					themes.Default.BtnDisabledBorder.Hex()),
			),
		),
		gs.S("&[disabled]",
			gs.S("&",
				btnColor(themes.Default.BtnDisabledColor.String(),
					themes.Default.BtnDisabledBG.Hex(),
					themes.Default.BtnDisabledBorder.Hex()),
			),
			gs.S("&:hover",
				btnColor(themes.Default.BtnDisabledColor.String(),
					themes.Default.BtnDisabledBG.Hex(),
					themes.Default.BtnDisabledBorder.Hex()),
			),
			gs.S("&:focus",
				btnColor(themes.Default.BtnDisabledColor.String(),
					themes.Default.BtnDisabledBG.Hex(),
					themes.Default.BtnDisabledBorder.Hex()),
			),
			gs.S("&:active",
				btnColor(themes.Default.BtnDisabledColor.String(),
					themes.Default.BtnDisabledBG.Hex(),
					themes.Default.BtnDisabledBorder.Hex()),
			),
			gs.S("&.active",
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
		gs.P("background-color", background),
		gs.P("border-color", border),
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
		gs.S("&:hover", btnColor(c, a, a)),
		gs.S("&:focus", btnColor(c, a, a)),
		gs.S("&:active", btnColor(c, b, b)),
		gs.S("&.active", btnColor(c, b, b)),
		disabled(),
	)
}

func variantOther(color, background, border string) gs.CSSRule {
	a := themes.Default.Primary[4]
	b := themes.Default.Primary[6]
	return gs.CSS(
		btnColor(color, background, border),
		gs.S("&:hover", btnColor(a.Hex(), background, a.Hex())),
		gs.S("&:focus", btnColor(a.Hex(), background, a.Hex())),
		gs.S("&:active", btnColor(b.Hex(), background, b.Hex())),
		gs.S("&.active", btnColor(b.Hex(), background, b.Hex())),
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
				themes.Default.BtnFontSizeLG.String(), "0"),
			gs.P("line-height", themes.Default.BtnHeighLG.Sub(gs.U("2px")).String()),
		),
		gs.S("&-lg > span > "+klass,
			size(themes.Default.BtnHeighLG.String(), themes.Default.BtnPaddingLG,
				themes.Default.BtnFontSizeLG.String(), "0"),
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
		gs.S("&, &:active,&:focus",
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
		),
		gs.S("&.disabled > *,&[disabled] > *",
			gs.P("pointer-events", "none"),
		),
		gs.S("&-lg",
			size(themes.Default.BtnHeighLG.String(),
				themes.Default.BtnPaddingLG, themes.Default.BtnFontSizeLG.String(),
				themes.Default.BtnBorderRadiusBase),
		),
		gs.S("&-sm",
			size(themes.Default.BtnHeighSM.String(),
				themes.Default.BtnPaddingSM, themes.Default.BtnFontSizeSM.String(),
				themes.Default.BtnBorderRadiusSM),
		),
	)
}

func primary() gs.CSSRule {
	return variantPrimary(themes.Default.BtnPrimaryColor.Hex(),
		themes.Default.BtnPrimaryBG)
}

func btnDefault() gs.CSSRule {
	return gs.CSS(
		variantOther(themes.Default.BtnDefaultColor.Hex(),
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

func circle(klass string) gs.CSSRule {
	return gs.CSS(
		mixins.Square(themes.Default.BtnCircleSize.String()),
		size(themes.Default.BtnCircleSize.String(),
			"0", themes.Default.FontSizeBase.Add(gs.U("2px")).String(), "50%"),
		gs.S("&"+klass+"-lg",
			mixins.Square(themes.Default.BtnCircleSizeLG.String()),
			size(themes.Default.BtnCircleSizeLG.String(),
				"0", themes.Default.BtnFontSizeLG.Add(gs.U("2px")).String(), "50%"),
		),
		gs.S("&"+klass+"-sm",
			mixins.Square(themes.Default.BtnCircleSizeSM.String()),
			size(themes.Default.BtnCircleSizeSM.String(),
				"0", themes.Default.FontSizeBase.Add(gs.U("2px")).String(), "50%"),
		),
	)
}

func group(klass string) gs.CSSRule {
	s := "& %s+%s , %s+&,& span+ %s,& %s+span,& > span+span, &+%s ,&+&"
	p := "%s-primary+%s:not(%s-primary):not([disabled])"
	r := "&> %s:first-child:not(:last-child),& > span:first-child:not(:last-child) > %s"
	return gs.CSS(
		groupBase(klass),
		gs.S(fmt.Sprintf(s, klass, klass, klass, klass, klass, klass),
			gs.P("margin-left", "-1px"),
		),
		gs.S(fmt.Sprintf(p, klass, klass, klass),
			gs.P("border-left-color", "transparent"),
		),
		gs.S(klass,
			gs.P("border-radius", "0"),
		),
		gs.S(fmt.Sprintf("& > %s:first-child,& > span:first-child >%s", klass, klass),
			gs.P("margin-left", "0"),
		),
		gs.S(fmt.Sprintf("&> %s:only-child", klass),
			gs.P("border-radius", themes.Default.BtnBorderRadiusBase),
		),
		gs.S(fmt.Sprintf("&> span:only-child > %s", klass),
			gs.P("border-radius", themes.Default.BtnBorderRadiusBase),
		),
		gs.S(fmt.Sprintf(r, klass, klass),
			gs.P("border-bottom-left-radius", themes.Default.BtnBorderRadiusBase),
			gs.P("border-top-left-radius", themes.Default.BtnBorderRadiusBase),
		),
	)
}

// Button returns antd button style
func Button() gs.CSSRule {
	return gs.CSS(
		gs.S(prefix,
			gs.P("line-height", themes.Default.LineHeightBase.String()),
			btn(),
			btnDefault(),
			gs.S("& >i,&>span",
				gs.P("pointer-events", "none"),
			),
			gs.S("&-primary",
				primary(),
				gs.S(prefix+"-group &:not(:first-child):not(:last-child)",
					gs.P("border-right-color", themes.Default.BtnGroupBorder.Hex()),
					gs.P("border-left-color", themes.Default.BtnGroupBorder.Hex()),
					gs.S("&:disabled",
						gs.P("border-color", themes.Default.BtnDefaultBorder.Hex()),
					),
				),
				gs.S(prefix+"-group &:first-child",
					gs.S("&:not(:last-child)",
						gs.P("border-right-color", themes.Default.BtnGroupBorder.Hex()),
						gs.S("&[disabled]",
							gs.P("border-right-color", themes.Default.BtnDefaultBorder.Hex()),
						),
					),
				),
				gs.S(prefix+"-group &:last-child:not(:first-child)",
					gs.P("border-left-color", themes.Default.BtnGroupBorder.Hex()),
					gs.S("&[disabled]",
						gs.P("border-left-color", themes.Default.BtnDefaultBorder.Hex()),
					),
				),
				gs.S(prefix+"-group & + &",
					gs.P("border-left-color", themes.Default.BtnGroupBorder.Hex()),
					gs.S("&[disabled]",
						gs.P("border-left-color", themes.Default.BtnDefaultBorder.Hex()),
					),
				),
			),
			gs.S("&-ghost",
				ghost(),
			),
			gs.S("&-dashed",
				dashed(),
			),
			gs.S("&-danger",
				danger(),
			),
			gs.S("&-circle",
				circle(prefix),
			),
			gs.S("&-circle-outline",
				circle(prefix),
			),
			gs.S(" &:before",
				gs.P("position", "absolute"),
				gs.P("top", "-1px"),
				gs.P("left", "-1px"),
				gs.P("bottom", "-1px"),
				gs.P("right", "-1px"),
				gs.P("background", "#fff"),
				gs.P("opacity", "0.35"),
				gs.P("content", "''"),
				gs.P("border-radius", "inherit"),
				gs.P("z-index", "1"),
				gs.P("transition", "opacity .2s"),
				gs.P("pointer-events", "none"),
				gs.P("display", "none"),
			),
			gs.S(themes.Default.IconFontPrefix,
				gs.P("transition", "margin-left .3s "+themes.Default.EaseInOut),
			),
			gs.S("&&-loading:before",
				gs.P("display", "block"),
			),
			gs.S(" &&-loading:not(&-circle):not(&-circle-outline):not(&-icon-only)",
				gs.P("padding-left", "29px"),
				gs.P("pointer-events", "none"),
				gs.P("position", "relative"),
				gs.S(themes.Default.IconFontPrefix,
					gs.P("margin-left", "-14px"),
				),
			),
			gs.S("&-sm&-loading:not(&-circle):not(&-circle-outline):not(&-icon-only) ",
				gs.P("padding-left", "24px"),
				gs.S(themes.Default.IconFontPrefix,
					gs.P("margin-left", "-17px"),
				),
			),
			gs.S("&-group",
				group(prefix),
			),
		),
	)
}
