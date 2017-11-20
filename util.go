package goss

import "strings"

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
	for k, v := range style.Fallbacks {
		if k == 0 {
			r = IndentStr(v.ToString(opts), indent)
		} else {
			r += "\n" + IndentStr(v.ToString(opts), indent)
		}
	}
	for _, v := range style.Rules {
		if vt, ok := v.(*Style); ok {
			if nested == "" {
				nested = ToCSS(vt, opts)
			} else {
				nested += "\n" + ToCSS(vt, opts)
			}
		} else {
			if style.Selector == "" {
				if r == "" {
					r = v.ToString(opts)
				} else {
					r += "\n" + v.ToString(opts)
				}
			} else {
				if r == "" {
					r = IndentStr(v.ToString(opts), indent)
				} else {
					r += "\n" + IndentStr(v.ToString(opts), indent)
				}
			}
		}
	}
	indent--
	result := r
	if style.Selector != "" {
		result = IndentStr(EndNewLine(style.Selector+" {")+r, indent) + IndentStr("\n}", indent)
	}
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

func replace(str string, old, new string) string {
	return strings.Replace(str, old, new, -1)
}

type match struct {
	begin, end int
}
