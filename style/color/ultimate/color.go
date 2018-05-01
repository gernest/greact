package ultimate

import (
	"encoding/hex"
	"errors"
	"math"
	"strconv"
)

var commonColors = map[string]string{
	"aliceblue":            "#f0f8ff",
	"antiquewhite":         "#faebd7",
	"aqua":                 "#00ffff",
	"aquamarine":           "#7fffd4",
	"azure":                "#f0ffff",
	"beige":                "#f5f5dc",
	"bisque":               "#ffe4c4",
	"black":                "#000000",
	"blanchedalmond":       "#ffebcd",
	"blue":                 "#0000ff",
	"blueviolet":           "#8a2be2",
	"brown":                "#a52a2a",
	"burlywood":            "#deb887",
	"cadetblue":            "#5f9ea0",
	"chartreuse":           "#7fff00",
	"chocolate":            "#d2691e",
	"coral":                "#ff7f50",
	"cornflowerblue":       "#6495ed",
	"cornsilk":             "#fff8dc",
	"crimson":              "#dc143c",
	"cyan":                 "#00ffff",
	"darkblue":             "#00008b",
	"darkcyan":             "#008b8b",
	"darkgoldenrod":        "#b8860b",
	"darkgray":             "#a9a9a9",
	"darkgrey":             "#a9a9a9",
	"darkgreen":            "#006400",
	"darkkhaki":            "#bdb76b",
	"darkmagenta":          "#8b008b",
	"darkolivegreen":       "#556b2f",
	"darkorange":           "#ff8c00",
	"darkorchid":           "#9932cc",
	"darkred":              "#8b0000",
	"darksalmon":           "#e9967a",
	"darkseagreen":         "#8fbc8f",
	"darkslateblue":        "#483d8b",
	"darkslategray":        "#2f4f4f",
	"darkslategrey":        "#2f4f4f",
	"darkturquoise":        "#00ced1",
	"darkviolet":           "#9400d3",
	"deeppink":             "#ff1493",
	"deepskyblue":          "#00bfff",
	"dimgray":              "#696969",
	"dimgrey":              "#696969",
	"dodgerblue":           "#1e90ff",
	"firebrick":            "#b22222",
	"floralwhite":          "#fffaf0",
	"forestgreen":          "#228b22",
	"fuchsia":              "#ff00ff",
	"gainsboro":            "#dcdcdc",
	"ghostwhite":           "#f8f8ff",
	"gold":                 "#ffd700",
	"goldenrod":            "#daa520",
	"gray":                 "#808080",
	"grey":                 "#808080",
	"green":                "#008000",
	"greenyellow":          "#adff2f",
	"honeydew":             "#f0fff0",
	"hotpink":              "#ff69b4",
	"indianred":            "#cd5c5c",
	"indigo":               "#4b0082",
	"ivory":                "#fffff0",
	"khaki":                "#f0e68c",
	"lavender":             "#e6e6fa",
	"lavenderblush":        "#fff0f5",
	"lawngreen":            "#7cfc00",
	"lemonchiffon":         "#fffacd",
	"lightblue":            "#add8e6",
	"lightcoral":           "#f08080",
	"lightcyan":            "#e0ffff",
	"lightgoldenrodyellow": "#fafad2",
	"lightgray":            "#d3d3d3",
	"lightgrey":            "#d3d3d3",
	"lightgreen":           "#90ee90",
	"lightpink":            "#ffb6c1",
	"lightsalmon":          "#ffa07a",
	"lightseagreen":        "#20b2aa",
	"lightskyblue":         "#87cefa",
	"lightslategray":       "#778899",
	"lightslategrey":       "#778899",
	"lightsteelblue":       "#b0c4de",
	"lightyellow":          "#ffffe0",
	"lime":                 "#00ff00",
	"limegreen":            "#32cd32",
	"linen":                "#faf0e6",
	"magenta":              "#ff00ff",
	"maroon":               "#800000",
	"mediumaquamarine":     "#66cdaa",
	"mediumblue":           "#0000cd",
	"mediumorchid":         "#ba55d3",
	"mediumpurple":         "#9370d8",
	"mediumseagreen":       "#3cb371",
	"mediumslateblue":      "#7b68ee",
	"mediumspringgreen":    "#00fa9a",
	"mediumturquoise":      "#48d1cc",
	"mediumvioletred":      "#c71585",
	"midnightblue":         "#191970",
	"mintcream":            "#f5fffa",
	"mistyrose":            "#ffe4e1",
	"moccasin":             "#ffe4b5",
	"navajowhite":          "#ffdead",
	"navy":                 "#000080",
	"oldlace":              "#fdf5e6",
	"olive":                "#808000",
	"olivedrab":            "#6b8e23",
	"orange":               "#ffa500",
	"orangered":            "#ff4500",
	"orchid":               "#da70d6",
	"palegoldenrod":        "#eee8aa",
	"palegreen":            "#98fb98",
	"paleturquoise":        "#afeeee",
	"palevioletred":        "#d87093",
	"papayawhip":           "#ffefd5",
	"peachpuff":            "#ffdab9",
	"peru":                 "#cd853f",
	"pink":                 "#ffc0cb",
	"plum":                 "#dda0dd",
	"powderblue":           "#b0e0e6",
	"purple":               "#800080",
	"rebeccapurple":        "#663399",
	"red":                  "#ff0000",
	"rosybrown":            "#bc8f8f",
	"royalblue":            "#4169e1",
	"saddlebrown":          "#8b4513",
	"salmon":               "#fa8072",
	"sandybrown":           "#f4a460",
	"seagreen":             "#2e8b57",
	"seashell":             "#fff5ee",
	"sienna":               "#a0522d",
	"silver":               "#c0c0c0",
	"skyblue":              "#87ceeb",
	"slateblue":            "#6a5acd",
	"slategray":            "#708090",
	"slategrey":            "#708090",
	"snow":                 "#fffafa",
	"springgreen":          "#00ff7f",
	"steelblue":            "#4682b4",
	"tan":                  "#d2b48c",
	"teal":                 "#008080",
	"thistle":              "#d8bfd8",
	"tomato":               "#ff6347",
	"turquoise":            "#40e0d0",
	"violet":               "#ee82ee",
	"wheat":                "#f5deb3",
	"white":                "#ffffff",
	"whitesmoke":           "#f5f5f5",
	"yellow":               "#ffff00",
	"yellowgreen":          "#9acd32",
}

type Color struct {
	RGB   []uint8
	Alpha float64
	Value interface{}
}

func New(rgb interface{}, opts ...interface{}) *Color {
	c := &Color{Alpha: 1}
	switch e := rgb.(type) {
	case []uint8:
		c.RGB = e
	case string:
		s := e
		if s[0] == '#' {
			s = s[1:]
		}
		hx := ""
		switch len(s) {
		case 3:
			x, y, z := string(s[0]), string(s[1]), string(s[2])
			hx = x + x + y + y + z + z
		case 6:
			hx = s
		default:
			panic(errors.New(e + " is unknown hex value"))
		}
		x, y, z := parseHex(hx)
		c.RGB[0], c.RGB[1], c.RGB[2] = x, y, z
	default:
		panic(errors.New("unsupported type"))
	}
	if len(opts) > 0 {
		a := opts[0]
		if al, ok := a.(float64); ok {
			c.Alpha = al
		}
	}
	return c
}

func parseHex(src string) (uint8, uint8, uint8) {
	h, _ := hex.DecodeString(src)
	return uint8(h[0]), uint8(h[1]), uint8(h[2])
}

func clamp(v, max float64) float64 {
	return math.Min(math.Max(v, 0), max)
}

func toHex(n ...float64) string {
	h := "#"
	for _, v := range n {
		c := clamp(math.Round(v), 255)
		u := uint64(c)
		if c < 16 {
			h += "0" + strconv.FormatUint(u, 16)
		} else {
			h += strconv.FormatUint(u, 16)
		}
	}
	return h
}

func (c *Color) toFloats() []float64 {
	return []float64{
		float64(c.RGB[0]),
		float64(c.RGB[1]),
		float64(c.RGB[2]),
	}
}

func (c *Color) Hex() string {
	return toHex(c.toFloats()...)
}

func (c *Color) Luma() float64 {
	f := c.toFloats()
	r, g, b := f[0]/255, f[1]/255, f[2]/255
	if r <= 0.03928 {
		r = r / 12.92
	} else {
		r = math.Pow(((r + 0.055) / 1.055), 2.4)
	}
	if g <= 0.03928 {
		g = g / 12.92
	} else {
		g = math.Pow(((b + 0.055) / 1.055), 2.4)
	}
	if b <= 0.03928 {
		b = b / 12.92
	} else {
		b = math.Pow(((b + 0.055) / 1.055), 2.4)
	}
	return 0.2126*r + 0.7152*g + 0.0722*b
}

func (c *Color) HSLA() (h, s, l, a float64) {
	f := c.toFloats()
	r, g, b := f[0]/255, f[1]/255, f[2]/255
	a = c.Alpha
	max := math.Max(math.Max(r, g), b)
	min := math.Min(math.Min(r, g), b)
	h = (max + min) / 2
	s = (max + min) / 2
	l = (max + min) / 2
	d := max - min
	if max == min {
		h, s = 0, 0
	} else {
		if l > 0.5 {
			s = d / (2 - max - min)
		} else {
			s = d / (max + min)
		}
		switch max {
		case r:
			var x float64
			if g < b {
				x = 6
			}
			h = (g-b)/d + x
		case g:
			h = (b-r)/d + 2
		case b:
			h = (r-g)/d + 4
		}
		h /= 6
	}
	h *= 360
	return
}

func (c *Color) HSVA() (h, s, v, a float64) {
	f := c.toFloats()
	r, g, b := f[0]/255, f[1]/255, f[2]/255
	a = c.Alpha
	max := math.Max(math.Max(r, g), b)
	min := math.Min(math.Min(r, g), b)
	h, s, v = max, max, max
	d := max - min
	if max == 0 {
		s = 0
	} else {
		s = d / max
	}
	if max == min {
		h = 0
	} else {
		switch max {
		case r:
			var x float64
			if g < b {
				x = 6
			}
			h = (g-b)/d + x
		case g:
			h = (b-r)/d + 2
		case b:
			h = (r-g)/d + 4
		}
		h /= 6
	}
	h *= 360
	return
}

func RGB(r, g, b uint8) *Color {
	return RGBA(r, g, b, 1.0)
}

func RGBA(r, g, b uint8, a float64) *Color {
	return New([3]uint8{r, g, b}, a)
}

func HSL(h, s, l float64) *Color {
	return HSLA(h, s, l, 1.0)
}

func HSLA(h, s, l, a float64) *Color {
	h = float64((int64(h) % 360) / 360)
	s = clamp0(s)
	l = clamp0(l)
	a = clamp0(a)

	var m2 float64
	if l <= 0.5 {
		m2 = l * (s + 1)
	} else {
		m2 = l + s - l*s
	}
	m1 := l*2 - m2
	r := hue(h+1/3, m1, m2) * 255
	g := hue(h, m1, m2) * 255
	b := hue(h-1/3, m1, m2) * 255
	return RGBA(
		uint8(r), uint8(g), uint8(b), a,
	)
}

func clamp0(v float64) float64 {
	return math.Min(1, math.Max(0, v))
}

func hue(h, m1, m2 float64) float64 {
	if h < 0 {
		h = h + 1
	} else {
		if h > 1 {
			h = h - 1
		}
	}
	if h*6 < 1 {
		return m2
	}
	if h*2 < 1 {
		return m2
	}
	if h*3 < 2 {
		return m1 + (m2-m1)*(2/3-h)*6
	}
	return m1
}

func HSV(h, s, v float64) *Color {
	return HSVA(h, s, v, 1.0)
}

func HSVA(h, s, v, a float64) *Color {
	h = float64((int64(h) % 360) / 360)
	i := (int(math.Floor(h)) / 60) % 6
	f := (h / 60) - float64(i)
	vs := []float64{
		v, v * (1 - s), v * (1 - f*s), v * (1 - (1-f)*s),
	}
	perm := [][]int{
		[]int{0, 3, 1},
		[]int{2, 0, 1},
		[]int{1, 0, 3},
		[]int{1, 2, 0},
		[]int{3, 1, 0},
		[]int{0, 1, 2},
	}
	r := vs[perm[i][0]] * 255
	g := vs[perm[i][1]] * 255
	b := vs[perm[i][2]] * 255
	return RGBA(uint8(r), uint8(g), uint8(b), a)
}

func (c *Color) Hue() float64 {
	h, _, _, _ := c.HSLA()
	return h
}
func (c *Color) Saturation() float64 {
	_, s, _, _ := c.HSLA()
	return s * 100
}

func (c *Color) Lightness() float64 {
	_, _, l, _ := c.HSLA()
	return l * 100
}

func (c *Color) Luminance() float64 {
	f := c.toFloats()
	l := (0.2126 * f[0] / 255) +
		(0.7152 * f[1] / 255) +
		(0.0722 * f[2] / 255)
	return l * c.Alpha * 100
}

func Saturate(c *Color, amount float64, method string) *Color {
	if c.RGB == nil {
		return nil
	}
	h, s, l, a := c.HSLA()
	if method == "relative" {
		s += s * amount / 100
	} else {
		s += amount / 100
	}
	s = clamp0(s)
	return HSLA(h, s, l, a)
}

func Desaturare(c *Color, amount float64, method string) *Color {
	if c.RGB == nil {
		return nil
	}
	h, s, l, a := c.HSLA()
	if method == "relative" {
		s -= s * amount / 100
	} else {
		s -= amount / 100
	}
	s = clamp0(s)
	return HSLA(h, s, l, a)
}

func Lighten(c *Color, amount float64, method string) *Color {
	h, s, l, a := c.HSLA()
	if method == "relative" {
		l += l * amount / 100
	} else {
		l += amount / 100
	}
	l = clamp0(l)
	return HSLA(h, s, l, a)
}

func Darken(c *Color, amount float64, method string) *Color {
	h, s, l, a := c.HSLA()
	if method == "relative" {
		l -= l * amount / 100
	} else {
		l -= amount / 100
	}
	l = clamp0(l)
	return HSLA(h, s, l, a)
}

func FadeIn(c *Color, amount float64, method string) *Color {
	h, s, l, a := c.HSLA()
	if method == "relative" {
		a += a * amount / 100
	} else {
		a += amount / 100
	}
	a = clamp0(a)
	return HSLA(h, s, l, a)
}

func FadeOut(c *Color, amount float64, method string) *Color {
	h, s, l, a := c.HSLA()
	if method == "relative" {
		a -= a * amount / 100
	} else {
		a -= amount / 100
	}
	a = clamp0(a)
	return HSLA(h, s, l, a)
}

func Fade(c *Color, amount float64) *Color {
	h, s, l, a := c.HSLA()
	a = amount / 100
	a = clamp0(a)
	return HSLA(h, s, l, a)
}

func Spin(c *Color, amount float64) *Color {
	h, s, l, a := c.HSLA()
	a = amount / 100
	hue := int64(h+amount) % 360
	if hue < 0 {
		h = 360 + float64(hue)
	} else {
		h = float64(hue)
	}
	return HSLA(h, s, l, a)
}

// Copyright (c) 2006-2009 Hampton Catlin, Natalie Weizenbaum, and Chris Eppstein
// http://sass-lang.com
//
func Mix(color1, color2 *Color, weight float64) *Color {
	if weight == 0 {
		weight = 59
	}
	p := weight / 100
	w := p*2 - 1
	_, _, _, a1 := color1.HSLA()
	_, _, _, a2 := color2.HSLA()
	a := a1 - a2
	var x float64
	if w*a == -1 {
		x = w
	} else {
		x = (w + a) / (1 + w*a)
	}
	w1 := x / 2.0
	w2 := 1 - w1
	r := color1.RGB[0]*uint8(w1) + color2.RGB[0]*uint8(w2)
	g := color1.RGB[1]*uint8(w1) + color2.RGB[1]*uint8(w2)
	b := color1.RGB[2]*uint8(w1) + color2.RGB[2]*uint8(w2)
	alpha := color1.Alpha*p + color2.Alpha*(1-p)
	return New([]uint8{r, g, b}, alpha)
}

func GreyScale(c *Color) *Color {
	return Desaturare(c, 100, "")
}

func Contrast(c *Color, dark, light *Color, threeshold float64) *Color {
	if light == nil {
		light = RGBA(255, 255, 255, 1.0)
	}
	if dark == nil {
		dark = RGBA(0, 0, 0, 1.0)
	}
	if dark.Luma() > light.Luma() {
		light, dark = dark, light
	}
	if threeshold == 0 {
		threeshold = 0.43
	}
	if c.Luma() < threeshold {
		return light
	}
	return dark
}

func Tint(c *Color, amount float64) *Color {
	return Mix(
		RGB(255, 255, 255), c, amount,
	)
}

func Shade(c *Color, amount float64) *Color {
	return Mix(
		RGB(0, 0, 0), c, amount,
	)
}
