package goss

// Item defines a css item.
type Item interface {
	ToString(opts *Options) string
}

// RuleType css rule type
type RuleType uint

// supported rules
const (
	Simple RuleType = iota + 1
	KeyFrame
	ConditionalRule
	StyleRule
	ViewportRule
	ListRule
)

// Rule defines a css rule.
type Rule interface {
	Item
	Type() RuleType
}

// Style defines a css style.
type Style struct {
	Selector  string
	Rules     []Rule
	Fallbacks []Rule
}

// ToString implements Item interface.
func (s *Style) ToString(o *Options) string {
	return ToCSS(s, o)
}

// Type implements Rule interface.
func (s *Style) Type() RuleType {
	return StyleRule
}
