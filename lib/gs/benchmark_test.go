package gs_test

import (
	"fmt"
	"testing"

	"github.com/gernest/vected/lib/gs"
)

func BenchmarkToString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		x := 100
		if i < 100 {
			x = i
		}
		css := gs.S("root", nest(x))
		b.StartTimer()
		gs.ToString(css)
	}
}

func nest(n int) gs.CSSRule {
	if n == 0 {
		return gs.S(fmt.Sprintf("& %d", n),
			gs.P("key", "value"),
		)
	}
	return gs.S(fmt.Sprintf("& %d", n), nest(n-1))
}
