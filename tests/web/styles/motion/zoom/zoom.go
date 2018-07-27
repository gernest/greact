package zoom

import (
	"github.com/gernest/vected/lib/gs"
	"github.com/gernest/vected/lib/mad"
	"github.com/gernest/vected/web/style/core/themes"
	"github.com/gernest/vected/web/style/motion/zoom"
)

func TestZoomMotion() mad.Test {
	return mad.It("generates zoom motion sheet", func(t mad.T) {
		css := gs.ToString(zoom.Motion(".zoom", zoom.Zoom, themes.Default.AnimationDurationBase))
		expect := `.zoom-enter, .zoom-appear {
  animation-duration:.2s;
  animation-fill-mode::both;
  animation-play-state:paused;
}
.zoom-leave {
  animation-duration:.2s;
  animation-fill-mode::both;
  animation-play-state:paused;
}
.zoom-enter.zoom-enter-active, .zoom-appear.zoom-appear-active {
  animation-name:~zoomIn;
  animation-play-state:running;
}
.zoom-leave.zoom-leave-active {
  animation-name:~zoomOut;
  animation-play-state:running;
  pointer-events:none;
}
.zoom-enter, .zoom-appear {
  transform:scale(0);
  animation-timing-function:cubic-bezier(0.08, 0.82, 0.17, 1);
}
.zoom-leave {
  animation-timing-function:cubic-bezier(0.78, 0.14, 0.15, 0.86);
}`
		if css != expect {
			t.Error("got wrong styles")
		}
	})
}

func TestKeyFrames() mad.Test {
	return mad.It("generates @keyframe styles for zoom", func(t mad.T) {
		css := gs.ToString(zoom.KeyFrames())
		expect := `@keyframes zoomIn {
  0% {
    opacity:0;
    transform:scale(0.2);
  }
  100% {
    opacity:1;
    transform:scale(1);
  }
}
@keyframes zoomOut {
  0% {
    transform:scale(1);
  }
  100% {
    opacity:0;
    transform:scale(0.2);
  }
}
@keyframes zoomBigIn {
  0% {
    opacity:0;
    transform:scale(.8);
  }
  100% {
    transform:scale(1);
  }
}
@keyframes zoomBigOut {
  0% {
    transform:scale(1);
  }
  100% {
    opacity:0;
    transform:scale(.8);
  }
}
@keyframes zoomUpIn {
  0% {
    opacity:0;
    transform-origin:50% 0%;
    transform:scale(.8);
  }
  100% {
    transform-origin:50% 0%;
    transform:scale(1);
  }
}
@keyframes zoomUpOut {
  0% {
    transform-origin:50% 0%;
    transform:scale(1);
  }
  100% {
    opacity:0;
    transform-origin:50% 0%;
    transform:scale(.8);
  }
}
@keyframes zoomLeftIn {
  0% {
    opacity:0;
    transform-origin:0% 50%;
    transform:scale(.8);
  }
  100% {
    transform-origin:0% 50%;
    transform:scale(1);
  }
}
@keyframes zoomLeftOut {
  0% {
    transform-origin:0% 50%;
    transform:scale(1);
  }
  100% {
    opacity:0;
    transform-origin:0% 50%;
    transform:scale(.8);
  }
}
@keyframes zoomRightIn {
  0% {
    opacity:0;
    transform-origin:100% 50%;
    transform:scale(.8);
  }
  100% {
    transform-origin:100% 50%;
    transform:scale(1);
  }
}
@keyframes zoomRightOut {
  0% {
    transform-origin:100% 50%;
    transform:scale(1);
  }
  100% {
    opacity:0;
    transform-origin:100% 50%;
    transform:scale(.8);
  }
}
@keyframes zoomDownIn {
  0% {
    opacity:0;
    transform-origin:50% 100%;
    transform:scale(.8);
  }
  100% {
    transform-origin:50% 100%;
    transform:scale(1);
  }
}
@keyframes zoomDownOut {
  0% {
    transform-origin:50% 100%;
    transform:scale(1);
  }
  100% {
    opacity:0;
    transform-origin:50% 100%;
    transform:scale(.8);
  }
}`
		if css != expect {
			t.Error("wrong styles")
		}
	})
}
