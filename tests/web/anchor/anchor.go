package anchor

import (
	"github.com/gernest/vected/lib/gs"
	"github.com/gernest/vected/lib/mad"
	"github.com/gernest/vected/web/anchor/style"
)

func TestStyle() mad.Test {
	return mad.It("generates css for anchor", func(t mad.T) {
		css := gs.ToString(style.Anchor())
		if css != expect {
			t.Error("got wrong styles")
		}
	})
}

const expect = `.ant-anchor {
  font-family:"Monospaced Number","Chinese Quote", -apple-system, BlinkMacSystemFont, "Segoe UI", Roboto, "PingFang SC", "Hiragino Sans GB", "Microsoft YaHei", "Helvetica Neue", Helvetica, Arial, sans-serif;
  font-size:14px;
  line-height:1.5;
  color:rgba(0,0,0,0.65);
  box-sizing:border-box;
  margin:0;
  padding:0;
  list-style:none;
  position:relative;
  padding-left:2px;
}
.ant-anchor-wrapper {
  background-color:#ffffff;
  overflow:auto;
  padding-left:4px;
  margin-left:-4px;
}
.ant-anchor-ink {
  position:absolute;
  height:100%;
  left:0;
  top:0;
}
.ant-anchor-ink:before {
  content:' ';
  position:relative;
  width:2px;
  height:100%;
  display:block;
  background-color:#e8e8e8;
  margin:auto;
}
.ant-anchor-ball {
  display:none;
  position:absolute;
  width:8px;
  height:8px;
  border-radius:8px;
  border:2px solid #1890ff;
  background-color:#ffffff;
  left:50%;
  transition:top .3s ease-in-out;
  transform:translateX(-50%);
}
.ant-anchor-ball.visible {
  display:inline-block;
}
.ant-anchor.fixed .ant-anchor-ink .ant-anchor-ink-ball {
  display:none;
}
.ant-anchor-link {
  padding:8px 0 8px 16px;
  line-height:1;
}
.ant-anchor-link-title {
  display:block;
  position:relative;
  transition:all .3s;
  color:rgba(0,0,0,0.65);
  white-space:nowrap;
  overflow:hidden;
  text-overflow:ellipsis;
  margin-bottom:8px;
}
.ant-anchor-link-title:only-child {
  margin-bottom:0;
}
.ant-anchor-link-active > .ant-anchor-link-title {
  color:#1890ff;
}
.ant-anchor-link .ant-anchor-link {
  padding-top:6px;
  padding-bottom:6px;
}`
