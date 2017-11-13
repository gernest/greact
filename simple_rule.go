package goss

import "errors"

type SimpleRule struct {
	Key   string
	Value string
}

func (s *SimpleRule) Type() RuleType {
	return Simple
}
func (s *SimpleRule) STring() string {
	return s.Key + " " + s.Value
}

func (s *SimpleRule) Valid() (bool, error) {
	if s.Key == "" {
		return false, errors.New("can't use empty string as key")
	}
	if s.Value == "" {
		return false, errors.New("can't use empty string as calue")
	}
	return true, nil
}
