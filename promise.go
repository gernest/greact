package prom

import (
	"fmt"
	"runtime"
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
	Error(err interface{})
	Errorf(string, ...interface{})
}

type ExecCommand struct {
	Desc string
	Func func(Result)
}

func (*ExecCommand) run() {}

type ResultInfo struct {
	Case         string
	Failed       bool
	FailMessages []Error
}

type Error struct {
	Message  string
	Location string
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
		}
	}
	return rs

}

type baseResult struct {
	err []string
}

func (b *baseResult) Error(v interface{}) {
	b.err = append(b.err, fmt.Sprint(v))
	fmt.Println(runtime.Caller(1))
}

func (b *baseResult) Errorf(s string, v ...interface{}) {
	b.err = append(b.err, fmt.Sprintf(s, v...))
}

func execute(e *ExecCommand) *ResultInfo {
	r := &baseResult{}
	if e.Func != nil {
		e.Func(r)
	}
	rs := &ResultInfo{Case: e.Desc}
	if r.err != nil {
		rs.Failed = true
		for _, v := range r.err {
			rs.FailMessages = append(rs.FailMessages, Error{
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
