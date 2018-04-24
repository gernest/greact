package normalize

import (
	"strings"

	"github.com/gernest/gs"
	"github.com/gopherjs/gopherjs/js"
)

func New() gs.CSSRule {
	return gs.CSS(
		gs.S("html",
			gs.P("line-height", "1.15"),
			gs.P("-webkit-text-size-adjust", "100%"),
		),
		removeMargins(), correctFont(), correctSizing(),
		correctInheritance(), removeGray(), fixBorder(),
		fontWeight(), scaling(), correctFontSize(), preventSub(),
		fixImages(), forms(), interactives(), misc(), vendors(),
	)
}

// Remove the margin in all browsers.
func removeMargins() gs.CSSRule {
	return gs.S("body",
		gs.P("margin", "0"),
	)
}

// * Correct the font size and margin on `h1` elements within `section` and
// * `article` contexts in Chrome, Firefox, and Safari.
func correctFont() gs.CSSRule {
	return gs.S("h1",
		gs.P("font-size", "2em"),
		gs.P("margin", "0.67em 0"),
	)
}

// * 1. Add the correct box sizing in Firefox.
// * 2. Show the overflow in Edge and IE.
func correctSizing() gs.CSSRule {
	return gs.S("br",
		gs.P("box-sizing", "content-box"),
		gs.P("height", "0"),
		gs.P("overflow", "visible"),
	)
}

// * 1. Correct the inheritance and scaling of font size in all browsers.
// * 2. Correct the odd `em` font sizing in all browsers.
func correctInheritance() gs.CSSRule {
	return gs.S("pre",
		gs.P("font-family", "monospace, monospace"),
		gs.P("font-size", "1em"),
	)
}

// * Remove the gray background on active links in IE 10.
func removeGray() gs.CSSRule {
	return gs.S("a",
		gs.P("background-color", "transparent"),
	)
}

// * 1. Remove the bottom border in Chrome 57-
// * 2. Add the correct text decoration in Chrome, Edge, IE, Opera, and Safari
func fixBorder() gs.CSSRule {
	return gs.S("abbr[title]",
		gs.P("border-bottom", "none"),
		gs.P("text-decoration", "underline"),
		gs.P("text-decoration", "underline dotted"),
	)
}

// * Add the correct font weight in Chrome, Edge, and Safari.
func fontWeight() gs.CSSRule {
	return gs.S("b", gs.S("&,strong",
		gs.P("font-weight", "bolder"),
	))
}

// * 1. Correct the inheritance and scaling of font size in all browsers.
// * 2. Correct the odd `em` font sizing in all browsers.
func scaling() gs.CSSRule {
	return gs.S("code",
		gs.S("&,kbd",
			gs.S("&,samp",
				gs.P("font-family", "monospace, monospace;"),
				gs.P("font-size", "1em"),
			),
		),
	)
}

// * Add the correct font size in all browsers.
func correctFontSize() gs.CSSRule {
	return gs.S("small",
		gs.P("font-size", "80%"),
	)
}

// * Prevent `sub` and `sup` elements from affecting the line height in
// * all browsers.
func preventSub() gs.CSSRule {
	return gs.CSS(
		gs.S("sub", gs.S("&,sup",
			gs.P("font-size", "75%"),
			gs.P("line-height", "0"),
			gs.P("position", "relative"),
			gs.P("vertical-align", "baseline"),
		)),
		gs.S("sub", gs.P("bottom", "-0.25em")),
		gs.S("sup", gs.P("top", "-0.5e")),
	)
}

// * Remove the border on images inside links in IE 10.
func fixImages() gs.CSSRule {
	return gs.S("img", gs.P("border-style", "none"))
}

func forms() gs.CSSRule {
	return gs.CSS(
		gs.S("button",
			gs.S("&,input",
				gs.S("&,optgroup",
					gs.S("&,select",
						gs.S("&,textarea",
							gs.P("font-family", "inherit"),
							gs.P("font-size", "100%"),
							gs.P("line-height", "1.15"),
							gs.P("margin", "0"),
						),
					),
				),
			),
		),
		gs.S("button", gs.S("&,input",
			gs.P("overflow", "visible"),
		)),
		gs.S("button", gs.S("&,select",
			gs.P("text-transform", "none"),
		)),
		gs.S("button",
			gs.S(`&,[type="button"]`,
				gs.S(`&,[type="reset"]`,
					gs.S(`&,[type="submit"]`,
						gs.P("-webkit-appearance", "button"),
					),
				),
			),
		),
		gs.S("fieldset",
			gs.P("padding", "0.35em 0.75em 0.625em"),
		),
		gs.S("legend",
			gs.P("box-sizing", "border-box"),
			gs.P("color", "inherit"),
			gs.P("display", "table"),
			gs.P("max-width", "100%"),
			gs.P("padding", "0"),
			gs.P("white-space", "normal"),
		),
		gs.S("progress", gs.P("vertical-align", "baseline")),
		gs.S("textarea", gs.P("overflow", "auto")),
		gs.S(`[type="checkbox"]`,
			gs.S(`&,[type="radio"]`,
				gs.P("box-sizing", "border-box"),
				gs.P("padding", "0"),
			),
		),
		gs.S(`[type="search"]`,
			gs.P("-webkit-appearance", "textfield"),
			gs.P("outline-offset", "-2px"),
		),
	)
}

func agent() string {
	if js.Global != nil {
		txt := js.Global.Get("navigator").Get("userAgent").String()
		if strings.Contains(txt, "Chrome") {
			return "chrome"
		}
		if strings.Contains(txt, "Firefox") {
			return "firefox"
		}
	}
	return ""
}

func vendors() gs.CSSRule {
	switch agent() {
	case "chrome":
		return chromePrefix()
	case "firefox":
		return mozPrefix()
	default:
		return gs.CSS(
			chromePrefix(),
			mozPrefix(),
		)
	}
}

func chromePrefix() gs.CSSRule {
	return gs.CSS(
		gs.S(`[type="number"]::-webkit-inner-spin-button`,
			gs.S(`&,[type="number"]::-webkit-outer-spin-button`,
				gs.P("height", "auto"),
			),
		),
		gs.S(`[type="search"]::-webkit-search-decoration`,
			gs.P("-webkit-appearance", "none"),
		),
		gs.S("::-webkit-file-upload-button",
			gs.P("-webkit-appearance", "button"),
			gs.P("font", "inherit"),
		),
	)
}

func mozPrefix() gs.CSSRule {
	return gs.CSS(gs.S("button::-moz-focus-inner",
		gs.S(`&,[type="button"]::-moz-focus-inner`,
			gs.S(`&,[type="reset"]::-moz-focus-inner`,
				gs.S(`&,[type="submit"]::-moz-focus-inner`,
					gs.P("border-style", "none"),
					gs.P("padding", "0"),
				),
			),
		),
	),
		gs.S("button:-moz-focusring",
			gs.S(`&,[type="button"]:-moz-focusring`,
				gs.S(`&,[type="reset"]:-moz-focusring`,
					gs.S(`&,[type="submit"]:-moz-focusring`,
						gs.P("outline", "1px dotted ButtonText"),
					),
				),
			),
		))
}

func interactives() gs.CSSRule {
	return gs.CSS(
		gs.S("details",
			gs.P("display", "block"),
		),
		gs.S("summary",
			gs.P("display", "list-item"),
		),
	)
}

func misc() gs.CSSRule {
	return gs.CSS(
		gs.S("template",
			gs.P("display", "none"),
		),
		gs.S("[hidden]",
			gs.P("display", "none"),
		),
	)
}
