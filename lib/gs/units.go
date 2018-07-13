package gs

import "strconv"

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
	if u[size-1] == '%' {
		return "%"
	}
	if len(u) > 2 {
		v := u[size-2:]
		if supportedUnits[string(v)] {
			return string(v)
		}
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
