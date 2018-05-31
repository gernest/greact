package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"
	"sync"
	"time"

	"github.com/mafredri/cdp/protocol/console"
	"github.com/mafredri/cdp/protocol/target"
	"github.com/mafredri/cdp/session"

	"github.com/gernest/mad/config"
	"github.com/gernest/mad/cover"
	"github.com/gernest/mad/launcher"
	"github.com/mafredri/cdp"
	"github.com/mafredri/cdp/devtool"
	"github.com/mafredri/cdp/rpcc"
)

// streamResponse runs the compiled tests in a web browser and displays the
// results to stdout.
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

	// keep track on which functions have been executed. The trick here is, we
	// first store all functions which are eligible for execution, then whenever we
	// receive report for a successful function execution, we remove the function
	// from this map until we have no function left on the map.
	//
	// So, by checking if there is no more keys stored in the map we will be sure
	// that the execution of all tests was complete.
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
	profiles := make(chan []cover.Profile)
	// This is the channel we use to send the profile data collected from test
	// running tabs.
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
			if cfg.Verbose || cfg.Cover {
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
						if cfg.Cover {
							txt := msg.Message.Text
							if strings.HasPrefix(txt, cover.Key) {
								txt := strings.TrimPrefix(txt, cover.Key)
								prof := []cover.Profile{}
								err := json.Unmarshal([]byte(txt), &prof)
								if err != nil {
									fmt.Println(err)
								} else {
									profiles <- prof
								}
							}
						}
					}
				}(csLog)
			}
			for {
				select {
				case <-ctx.Done():
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
			// Do not keep tabs on this function. This function was successful executed.
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
					totalProfiles := len(cfg.IntegrationFuncs)
					if len(cfg.UnitFuncs) > 0 {

						// All unit functions are executed in a single package. Which means they
						// will only give one profile.
						totalProfiles++
					}
					count := 1
					var collect []cover.Profile
					for p := range profiles {
						if collect != nil {
							collect = mergeProfiles(collect, p)
						} else {
							collect = p
						}
						if count == totalProfiles {
							break
						}
						count++
					}
					if cfg.Coverfile != "" {
						b, _ := json.Marshal(collect)
						err := ioutil.WriteFile(cfg.Coverfile, b, 0600)
						if err != nil {
							fmt.Println(err)
						}
					}
					printCoverage(collect)
				}
				chrome.Stop()
				cancel()
				return nil
			}
		}
	}
}
