package launcher

import "context"

// Browser this is an interface for launching a browser which supports chrome
// debugging protocol.
type Browser interface {
	Run(context.Context) error
	Ready() error
	Stop() error
}
