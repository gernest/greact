package prefix

import (
	"reflect"
	"sort"

	"github.com/gernest/gs/ciu/agents"
	"github.com/gernest/gs/data"
	"github.com/gernest/gs/prefix"
	"github.com/gernest/mad"
)

func TestPrefixes() mad.Test {
	tdata := struct {
		prefixes map[string]data.Data
	}{
		prefixes: map[string]data.Data{
			"a": data.Data{
				Browsers: []string{"firefox 21", "firefox 20 old", "chrome 30", "ie 6"},
			},
			"b": data.Data{
				Browsers: []string{"ie 7 new", "firefox 20"},
				Mistakes: []string{"-webkit-"},
				Props:    []string{"a", "*"},
			},
			"c": data.Data{
				Browsers: []string{"ie 7", "firefox 20"},
				Selector: true,
			},
		},
	}
	return mad.List{
		mad.Describe("Select",
			mad.It("selects necessary prefixes", func(t mad.T) {
				b, err := prefix.NewBrowser(agents.New(), "firefox 21", "ie 7")
				if err != nil {
					t.Fatal(err)
				}
				fill := &prefix.Prefixes{
					Browsers: b,
					Data:     tdata.prefixes,
				}
				sample := []struct {
					key   string
					add   bool
					value []string
				}{
					{key: "a", add: true, value: []string{"-moz-"}},
					{key: "b", add: true, value: []string{"-ms- new"}},
					{key: "c", add: true, value: []string{"-ms-"}},
					{key: "a", add: false, value: []string{"-webkit-", "-ms-", "-moz- old"}},
					{key: "b", add: false, value: []string{"-ms-", "-moz-", "-webkit-"}},
					{key: "c", add: false, value: []string{"-moz-"}},
				}

				sel := fill.Select(tdata.prefixes)
				for _, v := range sample {
					var g []string
					if v.add {
						g = sel.Add[v.key]
					} else {
						g = sel.Remove[v.key]
					}
					sort.Strings(g)
					sort.Strings(v.value)
					if !reflect.DeepEqual(g, v.value) {
						t.Errorf("expected %v got %v", v.value, g)
					}
				}
			}),
		),
	}
}
