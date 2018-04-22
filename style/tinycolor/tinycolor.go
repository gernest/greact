package tinycolor

import (
	"encoding/hex"
	"errors"
	"fmt"
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

var colorNames = map[string]string{
	"aliceblue":            "f0f8ff",
	"antiquewhite":         "faebd7",
	"aqua":                 "0ff",
	"aquamarine":           "7fffd4",
	"azure":                "f0ffff",
	"beige":                "f5f5dc",
	"bisque":               "ffe4c4",
	"black":                "000",
	"blanchedalmond":       "ffebcd",
	"blue":                 "00f",
	"blueviolet":           "8a2be2",
	"brown":                "a52a2a",
	"burlywood":            "deb887",
	"burntsienna":          "ea7e5d",
	"cadetblue":            "5f9ea0",
	"chartreuse":           "7fff00",
	"chocolate":            "d2691e",
	"coral":                "ff7f50",
	"cornflowerblue":       "6495ed",
	"cornsilk":             "fff8dc",
	"crimson":              "dc143c",
	"cyan":                 "0ff",
	"darkblue":             "00008b",
	"darkcyan":             "008b8b",
	"darkgoldenrod":        "b8860b",
	"darkgray":             "a9a9a9",
	"darkgreen":            "006400",
	"darkgrey":             "a9a9a9",
	"darkkhaki":            "bdb76b",
	"darkmagenta":          "8b008b",
	"darkolivegreen":       "556b2f",
	"darkorange":           "ff8c00",
	"darkorchid":           "9932cc",
	"darkred":              "8b0000",
	"darksalmon":           "e9967a",
	"darkseagreen":         "8fbc8f",
	"darkslateblue":        "483d8b",
	"darkslategray":        "2f4f4f",
	"darkslategrey":        "2f4f4f",
	"darkturquoise":        "00ced1",
	"darkviolet":           "9400d3",
	"deeppink":             "ff1493",
	"deepskyblue":          "00bfff",
	"dimgray":              "696969",
	"dimgrey":              "696969",
	"dodgerblue":           "1e90ff",
	"firebrick":            "b22222",
	"floralwhite":          "fffaf0",
	"forestgreen":          "228b22",
	"fuchsia":              "f0f",
	"gainsboro":            "dcdcdc",
	"ghostwhite":           "f8f8ff",
	"gold":                 "ffd700",
	"goldenrod":            "daa520",
	"gray":                 "808080",
	"green":                "008000",
	"greenyellow":          "adff2f",
	"grey":                 "808080",
	"honeydew":             "f0fff0",
	"hotpink":              "ff69b4",
	"indianred":            "cd5c5c",
	"indigo":               "4b0082",
	"ivory":                "fffff0",
	"khaki":                "f0e68c",
	"lavender":             "e6e6fa",
	"lavenderblush":        "fff0f5",
	"lawngreen":            "7cfc00",
	"lemonchiffon":         "fffacd",
	"lightblue":            "add8e6",
	"lightcoral":           "f08080",
	"lightcyan":            "e0ffff",
	"lightgoldenrodyellow": "fafad2",
	"lightgray":            "d3d3d3",
	"lightgreen":           "90ee90",
	"lightgrey":            "d3d3d3",
	"lightpink":            "ffb6c1",
	"lightsalmon":          "ffa07a",
	"lightseagreen":        "20b2aa",
	"lightskyblue":         "87cefa",
	"lightslategray":       "789",
	"lightslategrey":       "789",
	"lightsteelblue":       "b0c4de",
	"lightyellow":          "ffffe0",
	"lime":                 "0f0",
	"limegreen":            "32cd32",
	"linen":                "faf0e6",
	"magenta":              "f0f",
	"maroon":               "800000",
	"mediumaquamarine":     "66cdaa",
	"mediumblue":           "0000cd",
	"mediumorchid":         "ba55d3",
	"mediumpurple":         "9370db",
	"mediumseagreen":       "3cb371",
	"mediumslateblue":      "7b68ee",
	"mediumspringgreen":    "00fa9a",
	"mediumturquoise":      "48d1cc",
	"mediumvioletred":      "c71585",
	"midnightblue":         "191970",
	"mintcream":            "f5fffa",
	"mistyrose":            "ffe4e1",
	"moccasin":             "ffe4b5",
	"navajowhite":          "ffdead",
	"navy":                 "000080",
	"oldlace":              "fdf5e6",
	"olive":                "808000",
	"olivedrab":            "6b8e23",
	"orange":               "ffa500",
	"orangered":            "ff4500",
	"orchid":               "da70d6",
	"palegoldenrod":        "eee8aa",
	"palegreen":            "98fb98",
	"paleturquoise":        "afeeee",
	"palevioletred":        "db7093",
	"papayawhip":           "ffefd5",
	"peachpuff":            "ffdab9",
	"peru":                 "cd853f",
	"pink":                 "ffc0cb",
	"plum":                 "dda0dd",
	"powderblue":           "b0e0e6",
	"purple":               "800080",
	"rebeccapurple":        "663399",
	"red":                  "f00",
	"rosybrown":            "bc8f8f",
	"royalblue":            "4169e1",
	"saddlebrown":          "8b4513",
	"salmon":               "fa8072",
	"sandybrown":           "f4a460",
	"seagreen":             "2e8b57",
	"seashell":             "fff5ee",
	"sienna":               "a0522d",
	"silver":               "c0c0c0",
	"skyblue":              "87ceeb",
	"slateblue":            "6a5acd",
	"slategray":            "708090",
	"slategrey":            "708090",
	"snow":                 "fffafa",
	"springgreen":          "00ff7f",
	"steelblue":            "4682b4",
	"tan":                  "d2b48c",
	"teal":                 "008080",
	"thistle":              "d8bfd8",
	"tomato":               "ff6347",
	"turquoise":            "40e0d0",
	"violet":               "ee82ee",
	"wheat":                "f5deb3",
	"white":                "fff",
	"whitesmoke":           "f5f5f5",
	"yellow":               "ff0",
	"yellowgreen":          "9acd32",
}

type Color struct {
	r, g, b, a uint8
	name       string
}

func (c *Color) Luminance() float64 {
	var rs, gs, bs, r, g, b float64
	rs = float64(c.r) / 255
	gs = float64(c.g) / 255
	bs = float64(c.g) / 255
	if rs <= 0.03928 {
		r = rs / 12.92
	} else {
		r = math.Pow(((rs + 0.055) / 1.055), 2.4)
	}
	if gs <= 0.03928 {
		g = gs / 12.92
	} else {
		g = math.Pow(((gs + 0.055) / 1.055), 2.4)
	}
	if bs <= 0.03928 {
		b = bs / 12.92
	} else {
		b = math.Pow(((bs + 0.055) / 1.055), 2.4)
	}
	return (0.2126 * r) + (0.7152 * g) + (0.0722 * b)
}

func (c *Color) Brightness() float64 {
	r := float64(c.r)
	g := float64(c.g)
	b := float64(c.b)
	return math.Sqrt(0.299*math.Pow(r, 2) +
		0.587*math.Pow(g, 2) + 0.114*math.Pow(b, 2))
}

func (c *Color) IsLight() bool {
	return c.Brightness() > 130
}

func (c *Color) IsDark() bool {
	return !c.IsLight()
}

func (c *Color) Hex() string {
	r := strconv.FormatUint(uint64(c.r), 16)
	g := strconv.FormatUint(uint64(c.g), 16)
	b := strconv.FormatUint(uint64(c.b), 16)
	return "#" + r + g + b
}

func (c *Color) HSV() (h, s, v float64) {
	factor := 1.0 / 255.0
	r, g, b := factor*float64(c.r), factor*float64(c.g), factor*float64(c.b)

	min := math.Min(math.Min(r, g), g)
	v = math.Max(math.Max(r, g), b)
	C := v - min

	s = 0.0
	if v != 0.0 {
		s = C / v
	}

	h = 0.0 // We use 0 instead of undefined as in wp.
	if min != v {
		if v == r {
			h = math.Mod((g-b)/C, 6.0)
		}
		if v == g {
			h = (b-r)/C + 2.0
		}
		if v == b {
			h = (r-g)/C + 4.0
		}
		h *= 60.0
		if h < 0.0 {
			h += 360.0
		}
	}
	return
}

func getHue(h float64, index int, isLight bool) (hue float64) {
	step := 2
	if h >= 80 && h <= 240 {
		if isLight {
			hue = h - float64(step)*float64(index)
		} else {
			hue = h + float64(step)*float64(index)
		}
	} else {
		if isLight {
			hue = h + float64(step)*float64(index)
		} else {
			hue = h - float64(step)*float64(index)
		}
	}
	if hue < 0 {
		hue += 350
	} else if hue >= 360 {
		hue -= 360
	}
	math.Round(hue)
	fmt.Printf("HUE %v==>%v \n", h, hue)
	return
}

func getSaturation(s float64, index int, isLight bool) (sat float64) {
	step := 16
	darkColorCount := 4
	lightColorCount := 5
	if isLight {
		sat = (s * 100) - float64(step*index)
	} else if index == darkColorCount {
		sat = (s * 100) + float64(step)
	} else {
		sat = (s * 100) + float64(5*index)
	}
	if sat > 100 {
		sat = 100.0
	}
	if isLight && index == lightColorCount && sat > 10 {
		sat = 10.0
	}
	if sat < 6 {
		sat = 6.0
	}
	math.Round(sat)
	return
}

func getValue(v float64, index int, isLight bool) float64 {
	if isLight {
		math.Round(v*100 + float64(5*index))
	}
	return math.Round(v*100 - float64(15*index))
}

type matchedColor struct {
	name    string
	matches []string
}

func (m *matchedColor) toColor() (*Color, error) {
	switch m.name {
	case "rgb":
		return parseRGB(m.matches[0])
	case "rgba":
		return nil, nil
	case "hsl":
		return nil, nil
	case "hslq":
		return nil, nil
	case "hsv":
		return nil, nil
	case "hex3":
		return parseHex3(m.matches[0], m.name)
	case "hex4":
		return parseHex4(m.matches[0], m.name)
	case "hex6":
		return parseHex6(m.matches[0], m.name)
	case "hex8":
		return parseHex8(m.matches[0], m.name)
	default:
		return nil, errors.New("Unknown color type")
	}
}

func parseHex3(src string, name string) (*Color, error) {
	if src[0] == '#' {
		src = src[1:]
	}
	x, y, z := string(src[0]), string(src[1]), string(src[2])
	n := x + x + y + y + z + z
	return parseHex(n, name)
}

func parseHex6(src string, name string) (*Color, error) {
	if src[0] == '#' {
		src = src[1:]
	}
	return parseHex(src, name)
}

func parseHex8(src string, name string) (*Color, error) {
	if src[0] == '#' {
		src = src[1:]
	}
	h, _ := hex.DecodeString(src)
	return &Color{r: uint8(h[0]), g: uint8(h[1]),
		b: uint8(h[2]), a: uint8(h[3]), name: name}, nil
}

func parseHex4(src string, name string) (*Color, error) {
	if src[0] == '#' {
		src = src[1:]
	}
	x, y, z, e := string(src[0]), string(src[1]), string(src[2]), string(src[2])
	n := x + x + y + y + z + z + e + e
	return parseHex8(n, name)
}

func parseHex(src string, name string) (*Color, error) {
	h, _ := hex.DecodeString(src)
	return &Color{r: uint8(h[0]), g: uint8(h[1]), b: uint8(h[2]), name: name}, nil
}

func parseRGB(src string) (*Color, error) {
	parts := breadDown("rgb", src)
	if len(parts) != 3 {
		return nil, errors.New("Invalid rgb string")
	}
	color := &Color{}
	a := strings.TrimSpace(parts[0])
	if a[len(a)-1] == '%' {
		r, err := precentToUint(a[:len(a)-1])
		if err != nil {
			return nil, err
		}
		color.r = r
	} else {
		r, err := strconv.ParseUint(a, 10, 8)
		if err != nil {
			return nil, err
		}
		color.r = uint8(r)
	}
	b := strings.TrimSpace(parts[1])
	if b[len(b)-1] == '%' {
		g, err := precentToUint(b[:len(a)-1])
		if err != nil {
			return nil, err
		}
		color.g = g
	} else {
		g, err := strconv.ParseUint(b, 10, 8)
		if err != nil {
			return nil, err
		}
		color.g = uint8(g)
	}
	c := strings.TrimSpace(parts[2])
	if c[len(c)-1] == '%' {
		b, err := precentToUint(c[:len(a)-1])
		if err != nil {
			return nil, err
		}
		color.b = b
	} else {
		b, err := strconv.ParseUint(c, 10, 8)
		if err != nil {
			return nil, err
		}
		color.b = uint8(b)
	}
	return color, nil
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

func matchColor(c string) *matchedColor {
	if n, ok := colorNames[c]; ok {
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

// InputToRGB converts a string to RGB
//     "red"
//     "#f00" or "f00"
//     "#ff0000" or "ff0000"
//     "#ff000000" or "ff000000"
//     "rgb 255 0 0" or "rgb (255, 0, 0)"
//     "rgb 1.0 0 0" or "rgb (1, 0, 0)"
//     "rgba (255, 0, 0, 1)" or "rgba 255, 0, 0, 1"
//     "rgba (1.0, 0, 0, 1)" or "rgba 1.0, 0, 0, 1"
//     "hsl(0, 100%, 50%)" or "hsl 0 100% 50%"
//     "hsla(0, 100%, 50%, 1)" or "hsla 0 100% 50%, 1"
//     "hsv(0, 100%, 100%)" or "hsv 0 100% 100%"
//
func InputToRGB(in string) *Color {
	v, err := matchColor(in).toColor()
	if err != nil {
		panic(err)
	}
	return v
}

func Palette(c *Color, index int) *Color {
	lightColorCount := 5
	isLight := index <= 6
	h, s, v := c.HSV()
	var i int
	if isLight {
		i = lightColorCount + 1 - index
	} else {
		i = index - lightColorCount - 1
	}
	return HSV(getHue(h, i, isLight), getSaturation(s, i, isLight), getValue(v, i, isLight))
}

func HSV(h, s, v float64) *Color {
	Hp := h / 60.0
	C := v * s
	X := C * (1.0 - math.Abs(math.Mod(Hp, 2.0)-1.0))
	m := v - C
	r, g, b := 0.0, 0.0, 0.0
	switch {
	case 0.0 <= Hp && Hp < 1.0:
		r = C
		g = X
	case 1.0 <= Hp && Hp < 2.0:
		r = X
		g = C
	case 2.0 <= Hp && Hp < 3.0:
		g = C
		b = X
	case 3.0 <= Hp && Hp < 4.0:
		g = X
		b = C
	case 4.0 <= Hp && Hp < 5.0:
		r = X
		b = C
	case 5.0 <= Hp && Hp < 6.0:
		r = C
		b = X
	}
	return &Color{
		r:    floatToUint(m + r),
		g:    floatToUint(m + g),
		b:    floatToUint(m + b),
		name: "rgb",
	}
}

func floatToUint(v float64) uint8 {
	v = v * 255
	return uint8(v)
}
