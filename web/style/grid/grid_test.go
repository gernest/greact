package grid

import (
	"io/ioutil"
	"testing"

	"github.com/gernest/vected/lib/gs"
)

func TestColumn(t *testing.T) {
	opts := ColOptions{
		Span:   24,
		Pull:   24,
		Push:   24,
		Offset: 24,
		Order:  24,
	}
	s := Column(&opts,
		MediaOption{SM, &opts},
		MediaOption{LG, &opts},
	)
	v := gs.ToString(s)
	b, err := ioutil.ReadFile("grid.css")
	if err != nil {
		t.Fatal(err)
	}
	expect := string(b)
	if v != expect {
		t.Errorf("expected %s got %s", expect, v)
	}
	// ioutil.WriteFile("grid.css", []byte(v), 0600)
}

type mockSheetObject struct {
	rules    []string
	detached bool
}

func (m *mockSheetObject) InsertRule(rule string) {
	m.rules = append(m.rules, rule)
}

func (m *mockSheetObject) Detach() {
	m.detached = true
}
