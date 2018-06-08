package firefox

import (
	"context"
	"testing"
)

func TestGet(t *testing.T) {
	l, err := New(Options{
		Type:         Firefox,
		Verbose:      true,
		FirefoxFlags: []string{"-headless"}})
	if err != nil {
		t.Fatal(err)
	}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	go func() {
		err = l.Run(ctx)
		if err != nil {
			t.Fatal(err)
		}
	}()
	err = l.Ready()
	if err != nil {
		t.Fatal(err)
	}
	err = l.Stop()
	if err != nil {
		t.Fatal(err)
	}
}
