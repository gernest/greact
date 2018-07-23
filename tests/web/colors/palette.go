package color

import (
	"math"

	"github.com/gernest/vected/lib/mad"
	"github.com/gernest/vected/web/style/color"
)

type hsvSample struct {
	index   int
	isLight bool
	h       float64
	s       float64
	v       float64
	hex     string
	r, g, b uint8
}

func TestPalette2() mad.Test {
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
				{index: 1, isLight: true, h: 4, s: 6, v: 100, hex: "#fff1f0",
					r: 255, g: 241, b: 240,
				},
				{index: 2, isLight: true, h: 5, s: 22, v: 100, hex: "#ffccc7",
					r: 255, g: 204, b: 199,
				},
				{index: 3, isLight: true, h: 3, s: 38, v: 100, hex: "#ffa39e",
					r: 255, g: 163, b: 158,
				},
				{index: 4, isLight: true, h: 1, s: 54, v: 100, hex: "#ff7875",
					r: 255, g: 120, b: 117,
				},
				{index: 5, isLight: true, h: 359, s: 70, v: 100, hex: "#ff4d4f",
					r: 255, g: 77, b: 79,
				},
				{index: 7, isLight: false, h: 355, s: 91, v: 81, hex: "#cf1322",
					r: 207, g: 19, b: 34,
				},
				{index: 8, isLight: false, h: 353, s: 96, v: 66, hex: "#a8071a",
					r: 168, g: 7, b: 76,
				},
				{index: 9, isLight: false, h: 351, s: 100, v: 51, hex: "#820014",
					r: 130, g: 0, b: 20,
				},
				{index: 10, isLight: false, h: 349, s: 100, v: 36, hex: "#5c0011",
					r: 92, g: 0, b: 17,
				},
			}},
	}

	var cases []mad.Test
	for _, v := range s {
		cases = append(cases, mad.It(v.base, func(t mad.T) {
			o := color.New(v.base)
			h, s, ve, _ := o.HSVA()
			h = math.Round(h)
			s = math.Round(s * 100)
			ve = math.Round(ve * 100)
			if h != v.h {
				t.Errorf("expected %v got %v", v.h, h)
			}
			if s != v.s {
				t.Errorf("expected %v got %v", v.s, s)
			}
			if ve != v.v {
				t.Errorf("expected %v got %v", v.v, ve)
			}

			if v.r != o.RGB[0] {
				t.Errorf("R: expected %v got %v", v.r, o.RGB[0])
			}
			if v.g != o.RGB[1] {
				t.Errorf("G: expected %v got %v", v.g, o.RGB[1])
			}
			if v.b != o.RGB[2] {
				t.Errorf("B: expected %v got %v", v.b, o.RGB[2])
			}
			for _, hx := range v.hues {
				nc := color.GenerateColor(o, hx.index)
				if nc.Hex() != hx.hex {
					t.Errorf("%v:%v expected %s got %s", hx.index, hx.isLight, hx.hex, nc.Hex())
				}
			}
		}))
	}
	return mad.Describe("Generates correct color pellete", cases...)
}

func TestGenerateColor() mad.Test {
	return mad.It("generates colors", func(t mad.T) {
		base := color.New("#f5222d")
		c := color.GenerateColor(base, 9)
		expect := color.New("#820014")
		if c.Hex() != expect.Hex() {
			t.Errorf("expected %v got %v", expect.Hex(), c.Hex())
		}
		h := c.ToHSV()
		e := expect.ToHSV()
		if h != e {
			t.Errorf("expected %s got %s", e, h)
		}
	})
}

func TestGenerate_bug_blue_pallete() mad.Test {
	base := color.New("#1890ff")
	pallete := color.Generate(base)
	return mad.List{
		mad.It("must contain the base color", func(t mad.T) {
			i := pallete[5]
			bx := base.Hex()
			ix := i.Hex()
			if ix != bx {
				t.Errorf("expected %s got %s", bx, ix)
			}
		}),
	}
}
