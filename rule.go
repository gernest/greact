package goss

import (
	"bytes"
	"errors"
	"strconv"
)

type Item interface {
	ToString(opts *Options) string
}

// RuleType css rule type
type RuleType uint

// supported rules
const (
	Simple RuleType = iota + 1
	KeyFrame
	ConditionalRule
	StyleRule
	ViewportRule
	ListRule
)

type Rule interface {
	Item
	Type() RuleType
}

type Sheet struct {
	Styles   []Style
	ClassMap ClassMap
	buf      *bytes.Buffer
	index    int64
}

// NewSheet process css and returns a new *Sheet instance ready for use.
func NewSheet(css CSS) (*Sheet, error) {
	s := &Sheet{ClassMap: make(ClassMap)}
	for k, v := range css {
		c, ok := v.(CSS)
		if !ok {
			return nil, errors.New("can't have non CSS type as value")
		}
		r, err := ParseCSS(k, c)
		if err != nil {
			return nil, err
		}
		opts := NewOpts()
		opts.ClassNamer = s.ClassName
		opts.ClassMap = s.ClassMap
		r.Selector = k
		if s.ShouldGenClass(k) {
			n := s.ClassName(k)
			s.ClassMap[k] = n
			r.Selector = n
		}
		o := ToCSS(r, opts)
		if s.buf.Len() == 0 {
			s.buf.WriteString(o)
		} else {
			s.buf.WriteRune('\n')
			s.buf.WriteString(o)
		}
	}
	return s, nil
}

// ClassName assigns a new classname for selector c
func (s *Sheet) ClassName(c string) string {
	s.index++
	return c + strconv.FormatInt(s.index, 10)
}

func (s *Sheet) String() string {
	return s.buf.String()
}

// ShouldGenClass returns true if we should generate a new classname for the
// selector.
func (s *Sheet) ShouldGenClass(c string) bool {
	return false
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
