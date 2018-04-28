package complete

import (
	"reflect"
	"testing"
)

func TestCompleter(t *testing.T) {
	c, err := New()
	if err != nil {
		t.Fatal(err)
	}
	g, err := c.FindHTML("a")
	if err != nil {
		t.Fatal(err)
	}
	e := []string{"abbr", "address", "area", "article", "aside", "audio"}
	if !reflect.DeepEqual(e, g) {
		t.Errorf("expected %v got %v", e, g)
	}
}
