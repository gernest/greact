package test

import "github.com/gernest/prom"

func TestBefore() prom.Test {
	return prom.Describe("Before",
		prom.It("be called before the testcase", func(t prom.T) {
			before := 500
			ts := prom.Describe("TestT_Before",
				prom.Before(func() {
					before = 200
				}),
				prom.It("must set before value", func(rs prom.T) {
				}),
			)
			prom.Exec(ts)
			if before != 200 {
				t.Errorf("expected %v got %v", 200, before)
			}
		}),
	)
}

func TestAfter() prom.Test {
	return prom.Describe("After",
		prom.It("should be called after the testcase", func(t prom.T) {
			after := 500
			ts := prom.Describe("TestAfter",
				prom.After(func() {
					after = 200
				}),
				prom.It("must set before value", func(rs prom.T) {
					after = 0
				}),
			)
			prom.Exec(ts)
			if after != 200 {
				t.Errorf("expected %v got %v", 200, after)
			}
		}),
	)
}
