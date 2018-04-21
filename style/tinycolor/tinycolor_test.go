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
