package style

import (
	"github.com/gernest/vected/lib/gs"
	"github.com/gernest/vected/web/style/core/themes"
	"github.com/gernest/vected/web/style/mixins"
)

var prefix = themes.Default.AntPrefix + "-breadcrumb"

// Breadcrumb defines styles for antd Breadcrumb component
func Breadcrumb() gs.CSSRule {
	return gs.CSS(
		gs.S(prefix,
			mixins.ResetComponent(),
			gs.P("color", themes.Default.BreadcrumbBaseColor.String()),
			gs.P("font-size", themes.Default.BreadcrumbFontSize.String()),
			gs.S("."+themes.Default.IconFontPrefix,
				gs.P("font-size", themes.Default.BreadcrumbIconFontSize.String()),
			),
			gs.S("a",
				gs.P("color", themes.Default.BreadcrumbLinkColor.String()),
				gs.P("transition", "color .3s"),
				gs.S("&:hover",
					gs.P("color", themes.Default.BreadcrumbLinkColorHover.Hex()),
				),
			),
			gs.S("& > span:last-child",
				gs.P("color", themes.Default.BreadcrumbLastItemColor.String()),
			),
			gs.S("& > span:last-child &-separator",
				gs.P("display", "none"),
			),
			gs.S("&-separator",
				gs.P("margin", themes.Default.BreadcrumbSeparatorMargin),
				gs.P("color", themes.Default.BreadcrumbSeparatorColor.String()),
			),
			gs.S("&-link",
				gs.S("> ."+themes.Default.IconFontPrefix+" + span",
					gs.P("margin-left", "4px"),
				),
			),
		),
	)
}
