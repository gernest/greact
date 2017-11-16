package goss

import (
	"errors"
	"strconv"
)

//go:generate go run bin/cssprops/main.go

// CSS generic css representation
type CSS map[string]interface{}

// ParseCSS takes a CSS object and transforms to Style.
func ParseCSS(css CSS) (*Style, error) {
	s := &Style{}
	for k, v := range css {
		switch {
		case k[0] == '@':
			r, err := parseSpecialRule(k, v)
			if err != nil {
				return nil, err
			}
			s.Rules = append(s.Rules, r)
		case k == "fallback":
			r, err := parseFallBack(v)
			if err != nil {
				return nil, err
			}
			s.Fallbacks = append(s.Fallbacks, r...)
		case cssprops[k]:
			r, err := parseSimpleRule(k, v)
			if err != nil {
				return nil, err
			}
			s.Rules = append(s.Rules, r)
		default:
			return nil, errors.New("Unknown format for " + k)
		}

	}
	return s, nil
}

func parseSpecialRule(key string, value interface{}) (Rule, error) {
	switch key {
	case "@viewport":
		if v, ok := value.(CSS); ok {
			s, err := ParseCSS(v)
			if err != nil {
				return nil, err
			}
			return &ViewPort{Key: key, Style: s}, nil
		}
		return nil, errors.New("@viewport accepts only CSS object")
	}
	return nil, nil
}

func parseFallBack(v interface{}) ([]Rule, error) {
	return nil, nil
}

func parseSimpleRule(key string, value interface{}) (Rule, error) {
	r := &SimpleRule{Key: key}
	v, err := toString(value)
	if err != nil {
		return nil, err
	}
	r.Value = v
	return r, nil
}

func toString(value interface{}) (string, error) {
	switch v := value.(type) {
	case string:
		return v, nil
	case int:
		return strconv.Itoa(v), nil
	case int64:
		return strconv.FormatInt(v, 10), nil
	case bool:
		return strconv.FormatBool(v), nil
	case []string:
		r := ""
		for i, item := range v {
			if i == 0 {
				r += item
			} else {
				r += ", " + item
			}
		}
		return r, nil
	default:
		return "", errors.New("Value not supported")
	}
}
