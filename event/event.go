package event

import (
	"github.com/gernest/vected/dom"
	"github.com/gopherjs/gopherjs/js"
)

type Type string

const (
	AnimationEvent    Type = "AnimationEvent"
	BeforeUnload      Type = "BeforeUnloadEvent"
	Composition       Type = "CompositionEvent"
	DeviceOrientation Type = "DeviceOrientationEvent"
	Drag              Type = "DragEvent"
	EventBase         Type = "Event"
	CustomEvent       Type = "CustomEvent"
	Focus             Type = "FocusEvent"
	HashChange        Type = "HashChangeEvent"
	Keyboard          Type = "KeyboardEvent"
	Message           Type = "MessageEvent"
	Storage           Type = "StorageEvent"
	Touch             Type = "TouchEvent"
	UI                Type = "UIEvent"
)

// Listener implements EventListener interface
type Listener struct {
	*js.Object
	HandleEvent func(*dom.Event) `js:"handleEvent"`
}

func Create(doc *js.Object, typ string, bubbles, cancelable bool, composed ...bool) *dom.Event {
	args := []interface{}{typ, bubbles, cancelable}
	if len(composed) > 0 {
		args = append(args, composed[0])
	}
	o := js.Global.Get(string(EventBase)).New(args...)
	e := &dom.Event{Object: o}
	return e
}

func CreateCustom(doc *js.Object, typ string, detail ...string) *dom.CustomEvent {
	args := []interface{}{typ}
	if len(detail) > 0 {
		args = append(args, map[string]string{
			"detail": detail[0],
		})
	}
	o := js.Global.Get(string(CustomEvent)).New(args...)
	c := &dom.CustomEvent{Object: o}
	c.Event = &dom.Event{Object: o}
	return c
}
