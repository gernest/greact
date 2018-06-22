package iconfont

import (
	"github.com/gernest/gs"
	"github.com/gernest/mad"
	"github.com/gernest/vected/web/style/iconfont"
)

func TestFont() mad.Test {
	return mad.It("generates  iconfont  class style", func(t mad.T) {
		css := gs.ToString(iconfont.Font())
		expect := `.anticon {
  display:inline-block;
  font-style:normal;
  vertical-align:baseline;
  text-align:center;
  text-transform::none;
  line-height:1;
  text-rendering:optimizeLegibility;
  -webkit-font-smoothing:antialiased;
  -moz-osx-font-smoothing:grayscale;
}
.anticon:before {
  display:block;
  font-family:"anticon" !important;
}`
		if css != expect {
			t.Errorf("expected %s got %s", expect, css)
		}
	})
}
