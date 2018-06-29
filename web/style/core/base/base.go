package base

import (
	"github.com/gernest/gs"
	"github.com/gernest/vected/web/mixins"
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
	)
}
