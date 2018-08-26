package vected

import (
	"context"
	"reflect"

	"github.com/gernest/vected/vdom/value"

	"github.com/gernest/vected/prop"
	"github.com/gernest/vected/state"
	"github.com/gernest/vected/vdom"
	"github.com/gernest/vected/vdom/dom"
)

const (
	componentKey         = "_component"
	componentConstructor = "_componentConstructor"
)

// Creates a new component instance. The component is assigned a unique id and
// cached for future retrieval.
//
// caching is important because we can't pass object references to the dom yet,
// instead we will pass  the id which will be used to reference the
// component.
func (v *Vected) createComponent(ctx context.Context, cmp Component, props prop.Props) Component {
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
	core.id = idPool.Get().(int)
	v.cache[core.id] = ncmp
	return ncmp
}

// setProps sets cmp props and possibly re renders. Props can contain key,ref
// props where key will be registered as the component key and ref is a function
// receiving an interface{}, the ref function is a callback which will be passed
// either component's instance if it is a higher order component or dom.Element
// if it is a regular dom node.
//
// Disabled components are ignored.
func (v *Vected) setProps(ctx context.Context, cmp Component, props prop.Props, mode RenderMode, mountAll bool) {
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
			v.renderComponent(cmp, Sync, mountAll, false)
		} else {
			v.enqueueRender(cmp)
		}
	}
	if core.ref != nil {
		core.ref(cmp)
	}
}

func (v *Vected) isHigherOrder(node *vdom.Node) bool {
	if node.Type == vdom.ElementNode {
		if _, ok := v.Components[node.Data]; ok {
			return true
		}
	}
	return false
}

func (v *Vected) getComponent(node *vdom.Node) Component {
	return v.Components[node.Data]
}

func (v *Vected) renderComponent(cmp Component, mode RenderMode, mountAll bool, isChild bool) {
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
		childComponent := v.getComponent(rendered)
		var toUnmount Component
		var base dom.Element
		if v.isHigherOrder(rendered) {
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
				v.setProps(context, inst, childProps, Sync, false)
			} else {
				// We must create a new initialChildComponent and set the current cmp as
				// parent.
				toUnmount = inst
				inst = v.createComponent(context, childComponent, childProps)
				core.component = inst
				instanceCore := inst.core()
				instanceCore.component = inst
				if instanceCore.nextBase == nil {
					instanceCore.nextBase = nextBase
				}
				instanceCore.parentComponent = cmp
				v.setProps(context, inst, childProps, No, false)
				v.renderComponent(inst, Sync, mountAll, true)
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
					cbase.Set(componentKey, 0)
				}
				var parent dom.Element
				if dom.Valid(initialBase) {
					parent = initialBase.Get("parentNode")
				}
				base = v.diff(context, cbase, rendered, parent, mountAll || !dom.Valid(isUpdate), true)
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
					v.removeComponentRef(initialBase)
					v.recollectNodeTree(initialBase, false)
				}
			}
		}
		if toUnmount != nil {
			v.unmountComponent(toUnmount)
		}
		core.base = base
		if dom.Valid(base) && !isChild {
			componentRef := cmp
			t := cmp
			for {
				t = t.core().parentComponent
				if t == nil {
					break
				}
				t.core().base = base
				componentRef = t
			}
			v.addComponentRef(base, componentRef)
		}
	}
	if !dom.Valid(isUpdate) || mountAll {
		//TODO mounts.unshift(component);
	} else if !skip {
		// Ensure that pending componentDidMount() hooks of child components
		// are called before the componentDidUpdate() hook in the parent.
		// Note: disabled as it causes duplicate hooks, see https://github.com/developit/preact/issues/750
		// flushMounts();
		if u, ok := cmp.(DidUpdate); ok {
			u.ComponentDidUpdate(prevProps, prevState)
		}
	}
	if v.diffLevel == 0 && !isChild {
		v.flushMounts()
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
func (v *Vected) findComponent(node dom.Element) Component {
	if dom.Valid(node) {
		id := node.Get(componentKey)
		if id.Type() == value.TypeNumber {
			i := id.Int()
			if c, ok := v.cache[i]; ok {
				return c
			}
		}
	}
	return nil
}

// removeComponentRef removes the reference to a component from the dom element.
func (v *Vected) removeComponentRef(e dom.Element) {
	if dom.Valid(e) {
		id := e.Get(componentKey)
		if id.Type() == value.TypeNumber {
			i := id.Int()
			v.refs[i]--
		}
		e.Set(componentKey, 0)
	}
}

func (v *Vected) addComponentRef(e dom.Element, cmp Component) {
	if dom.Valid(e) {
		e.Set(componentKey, cmp.core().id)
		e.Set(componentConstructor, cmp.core().constructor)
		v.refs[cmp.core().id]++
	}
}

func (v *Vected) unmountComponent(cmp Component) {
	core := cmp.core()
	core.disable = true
	base := core.base
	if wm, ok := cmp.(WillUnmount); ok {
		wm.ComponentWillUnmount()
	}
	core.base = nil
	if core.component != nil {
		v.unmountComponent(core.component)
	} else if base != nil {
		core.nextBase = base
		dom.RemoveNode(base)
		v.removeChildren(base)
	}
}

func (v *Vected) removeChildren(node dom.Element) {
	node = node.Get("lastChild")
	for {
		if !dom.Valid(node) {
			break
		}
		next := node.Get("previousSibling")
		v.recollectNodeTree(node, true)
		node = next
	}
}

// componentCache this stores references to rendered components.
type componentCache struct {
	cache map[int64]Component
	refs  map[int64]int64
}
