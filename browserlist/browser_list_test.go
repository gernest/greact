package browserlist

import (
	"testing"

	"github.com/kr/pretty"
)

func TestRegexp(t *testing.T) {
	// d := getData()
	// pretty.Println(d)
	txt := "last 1 major version"
	v, err := Query(txt)
	if err != nil {
		t.Fatal(err)
	}
	pretty.Println(v)
	t.Error("yay")
}
