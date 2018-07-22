package grid

import (
	"github.com/gernest/vected/lib/gs"
	"github.com/gernest/vected/lib/mad"
	"github.com/gernest/vected/web/grid/style"
)

func TestGrid() mad.Test {
	return mad.List{
		mad.Describe("RowStyle",
			mad.It("generates all the styles for grid rows", func(t mad.T) {
				css := gs.ToString(style.RowStyle())
				if css != expectRowStyle {
					t.Errorf("expected %s got %s", expectRowStyle, css)
				}
			}),
		),
	}
}

const expectRowStyle = `.ant-row {
  position:relative;
  margin-left:0;
  margin-right:0;
  height:auto;
  zoom:1;
  display:block;
  box-sizing:border-box;
}
.ant-row:before {
  content:"";
  display:table;
}
.ant-row:after {
  content:"";
  display:table;
  clear:both;
  visibility:hidden;
  font-size:0;
  height:0;
}
.ant-row-flex {
  display:flex;
  flex-flow:row wrap;
}
.ant-row-flex:before {
  display:flex;
}
.ant-row-flex:after {
  display:flex;
}
.ant-row-flex-start {
  justify-content:flex-start;
}
.ant-row-flex-center {
  justify-content:center;
}
.ant-row-flex-end {
  justify-content:flex-end;
}
.ant-row-flex-space-between {
  justify-content:space-between;
}
.ant-row-flex-space-around {
  justify-content:space-around;
}
.ant-row-flex-top {
  justify-content:flex-start;
}
.ant-row-flex-middle {
  justify-content:center;
}
.ant-row-flex-bottom {
  justify-content:flex-bottom;
}`
