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
func ToCSS(style *Style, opts *Options) string {
	r := ""
	if style == nil {
		return r
	}
	nested := ""
	indent := opts.Indent
	indent++
	for _, v := range style.Fallbacks {
		r += IndentStr(EndNewLine(v.ToString(opts)), indent)
	}
	for _, v := range style.Rules {
		if vt, ok := v.(*Style); ok {
			nested += EndNewLine(ToCSS(vt, opts))
		} else {
			r += IndentStr(EndNewLine(v.ToString(opts)), indent)
		}
	}
	indent--
	result := IndentStr(EndNewLine(style.Selector+" {")+r, indent) + IndentStr("}", indent)
	if nested != "" {
		return result + BeginNewLine(nested)
	}
	return result
}

func hasPrefix(str string, prefix string) bool {
	if len(prefix) > len(str) || str == "" || prefix == "" {
		return false
	}
	for i := 0; i < len(prefix); i++ {
		v := prefix[i]
		e := str[i]
		if v != e {
			return false
		}
	}
	return true
}
