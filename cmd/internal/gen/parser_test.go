package gen

import (
	"bytes"
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
		expect := `fmt.Sprint("hello")`
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
				expect: `fmt.Sprint("hello,", props.String("name"))`,
			},
			{
				src:    `{props.String("initialName")}/{s.State().String("name")}`,
				expect: `fmt.Sprint(props.String("initialName"), "/", s.State().String("name"))`,
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
}

func TestGenerate(t *testing.T) {
	geneateTest(t, "fixture/generate/basic.html")
	geneateTest(t, "fixture/generate/custom.html")
	geneateTest(t, "fixture/generate/custom_with_mapping.html", map[string]string{
		"custom": "Custom",
	})
}

func geneateTest(t *testing.T, file string, m ...map[string]string) {
	t.Run(file, func(t *testing.T) {
		b, err := ioutil.ReadFile(file)
		if err != nil {
			t.Fatal(err)
		}
		n, err := ParseString(string(b))
		if err != nil {
			t.Fatal(err)
		}
		ctx := GeneratorContext{
			StructName: "Hello",
			Recv:       "t",
			Node:       n,
		}

		var buf bytes.Buffer
		var ma map[string]string
		if len(m) > 0 {
			ma = m[0]
		}
		err = Generate(&buf, "generate", ma, ctx)
		if err != nil {
			t.Fatal(err)
		}
		output := file + ".go.out"
		ioutil.WriteFile(output, buf.Bytes(), 0600)
		o, err := ioutil.ReadFile(output)
		if err != nil {
			t.Fatal(err)
		}
		if !bytes.Equal(buf.Bytes(), o) {
			t.Errorf("expected:\n %s \n got:\n%s", string(o), buf.String())
		}
	})
}
