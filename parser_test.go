package vected

import (
	"io/ioutil"
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
	// ioutil.WriteFile("fixture/parser/gen0/test.gen.go", v, 0600)
	s := string(v)
	b, err := ioutil.ReadFile("fixture/parser/gen0/test.gen.go")
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

func TestInterpretText(t *testing.T) {
	t.Run("must do nothing if there is no templates", func(ts *testing.T) {
		n := "hello"
		v := interpretText(n)
		if v != n {
			ts.Errorf("expected %s got %s", n, v)
		}
	})
	t.Run("must transform templates", func(ts *testing.T) {
		sample := []struct {
			src, expect string
		}{
			{
				src:    `hello, {props.String("name")}`,
				expect: `fmt.Sprintf("%v%v","hello, ",props.String("name"))`,
			},
			{
				src:    `{props.String("initialName")}/{s.State().String("name)}`,
				expect: `fmt.Sprintf("%v%v%s%v","",props.String("initialName"),"/",s.State().String("name))`,
			},
		}
		for _, v := range sample {
			x := interpretText(v.src)
			if x != v.expect {
				t.Errorf("expected %s got %s", v.expect, x)
			}
		}
	})
}

// func TestFrags(t *testing.T) {
// 	sample := `<div className={props["classNames"]} key=value>
// 	<ul>
// 		<li>1</li>
// 		<li>2</li>
// 		<li>3</li>
// 		<li>4</li>
// 		<li>5</li>
// 	</ul>
// 	<ol>
// 		<li>1</li>
// 		<li>2</li>
// 		<li>3</li>
// 		<li>4</li>
// 		<li>5</li>
// 	</ol>
// </div>`
// 	n, err := ParseString(sample)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	t.Error(pretty.Sprint(n.Attr))
// }
