package ws

import (
	"encoding/json"
	"net"
	"net/url"
	"path/filepath"

	"github.com/gernest/prom"
	"github.com/gopherjs/gopherjs/js"
	"github.com/gopherjs/websocket"
)

func New() (*WS, error) {
	h := js.Global.Get("location").Get("href").String()
	u, err := url.Parse(h)
	if err != nil {
		return nil, err
	}
	u.Path = "/test"
	u.Scheme = "ws"
	q := u.Query()
	src, err := url.QueryUnescape(q.Get("src"))
	if err != nil {
		return nil, err
	}
	src = filepath.Dir(src)
	src = filepath.Dir(src)
	q.Del("src")
	q.Set("pkg", src)
	u.RawQuery = q.Encode()
	println(u.String())
	conn, err := websocket.Dial(u.String())
	if err != nil {
		return nil, err
	}
	w := &WS{Conn: conn}
	w.enc = json.NewEncoder(w)
	return w, nil
}

type WS struct {
	net.Conn
	enc *json.Encoder
}

func (w *WS) Report(ts prom.Test) error {
	rs := toResult(ts)
	for _, v := range rs {
		if err := w.WriteResponse(v); err != nil {
			return err
		}
	}
	return nil
}

func (w *WS) WriteResponse(rs *prom.SpecResult) error {
	return w.enc.Encode(rs)
}

func toResult(rs prom.Test) []*prom.SpecResult {
	var results []*prom.SpecResult
	switch e := rs.(type) {
	case *prom.Suite:
		results = append(results, e.Result())
	case prom.List:
		for _, v := range e {
			results = append(results, toResult(v)...)
		}
	}
	return results
}
