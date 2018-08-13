package main

import (
	"context"
	"syscall/js"

	"github.com/gernest/vected/vdom/dom"
)

var doc = js.Global().Get("document")

func main() {
	ctx, _ := context.WithCancel(context.TODO())
	root := doc.Call("getElementById", "main")
	btn := dom.CreateNode("button")
	txt := doc.Call("createTextNode", "Click me")
	btn.Call("appendChild", txt)
	cb := dom.NewEvent(0, func(args js.Value) {
		e := args.Get("target").Get("_listeners")
		js.Global().Get("console").Call("log", e)
		// fmt.Println(v)
	})
	dom.SetAccessor(btn, "onClick", nil, cb, false)
	root.Call("appendChild", btn)
	<-ctx.Done()
}
