package move

import (
	"github.com/gernest/gs"
	"github.com/gernest/mad"
	"github.com/gernest/vected/web/style/motion/move"
)

func TestKeyFrames() mad.Test {
	return mad.It("generates css for move @keyframes", func(t mad.T) {
		css := gs.ToString(move.KeyFrames())
		expect := `@keyframes moveDownIn {
  0% {
    transform-origin:0 0;
    transform:translateY(100%);
    opacity:0;
  }
  100% {
    transform-origin:0 0;
    transform:translateY(0%);
    opacity:1;
  }
}
@keyframes moveDownOut {
  0% {
    transform-origin:0 0;
    transform:translateY(100%);
    opacity:1;
  }
  100% {
    transform-origin:0 0;
    transform:translateY(0%);
    opacity:0;
  }
}
@keyframes moveLeftIn {
  0% {
    transform-origin:0 0;
    transform:translateX(-100%);
    opacity:0;
  }
  100% {
    transform-origin:0 0;
    transform:translateX(0%);
    opacity:1;
  }
}
@keyframes moveLeftOut {
  0% {
    transform-origin:0 0;
    transform:translateX(0%);
    opacity:1;
  }
  100% {
    transform-origin:0 0;
    transform:translateX(-100%);
    opacity:0;
  }
}
@keyframes moveRightIn {
  0% {
    transform-origin:0 0;
    transform:translateX(100%);
    opacity:0;
  }
  100% {
    transform-origin:0 0;
    transform:translateX(0%);
    opacity:1;
  }
}
@keyframes moveRightOut {
  0% {
    transform-origin:0 0;
    transform:translateX(0%);
    opacity:1;
  }
  100% {
    transform-origin:0 0;
    transform:translateX(100%);
    opacity:0;
  }
}
@keyframes moveUpIn {
  0% {
    transform-origin:0 0;
    transform:translateY(-100%);
    opacity:0;
  }
  100% {
    transform-origin:0 0;
    transform:translateY(0%);
    opacity:1;
  }
}
@keyframes moveUpOut {
  0% {
    transform-origin:0 0;
    transform:translateY(0%);
    opacity:1;
  }
  100% {
    transform-origin:0 0;
    transform:translateY(-100%);
    opacity:0;
  }
}`
		if css != expect {
			t.Errorf("expected %s got %s", expect, css)
		}
	})
}
