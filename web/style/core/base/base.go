package base

import (
	"github.com/gernest/gs"
	"github.com/gernest/vected/web/mixins"
	"github.com/gernest/vected/web/themes"
)

// Base is a direct port of https://github.com/ant-design/ant-design/blob/master/components/style/core/base.less
// to gs
//
// Reboot
//
// Normalization of HTML elements, manually forked from Normalize.css to remove
// styles targeting irrelevant browsers while applying new styles.
//
// Normalize is licensed MIT. https://github.com/necolas/normalize.css
//
// http://stackoverflow.com/a/13611748/3040605
func Base() gs.CSSRule {
	return gs.CSS(
		gs.FontFace(
			gs.P("font-family", `"Monospaced Number"`),
			gs.P("src", `local("Tahoma")`),
			gs.P("unicode-range", "U+30-39"),
		),
		gs.FontFace(
			gs.P("font-family", `"Monospaced Number"`),
			gs.P("font-weight", "bold"),
			gs.P("src", `local("Tahoma-Bold")`),
			gs.P("unicode-range", "U+30-39"),
		),
		gs.FontFace(
			gs.P("font-family", `"Chinese Quote"`),
			gs.P("font-weight", "bold"),
			gs.P("src", `local("PingFang SC"), local("SimSun")`),
			gs.P("unicode-range", "U+2018, U+2019, U+201c, U+201d"),
		),
		// HTML & Body reset
		gs.S("html", gs.S("&,body"),
			mixins.Square("100%"),
		),
		// remove the clear button of a text input control in IE10+
		gs.S("input::-ms-clear", gs.S("&,input::-ms-reveal",
			gs.P("display", "none"),
		)),
		// Document
		//
		// 1. Change from `box-sizing: content-box` so that `width` is not affected by `padding` or `border`.
		// 2. Change the default font family in all browsers.
		// 3. Correct the line height in all browsers.
		// 4. Prevent adjustments of font size after orientation changes in IE on Windows Phone and in iOS.
		// 5. Setting @viewport causes scrollbars to overlap content in IE11 and Edge, so
		//    we force a non-overlapping, non-auto-hiding scrollbar to counteract.
		// 6. Change the default tap highlight to be completely transparent in iOS.
		gs.S("*", gs.S("&,*:::before", gs.S("&,*::after",
			gs.P("box-sizing", "border-box"),
		))),
		gs.S("html",
			gs.P("font-family", "sans-serif"),
			gs.P("line-height", "1.15"),
			gs.P("-webkit-text-size-adjust", "100%"),
			gs.P("-ms-text-size-adjust", "100%"),
			gs.P("-ms-overflow-style", "scrollbar"),
			gs.P("-webkit-tap-highlight-color", "rgba(0, 0, 0, 0)"),
		),
		// IE10+ doesn't honor `<meta name="viewport">` in some cases.
		gs.Cond("@-ms-viewport ",
			gs.P("width", "device-width"),
		),
		// Shim for "new" HTML5 structural elements to display correctly (IE10, older browsers)
		gs.S("article", gs.S("&,aside", gs.S("&,dialog", gs.S("&,figcaption", gs.S("&,figure", gs.S("&,footer", gs.S("&,header", gs.S("&,hgroup", gs.S("&,main", gs.S("&,nav", gs.S("&,section",
			gs.P("display", "block"),
		))))))))))),
		// Body
		//
		// 1. remove the margin in all browsers.
		// 2. As a best practice, apply a default `body-background`.
		gs.S("body",
			gs.P("margin", "0"),
			gs.P("font-family", themes.Default.FontFamily),
			gs.P("font-size", themes.Default.FontSizeBase),
			gs.P("line-height", themes.Default.LineHeightBase),
			gs.P("color", themes.Default.TextColor.String()),
			gs.P("background-color", themes.Default.BodyBackground.String()),
		),
		// Suppress the focus outline on elements that cannot be accessed via keyboard.
		// This prevents an unwanted focus outline from appearing around elements that
		// might still respond to pointer events.
		//
		// Credit: https://github.com/suitcss/base
		gs.S(`[tabindex="-1"]`,
			gs.S("&:focus",
				gs.P("outline", `one !important`),
			),
		),
		// Content grouping
		//
		// 1. Add the correct box sizing in Firefox.
		// 2. Show the overflow in Edge and IE.
		gs.S("hr",
			gs.P("box-sizing", "content-box"),
			gs.P("height", "0"),
			gs.P("overflow", "visible"),
		),
	)
}
