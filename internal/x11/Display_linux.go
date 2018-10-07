package x11

import (
	// #cgo pkg-config: x11
	// #include <X11/Xlib.h>
	"C"
)

const (
	KeyPressMask = 1 << iota
	KeyReleaseMask
	ButtonPressMask
	ButtonReleaseMask
	EnterWindowMask
	LeaveWindowMask
	PointerMotionMask
	PointerMotionHintMask
	Button1MotionMask
	Button2MotionMask
	Button3MotionMask
	Button4MotionMask
	Button5MotionMask
	ButtonMotionMask
	KeymapStateMask
	ExposureMask
	VisibilityChangeMask
	StructureNotifyMask
	ResizeRedirectMask
	SubstructureNotifyMask
	SubstructureRedirectMask
	FocusChangeMask
	PropertyChangeMask
	ColormapChangeMask
	OwnerGrabButtonMask
	NoEventMask = 0
)

var (
	display       *C.Display
	lastEventTime C.Time
)

type Visual C.Visual

func OpenDisplay() {
	if display != nil {
		panic("Cannot open the X11 display again")
	}
	C.XInitThreads()
	if display = C.XOpenDisplay(nil); display == nil {
		panic("Failed to open the X11 display")
	}
	initAtoms()
	initClipboard()
}

func CloseDisplay() {
	C.XCloseDisplay(display)
	display = nil
}

func Running() bool {
	return display != nil
}

func NextEvent() *Event {
	var event Event
	C.XNextEvent(display, (*C.XEvent)(&event))
	switch event.Type() {
	case KeyPressType, KeyReleaseType:
		lastEventTime = event.ToKeyEvent().When()
	case ButtonPressType, ButtonReleaseType:
		lastEventTime = event.ToButtonEvent().When()
	}
	return &event
}

func DefaultRootWindow() Window {
	return Window(C.XDefaultRootWindow(display))
}

func DefaultScreen() int {
	return int(C.XDefaultScreen(display))
}

func DefaultVisual() *Visual {
	return (*Visual)(C.XDefaultVisual(display, C.int(DefaultScreen())))
}

func InputFocus() Window {
	var focus Window
	var revert C.int
	C.XGetInputFocus(display, (*C.Window)(&focus), &revert)
	return focus
}

func Flush() {
	C.XFlush(display)
}
