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
	hex     string
	r, g, b uint8
}

func TestPalette(t *testing.T) {
	s := []struct {
		base    string
		h       float64
		s       float64
		v       float64
		r, g, b uint8
		hues    []hsvSample
	}{
		{base: "#f5222d",
			h: 357, s: 86, v: 96,
			r: 245, g: 34, b: 45,
			hues: []hsvSample{
				{index: 5, isLight: true, h: 4, s: 6, v: 100, hex: "#fff1f0",
					r: 255, g: 241, b: 240,
				},
				{index: 4, isLight: true, h: 5, s: 22, v: 100, hex: "#ffccc7",
					r: 255, g: 204, b: 199,
				},
				{index: 3, isLight: true, h: 3, s: 38, v: 100, hex: "#ffa39e",
					r: 255, g: 163, b: 158,
				},
				{index: 2, isLight: true, h: 1, s: 54, v: 100, hex: "#ff7875",
					r: 255, g: 120, b: 117,
				},
				{index: 1, isLight: true, h: 359, s: 70, v: 100, hex: "#ff4d4f",
					r: 255, g: 77, b: 79,
				},
				{index: 1, isLight: false, h: 355, s: 91, v: 81, hex: "#cf1322",
					r: 207, g: 19, b: 34,
				},
				{index: 2, isLight: false, h: 353, s: 96, v: 66, hex: "#a8071a",
					r: 168, g: 7, b: 76,
				},
				{index: 3, isLight: false, h: 351, s: 100, v: 51, hex: "#820014",
					r: 130, g: 0, b: 20,
				},
				{index: 4, isLight: false, h: 349, s: 100, v: 36, hex: "#5c0011",
					r: 92, g: 0, b: 17,
				},
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

			if v.r != o.r {
				t.Errorf("R: expected %v got %v", v.r, o.r)
			}
			if v.g != o.g {
				t.Errorf("G: expected %v got %v", v.g, o.g)
			}
			if v.b != o.b {
				t.Errorf("B: expected %v got %v", v.b, o.b)
			}
			testHSV(ts, o, v.hues)
		})
	}
}

func testHSV(t *testing.T, base *Color, s []hsvSample) {
	hh, ss, vv := base.HSV()
	for _, v := range s {
		h := getHue(hh, v.index, v.isLight)
		if h != v.h {
			t.Errorf("HUE [%v|%v] expected %v got %v", v.index, v.isLight, v.h, h)
		}
		s := getSaturation(ss, v.index, v.isLight)
		if s != v.s {
			t.Errorf("SATURATION [%v|%v] expected %v got %v", v.index, v.isLight, v.s, s)
		}
		ve := getValue(vv, v.index, v.isLight)
		if ve != v.v {
			t.Errorf("VALUE [%v|%v] expected %v got %v", v.index, v.isLight, v.v, ve)
		}
		o := HSV(h, s, ve)
		if v.r != o.r {
			t.Errorf("R [%v|%v]: expected %v got %v", v.index, v.isLight, v.r, o.r)
		}
		if v.g != o.g {
			t.Errorf("G [%v|%v]: expected %v got %v", v.index, v.isLight, v.g, o.g)
		}
		if v.b != o.b {
			t.Errorf("B [%v|%v]: expected %v got %v", v.index, v.isLight, v.b, o.b)
		}
		nh, ns, nv := o.HSV()
		if nh != h {
			t.Errorf("REPRO HUE [%v|%v] expected %v got %v", v.index, v.isLight, h, nh)
		}
		if ns != s {
			t.Errorf("REPRO SAT [%v|%v] expected %v got %v", v.index, v.isLight, s, ns)
		}
		if nv != ve {
			t.Errorf("REPRO VALUE [%v|%v] expected %v got %v", v.index, v.isLight, ve, nv)
		}
		hex := o.Hex()
		if hex != v.hex {
			t.Errorf("REPRO HEX [%v|%v] expected %v got %v", v.index, v.isLight, v.hex, hex)
		}
	}
}

func TestSpecialCase(t *testing.T) {
	s := hsvSample{index: 2, isLight: false, h: 353, s: 96, v: 66, hex: "#a8071a",
		r: 168, g: 7, b: 76,
	}
	o := InputToRGB(s.hex)
	if s.r != o.r {
		t.Errorf("expected %v got %v", s.r, o.r)
	}
	if s.g != o.g {
		t.Errorf("expected %v got %v", s.g, o.g)
	}

	oo := HSV(s.h, s.s, s.v)
	if s.r != oo.r {
		t.Errorf("R expected %v got %v", s.r, oo.r)
	}
	if s.g != oo.g {
		t.Errorf("G expected %v got %v", s.g, oo.g)
	}
	if s.b != oo.b {
		t.Errorf("B expected %v got %v", s.b, oo.b)
	}

}
