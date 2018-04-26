package gs

import (
	"testing"
)

func TestConditional(t *testing.T) {
	c := Cond("@media",
		S("0%",
			P("key1", "value1"),
		),
		S("10%",
			P("key2", "value2"),
		),
	)
	e := "@media {\n  0% {\n     key1 : value1;\n  }\n  10% {\n     key2 : value2;\n  }\n  \n\n}"
	v := toString(c)
	if v != e {
		t.Errorf("expected %s got %s", e, v)
	}
}
