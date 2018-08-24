// Package vected is a component based frontend framework for golang. This
// framework delivers high performance and sleek ui's, that works both on the
// serverside and the frontend.
//
// Go templates are used as the main templating system. The framework is
// inspired by react, especially preact which I used to learn more about how
// react works.
//
// Also, this borrows from vue js. The templates are just normal go templates so
// no need to learn a different syntax.
//
// The user intrface styles is ant design
// see https://github.com/ant-design/ant-design to learn more about ant design.
package vected

import (
	"container/list"
	"context"
	"strings"
	"sync"

	"github.com/gernest/vected/elements"
	"github.com/gernest/vected/vdom/value"

	"github.com/gernest/vected/prop"
	"github.com/gernest/vected/state"
	"github.com/gernest/vected/vdom"
	"github.com/gernest/vected/vdom/dom"
)

// RenderMode is a flag determining how a component is rendered.
type RenderMode uint

//supported render mode
const (
	No RenderMode = iota
	Force
	Sync
	Async
)

// AttrKeys is a key used to store node's attributes/props
const AttrKey = "__vected_attr__"

var queue = NeqQueueRenderer()
var mounts = list.New()
var isSVGMode bool
var hydrating bool
var diffLevel int

// Component is an interface which defines a unit of user interface.
type Component interface {
	Render(context.Context, prop.Props, state.State) *vdom.Node
	core() *Core
}

type Templater interface {
	Template() string
}

type Constructor interface {
	New() Component
}

// Core is th base struct that every struct that wants to implement Component
// interface must embed.
//
// This is used to make Props available to the component.
type Core struct {
	props           prop.Props
	state           state.State
	prevProps       prop.Props
	prevState       state.State
	disable         bool
	renderCallbacks []func()
	context         context.Context
	prevContext     context.Context
	component       Component
	parentComponent Component
	base            dom.Element
	nextBase        dom.Element
	dirty           bool
	key             prop.NullString

	// This is a callback used to receive instance of Component or the Dom element.
	// after they have been mounted.
	ref func(interface{})

	// priority this is a number indicating how important this component is in the
	// re rendering queue. The higher the number the more urgent re renders.
	priority int
}

func (c *Core) core() *Core { return c }

// SetState updates component state and schedule re rendering.
func (c *Core) SetState(newState state.State, callback ...func()) {
	prev := c.prevState
	c.prevState = newState
	c.state = state.Merge(prev, newState)
	if len(callback) > 0 {
		c.renderCallbacks = append(c.renderCallbacks, callback...)
	}
	//TODO enqueue this for re rendering.
}

// InitState is an interface for exposing initial state.
// Component should implement this interface if they want to set initial state
// when the component is first created before being rendered.
type InitState interface {
	InitState() state.State
}

// InitProps is an interface for exposing default props. This will be merged
// with other props before being sent to render.
type InitProps interface {
	InitProps() prop.Props
}

// WillMount is an interface defining a callback which is invoked before the
// component is mounted on the dom.
type WillMount interface {
	ComponentWillMount()
}

// DidMount is an interface defining a callback that is invoked after the
// component has been mounted to the dom.
type DidMount interface {
	ComponentDidMount()
}

// WillUnmount is an interface defining a callback that is invoked prior to
// removal of the rendered component from the dom.
type WillUnmount interface {
	ComponentWillUnmount()
}

// WillReceiveProps is an interface defining a callback that will be called with
// the new props before they are accepted and passed to be rendered.
type WillReceiveProps interface {
	ComponentWillReceiveProps(context.Context, prop.Props)
}

// ShouldUpdate is an interface defining callback that is called before render
// determine if re render is necessary.
type ShouldUpdate interface {
	// If this returns false then re rendering for the component is skipped.
	ShouldComponentUpdate(context.Context, prop.Props, state.State) bool
}

// WillUpdate is an interface defining a callback that is called before rendering
type WillUpdate interface {
	// If returned props are not nil, then it will be merged with nextprops then
	// passed to render for rendering.
	ComponentWillUpdate(context.Context, prop.Props, state.State) prop.Props
}

// DidUpdate defines a callback that is invoked after rendering.
type DidUpdate interface {
	ComponentDidUpdate()
}

// DerivedState is an interface which can be used to derive state from props.
type DerivedState interface {
	DeriveState(prop.Props, state.State) state.State
}

// WithContext is an interface used to update the context that is passed to
// component's children.
type WithContext interface {
	WithContext(context.Context) context.Context
}

type QueuedRender struct {
	components *list.List
	mu         sync.RWMutex
	closed     bool
	r          *Renderer
}

func (q *QueuedRender) Push(v Component) {
	q.mu.Lock()
	defer q.mu.Unlock()
	q.components.PushBack(v)
}

// Pop returns the last added component and removes it from the queue.
func (q *QueuedRender) Pop() Component {
	e := q.pop()
	if e != nil {
		return e.Value.(Component)
	}
	return nil
}

func (q *QueuedRender) pop() *list.Element {
	e := q.last()
	q.mu.Lock()
	if e != nil {
		q.components.Remove(e)
	}
	q.mu.Unlock()
	return e
}

func (q *QueuedRender) last() *list.Element {
	q.mu.RLock()
	e := q.components.Back()
	q.mu.RUnlock()
	return e
}

// Last returns the last added component to the queue.
func (q *QueuedRender) Last() Component {
	e := q.last()
	if e != nil {
		return e.Value.(Component)
	}
	return nil
}

func NeqQueueRenderer() *QueuedRender {
	return &QueuedRender{
		components: list.New(),
	}
}

func enqueueRender(cmp Component) {
	if cmp.core().dirty {
		queue.Push(cmp)
		queue.Rerender()
	}
}

// Rerender re renders all enqueued dirty components.
func (q *QueuedRender) Rerender() {
	go q.rerender()
}

func (q *QueuedRender) rerender() {
	for cmp := q.Pop(); cmp != nil; cmp = q.Pop() {
		if cmp.core().dirty {
			q.r.renderComponent(cmp, 0, false, false)
		}
	}
}

func flushMounts() {
	for c := mounts.Back(); c != nil; c = mounts.Back() {
		if cmp, ok := c.Value.(Component); ok {
			if m, ok := cmp.(DidMount); ok {
				m.ComponentDidMount()
			}
		}
		mounts.Remove(c)
	}
}

func recollectNodeTree(node dom.Element, unmountOnly bool) {
	cmp := findComponent(node)
	if cmp != nil {
		unmountComponent(cmp)
	} else {
		if !unmountOnly {
			dom.RemoveNode(node)
		}
		removeChildren(node)
	}
}

// findComponent returns the component that rendered the node element. This
// returns nil if the node wasn't a component.
func findComponent(node dom.Element) Component {
	return nil
}

func unmountComponent(cmp Component) {
	core := cmp.core()
	core.disable = true
	base := core.base
	if wm, ok := cmp.(WillUnmount); ok {
		wm.ComponentWillUnmount()
	}
	core.base = nil
	if core.component != nil {
		unmountComponent(core.component)
	} else if base != nil {
		core.nextBase = base
		dom.RemoveNode(base)
		removeChildren(base)
	}
}

func removeChildren(node dom.Element) {
	node = node.Get("lastChild")
	for {
		if !dom.Valid(node) {
			return
		}
		next := node.Get("previousSibling")
		recollectNodeTree(node, true)
		node = next
	}
}

// UndefinedFunc is a function  that returns a javascript undefined value.
type UndefinedFunc func() value.Value

// Undefined is a work around to allow the library to work with/without wasm
// support.
//
// TODO: find a better way to handle this.
var Undefined UndefinedFunc

// Callback this is supposed to be defined by the package consumers.
var Callback dom.CallbackGenerator

func diffAttributes(node dom.Element, attrs, old []vdom.Attribute) {
	a := mapAtts(attrs)
	b := mapAtts(old)
	for k, v := range b {
		if _, ok := a[k]; !ok {
			dom.SetAccessor(Callback, node, k, v, Undefined(), isSVGMode)
		}
	}
	for k := range a {
		switch k {
		case "children", "innerHTML":
			continue
		default:
			dom.SetAccessor(Callback, node, k, b[k], a[k], isSVGMode)
		}
	}
}

func mapAtts(attrs []vdom.Attribute) map[string]vdom.Attribute {
	m := make(map[string]vdom.Attribute)
	for _, v := range attrs {
		m[v.Key] = v
	}
	return m
}

func diff(ctx context.Context, elem dom.Element, node *vdom.Node, parent dom.Element, mountAll, componentRoot bool) dom.Element {
	if diffLevel == 0 {
		diffLevel++
		// when first starting the diff, check if we're diffing an SVG or within an SVG
		isSVGMode = parent != nil && parent.Type() != value.TypeNull &&
			dom.Valid(parent.Get("ownerSVGElement"))

		// hydration is indicated by the existing element to be diffed not having a
		// prop cache
		hydrating = dom.Valid(elem) && dom.Valid(elem.Get(AttrKey))
	}
	ret := idiff(ctx, elem, node, mountAll, componentRoot)

	// append the element if its a new parent
	if dom.Valid(parent) &&
		!dom.IsEqual(ret.Get("parentNode"), parent) {
		parent.Call("appendChild", ret)
	}
	diffLevel--
	if diffLevel == 0 {
		hydrating = false
		if !componentRoot {
			flushMounts()
		}
	}
	return ret
}

func idiff(ctx context.Context, elem dom.Element, node *vdom.Node, mountAll, componentRoot bool) dom.Element {
	out := elem
	prevSVGMode := isSVGMode
	switch node.Type {
	case vdom.TextNode:
		if dom.Valid(elem) && dom.Valid(elem.Get("splitText")) &&
			dom.Valid(elem.Get("parentNode")) {
			v := elem.Get("nodeValue").String()
			if v != node.Data {
				elem.Set("nodeValue", node.Data)
			}

		} else {
			out = dom.Document.Call("createTextNode", node.Data)
			if dom.Valid(elem) {
				if dom.Valid(elem.Get("parentNode")) {
					elem.Get("parentNode").Call("replaceChild", out, elem)
				}
				recollectNodeTree(elem, true)
			}
		}
		out.Set(AttrKey, true)
		return out
	case vdom.ElementNode:
		if !elements.Valid(node.Data) {
			if node.Data == "svg" {
				isSVGMode = true
			} else if node.Data == "foreignObject" {
				isSVGMode = false
			}
		}
		nodeName := node.Data
		if !dom.Valid(elem) || !isNamedNode(elem, node) {
			out = dom.CreateNode(nodeName)
			if dom.Valid(elem) {
				if dom.Valid(elem.Get("firstChild")) {
					out.Call("appendChild", elem.Get("firstChild"))
				}
				if e := elem.Get("parentNode"); dom.Valid(e) {
					elem.Get("parentNode").Call("replaceChild", out, elem)
				}
				recollectNodeTree(elem, true)
			}
		}
		fc := out.Get("firstChild")
		props := out.Get(AttrKey)
		var old []vdom.Attribute
		if !dom.Valid(props) {
			a := elem.Get("attributes")
			for _, v := range value.Keys(a) {
				old = append(old, vdom.Attribute{
					Key: v,
					Val: a.Get(v).String(),
				})
			}
		}
		if !hydrating && len(node.Children) == 1 &&
			node.Children[0].Type == vdom.TextNode && dom.Valid(fc) &&
			dom.Valid(fc.Get("splitText")) &&
			fc.Get("nextSibling").Type() == value.TypeNull {
			nv := node.Children[0].Data
			fv := fc.Get("nodeValue").String()
			if fv != nv {
				fc.Set("nodeValue", nv)
			}
		} else if len(node.Children) > 0 || dom.Valid(fc) {
			innerDiffMode(ctx, out, node.Children, mountAll, hydrating)
		}
		diffAttributes(out, node.Attr, old)
		isSVGMode = prevSVGMode
		return out
	default:
		panic("Un supported node")
	}
}
func buildComponentFromVNode(ctx context.Context, elem dom.Element, node *vdom.Node, mountAll, componentRoot bool) dom.Element {
	//TODO
	//
	// port buildComponentFromVNode
	return nil
}
func innerDiffMode(ctx context.Context, elem dom.Element, vchildrens []*vdom.Node, mountAll, isHydrating bool) {
	original := elem.Get("childNodes")
	length := original.Get("length").Int()
	keys := make(map[string]dom.Element)
	var children []dom.Element
	var min int
	if length > 0 {
		for i := 0; i < length; i++ {
			child := original.Index(i)
			cmp := findComponent(child)
			var key prop.NullString
			if cmp != nil {
				key = cmp.core().key
			}
			if !key.IsNull {
				keys[key.Value] = child
			} else {
				var x bool
				if cmp != nil || dom.Valid(child.Get("splitText")) {
					v := child.Get("nodeValue").String()
					v = strings.TrimSpace(v)
					if isHydrating {
						x = v != ""
					} else {
						x = true
					}
				} else {
					x = isHydrating
				}
				if x {
					children = append(children, child)
				}
			}
		}
	}
	for i := 0; i < len(vchildrens); i++ {
		vchild := vchildrens[i]
		key := vchild.Key()
		var child dom.Element
		if key != "" {
			if ch, ok := keys[key]; ok {
				delete(keys, key)
				child = ch
			}
		} else if min < len(children) {
			for j := min; j < len(children); j++ {
				c := children[j]
				if c != nil && dom.Valid(c) && isSameNodeType(c, vchild, isHydrating) {
					child = c
					children[j] = nil
					if j == min {
						min++
					}
					break
				}
			}
		}
		child = idiff(ctx, child, vchild, mountAll, false)
		f := original.Index(i)
		if dom.Valid(child) && !dom.IsEqual(child, elem) && !dom.IsEqual(child, f) {
			if f.Type() == value.TypeNull {
				elem.Call("appendChild", child)
			} else if dom.IsEqual(child, f.Get("nextSibling")) {
				dom.RemoveNode(f)
			} else {
				elem.Call("insertBefore", child, f)
			}
		}
	}

	// removing unused keyed  children
	for _, v := range keys {
		recollectNodeTree(v, false)
	}
	for i := min; i < len(children); i++ {
		ch := children[i]
		if ch != nil {
			recollectNodeTree(ch, false)
		}
	}
}

// isSameNodeType compares elem to vnode and returns true if thy are of the same
// type.
//
// There are only two types of nodes supported , TextNode and ElementNode.
func isSameNodeType(elem dom.Element, vnode *vdom.Node, isHydrating bool) bool {
	switch vnode.Type {
	case vdom.TextNode:
		return dom.Valid(elem.Get("splitText"))
	case vdom.ElementNode:
		return isNamedNode(elem, vnode)
	default:
		return false
	}
}

// isNamedNode compares elem to vnode to see if elem was created from the
// virtual node of the same type as vnode..
func isNamedNode(elem dom.Element, vnode *vdom.Node) bool {
	v := elem.Get("normalizedNodeName")
	if dom.Valid(v) {
		name := v.String()
		return name == vnode.Data
	}
	return false
}
