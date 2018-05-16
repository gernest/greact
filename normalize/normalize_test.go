package normalize

import (
	"io/ioutil"
	"testing"

	"github.com/gernest/gs"
)

func TestNew(t *testing.T) {
	c := New()
	s := gs.Process(c)
	v := s.String()
	b, _ := ioutil.ReadFile("normalize.css")
	e := string(b)
	if v != e {
		t.Errorf("expected %s got %s", e, v)
	}
	// ioutil.WriteFile("normalize.css", []byte(s.String()), 0600)
}
