package prom

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/gopherjs/gopherjs/js"
)

var (
	_ Test = (*Suite)(nil)
	_ Test = (*Expectation)(nil)
	_ Test = (List)(nil)
	_ Test = (*BeforeFuncs)(nil)
	_ Test = (*AfterFuncs)(nil)
	_ T    = (*baseT)(nil)
	_ T    = (*RSWithNode)(nil)
)

type Test interface {
	Exec()
	run()
}

// Describe describe what you want to test.
func Describe(desc string, tc ...Test) Test {
	t := &Suite{Desc: desc}
	for _, v := range tc {
		switch e := v.(type) {
		case *BeforeFuncs:
			if t.BeforeFuncs != nil {
				t.BeforeFuncs.Funcs =
					append(t.BeforeFuncs.Funcs, e.Funcs...)
			} else {
				t.BeforeFuncs = e
			}
		case *AfterFuncs:
			if t.AfterFuncs != nil {
				t.AfterFuncs.Funcs =
					append(t.AfterFuncs.Funcs, e.Funcs...)
			} else {
				t.AfterFuncs = e
			}
		case *Suite:
			e.Parent = t
			t.Children = append(t.Children, e)
		case *Expectation:
			e.Parent = t
			t.Expectations = append(t.Expectations, e)
		}
	}
	return t
}

type List []Test

func (ls List) run() {}
func (ls List) Exec() {
	for _, v := range ls {
		v.Exec()
	}
}

type SpecResult struct {
	Desc               string
	FullName           string
	FailedExpectations []*ExpectResult
	PassedExpectations []*ExpectResult
}

type ExpectResult struct {
	Desc     string
	Messages []string
}

type Suite struct {
	Parent             *Suite
	Desc               string
	BeforeFuncs        *BeforeFuncs
	AfterFuncs         *AfterFuncs
	MarkedSKip         bool
	MarkedSkipMessage  string
	Duration           time.Duration
	Expectations       []*Expectation
	FailedExpectations []*Expectation
	PassedExpectations []*Expectation
	Children           []*Suite
}

func (s *Suite) FullName() string {
	var names []string
	p := s
	for p != nil {
		names = append(names, p.Desc)
		p = p.Parent
	}
	size := len(names)
	rvs := make([]string, size)
	for i := 0; i < size; i++ {
		rvs[i] = names[size-i-1]
	}
	return strings.Join(rvs, " ")
}

func (s *Suite) Exec() {
	start := time.Now()
	if s.BeforeFuncs != nil {
		s.BeforeFuncs.Exec()
	}
	defer func() {
		s.Duration = time.Now().Sub(start)
	}()
	defer func() {
		if e := recover(); e != nil {
			if err, ok := e.(*Error); ok {
				if err.Pending {
					s.MarkedSKip = true
					s.MarkedSkipMessage = err.Message.Error()
				}
			}
		} else {
			if s.AfterFuncs != nil {
				s.AfterFuncs.Exec()
			}
		}

	}()
	if len(s.Expectations) > 0 {
		for _, v := range s.Expectations {
			v.Exec()
			if !v.Passed {
				s.FailedExpectations = append(s.FailedExpectations, v)
			} else {
				s.PassedExpectations = append(s.PassedExpectations, v)
			}
		}
	}
	if len(s.Children) > 0 {
		for _, v := range s.Children {
			v.Exec()
		}
	}
}

func Exec(ts ...Test) Test {
	ls := List(ts)
	ls.Exec()
	return ls
}

// Skip marks test suite as skipped
func Skip(message string) {
	panic(&Error{Pending: true, Message: errors.New(message)})
}

func (*Suite) run() {}

func (s *Suite) Result() *SpecResult {
	r := &SpecResult{
		Desc:     s.Desc,
		FullName: s.FullName(),
	}
	for _, v := range s.FailedExpectations {
		r.FailedExpectations = append(r.FailedExpectations, v.Result())
	}
	for _, v := range s.PassedExpectations {
		r.PassedExpectations = append(r.PassedExpectations, v.Result())
	}
	return r
}

func It(desc string, fn func(T)) Test {
	return &Expectation{Desc: desc, Func: fn}
}

type T interface {
	Error(...interface{})
	Errorf(string, ...interface{})
	Fatal(...interface{})
	FatalF(string, ...interface{})
	Errors() []error
}

type Expectation struct {
	Parent       *Suite
	Desc         string
	Func         func(T)
	Passed       bool
	FailMessages []string
}

func (*Expectation) run() {}
func (e *Expectation) Result() *ExpectResult {
	return &ExpectResult{
		Desc:     e.Desc,
		Messages: e.FailMessages,
	}
}

func (e *Expectation) Exec() {
	defer func() {
		if ev := recover(); ev != nil {
			if err, ok := ev.(*Error); ok {
				if !err.Pending {
					e.Passed = false
					e.FailMessages = append(e.FailMessages, err.Message.Error())
				}
			}
		}
	}()
	rs := &baseT{}
	if e.Func != nil {
		e.Func(rs)
	}
	errs := rs.Errors()
	if errs != nil {
		for _, v := range errs {
			e.FailMessages = append(e.FailMessages, v.Error())
		}
	} else {
		e.Passed = true
	}
}

type TInfo struct {
	Case         string   `json:"case"`
	Failed       bool     `json:"failed"`
	FailMessages []string `json:"fail_messages"`
}

type Error struct {
	Message error
	Pending bool
}

func (e *Error) Error() string {
	if e.Message != nil {
		return e.Error()
	}
	return ""
}

type baseT struct {
	err []error
}

func (b *baseT) Error(v ...interface{}) {
	b.err = append(b.err, errors.New(fmt.Sprint(v...)))
}
func (b *baseT) Fatal(v ...interface{}) {
	panic(&Error{Message: errors.New(fmt.Sprint(v...))})
}

func (b *baseT) Errorf(s string, v ...interface{}) {
	b.err = append(b.err, fmt.Errorf(s, v...))
}

func (b *baseT) FatalF(s string, v ...interface{}) {
	panic(&Error{Message: fmt.Errorf(s, v...)})
}

func (b *baseT) Errors() []error {
	return b.err
}

type Component struct {
	ID        string
	Component func() interface{}
	IsBody    bool
	Cases     List
}

func (c *Component) runIntegration() {}

type Integration interface {
	runIntegration()
}

// Node is an interface for retrieving a rendered Component node.
type Node interface {
	Node() *js.Object
}

type RSWithNode struct {
	baseT
	Object *js.Object
}

func (rs *RSWithNode) Node() *js.Object {
	return rs.Object
}

// Render returns an integration test for non body Components. Use this to test
// Components that renders spans,div etc.
//
// NOTE: func()interface{} was supposed to be func()vecty.ComponentOrHTML this
// is a workaround, because importing vecty in this package will make it
// impossible to run the commandline tools since vecty only works with the
// browser.
func Render(desc string, c func() interface{}, cases ...Test) Integration {
	return &Component{
		ID: desc, Component: c, Cases: cases,
	}
}

// RenderBody is like Render but the Component is expected to be elem.Body
func RenderBody(desc string, c func() interface{}, cases ...Test) Integration {
	return &Component{
		ID: desc, Component: c, Cases: cases, IsBody: true,
	}
}

type BeforeFuncs struct {
	Funcs []func()
}

func (*BeforeFuncs) run() {}
func (b *BeforeFuncs) Exec() {
	for _, v := range b.Funcs {
		v()
	}
}

// Before is a list of functions that will be executed before the actual test
// suite is run.
func Before(fn ...func()) Test {
	return &BeforeFuncs{Funcs: fn}
}

type AfterFuncs struct {
	Funcs []func()
}

func (*AfterFuncs) run() {}
func (b *AfterFuncs) Exec() {
	for _, v := range b.Funcs {
		v()
	}
}

// After is a list of functions that will be executed after the actual test
// suite is run.
// You can use this to release resources/cleanup after the tests are done.
func After(fn ...func()) Test {
	return &AfterFuncs{Funcs: fn}
}
