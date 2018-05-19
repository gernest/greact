package main

import (
	"context"
	"fmt"
	"net"
	"net/url"
	"time"

	"github.com/gernest/mad"
	"github.com/gernest/mad/api"
	"github.com/gernest/mad/config"

	"github.com/gorilla/websocket"
)

// Opens a websocket connection using ws as url and reads the received messages
// as json of type *api.TestSuite.
//
// If handler is not nil, for every successful read the handler will be invoked
// passing the decoded *api.TestSuite as argument.
func streamResponse(ctx context.Context, cfg *config.Config, res *api.TestResponse, h respHandler) error {
	nctx, cancel := context.WithCancel(ctx)
	go func() {
		err := newBrowser(nctx, res.IndexURL, 30*time.Second)
		if err != nil {
			fmt.Println(err)
		}
	}()
	u, err := url.Parse(res.WebsocketURL)
	if err != nil {
		return err
	}
	// u.Scheme = "tcp"
	m := make(map[string]bool)
	for _, v := range cfg.UnitFuncs {
		m[v] = true
	}
	d, err := net.Dial("tcp", u.Host)
	if err != nil {
		return err
	}
	conn, _, err := websocket.NewClient(d, u, nil, 1024, 1024)
	if err != nil {
		return err
	}
	for {
		select {
		case <-ctx.Done():
			cancel()
			return conn.Close()
		default:
			if len(m) == 0 {
				if h != nil {
					h.Done()
					cancel()
				}
				return nil
			}
			ts := &mad.SpecResult{}
			err := conn.ReadJSON(ts)
			if err != nil {
				return err
			}
			if h != nil {
				h.Handle(ts)
			}
			delete(m, ts.Desc)
		}
	}
}
