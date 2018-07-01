package vdom

import (
	"fmt"
	"html"
	"reflect"
)

// A Tree is a virtual, in-memory representation of a DOM tree
type Tree struct {
	// Children is the first-level child nodes for the tree
	Children []Node
	reader   *IndexedByteReader
	src      []byte
}

// HTML returns the html of this tree and recursively its children
// as a slice of bytes.
func (t *Tree) HTML() []byte {
	escaped := string(t.src)
	return []byte(html.UnescapeString(escaped))
}

// A Node is an element inside a tree.
type Node interface {
	// Parent returns the parent node or nil if there is none
	Parent() *Element
	// Children returns a slice of child nodes or nil if there
	// are none
	Children() []Node
	// HTML returns the unescaped html of this node and its
	// children as a slice of bytes.
	HTML() []byte
	// Index returns the child indexes starting at the root of the
	// virtual tree that can be used to get to this node. So if this
	// node is the second child of its parent, and its parent is the first
	// child of some root node, Index should return [0, 1]. This means we
	// can get to this node via root.ChildNodes()[0].ChildNodes()[1].
	Index() []int
}

// Attr is an html attribute
type Attr struct {
	Name  string
	Value string
}

// Element is an html element, e.g., <div></div>. Name does not include the
// <, >, or / symbols.
type Element struct {
	Name          string
	Attrs         []Attr
	parent        *Element
	children      []Node
	tree          *Tree
	srcStart      int
	srcEnd        int
	srcInnerStart int
	srcInnerEnd   int
	autoClosed    bool
	index         []int
}

func (e *Element) Parent() *Element {
	return e.parent
}

func (e *Element) Children() []Node {
	return e.children
}

func (e *Element) HTML() []byte {
	if e.autoClosed {
		// If the tag was autoclosed, it has no children. Just construct the html manually
		result := []byte(fmt.Sprintf("<%s", e.Name))
		for _, attr := range e.Attrs {
			result = append(result, []byte(fmt.Sprintf(` %s="%s"`, attr.Name, attr.Value))...)
		}
		result = append(result, '>')
		return result
	} else {
		escaped := string(e.tree.src[e.srcStart:e.srcEnd])
		return []byte(html.UnescapeString(escaped))
	}
}

// AttrMap returns this element's attributes as a map
// of attribute name to attribute value
func (e *Element) AttrMap() map[string]string {
	m := map[string]string{}
	for _, attr := range e.Attrs {
		m[attr.Name] = attr.Value
	}
	return m
}

// InnerHTML returns the unescaped html inside of e. So if e
// is <ul><li>one</li><li>two</li></ul>, it will return
// <li>one</li><li>two</li>. Since Element is the only type that
// can have children, this only makes sense for the Element type.
func (e *Element) InnerHTML() []byte {
	if e.autoClosed {
		// If the tag was autoclosed, it has no children, and therefore no inner html.
		return nil
	} else {
		escaped := string(e.tree.src[e.srcInnerStart:e.srcInnerEnd])
		return []byte(html.UnescapeString(escaped))
	}
}

// Selector returns a css selector which can be used to find
// the corresponding element in the actual DOM. The selector
// should be applied to the root of the tree, i.e. the starting
// point for the virtual tree in the actual DOM.
func (e *Element) Selector() string {
	// Simply use the index field to construct a selector with nth-child.
	selector := fmt.Sprintf("*:nth-child(%d)", e.index[0]+1)
	for _, i := range e.index[1:] {
		selector += fmt.Sprintf(" > *:nth-child(%d)", i+1)
	}
	return selector
}

func (e *Element) Index() []int {
	return e.index
}

// Compare non-recursively compares e to other. It does not check
// the child nodes since they can be a Node with any underlying type.
// If you want to compare the parent and children fields, use CompareNodes.
func (e *Element) Compare(other *Element, compareAttrs bool) (bool, string) {
	if e.Name != other.Name {
		return false, fmt.Sprintf("e.Name was %s but other.Name was %s", e.Name, other.Name)
	}
	if !compareAttrs {
		return true, ""
	}
	attrs := e.Attrs
	otherAttrs := other.Attrs
	if len(attrs) != len(otherAttrs) {
		return false, fmt.Sprintf("n has %d attrs but other has %d attrs.", len(attrs), len(otherAttrs))
	}
	for i, attr := range attrs {
		otherAttr := otherAttrs[i]
		if attr != otherAttr {
			return false, fmt.Sprintf("e.Attrs[%d] was %s but other.Attrs[%d] was %s", i, attr, i, otherAttr)
		}
	}
	return true, ""
}

// Text is a text node inside an xml/html document, i.e. anything
// not surrounded by tags.
type Text struct {
	Value  []byte
	parent *Element
	index  []int
}

func (t *Text) Parent() *Element {
	return t.parent
}

func (t *Text) Children() []Node {
	// A text node can't have any children
	return nil
}

func (t *Text) HTML() []byte {
	return t.Value
}

func (t *Text) Index() []int {
	return t.index
}

// Compare non-recursively compares t to other. It does not check
// the child nodes since they can be a Node with any underlying type.
// If you want to compare the parent and children fields, use CompareNodes.
func (t *Text) Compare(other *Text) (bool, string) {
	if string(t.Value) != string(other.Value) {
		return false, fmt.Sprintf("t.Value was %s but other.Value was %s", string(t.Value), string(other.Value))
	}
	return true, ""
}

// Comment is an xml/html comment of the form <!-- value -->.
// Value does not include the <!-- and --> markers.
type Comment struct {
	Value  []byte
	parent *Element
	index  []int
}

func (c *Comment) Parent() *Element {
	return c.parent
}

func (c *Comment) Children() []Node {
	// A commet node can't have any children
	return nil
}

func (c *Comment) HTML() []byte {
	// Re-add the open and close for the tag
	result := []byte("<!--")
	result = append(result, c.Value...)
	result = append(result, []byte("-->")...)
	return result
}

func (c *Comment) Index() []int {
	return c.index
}

// Compare non-recursively compares c to other. It does not check
// the child nodes since they can be a Node with any underlying type.
// If you want to compare the parent and children fields, use CompareNodes.
func (c *Comment) Compare(other *Comment) (bool, string) {
	if string(c.Value) != string(other.Value) {
		return false, fmt.Sprintf("c.Value was %s but other.Value was %s", string(c.Value), string(other.Value))
	}
	return true, ""
}

// Compare recursively compares t to other. It returns false and a detailed
// message if n does not equal other. Otherwise, it returns true and an empty
// string. NOTE: Comare never checks the parent properties of t's
// children. This is so you can construct a comparable tree inside a literal.
// (You can't set the parent field inside a literal).
func (t *Tree) Compare(other *Tree, compareAttrs bool) (bool, string) {
	if len(t.Children) != len(other.Children) {
		return false, fmt.Sprintf("t had %d first-level children but other had %d", len(t.Children), len(other.Children))
	}
	for i, root := range t.Children {
		otherRoot := other.Children[i]
		if match, msg := CompareNodesRecursive(root, otherRoot, compareAttrs); !match {
			return false, msg
		}
	}
	return true, ""
}

// CompareNodes non-recursively compares n to other. It returns false and
// a detailed message if n does not equal other. Otherwise, it returns true and
// an empty string. NOTE: CompareNodes never checks the parent properties of n
// or n's children. This is so you can construct a comparable tree inside a literal.
// (You can't set the parent field inside a literal).
func CompareNodes(n Node, other Node, compareAttrs bool) (bool, string) {
	if reflect.TypeOf(n) != reflect.TypeOf(other) {
		return false, fmt.Sprintf("n has underlying type %T but the other node has underlying type %T", n, other)
	}
	switch n.(type) {
	case *Element:
		el := n.(*Element)
		otherEl := other.(*Element)
		if match, msg := el.Compare(otherEl, compareAttrs); !match {
			return false, msg
		}
	case *Text:
		text := n.(*Text)
		otherText := other.(*Text)
		if match, msg := text.Compare(otherText); !match {
			return false, msg
		}
	case *Comment:
		comment := n.(*Comment)
		otherComment := other.(*Comment)
		if match, msg := comment.Compare(otherComment); !match {
			return false, msg
		}
	default:
		return false, fmt.Sprintf("Don't know how to compare n of underlying type %T", n)
	}
	return true, ""
}

// CompareNodesRecursive recursively compares n to other. It returns false and
// a detailed message if n does not equal other. Otherwise, it returns true and
// an empty string. NOTE: CompareNodesRecursive never checks the parent properties
// of n or n's children. This is so you can construct a comparable tree inside a
// literal. (You can't set the parent field inside a literal).
func CompareNodesRecursive(n Node, other Node, compareAttrs bool) (bool, string) {
	if match, msg := CompareNodes(n, other, compareAttrs); !match {
		return false, msg
	}
	children := n.Children()
	otherChildren := other.Children()
	if len(children) != len(otherChildren) {
		return false, fmt.Sprintf("n has %d children but other has %d children.", len(children), len(otherChildren))
	}
	for i, child := range children {
		otherChild := otherChildren[i]
		if match, msg := CompareNodesRecursive(child, otherChild, compareAttrs); !match {
			return false, msg
		}
	}
	return true, ""
}
