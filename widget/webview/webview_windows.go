package webview

import (
	"unsafe"

	"github.com/richardwilkes/toolbox/xmath/geom"
	"github.com/richardwilkes/ui"
)

type platformWebView unsafe.Pointer

func platformNewWebView(wnd ui.Window) platformWebView {
	panic("unimplemented")
}

func (w *WebView) platformDispose() {
	panic("unimplemented")
}

func (w *WebView) platformSetViewFrame(bounds geom.Rect) {
	panic("unimplemented")
}

func (w *WebView) platformLoadURL() {
	panic("unimplemented")
}
