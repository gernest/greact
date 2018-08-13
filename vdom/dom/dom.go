package dom

import (
	"fmt"
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
	if v.Type() == js.TypeUndefined {
		return false
	}
	return v.Type() != js.TypeNull
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
			useCapture := name != strings.TrimSuffix(name, "Capture")
			name = eventName(name)
			fmt.Println(name)
			if ev, ok := value.(Event); ok {
				cb := js.NewEventCallback(ev.Flags(), ev.Call)
				if old == nil {
					var release js.Callback
					release = js.NewCallback(func(args []js.Value) {
						node.Call("removeEventListener", name, cb, useCapture)
						cb.Release()
						release.Release()
					})
					node.Call("addEventListener", name, cb, useCapture)
					// To release resources allocated for the callback we keep track of of all
					// callbacks added to this node.
					//
					// These can be later removed by calling the functions.
					releaseList := node.Get("_listeners")
					if releaseList.Type() == js.TypeUndefined {
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
					releaseList.Set(name, js.Undefined())
				}
			}
		}
	}
}

type Event interface {
	Flags() js.EventCallbackFlag
	Call(js.Value)
}
type eventWrapper struct {
	flags js.EventCallbackFlag
	event func(js.Value)
}

func (e eventWrapper) Flags() js.EventCallbackFlag {
	return e.flags
}
func (e eventWrapper) Call(v js.Value) {
	e.event(v)
}

// NewEvent returns new Event call back.
func NewEvent(flags js.EventCallbackFlag, fn func(event js.Value)) Event {
	return eventWrapper{flags: flags, event: fn}
}

// eventName takes a props event name and returns a string suitable for
// registering the event on the dom.
func eventName(name string) string {
	name = strings.ToLower(name)
	return name[2:]
}
