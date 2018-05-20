package main

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net"
	"net/url"
	"path/filepath"

	"github.com/gernest/mad"
	"github.com/gernest/mad/api"
	"github.com/gernest/mad/config"
	"github.com/mafredri/cdp"
	"github.com/mafredri/cdp/devtool"
	"github.com/mafredri/cdp/protocol/page"
	"github.com/mafredri/cdp/rpcc"

	"github.com/gorilla/websocket"
)

// Opens a websocket connection using ws as url and reads the received messages
// as json of type *api.TestSuite.
//
// If handler is not nil, for every successful read the handler will be invoked
// passing the decoded *api.TestSuite as argument.
func streamResponse(ctx context.Context, cfg *config.Config, res *api.TestResponse, h respHandler) error {
	u, err := url.Parse(res.WebsocketURL)
	if err != nil {
		return err
	}
	m := make(map[string]bool)
	for _, v := range cfg.UnitFuncs {
		m[v] = true
	}
	d, err := net.Dial("tcp", u.Host)
	if err != nil {
		return err
	}
	ws, _, err := websocket.NewClient(d, u, nil, 1024, 1024)
	if err != nil {
		return err
	}
	defer ws.Close()
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
	if cfg.Cover {
		err = c.Profiler.Enable(ctx)
		if err != nil {
			return err
		}
		err = c.Profiler.Start(ctx)
		if err != nil {
			return err
		}
	}
	navArgs := page.NewNavigateArgs(res.IndexURL)
	_, err = c.Page.Navigate(ctx, navArgs)
	if err != nil {
		return err
	}
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			if len(m) == 0 {
				if h != nil {
					h.Done()
				}
				if cfg.Cover {
					s, err := c.Profiler.Stop(ctx)
					if err != nil {
						return err
					}
					b, _ := json.Marshal(s.Profile)
					err = ioutil.WriteFile(
						filepath.Join(cfg.OutputPath, cfg.Coverfile), b, 0600)
					if err != nil {
						return err
					}
				}

				return nil
			}
			ts := &mad.SpecResult{}
			err = ws.ReadJSON(ts)
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
