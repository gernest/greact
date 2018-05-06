package dom

import "github.com/gopherjs/gopherjs/js"

// EventPhase are values which describe which phase the event flow is currently
// being evaluated.
type EventPhase int

const (
	// None is set when no event is being processed at this time
	None EventPhase = iota

	// Capturing is set when the event is being propagated through the target's
	// ancestor objects. This process starts with the Window, then Document, then
	// the HTMLHtmlElement, and so on through the elements until the target's
	// parent is reached. Event listeners registered for capture mode when
	// EventTarget.addEventListener() was called are triggered during this phase.
	Capturing

	// AtTarget is set when the event has arrived at the event's target. Event
	// listeners registered for this phase are called at this time. If
	// Event.bubbles is false, processing the event is finished after this phase is
	// complete.
	AtTarget

	// Bubbling is set when the event is propagating back up through the target's
	// ancestors in reverse order, starting with the parent, and eventually
	// reaching the containing Window. This is known as bubbling, and occurs only
	// if Event.bubbles is true. Event listeners registered for this phase are
	// triggered during this process.
	Bubbling
)

type Event struct {
	*js.Object

	// A Boolean indicating whether the event bubbles up through the DOM or not.
	Bubbles bool `js:"bubbles"`

	// A historical alias to Event.stopPropagation(). Setting its value to true
	// before returning from an event handler prevents propagation of the event.
	CancelBubbles bool `js:"cancelBubbles"`

	// A Boolean indicating whether the event is cancelable
	Cancelable bool `js:"cancelable"`

	// A Boolean value indicating whether or not the event can bubble across the
	// boundary between the shadow DOM and the regular DOM
	Composed bool `js:"composed"`

	// A reference to the currently registered target for the event. This is the
	// object to which the event is currently slated to be sent; it's possible this
	// has been changed along the way through retargeting.
	CurrentTarget *js.Object `js:"currentTarget"`

	// Indicates whether or not event.preventDefault() has been called on the event.
	DefaultPrevented bool `js:"defaultPrevented "`

	// Indicates which phase of the event flow is being processed.
	Phase EventPhase `js:"eventPhase"`

	// A reference to the target to which the event was originally dispatched.
	Target *js.Object `js:"target"`

	// The time at which the event was created (in milliseconds). By specification,
	// this value is time since epoch, but in reality browsers' definitions vary;
	// in addition, work is underway to change this to be a DOMHighResTimeStamp
	// instead.
	Timestamp int64 `js:"timestamp"`

	// The name of the event (case-insensitive).
	Type string `js:"type"`

	// Indicates whether or not the event was initiated by the browser (after a
	// user click for instance) or by a script (using an event creation method,
	// like event.initEvent).
	Trusted bool `js:"isTrusted"`
}

type CustomEvent struct {
	*js.Object
	*Event
	Detail string `js:"detail"`
}

// AnimationEvent  represents events providing information related
// to animations.
type AnimationEvent struct {
	*js.Object
	*Event

	// Is a string containing the value of the animation-name CSS property
	// associated with the transition.
	Name string `js:"animationName"`

	// Is a float giving the amount of time the animation has been running, in
	// seconds, when this event fired, excluding any time the animation was paused.
	// For an "animationstart" event, elapsedTime is 0.0 unless there was a
	// negative value for animation-delay, in which case the event will be fired
	// with elapsedTime containing (-1 * delay).
	Elapsed float64 `js:"elapsedTime"`

	// Is a string, starting with '::', containing the name of the
	// pseudo-element the animation runs on. If the animation doesn't run on a
	// pseudo-element but on the element, an empty string: ''.
	PseudoElement float64 `js:"pseudoElement"`
}

// BlobEvent  represents events associated with a Blob. These blobs are
// typically, but not necessarily, associated with media content.
type BlobEvent struct {
	*js.Object
	*Event

	// A Blob representing the data associated with the event. The event was fired
	// on the EventTarget because of something happening on that specific Blob.
	Data []byte `js:"data"`

	// A DOMHighResTimeStamp indicating the difference between the timestamp of the
	// first chunk in data and the timestamp of the first chunk in the first
	// BlobEvent produced by this recorder. Note that the timecode in the first
	// produced BlobEvent does not need to be zero.
	Timecode int64 `js:"timecode"`
}

type Clipboard struct {
	*js.Object
}

type DragEvent struct {
	*js.Object

	DataTransfer *DataTransfer `js:"dataTransFer"`
}

func ToDragEvent(ev *js.Object) *DragEvent {
	dt := ev.Get("dataTransfer")
	e := &DragEvent{Object: ev}
	e.DataTransfer = &DataTransfer{Object: dt}
	return e
}
