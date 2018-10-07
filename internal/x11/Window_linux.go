package x11

import (
	"math"
	"os"
	"unsafe"

	"github.com/richardwilkes/toolbox/xmath/geom"
	"github.com/richardwilkes/ui/draw"

	// #cgo pkg-config: x11 cairo
	// #include <stdlib.h>
	// #include <X11/Xlib.h>
	// #include <X11/Xatom.h>
	// #include <X11/Xutil.h>
	// #include <cairo/cairo-xlib.h>
	"C"
)

const (
	//_NET_WM_STATE_REMOVE = iota
	//_NET_WM_STATE_ADD
	_NET_WM_STATE_TOGGLE = 2
)

const (
	PropModeReplace = iota
	PropModePrepend
	PropModeAppend
)

const (
	CWBackPixmap = 1 << iota
	CWBackPixel
	CWBorderPixmap
	CWBorderPixel
	CWBitGravity
	CWWinGravity
	CWBackingStore
	CWBackingPlanes
	CWBackingPixel
	CWOverrideRedirect
	CWSaveUnder
	CWEventMask
	CWDontPropagate
	CWColormap
	CWCursor
)

type Window C.Window

func NewWindow(bounds geom.Rect) Window {
	attr, mask := prepareCommonWindowAttributes()
	wnd := createWindow(bounds, mask, attr)
	wnd.applyCommonSetup()
	wnd.ChangeProperty(wmWindowTypeAtom, C.XA_ATOM, 32, PropModeReplace, unsafe.Pointer(&wmWindowTypeNormalAtom), 1)
	wnd.setWindowHints(bounds)
	return wnd
}

func NewPopupWindow(parent Window, bounds geom.Rect) Window {
	attr, mask := prepareCommonWindowAttributes()
	attr.override_redirect = C.True
	wnd := createWindow(bounds, mask|CWOverrideRedirect, attr)
	wnd.applyCommonSetup()
	wnd.ChangeProperty(wmWindowTypeAtom, C.XA_ATOM, 32, PropModeReplace, unsafe.Pointer(&wmWindowTypeDropDownMenuAtom), 1)
	wnd.ChangeProperty(WindowStateAtom, C.XA_ATOM, 32, PropModeReplace, unsafe.Pointer(&wmWindowStateSkipTaskBarAtom), 1)
	C.XSetTransientForHint(display, C.Window(wnd), C.Window(parent))
	wnd.setWindowHints(bounds)
	return wnd
}

func createWindow(bounds geom.Rect, mask int, attr *C.XSetWindowAttributes) Window {
	return Window(C.XCreateWindow(display, C.Window(DefaultRootWindow()), C.int(bounds.X), C.int(bounds.Y), C.uint(bounds.Width), C.uint(bounds.Height), 0, C.CopyFromParent, C.InputOutput, nil, C.ulong(mask), attr))
}

func prepareCommonWindowAttributes() (attr *C.XSetWindowAttributes, mask int) {
	return &C.XSetWindowAttributes{background_pixmap: C.None, backing_store: C.WhenMapped}, CWBackPixmap | CWBackingStore
}

func (wnd Window) setWindowHints(bounds geom.Rect) {
	var sizeHints C.XSizeHints
	sizeHints.x = C.int(bounds.X)
	sizeHints.y = C.int(bounds.Y)
	sizeHints.width = C.int(bounds.Width)
	sizeHints.height = C.int(bounds.Height)
	sizeHints.flags = C.PPosition | C.PSize
	C.XSetWMNormalHints(display, C.Window(wnd), &sizeHints)
}

func (wnd Window) applyCommonSetup() {
	wnd.SelectInput(KeyPressMask | KeyReleaseMask | ButtonPressMask | ButtonReleaseMask | EnterWindowMask | LeaveWindowMask | ExposureMask | PointerMotionMask | ExposureMask | VisibilityChangeMask | StructureNotifyMask | FocusChangeMask)
	wnd.SetProtocols(DeleteWindowSubType)
	pid := os.Getpid()
	wnd.ChangeProperty(wmPidAtom, C.XA_CARDINAL, 32, PropModeReplace, unsafe.Pointer(&pid), 1)
}

func (wnd Window) Destroy() {
	C.XDestroyWindow(display, C.Window(wnd))
}

func (wnd Window) NewSurface(size geom.Size) *draw.Surface {
	return draw.NewSurface(unsafe.Pointer(C.cairo_xlib_surface_create(display, C.Drawable(uintptr(wnd)), (*C.Visual)(DefaultVisual()), C.int(size.Width), C.int(size.Height))), size)
}

func (wnd Window) SelectInput(mask int) {
	C.XSelectInput(display, C.Window(wnd), C.long(mask))
}

func (wnd Window) SetProtocols(protocols ...Atom) {
	C.XSetWMProtocols(display, C.Window(wnd), (*C.Atom)(&protocols[0]), C.int(len(protocols)))
}

func (wnd Window) Title() string {
	var result *C.char
	C.XFetchName(display, C.Window(wnd), &result)
	if result == nil {
		return ""
	}
	defer C.XFree(unsafe.Pointer(result))
	return C.GoString(result)
}

func (wnd Window) SetTitle(title string) {
	cTitle := C.CString(title)
	C.XStoreName(display, C.Window(wnd), cTitle)
	C.free(unsafe.Pointer(cTitle))
}

func (wnd Window) Property(property Atom, requestedType Atom) (actualType Atom, actualFormat, count int, data unsafe.Pointer) {
	var format C.int
	var nitems C.ulong
	var rem C.ulong
	var dp *C.uchar
	if C.XGetWindowProperty(display, C.Window(wnd), C.Atom(property), 0, math.MaxInt32, C.False, C.Atom(requestedType), (*C.Atom)(&actualType), &format, &nitems, &rem, &dp) == C.Success {
		actualFormat = int(format)
		count = int(nitems)
		data = unsafe.Pointer(dp)
	}
	return
}

func (wnd Window) ChangeProperty(property Atom, propertyType Atom, format, mode int, data unsafe.Pointer, elements int) {
	C.XChangeProperty(display, C.Window(wnd), C.Atom(property), C.Atom(propertyType), C.int(format), C.int(mode), (*C.uchar)(data), C.int(elements))
}

func (wnd Window) DeleteProperty(property Atom) {
	C.XDeleteProperty(display, C.Window(wnd), C.Atom(property))
}

func (wnd Window) FrameDecorationSpace() (top, left, bottom, right float64) {
	actualType, actualFormat, count, data := wnd.Property(wmWindowFrameExtentsAtom, C.XA_CARDINAL)
	if data != nil {
		if actualType == C.XA_CARDINAL && actualFormat == 32 && count == 4 {
			fields := (*[4]C.long)(data)
			left = float64(fields[0])
			right = float64(fields[1])
			top = float64(fields[2])
			bottom = float64(fields[3])
		}
		C.XFree(data)
	}
	return
}

func (wnd Window) SetFrame(bounds geom.Rect) {
	C.XMoveResizeWindow(display, C.Window(wnd), C.int(bounds.X), C.int(bounds.Y), C.uint(bounds.Width), C.uint(bounds.Height))
}

func (wnd Window) ContentFrame() geom.Rect {
	var xwa C.XWindowAttributes
	if C.XGetWindowAttributes(display, C.Window(wnd), &xwa) != 0 {
		var x, y C.int
		var child C.Window
		if C.XTranslateCoordinates(display, C.Window(wnd), xwa.root, 0, 0, &x, &y, &child) != 0 {
			return geom.Rect{Point: geom.Point{X: float64(x - xwa.x), Y: float64(y - xwa.y)}, Size: geom.Size{Width: float64(xwa.width), Height: float64(xwa.height)}}
		}
	}
	return geom.Rect{}
}

func (wnd Window) Raise() {
	C.XRaiseWindow(display, C.Window(wnd))
}

func (wnd Window) Show() {
	C.XMapWindow(display, C.Window(wnd))
}

func (wnd Window) Move(where geom.Point) {
	C.XMoveWindow(display, C.Window(wnd), C.int(where.X), C.int(where.Y))
}

func (wnd Window) RequestFocus() {
	C.XSetInputFocus(display, C.Window(wnd), C.RevertToNone, C.CurrentTime)
}

func (wnd Window) Repaint(bounds geom.Rect) {
	wnd.Send(ExposureMask, NewExposeEvent(wnd, bounds))
}

func (wnd Window) Minimize() {
	C.XIconifyWindow(display, C.Window(wnd), C.int(DefaultScreen()))
}

func (wnd Window) Zoom() {
	evt := NewClientMessageEvent(wnd, WindowStateAtom, 32)
	data := (*[5]Atom)(unsafe.Pointer(&evt.data))
	data[0] = _NET_WM_STATE_TOGGLE
	data[1] = WindowStateMaximizedHAtom
	data[2] = WindowStateMaximizedVAtom
	DefaultRootWindow().Send(SubstructureRedirectMask|SubstructureNotifyMask, evt)
}

func (wnd Window) InvokeTask(id uint64) {
	evt := NewClientMessageEvent(wnd, TaskSubType, 32)
	data := (*uint64)(unsafe.Pointer(&evt.data))
	*data = id
	wnd.Send(NoEventMask, evt)
	Flush()
}

func (wnd Window) SetCursor(cursor Cursor) {
	C.XDefineCursor(display, C.Window(wnd), C.Cursor(cursor))
}

func (wnd Window) NextEventOfType(eventType int) *Event {
	var event Event
	if C.XCheckTypedWindowEvent(display, C.Window(wnd), C.int(eventType), (*C.XEvent)(&event)) != 0 {
		return &event
	}
	return nil
}

func (wnd Window) Send(eventMask int, evt Eventable) {
	C.XSendEvent(display, C.Window(wnd), C.False, C.long(eventMask), (*C.XEvent)(evt.ToEvent()))
}
