package gs

import (
	"testing"
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

func TestRegistry(t *testing.T) {
	var o *mockSheetObject
	r := NewSimpleRegistry(func() SheetObject {
		o = &mockSheetObject{}
		return o
	})
	s := r.NewSheet()
	s.AddRule(CSS(
		S(".hello",
			P("key1", "value1"),
			P("key2", "value2"),
		),
	))
	s.Attach()
	if o.detached {
		t.Errorf("expected to be attached")
	}
	e := ".hello-1-1{key1:value1;key2:value2;}"
	if o.rules[0] != e {
		t.Errorf("expected %s got %s", e, o.rules[0])
	}
	s.Detach()
	if !o.detached {
		t.Errorf("expected to be detached")
	}
}
