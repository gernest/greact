package other

import (
	"github.com/gernest/vected/lib/gs"
	"github.com/gernest/vected/lib/mad"
	"github.com/gernest/vected/web/style/motion/other"
)

func TestLoadingCircle() mad.Test {
	return mad.It("generates loadingCircle @keyframes", func(t mad.T) {
		css := gs.ToString(other.LoadingCircle())
		expect := `@keyframes loadingCircle {
  0% {
    transform-origin:50% 50%;
    transform:rotate(0deg);
  }
  100% {
    transform-origin:50% 50%;
    transform:rotate(360deg);
  }
}`
		if css != expect {
			t.Errorf("expected %s got %s", expect, css)
		}
	})
}
