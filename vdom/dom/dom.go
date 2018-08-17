package dom

import (
	"strings"

	"github.com/gernest/vected/vdom/value"
)

// Element is an alias for the dom node.
type Element value.Value

// HasProperty returns true if e has property.
func HasProperty(e Element, v string) bool {
	return e.Call("hasOwnProperty", v).Bool()
}

// CreateNode creates a dom element.
func CreateNode(doc value.Value, name string) Element {
	node := doc.Call("createElement", name)
	node.Set("normalizedNodeName", name)
	return node
}

const svg = "http://www.w3.org/2000/svg'"

// CreateSVGNode creates svg dom element.
func CreateSVGNode(doc value.Value, name string) Element {
	node := doc.Call("createElementNS", svg, name)
	node.Set("normalizedNodeName", name)
	return node
}

// returns true if value is not null or undefined.
func valid(v value.Value) bool {
	if v.Type() == value.TypeUndefined {
		return false
	}
	return v.Type() != value.TypeNull
}

// RemoveNode removes node from its parent if attached.
func RemoveNode(node value.Value) {
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
func SetAccessor(gen CallbackGenerator, node Element, name string, old, val interface{}, isSVG bool) {
	if name == "className" {
		name = "class"
	}
	switch name {
	case "class":
		v := val
		if v == nil {
			v = ""
		}
		node.Set("className", v)
	case "style":
		style := node.Get("style")
		switch e := val.(type) {
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
		node.Set("innerHTML", val)
	default:
		switch {
		case strings.HasPrefix(name, "on"):
			useCapture := name != strings.TrimSuffix(name, "Capture")
			name = eventName(name)
			if ev, ok := val.(func([]value.Value)); ok {
				cb := gen(ev)
				if old == nil {
					node.Call("addEventListener", name, cb, useCapture)
					// To release resources allocated for the callback we keep track of of all
					// callbacks added to this node.
					//
					// These can be later removed by calling the functions.
					var release value.Callback
					release = gen(func(args []value.Value) {
						node.Call("removeEventListener", name, cb, useCapture)
						cb.Release()
						release.Release()
					})
					releaseList := node.Get("_listeners")
					if releaseList.Type() == value.TypeUndefined {
						node.Set("_listeners", make(map[string]interface{}))
						releaseList = node.Get("_listeners")
					}
					releaseList.Set(name, release)
				}
			} else {
				// If we don't supply the event call back it is the same as saying remove
				// this event.
				//
				// We release the resources allocated for the event callback and free up the
				// event reference by setting its value to undefined.
				releaseList := node.Get("_listeners")
				if valid(releaseList) {
					releaseList.Call(name)
					releaseList.Set(name, "")
				}
			}
		case name != "list" && name != "type" && !isSVG && HasProperty(node, name):
			func() {
				defer recover()
				if val != nil {
					node.Set(name, val)
				} else {
					node.Set(name, "")
				}
			}()
			if (val == nil || !toBool(val)) && name != "spellcheck" {
				node.Call("removeAttribute", name)
			}
		default:
			//TODO handle namespace
		}
	}
}

// CallbackGenerator is a function that returns callbacks.
type CallbackGenerator func(fn func([]value.Value)) value.Callback

func toBool(v interface{}) bool {
	if v, ok := v.(bool); ok {
		return v
	}
	return false
}

// eventName takes a props event name and returns a string suitable for
// registering the event on the dom.
func eventName(name string) string {
	name = strings.ToLower(name)
	return name[2:]
}
