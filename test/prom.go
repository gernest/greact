package test

import "github.com/gernest/prom"

func TestT_Before() prom.Test {
	return prom.Describe("T.Before",
		prom.It("be called before the testcase", func(rs prom.Result) {
			before := 500
			ts := prom.Describe("TestT_Before",
				prom.Before(func() {
					before = 200
				}),
				prom.It("must set before value", func(rs prom.Result) {
				}),
			)
			prom.Exec(ts)
			if before != 200 {
				rs.Errorf("expected %v got %v", 200, before)
			}
		}),
	)
}

func TestT_After() prom.Test {
	return prom.Describe("T.After",
		prom.It("should be called after the testcase", func(rs prom.Result) {
			after := 500
			ts := prom.Describe("TestT_After",
				prom.After(func() {
					after = 200
				}),
				prom.It("must set before value", func(rs prom.Result) {
					after = 0
				}),
			)
			prom.Exec(ts)
			if after != 200 {
				rs.Errorf("expected %v got %v", 200, after)
			}
		}),
	)
}

func TestReverse() {

}
