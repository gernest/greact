package tests

import (
	"github.com/gernest/vected/lib/gs"
	"github.com/gernest/vected/lib/mad"
)

func TestToString() mad.Test {
	return mad.It("returns a valid css string", func(t mad.T) {
		css := gs.S(".test",
			gs.P("display", "block"),
			gs.P("box-sizing", "border-box"),
			gs.CSS(
				gs.P("position", "relative"),
				gs.P("margin-left", "0px"),
				gs.P("margin-right", "0px"),
				gs.P("height", "auto"),
			),
			gs.CSS(
				gs.P("zoom", "1"),
				gs.S("&:before",
					gs.P("content", " "),
					gs.P("display", "table"),
				),
				gs.S("&:after",
					gs.P("content", " "),
					gs.P("display", "table"),
					gs.P("clear", "both"),
					gs.P("visibility", "hidden"),
					gs.P("font-size", "0"),
					gs.P("height", "0"),
				),
			),
		)
		txt := gs.ToString(css)
		expect := `.test {
  display:block;
  box-sizing:border-box;
  position:relative;
  margin-left:0px;
  margin-right:0px;
  height:auto;
  zoom:1;
}
.test:before {
  content: ;
  display:table;
}
.test:after {
  content: ;
  display:table;
  clear:both;
  visibility:hidden;
  font-size:0;
  height:0;
}`
		if txt != expect {
			t.Errorf("expected %s got %s", expect, txt)
		}
	})
}
