package icon

import (
	"io/ioutil"
	"testing"

	"github.com/gernest/gs"
)

func TestStyle(t *testing.T) {
	b, _ := ioutil.ReadFile("icon.css")
	expect := string(b)
	s := gs.ToString(Style("instagram", `"\e70b"`, false))
	if s != expect {
		t.Errorf("expected %s got %s", expect, s)
	}

	b, _ = ioutil.ReadFile("icon_spin.css")
	expect = string(b)
	s = gs.ToString(Style("spin", `"\e70b"`, true))
	if s != expect {
		t.Errorf("expected %s got %s", expect, s)
	}
}
