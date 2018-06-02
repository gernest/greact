package tests

import (
	"reflect"
	"sort"
	"strconv"

	"github.com/gernest/gs/agents"
	"github.com/gernest/gs/data"
	"github.com/gernest/gs/prefix"
	"github.com/gernest/mad"
)

func TestPrefix() mad.Test {
	return mad.It("adds vendor prefix", func(t mad.T) {
		s := []struct {
			src, expect string
		}{
			{src: "-moz-tab-size", expect: "-moz-"},
		}

		for _, v := range s {
			g := prefix.Prefix(v.src)
			if g != v.expect {
				t.Errorf("expected %s got %s", v.expect, g)
			}
		}
	})
}

func TestUnPrefix() mad.Test {
	return mad.It("removes vendor prefix", func(t mad.T) {
		s := []struct {
			src, expect string
		}{
			{src: "-moz-tab-size", expect: "tab-size"},
		}

		for _, v := range s {
			g := prefix.UnPrefix(v.src)
			if g != v.expect {
				t.Errorf("expected %s got %s", v.expect, g)
			}
		}
	})
}

func parseFloat(v string) float64 {
	f, _ := strconv.ParseFloat(v, 64)
	return f
}

func TestBrowser() mad.Test {
	return mad.List{
		mad.Describe("Prefixes",
			mad.It("must contain all browser vendor prefixes", func(t mad.T) {
				b := prefix.NewBrowser()
				expect := []string{"-moz-", "-ms-", "-o-", "-webkit-"}
				g := b.Prefixcache
				for k, v := range expect {
					if g[k] != v {
						t.Fatalf("expected %v got %v", expect, g)
					}
				}
			}),
		),
		mad.Describe("WithPrefix",
			mad.It("must return true when the prefix exist", func(t mad.T) {
				s := "1 -o-calc(1)"
				b := prefix.NewBrowser()
				if !b.WithPrefix(s) {
					t.Error("expected to be true")
				}
			}),
			mad.It("must return false when the prefix does not exist", func(t mad.T) {
				s := "1 calc(1)"
				b := prefix.NewBrowser()
				if b.WithPrefix(s) {
					t.Error("expected to be false")
				}
			}),
		),
		mad.Describe("selected",
			mad.It("must select the right browser", func(t mad.T) {
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
					b := prefix.NewBrowser(v.src)
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
			}),
		),
		mad.Describe("Prefix",
			mad.It("must return browser prefix", func(t mad.T) {
				s := []struct {
					src    string
					expect string
				}{
					{"chrome 30", "-webkit-"},
				}
				b := prefix.NewBrowser()
				for _, v := range s {
					g := b.Prefix(v.src)
					if g != v.expect {
						t.Errorf("expected %s got %s", v.expect, g)
					}
				}
			}),
		),
		mad.Describe("IsSelected",
			mad.It("must be selected", func(t mad.T) {
				b := prefix.NewBrowser(func(name, ver string) bool {
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
			}),
		),
	}
}

func TestPrefixes() mad.Test {
	return mad.Describe("Select",
		mad.It("must select stuff", func(t mad.T) {
			tdata := struct {
				prefixes map[string]data.Data
			}{
				prefixes: map[string]data.Data{
					"a": data.Data{
						Browsers: []string{"firefox 21", "firefox 20 old", "chrome 30", "ie 6"},
					},
					"b": data.Data{
						Browsers: []string{"ie 7 new", "firefox 20"},
						Mistakes: []string{"-webkit-"},
						Props:    []string{"a", "*"},
					},
					"c": data.Data{
						Browsers: []string{"ie 7", "firefox 20"},
						Selector: true,
					},
				},
			}

			fill := &prefix.Prefixes{
				Browsers: prefix.NewBrowser(func(name, version string) bool {
					return name == "firefox" && version == "21" ||
						name == "ie" && version == "7"
				}),
				Data: tdata.prefixes,
			}

			sample := []struct {
				key   string
				add   bool
				value []string
			}{
				{key: "a", add: true, value: []string{"-moz-"}},
				{key: "b", add: true, value: []string{"-ms- new"}},
				{key: "c", add: true, value: []string{"-ms-"}},
				{key: "a", add: false, value: []string{"-webkit-", "-ms-", "-moz- old"}},
				{key: "b", add: false, value: []string{"-ms-", "-moz-", "-webkit-"}},
				{key: "c", add: false, value: []string{"-moz-"}},
			}

			sel := fill.Select(tdata.prefixes)
			for _, v := range sample {
				var g []string
				if v.add {
					g = sel.Add[v.key]
				} else {
					g = sel.Remove[v.key]
				}
				sort.Strings(g)
				sort.Strings(v.value)
				if !reflect.DeepEqual(g, v.value) {
					t.Errorf("expected %v got %v", v.value, g)
				}
			}
		}),
	)
}
