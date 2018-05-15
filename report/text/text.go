package text

import (
	"encoding/json"

	"github.com/gernest/prom"
)

func JSON(rs prom.Test, pkg, id string) {
	v, err := json.MarshalIndent(reportJSON(rs, pkg, id), "", "  ")
	if err != nil {
		panic(err)
	}
	println(string(v))
}

func reportJSON(rs prom.Test, pkg, id string) []*prom.SpecResult {
	var results []*prom.SpecResult
	switch e := rs.(type) {
	case *prom.Suite:
		e.ID = id
		e.Package = pkg
		results = append(results, e.Result())
	case prom.List:
		for _, v := range e {
			results = append(results, reportJSON(v, pkg, id)...)
		}
	}
	return results
}
