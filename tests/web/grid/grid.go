package grid

import (
	"github.com/gernest/gs"
	"github.com/gernest/mad"
	"github.com/gernest/vected/web/components/grid"
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

func TestRow() mad.Test {
	return mad.List{
		mad.Describe("MakeRow", makeRow()),
		mad.Describe("RowStyle", rowStyle()),
	}
}

func makeRow() mad.Test {
	return mad.List{
		mad.It("must use default gutter width", func(t mad.T) {
			expect := `position:relative;

margin-left:0px;

margin-right:0px;

height:auto;

zoom:1;

&:before {
  content: ;
  display:table;
}

&:after {
  content: ;
  display:table;
  clear:both;
  visibility:hidden;
  font-size:0;
  height:0;
}`
			css := grid.MakeRow(0)
			str := gs.ToString(css)
			if str != expect {
				t.Errorf("expected %s got %s", expect, str)
			}
		}),
		mad.It("must calculate the gutter", func(t mad.T) {
			expect := `position:relative;

margin-left:-8px;

margin-right:-8px;

height:auto;

zoom:1;

&:before {
  content: ;
  display:table;
}

&:after {
  content: ;
  display:table;
  clear:both;
  visibility:hidden;
  font-size:0;
  height:0;
}`
			css := grid.MakeRow(16)
			str := gs.ToString(css)
			if str != expect {
				t.Errorf("expected %s got %s", expect, str)
			}
		}),
	}
}

func rowStyle() mad.Test {
	return mad.It("creates base row style", func(t mad.T) {
		css := grid.RowStyle(15)
		txt := gs.ToString(css)
		expect := `.vected-row {
  position:relative;
  margin-left:-7px;
  margin-right:-7px;
  height:auto;
  zoom:1;
  .vected-row:before {
    content: ;
    display:table;
  }
  .vected-row:after {
    content: ;
    display:table;
    clear:both;
    visibility:hidden;
    font-size:0;
    height:0;
  }
  display:block;
  box-sizing:border-box;
}`
		if txt != expect {
			t.Errorf("expected %s got %s", expect, txt)
		}
	})
}
