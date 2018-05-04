package test

import "github.com/gernest/prom"

func DescriptionTest(t *prom.T) {
	t.Before(func() {
		// pre hooks
	})
	t.After(func() {
		// cleanup
	})
	t.Describe("Describe",
		prom.It("Should Pass", func(rs prom.Result) {
		}),
		prom.It("Should Fail", func(rs prom.Result) {
			rs.Error("Fail")
		}),
	)
}
