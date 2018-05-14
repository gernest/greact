package ws

import (
	"github.com/gopherjs/gopherjs/js"
)

func New() {
	h := js.Global.Get("location").Get("href")
	println(h)
}
