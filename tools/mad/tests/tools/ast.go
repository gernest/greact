package tools

import (
	"bytes"
	"go/parser"
	"go/printer"
	"go/token"
	"strings"

	"github.com/gernest/vected/lib/mad"

	"github.com/gernest/vected/tools/mad/tools"
)

const sample = `
package sample

import "github.com/gernest/vected/lib/mad"

func TestwithWrongNaming() mad.Test { return mad.List{} }
func TestWithWrongNaming() mad.Test { return mad.List{} }
func TestWithWrongReturn()          {}

func TestAddLine() mad.Test {
	return mad.It("adds line numbers", func(ts mad.T) {
		ts.Error("here ")
		ts.Errorf("here %d", 2)
		ts.Fatal("here 3")
		ts.Fatalf("here %s", 4)
	})
}
`

func TestAddLine() mad.Test {
	return mad.It("adds file and line number on error calls", func(t mad.T) {
		set := token.NewFileSet()
		name := "tests.go"
		f, err := parser.ParseFile(set, name, sample, 0)
		if err != nil {
			t.Fatal(err)
		}
		names := tools.AddFileNumber(set, f)
		if len(names.Unit) != 2 {
			t.Fatalf("expected 2 valid names got %d", len(names.Unit))
		}
		firstMatch := "TestWithWrongNaming"
		if names.Unit[0] != firstMatch {
			t.Errorf("expected %s got %s", firstMatch, names.Unit[0])
		}
		secondMatch := "TestAddLine"
		if names.Unit[1] != secondMatch {
			t.Errorf("expected %s got %s", secondMatch, names.Unit[1])
		}
		var buf bytes.Buffer
		printer.Fprint(&buf, set, f)

		expectedLines := []string{
			"tests.go:12", "tests.go:13", "tests.go:14", "tests.go:15",
		}
		txt := buf.String()
		for _, v := range expectedLines {
			if !strings.Contains(txt, v) {
				t.Errorf("expected %s to be added", v)
			}
		}
	})
}

const sample2 = `
package sample

import "github.com/gernest/vected/lib/mad"


func TestAddLine() mad.Test {
	return mad.List{
		mad.It("adds line numbers", func(ts mad.T) {
			ts.Error("here ")
			ts.Errorf("here %d", 2)
			wrap(ts)
		}),
		wrap2(),
	}
}

func wrap(t mad.T)  {
	t.Fatal("here 3")
	t.Fatalf("here %s", 4)
}

func wrap2() mad.Test {
	return mad.It("wrap2", func(t mad.T) {
		t.Fatal("here 5")
		t.Fatalf("here %s", 6)
	})
}

func wrap3() mad.Test {
	h := func(t mad.T) {
		t.Fatal("here 5")
		t.Fatalf("here %s", 6)
	}
	return mad.It("wrap2", h)
}


`

func TestWrap() mad.Test {
	return mad.It("wraps ", func(t mad.T) {
		set := token.NewFileSet()
		name := "tests.go"
		f, err := parser.ParseFile(set, name, sample2, 0)
		if err != nil {
			t.Fatal(err)
		}
		tools.AddFileNumber(set, f)
		var buf bytes.Buffer
		printer.Fprint(&buf, set, f)
		expectedLines := []string{
			"tests.go:10", "tests.go:11",
			"tests.go:19", "tests.go:20",
			"tests.go:25", "tests.go:26",
			"tests.go:32", "tests.go:33",
		}
		txt := buf.String()
		for _, v := range expectedLines {
			if !strings.Contains(txt, v) {
				t.Errorf("expected %s to be added", v)
			}
		}
	})
}
