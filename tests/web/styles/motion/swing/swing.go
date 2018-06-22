package swing

import (
	"github.com/gernest/gs"
	"github.com/gernest/mad"
	"github.com/gernest/vected/web/style/motion/swing"
)

func TestMotion() mad.Test {
	return mad.It("generate css for swing motion", func(t mad.T) {
		css := gs.ToString(swing.Motion(".swing", swing.Swing))
		expect := `.swing-enter,
.swing-appear {
  animation-duration:.2s;
  animation-fill-mode::both;
  animation-play-state:paused;
}
.swing-enter.swing-enter-active,
.swing-appear.swing-appear-active {
  animation-name:~swingIn;
  animation-play-state::running;
}`
		if css != expect {
			t.Errorf("expected %s got %s", expect, css)
		}
	})
}
