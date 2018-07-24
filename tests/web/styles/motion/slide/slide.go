package slide

import (
	"github.com/gernest/vected/lib/gs"
	"github.com/gernest/vected/lib/mad"
	"github.com/gernest/vected/web/style/motion/slide"
)

func TestSlideStyle() mad.Test {
	return mad.It("generates css for slide @keyframes", func(t mad.T) {
		css := gs.ToString(slide.KeyFrames())
		expect := `@keyframes slideUpIn {
  0% {
    opacity:0;
    transform-origin:0% 0%;
    transform:scaleY(.8);
  }
  100% {
    opacity:1;
    transform-origin:0% 0%;
    transform:scaleY(1);
  }
}
@keyframes slideUpOut {
  0% {
    opacity:1;
    transform-origin:0% 0%;
    transform:scaleY(1);
  }
  100% {
    opacity:0;
    transform-origin:0% 0%;
    transform:scaleY(.8);
  }
}
@keyframes slideDownIn {
  0% {
    opacity:0;
    transform-origin:100% 100%;
    transform:scaleY(.8);
  }
  100% {
    opacity:1;
    transform-origin:100% 100%;
    transform:scaleY(1);
  }
}
@keyframes slideDownOut {
  0% {
    opacity:1;
    transform-origin:100% 100%;
    transform:scaleY(1);
  }
  100% {
    opacity:0;
    transform-origin:100% 100%;
    transform:scaleY(.8);
  }
}
@keyframes slideLeftIn {
  0% {
    opacity:0;
    transform-origin:0% 0%;
    transform:scaleX(.8);
  }
  100% {
    opacity:1;
    transform-origin:0% 0%;
    transform:scaleX(1);
  }
}
@keyframes slideLeftOut {
  0% {
    opacity:1;
    transform-origin:0% 0%;
    transform:scaleX(1);
  }
  100% {
    opacity:0;
    transform-origin:0% 0%;
    transform:scaleX(.8);
  }
}
@keyframes slideLeftIn {
  0% {
    opacity:0;
    transform-origin:100% 0%;
    transform:scaleX(.8);
  }
  100% {
    opacity:1;
    transform-origin:100% 0%;
    transform:scaleX(1);
  }
}
@keyframes slideLeftOut {
  0% {
    opacity:1;
    transform-origin:100% 0%;
    transform:scaleX(1);
  }
  100% {
    opacity:0;
    transform-origin:100% 0%;
    transform:scaleX(.8);
  }
}`
		if css != expect {
			t.Error(css)
			// t.Errorf("expected %s got %s", expect, css)
		}
	})
}

func TestMotion() mad.Test {
	return mad.It("generates css for slide motion", func(t mad.T) {
		css := gs.ToString(slide.Motion(".slide-up", slide.Up))
		expect := `.slide-up-enter,
.slide-up-appear {
  animation-duration:.2s;
  animation-fill-mode::both;
  animation-play-state:paused;
}
.slide-up-leave {
  animation-duration:.2s;
  animation-fill-mode::both;
  animation-play-state:paused;
}
.slide-up-enter.slide-up-enter-active,
.slide-up-enter.slide-up-appear-active {
  animation-name:~slideUpIn;
  animation-play-state:running;
}
.slide-up-leave.slide-up-leaveactive {
  animation-name:~slideUpOut;
  animation-play-state:running;
  pointer-events:none;
}
.slide-up-enter,
.slide-up-appear {
  opacity:0;
  animation-timing-function:cubic-bezier(0.23, 1, 0.32, 1);
}
.slide-up-leave {
  animation-timing-function:cubic-bezier(0.755, 0.05, 0.855, 0.06);
}`
		if css != expect {
			t.Errorf("expected %s got %s", expect, css)
		}
	})
}
