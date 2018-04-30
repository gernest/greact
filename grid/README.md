---
category: Components
type: Layout
cols: 1
title: Grid
---

24 Grids System。

## Design concept

![](media/grid1.png)
In most business situations, Ant Design needs to solve a lot of information
storage problems within the design area, so based on 12 Grids System, we
divided the design area into 24 aliquots.

We name the divided area 'box'. We suggest four boxes for horizontal
arrangement at most, one at least. Boxes are proportional to the entire screen
as shown in the picture above. To ensure a high level of visual comfort, we
customize the typography inside of the box based on the box unit.

## Outline

In the grid system, we define the frame outside the information area based on `row` and `column`, to ensure that every area can have stable arrangement.

Rows are defined by the `Row` component and columns are defined by `Column`
component. These components are standard vecty components which implements
`vecty.Component` interface.
Following is a brief look at how it works:

- Establish a set of `column` in the horizontal space defined by `row` that is `Column` component instance.
- Your content elements should be placed directly in the `Column` component's
 `Children` field. We use `func()vecty.MarkupOrChild` to avoid keeping instances
 of children's when rendering , and only `Column` components should be placed
 directly in `Row.Children`. For a list of Columns use `vecty.List`( You will see in the examples below).
- The column grid system is a value of 1-24 to represent its range spans. For example, three columns of equal width can be created by `Span=8`.
- If the sum of `Column` spans in a `Row` are more than 24, then the overflowing `Column` as a whole will start a new line arrangement.

## Flex layout

Our grid systems support Flex layout to allow the elements within the parent to be aligned horizontally - left, center, right, wide arrangement, and decentralized arrangement. The Grid system also supports vertical alignment - top aligned, vertically centered, bottom-aligned. You can also define the order of elements by using `order`.

Flex layout uses a 24 grid layout to define the width of each "box", but does not rigidly adhere to the grid layout.

# Examples

## Basic grid

 We want to build a grid that looks like this using vected grid components.
![](media/grid1.png)

vected is just a component library built on top of vecty. This means it doesn't require additional concepts or tooling to use, just plug and play. In this first example I will show you how to achieve the grid layout from the figure above using vected.

Let's start by defining out component. We will call the component `BasicGrid`.
have created a directory named `demo` where I created a subdirectory components  file `basic_grid.go`

`demo/components/basic_grid.go`
```go
package components

import (
	"github.com/gopherjs/vecty"
)

type BasicGrid struct {
	vecty.Core
}
```

Pretty standard vecty component. We don't plan to keep any states so there is no need for fields.

Let's go ahead and add a dummy `Render` method. The goal is to satisfy the `vecty.Component` interface.


```diff
diff --git a/grid/demo/components/basic_grid.go b/grid/demo/components/basic_grid.go
index c1df593..b5ab13d 100644
--- a/grid/demo/components/basic_grid.go
+++ b/grid/demo/components/basic_grid.go
@@ -7,3 +7,7 @@ import (
 type BasicGrid struct {
        vecty.Core
 }
+
+func (BasicGrid) Render() vecty.ComponentOrHTML {
+       return nil
+}
```

So from the diagram the structure of the expected layout should be something like this

```html
    <Row>
      <Col span={12}>col-12</Col>
      <Col span={12}>col-12</Col>
    </Row>
    <Row>
      <Col span={8}>col-8</Col>
      <Col span={8}>col-8</Col>
      <Col span={8}>col-8</Col>
    </Row>
    <Row>
      <Col span={6}>col-6</Col>
      <Col span={6}>col-6</Col>
      <Col span={6}>col-6</Col>
      <Col span={6}>col-6</Col>
    </Row>
  </div>,
```

Here is the equivalent code in vected

```diff
diff --git a/grid/demo/components/basic_grid.go b/grid/demo/components/basic_grid.go
index b5ab13d..361da3f 100644
--- a/grid/demo/components/basic_grid.go
+++ b/grid/demo/components/basic_grid.go
@@ -1,13 +1,54 @@
 package components
 
 import (
+	"github.com/gernest/vected/grid"
 	"github.com/gopherjs/vecty"
+	"github.com/gopherjs/vecty/elem"
 )
 
 type BasicGrid struct {
 	vecty.Core
 }
 
+func BasicGridRow(span grid.Number) vecty.Component {
+	return &grid.Column{
+		Span: span,
+		Children: func() vecty.MarkupOrChild {
+			return vecty.Text(span.String())
+		},
+	}
+}
 func (BasicGrid) Render() vecty.ComponentOrHTML {
-	return nil
+	return elem.Div(
+		&grid.Row{
+			CSS: codeBoxDemo(),
+			Children: func() vecty.MarkupOrChild {
+				return vecty.List{
+					BasicGridRow(grid.G12),
+					BasicGridRow(grid.G12),
+				}
+			},
+		},
+		&grid.Row{
+			CSS: codeBoxDemo(),
+			Children: func() vecty.MarkupOrChild {
+				return vecty.List{
+					BasicGridRow(grid.G8),
+					BasicGridRow(grid.G8),
+					BasicGridRow(grid.G8),
+				}
+			},
+		},
+		&grid.Row{
+			CSS: codeBoxDemo(),
+			Children: func() vecty.MarkupOrChild {
+				return vecty.List{
+					BasicGridRow(grid.G6),
+					BasicGridRow(grid.G6),
+					BasicGridRow(grid.G6),
+					BasicGridRow(grid.G6),
+				}
+			},
+		},
+	)
 }
```

 `BasicGridRow` is a helper function that returns a new `*Column` component. We
are passing span argument which tells the component how many cells of the grid
the component will occupy.

Now rendering this component will give us something like on this image
![](media/basic_grid_1.png)

Promising results. So, what we want now is to add some styles to our columns. Like height,padding etc for visibility. Thanks to vecty, components can be composed with functions.

 `vected` uses `gs` package for styling see [github.com/gernest/gs]()
internally. You can use this if you like since both `Column` and `Row`
components have `CSS` field, which you can use to pass `gs.CSSRule` that will
be applied to the component when mounted. However this is not mendatory, you
can optionally use `Style` field to pass `vecty.Applyer` that will be applied
to the component's top `<div>`

```diff
diff --git a/grid/demo/components/basic_grid.go b/grid/demo/components/basic_grid.go
index 361da3f..8e9fade 100644
--- a/grid/demo/components/basic_grid.go
+++ b/grid/demo/components/basic_grid.go
@@ -1,6 +1,7 @@
 package components
 
 import (
+	"github.com/gernest/gs"
 	"github.com/gernest/vected/grid"
 	"github.com/gopherjs/vecty"
 	"github.com/gopherjs/vecty/elem"
@@ -52,3 +53,17 @@ func (BasicGrid) Render() vecty.ComponentOrHTML {
 		},
 	)
 }
+
+func styleBasic() gs.CSSRule {
+	return gs.S(".BasicStyle",
+		gs.P("color", "#fff"),
+		gs.P("background", "#00a0e9"),
+		gs.P("text-align", "center"),
+		gs.P("padding", "30px 0"),
+		gs.P("font-size", "18px"),
+		gs.P("border", "none"),
+		gs.P("margin-top", "8px"),
+		gs.P("margin-bottom", "8px"),
+		gs.P("height", "15px"),
+	)
+}
```

Here I am using gs to define a class `BasicStyle` which ssets text color to white, background color and a bunch of css properties. We will then apply this to the Column components to give them good looks.

Now we apply the style when creating the Column.
```diff
diff --git a/grid/demo/components/basic_grid.go b/grid/demo/components/basic_grid.go
index 8e9fade..5d561ee 100644
--- a/grid/demo/components/basic_grid.go
+++ b/grid/demo/components/basic_grid.go
@@ -14,6 +14,7 @@ type BasicGrid struct {
 func BasicGridRow(span grid.Number) vecty.Component {
 	return &grid.Column{
 		Span: span,
+		CSS:  styleBasic(),
 		Children: func() vecty.MarkupOrChild {
 			return vecty.Text(span.String())
 		},

```

This is how it looks like now.
![](media/basic_grid_2.png)

Much better , but there is one thing a miss, we can't distinguish visualy the columns because they both have same color.

We can resolve this by applying different backgrounds on different columns.

```diff
diff --git a/grid/demo/components/basic_grid.go b/grid/demo/components/basic_grid.go
index 8e9fade..2d2a51a 100644
--- a/grid/demo/components/basic_grid.go
+++ b/grid/demo/components/basic_grid.go
@@ -11,22 +11,24 @@ type BasicGrid struct {
 	vecty.Core
 }
 
-func BasicGridRow(span grid.Number) vecty.Component {
+func BasicGridRow(span grid.Number, fn func() gs.CSSRule) vecty.Component {
 	return &grid.Column{
 		Span: span,
+		CSS:  fn(),
 		Children: func() vecty.MarkupOrChild {
 			return vecty.Text(span.String())
 		},
 	}
 }
 func (BasicGrid) Render() vecty.ComponentOrHTML {
+	style := styleBasic()
 	return elem.Div(
 		&grid.Row{
 			CSS: codeBoxDemo(),
 			Children: func() vecty.MarkupOrChild {
 				return vecty.List{
-					BasicGridRow(grid.G12),
-					BasicGridRow(grid.G12),
+					BasicGridRow(grid.G12, style),
+					BasicGridRow(grid.G12, style),
 				}
 			},
 		},
@@ -34,9 +36,9 @@ func (BasicGrid) Render() vecty.ComponentOrHTML {
 			CSS: codeBoxDemo(),
 			Children: func() vecty.MarkupOrChild {
 				return vecty.List{
-					BasicGridRow(grid.G8),
-					BasicGridRow(grid.G8),
-					BasicGridRow(grid.G8),
+					BasicGridRow(grid.G8, style),
+					BasicGridRow(grid.G8, style),
+					BasicGridRow(grid.G8, style),
 				}
 			},
 		},
@@ -44,26 +46,37 @@ func (BasicGrid) Render() vecty.ComponentOrHTML {
 			CSS: codeBoxDemo(),
 			Children: func() vecty.MarkupOrChild {
 				return vecty.List{
-					BasicGridRow(grid.G6),
-					BasicGridRow(grid.G6),
-					BasicGridRow(grid.G6),
-					BasicGridRow(grid.G6),
+					BasicGridRow(grid.G6, style),
+					BasicGridRow(grid.G6, style),
+					BasicGridRow(grid.G6, style),
+					BasicGridRow(grid.G6, style),
 				}
 			},
 		},
 	)
 }
 
-func styleBasic() gs.CSSRule {
-	return gs.S(".BasicStyle",
-		gs.P("color", "#fff"),
-		gs.P("background", "#00a0e9"),
-		gs.P("text-align", "center"),
-		gs.P("padding", "30px 0"),
-		gs.P("font-size", "18px"),
-		gs.P("border", "none"),
-		gs.P("margin-top", "8px"),
-		gs.P("margin-bottom", "8px"),
-		gs.P("height", "15px"),
-	)
+func styleBasic() func() gs.CSSRule {
+	on := false
+	bg1 := gs.P("background", "rgba(0,160,233,.7)")
+	bg2 := gs.P("background", "#00a0e9")
+	return func() gs.CSSRule {
+		bg := bg1
+		if on {
+			bg = bg2
+		}
+		on = !on
+		return gs.S(".BasicStyle",
+			gs.P("color", "#fff"),
+			bg,
+			gs.P("text-align", "center"),
+			gs.P("padding", "30px 0"),
+			gs.P("font-size", "18px"),
+			gs.P("border", "none"),
+			gs.P("margin-top", "8px"),
+			gs.P("margin-bottom", "8px"),
+			gs.P("height", "15px"),
+		)
+	}
+
 }
```

Which now gives us our final desired output.

![](media/grid1.png)

Final code snippet

```go
package components

import (
	"github.com/gernest/gs"
	"github.com/gernest/vected/grid"
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
)

type BasicGrid struct {
	vecty.Core
}

func BasicGridRow(span grid.Number, fn func() gs.CSSRule) vecty.Component {
	return &grid.Column{
		Span: span,
		CSS:  fn(),
		Children: func() vecty.MarkupOrChild {
			return vecty.Text(span.String())
		},
	}
}
func (BasicGrid) Render() vecty.ComponentOrHTML {
	style := styleBasic()
	return elem.Div(
		&grid.Row{
			CSS: codeBoxDemo(),
			Children: func() vecty.MarkupOrChild {
				return vecty.List{
					BasicGridRow(grid.G12, style),
					BasicGridRow(grid.G12, style),
				}
			},
		},
		&grid.Row{
			CSS: codeBoxDemo(),
			Children: func() vecty.MarkupOrChild {
				return vecty.List{
					BasicGridRow(grid.G8, style),
					BasicGridRow(grid.G8, style),
					BasicGridRow(grid.G8, style),
				}
			},
		},
		&grid.Row{
			CSS: codeBoxDemo(),
			Children: func() vecty.MarkupOrChild {
				return vecty.List{
					BasicGridRow(grid.G6, style),
					BasicGridRow(grid.G6, style),
					BasicGridRow(grid.G6, style),
					BasicGridRow(grid.G6, style),
				}
			},
		},
	)
}

func styleBasic() func() gs.CSSRule {
	on := false
	bg1 := gs.P("background", "rgba(0,160,233,.7)")
	bg2 := gs.P("background", "#00a0e9")
	return func() gs.CSSRule {
		bg := bg1
		if on {
			bg = bg2
		}
		on = !on
		return gs.S(".BasicStyle",
			gs.P("color", "#fff"),
			bg,
			gs.P("text-align", "center"),
			gs.P("padding", "30px 0"),
			gs.P("font-size", "18px"),
			gs.P("border", "none"),
			gs.P("margin-top", "8px"),
			gs.P("margin-bottom", "8px"),
			gs.P("height", "15px"),
		)
	}

}
