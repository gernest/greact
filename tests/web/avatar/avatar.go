package avatar

import (
	"github.com/gernest/vected/lib/gs"
	"github.com/gernest/vected/lib/mad"
	"github.com/gernest/vected/web/avatar/style"
)

func TestStyle() mad.Test {
	return mad.It("generates css for avatar", func(t mad.T) {
		css := gs.ToString(style.Avatar())
		if css != expect {
			t.Error("got wrong styles")
		}
	})
}

const expect = `.ant-avatar {
  font-family:"Monospaced Number","Chinese Quote", -apple-system, BlinkMacSystemFont, "Segoe UI", Roboto, "PingFang SC", "Hiragino Sans GB", "Microsoft YaHei", "Helvetica Neue", Helvetica, Arial, sans-serif;
  font-size:14px;
  line-height:1.5;
  color:rgba(0,0,0,0.65);
  box-sizing:border-box;
  margin:0;
  padding:0;
  list-style:none;
  display:inline-block;
  text-align:center;
  background:#cccccc;
  color:#ffffff;
  white-space:nowrap;
  position:relative;
  overflow:hidden;
  vertical-align:middle;
  width:32px;
  height:32px;
  line-height:32px;
  border-radius:50%;
}
.ant-avatar-image {
  background:transparent;
}
.ant-avatar > * {
  line-height:32px;
}
.ant-avatar.ant-avatar-icon {
  font-size:18px;
}
.ant-avatar-lg {
  width:40px;
  height:40px;
  line-height:40px;
  border-radius:50%;
}
.ant-avatar-lg > * {
  line-height:40px;
}
.ant-avatar-lg.ant-avatar-icon {
  font-size:24px;
}
.ant-avatar-sm {
  width:24px;
  height:24px;
  line-height:24px;
  border-radius:50%;
}
.ant-avatar-sm > * {
  line-height:24px;
}
.ant-avatar-sm.ant-avatar-icon {
  font-size:14px;
}
.ant-avatar-square {
  border-radius:4px;
}
.ant-avatar > img {
  width:100%;
  height:100%;
  display:block;
}`
