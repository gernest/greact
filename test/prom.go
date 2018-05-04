package test

import "github.com/gernest/prom"

func TestT_Before(t *prom.T) {
	var before bool
	t.Before(func() {
		before = true
	})
	t.Describe("T.Before",
		prom.It("be called before the testcase", func(rs prom.Result) {
			if !before {
				rs.Error("expected before to be true")
			}
		}),
	)
}

func TestDescribe(t *prom.T) {
}
