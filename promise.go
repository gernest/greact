package prom

import (
	"errors"
	"fmt"
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
	Case         string
	Failed       bool
	FailMessages []*Error
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
	Parent   *ResultCtx
	Desc     string
	Children []*ResultCtx
	Results  []*ResultInfo
}

func Exec(ctx ...*T) (*ResultCtx, error) {
	rs := &ResultCtx{}
	for _, v := range ctx {
		rs.Children = append(rs.Children, v.exec())
	}
	return rs, nil
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
			if nv, ok := err.(*Error); ok {
				rs.Failed = true
				rs.FailMessages = append(rs.FailMessages, nv)
				panicked = true
			} else {
				panic(err)
			}
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
			rs.FailMessages = append(rs.FailMessages, &Error{
				Message: v,
			})
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

func NewTest(name string) *T {
	return &T{suit: &Suite{Desc: name}}
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
