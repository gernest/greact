package grid

import (
	"github.com/gernest/gs"
	"github.com/gernest/vected/style/grid"
	"github.com/gernest/vected/ui"
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
)

// Number an alias for int64 used to represent a grid cell.
type Number = grid.Number

// helper keys for the 24 grid cells.You can use the keyscof column
// span,offset,pull or push.
const (
	// Empty means display: none
	Empty Number = iota //display: none
	G1                  //span 1
	G2                  //span 2
	G3                  //span 3
	G4                  //span 4
	G5                  //span 5
	G6                  //span 6
	G7                  //span 7
	G8                  //span 8
	G9                  //span 9
	G10                 //span 10
	G11                 //span 11
	G12                 //span 12
	G13                 //span 13
	G14                 //span 14
	G15                 //span 15
	G16                 //span 16
	G17                 //span 17
	G18                 //span 18
	G19                 //span 19
	G20                 //span 20
	G21                 //span 21
	G22                 //span 22
	G23                 //span 123
	G24                 //span 124
)

// ColOptions options for styling a grid Column
type ColOptions = grid.ColOptions

// Column implements vecty.Component. This uses ant design language to style the
// grid column offering 24 cells grid column.
//
// Styles are bundled with the component using gs library so no additional
// dependency is required.
type Column struct {
	vecty.Core

	// Span : raster number of cells to occupy, 0 corresponds to display: none
	Span Number
	// Order: raster order, used in flex layout mode
	Order Number
	// Offset: the number of cells to offset Col from the left
	Offset Number
	// Push : the number of cells that raster is moved to the right
	Push Number
	// Pull : the number of cells that raster is moved to the left
	Pull Number

	// Gutter is the space between grid cells
	Gutter int64

	// The style will be applied in the column's dive inside vecty.Markup(). This
	// is optional.
	Style vecty.Applyer

	// Provide optional styles using the the gs library. It makes more sense to
	// define selectors (classes) which will be applied the the column's top level
	// <div> element together with the default column styles.
	//
	// The styles will be cleared after the component has been unmounted.
	CSS gs.CSSRule

	// This component will be rendered as children of the column's <div>
	//
	// This function will be called whenever this component is rendered. We don't
	// want to store the child components so that we can avoid vecty complains that
	// we stashed component address..
	Children func() vecty.MarkupOrChild

	// Media queries
	XS  *ColOptions // <576px
	SM  *ColOptions //≥576px,
	MD  *ColOptions //≥768px
	LG  *ColOptions //≥992px
	XL  *ColOptions //≥1200px
	XXL *ColOptions //≥1600px

	sheet *gs.Sheet
}

// Mount attaches the loaded stylesheets for this component.
func (c *Column) Mount() {
	c.sheet.Attach()
}

// Render adds the default ant design styles for the column and options the
// style rules provided in the CSS field. The stylesheet is not attached to the
// dom until the component is mounted.
func (c *Column) Render() vecty.ComponentOrHTML {
	if c.sheet == nil {
		c.sheet = ui.NewSheet()
		c.sheet.AddRule(c.style())
		if c.CSS != nil {
			c.sheet.AddRule(c.CSS)
		}
	}
	ch := c.getChildren()
	style := c.Style
	if c.Gutter > 0 {
		style = vecty.Markup(
			vecty.Style("padding-left", format(c.Gutter/2)+"px"),
			vecty.Style("padding-right", format(c.Gutter/2)+"px"),
		)
	}
	classes := vecty.ClassMap(c.sheet.CLasses.Classes())
	return elem.Div(vecty.Markup(classes, style), ch)
}

func (c *Column) getChildren() vecty.MarkupOrChild {
	if c.Children != nil {
		return c.Children()
	}
	return nil
}

func (c *Column) style() gs.CSSRule {
	var media []grid.MediaOption
	if c.XS != nil {
		media = append(media, grid.MediaOption{
			Type: grid.XS,
			Opts: c.XS,
		})
	}
	if c.SM != nil {
		media = append(media, grid.MediaOption{
			Type: grid.SM,
			Opts: c.SM,
		})
	}
	if c.MD != nil {
		media = append(media, grid.MediaOption{
			Type: grid.MD,
			Opts: c.MD,
		})
	}
	if c.LG != nil {
		media = append(media, grid.MediaOption{
			Type: grid.LG,
			Opts: c.LG,
		})
	}

	if c.XL != nil {
		media = append(media, grid.MediaOption{
			Type: grid.XL,
			Opts: c.XL,
		})
	}
	if c.XXL != nil {
		media = append(media, grid.MediaOption{
			Type: grid.XXL,
			Opts: c.XXL,
		})
	}
	return grid.Column(c.options(), media...)
}

func (c *Column) options() *ColOptions {
	return &ColOptions{
		Span:   c.Span,
		Order:  c.Order,
		Offset: c.Offset,
		Push:   c.Push,
		Pull:   c.Pull,
		Gutter: c.Gutter,
	}
}

func join(s ...string) string {
	o := ""
	for _, v := range s {
		o += v
	}
	return o
}

// Unmount cleanups by detaching the loaded styles for the component.
func (c *Column) Unmount() {
	c.sheet.Detach()
}
