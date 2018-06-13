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
	t.Run("case insensitive", func(ts *testing.T) {
		v, err := QueryWith(customData, "Last 01 Version")
		if err != nil {
			ts.Fatal(err)
		}
		e := []string{"bb 8", "chrome 39", "edge 12", "ie 11"}
		if !reflect.DeepEqual(v, e) {
			ts.Errorf("expected %v got %v", e, fmt.Sprint(v))
		}
	})

}

func TestSince(t *testing.T) {
	ctx := map[string]data{
		"ie": data{
			name:     "ie",
			versions: []string{"1", "2", "3"},
			releaseDate: map[string]int64{
				"1": 0,          // Thu, 01 Jan 1970 00:00:00 +0000
				"2": 1483228800, // Sun, 01 Jan 2017 00:00:00 +0000
				"3": 1485907200, // Wed, 01 Feb 2017 00:00:00 +0000
			},
		},
	}
	t.Run("selects versions released on year boundaries", func(ts *testing.T) {
		e := []string{"ie 1", "ie 2", "ie 3"}
		g, err := QueryWith(ctx, "since 1970")
		if err != nil {
			ts.Fatal(err)
		}
		if !reflect.DeepEqual(e, g) {
			ts.Errorf("expected %v got %v", e, g)
		}
	})
	t.Run("is case insensitive", func(ts *testing.T) {
		e := []string{"ie 1", "ie 2", "ie 3"}
		g, err := QueryWith(ctx, "Since 1970")
		if err != nil {
			ts.Fatal(err)
		}
		if !reflect.DeepEqual(e, g) {
			ts.Errorf("expected %v got %v", e, g)
		}
	})
	t.Run("selects versions released on year and month boundaries", func(ts *testing.T) {
		e := []string{"ie 2", "ie 3"}
		g, err := QueryWith(ctx, "since 2017-01")
		if err != nil {
			ts.Fatal(err)
		}
		if !reflect.DeepEqual(e, g) {
			ts.Errorf("expected %v got %v", e, g)
		}
	})
	t.Run("selects versions released on date boundaries", func(ts *testing.T) {
		e := []string{"ie 3"}
		g, err := QueryWith(ctx, "since 2017-02-01")
		if err != nil {
			ts.Fatal(err)
		}
		if !reflect.DeepEqual(e, g) {
			ts.Errorf("expected %v got %v", e, g)
		}
	})
}

func TestGlobal(t *testing.T) {
	ctx := map[string]data{
		"ie": data{
			usage: map[string]float64{
				"8":  1,
				"9":  5,
				"10": 10.1,
				"11": 75,
			},
		},
	}
	t.Run("selects browsers by popularity", func(ts *testing.T) {
		e := []string{"ie 10", "ie 11"}
		g, err := QueryWith(ctx, "> 10%")
		if err != nil {
			ts.Fatal(err)
		}
		if !reflect.DeepEqual(e, g) {
			t.Errorf("expected %v got %v", e, g)
		}
	})
	t.Run("selects popularity by more or equal", func(ts *testing.T) {
		e := []string{"ie 10", "ie 11", "ie 9"}
		g, err := QueryWith(ctx, ">= 5%")
		if err != nil {
			ts.Fatal(err)
		}
		if !reflect.DeepEqual(e, g) {
			t.Errorf("expected %v got %v", e, g)
		}
	})
	t.Run("selects browsers by unpopularity", func(ts *testing.T) {
		e := []string{"ie 8"}
		g, err := QueryWith(ctx, "< 5%")
		if err != nil {
			ts.Fatal(err)
		}
		if !reflect.DeepEqual(e, g) {
			t.Errorf("expected %v got %v", e, g)
		}
	})
	t.Run("selects unpopularity by less or equal", func(ts *testing.T) {
		e := []string{"ie 8", "ie 9"}
		g, err := QueryWith(ctx, "<= 5%")
		if err != nil {
			ts.Fatal(err)
		}
		if !reflect.DeepEqual(e, g) {
			t.Errorf("expected %v got %v", e, g)
		}
	})
	t.Run("accepts non-space query", func(ts *testing.T) {
		e := []string{"ie 10", "ie 11"}
		g, err := QueryWith(ctx, ">10%")
		if err != nil {
			ts.Fatal(err)
		}
		if !reflect.DeepEqual(e, g) {
			t.Errorf("expected %v got %v", e, g)
		}
	})
	t.Run("works with float", func(ts *testing.T) {
		e := []string{"ie 11"}
		g, err := QueryWith(ctx, "> 10.2%")
		if err != nil {
			ts.Fatal(err)
		}
		if !reflect.DeepEqual(e, g) {
			t.Errorf("expected %v got %v", e, g)
		}
	})
	t.Run("works with float that has a leading dot", func(ts *testing.T) {
		e := []string{"ie 10", "ie 11", "ie 8", "ie 9"}
		g, err := QueryWith(ctx, "> .2%")
		if err != nil {
			ts.Fatal(err)
		}
		if !reflect.DeepEqual(e, g) {
			t.Errorf("expected %v got %v", e, g)
		}
	})
}
