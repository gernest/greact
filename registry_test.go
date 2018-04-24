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
			P("key", "value"),
		),
	))

	s.Attach()
	if o.detached {
		t.Errorf("expected to be attached")
	}

	s.Detach()
	if !o.detached {
		t.Errorf("expected to be detached")
	}
}
