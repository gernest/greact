package backtop

import (
	"github.com/gernest/vected/lib/gs"
	"github.com/gernest/vected/lib/mad"
	"github.com/gernest/vected/web/backtop/style"
)

func TestStyle() mad.Test {
	return mad.It("generate backtop css", func(t mad.T) {
		css := gs.ToString(style.BackTop())
		if css != expect {
			t.Error("got wrong styles")
		}
	})
}

const expect = `.ant-back-top {
  font-family:"Monospaced Number","Chinese Quote", -apple-system, BlinkMacSystemFont, "Segoe UI", Roboto, "PingFang SC", "Hiragino Sans GB", "Microsoft YaHei", "Helvetica Neue", Helvetica, Arial, sans-serif;
  font-size:14px;
  line-height:1.5;
  color:rgba(0,0,0,0.65);
  box-sizing:border-box;
  margin:0;
  padding:0;
  list-style:none;
  z-index:10;
  position:fixed;
  right:100px;
  bottom:50px;
  height:40px;
  width:40px;
  cursor:pointer;
}
.ant-back-top-content {
  height:40px;
  width:40px;
  border-radius:20px;
  background-color:rgba(0,0,0,0.45);
  color:#ffffff;
  text-align:center;
  transition:all .3s cubic-bezier(0.645, 0.045, 0.355, 1);
  overflow:hidden;
}
.ant-back-top-content:hover {
  background-color:rgba(0,0,0,0.65);
  transition:all .3s cubic-bezier(0.645, 0.045, 0.355, 1);
}
.ant-back-top-icon {
  margin:12px auto;
  width:14px;
  height:16px;
  background:url(data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAACQAAAAoCAYAAACWwljjAAAABGdBTUEAALGPC/xhBQAAAbtJREFUWAntmMtKw0AUhhMvS5cuxILgQlRUpIggIoKIIoigG1eC+AA+jo+i6FIXBfeuXIgoeKVeitVWJX5HWhhDksnUpp3FDPyZk3Nm5nycmZKkXhAEOXSA3lG7muTeRzmfy6HneUvIhnYkQK+Q9NhAA0Opg0vBEhjBKHiyb8iGMyQMOYuK41BcBSypAL+MYXSKjtFAW7EAGEO3qN4uMQbbAkXiSfRQJ1H6a+yhlkKRcAoVFYiweYNjtCVQJJpBz2GCiPt7fBOZQpFgDpUikse5HgnkM4Fi4QX0Fpc5wf9EbLqpUCy4jMoJSXWhFwbMNgWKhVbRhy5jirhs9fy/oFhgHVVTJEs7RLZ8sSEoJm6iz7SZDMbJ+/OKERQTttCXQRLToRUmrKWCYuA2+jbN0MB4OQobYShfdTCgn/sL1K36M7TLrN3n+758aPy2rrpR6+/od5E8tf/A1uLS9aId5T7J3CNYihkQ4D9PiMdMC7mp4rjB9kjFjZp8BlnVHJBuO1yFXIV0FdDF3RlyFdJVQBdv5AxVdIsq8apiZ2PyYO1EVykesGfZEESsCkweyR8MUW+V8uJ1gkYipmpdP1pm2aJVPEGzAAAAAElFTkSuQmCC) ~"100%/100%" no-repeat
  ;
}
@media screen and (max-width: @screen-md)  {
  .ant-back-top {
    right:60px;
  }
}
@media screen and (max-width: @screen-xs) {
  .ant-back-top {
    right:20px;
  }
}`
