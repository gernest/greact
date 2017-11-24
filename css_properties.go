package goss

var cssprops = map[string]bool{
	"align-content":              true,
	"align-items":                true,
	"align-self":                 true,
	"all":                        true,
	"animation":                  true,
	"animation-delay":            true,
	"animation-direction":        true,
	"animation-duration":         true,
	"animation-fill-mode":        true,
	"animation-iteration-count":  true,
	"animation-name":             true,
	"animation-play-state":       true,
	"animation-timing-function":  true,
	"backface-visibility":        true,
	"background":                 true,
	"background-attachment":      true,
	"background-blend-mode":      true,
	"background-clip":            true,
	"background-color":           true,
	"background-image":           true,
	"background-origin":          true,
	"background-position":        true,
	"background-repeat":          true,
	"background-size":            true,
	"border":                     true,
	"border-bottom":              true,
	"border-bottom-color":        true,
	"border-bottom-left-radius":  true,
	"border-bottom-right-radius": true,
	"border-bottom-style":        true,
	"border-bottom-width":        true,
	"border-collapse":            true,
	"border-color":               true,
	"border-image":               true,
	"border-image-outset":        true,
	"border-image-repeat":        true,
	"border-image-slice":         true,
	"border-image-source":        true,
	"border-image-width":         true,
	"border-left":                true,
	"border-left-color":          true,
	"border-left-style":          true,
	"border-left-width":          true,
	"border-radius":              true,
	"border-right":               true,
	"border-right-color":         true,
	"border-right-style":         true,
	"border-right-width":         true,
	"border-spacing":             true,
	"border-style":               true,
	"border-top":                 true,
	"border-top-color":           true,
	"border-top-left-radius":     true,
	"border-top-right-radius":    true,
	"border-top-style":           true,
	"border-top-width":           true,
	"border-width":               true,
	"bottom":                     true,
	"box-shadow":                 true,
	"box-sizing":                 true,
	"caption-side":               true,
	"clear":                      true,
	"clip":                       true,
	"color":                      true,
	"column-count":               true,
	"column-fill":                true,
	"column-gap":                 true,
	"column-rule":                true,
	"column-rule-color":          true,
	"column-rule-style":          true,
	"column-rule-width":          true,
	"column-span":                true,
	"column-width":               true,
	"columns":                    true,
	"content":                    true,
	"counter-increment":          true,
	"counter-reset":              true,
	"cursor":                     true,
	"direction":                  true,
	"display":                    true,
	"empty-cells":                true,
	"filter":                     true,
	"flex":                       true,
	"flex-basis":                 true,
	"flex-direction":             true,
	"flex-flow":                  true,
	"flex-grow":                  true,
	"flex-shrink":                true,
	"flex-wrap":                  true,
	"float":                      true,
	"font":                       true,
	"@font-face":                 true,
	"font-family":                true,
	"font-size":                  true,
	"font-size-adjust":           true,
	"font-stretch":               true,
	"font-style":                 true,
	"font-variant":               true,
	"font-weight":                true,
	"hanging-punctuation":        true,
	"height":                     true,
	"justify-content":            true,
	"@keyframes":                 true,
	"left":                       true,
	"letter-spacing":             true,
	"line-height":                true,
	"list-style":                 true,
	"list-style-image":           true,
	"list-style-position":        true,
	"list-style-type":            true,
	"margin":                     true,
	"margin-bottom":              true,
	"margin-left":                true,
	"margin-right":               true,
	"margin-top":                 true,
	"max-height":                 true,
	"max-width":                  true,
	"@media":                     true,
	"min-height":                 true,
	"min-width":                  true,
	"nav-down":                   true,
	"nav-index":                  true,
	"nav-left":                   true,
	"nav-right":                  true,
	"nav-up":                     true,
	"opacity":                    true,
	"order":                      true,
	"outline":                    true,
	"outline-color":              true,
	"outline-offset":             true,
	"outline-style":              true,
	"outline-width":              true,
	"overflow":                   true,
	"overflow-x":                 true,
	"overflow-y":                 true,
	"padding":                    true,
	"padding-bottom":             true,
	"padding-left":               true,
	"padding-right":              true,
	"padding-top":                true,
	"page-break-after":           true,
	"page-break-before":          true,
	"page-break-inside":          true,
	"perspective":                true,
	"perspective-origin":         true,
	"position":                   true,
	"quotes":                     true,
	"resize":                     true,
	"right":                      true,
	"tab-size":                   true,
	"table-layout":               true,
	"text-align":                 true,
	"text-align-last":            true,
	"text-decoration":            true,
	"text-decoration-color":      true,
	"text-decoration-line":       true,
	"text-decoration-style":      true,
	"text-indent":                true,
	"text-justify":               true,
	"text-overflow":              true,
	"text-shadow":                true,
	"text-transform":             true,
	"top":                        true,
	"transform":                  true,
	"transform-origin":           true,
	"transform-style":            true,
	"transition":                 true,
	"transition-delay":           true,
	"transition-duration":        true,
	"transition-property":        true,
	"transition-timing-function": true,
	"unicode-bidi":               true,
	"vertical-align":             true,
	"visibility":                 true,
	"white-space":                true,
	"width":                      true,
	"word-break":                 true,
	"word-spacing":               true,
	"word-wrap":                  true,
	"z-index":                    true,
}

// useful for avoiding typing strings all the time.
const (
	AlignContent             = "align-content"
	AlignItems               = "align-items"
	AlignSelf                = "align-self"
	All                      = "all"
	Animation                = "animation"
	AnimationDelay           = "animation-delay"
	AnimationDirection       = "animation-direction"
	AnimationDuration        = "animation-duration"
	AnimationFillMode        = "animation-fill-mode"
	AnimationIterationCount  = "animation-iteration-count"
	AnimationName            = "animation-name"
	AnimationPlayState       = "animation-play-state"
	AnimationTimingFunction  = "animation-timing-function"
	BackfaceVisibility       = "backface-visibility"
	Background               = "background"
	BackgroundAttachment     = "background-attachment"
	BackgroundBlendMode      = "background-blend-mode"
	BackgroundClip           = "background-clip"
	BackgroundColor          = "background-color"
	BackgroundImage          = "background-image"
	BackgroundOrigin         = "background-origin"
	BackgroundPosition       = "background-position"
	BackgroundRepeat         = "background-repeat"
	BackgroundSize           = "background-size"
	Border                   = "border"
	BorderBottom             = "border-bottom"
	BorderBottomColor        = "border-bottom-color"
	BorderBottomLeftRadius   = "border-bottom-left-radius"
	BorderBottomRightRadius  = "border-bottom-right-radius"
	BorderBottomStyle        = "border-bottom-style"
	BorderBottomWidth        = "border-bottom-width"
	BorderCollapse           = "border-collapse"
	BorderColor              = "border-color"
	BorderImage              = "border-image"
	BorderImageOutset        = "border-image-outset"
	BorderImageRepeat        = "border-image-repeat"
	BorderImageSlice         = "border-image-slice"
	BorderImageSource        = "border-image-source"
	BorderImageWidth         = "border-image-width"
	BorderLeft               = "border-left"
	BorderLeftColor          = "border-left-color"
	BorderLeftStyle          = "border-left-style"
	BorderLeftWidth          = "border-left-width"
	BorderRadius             = "border-radius"
	BorderRight              = "border-right"
	BorderRightColor         = "border-right-color"
	BorderRightStyle         = "border-right-style"
	BorderRightWidth         = "border-right-width"
	BorderSpacing            = "border-spacing"
	BorderStyle              = "border-style"
	BorderTop                = "border-top"
	BorderTopColor           = "border-top-color"
	BorderTopLeftRadius      = "border-top-left-radius"
	BorderTopRightRadius     = "border-top-right-radius"
	BorderTopStyle           = "border-top-style"
	BorderTopWidth           = "border-top-width"
	BorderWidth              = "border-width"
	Bottom                   = "bottom"
	BoxShadow                = "box-shadow"
	BoxSizing                = "box-sizing"
	CaptionSide              = "caption-side"
	Clear                    = "clear"
	Clip                     = "clip"
	Color                    = "color"
	ColumnCount              = "column-count"
	ColumnFill               = "column-fill"
	ColumnGap                = "column-gap"
	ColumnRule               = "column-rule"
	ColumnRuleColor          = "column-rule-color"
	ColumnRuleStyle          = "column-rule-style"
	ColumnRuleWidth          = "column-rule-width"
	ColumnSpan               = "column-span"
	ColumnWidth              = "column-width"
	Columns                  = "columns"
	Content                  = "content"
	CounterIncrement         = "counter-increment"
	CounterReset             = "counter-reset"
	Cursor                   = "cursor"
	Direction                = "direction"
	Display                  = "display"
	EmptyCells               = "empty-cells"
	Filter                   = "filter"
	Flex                     = "flex"
	FlexBasis                = "flex-basis"
	FlexDirection            = "flex-direction"
	FlexFlow                 = "flex-flow"
	FlexGrow                 = "flex-grow"
	FlexShrink               = "flex-shrink"
	FlexWrap                 = "flex-wrap"
	Float                    = "float"
	Font                     = "font"
	FontFace                 = "@font-face"
	FontFamily               = "font-family"
	FontSize                 = "font-size"
	FontSizeAdjust           = "font-size-adjust"
	FontStretch              = "font-stretch"
	FontStyle                = "font-style"
	FontVariant              = "font-variant"
	FontWeight               = "font-weight"
	HangingPunctuation       = "hanging-punctuation"
	Height                   = "height"
	JustifyContent           = "justify-content"
	Keyframes                = "@keyframes"
	Left                     = "left"
	LetterSpacing            = "letter-spacing"
	LineHeight               = "line-height"
	ListStyle                = "list-style"
	ListStyleImage           = "list-style-image"
	ListStylePosition        = "list-style-position"
	ListStyleType            = "list-style-type"
	Margin                   = "margin"
	MarginBottom             = "margin-bottom"
	MarginLeft               = "margin-left"
	MarginRight              = "margin-right"
	MarginTop                = "margin-top"
	MaxHeight                = "max-height"
	MaxWidth                 = "max-width"
	Media                    = "@media"
	MinHeight                = "min-height"
	MinWidth                 = "min-width"
	NavDown                  = "nav-down"
	NavIndex                 = "nav-index"
	NavLeft                  = "nav-left"
	NavRight                 = "nav-right"
	NavUp                    = "nav-up"
	Opacity                  = "opacity"
	Order                    = "order"
	Outline                  = "outline"
	OutlineColor             = "outline-color"
	OutlineOffset            = "outline-offset"
	OutlineStyle             = "outline-style"
	OutlineWidth             = "outline-width"
	Overflow                 = "overflow"
	OverflowX                = "overflow-x"
	OverflowY                = "overflow-y"
	Padding                  = "padding"
	PaddingBottom            = "padding-bottom"
	PaddingLeft              = "padding-left"
	PaddingRight             = "padding-right"
	PaddingTop               = "padding-top"
	PageBreakAfter           = "page-break-after"
	PageBreakBefore          = "page-break-before"
	PageBreakInside          = "page-break-inside"
	Perspective              = "perspective"
	PerspectiveOrigin        = "perspective-origin"
	Position                 = "position"
	Quotes                   = "quotes"
	Resize                   = "resize"
	Right                    = "right"
	TabSize                  = "tab-size"
	TableLayout              = "table-layout"
	TextAlign                = "text-align"
	TextAlignLast            = "text-align-last"
	TextDecoration           = "text-decoration"
	TextDecorationColor      = "text-decoration-color"
	TextDecorationLine       = "text-decoration-line"
	TextDecorationStyle      = "text-decoration-style"
	TextIndent               = "text-indent"
	TextJustify              = "text-justify"
	TextOverflow             = "text-overflow"
	TextShadow               = "text-shadow"
	TextTransform            = "text-transform"
	Top                      = "top"
	Transform                = "transform"
	TransformOrigin          = "transform-origin"
	TransformStyle           = "transform-style"
	Transition               = "transition"
	TransitionDelay          = "transition-delay"
	TransitionDuration       = "transition-duration"
	TransitionProperty       = "transition-property"
	TransitionTimingFunction = "transition-timing-function"
	UnicodeBidi              = "unicode-bidi"
	VerticalAlign            = "vertical-align"
	Visibility               = "visibility"
	WhiteSpace               = "white-space"
	Width                    = "width"
	WordBreak                = "word-break"
	WordSpacing              = "word-spacing"
	WordWrap                 = "word-wrap"
	ZIndex                   = "z-index"
)
