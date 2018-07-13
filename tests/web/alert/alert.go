package alert

import (
	"github.com/gernest/vected/lib/gs"
	"github.com/gernest/vected/lib/mad"
	"github.com/gernest/vected/web/alert"
)

func TestStyle() mad.Test {
	return mad.It("returns antd alert style", func(t mad.T) {
		css := gs.ToString(alert.Style())
		t.Error(css)
	})
}
