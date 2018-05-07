package prom

import (
	"encoding/json"
	"errors"
	"fmt"
)

var (
	_ Test = (*T)(nil)
	_ Test = (*Suite)(nil)
	_ Test = (*ExecCommand)(nil)
	_ Test = (List)(nil)
)

type Test interface {
	run()
}

func Describe(desc string, ctx ...Test) Test {
	return &Suite{Desc: desc, Cases: List(ctx)}
}

type List []Test

func (ls List) run() {}

type Suite struct {
	Desc  string
	Cases List
}

func (*Suite) run() {}

func It(desc string, fn func(Result)) Test {
	return &ExecCommand{Desc: desc, Func: fn}
}

type Result interface {
	Error(...interface{})
	Errorf(string, ...interface{})
	Fatal(...interface{})
	FatalF(string, ...interface{})
}

type ExecCommand struct {
	Desc string
	Func func(Result)
}

func (*ExecCommand) run() {}

type ResultInfo struct {
	Case         string   `json:"case"`
	Failed       bool     `json:"failed"`
	FailMessages []string `json:"fail_messages"`
}

type Error struct {
	Message error
}

func (e *Error) Error() string {
	if e.Message != nil {
		return e.Error()
	}
	return ""
}

type ResultCtx struct {
	Parent   *ResultCtx    `json:"-"`
	Desc     string        `json:"description"`
	Children []*ResultCtx  `json:"children"`
	Results  []*ResultInfo `json:"results"`
}

func (r ResultCtx) ToJson() string {
	v, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(v)
}

func Exec(ctx ...*T) *ResultCtx {
	rs := &ResultCtx{}
	for _, v := range ctx {
		rs.Children = append(rs.Children, v.exec())
	}
	return rs
}

func execSuite(s *Suite) *ResultCtx {
	rs := &ResultCtx{
		Desc: s.Desc,
	}
	for _, v := range s.Cases {
		switch e := v.(type) {
		case *Suite:
			ch := execSuite(e)
			ch.Parent = rs
			rs.Children = append(rs.Children, ch)
		case *ExecCommand:
			rs.Results = append(rs.Results, execute(e))
			rv, ok := wrapPanic(e)
			rs.Results = append(rs.Results, rv)
			if ok {
				// call to Fatal or Fatalf halts the whole suite.
				return rs
			}
		}
	}
	return rs
}

func wrapPanic(e *ExecCommand) (rs *ResultInfo, panicked bool) {
	defer func() {
		if err := recover(); err != nil {
			rs.Failed = true
			rs.FailMessages = append(rs.FailMessages, fmt.Sprint(err))
			panicked = true
		}
	}()
	rs = execute(e)
	return
}

type baseResult struct {
	err []error
}

func (b *baseResult) Error(v ...interface{}) {
	b.err = append(b.err, errors.New(fmt.Sprint(v...)))
}
func (b *baseResult) Fatal(v ...interface{}) {
	panic(&Error{Message: errors.New(fmt.Sprint(v...))})
}

func (b *baseResult) Errorf(s string, v ...interface{}) {
	b.err = append(b.err, fmt.Errorf(s, v...))
}

func (b *baseResult) FatalF(s string, v ...interface{}) {
	panic(&Error{Message: fmt.Errorf(s, v...)})
}

// execute calls the function e.Func and register results.
func execute(e *ExecCommand) (rs *ResultInfo) {
	r := &baseResult{}
	if e.Func != nil {
		e.Func(r)
	}
	rs = &ResultInfo{Case: e.Desc}

	if r.err != nil {
		rs.Failed = true
		for _, v := range r.err {
			rs.FailMessages = append(rs.FailMessages, v.Error())
		}
	}
	return rs
}

type T struct {
	before func()
	after  func()
	suit   *Suite
	base   List
}

func NewTest(name string, hooks ...func()) *T {
	t := &T{suit: &Suite{Desc: name}}
	switch len(hooks) {
	case 1:
		t.before = hooks[0]
	case 2:
		t.before, t.after = hooks[0], hooks[1]
	}
	return t
}

func (t *T) Before(fn ...func()) {
	if len(fn) > 0 {
		t.before = fn[0]
	}
}

func (t *T) After(fn ...func()) {
	if len(fn) > 0 {
		t.after = fn[0]
	}
}

func (t *T) run() {}

func (t *T) Describe(desc string, cases ...Test) {
	t.suit.Cases = append(t.suit.Cases, Describe(desc, cases...))
}

func (t *T) exec() *ResultCtx {
	if t.before != nil {
		t.before()
	}
	rs := execSuite(t.suit)
	if t.after != nil {
		t.after()
	}
	return rs
}

func (t *T) Cases(tc ...Test) *T {
	t.suit.Cases = append(t.suit.Cases, tc...)
	return t
}
