package goss

// SimpleRule basic simple css property. Contains key/value .
type SimpleRule struct {
	Key   string
	Value string
}

// Type returns Simple rule type
func (s *SimpleRule) Type() RuleType {
	return Simple
}

// ToString returns string representation of the rule.
func (s *SimpleRule) ToString(opts *Options) string {
	return s.Key + ": " + s.Value + ";"
}

// Valid returns false and an error when key/value is an empty string.
func (s *SimpleRule) Valid() (bool, string) {
	if s.Key == "" {
		return false, "can't use empty string as key"
	}
	if s.Value == "" {
		return false, "can't use empty string as calue"
	}
	return true, ""
}
