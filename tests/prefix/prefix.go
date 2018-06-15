package prefix

import (
	"reflect"
	"sort"
	"strconv"

	"github.com/gernest/gs/ciu/agents"

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
				b, err := prefix.NewBrowser(agents.New())
				if err != nil {
					t.Fatal(err)
				}
				expect := []string{"-moz-", "-ms-", "-o-", "-webkit-"}
				g := b.Prefixcache
				if !reflect.DeepEqual(expect, g) {
					t.Errorf("expected %v got %v", expect, g)
				}
			}),
		),
		mad.Describe("WithPrefix",
			mad.It("must return true when the prefix exist", func(t mad.T) {
				s := "1 -o-calc(1)"
				b, err := prefix.NewBrowser(agents.New())
				if err != nil {
					t.Fatal(err)
				}
				if !b.WithPrefix(s) {
					t.Error("expected to be true")
				}
			}),
			mad.It("must return false when the prefix does not exist", func(t mad.T) {
				s := "1 calc(1)"
				b, err := prefix.NewBrowser(agents.New())
				if err != nil {
					t.Fatal(err)
				}
				if b.WithPrefix(s) {
					t.Error("expected to be false")
				}
			}),
		),
		mad.Describe("selected",
			mad.It("must select the right browser", func(t mad.T) {
				b, err := prefix.NewBrowser(agents.New(), "chrome 30", "chrome 31")
				if err != nil {
					t.Fatal(err)
				}
				if !b.IsSelected("chrome 30") {
					t.Error("expected to be true")
				}
				if !b.IsSelected("chrome 31") {
					t.Error("expected to be true")
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
				b, err := prefix.NewBrowser(agents.New(), "chrome 30")
				if err != nil {
					t.Fatal(err)
				}
				for _, v := range s {
					g := b.Prefix(v.src)
					if g != v.expect {
						t.Errorf("expected %s got %s", v.expect, g)
					}
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

			b, err := prefix.NewBrowser(agents.New(), "firefox 21", "ie 7")
			if err != nil {
				t.Fatal(err)
			}
			fill := &prefix.Prefixes{
				Browsers: b,
				Data:     tdata.prefixes,
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
