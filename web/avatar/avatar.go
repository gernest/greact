package avatar

import (
	"github.com/gernest/vected"
	"github.com/gernest/vected/lib/props"
	"github.com/gernest/vected/lib/state"
)

var _ vected.InitState = (*Avatar)(nil)
var _ vected.InitProps = (*Avatar)(nil)

// Avatar is a vected component for antd avatar.
type Avatar struct {
	vected.Core
}

// InitState implements vected.InitState interface. This provides initial state
// values of the avatar component.
func (*Avatar) InitState() state.State {
	return state.State{
		"scale":      1,
		"isImgExist": true,
	}
}

// InitProps returns default props.
func (*Avatar) InitProps() props.Props {
	return props.Props{
		"prefixCls": "ant-avatar",
		"shape":     "circle",
		"size":      "default",
	}
}
