# Integration tests

These are tests aimed to render a component/component aon the real dom and
assert some things on it.

This was specifically designed to work with the vecty package so integration
tests works only with vecty by design.

Example integration test 

```go
func TestRenderBody() mad.Integration {
	txt := "hello,world"
	return mad.RenderBody("mad.RenderBody",
		func() interface{} {
			return elem.Body(
				vecty.Text(txt),
			)
		},
		mad.It("must have text node", func(t mad.T) {
			defer func() {
				if err := recover(); err != nil {
					t.Error(err)
				}
			}()
			o := js.Global.Get("document").Get("body").Get("textContent").String()
			if o != txt {
				t.Errorf("expected %s got %s", txt, o)
			}
		}),
	)
}
```