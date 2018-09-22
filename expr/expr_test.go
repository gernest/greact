package expr

import (
	"reflect"
	"testing"
)

func TestExtractExpression(t *testing.T) {
	sample := []struct {
		src    string
		expect []Expression
	}{
		{"{a}{b}", []Expression{
			{Text: "a", Plain: false},
			{Text: "b", Plain: false},
		}},
		// Ignore white space
		{" {a} {b} ", []Expression{
			{Text: "a", Plain: false},
			{Text: "b", Plain: false},
		}},

		// strings
		{"hello", []Expression{
			{Text: "hello", Plain: true},
		}},

		// string with expr
		{"hello{a}", []Expression{
			{Text: "hello", Plain: true},
			{Text: "a", Plain: false},
		}},

		// string with expr and space
		{"hello, world {a}", []Expression{
			{Text: "hello, world", Plain: true},
			{Text: "a", Plain: false},
		}},
	}
	for _, v := range sample {
		e, err := ExtractExpressions(v.src, '{', '}')
		if err != nil {
			t.Error(err)
		}
		if !reflect.DeepEqual(e, v.expect) {
			t.Errorf("expected %v got %v", v.expect, e)
		}
	}
}
