package webidl

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"strings"
	"testing"

	"github.com/gernest/vected/lib/webidl/lexer"
	"golang.org/x/exp/ebnf"
)

func TestGrammar(t *testing.T) {
	f, err := ioutil.ReadFile("grammar.ebnf")
	if err != nil {
		t.Fatal(err)
	}
	g, err := ebnf.Parse("grammer", bytes.NewReader(f))
	if err != nil {
		t.Fatal(err)
	}
	err = ebnf.Verify(g, "definitions")
	if err != nil {
		t.Fatal(err)
	}
	def, err := lexer.EBNF(string(f))
	if err != nil {
		t.Fatal(err)
	}
	for k := range def.Symbols() {
		fmt.Println(k)
	}
	sample := `callback AbortCallback = void ();`
	_, err = lexer.ConsumeAll(def.Lex(strings.NewReader(sample)))
	if err != nil {
		t.Fatal(err)
	}
}
