package main

import (
	"context"
	"fmt"
	"net"
	"net/url"

	"github.com/gernest/mad"
	"github.com/gernest/mad/config"

	"github.com/gorilla/websocket"
)

// Opens a websocket connection using ws as url and reads the received messages
// as json of type *api.TestSuite.
//
// If handler is not nil, for every successful read the handler will be invoked
// passing the decoded *api.TestSuite as argument.
func streamResponse(ctx context.Context, cfg *config.Config, ws string, handler func(*mad.SpecResult)) error {
	u, err := url.Parse(ws)
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
	fmt.Printf("socket %s\n", ws)
	conn, _, err := websocket.NewClient(d, u, nil, 1024, 1024)
	if err != nil {
		return err
	}
	for {
		select {
		case <-ctx.Done():
			return conn.Close()
		default:
			if len(m) == 0 {
				return nil
			}
			ts := &mad.SpecResult{}
			err := conn.ReadJSON(ts)
			if err != nil {
				return err
			}
			if handler != nil {
				handler(ts)
			}
			delete(m, ts.Desc)
		}
	}
}
