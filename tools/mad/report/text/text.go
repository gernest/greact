package text

import (
	"encoding/json"

	"github.com/gernest/vected/lib/mad"
)

func JSON(rs mad.Test, pkg, id string) {
	v, err := json.MarshalIndent(reportJSON(rs, pkg, id), "", "  ")
	if err != nil {
		panic(err)
	}
	println(string(v))
}

func reportJSON(rs mad.Test, pkg, id string) []*mad.SpecResult {
	var results []*mad.SpecResult
	switch e := rs.(type) {
	case *mad.Suite:
		e.ID = id
		e.Package = pkg
		results = append(results, e.Result())
	case mad.List:
		for _, v := range e {
			results = append(results, reportJSON(v, pkg, id)...)
		}
	}
	return results
}
