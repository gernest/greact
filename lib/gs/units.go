package gs

import (
	"fmt"
	"strconv"
	"unicode"
)

//U wraps strings to allow basic operations on css units.
type U string

var supportedUnits = map[string]bool{
	"%":  true,
	"cm": true,
	"em": true,
	"ex": true,
	"in": true,
	"mm": true,
	"pc": true,
	"pt": true,
	"px": true,
	"vh": true,
	"vw": true,
}

// Unit returns the string representation on the css unit.
func (u U) Unit() string {
	size := len(u)
	if size == 0 {
		return ""
	}
	v := u.Measure()
	if supportedUnits[string(v)] {
		return string(v)
	}
	return ""
}

// Value returns float64 of the value in the unit. Percent units are divided by
// 100, the rest are interpreted as floats.
func (u U) Value() float64 {
	if u == "" {
		return 0
	}
	if len(u) > 2 {
		p := u.Unit()
		v, _ := strconv.ParseFloat(string(u[:len(u)-len(p)]), 64)
		if p == "%" {
			return v / 100
		}
		return v
	}
	v, _ := strconv.ParseFloat(string(u), 64)
	return v
}

func (u U) String() string {
	return string(u)
}

// Measure returns measurement of the unit.
func (u U) Measure() string {
	for k, v := range u {
		if v == '%' {
			return "%"
		}
		if unicode.IsLetter(v) {
			return string(u[k:])
		}
	}
	return ""
}

// Div performs division between two units.
func (u U) Div(n U) U {
	a, au := u.Value(), u.Unit()
	b, bu := n.Value(), n.Unit()
	switch {
	case au == bu || (au != "" && bu == ""):
		return U(format(a/b, au))
	case au == "" && bu != "":
		return U(format(a/b, bu))
	}
	return U("")
}

func format(value float64, unit string) string {
	return fmt.Sprintf("%v%s", value, unit)
}
