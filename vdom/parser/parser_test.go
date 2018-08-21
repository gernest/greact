package parser

import (
	"io/ioutil"
	"testing"

	"github.com/gernest/vected/vdom"
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
		if n.Data != vdom.ContainerNode {
			t.Errorf("expected %s got %s", vdom.ContainerNode, n.Data)
		}
	})
}

func TestGenerate(t *testing.T) {
	sample := `<div className={props["classNames"]} key=value>
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
	n, err := ParseString(sample)
	if err != nil {
		t.Fatal(err)
	}
	v, err := GenerateRenderMethod("test", Context{
		Recv:       "t",
		StructName: "Hello",
		Node:       n,
	})
	if err != nil {
		t.Fatal(err)
	}
	// ioutil.WriteFile("test/test.gen.go", v, 0600)
	s := string(v)
	b, err := ioutil.ReadFile("test/test.gen.go")
	if err != nil {
		t.Fatal(err)
	}
	expected1 := string(b)
	if s != expected1 {
		t.Errorf("got wrong generated output")
	}
}

func TestInterpret(t *testing.T) {
	sample := []struct {
		expr, expect string
	}{
		{`{"hello"}`, `"hello"`},
		{"{props.class}", "props.class"},
	}
	for _, v := range sample {
		got := interpret(v.expr)
		if got != v.expect {
			t.Errorf("expected %s got %s", v.expect, got)
		}
	}
}
