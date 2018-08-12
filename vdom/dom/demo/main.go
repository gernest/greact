package main

import (
	"context"
	"fmt"
	"syscall/js"
)

var doc = js.Global().Get("document")

func main() {
	ctx, _ := context.WithCancel(context.TODO())
	root := doc.Call("getElementById", "main")
	btn := doc.Call("createElement", "button")
	txt := doc.Call("createTextNode", "Click me")
	btn.Call("appendChild", txt)
	cb := js.NewCallback(func(args []js.Value) {
		fmt.Println("button clicked")
	})
	btn.Call("addEventListener", "click", cb)
	root.Call("appendChild", btn)
	<-ctx.Done()
}
