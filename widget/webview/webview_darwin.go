package webview

import (
	// #cgo CFLAGS: -x objective-c
	// #cgo LDFLAGS: -framework Cocoa -framework WebKit -framework Security
	// #include "webview_darwin.h"
	"C"
	"unsafe"

	"github.com/richardwilkes/toolbox/xmath/geom"
	"github.com/richardwilkes/ui"
)

type platformWebView unsafe.Pointer

func platformNewWebView(wnd ui.Window) platformWebView {
	return platformWebView(C.newWebView(wnd.PlatformPtr()))
}

func (w *WebView) platformDispose() {
	C.disposeWebView(C.platformWebView(w.webview))
}

func (w *WebView) platformSetViewFrame(bounds geom.Rect) {
	C.setWebViewFrame(C.platformWebView(w.webview), C.double(bounds.X), C.double(bounds.Y), C.double(bounds.Width), C.double(bounds.Height))
}

func (w *WebView) platformLoadURL() {
	url := C.CString(w.url)
	defer C.free(unsafe.Pointer(url))
	C.loadWebViewURL(C.platformWebView(w.webview), url)
}
