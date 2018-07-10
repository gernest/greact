package vdom

import (
	"fmt"
	"github.com/gopherjs/gopherjs/js"
	"honnef.co/go/js/dom"
)

var document dom.Document

func init() {
	// We only want to initialize document if we are running in the browser.
	// We can detect this by checking if the document is defined.
	if js.Global != nil && js.Global.Get("document") != js.Undefined {
		document = dom.GetWindow().Document()
	}
}

// Patcher represents changes that can be made to the DOM.
type Patcher interface {
	// Patch applies the given patch to the DOM. The given root
	// is a relative starting point for the virtual tree in the
	// actual DOM.
	Patch(root dom.Element) error
}

// PatchSet is a set of zero or more Patchers
type PatchSet []Patcher

// Patch satisfies the Patcher interface and sequentially applies
// all the patches in the patch set.
func (ps PatchSet) Patch(root dom.Element) error {
	for _, patch := range ps {
		if err := patch.Patch(root); err != nil {
			return err
		}
	}
	return nil
}

// Append is a Patcher which will append a child Node to a parent Node.
type Append struct {
	Child  Node
	Parent *Element
}

// Patch satisfies the Patcher interface and applies the change to the
// actual DOM.
func (p *Append) Patch(root dom.Element) error {
	// fmt.Println("Got parent: ", p.Parent)
	// fmt.Println("Parent == nil: ", p.Parent == nil)
	// fmt.Println("Parent != nil: ", p.Parent != nil)
	var parent dom.Node
	if p.Parent != nil {
		// fmt.Println("Finding parent in DOM")
		parent = findInDOM(p.Parent, root)
	} else {
		// fmt.Println("Setting parent as root")
		parent = root
	}
	// fmt.Println("Computed parent: ", parent)
	child := createForDOM(p.Child)
	// fmt.Println("Created child: ", child)
	parent.AppendChild(child)
	// fmt.Println("Successfully appended")
	return nil
}

// Replace is a Patcher will will replace an old Node with a new Node.
type Replace struct {
	Old Node
	New Node
}

// Patch satisfies the Patcher interface and applies the change to the
// actual DOM.
func (p *Replace) Patch(root dom.Element) error {
	var parent dom.Node
	if p.Old.Parent() != nil {
		parent = findInDOM(p.Old.Parent(), root)
	} else {
		parent = root
	}
	oldChild := findInDOM(p.Old, root)
	newChild := createForDOM(p.New)
	parent.ReplaceChild(newChild, oldChild)
	return nil
}

// Remove is a Patcher which will remove the given Node.
type Remove struct {
	Node Node
}

// Patch satisfies the Patcher interface and applies the change to the
// actual DOM.
func (p *Remove) Patch(root dom.Element) error {
	var parent dom.Node
	if p.Node.Parent() != nil {
		parent = findInDOM(p.Node.Parent(), root)
	} else {
		parent = root
	}
	self := findInDOM(p.Node, root)
	parent.RemoveChild(self)

	// p.Node was removed, so subtract one from the final index for all
	// siblings that come after it.
	if p.Node.Parent() != nil {
		lastIndex := p.Node.Index()[len(p.Node.Index())-1]
		for _, sibling := range p.Node.Parent().Children()[lastIndex:] {
			switch t := sibling.(type) {
			case *Element:
				t.index[len(t.index)-1] = t.index[len(t.index)-1] - 1
			case *Text:
				t.index[len(t.index)-1] = t.index[len(t.index)-1] - 1
			case *Comment:
				t.index[len(t.index)-1] = t.index[len(t.index)-1] - 1
			default:
				panic("unreachable")
			}
		}
	}

	return nil
}

// SettAttr is a Patcher which will set the attribute of the given Node to
// the given Attr. It will overwrite any previous values for the given Attr.
type SetAttr struct {
	Node Node
	Attr *Attr
}

// Patch satisfies the Patcher interface and applies the change to the
// actual DOM.
func (p *SetAttr) Patch(root dom.Element) error {
	self := findInDOM(p.Node, root).(dom.Element)
	self.SetAttribute(p.Attr.Name, p.Attr.Value)
	return nil
}

// RemoveAttr is a Patcher which will remove the attribute with the given
// name from the given Node.
type RemoveAttr struct {
	Node     Node
	AttrName string
}

// Patch satisfies the Patcher interface and applies the change to the
// actual DOM.
func (p *RemoveAttr) Patch(root dom.Element) error {
	self := findInDOM(p.Node, root).(dom.Element)
	self.RemoveAttribute(p.AttrName)
	return nil
}

// findInDOM finds the node in the actual DOM corresponding
// to the given virtual node, using the given root as a relative
// starting point.
func findInDOM(node Node, root dom.Element) dom.Node {
	el := root.ChildNodes()[node.Index()[0]]
	for _, i := range node.Index()[1:] {
		el = el.ChildNodes()[i]
	}
	return el
}

// createForDOM creates a real node corresponding to the given
// virtual node. It does not insert it into the actual DOM.
func createForDOM(node Node) dom.Node {
	switch node.(type) {
	case *Element:
		vEl := node.(*Element)
		el := document.CreateElement(vEl.Name)
		for _, attr := range vEl.Attrs {
			el.SetAttribute(attr.Name, attr.Value)
		}
		el.SetInnerHTML(string(vEl.InnerHTML()))
		return el
	case *Text:
		vText := node.(*Text)
		textNode := document.CreateTextNode(string(vText.Value))
		return textNode
	case *Comment:
		vComment := node.(*Comment)
		commentNode := document.Underlying().Call("createComment", string(vComment.Value))
		return dom.WrapNode(commentNode)
	default:
		msg := fmt.Sprintf("Don't know how to create node for type %T", node)
		panic(msg)
	}
}
