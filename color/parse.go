package color

import (
	"errors"
	"math"
	"regexp"
	"strconv"
	"strings"
)

const (
	cssInt           = "[-\\+]?\\d+%?"
	cssNumber        = "[-\\+]?\\d*\\.\\d+%?"
	cssUnit          = "(?:" + cssNumber + ")|(?:" + cssInt + ")"
	permissiveMatch3 = "[\\s|\\(]+(" + cssUnit + ")[,|\\s]+(" + cssUnit + ")[,|\\s]+(" + cssUnit + ")\\s*\\)?"
	permissiveMatch4 = "[\\s|\\(]+(" + cssUnit + ")[,|\\s]+(" + cssUnit + ")[,|\\s]+(" + cssUnit + ")[,|\\s]+(" + cssUnit + ")\\s*\\)?"
)

var (
	matchUnit = regexp.MustCompile(cssUnit)
	matchRGB  = regexp.MustCompile("rgb" + permissiveMatch3)
	matchRGBA = regexp.MustCompile("rgba" + permissiveMatch4)
	matchHSL  = regexp.MustCompile("hsl" + permissiveMatch3)
	matchHSLA = regexp.MustCompile("hsla" + permissiveMatch4)
	matchHSV  = regexp.MustCompile("hsv" + permissiveMatch3)
	matchHSVA = regexp.MustCompile("hsva" + permissiveMatch4)
	matchHex3 = regexp.MustCompile("^#?([0-9a-fA-F]{1})([0-9a-fA-F]{1})([0-9a-fA-F]{1})$")
	matchHex4 = regexp.MustCompile("^#?([0-9a-fA-F]{1})([0-9a-fA-F]{1})([0-9a-fA-F]{1})([0-9a-fA-F]{1})$")
	matchHex6 = regexp.MustCompile("^#?([0-9a-fA-F]{2})([0-9a-fA-F]{2})([0-9a-fA-F]{2})$")
	matchHex8 = regexp.MustCompile("^#?([0-9a-fA-F]{2})([0-9a-fA-F]{2})([0-9a-fA-F]{2})([0-9a-fA-F]{2})$")
)

func matchColor(c string) *matchedColor {
	if n, ok := CommonColors[c]; ok {
		return matchColor(n)
	}
	return execMatchers(c,
		wrapMatcher("rgb", matchRGB),
		wrapMatcher("rgba", matchRGBA),
		wrapMatcher("hsl", matchHSL),
		wrapMatcher("hsla", matchHSLA),
		wrapMatcher("hsv", matchHSV),
		wrapMatcher("hsva", matchHSVA),
		wrapMatcher("hex8", matchHex8),
		wrapMatcher("hex6", matchHex6),
		wrapMatcher("hex4", matchHex4),
		wrapMatcher("hex3", matchHex3),
	)
}

func wrapMatcher(name string, re *regexp.Regexp) func(string) *matchedColor {
	return func(c string) *matchedColor {
		if o := re.FindAll([]byte(c), -1); o != nil {
			m := &matchedColor{name: name}
			for _, v := range o {
				m.matches = append(m.matches, string(v))
			}
			return m
		}
		return nil
	}
}

func execMatchers(src string, fn ...func(string) *matchedColor) *matchedColor {
	for _, v := range fn {
		if o := v(src); o != nil {
			return o
		}
	}
	return nil
}

type matchedColor struct {
	name    string
	matches []string
}

func (m *matchedColor) toColor() *Color {

	switch m.name {
	case "rgb":
		return parseRGB(m.matches[0])
	case "rgba":
		return nil
	case "hsl":
		return nil
	case "hslq":
		return nil
	case "hsv":
		return nil
	case "hex3":
		return parseHex3(m.matches[0])
	case "hex6":
		return parseHex6(m.matches[0])
	default:

		return nil
	}
}
func parseHex3(src string) *Color {
	if src[0] == '#' {
		src = src[1:]
	}
	x, y, z := string(src[0]), string(src[1]), string(src[2])
	n := x + x + y + y + z + z
	return RGB(parseHex(n))
}

func parseHex6(src string) *Color {
	if src[0] == '#' {
		src = src[1:]
	}
	return RGB(parseHex(src))
}

func parseRGB(src string) *Color {
	parts := breadDown("rgb", src)
	if len(parts) != 3 {
		panic(errors.New(src + "is invalid rgb string"))
	}
	a := strings.TrimSpace(parts[0])
	var red, green, blue uint8
	if a[len(a)-1] == '%' {
		r, err := precentToUint(a[:len(a)-1])
		if err != nil {
			panic(err)
		}
		red = r
	} else {
		r, err := strconv.ParseUint(a, 10, 8)
		if err != nil {
			panic(err)
		}
		red = uint8(r)
	}
	b := strings.TrimSpace(parts[1])
	if b[len(b)-1] == '%' {
		g, err := precentToUint(b[:len(a)-1])
		if err != nil {
			panic(err)
		}
		green = g
	} else {
		g, err := strconv.ParseUint(b, 10, 8)
		if err != nil {
			panic(err)
		}
		green = uint8(g)
	}
	c := strings.TrimSpace(parts[2])
	if c[len(c)-1] == '%' {
		b, err := precentToUint(c[:len(a)-1])
		if err != nil {
			panic(err)
		}
		blue = b
	} else {
		b, err := strconv.ParseUint(c, 10, 8)
		if err != nil {
			panic(err)
		}
		blue = uint8(b)
	}
	return RGB(red, green, blue)
}

func precentToUint(v string) (uint8, error) {
	a, err := strconv.ParseUint(v, 10, 8)
	if err != nil {
		return 0, err
	}
	r := float64(a) * .01 * 255
	r = math.Round(r)
	return uint8(r), nil
}

func breadDown(prefix string, src string) []string {
	if len(prefix) > len(src) {
		return nil
	}
	s := src[len(prefix):]
	s = strings.TrimSpace(s)
	if s[0] == '(' && s[len(s)-1] == ')' {
		s = s[1 : len(s)-1]
	}
	if strings.Contains(s, ",") {
		return strings.Split(s, ",")
	}
	return strings.Split(s, " ")
}
