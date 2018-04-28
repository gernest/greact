package prefix

import (
	"strconv"
	"testing"

	"github.com/gernest/gs/agents"
)

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

func TestBrowser_Prefixes(t *testing.T) {
	b := NewBrowser()
	expect := []string{"-moz-", "-ms-", "-o-", "-webkit-"}
	g := b.Prefixcache
	for k, v := range expect {
		if g[k] != v {
			t.Fatalf("expected %v got %v", expect, g)
		}
	}
}
func TestBrowser_WithPrefix(t *testing.T) {
	b := NewBrowser()
	s := []struct {
		src    string
		expect bool
	}{
		{"1 -o-calc(1)", true},
		{"1 calc(1)", false},
	}

	for _, v := range s {
		g := b.WithPrefix(v.src)
		if g != v.expect {
			t.Errorf("%s: expected %v got %v", v.src, v.expect, g)
		}
	}
}

func parseFloat(v string) float64 {
	f, _ := strconv.ParseFloat(v, 64)
	return f
}

func TestBrowser_Selected(t *testing.T) {
	var less = func(name, ver string) bool {
		if ver == "" {
			return false
		}
		v := parseFloat(ver)
		return name == agents.InternetExplorer.Name && v < 7
	}

	var combined = func(name, ver string) bool {
		if ver == "" {
			return false
		}
		v := parseFloat(ver)
		return name == agents.InternetExplorer.Name && v == 10 ||
			name == agents.InternetExplorer.Name && v < 6
	}

	var nothing = func(name, ver string) bool {
		return false
	}
	s := []struct {
		src    func(string, string) bool
		expect []string
	}{
		{nothing, nil},
		{less, []string{"ie 5.5", "ie 6"}},
		{combined, []string{"ie 5.5", "ie 10"}},
	}

	for _, v := range s {
		b := NewBrowser(v.src)
		if v.expect == nil {
			if b.Selected != nil {
				t.Errorf("expected nil got %v", b.Selected)
			}
		} else {
			if len(v.expect) != len(b.Selected) {
				t.Fatalf("expected %v got %v", v.expect, b.Selected)
			}
			for k, val := range v.expect {
				if b.Selected[k] != val {
					t.Fatalf("expected %v got %v", v.expect, b.Selected)
				}
			}
		}
	}
}

func TestBrowser_Prefix(t *testing.T) {
	s := []struct {
		src    string
		expect string
	}{
		{"chrome 30", "-webkit-"},
	}
	b := NewBrowser()
	for _, v := range s {
		g := b.Prefix(v.src)
		if g != v.expect {
			t.Errorf("expected %s got %s", v.expect, g)
		}
	}
}

func TestBrowser_IsSelected(t *testing.T) {
	b := NewBrowser(func(name, ver string) bool {
		if ver == "" {
			return false
		}
		v := parseFloat(ver)
		return name == agents.Chrome.Name && v == 30 ||
			name == agents.Chrome.Name && v == 31
	})

	if !b.IsSelected("chrome 30") {
		t.Error("expected to be true")
	}
	if b.IsSelected("ie 6") {
		t.Error("expected to be false")
	}
}
