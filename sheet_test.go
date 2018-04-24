package gs

import (
	"testing"
)

func TestSheet(t *testing.T) {
	var dummyIdGen = func() string {
		return "id"
	}
	s := CSS(
		S(".hello",
			P("key", "value"),
		),
	)
	sheet := NewSheet(s, dummyIdGen)
	e := ".hello-id"
	g := sheet.CLasses["hello"]
	if g != e {
		t.Errorf("expected %s got %s", e, g)
	}
}
