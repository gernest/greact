package webidl

import (
	"os"
	"testing"

	"golang.org/x/exp/ebnf"
)

func TestGrammar(t *testing.T) {
	f, err := os.Open("grammar.ebnf")
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()
	g, err := ebnf.Parse("grammer", f)
	if err != nil {
		t.Fatal(err)
	}
	err = ebnf.Verify(g, "definitions")
	if err != nil {
		t.Fatal(err)
	}
}
