package goss

func IndentStr(src string, indent int) string {
	r := ""
	for i := 0; i < indent; i++ {
		r += "  "
	}
	return r + src
}

func BeginNewLine(str string) string {
	return "\n" + str
}

func EndNewLine(str string) string {
	return str + "\n"
}

type Options struct {
	Indent int
}

// ToCSS returns css string representation for style
func ToCSS(selector string, style *Style, opts *Options) string {
	r := ""
	if style == nil {
		return r
	}
	indent := opts.Indent
	indent++
	for _, v := range style.Fallbacks {
		r += IndentStr(EndNewLine(v.ToString(opts)), indent)
	}
	for _, v := range style.Rules {
		r += IndentStr(EndNewLine(v.ToString(opts)), indent)
	}
	indent--
	return IndentStr(EndNewLine(selector+" {")+r, indent) + IndentStr("}", indent)
}
