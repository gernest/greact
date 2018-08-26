package vected

// var H = vdom.New
// var HA = vdom.Attr
// var HAT = vdom.Attrs

// type Hello struct {
// 	Core
// }

// func (*Hello) New() Component {
// 	return &Hello{}
// }

// // Render implements vected.Renderer interface.
// func (h *Hello) Render(ctx context.Context, props prop.Props, state state.State) *vdom.Node {
// 	return H(3, "", "div", nil,
// 		H(1, "", "hello,world", nil))
// }

// type HelloPlain struct {
// 	Core
// }

// // Render implements vected.Renderer interface.
// func (h HelloPlain) Render(ctx context.Context, props prop.Props, state state.State) *vdom.Node {
// 	return H(3, "", "div", nil,
// 		H(1, "", "hello,world", nil))
// }

// func TestCreateComponent(t *testing.T) {
// 	t.Run("must set initial context and props", func(ts *testing.T) {
// 		n := createComponent(context.Background(), &Hello{}, make(prop.Props))
// 		core := n.core()
// 		if core.context == nil {
// 			ts.Error("expected context to be set")
// 		}
// 		if core.props == nil {
// 			ts.Error("expected initial props to be set")
// 		}
// 	})
// 	t.Run("must initialize a new instance without Constructor", func(ts *testing.T) {
// 		n := createComponent(context.Background(), &HelloPlain{}, prop.Props{
// 			"key": "value",
// 		})
// 		core := n.core()
// 		if core.context == nil {
// 			ts.Error("expected context to be set")
// 		}
// 		if core.props == nil {
// 			ts.Error("expected initial props to be set")
// 		}
// 		if core.props.String("key").Value != "value" {
// 			t.Error("expected initial props to be set")
// 		}
// 	})
// }
