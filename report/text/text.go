package text

import (
	"encoding/json"

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

func JSON(rs prom.Test) {
	v, err := json.MarshalIndent(reportJSON(rs), " ", "  ")
	if err != nil {
		panic(err)
	}
	println(string(v))
}

func reportJSON(rs prom.Test) []*prom.SpecResult {
	var results []*prom.SpecResult
	switch e := rs.(type) {
	case *prom.Suite:
		results = append(results, e.Result())
	case prom.List:
		for _, v := range e {
			results = append(results, reportJSON(v)...)
		}
	}
	return results
}
