package goss

import (
	"errors"
	"strconv"
)

//go:generate go run bin/cssprops/main.go
//go:generate go run bin/tags/main.go

// CSS generic css representation
type CSS map[string]interface{}

// ParseCSS takes a CSS object and transforms to Style.
func ParseCSS(parent string, css CSS) (*Style, error) {
	s := &Style{
		Selector: parent,
	}
	for k, v := range css {
		switch {
		case k[0] == '@':
			r, err := parseSpecialRule(parent, k, v)
			if err != nil {
				return nil, err
			}
			s.Rules = append(s.Rules, r)
		case k == "fallbacks":
			r, err := parseFallBack(parent, v)
			if err != nil {
				return nil, err
			}
			s.Fallbacks = append(s.Fallbacks, r...)
		case cssprops[k]:
			r, err := parseBaseStyleRule(parent, k, v)
			if err != nil {
				return nil, err
			}
			s.Rules = append(s.Rules, r)
		default:
			if vt, ok := v.(CSS); ok {
				vk := k
				vk = replace(vk, "&", parent)
				r, err := ParseCSS(vk, vt)
				if err != nil {
					return nil, err
				}
				r.Selector = vk
				s.Rules = append(s.Rules, r)
			} else {
				return nil, errors.New("Unknown format for " + k)
			}
		}
	}
	return s, nil
}

func parseSpecialRule(parent, key string, value interface{}) (Rule, error) {
	switch key {
	case "@viewport":
		if v, ok := value.(CSS); ok {
			s, err := ParseCSS(parent, v)
			if err != nil {
				return nil, err
			}
			return &ViewPort{Key: key, Style: s}, nil
		}
		return nil, errors.New("@viewport accepts only CSS object")
	case "@import":
		switch v := value.(type) {
		case []string:
			var li RuleList
			for _, item := range v {
				li = append(li, &SimpleRule{
					Key:   key,
					Value: item,
				})
			}
			return li, nil
		default:
			return nil, errors.New("Unknown value type for @import")
		}
	}
	return nil, nil
}

type RuleList []Rule

func (r RuleList) Type() RuleType {
	return ListRule
}

func (r RuleList) ToString(opts *Options) string {
	o := ""
	for k, v := range r {
		if k == 0 {
			o += v.ToString(opts)
		} else {
			o += "\n" + v.ToString(opts)
		}
	}
	return o
}

func parseFallBack(parent string, v interface{}) ([]Rule, error) {
	switch t := v.(type) {
	case CSS:
		s, err := ParseCSS(parent, t)
		if err != nil {
			return nil, err
		}
		return s.Rules, nil
	case []CSS:
		var o []Rule
		for _, c := range t {
			r, err := ParseCSS(parent, c)
			if err != nil {
				return nil, err
			}
			o = append(o, r.Rules...)
		}
		return o, nil
	default:
		return nil, errors.New("unsporrted fallback")
	}
}

func parseBaseStyleRule(parent, key string, value interface{}) (Rule, error) {
	r := &BaseStyleRule{Key: key}
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
		r := join(", ", v)
		return r, nil
	case [][]string:
		var i []string
		for _, item := range v {
			i = append(i, join(" ", item))
		}
		r := join(", ", i)
		return r, nil
	default:
		return "", errors.New("Value not supported")
	}
}

func join(sep string, v []string) string {
	r := ""
	for i, item := range v {
		if i == 0 {
			r += item
		} else {
			r += sep + item
		}
	}
	return r
}
