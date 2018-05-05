package button

import (
	"io/ioutil"
	"testing"

	"github.com/gernest/gs"
)

func TestButton(t *testing.T) {
	s := gs.ToString(disable())
	ioutil.WriteFile("button.css", []byte(s), 0600)
}
