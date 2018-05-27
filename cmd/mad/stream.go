package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"sync"
	"time"

	"github.com/mafredri/cdp/protocol/console"
	"github.com/mafredri/cdp/protocol/profiler"
	"github.com/mafredri/cdp/protocol/target"
	"github.com/mafredri/cdp/session"

	"github.com/gernest/mad/config"
	"github.com/gernest/mad/launcher"
	"github.com/mafredri/cdp"
	"github.com/mafredri/cdp/devtool"
	"github.com/mafredri/cdp/rpcc"
)

func streamResponse(ctx context.Context, cfg *config.Config, h respHandler) error {
	nctx, cancel := context.WithCancel(ctx)
	defer cancel()
	server := newServer(nctx, cfg)
	chrome, err := launcher.New(launcher.Options{
		Port:        cfg.DevtoolPort,
		ChromeFlags: []string{"--headless"},
	})
	if err != nil {
		return err
	}
	go chrome.Run()
	defer chrome.Stop()
	err = chrome.Wait(cfg.Verbose)
	if err != nil {
		return err
	}
	tabs := &sync.Map{}
	for _, v := range cfg.UnitFuncs {
		tabs.Store(v, true)
	}
	for _, v := range cfg.IntegrationFuncs {
		tabs.Store(v, true)
	}
	devt := devtool.New(fmt.Sprintf("%s:%d", cfg.DevtoolURL, cfg.DevtoolPort))
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
	m, err := session.NewManager(c)
	if err != nil {
		return err
	}
	defer m.Close()
	var pages []string
	pages = append(pages, cfg.UnitIndexPage)
	pages = append(pages, cfg.IntegrationIndexPages...)

	// We need a way to collect coverage profile before exiting the chrome tabs.
	// we use profileCtx to signal the tab execution goroutine to collect the
	// profiles.
	//
	// We then call <-profileCtx.Done() to trigger profile collection.
	profileCtx, cancelProfile := context.WithCancel(context.Background())

	// This is the channel we use to send the profile data collected from test
	// running tabs.
	profiles := make(chan profiler.Profile)
	for _, v := range pages {
		// Each test execution script is done in a separate tab. All unit tests are
		// compiled to a single execution script while each integration test is
		// compiled to a separate execution script.
		go func(idx string) {
			target, err := c.Target.CreateTarget(nctx,
				target.NewCreateTargetArgs(idx),
			)
			if err != nil {
				fmt.Printf("%s :%v\n", idx, err)
				return
			}
			pageConn, err := m.Dial(nctx, target.TargetID)
			if err != nil {
				fmt.Printf("%s :%v\n", idx, err)
				return
			}
			defer pageConn.Close()
			pageClient := cdp.NewClient(pageConn)
			if cfg.Verbose {
				if err = pageClient.Console.Enable(nctx); err != nil {
					fmt.Printf("%s :%v\n", idx, err)
					return
				}
				csLog, err := pageClient.Console.MessageAdded(nctx)
				if err != nil {
					fmt.Printf("%s :%v\n", idx, err)
					return
				}
				go func(cs console.MessageAddedClient) {
					for {
						msg, err := cs.Recv()
						if err != nil {
							return
						}
						fmt.Println(msg.Message.Text)
					}
				}(csLog)
			}
			if cfg.Cover {
				err := pageClient.Profiler.Enable(nctx)
				if err != nil {
					fmt.Printf("%s :%v\n", idx, err)
					return
				}
				err = pageClient.Profiler.Start(nctx)
				if err != nil {
					fmt.Printf("%s :%v\n", idx, err)
					return
				}
			}
			for {
				select {
				case <-ctx.Done():
					return
				case <-profileCtx.Done():
					s, err := pageClient.Profiler.Stop(nctx)
					if err != nil {
						fmt.Println(err)
						return
					}
					profiles <- s.Profile
					return
				}
			}
		}(v)
	}
	timeout := time.NewTimer(cfg.Timeout)
	defer timeout.Stop()
	for {
		select {
		case <-timeout.C:
			cancel()
		case <-nctx.Done():
			return ctx.Err()
		case ts := <-server:
			if h != nil {
				h.Handle(ts)
			}
			tabs.Delete(ts.Desc)
		default:
			complete := true
			tabs.Range(func(_, _ interface{}) bool {
				complete = false
				return false
			})
			if complete {
				if h != nil {
					h.Done()
				}
				if cfg.Cover {
					cancelProfile()
					n := 1
					var p []profiler.Profile
					for v := range profiles {
						p = append(p, v)
						if n == len(pages) {
							break
						}
						n++
					}
					data, _ := json.Marshal(p)
					err := ioutil.WriteFile(filepath.Join(cfg.OutputPath, cfg.Coverfile), data, 0600)
					if err != nil {
						fmt.Println(err)
					}
				}
				chrome.Stop()
				cancel()
				return nil
			}
		}
	}
}
