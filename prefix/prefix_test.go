package prefix

import "testing"

func TestPrefix(t *testing.T) {
	s := []struct {
		src, expect string
	}{
		{src: "-moz-tab-size", expect: "-moz-"},
	}

	for _, v := range s {
		g := Prefix(v.src)
		if g != v.expect {
			t.Errorf("expected %s got %s", v.expect, g)
		}
	}
}
func TestUnPrefix(t *testing.T) {
	s := []struct {
		src, expect string
	}{
		{src: "-moz-tab-size", expect: "tab-size"},
	}

	for _, v := range s {
		g := UnPrefix(v.src)
		if g != v.expect {
			t.Errorf("expected %s got %s", v.expect, g)
		}
	}
}
