package grid

import (
	"github.com/gernest/mad"
	"github.com/gernest/vected/components/grid"
)

func TestMediaType() mad.Test {
	return mad.Describe("Screen",
		mad.It("returns a media query string", func(t mad.T) {
			sample := []struct {
				media  grid.MediaType
				expect string
			}{
				{media: grid.XS, expect: "@media (min-width:480px)"},
				{media: grid.SM, expect: "@media (min-width:576px)"},
				{media: grid.MD, expect: "@media (min-width:768px)"},
				{media: grid.LG, expect: "@media (min-width:992px)"},
				{media: grid.XL, expect: "@media (min-width:1200px)"},
				{media: grid.XXL, expect: "@media (min-width:1600px)"},
			}
			for _, v := range sample {
				g := v.media.Screen()
				if g != v.expect {
					t.Errorf("%s expected %s got %s", v.media, v.expect, g)
				}
			}
		}),
	)
}
