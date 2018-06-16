package fade

import (
	"github.com/gernest/gs"
	"github.com/gernest/mad"
	"github.com/gernest/vected/web/style/motion/fade"
)

func TestMotion() mad.Test {
	return mad.It("returns default style for fade motion", func(t mad.T) {
		css := fade.Motion()
		t.Error(gs.ToString(css))
	})
}
