package tinycolor

import (
	"testing"
)

func TestRegex(t *testing.T) {
	s := []struct {
		src  string
		name string
	}{
		{"red", "hex3"},
		{"#f00", "hex3"},
		{"f00", "hex3"},
		{"#ff0000", "hex6"},
		{"ff0000", "hex6"},
		{"#ff000000", "hex8"},
		{"ff000000", "hex8"},
		{"rgb 255 0 0", "rgb"},
		{"rgb (255, 0, 0)", "rgb"},
		{"rgb 1.0 0 0", "rgb"},
		{"rgb (1, 0, 0)", "rgb"},
		{"rgba (255, 0, 0, 1)", "rgba"},
		{"rgba 255, 0, 0, 1", "rgba"},
		{"rgba (1.0, 0, 0, 1)", "rgba"},
		{"rgba 1.0, 0, 0, 1", "rgba"},
		{"hsl(0, 100%, 50%)", "hsl"},
		{"hsl 0 100% 50%", "hsl"},
		{"hsla(0, 100%, 50%, 1)", "hsla"},
		{"hsla 0 100% 50%, 1", "hsla"},
		{"hsv(0, 100%, 100%)", "hsv"},
		{"hsv 0 100% 100%", "hsv"},
	}

	for _, v := range s {
		o := matchColor(v.src)
		if o.name != v.name {
			t.Errorf("%s: expected %s got %s", v.src, v.name, o.name)
		}
	}
}

func TestBreakDown(t *testing.T) {
	s := []struct {
		src    string
		name   string
		expect []string
	}{
		{"rgb 255 0 0", "rgb", []string{"255", "0", "0"}},
		{"rgb (255, 0, 0)", "rgb", []string{"255", " 0", " 0"}},
		{"rgb 1.0 0 0", "rgb", []string{"1.0", "0", "0"}},
		{"rgb (1, 0, 0)", "rgb", []string{"1", " 0", " 0"}},
		{"rgba (255, 0, 0, 1)", "rgba", []string{"255", " 0", " 0", " 1"}},
		{"rgba 255, 0, 0, 1", "rgba", []string{"255", " 0", " 0", " 1"}},
		{"rgba (1.0, 0, 0, 1)", "rgba", []string{"1.0", " 0", " 0", " 1"}},
		{"rgba 1.0, 0, 0, 1", "rgba", []string{"1.0", " 0", " 0", " 1"}},
		{"hsl(0, 100%, 50%)", "hsl", []string{"0", " 100%", " 50%"}},
		{"hsl 0 100% 50%", "hsl", []string{"0", "100%", "50%"}},
		{"hsla(0, 100%, 50%, 1)", "hsla", []string{"0", " 100%", " 50%", " 1"}},
		{"hsla 0 100% 50%, 1", "hsla", []string{"0 100% 50%", " 1"}},
		{"hsv(0, 100%, 100%)", "hsv", []string{"0", " 100%", " 100%"}},
		{"hsv 0 100% 100%", "hsv", []string{"0", "100%", "100%"}},
	}

	for _, v := range s {
		o := breadDown(v.name, v.src)
		if !equalSlice(v.expect, o) {
			t.Errorf("%x: expected %s got %s", v.src, v.expect, o)
		}
	}
}

func equalSlice(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for k, v := range a {
		bv := b[k]
		if v != bv {
			return false
		}
	}
	return true
}

func TestParseHex(t *testing.T) {
	s := []struct {
		hex        string
		r, g, b, a uint8
	}{
		{"#5f5", 85, 255, 85, 0},
		{"#5f55f5", 95, 85, 245, 0},
	}
	for _, v := range s {
		o := matchColor(v.hex)
		c, err := o.toColor()
		if err != nil {
			t.Fatal(err)
		}
		if c.r != v.r {
			t.Errorf("%s [R]: expected %v got %v", v.hex, v.r, c.r)
		}
		if c.g != v.g {
			t.Errorf("%s [G]: expected %v got %v", v.hex, v.g, c.g)
		}
		if c.b != v.b {
			t.Errorf("%s [B]: expected %v got %v", v.hex, v.b, c.b)
		}
		if c.a != v.a {
			t.Errorf("%s [A]: expected %v got %v", v.hex, v.a, c.a)
		}
	}
}

func TestParseRGB(t *testing.T) {
	s := []struct {
		hex        string
		r, g, b, a uint8
	}{
		{"rgb (85,255,85)", 85, 255, 85, 0},
		{"rgb 95,85,245", 95, 85, 245, 0},
	}
	for _, v := range s {
		o := matchColor(v.hex)
		c, err := o.toColor()
		if err != nil {
			t.Fatal(err)
		}
		if c.r != v.r {
			t.Errorf("%s [R]: expected %v got %v", v.hex, v.r, c.r)
		}
		if c.g != v.g {
			t.Errorf("%s [G]: expected %v got %v", v.hex, v.g, c.g)
		}
		if c.b != v.b {
			t.Errorf("%s [B]: expected %v got %v", v.hex, v.b, c.b)
		}
		if c.a != v.a {
			t.Errorf("%s [A]: expected %v got %v", v.hex, v.a, c.a)
		}
	}
}

type hsvSample struct {
	index   int
	isLight bool
	h       float64
	s       float64
	v       float64
}

func TestPalette(t *testing.T) {
	s := []struct {
		base string
		h    float64
		s    float64
		v    float64
		hues []hsvSample
	}{
		{base: "#f5222d",
			h: 357, s: 86, v: 96,
			hues: []hsvSample{
				{index: 5, isLight: true, h: 4, s: 6, v: 100},
				{index: 4, isLight: true, h: 5, s: 22, v: 100},
				{index: 3, isLight: true, h: 3, s: 38, v: 100},
				{index: 2, isLight: true, h: 1, s: 54, v: 100},
				{index: 1, isLight: true, h: 359, s: 70, v: 100},
				{index: 1, isLight: false, h: 355, s: 91, v: .81},
				{index: 2, isLight: false, h: 363, s: 96, v: 66},
				{index: 4, isLight: false, h: 351, s: 100, v: 51},
				{index: 5, isLight: false, h: 349, s: 100, v: 36},
			}},
	}

	for _, v := range s {
		t.Run(v.base, func(ts *testing.T) {
			o := InputToRGB(v.base)
			h, s, ve := o.HSV()
			if h != v.h {
				t.Errorf("expected %v got %v", v.h, h)
			}
			if s != v.s {
				t.Errorf("expected %v got %v", v.s, s)
			}
			if ve != v.v {
				t.Errorf("expected %v got %v", v.v, ve)
			}
			testHSV(ts, o, v.hues)
		})
	}
}

func testHSV(t *testing.T, base *Color, s []hsvSample) {
	for _, v := range s {
		g := Palette(base, v.index)
		h, s, ve := g.HSV()
		if h != v.h {
			t.Errorf("expected %v got %v", v.h, h)
		}
		if s != v.s {
			t.Errorf("expected %v got %v", v.s, s)
		}
		if ve != v.v {
			t.Errorf("expected %v got %v", v.v, ve)
		}
	}
}
