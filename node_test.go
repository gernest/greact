package vected

import "testing"

func TestVNode(t *testing.T) {
	h := NewNode
	t.Run("merges adjacent text child nodes", func(ts *testing.T) {
		x := h(ElementNode, "", "foo", nil,
			h(TextNode, "", "one", nil),
			h(TextNode, "", "two", nil),
			h(ElementNode, "", "bar", nil),
			h(TextNode, "", "three", nil),
		)
		if len(x.Children) != 3 {
			t.Errorf("expected 3 children got %d", len(x.Children))
		}
		txt := "onetwo"
		if x.Children[0].Type != TextNode {
			t.Error("expected text node")
		}
		if x.Children[0].Data != txt {
			t.Errorf("expected %s got %s", txt, x.Children[0].Data)
		}
	})
}
