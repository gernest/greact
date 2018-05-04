package main

import "github.com/gernest/prom"

func main() {
	ctx, _ := prom.Exec(
		prom.Describe("Test",
			prom.It("Panics", func(rs prom.Result) {
				rs.Error("some fish")
			}),
			prom.It("Panics", func(yay prom.Result) {
				yay.Error("some yay")
			}),
		),
	)
	println(ctx)
}
