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

## API


### Row

| Property | Description | Type | Default |
| -------- | ----------- | ---- | ------- |
| align | the vertical alignment of the flex layout: `top` `middle` `bottom` | string | `top` |
| gutter | spacing between grids, could be a number or a object like `{ xs: 8, sm: 16, md: 24}` | number/object | 0 |
| justify | horizontal arrangement of the flex layout: `start` `end` `center` `space-around` `space-between` | string | `start` |
| type | layout mode, optional `flex`, [browser support](http://caniuse.com/#search=flex) | string |  |

### Col

| Property | Description | Type | Default |
| -------- | ----------- | ---- | ------- |
| offset | the number of cells to offset Col from the left | number | 0 |
| order | raster order, used in `flex` layout mode | number | 0 |
| pull | the number of cells that raster is moved to the left | number | 0 |
| push | the number of cells that raster is moved to the right | number | 0 |
| span | raster number of cells to occupy, 0 corresponds to `display: none` | number | none |
| xs | `<576px` and also default setting, could be a `span` value or an object containing above props | number\|object | - |
| sm | `≥576px`, could be a `span` value or an object containing above props | number\|object | - |
| md | `≥768px`, could be a `span` value or an object containing above props | number\|object | - |
| lg | `≥992px`, could be a `span` value or an object containing above props | number\|object | - |
| xl | `≥1200px`, could be a `span` value or an object containing above props | number\|object | - |
| xxl | `≥1600px`, could be a `span` value or an object containing above props | number\|object | - |

The breakpoints of responsive grid follow [BootStrap 4 media queries rules](https://getbootstrap.com/docs/4.0/layout/overview/#responsive-breakpoints)(not including `occasionally part`).