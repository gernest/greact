package swing

import (
	"github.com/gernest/vected/lib/gs"
	"github.com/gernest/vected/lib/mad"
	"github.com/gernest/vected/web/style/motion/swing"
)

func TestSwingMotion() mad.Test {
	return mad.It("generate css for swing motion", func(t mad.T) {
		css := gs.ToString(swing.Motion(".swing", swing.Swing))
		if css != expectedMotion {
			t.Error("got wrong styles")
		}
	})
}

const expectedMotion = `.swing-enter, .swing-appear {
  animation-duration:.2s;
  animation-fill-mode::both;
  animation-play-state:paused;
}
.swing-enter.swing-enter-active, .swing-appear.swing-appear-active {
  animation-name:~swingIn;
  animation-play-state::running;
}`

func TestKeyframe() mad.Test {
	return mad.It("generate css for swing keyframes", func(t mad.T) {
		css := gs.ToString(swing.KeyFrames())
		expect := `@keyframes swingIn {
  0%,
  100% {
    transform:translateX(0);
  }
  20% {
    transform:translateX(-10px);
  }
  40% {
    transform:translateX(10px);
  }
  60% {
    transform:translateX(-5px);
  }
  80% {
    transform:translateX(5px);
  }
}`
		if css != expect {
			t.Error(css)
			t.Error("got wrong styles")
		}
	})
}
