package test

import "github.com/gernest/prom"

func TestT_Before() prom.Test {
	return prom.Describe("T.Before",
		prom.It("be called before the testcase", func(rs prom.Result) {
			before := 500
			ts := prom.NewTest("TestT_Before")
			ts.Before(func() {
				before = 200
			})
			ts.Describe("Set before",
				prom.It("must set before value", func(rs prom.Result) {
					before += 100
				}),
			)
			_, err := prom.Exec(ts)
			if err != nil {
				rs.Errorf("expected no error got %v instead", err)
			} else {
				if before != 200 {
					rs.Errorf("expected %v got %v", 200, before)
				}
			}
		}),
	)
}

func TestT_After() prom.Test {
	return prom.Describe("T.After",
		prom.It("should be called after the testcase", func(rs prom.Result) {
			after := 500
			ts := prom.NewTest("TestT_Before")
			ts.Before(func() {
				after = after + 200
			})
			ts.Describe("Set before",
				prom.It("must set before value", func(rs prom.Result) {
					after = 0
				}),
			)
			_, err := prom.Exec(ts)
			if err != nil {
				rs.Errorf("expected no error got %v instead", err)
			} else {
				if after != 200 {
					rs.Errorf("expected %v got %v", 200, after)
				}
			}
		}),
	)
}
