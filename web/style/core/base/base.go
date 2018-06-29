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
		gs.S("html", gs.S("body"),
			mixins.Square("100%"),
		),
	)
}
