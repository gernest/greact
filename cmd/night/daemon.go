package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"path/filepath"
	"strings"
	"sync"
	"syscall"

	"github.com/gernest/alien"
	"github.com/gernest/prom/api"
	"github.com/gorilla/websocket"

	"github.com/takama/daemon"
	"github.com/urfave/cli"
)

const (
	serviceName  = "promnight"
	desc         = "Treat your vecty tests like your first date"
	port         = ":1955"
	testEndpoint = "test"
)

func startDaemon(ctx *cli.Context) error {
	srv, err := daemon.New(serviceName, desc)
	if err != nil {
		return err
	}
	msg, err := srv.Start()
	if err != nil {
		return fmt.Errorf("%s %v", msg, err)
	}
	fmt.Println(msg)
	return nil
}

func stopDaemon(ctx *cli.Context) error {
	srv, err := daemon.New(serviceName, desc)
	if err != nil {
		return err
	}
	msg, err := srv.Stop()
	if err != nil {
		return fmt.Errorf("%s %v", msg, err)
	}
	fmt.Println(msg)
	return nil
}

func installDaemon(ctx *cli.Context) error {
	srv, err := daemon.New(serviceName, desc)
	if err != nil {
		return err
	}
	msg, err := srv.Install()
	if err != nil {
		return fmt.Errorf("%s %v", msg, err)
	}
	fmt.Println(msg)
	return nil
}

func removeDaemon(ctx *cli.Context) error {
	srv, err := daemon.New(serviceName, desc)
	if err != nil {
		return err
	}
	msg, err := srv.Remove()
	if err != nil {
		return fmt.Errorf("%s %v", msg, err)
	}
	fmt.Println(msg)
	return nil
}

func statusDaemon(ctx *cli.Context) error {
	srv, err := daemon.New(serviceName, desc)
	if err != nil {
		return err
	}
	msg, err := srv.Status()
	if err != nil {
		return fmt.Errorf("%s %v", msg, err)
	}
	fmt.Println(msg)
	return nil
}

func daemonService(ctx *cli.Context) (err error) {
	host := ctx.String("host")
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, os.Kill, syscall.SIGTERM)
	rctx, cancel := context.WithCancel(context.Background())
	server := &http.Server{
		Addr:    port,
		Handler: apiServer(rctx, host),
	}
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		fmt.Printf("started test-server on %s\n", port)
		err = server.ListenAndServe()
		wg.Done()
	}()
	go func() {
		for {
			select {
			case <-rctx.Done():
				err = server.Shutdown(rctx)
			case <-interrupt:
				cancel()
			}
		}
	}()
	wg.Wait()
	return
}

var upgrade = websocket.Upgrader{}

// apiServer returns a *alien.Mux instance with endpoints registered for serving
// test suites.
func apiServer(ctx context.Context, host string) *alien.Mux {
	mux := alien.New()
	stats := &api.TestStats{}
	queue := make(chan *api.TestRequest, 50)
	cache := &sync.Map{}

	// we store channels that tracts results of a a particular test suite run.
	resultsCache := &sync.Map{}
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case req := <-queue:
				ts := &api.TestSuite{
					Package: req.Package,
					Path:    req.Path,
				}
				ts.Status = "queued"
				cache.Store(ts.Package, ts)
				stats.Queued = append(stats.Queued, ts)
			}
		}

	}()
	// we only display the server's test stats on GET /
	mux.Get("/", func(w http.ResponseWriter, r *http.Request) {
		b, _ := json.Marshal(stats)
		fmt.Fprint(w, string(b))
	})

	// Accept test-running request on POST /. On success we return a json object
	// with websocket url to connect for the test running events.
	mux.Post("/", func(w http.ResponseWriter, r *http.Request) {
		b, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		req := &api.TestRequest{}
		err = json.Unmarshal(b, req)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		res, err := homeResponse(host, req)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		queue <- req
		b, _ = json.MarshalIndent(res, "", "  ")
		fmt.Fprint(w, string(b))
	})
	mux.Get("/"+testEndpoint, func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query()
		pkg := query.Get("pkg")
		pkg, err := url.QueryUnescape(pkg)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		tsv, ok := cache.Load(pkg)
		if !ok {
			http.Error(w, "package not found", http.StatusNotFound)
			return
		}
		ts := tsv.(*api.TestSuite)
		ts.Status = "running"
		var rstChan chan *api.TestSuite
		if ch, ok := resultsCache.Load(pkg); ok {
			rstChan = ch.(chan *api.TestSuite)
		} else {
			rstChan = make(chan *api.TestSuite, 10)
			resultsCache.Store(pkg, rstChan)
		}
		rstChan <- ts
		conn, err := upgrade.Upgrade(w, r, nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		go func() {
			for {
				select {
				case <-ctx.Done():
					break
				default:
					if _, _, err := conn.NextReader(); err != nil {
						conn.Close()
						break
					}
				}
			}
		}()
		for {
			select {
			case v := <-rstChan:
				err := conn.WriteJSON(v)
				if err != nil {
					//log error
					fmt.Println(err)
					break
				}
			}
		}
	})
	mux.Get("/resource", func(w http.ResponseWriter, r *http.Request) {
		q := r.URL.Query()
		src := q.Get("src")
		if src == "" {
			http.Error(w, "file not found", http.StatusNotFound)
			return
		}
		src, err := url.QueryUnescape(src)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// we don't serve files which are outside the package test directory.
		inScope := false
		var pkg *api.TestSuite
		cache.Range(func(k, v interface{}) bool {
			key := k.(string)
			println(key)
			if strings.HasPrefix(src, key) {
				inScope = true
				pkg = v.(*api.TestSuite)
				return false
			}
			return true
		})

		if !inScope {
			http.Error(w, "files outside test scope are not allowed",
				http.StatusForbidden)
			return
		}
		rel, err := filepath.Rel(pkg.Package, src)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		path := filepath.Join(pkg.Path, rel)
		fmt.Println(path)
		http.ServeFile(w, r, path)
	})
	return mux
}

func inPkgScope(a, b string) bool {
	return strings.HasPrefix(a, b)
}

func homeResponse(base string, req *api.TestRequest) (*api.TestResponse, error) {
	if req.Path == "" {
		return nil, errors.New("Path can not be empty")
	}
	if !filepath.IsAbs(req.Path) {
		return nil, errors.New("Path must be absolute")
	}
	u, err := websocketURL(base, req.Package)
	if err != nil {
		return nil, err
	}
	return &api.TestResponse{WebsocketURL: u}, nil
}

func websocketURL(base string, pkg string) (string, error) {
	u, err := url.Parse(base)
	if err != nil {
		return "", err
	}
	query := make(url.Values)
	query.Set("pkg", pkg)
	u.Path = testEndpoint
	u.Scheme = "ws"
	u.RawQuery = query.Encode()
	return u.String(), nil
}
