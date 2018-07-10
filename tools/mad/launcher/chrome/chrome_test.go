package chrome

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"testing"
	"time"
)

func TestFindChrome(t *testing.T) {
	v, err := resolveChromePath()
	if err != nil {
		t.Fatal(err)
	}
	if len(v) == 0 {
		t.Fatal("expected absolute path to chrome")
	}
	for _, path := range v {
		_, err := os.Stat(path)
		if os.IsNotExist(err) {
			t.Errorf("expected file %s to exist", path)
		}
	}
}

func TestNew(t *testing.T) {
	p, err := randomPort()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(p)
	var buf bytes.Buffer
	l, err := New(Options{})
	if err != nil {
		t.Fatal(err)
	}
	l.Cmd.Stdout = &buf
	l.Cmd.Stderr = &buf
	fmt.Println(l.Cmd.Args)
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		err := l.Run()
		if err != nil {
			t.Error(err)
		}
		cancel()
	}()
	time.Sleep(5 * time.Second)
	l.Stop()
	<-ctx.Done()
	t.Error(buf.String())
}
