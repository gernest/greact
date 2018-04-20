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

var htmlTags = map[string]string{
	"a":          "A",
	"abbr":       "Abbr",
	"address":    "Address",
	"area":       "Area",
	"article":    "Article",
	"aside":      "Aside",
	"audio":      "Audio",
	"b":          "B",
	"base":       "Base",
	"bdi":        "Bdi",
	"bdo":        "Bdo",
	"blockquote": "Blockquote",
	"body":       "Body",
	"br":         "Br",
	"button":     "Button",
	"canvas":     "Canvas",
	"caption":    "Caption",
	"cite":       "Cite",
	"code":       "Code",
	"col":        "Col",
	"colgroup":   "Colgroup",
	"data":       "Data",
	"datalist":   "Datalist",
	"dd":         "Dd",
	"del":        "Del",
	"details":    "Details",
	"dfn":        "Dfn",
	"dialog":     "Dialog",
	"div":        "Div",
	"dl":         "Dl",
	"dt":         "Dt",
	"em":         "Em",
	"embed":      "Embed",
	"fieldset":   "Fieldset",
	"figcaption": "Figcaption",
	"figure":     "Figure",
	"footer":     "Footer",
	"form":       "Form",
	"h1":         "H1",
	"h2":         "H2",
	"h3":         "H3",
	"h4":         "H4",
	"h5":         "H5",
	"h6":         "H6",
	"head":       "Head",
	"header":     "Header",
	"hgroup":     "Hgroup",
	"hr":         "Hr",
	"html":       "Html",
	"i":          "I",
	"iframe":     "Iframe",
	"img":        "Img",
	"input":      "Input",
	"ins":        "Ins",
	"kbd":        "Kbd",
	"keygen":     "Keygen",
	"label":      "Label",
	"legend":     "Legend",
	"li":         "Li",
	"link":       "Link",
	"main":       "Main",
	"map":        "Map",
	"mark":       "Mark",
	"math":       "Math",
	"menu":       "Menu",
	"menuitem":   "Menuitem",
	"meta":       "Meta",
	"meter":      "Meter",
	"nav":        "Nav",
	"noscript":   "Noscript",
	"object":     "HTMLObject",
	"ol":         "Ol",
	"optgroup":   "Optgroup",
	"option":     "Option",
	"output":     "Output",
	"p":          "P",
	"param":      "Param",
	"picture":    "Picture",
	"pre":        "Pre",
	"progress":   "Progress",
	"q":          "Q",
	"rb":         "Rb",
	"rp":         "Rp",
	"rt":         "Rt",
	"rtc":        "Rtc",
	"ruby":       "Ruby",
	"s":          "S",
	"samp":       "Samp",
	"script":     "Script",
	"section":    "Section",
	"select":     "Select",
	"slot":       "Slot",
	"small":      "Small",
	"source":     "Source",
	"span":       "Span",
	"strong":     "Strong",
	"style":      "Style",
	"sub":        "Sub",
	"summary":    "Summary",
	"sup":        "Sup",
	"svg":        "Svg",
	"table":      "Table",
	"tbody":      "Tbody",
	"td":         "Td",
	"template":   "Template",
	"textarea":   "Textarea",
	"tfoot":      "Tfoot",
	"th":         "Th",
	"thead":      "Thead",
	"time":       "Time",
	"title":      "Title",
	"tr":         "Tr",
	"track":      "Track",
	"u":          "U",
	"ul":         "Ul",
	"var":        "Var",
	"video":      "Video",
	"wbr":        "Wbr",
}

func main() {
	b, err := ioutil.ReadFile("html-tags/html-tags.json")
	if err != nil {
		log.Fatal(err)
	}
	var v []string
	err = json.Unmarshal(b, &v)
	if err != nil {
		log.Fatal(err)
	}

	b, err = ioutil.ReadFile("html-tags/html-tags-void.json")
	if err != nil {
		log.Fatal(err)
	}
	var v2 []string
	err = json.Unmarshal(b, &v2)
	if err != nil {
		log.Fatal(err)
	}
	c := make(map[string][]string)
	c["tags"] = v
	c["void"] = v2
	s := `
package html
  import(
	  "github.com/gernest/goss"
  )
{{range .tags}}
func {{funcName .}} (c...interface{})*Object{
	return goss.WraoArgs(Selector("{{.}}"),c...)
}
{{end}}

var htmlTags=map[string]bool{
	{{range .tags -}}
	"{{.}}":true,
	{{end}}
}


`
	tpl, err := template.New("tags").Funcs(template.FuncMap{
		"camel":    camel,
		"funcName": funcName,
	}).Parse(s)
	if err != nil {
		log.Fatal(err)
	}
	var buf bytes.Buffer
	err = tpl.Execute(&buf, c)
	if err != nil {
		log.Fatal(err)
	}
	f, err := format.Source(buf.Bytes())
	if err != nil {
		log.Fatal(err)
	}
	ioutil.WriteFile("html/tags.go", f, 0600)
}

func camel(a string) string {
	if strings.HasPrefix(a, "@") {
		return inflect.Camelize(a[1:])
	}
	return inflect.Camelize(a)
}

func funcName(n string) string {
	if name, ok := htmlTags[n]; ok {
		return name
	}
	return camel(n)
}
