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
	m, err := session.NewManager(c)
	if err != nil {
		return err
	}
	defer m.Close()
	var pages []string
	pages = append(pages, cfg.UnitIndexPage)
	pages = append(pages, cfg.IntegrationIndexPages...)
	for _, v := range pages {
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
			<-nctx.Done()
		}(v)
	}
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
	timeout := time.NewTimer(30 * time.Second)
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
