package fade

import (
	"github.com/gernest/gs"
	"github.com/gernest/mad"
	"github.com/gernest/vected/web/style/motion/fade"
)

func TestMotion() mad.Test {
	return mad.It("returns default style for fade motion", func(t mad.T) {
		css := fade.Motion()
		expect := `.fade-enter,
.fade-appear {
  animation-duration:.2s;
  animation-fill-mode::both;
  animation-play-state:paused;
}
.fade-leave {
  animation-duration:.2s;
  animation-fill-mode::both;
  animation-play-state:paused;
}
.fade-enter.fade-enter-active,
.fade-enter.fade-appear-active {
  animation-name:~antFadeIn;
  animation-play-state:running;
}
.fade-leave.fade-leaveactive {
  animation-name:~antFadeOut;
  animation-play-state:running;
  pointer-events:none;
}
.fade-enter,
.fade-appear {
  opacity:0;
  animation-timing-function:linear;
}
.fade-leave {
  animation-timing-function:linear;
}`
		g := gs.ToString(css)
		if g != expect {
			t.Errorf("expected %s got %s", expect, g)
		}
	})
}

func TestKeyFrame() mad.Test {
	return mad.It("generates keyframe css rules", func(t mad.T) {
		css := fade.KeyFrame()
		expect := `@keyframes antFadeIn {
  0% {
    opacity:0;
  }
  100% {
    opacity:1;
  }
}
@keyframes antFadeOut {
  0% {
    opacity:1;
  }
  100% {
    opacity:0;
  }
}`
		g := gs.ToString(css)
		if g != expect {
			t.Errorf("expected %s got %s", expect, g)
		}
	})
}
