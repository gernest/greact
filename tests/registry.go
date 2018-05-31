package tests

import (
	"github.com/gernest/gs"
	"github.com/gernest/mad"
)

type mockSheetObject struct {
	rules    []string
	detached bool
}

func (m *mockSheetObject) InsertRule(rule string) {
	m.rules = append(m.rules, rule)
}

func (m *mockSheetObject) Detach() {
	m.detached = true
}

func TestRegistry() mad.Test {
	var o *mockSheetObject
	r := gs.NewSimpleRegistry(func() gs.SheetObject {
		o = &mockSheetObject{}
		return o
	})
	rules := gs.CSS(
		gs.S(".hello",
			gs.P("key1", "value1"),
			gs.P("key2", "value2"),
		))
	return mad.List{
		mad.Describe("Attach",
			mad.It("must call attach handler with the rule string", func(t mad.T) {
				s := r.NewSheet()
				s.AddRule(rules)
				s.Attach()
				if o.detached {
					t.Errorf("expected to be attached")
				}
				e := `.hello-1-1 {
  key1:value1;
  key2:value2;
}`
				if o.rules[0] != e {
					t.Errorf("expected %s got %s", e, o.rules[0])
				}
			}),
		),
		mad.Describe("Detach",
			mad.It("must call detach  handler", func(t mad.T) {
				s := r.NewSheet()
				s.AddRule(rules)
				s.Attach()
				s.Detach()
				if !o.detached {
					t.Errorf("expected to be detached")
				}
			}),
		),
	}
}
