package prefix

import (
	"github.com/gernest/gs"
	"github.com/gernest/gs/prefix/hacks"
	"github.com/gernest/mad"
)

func TestBorderImage() mad.Test {
	return mad.Describe("Set",
		mad.It("replacess fill with first submatch", func(t mad.T) {
			var b hacks.BorderImage
			g := b.Set(gs.SimpleRule{
				Value: "  fill x",
			}, "").(gs.SimpleRule)
			e := " x"
			if g.Value != e {
				t.Errorf("expected %s got %s", e, g)
			}
		}),
	)
}
