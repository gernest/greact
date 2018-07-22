package style

import (
	"github.com/gernest/vected/lib/gs"
	"github.com/gernest/vected/web/style/core/themes"
)

var prefix = themes.Default.AntPrefix

// Affix returns css style for antd affix.
func Affix() gs.CSSRule {
	return gs.S(prefix+"-affix",
		gs.P("position", "fixed"),
		gs.P("z-index", themes.Default.ZIndexAffix),
	)
}
