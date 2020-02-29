// package greact is a component based frontend framework for golang. This
// framework delivers high performance and responsive ui.
//
// This relies on the experimental wasm api to interact with dom. The project
// started as a port of preact to go, but has since evolved. It still borrows a
// similar API from react/preact.
package greact

import (
	"container/list"
	"context"
	"fmt"
	"reflect"
	"regexp"
	"strings"
	"sync"

	"github.com/gernest/greact/elements"
)

// used in code generation
const (
	ID         = "greact"
	TemplateFn = "Template"
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

// AttrKey is a key used to store node's attributes/props
const AttrKey = "__vected_attr__"
const svg = "http://www.w3.org/2000/svg'"

// This tracks the last id issued. We use sync pool to reuse component id's.
//
// TODO: come up with a better way that can scale.
var idx int
var idPool = &sync.Pool{
	New: func() interface{} {
		idx++
		return idx
	},
}

// Component is an interface which defines a unit of user interface.There are
// two ways to satisfy this interface.
//
// You can define a struct that embeds Core and implements Templater interface
// like this.
// 	type Foo struct {
// 		greact.Core
// 	}
//
// 	func (f Foo) Template() string {
// 		return `<div />`
// 	}
//
// Then run
//	vected render /path/to/foo package
// The command will automatically generate the Render method for you. For the
// example above it will generate something like this.
//
// 	var H = New
// 	func (f Foo) Render(ctx context.Context, props Props, state state.State) *Node {
// 		return H(3, "", "div", nil)
// 	}
//
// The second way is to implement Render method. I recommend you stick with only
// implementing Templater interface which is less error prone and reduces
// verbosity.
type Component interface {
	Render(context.Context, Props, State) *Node
	core() *Core
}

// Resource is an interface for a resource that will need to be freed manually,
// such as callback..
type Resource interface {
	Release()
}

// RemoveNode removes node from its parent if attached.
func RemoveNode(node Value) {
	parent := node.Get("parentNode")
	if Valid(parent) {
		parent.Call("removeChild", node)
	}
}

// Element is an alias for the dom node.
type Element = Value

// HasProperty returns true if e has property.
func HasProperty(e Element, v string) bool {
	return e.Call("hasOwnProperty", v).Bool()
}

// Templater is an interface for describing components with xml like markup. The
// markup is similar to jsx but tailored towards go constructs.
type Templater interface {
	// Template returns jsx like string template. The template is compiled to
	// Render method of the struct that implements this interface..
	Template() string
}

// Constructor is an interface for creating new component instance.
type Constructor interface {
	New(Props) Component
}

// Core is th base struct that every struct that wants to implement Component
// interface must embed.
//
// This is used to make Props available to the component.
type Core struct {
	id int

	// constructor is the name of the higher order component. This is can be
	// defined when registering components with greact.Register. This library uses
	// golang.org/x/net/html for parsing component template, which defaults all
	// elements to lower case, so the constructor name must be lower case.
	constructor string

	context context.Context
	props   Props
	state   State

	prevContext context.Context
	prevProps   Props
	prevState   State

	// A list of functions that will be called after the component has been
	// rendered.
	renderCallbacks []func()

	// This is the instance of the child component.
	component       Component
	parentComponent Component

	// The base dom node on which the component was rendered. When this is set it
	// signals for an update, this will be nil if the component hasn't been
	// rendered yet.
	base     Element
	nextBase Element

	dirty   bool
	disable bool

	// Optional prop that must be unique among child components for efficient
	// rendering of lists.
	key string

	// This is a callback used to receive instance of Component or the Dom element.
	// after they have been mounted.
	ref func(interface{})

	// priority this is a number indicating how important this component is in the
	// re rendering queue. The higher the number the more urgent re renders.
	priority int

	enqueue *queuedRender
}

func (c *Core) core() *Core { return c }

// SetState updates component state and schedule re rendering.
func (c *Core) SetState(newState State, callback ...func()) {
	prev := c.prevState
	c.prevState = newState
	c.state = MergeState(prev, newState)
	if len(callback) > 0 {
		c.renderCallbacks = append(c.renderCallbacks, callback...)
	}
	c.enqueue.enqueueCore(c)
}

// Props returns current props.s
func (c *Core) Props() Props {
	return c.props
}

// State returns current state.
func (c *Core) State() State {
	return c.state
}

// Context returns current context.
func (c *Core) Context() context.Context {
	return c.context
}

// InitState is an interface for exposing initial state.
// Component should implement this interface if they want to set initial state
// when the component is first created before being rendered.
type InitState interface {
	InitState() State
}

// InitProps is an interface for exposing default props. This will be merged
// with other props before being sent to render.
type InitProps interface {
	InitProps() Props
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
	ComponentWillReceiveProps(context.Context, Props)
}

// ShouldUpdate is an interface defining callback that is called before render
// determine if re render is necessary.
type ShouldUpdate interface {
	// If this returns false then re rendering for the component is skipped.
	ShouldComponentUpdate(context.Context, Props, State) bool
}

// WillUpdate is an interface defining a callback that is called before rendering
type WillUpdate interface {
	// If returned props are not nil, then it will be merged with nextprops then
	// passed to render for rendering.
	ComponentWillUpdate(context.Context, Props, State) Props
}

// DidUpdate defines a callback that is invoked after rendering.
type DidUpdate interface {
	ComponentDidUpdate(prevProps Props, prevState State)
}

// DerivedState is an interface which can be used to derive state from props.
type DerivedState interface {
	DeriveState(Props, State) State
}

// WithContext is an interface used to update the context that is passed to
// component's children.
type WithContext interface {
	WithContext(context.Context) context.Context
}

type queuedRender struct {
	components *list.List
	mu         sync.RWMutex
	closed     bool
	v          *Vected
}

func newQueuedRender(v *Vected) *queuedRender {
	return &queuedRender{
		components: list.New(),
		v:          v,
	}
}

func (q *queuedRender) Push(v Component) {
	q.mu.Lock()
	defer q.mu.Unlock()
	q.components.PushBack(v)
}

// Pop returns the last added component and removes it from the queue.
func (q *queuedRender) Pop() Component {
	e := q.pop()
	if e != nil {
		return e.Value.(Component)
	}
	return nil
}

func (q *queuedRender) pop() *list.Element {
	e := q.last()
	q.mu.Lock()
	if e != nil {
		q.components.Remove(e)
	}
	q.mu.Unlock()
	return e
}

func (q *queuedRender) last() *list.Element {
	q.mu.RLock()
	e := q.components.Back()
	q.mu.RUnlock()
	return e
}

// Last returns the last added component to the queue.
func (q *queuedRender) Last() Component {
	e := q.last()
	if e != nil {
		return e.Value.(Component)
	}
	return nil
}

// Rerender re renders all enqueued dirty components async.
func (q *queuedRender) Rerender() {
	go q.rerender()
}

func (q *queuedRender) enqueue(cmp Component) {
	if !cmp.core().dirty {
		cmp.core().dirty = true
	}
	q.Push(cmp)
	q.Rerender()
}

func (q *queuedRender) enqueueCore(core *Core) {
	cmp := q.v.cache[core.id]
	if !cmp.core().dirty {
		cmp.core().dirty = true
	}
	q.Push(cmp)
	q.Rerender()
}

func (q *queuedRender) rerender() {
	for cmp := q.Pop(); cmp != nil; cmp = q.Pop() {
		if cmp.core().dirty {
			q.v.renderComponent(cmp, 0, false, false)
		}
	}
}

// CallbackGenerator is a function that returns callbacks.
type CallbackGenerator func(fn func([]Value)) Resource

// Vected this is the ultimate struct that ports preact to work with go/was.
// This is not a direct port, the two languages are different. Although some
// portion of the methods are a direct translation, the working differs from
// preact.
type Vected struct {

	// queue this is q queue of components that are supposed to be rendered
	// asynchronously.
	queue *queuedRender

	// Component this is a mapping of component name to component instance. The
	// name is not case sensitive and must not be the same as the standard html
	// elements.
	//
	// This means you cant have a component with name div,p,h1 etc. Remember that
	// it  is case insensitive so Div is also not allowed.
	//
	// In case you are not so sure, use the github.com/gernest/elements package to
	// check if the name is a valid component.
	//
	// The registered components won't be used as is, instead new instance will be
	// created so please don't pass component's which have state in them
	// (initialized field values etc) here, because they will be ignored.
	components map[string]Component

	// Is the browser's document object. New document elements will be created from
	// this.
	Document Element

	// mounts is a list of components ready to be mounted.
	mounts *list.List

	isSVGMode bool
	hydrating bool
	diffLevel int

	cache map[int]Component
	refs  map[int]int

	cb CallbackGenerator
}

// New returns an initialized Vected instance.
func New() *Vected {
	v := &Vected{
		cache:      make(map[int]Component),
		refs:       make(map[int]int),
		mounts:     list.New(),
		components: make(map[string]Component),
	}
	v.queue = newQueuedRender(v)
	return v
}

func (v *Vected) enqueueRender(cmp Component) {
	if cmp.core().dirty {
		v.queue.Push(cmp)
		v.queue.Rerender()
	}
}

func (v *Vected) flushMounts() {
	for c := v.mounts.Back(); c != nil; c = v.mounts.Back() {
		if cmp, ok := c.Value.(Component); ok {
			if m, ok := cmp.(DidMount); ok {
				m.ComponentDidMount()
			}
		}
		v.mounts.Remove(c)
	}
}

func (v *Vected) recollectNodeTree(node Element, unmountOnly bool) {
	cmp := v.findComponent(node)
	if cmp != nil {
		v.unmountComponent(cmp)
	} else {
		if !unmountOnly || !Valid(node.Get(AttrKey)) {
			RemoveNode(node)
		}
		v.removeChildren(node)
	}
}

// UndefinedFunc is a function  that returns a javascript undefined value.
type UndefinedFunc func() Value

// Undefined is a work around to allow the library to work with/without wasm
// support.
//
// TODO: find a better way to handle this.
var Undefined UndefinedFunc

func (v *Vected) diffAttributes(node Element, attrs, old []Attribute) {
	a := mapAtts(attrs)
	b := mapAtts(old)
	for k, val := range b {
		if _, ok := a[k]; !ok {
			setAccessor(v.cb, node, k, val, Undefined(), v.isSVGMode)
		}
	}
	for k := range a {
		switch k {
		case "children", "innerHTML":
			continue
		default:
			setAccessor(v.cb, node, k, b[k], a[k], v.isSVGMode)
		}
	}
}

func mapAtts(attrs []Attribute) map[string]Attribute {
	m := make(map[string]Attribute)
	for _, v := range attrs {
		m[v.Key] = v
	}
	return m
}

func (v *Vected) diff(ctx context.Context, elem Element, node *Node, parent Element, mountAll, componentRoot bool) Element {
	if v.diffLevel == 0 {
		v.diffLevel++
		// when first starting the diff, check if we're diffing an SVG or within an SVG
		v.isSVGMode = parent.IsNull() && parent.IsNull() &&
			Valid(parent.Get("ownerSVGElement"))

		// hydration is indicated by the existing element to be diffed not having a
		// prop cache
		v.hydrating = Valid(elem) && Valid(elem.Get(AttrKey))
	}
	ret := v.idiff(ctx, elem, node, mountAll, componentRoot)

	// append the element if its a new parent
	if Valid(parent) &&
		!IsEqual(ret.Get("parentNode"), parent) {
		parent.Call("appendChild", ret)
	}
	v.diffLevel--
	if v.diffLevel == 0 {
		v.hydrating = false
		if !componentRoot {
			v.flushMounts()
		}
	}
	return ret
}

func (v *Vected) idiff(ctx context.Context, elem Element, node *Node, mountAll, componentRoot bool) Element {
	out := elem
	prevSVGMode := v.isSVGMode
	switch node.Type {
	case TextNode:
		if Valid(elem) && Valid(elem.Get("splitText")) &&
			Valid(elem.Get("parentNode")) {
			v := elem.Get("nodeValue").String()
			if v != node.Data {
				elem.Set("nodeValue", node.Data)
			}

		} else {
			out = v.Document.Call("createTextNode", node.Data)
			if Valid(elem) {
				if Valid(elem.Get("parentNode")) {
					elem.Get("parentNode").Call("replaceChild", out, elem)
				}
				v.recollectNodeTree(elem, true)
			}
		}
		out.Set(AttrKey, true)
		return out
	case ElementNode:
		fmt.Printf("rendering %s\n", node.Data)
		if v.isHigherOrder(node) {
			return v.buildComponentFromVNode(ctx, elem, node, mountAll, false)
		}
		if !elements.Valid(node.Data) {
			if node.Data == "svg" {
				v.isSVGMode = true
			} else if node.Data == "foreignObject" {
				v.isSVGMode = false
			}
		}
		nodeName := node.Data
		if !Valid(elem) || !isNamedNode(elem, node) {
			out = v.CreateNode(nodeName)
			if Valid(elem) {
				if Valid(elem.Get("firstChild")) {
					out.Call("appendChild", elem.Get("firstChild"))
				}
				if e := elem.Get("parentNode"); Valid(e) {
					elem.Get("parentNode").Call("replaceChild", out, elem)
				}
				v.recollectNodeTree(elem, true)
			}
		}
		fc := out.Get("firstChild")
		props := out.Get(AttrKey)
		var old []Attribute
		if !Valid(props) {
			a := out.Get("attributes")
			for _, v := range Keys(a) {
				old = append(old, Attribute{
					Key: v,
					Val: a.Get(v).String(),
				})
			}
		}
		if !v.hydrating && len(node.Children) == 1 &&
			node.Children[0].Type == TextNode && Valid(fc) &&
			Valid(fc.Get("splitText")) &&
			fc.Get("nextSibling").IsNull() {
			nv := node.Children[0].Data
			fv := fc.Get("nodeValue").String()
			if fv != nv {
				fc.Set("nodeValue", nv)
			}
		} else if len(node.Children) > 0 || Valid(fc) {
			v.innerDiffMode(ctx, out, node.Children, mountAll, v.hydrating)
		}
		v.diffAttributes(out, node.Attr, old)
		v.isSVGMode = prevSVGMode
		return out
	default:
		panic("Un supported node")
	}
}

func (v *Vected) buildComponentFromVNode(ctx context.Context, elem Element, node *Node, mountAll, componentRoot bool) Element {
	c := v.findComponent(elem)
	originalComponent := c
	oldElem := elem
	isDirectOwner := c != nil && c.core().constructor == node.Data
	isOwner := isDirectOwner
	props := getNodeProps(node)
	for {
		if c != nil && !isOwner {
			c = c.core().parentComponent
			if c != nil {
				isOwner = c.core().constructor == node.Data
				continue
			}
		}
		break
	}
	if c != nil && isOwner && (!mountAll || c.core().component != nil) {
		v.setProps(ctx, c, props, Async, mountAll)
		elem = c.core().base
	} else {
		if originalComponent != nil && !isDirectOwner {
			v.unmountComponent(originalComponent)
			elem = Null()
			oldElem = Null()
		}
		c = v.createComponentByName(ctx, node.Data, props)
		if !elem.IsNull() && !Valid(c.core().nextBase) {
			c.core().nextBase = elem
			oldElem = Null()
		}
		v.setProps(ctx, c, props, Sync, mountAll)
		elem = c.core().base
		if !oldElem.IsNull() && !IsEqual(elem, oldElem) {
			//TODO dereference the component.
			oldElem.Set(componentKey, 0)
			v.recollectNodeTree(oldElem, false)
		}
	}
	return elem
}

func (v *Vected) innerDiffMode(ctx context.Context, elem Element, vchildrens []*Node, mountAll, isHydrating bool) {
	original := elem.Get("childNodes")
	length := original.Get("length").Int()
	keys := make(map[string]Element)
	var children []Element
	var min int
	if length > 0 {
		for i := 0; i < length; i++ {
			child := original.Index(i)
			cmp := v.findComponent(child)
			var key string
			if cmp != nil {
				key = cmp.core().key
			}
			if key != "" {
				keys[key] = child
			} else {
				var x bool
				if cmp != nil || Valid(child.Get("splitText")) {
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
		var child Element
		if key != "" {
			if ch, ok := keys[key]; ok {
				delete(keys, key)
				child = ch
			}
		} else if min < len(children) {
			for j := min; j < len(children); j++ {
				c := children[j]
				if Valid(c) && isSameNodeType(c, vchild, isHydrating) {
					child = c
					children[j] = Null()
					if j == min {
						min++
					}
					break
				}
			}
		}
		child = v.idiff(ctx, child, vchild, mountAll, false)
		f := original.Index(i)
		if Valid(child) && !IsEqual(child, elem) && !IsEqual(child, f) {
			if !Valid(f) {
				elem.Call("appendChild", child)
			} else if IsEqual(child, f.Get("nextSibling")) {
				RemoveNode(f)
			} else {
				elem.Call("insertBefore", child, f)
			}
		}
	}

	// removing unused keyed  children
	for _, val := range keys {
		v.recollectNodeTree(val, false)
	}
	for i := min; i < len(children); i++ {
		ch := children[i]
		if !ch.IsNull() {
			v.recollectNodeTree(ch, false)
		}
	}
}

// isSameNodeType compares elem to vnode and returns true if thy are of the same
// type.
//
// There are only two types of nodes supported , TextNode and ElementNode.
func isSameNodeType(elem Element, vnode *Node, isHydrating bool) bool {
	switch vnode.Type {
	case TextNode:
		return Valid(elem.Get("splitText"))
	case ElementNode:
		return isNamedNode(elem, vnode)
	default:
		return false
	}
}

// isNamedNode compares elem to vnode to see if elem was created from the
// virtual node of the same type as vnode..
func isNamedNode(elem Element, vnode *Node) bool {
	v := elem.Get("normalizedNodeName")
	if Valid(v) {
		name := v.String()
		return name == vnode.Data
	}
	return false
}

// Render renders vected component.
func (v *Vected) Render(vnode *Node, parent Element, merge ...Element) Element {
	var elem Element
	if len(merge) > 0 {
		elem = merge[0]
	}
	return v.diff(context.Background(), elem, vnode, parent, false, false)
}

// RenderComponent compiles component cmp and renders it.
func (v *Vected) RenderComponent(cmp string, parent Element, merge ...Element) (Element, error) {
	node, err := ParseString(cmp)
	if err != nil {
		return Null(), err
	}
	return v.Render(node, parent, merge...), nil
}

// Register add cmp instance to a map of known higher order components There is
// no other way the Vected instance can know about Higher orrder components.
//
// So, components must be registered before being used. For technical reason the
// name is case insensitive, so you can register somecomponent and use
// SomeComponent in your templates and it will work like a charm.
//
// The reason behind this because the x/html library used to parse the templates
// resolves or element names to lowercase.
func (v *Vected) Register(name string, cmp Component) {
	name = strings.ToLower(name)
	if v.components == nil {
		v.components = make(map[string]Component)
	}
	v.components[name] = cmp
}

// CreateNode creates a dom element.
func (v *Vected) CreateNode(name string) Element {
	fmt.Printf("creating a new node %s\n", name)
	node := v.Document.Call("createElement", name)
	node.Set("normalizedNodeName", name)
	return node
}

// CreateSVGNode creates svg dom element.
func (v *Vected) CreateSVGNode(doc Value, name string) Element {
	node := v.Document.Call("createElementNS", svg, name)
	node.Set("normalizedNodeName", name)
	return node
}

var xlink = regexp.MustCompile(`^xlink:?`)

// setAccessor Set a named attribute on the given Node, with special behavior
// for some names and event handlers. If `value` is `null`, the
// attribute/handler will be removed.
// node An element to mutate
//
// name The name/key to set, such as an event or attribute name
// old The last value that was set for this name/node pair
// value An attribute value, such as a function to be used as an event handler
// isSVG Are we currently diffing inside an svg?
func setAccessor(gen CallbackGenerator, node Element, name string, old, val interface{}, isSVG bool) {
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
	case "ref":
		applyRef(old, nil)
		applyRef(val, node)
	case "dangerouslySetInnerHTML":
		node.Set("innerHTML", val)
	default:
		switch {
		case strings.HasPrefix(name, "on"):
			useCapture := name != strings.TrimSuffix(name, "Capture")
			name = eventName(name)
			if ev, ok := val.(func([]Value)); ok {
				cb := gen(ev)
				if old == nil {
					node.Call("addEventListener", name, cb, useCapture)
					// To release resources allocated for the callback we keep track of of all
					// callbacks added to this node.
					//
					// These can be later removed by calling the functions.
					var release Resource
					release = gen(func(args []Value) {
						node.Call("removeEventListener", name, cb, useCapture)
						cb.Release()
						release.Release()
					})
					releaseList := node.Get("_listeners")
					if releaseList.IsUndefined() {
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
				if Valid(releaseList) {
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
			ns := isSVG && (name != xlink.ReplaceAllString(name, ""))
			isFalse := func() bool {
				if v, ok := val.(bool); ok {
					return !v
				}
				return false
			}
			if val == nil || isFalse() {
				if ns {
					name := strings.ToLower(xlink.ReplaceAllString(name, ""))
					node.Call("removeAttributeNS", "http://www.w3.org/1999/xlink", name)
				} else {
					node.Call("removeAttribute", name)
				}
			} else {
				e := reflect.ValueOf(val)
				if validSVGValue(e.Kind()) {
					if ns {
						name := strings.ToLower(xlink.ReplaceAllString(name, ""))
						node.Call("setAttributeNS", "http://www.w3.org/1999/xlink", name, val)
					} else {
						node.Call("setAttribute", name, val)
					}
				}
			}
		}
	}
}

func applyRef(ref, value interface{}) {
	if r, ok := ref.(Element); ok {
		r.Set("current", value)
	}
}

func validSVGValue(v reflect.Kind) bool {
	switch v {
	case reflect.Int, reflect.Float64, reflect.String:
		return true
	default:
		return false
	}
}

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
