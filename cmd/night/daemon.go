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

	"github.com/dgraph-io/badger"
	"github.com/gernest/alien"
	"github.com/gernest/mad"
	"github.com/gernest/mad/api"
	"github.com/gorilla/websocket"
	"github.com/takama/daemon"

	"github.com/urfave/cli"
)

const (
	serviceName        = "madtitan"
	port               = ":1955"
	testEndpoint       = "test"
	testResultEndpoint = "/results"
	home               = "madtitan"
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

func deployDaemonService(ctx *cli.Context) (err error) {
	if err := prepareHomeDir(); err != nil {
		return err
	}
	db, err := openDatabase()
	if err != nil {
		return err
	}
	defer db.Close()
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, os.Kill, syscall.SIGTERM)
	rctx, cancel := context.WithCancel(context.Background())
	server := &http.Server{
		Addr:    port,
		Handler: apiServer(rctx, db, serverURL),
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
				return
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
func apiServer(ctx context.Context, db *badger.DB, host string) *alien.Mux {
	mux := alien.New()
	stats := &api.TestStats{}
	queue := make(chan *api.TestRequest, 50)
	results := make(chan *mad.SpecResult, 50)
	listeners := &sync.Map{}
	cache := &sync.Map{}
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case rst := <-results:
				err := saveCompletedTest(db, rst.ID, rst)
				if err != nil {

				}
				listeners.Range(func(_, v interface{}) bool {
					if fn, ok := v.(func(*mad.SpecResult)); ok {
						fn(rst)
					}
					return true
				})
			case req := <-queue:
				cache.Store(req.ID, req)
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
		id := query.Get("id")
		tsv, ok := cache.Load(id)
		if !ok {
			http.Error(w, "package not found", http.StatusNotFound)
			return
		}
		ts := tsv.(*api.TestRequest)
		rsChan := make(chan string, 5)
		listeners.Store(ts.ID, func(rs *mad.SpecResult) {
			b, _ := json.Marshal(rs)
			rsChan <- string(b)
		})
		conn, err := upgrade.Upgrade(w, r, nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer conn.Close()
		go func() {
			for {
				select {
				case <-ctx.Done():
					break
				default:
					_, _, err := conn.ReadMessage()
					if err != nil {
						conn.Close()
						fmt.Printf(" reading response %s\n", err)
						return
					}
				}
			}
		}()
		for {
			select {
			case v := <-rsChan:
				err := conn.WriteMessage(websocket.TextMessage, []byte(v))
				if err != nil {
					//log error
					fmt.Printf(" writing response %s\n", err)
					break
				}
			}
		}
	})
	mux.Get(testResultEndpoint, func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query()
		id := query.Get("id")
		_, ok := cache.Load(id)
		if !ok {
			http.Error(w, "package not found", http.StatusNotFound)
			return
		}
		conn, err := upgrade.Upgrade(w, r, nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer conn.Close()
		for {
			select {
			case <-ctx.Done():
				return
			default:
				rs := &mad.SpecResult{}
				if err := conn.ReadJSON(rs); err != nil {
					fmt.Printf(" reading response %s\n", err)
					return
				}
				results <- rs
			}
		}
	})
	mux.Get(resourcePath, func(w http.ResponseWriter, r *http.Request) {
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
		id := q.Get("id")
		tv, ok := cache.Load(id)
		if !ok {
			http.Error(w, "file not found", http.StatusNotFound)
			return
		}
		ts := tv.(*api.TestRequest)
		if !strings.HasPrefix(src, ts.Package) {
			http.Error(w, "files outside test scope are not allowed",
				http.StatusForbidden)
			return
		}
		rel, err := filepath.Rel(ts.Package, src)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		path := filepath.Join(ts.Path, rel)
		http.ServeFile(w, r, path)
	})
	return mux
}

func homeResponse(base string, req *api.TestRequest) (*api.TestResponse, error) {
	if req.Path == "" {
		return nil, errors.New("Path can not be empty")
	}
	if !filepath.IsAbs(req.Path) {
		return nil, errors.New("Path must be absolute")
	}
	u, err := websocketURL(base, req.ID)
	if err != nil {
		return nil, err
	}
	idx := indexHome(base, req)
	return &api.TestResponse{WebsocketURL: u, IndexURL: idx}, nil
}

func websocketURL(base string, id string) (string, error) {
	u, err := url.Parse(base)
	if err != nil {
		return "", err
	}
	query := make(url.Values)
	query.Set("id", id)
	u.Path = testEndpoint
	u.Scheme = "ws"
	u.RawQuery = query.Encode()
	return u.String(), nil
}

func indexHome(host string, req *api.TestRequest) string {
	query := make(url.Values)
	query.Set("src", filepath.Join(req.Package, testsOutDir, "index.html"))
	query.Set("id", req.ID)
	return fmt.Sprintf("%s%s?%s", host, resourcePath, query.Encode())
}

// Create the directory where the daemon will use to store data. In darwin this
// is in /usr/local/var/madnight.
func prepareHomeDir() error {
	p, err := homePath()
	if err != nil {
		return err
	}
	// this is the directory where we store data using the badger package. Fo clean
	// setup we can't just use the root homepath.
	h := filepath.Join(p, "data")
	_, err = os.Stat(h)
	if os.IsNotExist(err) {
		err = os.MkdirAll(h, 0755)
		if err != nil {
			return err
		}
	}
	return nil
}

func homePath() (string, error) {
	wd, err := os.Getwd()
	if err != nil {
		return "", err
	}
	h := filepath.Join(wd, home)
	return h, nil
}
