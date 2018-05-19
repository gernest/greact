package main

import (
	"context"
	"time"

	"github.com/mafredri/cdp"
	"github.com/mafredri/cdp/devtool"
	"github.com/mafredri/cdp/protocol/page"
	"github.com/mafredri/cdp/rpcc"
)

func newBrowser(ctx context.Context, index string, timeout time.Duration) error {
	devt := devtool.New("http://127.0.0.1:9222")
	pt, err := devt.Get(ctx, devtool.Page)
	if err != nil {
		return err
	}
	conn, err := rpcc.DialContext(ctx, pt.WebSocketDebuggerURL)
	if err != nil {
		return err
	}
	defer conn.Close()
	c := cdp.NewClient(conn)
	navArgs := page.NewNavigateArgs(index)
	_, err = c.Page.Navigate(ctx, navArgs)
	if err != nil {
		return err
	}
	<-ctx.Done()
	return ctx.Err()
}
