package goss

import (
	"bytes"
	"html/template"
	"sort"
	"strings"

	"github.com/gernest/classnames"
)

func IndentStr(src string, indent int) string {
	r := ""
	for i := 0; i < indent; i++ {
		r += "  "
	}
	return r + src
}

type Options struct {
	Indent     int
	ClassNamer func(string) string
	ClassMap   ClassMap
}

// NewOpts returns a new *Options instance with non nil CLassMap
func NewOpts() *Options {
	return &Options{
		ClassMap: make(ClassMap),
	}
}

// ClassMap is a map of selectors to generated classname
type ClassMap map[string]string

// Classes returns a string representation of css classes stored in this map.
func (c ClassMap) Classes() string {
	var v []interface{}
	for _, i := range c {
		v = append(v, i)
	}
	return classnames.Join(v...)
}

func (c ClassMap) Merge(cm ClassMap) {
	for k, v := range cm {
		c[k] = v
	}
}

type CSSTree struct {
	Selector string
	Parent   *CSSTree
	Children TreeList
	Text     string
}

type TreeList []*CSSTree

func (t TreeList) Len() int {
	return len(t)
}
func (t TreeList) Less(i, j int) bool {
	return t[i].Text < t[j].Text
}

func (t TreeList) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}

// ToCSS returns css string representation for style
func ToCSS(style *Style, opts *Options) string {
	v, err := FormatCSS(style, nil, opts).Print(opts)
	if err != nil {
		panic(err)
	}
	return v
}

func FormatCSS(style *Style, parent *CSSTree, opts *Options) *CSSTree {
	var fallback TreeList
	for _, v := range style.Fallbacks {
		fallback = append(fallback, &CSSTree{
			Parent: parent,
			Text:   v.ToString(opts),
		})
	}
	current := &CSSTree{
		Parent:   parent,
		Selector: style.Selector,
	}
	for _, v := range style.Rules {
		switch e := v.(type) {
		case *Style:
			current.Children = append(current.Children, FormatCSS(e, current, opts))
		default:
			current.Children = append(current.Children, &CSSTree{
				Parent: current,
				Text:   v.ToString(opts),
			})
		}

	}
	sort.Sort(fallback)
	sort.Sort(current.Children)
	current.Children = append(current.Children, fallback...)
	return current
}

func hasPrefix(str string, prefix string) bool {
	return strings.HasPrefix(str, prefix)
}

func replace(str string, old, new string) string {
	return strings.Replace(str, old, new, 1)
}

func (c *CSSTree) Print(opts *Options) (string, error) {
	src := c.print(opts)
	tpl, err := template.New("css").Parse(src)
	if err != nil {
		return "", err
	}
	var buf bytes.Buffer
	if err = tpl.Execute(&buf, opts.ClassMap); err != nil {
		return "", err
	}
	return strings.TrimSpace(buf.String()), nil
}

func (c *CSSTree) print(opts *Options) string {
	var values []string
	if c.Selector != "" {
		if len(c.Children) > 0 {
			o := c.Selector
			if opts.ClassNamer != nil {
				o = opts.ClassNamer(c.Selector)
				if opts.ClassMap != nil {
					if o != c.Selector {
						opts.ClassMap[c.Selector] = o
					}
				}
			}
			o += "{"
			for _, v := range c.Children {
				if v.Selector != "" {
					values = append(values, v.print(opts))
				} else {
					opts.Indent += 2
					o += "\n" + v.print(opts)
					opts.Indent -= 2
				}
			}
			o += "\n}"
			values = append(values, o)
		}
	} else if c.Text != "" {
		values = append(values, IndentStr(c.Text, opts.Indent))
	} else {
		for _, v := range c.Children {
			values = append(values, v.print(opts))
		}
	}
	sort.Strings(values)
	return strings.Join(values, "\n")
}

func IDNamer(c string) string {
	if strings.HasPrefix(c, "@") || c == "" || strings.Contains(c, "{{") {
		return c
	}
	return c + "-id"
}
