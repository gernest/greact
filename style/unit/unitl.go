package unit

import (
	"regexp"
	"strings"
)

type Unit struct {
	Numerator   []string
	Denominator []string
}

func (u *Unit) ToString() string {
	if u.Numerator != nil {
		s := strings.Join(u.Numerator, "*")
		for _, v := range u.Denominator {
			s += "/" + v
		}
		return s
	}
	return ""
}

func (u *Unit) Equal(n *Unit) bool {
	return u.ToString() == n.ToString()
}

func (u *Unit) Is(n string) bool {
	return strings.ToUpper(u.ToString()) == strings.ToUpper(n)
}

var lenRexp = regexp.MustCompile(`^(px|em|ex|ch|rem|in|cm|mm|pc|pt|ex|vw|vh|vmin|vmax)$`)

func (u *Unit) IsLength() bool {
	s := u.ToCSS()
	if s != "" {
		return lenRexp.Match([]byte(s))
	}
	return false
}

func (u *Unit) ToCSS() string {
	if len(u.Numerator) == 1 {
		return u.Numerator[0]
	}
	if len(u.Denominator) > 0 {
		return u.Denominator[0]
	}
	return ""
}
