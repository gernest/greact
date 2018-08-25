package vected

import (
	"container/list"
	"context"
	"reflect"

	"github.com/gernest/vected/prop"
	"github.com/gernest/vected/state"
	"github.com/gernest/vected/vdom"
	"github.com/gernest/vected/vdom/dom"
)

var recyclerComponents = list.New()

func createComponent(ctx context.Context, cmp Component, props prop.Props) Component {
	var ncmp Component
	if in, ok := cmp.(Constructor); ok {
		ncmp = in.New()
	} else {
		// we use reflection to create a new component
		v := reflect.ValueOf(cmp)
		if v.Kind() != reflect.Ptr || v.Elem().Kind() != reflect.Struct {
			panic("component must be pointer to struct")
		}
		ncmp = reflect.New(v.Elem().Type()).Interface().(Component)
	}
	core := ncmp.core()
	core.context = ctx
	core.props = props
	return ncmp
}

// SetProps sets cmp props and possibly re renders
func (r *Renderer) SetProps(ctx context.Context, cmp Component, props prop.Props, mode RenderMode, mountAll bool) {
	core := cmp.core()
	if core.disable {
		return
	}
	ref := props["ref"]
	if fn, ok := ref.(func(interface{})); ok {
		core.ref = fn
	}
	core.key = props.String("key")
	delete(props, "key")
	delete(props, "ref")
	_, ok := cmp.(DerivedState)
	if !ok {
		if core.base == nil || mountAll {
			if m, ok := cmp.(WillMount); ok {
				m.ComponentWillMount()
			}
		} else if m, ok := cmp.(WillReceiveProps); ok {
			m.ComponentWillReceiveProps(ctx, props)
		}
	}
	if core.prevProps == nil {
		core.prevProps = core.props
	}
	core.props = props
	core.disable = false
	if mode != No {
		if mode == Sync {
			r.renderComponent(cmp, Sync, mountAll, false)
		} else {
			enqueueRender(cmp)
		}
	}
	if core.ref != nil {
		core.ref(cmp)
	}
}

type Renderer struct {
	components map[string]Component
	queue      *QueuedRender
}

func (r *Renderer) isHigherOrder(node *vdom.Node) bool {
	if node.Type == vdom.ElementNode {
		if _, ok := r.components[node.Data]; ok {
			return true
		}
	}
	return false
}

func (r *Renderer) getComponent(node *vdom.Node) Component {
	return r.components[node.Data]
}

func (r *Renderer) renderComponent(cmp Component, mode RenderMode, mountAll bool, isChild bool) {
	core := cmp.core()
	if core.disable {
		return
	}
	props := core.props
	xstate := core.state
	context := core.context
	prevProps := core.prevProps
	if prevProps == nil {
		prevProps = props
	}
	prevState := core.prevState
	if prevState == nil {
		prevState = xstate
	}
	prevContext := core.prevContext
	if prevContext == nil {
		prevContext = context
	}
	isUpdate := core.base
	nextBase := core.nextBase
	initialBase := isUpdate
	if initialBase == nil {
		initialBase = nextBase
	}
	initialChildComponent := core.component
	var (
		skip  bool
		inst  Component
		cbase dom.Element
	)
	if c, ok := cmp.(DerivedState); ok {
		xstate = state.Merge(xstate, c.DeriveState(props, xstate))
		core.state = xstate
	}
	if isUpdate != nil {
		core.props = prevProps
		core.state = prevState
		core.context = prevContext

		up, ok := cmp.(ShouldUpdate)
		if mode != Force && ok &&
			!up.ShouldComponentUpdate(context, props, xstate) {
			skip = true
		} else if w, ok := cmp.(WillUpdate); ok {
			w.ComponentWillUpdate(context, props, xstate)
		}
		core.props = props
		core.state = xstate
		core.context = context
	}
	core.prevProps = nil
	core.prevState = nil
	core.prevContext = nil
	core.nextBase = nil
	core.dirty = false

	if !skip {
		rendered := cmp.Render(context, props, xstate)
		if ctx, ok := cmp.(WithContext); ok {
			context = ctx.WithContext(context)
		}
		childComponent := r.getComponent(rendered)
		var toUnmount Component
		var base dom.Element
		if r.isHigherOrder(rendered) {
			childProps := getNodeProps(rendered)
			inst = initialChildComponent

			var validForProps = func() bool {
				if inst != nil && sameConstructor(inst, childComponent) {
					key := childProps.String("key")
					ikey := inst.core().key
					if !key.IsNull && ikey.IsNull && key.Value == ikey.Value {
						return true
					}
				}
				return false
			}
			if validForProps() {
				r.SetProps(context, inst, childProps, Sync, false)
			} else {
				toUnmount = inst
				inst = createComponent(context, childComponent, childProps)
				icore := inst.core()
				icore.component = inst
				if icore.nextBase == nil {
					icore.nextBase = nextBase
				}
				icore.parentComponent = cmp
				r.SetProps(context, inst, childProps, No, false)
				r.renderComponent(inst, Sync, mountAll, true)
			}
			base = inst.core().base
		} else {
			cbase = initialBase
			toUnmount = initialChildComponent
			if toUnmount != nil {
				cbase = nil
				core.component = nil
			}
			if initialBase != nil || mode == Sync {
				if cbase != nil {
					//TODO : destroy the reference to the child component.
					//cbase._component=nil
					//
					// We can't do this right now because tehre is no way to move objects
					// references between wasm(go) and js
					//
					// One option is to have a unique id for every component and only store the
					// id in the dom node that is assigned to a component ijnstance.
				}
				var parent dom.Element
				if dom.Valid(initialBase) {
					parent = initialBase.Get("parentNode")
				}
				base = diff(context, cbase, rendered, parent, mountAll || !dom.Valid(isUpdate), true)
			}
		}
		if dom.Valid(initialBase) &&
			!dom.IsEqual(base, initialBase) {
			// TODO: add inst!==initialChildComponent to the if condition
			// Go doesnt support that operation on structs so I will need to use
			// reflection for that or comeup with something else.
			baseParent := initialBase.Get("parentNode")
			if dom.Valid(baseParent) && !dom.IsEqual(base, baseParent) {
				baseParent.Call("replaceChild", base, initialBase)

				if toUnmount == nil {
					//TODO : add initialBase._component = null;
					//
					recollectNodeTree(initialBase, false)
				}
			}
		}
		if toUnmount != nil {
			unmountComponent(toUnmount)
		}
		core.base = base
		if dom.Valid(base) && !isChild {

		}
	}
}

func getNodeProps(node *vdom.Node) prop.Props {
	props := make(prop.Props)
	for _, v := range node.Attr {
		props[v.Key] = v.Val
	}
	props["children"] = node.Children
	return props
}

// sameConstructor returns true if both a and b were created form ain instance
// of the same type.
//
// There is only one way a user can satisfy the Component interface which is by
// embedding Core struct. So, we know bya a fact that it must be a pointer to
// the struct.
func sameConstructor(a, b Component) bool {
	v1 := reflect.ValueOf(a)
	v2 := reflect.ValueOf(b)
	switch v1.Kind() {
	case reflect.Ptr:
		if v2.Kind() == reflect.Ptr {
			v1 = v1.Elem()
			v2 = v2.Elem()
			return v1.Type() == v2.Type()
		}
	}
	return false
}

// findComponent returns the component that rendered the node element. This
// returns nil if the node wasn't a component.
//
// There is a challenge of storing references to higher order components. We
// can't simply attach Components to the dom elements since there is no way to
// move objects between the dom and go runtime(wasm) yet.
//
// To work around this, a simple reference counting is used to decide what
// componen's to keep around long enough.
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
			break
		}
		next := node.Get("previousSibling")
		recollectNodeTree(node, true)
		node = next
	}
}

// componentCache this stores references to rendered components.
type componentCache struct {
	cache map[int64]Component
	refs  map[int64]int64
}
