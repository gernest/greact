package affix

import (
	"github.com/gernest/vected/lib/gs"
	"github.com/gernest/vected/web/style/core/themes"
)

var prefix = themes.Default.AntPrefix

// Style returns css style for antd affix.
func Style() gs.CSSRule {
	return gs.S(prefix+"-affix",
		gs.P("position", "fixed"),
		gs.P("z-index", themes.Default.ZIndexAffix),
	)
}
