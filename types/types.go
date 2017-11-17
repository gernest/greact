package types

// CSS generic css representation
type CSS map[string]interface{}

// RuleType types of supported css rules.
type RuleType uint32

// css rules types
const (
	StyleRule RuleType = iota + 1
	CharsetRule
	ImportRule
	MediaRule
	FontFaceRule
	PageRule
	KeyFramesRule
	KeyFrameRule
	NamespaceRule
	CounterStyleRule
	SupportsRule
	DocumentRule
	FontFeatureValuesRule
	ViewportRule
	RegionStyleRule
)

type CSSRuleBase struct {
	CSSText DOMString
	Type    RuleType
}

func (b CSSRuleBase) GetType() RuleType {
	return b.Type
}

type DOMString string

type CSSStyleRule struct {
	CSSRuleBase
	Selector DOMString
	Text     DOMString
	Style    CSS
}

type CSSRuleList []CSSStyleRule

type CSSKeyframeRule struct {
	CSSRuleBase
	Style   CSS
	KeyText DOMString
}

type CSSKeyframesRule struct {
	CSSRuleBase
	CSSRules CSSRuleList
	Name     DOMString
}

func (f *CSSKeyframesRule) Append(rule DOMString) error {
	return nil
}

func (f *CSSKeyframesRule) Delete(rule DOMString) error {
	return nil
}

func (f *CSSKeyframesRule) Find(rule DOMString) (*CSSKeyframeRule, error) {
	return nil, nil
}

type CSSMediaRule struct {
	MediaText DOMString
	Length    int64
	Item      DOMString
}

func (f *CSSMediaRule) Append(medium DOMString) error {
	return nil
}

func (f *CSSMediaRule) Delete(medium DOMString) error {
	return nil
}

type Options struct{}

type CSSOMRule interface {
	ToCSS(*Options) DOMString
	GetType() RuleType
}

type ClassNameGen func(CSSOMRule) string

type Class string

type RuleOptions struct {
	Selector     string
	Index        int64
	Classes      []Class
	ClassNameGen ClassNameGen
}
