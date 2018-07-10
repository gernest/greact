package grid

import (
	"fmt"
	"strconv"

	"github.com/gernest/vected/web/style/core/themes"

	"github.com/gernest/gs"
	"github.com/gernest/vected/web/style/mixins"
)

func makeRow(gutter int64) gs.CSSRule {
	return gs.CSS(
		gs.P("position", "relative"),
		gs.P("margin-left", formatInt(gutter/-2)),
		gs.P("margin-right", formatInt(gutter/-2)),
		gs.P("height", "auto"),
		mixins.ClearFix(),
	)
}

func formatInt(v int64) string {
	return strconv.FormatInt(v, 10)
}

var prefix = themes.Default.AntPrefix

// RowStyle complete styles for antd grid rows
func RowStyle() gs.CSSRule {
	return gs.CSS(
		gs.S(prefix+"-row",
			makeRow(themes.Default.GridGutterWidth),
			gs.P("display", "block"),
			gs.P("box-sizing", "border-box"),
		),
		gs.S(prefix+"-row-flex",
			gs.P("display", "flex"),
			gs.P("flex-flow", "row wrap"),
			gs.S("&:before", gs.P("display", "flex")),
			gs.S("&:after", gs.P("display", "flex")),
		),
		gs.S(prefix+"-row-flex-start", gs.P("justify-content", "flex-start")),
		gs.S(prefix+"-row-flex-center", gs.P("justify-content", "center")),
		gs.S(prefix+"-row-flex-end", gs.P("justify-content", "flex-end")),
		gs.S(prefix+"-row-flex-space-between", gs.P("justify-content", "space-between")),
		gs.S(prefix+"-row-flex-space-around", gs.P("justify-content", "space-around")),
		gs.S(prefix+"-row-flex-top", gs.P("justify-content", "flex-start")),
		gs.S(prefix+"-row-flex-middle", gs.P("justify-content", "center")),
		gs.S(prefix+"-row-flex-bottom", gs.P("justify-content", "flex-bottom")),
	)
}

func MakeCols() gs.CSSRule {
	return GenCol(1)
}
func GenCol(idx int64) gs.CSSRule {
	item := "%s-col-%d, %s-col-xs-%d, %s-col-sm-%d, %s-col-md-%d, %s-col-lg-%d"
	var v []interface{}
	for i := 0; i < 5; i++ {
		v = append(v, prefix, idx)
	}
	item = fmt.Sprintf(item, v...)
	return col(idx+1, item)
}

func col(idx int64, list string) gs.CSSRule {
	if idx <= themes.Default.GridColumns {
		items := "%s-col-%d, %s-col-xs-%d, %s-col-sm-%d, %s-col-md-%d, %s-col-lg-%d"
		var v []interface{}
		for i := 0; i < 5; i++ {
			v = append(v, prefix, idx)
		}
		items = fmt.Sprintf(items, v...)
		return col(idx+1, list+","+items)
	}
	return gs.S(list,
		gs.P("position", "relative"),
		// Prevent columns from collapsing when empty
		gs.P("min-height", "1px"),
		gs.P("padding-left", formatInt(themes.Default.GridGutterWidth/2)),
		gs.P("padding-right", formatInt(themes.Default.GridGutterWidth/2)),
	)
}

func floatCol(idx int64, klass string) gs.CSSRule {
	return floatColList(idx+1, klass, fmt.Sprintf("%s-col%s-%d", prefix, klass, idx))
}

func floatColList(idx int64, klass string, list string) gs.CSSRule {
	if idx <= themes.Default.GridColumns {
		floatColList(idx+1, klass, fmt.Sprintf("%s-col%s-%d", prefix, klass, idx))
	}
	return gs.S(list,
		gs.P("float", "left"),
		gs.P("flex", "0 0 auto"),
	)
}

func loopGridCol(idx int64, klass string) gs.CSSRule {
	var list gs.RuleList
	for i := idx; i >= 0; i-- {
		list = append(list, loopCol(i, klass))
	}
	return list
}

func loopCol(idx int64, klass string) gs.CSSRule {
	if idx == 0 {
		return gs.CSS(
			gs.S(fmt.Sprintf("%s-col%s-%d", prefix, klass, idx),
				gs.P("display", "none"),
			),
			gs.S(fmt.Sprintf("%s-col-push-%d", prefix, idx),
				gs.P("left", "auto"),
			),
			gs.S(fmt.Sprintf("%s-col-pull-%d", prefix, idx),
				gs.P("right", "auto"),
			),
			gs.S(fmt.Sprintf("%s-col%s-push-%d", prefix, klass, idx),
				gs.P("left", "auto"),
			),
			gs.S(fmt.Sprintf("%s-col%s-pull-%d", prefix, klass, idx),
				gs.P("right", "auto"),
			),
			gs.S(fmt.Sprintf("%s-col%s-offset-%d", prefix, klass, idx),
				gs.P("margin-left", "0"),
			),
			gs.S(fmt.Sprintf("%s-col%s-order-%d", prefix, klass, idx),
				gs.P("order", "0"),
			),
		)
	}
	return gs.CSS(
		gs.S(fmt.Sprintf("%s-col%s-%d", prefix, klass, idx),
			gs.P("display", "block"),
			gs.P("box-sizing", "border-box"),
			gs.P("width", percent(idx, themes.Default.GridColumns)),
		),
		gs.S(fmt.Sprintf("%s-col-push-%d", prefix, idx),
			gs.P("left", percent(idx, themes.Default.GridColumns)),
		),
		gs.S(fmt.Sprintf("%s-col-pull-%d", prefix, idx),
			gs.P("right", percent(idx, themes.Default.GridColumns)),
		),
		gs.S(fmt.Sprintf("%s-col%s-push-%d", prefix, klass, idx),
			gs.P("left", percent(idx, themes.Default.GridColumns)),
		),
		gs.S(fmt.Sprintf("%s-col%s-pull-%d", prefix, klass, idx),
			gs.P("right", percent(idx, themes.Default.GridColumns)),
		),
		gs.S(fmt.Sprintf("%s-col%s-offset-%d", prefix, klass, idx),
			gs.P("margin-left", percent(idx, themes.Default.GridColumns)),
		),
		gs.S(fmt.Sprintf("%s-col%s-order-%d", prefix, klass, idx),
			gs.P("order", formatInt(idx)),
		),
	)
}

func percent(a, b int64) string {
	v := float64(a) / float64(b)
	s := strconv.FormatFloat(v*100, 'f', 8, 64)
	return s + "%"
}

func MakeGrid(klass string) gs.CSSRule {
	return gs.CSS(
		floatCol(1, klass),
		loopGridCol(themes.Default.GridColumns, klass),
	)
}

// ColStyle complete antd  grid column styles.
func ColStyle() gs.CSSRule {
	return gs.CSS(
		MakeCols(),
		MakeGrid(""),
		// Extra small grid
		//
		// Columns, offsets, pushes, and pulls for extra small devices like
		// smartphones.
		MakeGrid("-xs"),
		// Small grid
		//
		// Columns, offsets, pushes, and pulls for the small device range, from phones
		// to tablets.
		gs.Cond("@media (min-width: "+themes.Default.ScreenSMMin+")",
			MakeGrid("-sm"),
		),
		// Medium grid
		//
		// Columns, offsets, pushes, and pulls for the desktop device range.
		gs.Cond("@media (min-width: "+themes.Default.ScreenMDMin+")",
			MakeGrid("-md"),
		),
		// Large grid
		//
		// Columns, offsets, pushes, and pulls for the large desktop device range.
		gs.Cond("@media (min-width: "+themes.Default.ScreenLGMin+")",
			MakeGrid("-lg"),
		),
		// Extra Large grid
		//
		// Columns, offsets, pushes, and pulls for the full hd device range.
		gs.Cond("@media (min-width: "+themes.Default.ScreenXLMin+")",
			MakeGrid("-xl"),
		),
		// Extra Extra Large grid
		//
		// Columns, offsets, pushes, and pulls for the full hd device range.
		gs.Cond("@media (min-width: "+themes.Default.ScreenXXLMin+")",
			MakeGrid("-xxl"),
		),
	)
}

// Style complete grid styles for antd
func Style() gs.CSSRule {
	return gs.CSS(
		RowStyle(),
		ColStyle(),
	)
}
