package goss

type Item interface {
	String() string
}

type Validator interface {
	Valid() (bool, error)
}

type RuleType uint

const (
	Simple RuleType = iota + 1
	KeyFrame
	Conditional
	StyleRule
	Viewport
)

type Rule interface {
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
