package prefix

import (
	"reflect"

	"github.com/gernest/gs/prefix/decl"
	"github.com/gernest/vected/lib/gs"
	"github.com/gernest/vected/lib/mad"
)

func TestDecl() mad.Test {
	var d decl.Decl
	return mad.List{
		mad.Describe("Set",
			mad.It("sets prefixes of simple rule", func(t mad.T) {
				v := gs.SimpleRule{Key: "tab-size", Value: "4"}
				g := d.Set(v, "-moz-").(gs.SimpleRule)
				expect := "-moz-tab-size"
				if g.Key != expect {
					t.Errorf("expected %s got %s", expect, g.Key)
				}
				if g.Value != v.Value {
					t.Errorf("expected %s got %s", v.Value, g.Value)
				}
			}),
		),
		mad.Describe("Insert",
			mad.It("returns the same rule if no prefix is given", func(t mad.T) {
				v := gs.SimpleRule{Key: "tab-size", Value: "4"}
				g := d.Insert(v).(gs.SimpleRule)
				if g != v {
					t.Error("expected the same value")
				}
			}),
			mad.Describe("gs.SimpleRule",
				mad.It("returns a list of prefixed rules", func(t mad.T) {
					v := gs.SimpleRule{Key: "tab-size", Value: "4"}
					g := d.Insert(v, "-moz-", "-ie-").(gs.RuleList)
					expect := gs.RuleList{
						gs.SimpleRule{Key: "-moz-tab-size", Value: "4"},
						gs.SimpleRule{Key: "-ie-tab-size", Value: "4"},
						gs.SimpleRule{Key: "tab-size", Value: "4"},
					}
					if !reflect.DeepEqual(g, expect) {
						t.Errorf("expected %v got %v", expect, g)
					}
				}),
			),
		),
	}
}
