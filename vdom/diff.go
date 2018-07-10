package vdom

func Diff(t, other *Tree) (PatchSet, error) {
	patches := []Patcher{}
	if err := recursiveDiff(&patches, t.Children, other.Children); err != nil {
		return nil, err
	}
	return patches, nil
}

func recursiveDiff(patches *[]Patcher, nodes, otherNodes []Node) error {
	numOtherNodes := len(otherNodes)
	numNodes := len(nodes)
	minNumNodes := numOtherNodes
	if numOtherNodes > numNodes {
		// There are more otherNodes than there are nodes.
		// We should append the additional nodes.
		for _, otherNode := range otherNodes[numNodes:] {
			*patches = append(*patches, &Append{
				Parent: otherNode.Parent(),
				Child:  otherNode,
			})
		}
		minNumNodes = numNodes
	} else if numNodes > numOtherNodes {
		// There are more nodes than there are otherNodes.
		// We should remove the additional children.
		for _, node := range nodes[numOtherNodes:] {
			*patches = append(*patches, &Remove{
				Node: node,
			})
		}
		minNumNodes = numOtherNodes
	}
	for i := 0; i < minNumNodes; i++ {
		otherNode := otherNodes[i]
		node := nodes[i]
		if match, _ := CompareNodes(node, otherNode, false); !match {
			// The nodes have different tag names or values. We should replace
			// node with otherNode
			*patches = append(*patches, &Replace{
				Old: node,
				New: otherNode,
			})
			continue
		}
		// NOTE: Since CompareNodes checks the type,
		// we can only reach here if the nodes are of
		// the same type.
		if otherEl, ok := otherNode.(*Element); ok {
			// Both nodes are elements. We need to treat them differently because
			// they have children and attributes.
			el := node.(*Element)
			// Add the patches needed to make the attributes match (if any)
			diffAttributes(patches, el, otherEl)
			// Recursively apply diff algorithm to each element's children
			recursiveDiff(patches, el.Children(), otherEl.Children())
		}
	}
	return nil
}

// diffAttributes compares the attributes in el to the attributes in otherEl
// and adds the necessary patches to make the attributes in el match those in
// otherEl
func diffAttributes(patches *[]Patcher, el, otherEl *Element) {
	otherAttrs := otherEl.AttrMap()
	attrs := el.AttrMap()
	for attrName := range attrs {
		// Remove any attributes in el that are not found in otherEl
		if _, found := otherAttrs[attrName]; !found {
			*patches = append(*patches, &RemoveAttr{
				Node:     el,
				AttrName: attrName,
			})
		}
	}
	// Now iterate through the attributes in otherEl
	for name, otherValue := range otherAttrs {
		value, found := attrs[name]
		if !found {
			// The attribute exists in otherEl but not in el,
			// we should add it.
			*patches = append(*patches, &SetAttr{
				Node: el,
				Attr: &Attr{
					Name:  name,
					Value: otherValue,
				},
			})
		} else if value != otherValue {
			// The attribute exists in el but has a different value
			// than it does in otherEl. We should set it to the value
			// in otherEl.
			*patches = append(*patches, &SetAttr{
				Node: el,
				Attr: &Attr{
					Name:  name,
					Value: otherValue,
				},
			})
		}
	}
}
