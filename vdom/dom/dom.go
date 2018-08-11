package dom

import (
	"strings"
	"syscall/js"
)

// Element is an alias for the dom node.
type Element = js.Value

var doc = js.Global().Get("document")

// CreateNode creates a dom element.
func CreateNode(name string) Element {
	node := doc.Call("createElement", name)
	node.Set("normalizedNodeName", name)
	return node
}

const svg = "http://www.w3.org/2000/svg'"

// CreateSVGNode creates svg dom element.
func CreateSVGNode(name string) Element {
	node := doc.Call("createElementNS", svg, name)
	node.Set("normalizedNodeName", name)
	return node
}

// returns true if value is not null or undefined.
func valid(v js.Value) bool {
	if v.InstanceOf(js.Undefined()) {
		return false
	}
	return !v.InstanceOf(js.Null())
}

// RemoveNode removes node from its parent if attached.
func RemoveNode(node js.Value) {
	parent := node.Get("parentNode")
	if valid(parent) {
		parent.Call("removeChild", node)
	}
}

// SetAccessor Set a named attribute on the given Node, with special behavior
// for some names and event handlers. If `value` is `null`, the
// attribute/handler will be removed.
// node An element to mutate
//
// name The name/key to set, such as an event or attribute name
// old The last value that was set for this name/node pair
// value An attribute value, such as a function to be used as an event handler
// isSVG Are we currently diffing inside an svg?
func SetAccessor(node Element, name string, old, value interface{}, isSVG bool) {
	if name == "className" {
		name = "class"
	}
	switch name {
	case "class":
		v := value
		if v == nil {
			v = ""
		}
		node.Set("className", v)
	case "style":
		style := node.Get("style")
		switch e := value.(type) {
		case string:
			style.Set("cssText", e)
		case map[string]string:
			if o, ok := old.(map[string]string); ok {
				for k := range o {
					if _, ok := e[k]; !ok {
						style.Set(k, "")
					}
				}
			}
			for k, v := range e {
				style.Set(k, v)
			}
		}
	case "dangerouslySetInnerHTML":
		node.Set("innerHTML", value)
	default:
		switch {
		case strings.HasPrefix(name, "on"):
			// TODO: register and handle event listeners
		}
	}
}
