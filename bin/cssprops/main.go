package main

import (
	"bytes"
	"encoding/json"
	"go/format"
	"io/ioutil"
	"log"
	"strings"
	"text/template"

	"github.com/gernest/inflect"
)

var cssprops = map[string]string{
	"align-content":              "AlignContent",
	"align-items":                "AlignItems",
	"align-self":                 "AlignSelf",
	"all":                        "All",
	"animation":                  "Animation",
	"animation-delay":            "AnimationDelay",
	"animation-direction":        "AnimationDirection",
	"animation-duration":         "AnimationDuration",
	"animation-fill-mode":        "AnimationFillMode",
	"animation-iteration-count":  "AnimationIterationCount",
	"animation-name":             "AnimationName",
	"animation-play-state":       "AnimationPlayState",
	"animation-timing-function":  "AnimationTimingFunction",
	"backface-visibility":        "BackfaceVisibility",
	"background":                 "Background",
	"background-attachment":      "BackgroundAttachment",
	"background-blend-mode":      "BackgroundBlendMode",
	"background-clip":            "BackgroundClip",
	"background-color":           "BackgroundColor",
	"background-image":           "BackgroundImage",
	"background-origin":          "BackgroundOrigin",
	"background-position":        "BackgroundPosition",
	"background-repeat":          "BackgroundRepeat",
	"background-size":            "BackgroundSize",
	"border":                     "Border",
	"border-bottom":              "BorderBottom",
	"border-bottom-color":        "BorderBottomColor",
	"border-bottom-left-radius":  "BorderBottomLeftRadius",
	"border-bottom-right-radius": "BorderBottomRightRadius",
	"border-bottom-style":        "BorderBottomStyle",
	"border-bottom-width":        "BorderBottomWidth",
	"border-collapse":            "BorderCollapse",
	"border-color":               "BorderColor",
	"border-image":               "BorderImage",
	"border-image-outset":        "BorderImageOutset",
	"border-image-repeat":        "BorderImageRepeat",
	"border-image-slice":         "BorderImageSlice",
	"border-image-source":        "BorderImageSource",
	"border-image-width":         "BorderImageWidth",
	"border-left":                "BorderLeft",
	"border-left-color":          "BorderLeftColor",
	"border-left-style":          "BorderLeftStyle",
	"border-left-width":          "BorderLeftWidth",
	"border-radius":              "BorderRadius",
	"border-right":               "BorderRight",
	"border-right-color":         "BorderRightColor",
	"border-right-style":         "BorderRightStyle",
	"border-right-width":         "BorderRightWidth",
	"border-spacing":             "BorderSpacing",
	"border-style":               "BorderStyle",
	"border-top":                 "BorderTop",
	"border-top-color":           "BorderTopColor",
	"border-top-left-radius":     "BorderTopLeftRadius",
	"border-top-right-radius":    "BorderTopRightRadius",
	"border-top-style":           "BorderTopStyle",
	"border-top-width":           "BorderTopWidth",
	"border-width":               "BorderWidth",
	"bottom":                     "Bottom",
	"box-shadow":                 "BoxShadow",
	"box-sizing":                 "BoxSizing",
	"caption-side":               "CaptionSide",
	"clear":                      "Clear",
	"clip":                       "Clip",
	"color":                      "Color",
	"column-count":               "ColumnCount",
	"column-fill":                "ColumnFill",
	"column-gap":                 "ColumnGap",
	"column-rule":                "ColumnRule",
	"column-rule-color":          "ColumnRuleColor",
	"column-rule-style":          "ColumnRuleStyle",
	"column-rule-width":          "ColumnRuleWidth",
	"column-span":                "ColumnSpan",
	"column-width":               "ColumnWidth",
	"columns":                    "Columns",
	"content":                    "Content",
	"counter-increment":          "CounterIncrement",
	"counter-reset":              "CounterReset",
	"cursor":                     "Cursor",
	"direction":                  "Direction",
	"display":                    "Display",
	"empty-cells":                "EmptyCells",
	"filter":                     "Filter",
	"flex":                       "Flex",
	"flex-basis":                 "FlexBasis",
	"flex-direction":             "FlexDirection",
	"flex-flow":                  "FlexFlow",
	"flex-grow":                  "FlexGrow",
	"flex-shrink":                "FlexShrink",
	"flex-wrap":                  "FlexWrap",
	"float":                      "Float",
	"font":                       "Font",
	"@font-face":                 "FontFace",
	"font-family":                "FontFamily",
	"font-size":                  "FontSize",
	"font-size-adjust":           "FontSizeAdjust",
	"font-stretch":               "FontStretch",
	"font-style":                 "FontStyle",
	"font-variant":               "FontVariant",
	"font-weight":                "FontWeight",
	"hanging-punctuation":        "HangingPunctuation",
	"height":                     "Height",
	"justify-content":            "JustifyContent",
	"@keyframes":                 "Keyframes",
	"left":                       "Left",
	"letter-spacing":             "LetterSpacing",
	"line-height":                "LineHeight",
	"list-style":                 "ListStyle",
	"list-style-image":           "ListStyleImage",
	"list-style-position":        "ListStylePosition",
	"list-style-type":            "ListStyleType",
	"margin":                     "Margin",
	"margin-bottom":              "MarginBottom",
	"margin-left":                "MarginLeft",
	"margin-right":               "MarginRight",
	"margin-top":                 "MarginTop",
	"max-height":                 "MaxHeight",
	"max-width":                  "MaxWidth",
	"@media":                     "Media",
	"min-height":                 "MinHeight",
	"min-width":                  "MinWidth",
	"nav-down":                   "NavDown",
	"nav-index":                  "NavIndex",
	"nav-left":                   "NavLeft",
	"nav-right":                  "NavRight",
	"nav-up":                     "NavUp",
	"opacity":                    "Opacity",
	"order":                      "Order",
	"outline":                    "Outline",
	"outline-color":              "OutlineColor",
	"outline-offset":             "OutlineOffset",
	"outline-style":              "OutlineStyle",
	"outline-width":              "OutlineWidth",
	"overflow":                   "Overflow",
	"overflow-x":                 "OverflowX",
	"overflow-y":                 "OverflowY",
	"padding":                    "Padding",
	"padding-bottom":             "PaddingBottom",
	"padding-left":               "PaddingLeft",
	"padding-right":              "PaddingRight",
	"padding-top":                "PaddingTop",
	"page-break-after":           "PageBreakAfter",
	"page-break-before":          "PageBreakBefore",
	"page-break-inside":          "PageBreakInside",
	"perspective":                "Perspective",
	"perspective-origin":         "PerspectiveOrigin",
	"position":                   "Position",
	"quotes":                     "Quotes",
	"resize":                     "Resize",
	"right":                      "Right",
	"tab-size":                   "TabSize",
	"table-layout":               "TableLayout",
	"text-align":                 "TextAlign",
	"text-align-last":            "TextAlignLast",
	"text-decoration":            "TextDecoration",
	"text-decoration-color":      "TextDecorationColor",
	"text-decoration-line":       "TextDecorationLine",
	"text-decoration-style":      "TextDecorationStyle",
	"text-indent":                "TextIndent",
	"text-justify":               "TextJustify",
	"text-overflow":              "TextOverflow",
	"text-shadow":                "TextShadow",
	"text-transform":             "TextTransform",
	"top":                        "Top",
	"transform":                  "Transform",
	"transform-origin":           "TransformOrigin",
	"transform-style":            "TransformStyle",
	"transition":                 "Transition",
	"transition-delay":           "TransitionDelay",
	"transition-duration":        "TransitionDuration",
	"transition-property":        "TransitionProperty",
	"transition-timing-function": "TransitionTimingFunction",
	"unicode-bidi":               "UnicodeBidi",
	"vertical-align":             "VerticalAlign",
	"visibility":                 "Visibility",
	"white-space":                "WhiteSpace",
	"width":                      "Width",
	"word-break":                 "WordBreak",
	"word-spacing":               "WordSpacing",
	"word-wrap":                  "WordWrap",
	"z-index":                    "ZIndex",
}

func main() {
	b, err := ioutil.ReadFile("css-properties/w3c-css-properties.json")
	if err != nil {
		log.Fatal(err)
	}
	var v []string
	err = json.Unmarshal(b, &v)
	if err != nil {
		log.Fatal(err)
	}
	s := `
package goss
import(
	"github.com/gernest/goss"
)
{{range .}}
func {{funcName .}}(value interface{})Styler{
	return goss.Prop("{{.}}",value)
}
{{end}}


`
	fu := template.FuncMap{
		"camel": camel,
		"funcName": func(n string) string {
			if name, ok := cssprops[n]; ok {
				return name
			}
			return camel(n)
		},
	}
	tpl, err := template.New("props").Funcs(fu).Parse(s)
	if err != nil {
		log.Fatal(err)
	}
	var buf bytes.Buffer
	err = tpl.Execute(&buf, v)
	if err != nil {
		log.Fatal(err)
	}
	f, err := format.Source(buf.Bytes())
	if err != nil {
		log.Fatal(err)
	}
	ioutil.WriteFile("css/properties.go", f, 0600)
}
func camel(a string) string {
	if strings.HasPrefix(a, "@") {
		return inflect.Camelize(a[1:])
	}
	return inflect.Camelize(a)
}
