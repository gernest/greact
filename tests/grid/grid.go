package grid

import (
	"github.com/gernest/gs"
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

func TestRowStyle() mad.Test {
	fixture := []string{
		`.Row {
  position:relative;
  margin-left:-8px;
  margin-right:-8px;
  box-sizing:border-box;
  display:block;
  height:auto;
  zoom:1;
}

.Row:before {
  content:;
  display:table;
}

.Row:after {
  content:;
  display:table;
  clear:both;
  visibility:hidden;
  font-size:0;
  height:0;
}`,
	}
	return mad.It("creates row styles", func(t mad.T) {
		sample := []struct {
			gutter  int64
			flex    bool
			justify grid.FlexStyle
			align   grid.FlexAlign
		}{
			{16, false, grid.Start, grid.Top},
		}
		for k, v := range sample {
			css := grid.RowStyle(v.gutter, v.flex, v.justify, v.align)
			got := gs.ToString(css)
			expect := fixture[k]
			if got != expect {
				t.Errorf("expected %s got %s", expect, got)
			}
		}
	})
}
