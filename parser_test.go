package vected

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
)

func TestClear(t *testing.T) {
	t.Run("should return  element", func(ts *testing.T) {
		e := `<div></div>`
		n, err := ParseString(e)
		if err != nil {
			ts.Fatal(err)
		}
		if n.Data != "div" {
			t.Errorf("expected div got %s", n.Data)
		}
	})
	t.Run("should return  container element", func(ts *testing.T) {
		e := `
		<div>
		</div>
		<div>
		</div>
		`
		n, err := ParseString(e)
		if err != nil {
			ts.Fatal(err)
		}
		if n.Data != "div" {
			t.Errorf("expected div got %s", n.Data)
		}
		if len(n.Children) != 2 {
			t.Errorf("expected 2 children got %d", len(n.Children))
		}
	})
}

func TestInterpret(t *testing.T) {
	sample := []struct {
		expr, expect string
	}{
		{`{"hello"}`, `"hello"`},
		{"{props.class}", "props.class"},
	}
	for _, v := range sample {
		got, err := interpret(v.expr)
		if err != nil {
			t.Fatal(err)
		}
		if got != v.expect {
			t.Errorf("expected %s got %s", v.expect, got)
		}
	}
}

func TestInterpretText(t *testing.T) {
	t.Run("must do nothing if there is no templates", func(ts *testing.T) {
		n := "hello"
		v, err := interpretText(n)
		if err != nil {
			ts.Fatal(err)
		}
		expect := `fmt.Println("hello")`
		if v != expect {
			ts.Errorf("expected %s got %s", expect, v)
		}
	})
	t.Run("must transform templates", func(ts *testing.T) {
		sample := []struct {
			src, expect string
		}{
			{
				src:    `hello, {props.String("name")}`,
				expect: `fmt.Println("hello,", props.String("name"))`,
			},
			{
				src:    `{props.String("initialName")}/{s.State().String("name")}`,
				expect: `fmt.Println(props.String("initialName"), "/", s.State().String("name"))`,
			},
		}
		for _, v := range sample {
			x, err := interpretText(v.src)
			if err != nil {
				ts.Fatal(err)
			}
			if x != v.expect {
				t.Errorf("expected %s got %s", v.expect, x)
			}
		}
	})
	fmt.Println(1, 2, 3)
	t.Error("yay")
}

func TestSomeShit(t *testing.T) {

	genSample := `<div className={props["classNames"]} key=value>
	<ul>
		<li>1</li>
		<li>2</li>
		<li>3</li>
		<li>4</li>
		<li>5</li>
	</ul>
	<ol>
		<li>1</li>
		<li>2</li>
		<li>3</li>
		<li>4</li>
		<li>5</li>
	</ol>
</div>`
	expect := `package hello

import "context"
import "fmt"
import "github.com/gernest/vected"

var H = vected.NewNode
var HA = vected.Attr
var HAT = vected.Attrs

func (t *Hello) Render(ctx context.Context, props vected.Props, state vected.State) *vected.Node {
	return H(3, "", "div", HAT(HA("", "classname", props["classNames"]), HA("", "key", "value")), H(3, "", "ul", nil, H(3, "", "li", nil, H(1, "", "1", nil)), H(3, "", "li", nil, H(1, "", "2", nil)), H(3, "", "li", nil, H(1, "", "3", nil)), H(3, "", "li", nil, H(1, "", "4", nil)), H(3, "", "li", nil, H(1, "", "5", nil))), H(3, "", "ol", nil, H(3, "", "li", nil, H(1, "", "1", nil)), H(3, "", "li", nil, H(1, "", "2", nil)), H(3, "", "li", nil, H(1, "", "3", nil)), H(3, "", "li", nil, H(1, "", "4", nil)), H(3, "", "li", nil, H(1, "", "5", nil))))
}
`
	n, err := ParseString(genSample)
	if err != nil {
		t.Fatal(err)
	}

	ctx := GeneratorContext{
		StructType: "Hello",
		Receiver:   "t",
		Node:       n,
	}

	var out bytes.Buffer
	err = Generate(&out, "hello", ctx)
	if err != nil {
		t.Fatal(err)
	}
	got := out.String()
	expect = strings.TrimSpace(expect)
	got = strings.TrimSpace(got)
	if got != expect {
		t.Error("got wrong output")
	}
	// ioutil.WriteFile("./tmp/hello/hello_render_gen.go", out.Bytes(), 0600)
	// printer.Fprint(os.Stdout, token.NewFileSet(), file)
}
