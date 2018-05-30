package tests

import (
	"github.com/gernest/gs/prefix"
	"github.com/gernest/mad"
)

func TestPrefix() mad.Test {
	return mad.Describe("Basic",
		mad.It("adds vendor prefix", func(t mad.T) {
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
		}),
	)
}

func TestUnPrefix() mad.Test {
	return mad.Describe("Basic",
		mad.It("removes vendor prefix", func(t mad.T) {
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
		}),
	)
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
	}
}
