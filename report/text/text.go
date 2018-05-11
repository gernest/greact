package text

import (
	"github.com/gernest/prom"
)

func Report(rs prom.Test) {
	switch e := rs.(type) {
	case *prom.Suite:
		println(e.FullName())
	case prom.List:
		for _, v := range e {
			Report(v)
		}
	}
}
