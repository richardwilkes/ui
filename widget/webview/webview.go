package webview

import (
	"fmt"

	"github.com/richardwilkes/toolbox/xmath/geom"
	"github.com/richardwilkes/ui"
	"github.com/richardwilkes/ui/widget"
)

// WebView represents a native web view. Since it is a native component, it
// does not respect the view hierarchy and effectively draws on top of all
// other components.
type WebView struct {
	widget.Block
	owner     ui.Window
	webview   platformWebView
	url       string
	urlLoaded bool
}

// NewWebView creates a new, empty web view.
//
// Note that this component behaves differently than other components because
// it is backed by a native widget. In particular, you must pass in a valid
// Window at construction time and you must manually dispose of it when no
// longer needed.
func NewWebView(wnd ui.Window) *WebView {
	w := &WebView{owner: wnd}
	w.InitTypeAndID(w)
	w.Describer = func() string { return fmt.Sprintf("WebView #%d", w.ID()) }
	w.webview = platformNewWebView(wnd)
	w.SetFocusable(true)
	return w
}

// Dispose of the web view.
func (w *WebView) Dispose() {
	w.RemoveFromParent()
	w.platformDispose()
	w.webview = nil
}

// LoadURL loads the specified URL into the view.
func (w *WebView) LoadURL(url string) {
	w.url = url
	w.urlLoaded = false
	if url != "" && w.Parent() != nil {
		w.platformLoadURL()
		w.urlLoaded = true
	}
}

// SetParent implements the Widget interface.
func (w *WebView) SetParent(parent ui.Widget) {
	w.Block.SetParent(parent)
	w.adjustNativeBounds()
	if !w.urlLoaded && w.url != "" && w.Parent() != nil {
		w.platformLoadURL()
		w.urlLoaded = true
	}
}

// SetBounds implements the Widget interface.
func (w *WebView) SetBounds(bounds geom.Rect) {
	if bounds != w.Bounds() {
		w.Block.SetBounds(bounds)
		w.adjustNativeBounds()
	}
}

// SetLocation implements the Widget interface.
func (w *WebView) SetLocation(pt geom.Point) {
	if pt != w.Location() {
		w.Block.SetLocation(pt)
		w.adjustNativeBounds()
	}
}

// SetSize implements the Widget interface.
func (w *WebView) SetSize(size geom.Size) {
	if size != w.Size() {
		w.Block.SetSize(size)
		w.adjustNativeBounds()
	}
}

func (w *WebView) adjustNativeBounds() {
	b := w.LocalBounds()
	b.Point = w.ToWindow(b.Point)
	w.platformSetViewFrame(b)
}
