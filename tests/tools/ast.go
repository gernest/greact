package tools

import (
	"bytes"
	"go/parser"
	"go/printer"
	"go/token"
	"strings"

	"github.com/gernest/mad"

	"github.com/gernest/mad/tools"
)

const sample = `
package sample

import "github.com/gernest/mad"

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
