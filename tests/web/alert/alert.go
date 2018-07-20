package alert

import (
	"github.com/gernest/vected/lib/gs"
	"github.com/gernest/vected/lib/mad"
	"github.com/gernest/vected/web/alert"
)

func TestStyle() mad.Test {
	return mad.It("returns antd alert style", func(t mad.T) {
		css := gs.ToString(alert.Style())
		if css != expect {
			t.Error("got wrong styles")
		}
	})
}

const expect = `.ant-alert {
  font-family:"Monospaced Number","Chinese Quote", -apple-system, BlinkMacSystemFont, "Segoe UI", Roboto, "PingFang SC", "Hiragino Sans GB", "Microsoft YaHei", "Helvetica Neue", Helvetica, Arial, sans-serif;
  font-size:14px;
  line-height:1.5;
  color:rgba(0,0,0,0.65);
  box-sizing:border-box;
  margin:0;
  padding:0;
  list-style:none;
  position:relative;
  padding:8px 15px 8px 37px;
  border-radius:4px;
}
.ant-alert.ant-alert-no-icon {
  padding:8px 15px;
}
.ant-alert-icon {
  top:12.5px;
  left:16px;
  position:absolute;
}
.ant-alert-description {
  font-size:14px;
  line-height:22px;
  display:none;
}
.ant-alert-success {
  border:1px solid #b7eb8f;
  background-color:#f6ffed;
}
.ant-alert-success .ant-alert-icon {
  color:#52c41a;
}
.ant-alert-info {
  border:1px solid #91d5ff;
  background-color:#e6f7ff;
}
.ant-alert-info .ant-alert-icon {
  color:#178fff;
}
.ant-alert-warning {
  border:1px solid #ffe58f;
  background-color:#fffbe6;
}
.ant-alert-warning .ant-alert-icon {
  color:#faad14;
}
.ant-alert-error {
  border:1px solid #ffa39e;
  background-color:#fff1f0;
}
.ant-alert-error .ant-alert-icon {
  color:#f5222d;
}`
