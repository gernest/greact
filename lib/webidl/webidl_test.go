package webidl

import (
	"io/ioutil"
	"path/filepath"
	"testing"

	"github.com/kr/pretty"

	"github.com/gernest/vected/lib/webidl/token"
)

func TestParse(t *testing.T) {
	base := "fixture"
	dir, err := ioutil.ReadDir(base)
	if err != nil {
		t.Fatal(err)
	}
	for _, info := range dir {
		if info.IsDir() || filepath.Ext(info.Name()) == ".out" {
			continue
		}
		t.Run(info.Name(), func(ts *testing.T) {
			name := filepath.Join(base, info.Name())
			b, err := ioutil.ReadFile(name)
			if err != nil {
				ts.Fatal(err)
			}
			fs := token.NewFileSet()
			node, err := Parse(fs, name, b)
			if err != nil {
				ts.Fatal(err)
			}
			out := name + ".out"
			b, err = ioutil.ReadFile(out)
			if err != nil {
				ts.Fatal(err)
			}
			expect := string(b)
			got := Dump(node)
			if got != expect {
				t.Errorf("expected %s got %s", expect, got)
			}
		})
	}
}

func TestScan(t *testing.T) {
	b, err := ioutil.ReadFile("./fixture/extended_attr_args_list")
	if err != nil {
		t.Fatal(err)
	}
	tkn, err := collectTokens(token.NewFileSet(), "", b)
	if err != nil {
		t.Fatal(err)
	}
	t.Error(pretty.Sprint(tkn))
}
