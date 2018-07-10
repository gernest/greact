package tests

import (
	"github.com/gernest/gs"
	"github.com/gernest/mad"
)

func TestSheet() mad.Test {
	var dummyIDGen = func() string {
		return "id"
	}
	return mad.Describe("AddRule",
		mad.It("generate classes with new id", func(t mad.T) {
			s := gs.CSS(
				gs.S(".hello",
					gs.P("key", "value"),
				),
			)
			sheet := gs.NewSheet(dummyIDGen)
			sheet.AddRule(s)
			e := ".hello-id"
			g := sheet.CLasses[".hello"]
			if g != e {
				t.Errorf("expected %s got %s", e, g)
			}
		}),
	)
}
