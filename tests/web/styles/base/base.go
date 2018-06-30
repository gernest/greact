package base

import (
	"github.com/gernest/gs"
	"github.com/gernest/mad"
	"github.com/gernest/vected/web/style/core/base"
)

func TestBase() mad.Test {
	return mad.It("generates normalize css", func(t mad.T) {
		css := gs.ToString(base.Base())
		t.Error(css)
	})
}
