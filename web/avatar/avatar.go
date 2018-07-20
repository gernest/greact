package avatar

import (
	"github.com/gernest/vected"
	"github.com/gernest/vected/lib/props"
)

var _ vected.InitState = (*Avatar)(nil)

// Avatar is a vected component for antd avatar.
type Avatar struct {
	vected.Core
}

// InitState implements vected.InitState interface.
func (*Avatar) InitState(_ props.Props) map[string]interface{} {
	return map[string]interface{}{
		"scale":      1,
		"isImgExist": true,
	}
}
