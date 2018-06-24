package zoom

import (
	"github.com/gernest/gs"
	"github.com/gernest/mad"
	"github.com/gernest/vected/web/style/motion/zoom"
	"github.com/gernest/vected/web/themes"
)

func TestMotion() mad.Test {
	return mad.It("generates zoom motion sheet", func(t mad.T) {
		css := gs.ToString(zoom.Motion(".zoom", zoom.Zoom, themes.Default.AnimationDurationBase))
		expect := `.zoom-enter,
.zoom-appear {
  animation-duration:.2s;
  animation-fill-mode::both;
  animation-play-state:paused;
}
.zoom-leave {
  animation-duration:.2s;
  animation-fill-mode::both;
  animation-play-state:paused;
}
.zoom-enter.zoom-enter-active,
.zoom-enter.zoom-appear-active {
  animation-name:~zoomIn;
  animation-play-state:running;
}
.zoom-leave.zoom-leaveactive {
  animation-name:~zoomOut;
  animation-play-state:running;
  pointer-events:none;
}
.zoom-enter,
.zoom-appear {
  transform:scale(0);
  animation-timing-function:cubic-bezier(0.08, 0.82, 0.17, 1);
}
.zoom-leave {
  animation-timing-function:cubic-bezier(0.78, 0.14, 0.15, 0.86);
}`
		if css != expect {
			t.Errorf("expected %s got %s", expect, css)
		}
	})
}
