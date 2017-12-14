package unit

import (
	"strconv"
	"unicode/utf8"
)

type Unit interface{}

func Px(u Unit) string {
	if s, ok := u.(string); ok {
		if !isPx(s) {
			panic("trying to convert non px string unit: " + s)
		}
		return s
	}
	return toString(u) + "px"
}

func Rem(u Unit) string {
	if s, ok := u.(string); ok {
		if !isPx(s) {
			panic("trying to convert non rem string unit: " + s)
		}
		return s
	}
	return toString(u) + "rem"
}

func toString(v interface{}) string {
	switch t := v.(type) {
	case string:
		return t
	case int:
		return FormatInt(t)
	case int64:
		return FormatInt64(t)
	default:
		panic("Unkown unit value")
	}
}

func isPx(s string) bool {
	if len(s) < 3 {
		return false
	}
	x, size := utf8.DecodeLastRuneInString(s)
	if x != 'x' {
		return false
	}
	p, _ := utf8.DecodeLastRuneInString(s[:len(s)-size])
	return p == 'p'
}

func isRem(s string) bool {
	if len(s) < 4 {
		return false
	}
	m, size := utf8.DecodeLastRuneInString(s)
	if m != 'm' {
		return false
	}
	e, size := utf8.DecodeLastRuneInString(s[:len(s)-size])
	if e != 'e' {
		return false
	}
	r, _ := utf8.DecodeLastRuneInString(s[:len(s)-size])
	return r == 'r'
}

func Em(u Unit) string {
	if f, ok := u.(float64); ok {
		FormatFloat(f)
	}
	return toString(u) + "em"
}

func Format(u Unit) string {
	if s, ok := u.(string); ok {
		return s
	}
	return Px(u)
}

func FormatInt(v int) string {
	return FormatInt64(int64(v))
}

func FormatInt64(v int64) string {
	return strconv.FormatInt(v, 10)
}

func FormatFloat(f float64) string {
	return strconv.FormatFloat(f, 'f', 2, 64)
}
