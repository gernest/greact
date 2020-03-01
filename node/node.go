package node

import "github.com/gernest/greact/expr"

// A NodeType is the type of a Node.
type NodeType uint32

// common HTML node types
const (
	ErrorNode NodeType = iota
	TextNode
	DocumentNode
	ElementNode
	CommentNode
	DoctypeNode
)

func (n NodeType) String() string {
	switch n {
	case ErrorNode:
		return "ErrorNode"
	case TextNode:
		return "TextNode"
	case DocumentNode:
		return "DocumentNode"
	case ElementNode:
		return "ElementNode"
	case CommentNode:
		return "CommentNode"
	case DoctypeNode:
		return "DoctypeNode"
	default:
		return "ErrorNode"
	}
}

// Attribute represent html attribute. Val can be any valid go expression.
type Attribute struct {
	Namespace, Key string
	Val            interface{}
}

// Node represents a virtual dom node. This is a go object that represents a dom
// object.
type Node struct {
	Type      NodeType
	Key       string
	Data      string
	Namespace string
	Attr      []Attribute
	Children  []*Node
}

// New is a wrapper for creating new node. If children are provided adjacent
// text nodes will be merged to a single node.
func New(typ NodeType, ns, name string, attrs []Attribute, children ...*Node) *Node {
	var norm []Attribute
	var key string
	for _, v := range attrs {
		if v.Key == "key" {
			key = expr.Eval(v.Val)
		} else {
			norm = append(norm, v)
		}
	}
	if len(children) > 0 {
		norm = append(norm, Attribute{
			Key: "children",
			Val: children,
		})
	}
	n := &Node{
		Type:      typ,
		Namespace: ns,
		Key:       key,
		Data:      name,
		Attr:      norm,
	}
	var lastText *Node
	for _, v := range children {
		switch v.Type {
		case TextNode:
			if lastText != nil {
				lastText.Data += v.Data
			} else {
				lastText = v
				n.Children = append(n.Children, lastText)
			}
		default:
			lastText = nil
			n.Children = append(n.Children, v)
		}
	}
	return n
}

// Attr returns Attribute from the arguments. This doesn't do much appart from
// wrapping the arguments.
func Attr(ns, key string, val interface{}) Attribute {
	return Attribute{
		Namespace: ns,
		Key:       key,
		Val:       val,
	}
}

// Attrs is a wrapper/shortcut for optionally providing Attributes. Due tto the
// nature of composition for components, this saves space and improves
// readability.
func Attrs(attr ...Attribute) []Attribute {
	return attr
}
