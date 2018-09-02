package cmp

import (
	"github.com/gernest/vected"
)

type SimpleComponent struct {
	vected.Core
}

func (s SimpleComponent) Template() string {
	return `
	<div>
		<span>
			{props.String("initialName")}/{s.State().String("name")}
		</span>
	</div>
	`
}
