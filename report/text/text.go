package text

import (
	"bytes"
	"fmt"

	"github.com/gernest/prom"
)

func Report(rs *prom.ResultCtx) {
	var buf bytes.Buffer
	report(&buf, 0, rs)
	println(buf.String())
}

func report(o *bytes.Buffer, level int, rs *prom.ResultCtx) {
	if rs.Parent == nil {
		// we are at the root node
		for _, v := range rs.Children {
			report(o, level+1, v)
		}
	} else {
		fmt.Fprintf(o, "%s%s\n", ident(level), rs.Desc)
		if len(rs.Children) > 0 {
			for _, v := range rs.Children {
				report(o, level+2, v)
			}
		}
		fail := false
		for _, v := range rs.Results {
			if v.Failed {
				fail = true
				fmt.Fprintf(o, "%s FAIL %s\n", ident(level+1), v.Case)
				for _, v := range v.FailMessages {
					fmt.Fprintf(o, "%s --- %s\n", ident(level+1), v)
				}
			}
		}
		if fail {
			fmt.Fprintf(o, "%s FAILED \n", ident(level))
		} else {
			fmt.Fprintf(o, "%s PASSED \n", ident(level))
		}
	}
}

func ident(n int) string {
	v := ""
	if n == 0 {
		return v
	}
	for i := 0; i < n; i++ {
		v += " "
	}
	return v
}
