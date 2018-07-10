package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"
	"sync"

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
func streamResponse(ctx context.Context, cfg *config.Config, browser launcher.Browser, h respHandler) error {
	nctx, cancel := context.WithCancel(ctx)
	defer cancel()
	server := newServer(nctx, cfg)
	go browser.Run(nctx)
	defer browser.Stop()
	err := browser.Ready()
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
	totalProfiles := 0
	hasUnitTest := false
	for k, v := range cfg.TestNames {
		for _, fn := range v.Unit {
			if !hasUnitTest {
				hasUnitTest = true
			}
			tabs.Store(k.FormatName(fn), true)
		}
		for _, fn := range v.Integration {
			tabs.Store(k.FormatName(fn), true)
			totalProfiles++
		}
	}
	if hasUnitTest {
		totalProfiles++
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
			err := executeInTab(nctx, cfg, idx, c, m, profiles)
			if err != nil {
				fmt.Println(err)
			}
		}(v)
	}

	for {
		select {
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
						ext := filepath.Ext(cfg.Coverfile)
						if ext == ".json" {
							b, _ := json.Marshal(collect)
							err := ioutil.WriteFile(cfg.Coverfile, b, 0600)
							if err != nil {
								fmt.Println(err)
							}
						} else {
							var buf bytes.Buffer
							err := cover.FormatLine(&buf, cfg.Covermode, collect)
							if err != nil {
								fmt.Println(err)
							} else {
								err = ioutil.WriteFile(cfg.Coverfile, buf.Bytes(), 0600)
								if err != nil {
									fmt.Println(err)
								}
							}
						}
					}
					printCoverage(collect)
				}
				browser.Stop()
				cancel()
				return nil
			}
		}
	}
}

func executeInTab(ctx context.Context, cfg *config.Config, idx string, c *cdp.Client, m *session.Manager, profiles chan []cover.Profile) error {
	target, err := c.Target.CreateTarget(ctx,
		target.NewCreateTargetArgs(idx),
	)
	if err != nil {
		return fmt.Errorf("trouble creating target %s:%v\n", idx, err)
	}
	pageConn, err := m.Dial(ctx, target.TargetID)
	if err != nil {
		return fmt.Errorf("trouble dialing session manager  %s:%v\n", idx, err)
	}
	defer pageConn.Close()
	pageClient := cdp.NewClient(pageConn)
	if cfg.Verbose || cfg.Cover {
		if err = pageClient.Console.Enable(ctx); err != nil {
			return fmt.Errorf("trouble enabling  console   %s:%v\n", idx, err)
		}
		csLog, err := pageClient.Console.MessageAdded(ctx)
		if err != nil {
			return fmt.Errorf("trouble receiving console client  %s:%v\n", idx, err)
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
	<-ctx.Done()
	return nil
}
