package goss

type ViewPort struct {
	Key   string
	Style *Style
}

func (v *ViewPort) Type() RuleType {
	return ViewportRule
}

func (v *ViewPort) Valid() (bool, string) {
	if v.Key != "@viewport" {
		return false, v.Key + " is a wrong viewport key"
	}
	return false, ""
}

func (v *ViewPort) ToString(opts *Options) string {
	return ToCSS(v.Key, v.Style, opts)
}
