package dom

import "github.com/gopherjs/gopherjs/js"

// DataTransfer object is used to hold the data that is being dragged during a
// drag and drop operation. It may hold one or more data items, each of one or
// more data types. For more information about drag and drop, see HTML Drag and
// Drop API.
//
// This object is available from the dataTransfer property of all drag events.
// It cannot be created separately (i.e. there is no constructor for this
// object).
type DataTransfer struct {
	*js.Object

	// The DataTransfer.dropEffect property controls the feedback (typically visual)
	// the user is given during a drag and drop operation. It will affect which
	// cursor is displayed while dragging. For example, when the user hovers over a
	// target drop element, the browser's cursor may indicate which type of
	// operation will occur.
	//
	// When the DataTransfer object is created, dropEffect is set to a string value.
	// On getting, it returns its current value. On setting, if the new value is one
	// of the values listed below, then the property's current value will be set to
	// the new value and other values will be ignored.
	//
	// For the dragenter and dragover events, dropEffect will be initialized based
	// on what action the user is requesting. How this is determined is platform
	// specific, but typically the user can press modifier keys such as the alt key
	// to adjust the desired action. Within event handlers for dragenter and
	// dragover events, dropEffect should be modified if a different action is
	// desired than the action that the user is requesting.
	//
	// For the drop and dragend events, dropEffect will be set to the action that
	// was desired, which will be the value dropEffect had after the last dragenter
	// or dragover event. In a dragend event, for instance, if the desired
	// dropEffect is "move", then the data being dragged should be removed from the
	// source.
	//
	// The meaning of values
	// copy
	// 	A copy of the source item is made at the new location.
	// move
	// 	An item is moved to a new location.
	// link
	// 	A link is established to the source at the new location.
	// none
	// 	The item may not be dropped.
	DropEffect string `js:"dropEffect"`

	// The DataTransfer.effectAllowed property specifies the effect that is allowed
	// for a drag operation. The copy operation is used to indicate that the data
	// being dragged will be copied from its present location to the drop location.
	// The move operation is used to indicate that the data being dragged will be
	// moved, and the link operation is used to indicate that some form of
	// relationship or connection will be created between the source and drop
	// locations.
	//
	// This property should be set in the dragstart event to set the desired drag
	// effect for the drag source. Within the dragenter and dragover event
	// handlers, this property will be set to whatever value was assigned during
	// the dragstart event, thus effectAllowed may be used to determine which
	// effect is permitted.
	//
	// Assigning a value to effectAllowed in events other than dragstart has no
	// effect.
	//
	// values
	// 	none
	// 	The item may not be dropped.
	// copy
	// 	A copy of the source item may be made at the new location.
	// copyLink
	// 	A copy or link operation is permitted.
	// copyMove
	// 	A copy or move operation is permitted.
	// link
	// 	A link may be established to the source at the new location.
	// linkMove
	// 	A link or move operation is permitted.
	// move
	// 	An item may be moved to a new location.
	// all
	// 	All operations are permitted.
	// uninitialized
	// 	The default value when the effect has not been set, equivalent to all.
	EffectAllowed string `js:"effectAllowed"`

	// Contains a list of all the local files available on the data transfer. If
	// the drag operation doesn't involve dragging files, this property is an empty
	// list.
	Files *FileList `js:"files"`

	// is a list of all of the drag data.
	Items *DataTransferItemList `js:"items"`

	Types []string `js:"types"`
}

func (dt *DataTransfer) SetData(format string, data string) {
	dt.Call("setData", format, data)
}

func (dt *DataTransfer) GetData(format string) string {
	return dt.Call("getData", format).String()
}

type File struct {
	*js.Object

	// Returns the last modified time of the file, in millisecond since the UNIX
	// epoch (January 1st, 1970 at Midnight).
	LastModified int64 `js:"lastModified"`

	// Returns the last modified Date of the file referenced by the File object.
	LastModifiedDate *Date `js:"lastModifiedDate"`

	// Returns the name of the file referenced by the File object.
	Name string `js:"name"`

	// Returns the path the URL of the File is relative to.
	WebkitRelativePath string `js:"webkitRelativePath "`

	// Returns the MIME type of the file.
	Type string `js:"type "`
}

// FileList returned by the files property of the HTML <input>
// element; this lets you access the list of files selected with the <input
// type="file"> element. It's also used for a list of files dropped into web
// content when using the drag and drop API; see the DataTransfer object for
// details on this usage.
type FileList struct {
	*js.Object
}

type Date struct {
	*js.Object
}

type DataTransferItemList struct {
	*js.Object

	// is the number of drag items in the list.
	Length uint64 `js:"length"`
}
