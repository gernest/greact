package vdom

const (
	// ContainerNode this is tha name of a node which acts as a container to a
	// slice of other nodes.
	ContainerNode = "__internal_container___"
)

// A NodeType is the type of a Node.
type NodeType uint32

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

type Node struct {
	Type      NodeType
	Data      string
	Namespace string
	Attr      []Attribute
	Children  []*Node
}

// New is a wrapper for creating new node
func New(typ NodeType, ns, name string, attrs []Attribute, children ...*Node) *Node {
	return &Node{
		Type:      typ,
		Namespace: ns,
		Data:      name,
		Attr:      attrs,
		Children:  newChildren(children...),
	}
}

// newChildren processes n nodes.
//
// Adjacent text nodes are merged.
func newChildren(n ...*Node) []*Node {
	if len(n) > 0 {
		var o []*Node
		var lastText *Node
		for _, v := range n {
			switch v.Type {
			case TextNode:
				if lastText != nil {
					lastText.Data += v.Data
				} else {
					lastText = v
					o = append(o, lastText)
				}
			default:
				lastText = nil
				o = append(o, v)
			}
		}
		return o
	}
	return nil
}

func Attr(ns, key string, val interface{}) Attribute {
	return Attribute{
		Namespace: ns,
		Key:       key,
		Val:       val,
	}
}

func Attrs(attr ...Attribute) []Attribute {
	return attr
}

func (v *Node) Key() string {
	for _, v := range v.Attr {
		if v.Key == "key" {
			return v.Val.(string)
		}
	}
	return ""
}
