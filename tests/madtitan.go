package tests

import (
	"github.com/gernest/mad"
	"github.com/gopherjs/gopherjs/js"
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
)

func TestBefore() mad.Test {
	return mad.Describe("Before",
		mad.It("should be called before the testcase", func(t mad.T) {
			before := 500
			ts := mad.Describe("TestTBefore",
				mad.Before(func() {
					before = 200
				}),
				mad.It("must set before value", func(rs mad.T) {
				}),
			)
			mad.Exec(ts)
			if before != 200 {
				t.Errorf("expected %v got %v", 200, before)
			}
		}),
	)
}

func TestAfter() mad.Test {
	return mad.Describe("After",
		mad.It("should be called after the testcase", func(t mad.T) {
			after := 500
			ts := mad.Describe("TestAfter",
				mad.After(func() {
					after = 200
				}),
				mad.It("must set before value", func(rs mad.T) {
					after = 0
				}),
			)
			mad.Exec(ts)
			if after != 200 {
				t.Errorf("expected %v got %v", 200, after)
			}
		}),
	)
}

func TestRenderBody() mad.Integration {
	txt := "hello,world"
	return mad.RenderBody("RenderBody",
		func() interface{} {
			return elem.Body(
				vecty.Text(txt),
			)
		},
		mad.It("must render the node in the html body", func(t mad.T) {
			defer func() {
				if err := recover(); err != nil {
					t.Error(err)
				}
			}()
			o := js.Global.Get("document").Get("body").Get("textContent").String()
			if o != txt {
				t.Errorf("expected %s got %s", txt, o)
			}
		}),
	)
}
