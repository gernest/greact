package base

import (
	"github.com/gernest/gs"
	"github.com/gernest/vected/web/mixins"
	"github.com/gernest/vected/web/style/core/themes"
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
		// Typography
		//

		// remove top margins from headings
		//
		// By default, `<h1>`-`<h6>` all receive top and bottom margins. We nuke the top
		// margin for easier control within type scales as it avoids margin collapsing.
		gs.S("h1", gs.S("&,h2", gs.S("&,h3", gs.S("&,h4", gs.S("&,h5",
			gs.P("margin-top", "0"),
			gs.P("margin-bottom:", ".5em"),
			gs.P("color", themes.Default.HeadingColor.String()),
			gs.P("font-weight", "500"),
		))))),
		// Reset margins on paragraphs
		//
		// Similarly, the top margin on `<p>`s get reset. However, we also reset the
		// bottom margin to use `em` units instead of `em`.
		gs.S("p",
			gs.P("margin-top", "0"),
			gs.P("margin-bottom", "1em"),
		),
		// Abbreviations
		//
		// 1. remove the bottom border in Firefox 39-.
		// 2. Add the correct text decoration in Chrome, Edge, IE, Opera, and Safari.
		// 3. Add explicit cursor to indicate changed behavior.
		// 4. Duplicate behavior to the data-* attribute for our tooltip plugin
		gs.S("abbr[title]",
			gs.S("&,abbr[data-original-title]",
				gs.P("text-decoration", "underline"),
				gs.P("text-decoration", "underline dotted"),
				gs.P("cursor", "help"),
				gs.P("border-bottom", "0"),
			),
		),
		gs.S("address",
			gs.P("margin-bottom", "1em"),
			gs.P("font-style", "normal"),
			gs.P("line-height", "inherit"),
		),
		gs.S(`input[type="text"]`, gs.S(`&,input[type="password"]`, gs.S(`&,input[type="number"]`, gs.S("&,textarea",
			gs.P("-webkit-appearance", "none"),
		)))),
		gs.S("ol", gs.S("&,ul", gs.S("&,dl",
			gs.P(" margin-top", "0"),
			gs.P("margin-bottom", "1em"),
		))),
		gs.S("ol ol", gs.S("&,ul ul", gs.S("&,ol ul", gs.S("&,ul ol",
			gs.P("margin-bottom", "1em"),
		)))),
		gs.S("dt",
			gs.P("font-weight", "500"),
		),
		gs.S("dd",
			gs.P("margin-bottom", ".5em"),
			gs.P("margin-left", "0"),
		),
		gs.S("blockquote",
			gs.P("margin", " 0 1em"),
		),
		gs.S("dfn",
			// Add the correct font style in Android 4.3-
			gs.P("font-style", "italic"),
		),
		gs.S("b", gs.S("&,strong",
			// Add the correct font weight in Chrome, Edge, and Safari
			gs.P("font-weight", "bolder"),
		)),
		gs.S("small",
			// Add the correct font size in all browsers
			gs.P("font-size", "80%"),
		),
		//
		// Prevent `sub` and `sup` elements from affecting the line height in
		// all browsers.
		//
		gs.S("sub", gs.S("&,sup",
			gs.P("position", "relative"),
			gs.P("font-size", "75%"),
			gs.P("line-height", "0"),
			gs.P("vertical-align", "baseline"),
		)),
		gs.S("sub", gs.P("bottom", "-.25em")),
		gs.S("sup", gs.P("top", "-.5em")),
		//
		// Links
		//
		gs.S("a",
			gs.P("color", themes.Default.LinkColor.String()),
			gs.P("background-color", "transparent"),
			gs.P("text-decoration", themes.Default.LinkDecoration),
			gs.P("outline", "none"),
			gs.P("cursor", "pointer"),
			gs.P("transition", "color .3s"),
			// remove gaps in links underline in iOS 8+ and Safari 8+.
			gs.P("-webkit-text-decoration-skip", "objects"),
			gs.S("&:focus",
				gs.P("text-decoration", "underline"),
				gs.P("text-decoration-skip", "ink"),
			),
			gs.S("&:hover", gs.P("color", themes.Default.LinkHoverColor.Hex())),
			gs.S("&:active", gs.P("color", themes.Default.LinkActiveColor.Hex())),
			gs.S("&:active", gs.S("&,a:hover",
				gs.P("outline", "0"),
				gs.P("text-decoration", themes.Default.LinkHoverDecoration),
			)),
			gs.S("&[disabled]",
				gs.P("color", themes.Default.DisabledColor.String()),
				gs.P("cursor", "not-allowed"),
				gs.P("pointer-events", "none"),
			),
		),
	)
}
