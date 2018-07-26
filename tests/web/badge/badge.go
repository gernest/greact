package badge

import (
	"github.com/gernest/vected/lib/gs"
	"github.com/gernest/vected/lib/mad"
	"github.com/gernest/vected/web/badge/style"
)

func TestBadgeStyle() mad.Test {
	return mad.It("generates antd badge css", func(t mad.T) {
		css := gs.ToString(style.Badge())
		if css != expect {
			t.Error("got wrong styles")
		}
	})
}

const expect = `.ant-badge {
  font-family:"Monospaced Number","Chinese Quote", -apple-system, BlinkMacSystemFont, "Segoe UI", Roboto, "PingFang SC", "Hiragino Sans GB", "Microsoft YaHei", "Helvetica Neue", Helvetica, Arial, sans-serif;
  font-size:14px;
  line-height:1.5;
  color:rgba(0,0,0,0.65);
  box-sizing:border-box;
  margin:0;
  padding:0;
  list-style:none;
  position:relative;
  display:inline-block;
  line-height:1;
  vertical-align:middle;
}
.ant-badge-count  {
  position:absolute;
  transform:translateX(-50%);
  top:-10px;
  height:20px;
  border-radius:10px;
  min-width:20px;
  background:#f5222d;
  color:#fff;
  line-height:20px;
  text-align:center;
  padding:0 6px;
  font-size:12px;
  font-weight:normal;
  white-space:nowrap;
  transform-origin:-10% center;
  box-shadow:0 0 0 1px #fff;
}
.ant-badge-count  a, .ant-badge-count  a:hover {
  color:#fff;
}
.ant-badge-multiple-words {
  padding:0 8px;
}
.ant-badge-dot {
  position:absolute;
  transform:translateX(-50%);
  transform-origin:0 center;
  top:-3px;
  height:6px;
  width:6px;
  border-radius:100%;
  background:#f5222d;
  z-index:10;
  box-shadow:0 0 0 1px #fff;
}
.ant-badge-status {
  ine-height:inherit;
  vertical-align:baseline;
}
.ant-badge-status-dot {
  width:6px;
  height:6px;
  display:inline-block;
  border-radius:50%;
  vertical-align:middle;
  position:relative;
  top:-1px;
}
.ant-badge-status-success {
  background:#52c41a;
}
.ant-badge-status-processing {
  background-color:#1890ff;
  position:relative;
}
.ant-badge-status-processing:after {
  position:absolute;
  top:0;
  left:0;
  width:100%;
  height:100%;
  border-radius:50%;
  border:1px solid #1890ff;
  content:'';
  animation:antStatusProcessing 1.2s infinite ease-in-out;
}
.ant-badge-status-default {
  background-color:#d9d9d9;
}
.ant-badge-status-error {
  background-color:#f5222d;
}
.ant-badge-status-warning {
  background-color:#faad14;
}
.ant-badge-status-text {
  color:rgba(0,0,0,0.65);
  font-size:14px;
  margin-left:8px;
}
.ant-badge-zoom-appear {
  animation:antZoomBadgeIn .3s cubic-bezier(0.12, 0.4, 0.29, 1.46);
  animation-fill-mode:both;
}
.ant-badge-zoom-enter {
  animation:antZoomBadgeIn .3s cubic-bezier(0.12, 0.4, 0.29, 1.46);
  animation-fill-mode:both;
}
.ant-badge-zoom-leave {
  animation:antZoomBadgeIn .3s cubic-bezier(0.71, -0.46, 0.88, 0.6);
  animation-fill-mode:both;
}
.ant-badge-not-a-wrapper .ant-scroll-number {
  top:auto;
  display:block;
  position:relative;
}
 .ant-badge-not-a-wrapper .ant-badge-count {
  transform:none;
}
@keyframes antStatusProcessing {
  0% {
    transform:scale(0.8);
    opacity:0.5;
  }
  100% {
    transform:scale(2.4);
    opacity:0;
  }
}
.ant-scroll-number {
  overflow:hidden;
}
.ant-scroll-number-only {
  display:inline-block;
  transition:all .3s cubic-bezier(0.645, 0.045, 0.355, 1);
  height:20px;
}
.ant-scroll-number-only,
> p {
  height:20px;
  margin:0;
}
@keyframes antZoomBadgeIn {
  0% {
    opacity:0;
    transform:scale(0) translateX(-50%);
  }
  100% {
    transform:scale(1) translateX(-50%);
  }
}
@keyframes antZoomBadgeOut {
  0% {
    transform:scale(1) translateX(-50%);
  }
  100% {
    opacity:0;
    transform:scale(0) translateX(-50%);
  }
}`
