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
	Validator
	Type() RuleType
}

type Sheet struct {
	Styles []Style
}

type Style struct {
	Rules     []Rule
	Fallbacks []Rule
}
