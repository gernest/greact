package goss

type Item interface {
	ToString(opts *Options) string
}

type Validator interface {
	Valid() (bool, string)
}

type RuleType uint

const (
	Simple RuleType = iota + 1
	KeyFrame
	ConditionalRule
	StyleRule
	ViewportRule
)

type Rule interface {
	Item
	Type() RuleType
}

type Sheet struct {
	Styles []Style
}

type Style struct {
	Selector  string
	Rules     []Rule
	Fallbacks []Rule
}

func (s *Style) ToString(o *Options) string {
	return ToCSS(s, o)
}

func (s *Style) Type() RuleType {
	return StyleRule
}
