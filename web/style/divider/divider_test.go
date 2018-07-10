package divider

import (
	"io/ioutil"
	"testing"

	"github.com/gernest/vected/lib/gs"
)

func TestDividerStyle(t *testing.T) {
	s := gs.ToString(Style())
	// t.Error("yay")
	ioutil.WriteFile("divider.css", []byte(s), 0600)
}
