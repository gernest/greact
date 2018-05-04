package prom

import (
	"errors"
	"fmt"
	"runtime"
)

type Test interface {
	run()
}

func Describe(desc string, ctx ...Test) Test {
	return &suite{desc: desc, cases: List(ctx)}
}

type List []Test

func (ls List) run() {}

type suite struct {
	desc  string
	cases List
}

func (*suite) run() {}

func It(desc string, fn func(Result)) Test {
	return &executioner{desc: desc, fn: fn}
}

type Result interface {
	Error(err interface{})
	Errorf(string, ...interface{})
}

type executioner struct {
	desc string
	fn   func(Result)
}

func (*executioner) run() {}

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

func Exec(ctx ...Test) (*ResultCtx, error) {
	rs := &ResultCtx{}
	for _, v := range ctx {
		s, ok := v.(*suite)
		if !ok {
			return nil, errors.New("only test suites are allowed on the top level")
		}
		ch := execSuite(s)
		ch.Parent = rs
		rs.Children = append(rs.Children, ch)
	}
	return rs, nil
}

func execSuite(s *suite) *ResultCtx {
	rs := &ResultCtx{
		Desc: s.desc,
	}
	for _, v := range s.cases {
		switch e := v.(type) {
		case *suite:
			ch := execSuite(e)
			ch.Parent = rs
			rs.Children = append(rs.Children, ch)
		case *executioner:
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

func execute(e *executioner) *ResultInfo {
	r := &baseResult{}
	if e.fn != nil {
		e.fn(r)
	}
	rs := &ResultInfo{Case: e.desc}
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
	base List
}

func (t *T) Before(fn ...func()) {

}

func (t *T) After(fn ...func()) {

}

func (t *T) Describe(desc string, cases ...Test) {
	t.base = append(t.base, Describe(desc, cases...))
}
