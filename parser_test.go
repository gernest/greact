package vected

import (
	"fmt"
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
	_, err := ParseString(sample)
	if err != nil {
		t.Fatal(err)
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
