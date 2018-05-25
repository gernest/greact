package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path/filepath"

	"github.com/gernest/mad/config"
	"github.com/gernest/mad/launcher"
	"github.com/mafredri/cdp"
	"github.com/mafredri/cdp/devtool"
	"github.com/mafredri/cdp/protocol/page"
	"github.com/mafredri/cdp/rpcc"
)

func streamResponse(ctx context.Context, cfg *config.Config, h respHandler) error {
	nctx, cancel := context.WithCancel(ctx)
	defer cancel()
	server := NewServer(nctx, cfg)
	chrome, err := launcher.New(launcher.Options{
		Port:        9222,
		ChromeFlags: []string{"--headless"},
	})
	if err != nil {
		return err
	}
	go chrome.Run()
	defer chrome.Stop()
	err = chrome.Wait()
	if err != nil {
		return err
	}
	m := make(map[string]bool)
	for _, v := range cfg.UnitFuncs {
		m[v] = true
	}
	devt := devtool.New(fmt.Sprintf("http://127.0.0.1:%d", 9222))
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
		err = c.Profiler.Enable(nctx)
		if err != nil {
			return err
		}
		err = c.Profiler.Start(nctx)
		if err != nil {
			return err
		}
	}
	navArgs := page.NewNavigateArgs(fmt.Sprintf("http://localhost:%d",
		cfg.Port) + resourcePath + "?src=index.html")
	_, err = c.Page.Navigate(ctx, navArgs)
	if err != nil {
		return err
	}
	for {
		select {
		case <-nctx.Done():
			return ctx.Err()
		case ts := <-server:

			if h != nil {
				h.Handle(ts)
			}
			delete(m, ts.Desc)
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
				chrome.Stop()
				cancel()
				return nil
			}
		}
	}
}
