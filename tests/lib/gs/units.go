package gs

import (
	"github.com/gernest/vected/lib/gs"
	"github.com/gernest/vected/lib/mad"
)

func TestGSUnits() mad.Test {
	return mad.List{
		mad.Describe("Unit",
			mad.It("it returns the string representation of css unit of measurement", func(t mad.T) {
				sample := []struct {
					src, expect string
				}{
					{"0px", "px"},
					{"10%", "%"},
				}
				for _, v := range sample {
					got := gs.U(v.src).Unit()
					if got != v.expect {
						t.Errorf("expected %s got %s", v.expect, got)
					}
				}
			}),
		),
		mad.Describe("Value",
			mad.It("returns value from css unit", func(t mad.T) {
				got := gs.U("10px").Value()
				expect := float64(10)
				if got != expect {
					t.Errorf("expected %v got %v", expect, got)
				}
			}),
			mad.It("converts % to floats", func(t mad.T) {
				got := gs.U("10%").Value()
				expect := 0.1
				if got != expect {
					t.Errorf("expected %v got %v", expect, got)
				}
			}),
		),
		mad.Describe("Div",
			mad.It("divides same unit", func(t mad.T) {
				a := gs.U("20px")
				b := gs.U("2")
				g := a.Div(b)
				got := g.String()
				expect := "10px"
				if got != expect {
					t.Errorf("expected %s got %s", expect, got)
				}
			}),
		),
	}
}
