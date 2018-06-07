package grid

import (
	"github.com/gernest/gs"
	"github.com/gopherjs/vecty"
)

// FlexStyle horizontal arrangement of the flex layout: start end center
// space-around space-between
type FlexStyle int64

const (
	// Start sets justify-content:flex-start
	Start FlexStyle = iota

	// End sets justify-content:flex-end
	End

	// Center sets justify-content:center
	Center

	// SpaceAround sets justify-content:space-around
	SpaceAround

	// SpaceBetween sets justify-content:space-between
	SpaceBetween
)

// FlexAlign the vertical alignment of the flex layout: top middle bottom
type FlexAlign int64

const (
	// Top sets align-items: lex-start
	Top FlexAlign = iota

	// Middle sets align-items: center
	Middle

	// Bottom sets align-items: lex-end
	Bottom
)

// Row is a vecty component using ant design to render a flex row grid layout.
// This component uses gs library to bundle the styles with the component so no
// need for external css.
type Row struct {
	vecty.Core

	// This will be  appkied to the row's <div>
	Style vecty.Applyer

	// Optional styles to be attached to this component's style sheet. Forbetter
	// results supply class selectors gs.S()
	CSS gs.CSSRule

	// Children is a function which returns components to be rendered inside to
	// row.
	// For consistency and better results this should return vecty.List of *Column
	// component. It is fine to mix Columns with other components and they will be
	// rendered correctly.
	Children func() vecty.MarkupOrChild

	// This is spacing between grids
	Gutter int64

	// Flex uses flex layout when this field is set to true. You can use Justify
	// and Align to control the layout.
	//
	// Default is false.
	Flex bool

	// Justify is horizontal arrangement of the flex layout: tart end center
	// space-around space-between
	//
	// Default is Start
	Justify FlexStyle

	// Align is the vertical alignment of the flex layout: top middle bottom
	//
	// Default is top
	Align FlexAlign
	sheet *gs.Sheet
}

// Render implements vecty.Component interface.
//
// If Gutter >0 margins are computed based on the gutter size and styles applied
// directly on the row's div. In case childerns are of type *Column then the new
// gutter size is applied before rendering of the children's.
func (r *Row) Render() vecty.ComponentOrHTML {
	return nil
}
