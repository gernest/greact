package console

import (
	"fmt"

	"github.com/gernest/mad"
)

func Report(ts *mad.SpecResult) {
	printResult(ts, 0)
}

func printResult(ts *mad.SpecResult, level int) {
	fmt.Printf("%s%s --> \n", ident(level), ts.Desc)
	for _, v := range ts.FailedExpectations {
		fmt.Printf("%s✖ %s :\n", ident(level+1), v.Desc)
		for _, msg := range v.Messages {
			fmt.Printf("%s-- %s \n", ident(level+2), msg)
		}
	}
	for _, v := range ts.PassedExpectations {
		fmt.Printf("%s✔ %s \n", ident(level+1), v.Desc)
	}
	for _, v := range ts.Children {
		printResult(v, level+1)
	}
}

func ident(level int) string {
	s := ""
	for i := 0; i < level; i++ {
		s += "  "
	}
	return s
}
