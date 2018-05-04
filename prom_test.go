package prom

import (
	"testing"

	"github.com/kr/pretty"
)

func TestSuite(t *testing.T) {
	ctx := Describe("Prom",
		Describe("Describe",
			It("sets description title", func(rs Result) {
				desc := "Case1"
				n := Describe(desc)
				s, ok := n.(*suite)
				if !ok {
					rs.Error("expected return type to be suite")
				} else {
					if s.desc != desc {
						rs.Errorf("expected %v got %v", desc, s.desc)
					}
				}
			}),
			It("Fails", func(rs Result) {
				rs.Error("some fish")
			}),
		),
	)
	t.Error(pretty.Sprint(ctx))
	t.Error(pretty.Sprint(Exec(ctx)))
}

func TestLineNumer(t *testing.T) {
	s := []struct {
		src string
	}{
		{src: `(rs prom.Result)`},
		{src: `func(rs prom.Result)`},
		{src: `func(rs prom.Result){`},
		{src: `func(rs   prom.Result){`},
	}

	for _, v := range s {
		m := re.FindStringSubmatch(v.src)
		t.Error(pretty.Sprint(m))
	}
}

func TestLines(t *testing.T) {
	src := `prom.Describe("Test",
		prom.It("Panics", func(rs prom.Result) {
			rs.Error("some fish")
		}),
		prom.It("Panics", func(yay prom.Result) {
			yay.Error("some yay")
		}),
	),`

	LineNumber([]byte(src))
	t.Error("yay")
}
