package breadcrumb

import (
	"strings"

	"github.com/gernest/vected/lib/gs"
	"github.com/gernest/vected/lib/mad"
	"github.com/gernest/vected/web/breadcrumb/style"
)

func TestBreadcrumbStyle() mad.Test {
	return mad.It("generate css", func(t mad.T) {
		css := gs.ToString(style.Breadcrumb())
		e := strings.TrimSpace(expect)
		if css != e {
			t.Error("got wrong styles")
		}
	})
}

const expect = `.ant-breadcrumb {
  font-family:"Monospaced Number","Chinese Quote", -apple-system, BlinkMacSystemFont, "Segoe UI", Roboto, "PingFang SC", "Hiragino Sans GB", "Microsoft YaHei", "Helvetica Neue", Helvetica, Arial, sans-serif;
  font-size:14px;
  line-height:1.5;
  color:rgba(0,0,0,0.65);
  box-sizing:border-box;
  margin:0;
  padding:0;
  list-style:none;
  color:rgba(0,0,0,0.45);
  font-size:14px;
}
.ant-breadcrumb .anticon {
  font-size:12px;
}
.ant-breadcrumb a {
  color:rgba(0,0,0,0.45);
  transition:color .3s;
}
.ant-breadcrumb a:hover {
  color:#40a9ff;
}
.ant-breadcrumb > span:last-child {
  color:rgba(0,0,0,0.65);
}
.ant-breadcrumb > span:last-child .ant-breadcrumb-separator {
  display:none;
}
.ant-breadcrumb-separator {
  margin:0 8px;
  color:rgba(0,0,0,0.45);
}
.ant-breadcrumb-link > .anticon + span {
  margin-left:4px;
}`
