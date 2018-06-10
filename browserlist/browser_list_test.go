package browserlist

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/kr/pretty"
)

func TestLast(t *testing.T) {
	customData := map[string]data{
		"ie": data{
			name:     "ie",
			released: []string{"9", "10", "11"},
			versions: []string{"9", "10", "11"},
		},
		"edge": data{
			name:     "edge",
			released: []string{"12"},
			versions: []string{"12", "13"},
		},
		"chrome": data{
			name:     "chrome",
			released: []string{"37", "38", "39"},
			versions: []string{"37", "38", "39", "40"},
		},
		"bb": data{
			name:     "bb",
			released: []string{"8"},
			versions: []string{"8"},
		},
		"firefox": data{},
	}
	t.Run("selects versions of each browser", func(ts *testing.T) {
		v, err := QueryWith(customData, "last 2 versions")
		if err != nil {
			ts.Fatal(err)
		}
		e := []string{"bb 8", "chrome 38", "chrome 39", "edge 12", "ie 10", "ie 11"}
		if !reflect.DeepEqual(v, e) {
			ts.Errorf("expected %v got %v", e, pretty.Sprint(v))
		}
	})
	t.Run("support pluralization", func(ts *testing.T) {
		v, err := QueryWith(customData, "last 1 version")
		if err != nil {
			ts.Fatal(err)
		}
		e := []string{"bb 8", "chrome 39", "edge 12", "ie 11"}
		if !reflect.DeepEqual(v, e) {
			ts.Errorf("expected %v got %v", e, fmt.Sprint(v))
		}
	})
}
