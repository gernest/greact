package vdom

import (
	"encoding/xml"
	"fmt"
	"io"
)

// Parse reads escaped html from src and returns a virtual tree structure
// representing it. It returns an error if there was a problem parsing the html.
func Parse(src []byte) (*Tree, error) {
	// Create a xml.Decoder to read from an IndexedByteReader
	r := NewIndexedByteReader(src)
	dec := xml.NewDecoder(r)
	dec.Entity = xml.HTMLEntity
	dec.Strict = false
	dec.AutoClose = xml.HTMLAutoClose

	// Iterate through each token and construct the tree
	tree := &Tree{src: src, reader: r}
	var currentParent *Element = nil
	for token, err := dec.Token(); ; token, err = dec.Token() {
		if err != nil {
			if err == io.EOF {
				// We reached the end of the document we were parsing
				break
			} else {
				// There was some unexpected error
				return nil, err
			}
		}
		if nextParent, err := parseToken(tree, token, currentParent); err != nil {
			return nil, err
		} else {
			currentParent = nextParent
		}
	}
	return tree, nil
}

// parseToken parses a single token and adds the appropriate node(s) to the tree. When calling
// parseToken iteratively, you should always capture the nextParent return and use it as the
// currentParent argument in the next iteration.
func parseToken(tree *Tree, token xml.Token, currentParent *Element) (nextParent *Element, err error) {
	var resultingNode Node
	switch token.(type) {
	case xml.StartElement:
		// Parse the name and attrs directly from the xml.StartElement
		startEl := token.(xml.StartElement)
		el := &Element{
			Name: parseName(startEl.Name),
			tree: tree,
		}
		for _, attr := range startEl.Attr {
			el.Attrs = append(el.Attrs, Attr{
				Name:  parseName(attr.Name),
				Value: attr.Value,
			})
		}
		if currentParent != nil {
			// Set the index based on how many children we've seen so far
			el.index = make([]int, len(currentParent.index)+1)
			copy(el.index, currentParent.index)
			el.index[len(currentParent.index)] = len(currentParent.children)
			// Set this element's parent
			el.parent = currentParent
			// Add this element to the currentParent's children
			currentParent.children = append(currentParent.children, el)
		} else {
			// There is no current parent, so set the index based on the
			// number of first-level child nodes we have seen so far for
			// this tree
			el.index = []int{len(tree.Children)}
		}
		// Set the srcStart to indicate where in tree.src the html for this element
		// starts. Since we don't know the exact length of the starting tag (might be extra whitespace
		// in between attributes), we can't just do arithmetic here. Instead, start from the current
		// offset and find the first preceding tag open (the '<' character)
		start, err := tree.reader.BackwardsSearch(0, tree.reader.Offset()-1, '<')
		if err != nil {
			return nil, err
		}
		el.srcStart = start
		// The innerHTML start is just the current offset
		el.srcInnerStart = tree.reader.Offset()
		// Set this element to the nextParent. The next node(s) we find
		// are children of this element until we reach xml.EndElement
		nextParent = el
		resultingNode = el
	case xml.EndElement:
		// Assuming the xml is well-formed, this marks the end of the current
		// parent
		endEl := token.(xml.EndElement)
		if currentParent == nil {
			// If we reach a closing tag without a corresponding start tag, the
			// xml is malformed
			return nil, fmt.Errorf("XML was malformed: Found closing tag %s before a corresponding opening tag.", parseName(endEl.Name))
		} else if currentParent.Name != parseName(endEl.Name) {
			// Make sure the name of the closing tag matches what we expect
			return nil, fmt.Errorf("XML was malformed: Found closing tag %s before the closing tag for %s", parseName(endEl.Name), currentParent.Name)
		}
		// The currentParent has been closed
		// Check whether it was autoclosed
		if wasAutoClosed(tree, currentParent.Name) {
			// There was not a corresponding closing tag, so the currentParent was
			// autoclosed and therefore can have no children. Don't worry about the
			// ending index, as our HTML method will do something different in this case.
			currentParent.autoClosed = true
		} else {
			// There was a corresponding closing tag, as indicated by the '/' symbol
			// This means we can use the underlying src buffer of the tree to get all
			// the bytes for the html of the currentParent and its children. The ending
			// index is the current offset.
			currentParent.srcEnd = tree.reader.Offset()
			// The innerHTML ends at the start of the closing tag
			// The closing tag has length of len(currentParent.Name) + 3
			// for the <, /, and > characters.
			closingTagLength := len(currentParent.Name) + 3
			currentParent.srcInnerEnd = tree.reader.Offset() - closingTagLength
		}
		// The currentParent has no more children.
		// The next node(s) we find must be children of currentParent.parent.
		if currentParent.parent != nil {
			nextParent = currentParent.parent
		} else {
			nextParent = nil
		}
	case xml.CharData:
		charData := token.(xml.CharData)
		// Parse the value from the xml.CharData
		text := &Text{
			Value: []byte(charData.Copy()),
		}
		if currentParent != nil {
			// Set the index based on how many children we've seen so far
			text.index = make([]int, len(currentParent.index)+1)
			copy(text.index, currentParent.index)
			text.index[len(currentParent.index)] = len(currentParent.children)
			// Set this text node's parent
			text.parent = currentParent
			// Add this text node to the currentParent's children
			currentParent.children = append(currentParent.children, text)
		} else {
			// There is no current parent, so set the index based on the
			// number of first-level child nodes we have seen so far for
			// this tree
			text.index = []int{len(tree.Children)}
		}
		resultingNode = text
		nextParent = currentParent
	case xml.Comment:
		xmlComment := token.(xml.Comment)
		// Parse the value from the xml.Comment
		comment := &Comment{
			Value: []byte(xmlComment.Copy()),
		}
		if currentParent != nil {
			// Set the index based on how many children we've seen so far
			comment.index = make([]int, len(currentParent.index)+1)
			copy(comment.index, currentParent.index)
			comment.index[len(currentParent.index)] = len(currentParent.children)
			// Set this comment node's parent
			comment.parent = currentParent
			// Add this comment node to the currentParent's children
			currentParent.children = append(currentParent.children, comment)
		} else {
			// There is no current parent, so set the index based on the
			// number of first-level child nodes we have seen so far for
			// this tree
			comment.index = []int{len(tree.Children)}
		}
		resultingNode = comment
		nextParent = currentParent
	case xml.ProcInst:
		return nil, fmt.Errorf("parse error: found token of type xml.ProcInst, which is not allowed in html")
	case xml.Directive:
		return nil, fmt.Errorf("parse error: found token of type xml.Directive, which is not allowed in html")
	}
	if resultingNode != nil && currentParent == nil {
		// If this node has no parents, it is one of the first-level children
		// of the tree.
		tree.Children = append(tree.Children, resultingNode)
	}
	return nextParent, nil
}

// parseName converts an xml.Name to a single string name. For our
// purposes we are not interested in the different namespaces, and
// just need to treat the name as a single string.
func parseName(name xml.Name) string {
	if name.Space != "" {
		return fmt.Sprintf("%s:%s", name.Space, name.Local)
	}
	return name.Local
}

// wasAutoClosed returns true if the tagName was autoclosed. It does
// this by reading the bytes backwards from the current offset of the
// tree's reader and comparing them to the expected closing tag.
func wasAutoClosed(tree *Tree, tagName string) bool {
	closingTag := fmt.Sprintf("</%s>", tagName)
	stop := tree.reader.Offset()
	start := stop - len(closingTag)
	if start < 0 {
		// The tag must have been autoclosed becuase there's
		// not enough space in the buffer before this point
		// to contain the entire closingTag.
		return true
	}
	// The tag was autoclosed iff the last bytes to be read
	// were not the closing tag.
	return string(tree.reader.buf[start:stop]) != closingTag
}
