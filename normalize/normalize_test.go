package normalize

import (
	"io/ioutil"
	"testing"

	"github.com/gernest/gs"
)

func TestNew(t *testing.T) {
	c := New()
	s := gs.ToString(c)
	b, _ := ioutil.ReadFile("normalize.css")
	e := string(b)
	if s != e {
		t.Errorf("expected %s got %s", e, s)
	}
}
