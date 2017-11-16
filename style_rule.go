package goss

// BaseStyleRule basic simple css property. Contains key/value .
type BaseStyleRule struct {
	Key   string
	Value string
}

// Type returns Simple rule type
func (s *BaseStyleRule) Type() RuleType {
	return StyleRule
}

// ToString returns string representation of the rule.
func (s *BaseStyleRule) ToString(opts *Options) string {
	return s.Key + ": " + s.Value + ";"
}

// Valid returns false and an error when key/value is an empty string.
func (s *BaseStyleRule) Valid() (bool, string) {
	if s.Key == "" {
		return false, "can't use empty string as key"
	}
	if s.Value == "" {
		return false, "can't use empty string as calue"
	}
	return true, ""
}
