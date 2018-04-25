package grid

import (
	"strconv"

	"github.com/gernest/gs"
	"github.com/gernest/vected/style/mixins"
)

func MakeRow(gutter int64) gs.CSSRule {
	return gs.CSS(
		gs.P("position", "relative"),
		gs.P("margin-left", format(gutter/-2)),
		gs.P("margin-right", format(gutter/-2)),
		gs.P("height", "auto"),
		mixins.ClearFix(),
	)
}

func format(v int64) string {
	return strconv.FormatInt(v, 10)
}
