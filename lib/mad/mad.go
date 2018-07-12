// Package mad is a library for defining and executing tests. There are two
// types of tests supported by this library.
//
// Unit tests
//
// These are tests which are on a functional basis. They are not supposed to be
// interacting with the dom.
//
// Unit tests are defined in a composition way using the Describe and It functions.
//
// Integration tests
//
// These are tests which involves dom interaction. For now the vecty framework
// is the only supported framework for defining and executing these kind of
// tests.
//
// Please see the mad commandline tool, which is the application responsible for
// running the tests defined by this library.
package mad

import (
	"errors"
	"fmt"
	"runtime/debug"
	"strings"
	"time"
)

var (
	_ Test = (*Suite)(nil)
	_ Test = (*Expectation)(nil)
	_ Test = (List)(nil)
	_ Test = (*BeforeFuncs)(nil)
	_ Test = (*AfterFuncs)(nil)
	_ T    = (*baseT)(nil)
)

// Test is an interface for a testable object. Note that this is supposed to be
// used internally so user's can't implement this interface.
type Test interface {
	Exec()
	run()
}

// Describe describe what you want to test. The first argument desc, can be a
// simple string for identifier you want to test, that can be a function, as
// struct method or anything.
//
// You can use this to organise your test into a nested tree like structure.
func Describe(desc string, testCases ...Test) Test {
	t := &Suite{Desc: desc}
	for _, v := range flatCases(testCases...) {
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

func flatCases(cases ...Test) []Test {
	var o []Test
	for _, v := range cases {
		switch e := v.(type) {
		case List:
			o = append(o, flatCases(e...)...)
		default:
			o = append(o, e)
		}
	}
	return o
}

// List is a list of tests.
type List []Test

func (ls List) run() {}

// Exec implements Test interface. This will call Exec method of all the list
// members.
func (ls List) Exec() {
	for _, v := range ls {
		v.Exec()
	}
}

// SpecResult contains result information after executing a test suite.
type SpecResult struct {
	ID                 string          `json:"id"`
	Package            string          `json:"package"`
	Desc               string          `json:"description"`
	FullName           string          `json:"fullname"`
	Duration           time.Duration   `json:"duration"`
	FailedExpectations []*ExpectResult `json:"failed_expectations,omitempty"`
	PassedExpectations []*ExpectResult `json:"passed_expectations,omitempty"`
	Children           []*SpecResult   `json:"children,omitempty"`
}

// ExpectResult contains reults of executing expectation.
type ExpectResult struct {
	Desc       string        `json:"description"`
	Duration   time.Duration `json:"duration"`
	Messages   []string      `json:"error_messages"`
	StackTrace string        `json:"stack_trace"`
}

type Suite struct {
	Parent             *Suite
	Package            string
	ID                 string
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

// Fullname returns a string depicting full tree descriptions from the parent
// root Suite to the current one.
func (s *Suite) Fullname() string {
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

//Exec implements Test interface.
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

// Exec Executes one or more tests.
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

// Result returns results of executing the suite.
func (s *Suite) Result() *SpecResult {
	r := &SpecResult{
		ID:       s.ID,
		Package:  s.Package,
		Desc:     s.Desc,
		FullName: s.Fullname(),
		Duration: s.Duration,
	}
	for _, v := range s.FailedExpectations {
		r.FailedExpectations = append(r.FailedExpectations, v.Result())
	}
	for _, v := range s.PassedExpectations {
		r.PassedExpectations = append(r.PassedExpectations, v.Result())
	}
	for _, v := range s.Children {
		r.Children = append(r.Children, v.Result())
	}
	return r
}

// It defines expectations. The test logic happens in the function fn.
func It(desc string, fn func(T)) Test {
	return &Expectation{Desc: desc, Func: fn}
}

// T is an interface for failing expectations.
type T interface {
	Error(...interface{})
	Errorf(string, ...interface{})
	Fatal(...interface{})
	Fatalf(string, ...interface{})
	Errors() []string
}

// Expectation contains the main test function that checks expectations. If the
// main function after execution happens not to call any method of the passed T
// object then the test has passed.
type Expectation struct {
	Parent       *Suite
	Desc         string
	Func         func(T)
	Passed       bool
	FailMessages []string
	StackTrace   string
	Duration     time.Duration
}

func (*Expectation) run() {}

// Result returns *ExpectResult from executing the expectation.
func (e *Expectation) Result() *ExpectResult {
	return &ExpectResult{
		Desc:       e.Desc,
		Messages:   e.FailMessages,
		StackTrace: e.StackTrace,
		Duration:   e.Duration,
	}
}

// Exec runs the test function and records the result.
//
// TODO: add timeout.
func (e *Expectation) Exec() {
	start := time.Now()
	defer func() {
		if ev := recover(); ev != nil {
			e.FailMessages = append(e.FailMessages, fmt.Sprint(ev))
			e.StackTrace = string(debug.Stack())
		}
	}()
	defer func() {
		e.Duration = time.Now().Sub(start)
	}()
	rs := &baseT{}
	if e.Func != nil {
		e.Func(rs)
	}
	errs := rs.Errors()
	if errs != nil {
		for _, v := range errs {
			e.FailMessages = append(e.FailMessages, v)
		}
	} else {
		e.Passed = true
	}
}

// Error implements error interface. This is useful for taking care of interupts
// via panics.
type Error struct {
	Message error
	Pending bool
}

func (e *Error) Error() string {
	if e.Message != nil {
		return e.Message.Error()
	}
	return ""
}

type baseT struct {
	err []string
}

func (b *baseT) Error(v ...interface{}) {
	b.err = append(b.err, fmt.Sprint(v...))
}
func (b *baseT) Fatal(v ...interface{}) {
	panic(&Error{Message: errors.New(fmt.Sprint(v...))})
}

func (b *baseT) Errorf(s string, v ...interface{}) {
	b.err = append(b.err, fmt.Sprintf(s, v...))
}

func (b *baseT) Fatalf(s string, v ...interface{}) {
	panic(&Error{Message: fmt.Errorf(s, v...)})
}

func (b *baseT) Errors() []string {
	return b.err
}

type Component struct {
	ID        string
	Component func() interface{}
	IsBody    bool
	Cases     Test
}

func (c *Component) runIntegration() {}

// Integration is an interface for integration tests.
type Integration interface {
	runIntegration()
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
		ID: desc, Component: c, Cases: Describe(desc, cases...),
	}
}

// RenderBody is like Render but the Component is expected to be elem.Body
func RenderBody(desc string, c func() interface{}, cases ...Test) Integration {
	return &Component{
		ID: desc, Component: c, Cases: Describe(desc, cases...), IsBody: true,
	}
}

// SetupIntegration returns an Integration test ready for execution. name is the
// name of the test function.
func SetupIntegration(name string, i Integration) Integration {
	v := i.(*Component)
	v.ID = name
	v.Cases = Describe(name, v.Cases)
	return v
}

// BeforeFuncs contains functions that are supposed to be executed before a
// test.
type BeforeFuncs struct {
	Funcs []func()
}

func (*BeforeFuncs) run() {}

// Exec implements Test interface. When called this will iterate and call every
// function that is stored in the Funcs field. Iteration is done by the order in
// which the functions are added.
//
// TODO: Have timeout to allow handling of long running functions. One option is
// to pass context.Context as the first argument of the functions.
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

// AfterFuncs is like BeforeFuncs but executed after the test has run.
type AfterFuncs struct {
	Funcs []func()
}

func (*AfterFuncs) run() {}

// Exec like BeforeFuncs.Exec but executed after the test run.
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
