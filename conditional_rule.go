package goss

type Conditional struct {
	Key   string
	Rules *Style
}

func (Conditional) Type() RuleType {
	return ConditionalRule
}

func (c Conditional) ToString(opts *Options) string {
	opts.Indent++
	inner := c.Rules.ToString(opts)
	if inner != "" {
		return c.Key + "{" + inner + "\n}"
	}
	return ""
}
