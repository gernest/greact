package prefix

import (
	"reflect"

	"github.com/gernest/vected/lib/gs"

	"github.com/gernest/vected/lib/ciu/agents"

	"github.com/gernest/vected/lib/mad"
	"github.com/gernest/vected/tools/prefix"
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

func TestFindPrefix() mad.Test {
	b, _ := prefix.NewBrowser(agents.New())
	return mad.List{
		mad.It("finds in at rules", func(t mad.T) {
			css := gs.Cond("@-ms-keyframes a",
				gs.S("to"),
			)
			g := prefix.FindPrefix(b, css)
			expect := "-ms-"
			if g != expect {
				t.Errorf("expected %s got %s", expect, g)
			}
		}),
		mad.It("finds in selectors", func(t mad.T) {
			css := gs.S(":-moz-full-screen")
			expect := "-moz-"
			g := prefix.FindPrefix(b, css)
			if g != expect {
				t.Errorf("expected %s got %s", expect, g)
			}
		}),
		mad.It("finds only browsers prefixes", func(t mad.T) {
			css := gs.Cond("@-dev-keyframes")
			expect := ""
			g := prefix.FindPrefix(b, css)
			if g != expect {
				t.Errorf("expected %s got %s", expect, g)
			}
		}),
	}
}
