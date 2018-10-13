package filemenu

import (
	"github.com/richardwilkes/toolbox/i18n"
	"github.com/richardwilkes/ui/event"
	"github.com/richardwilkes/ui/keys"
	"github.com/richardwilkes/ui/menu"
	"github.com/richardwilkes/ui/window"
)

// NewCloseKeyWindowItem creates the standard "Close" menu item that will
// close the current key window when chosen.
func NewCloseKeyWindowItem() menu.Item {
	item := menu.NewItemWithKey(i18n.Text("Close"), keys.VirtualKeyW, CloseKeyWindow)
	item.EventHandlers().Add(event.ValidateType, ValidateCloseKeyWindow)
	return item
}

// CloseKeyWindow attempts to close the current key window.
func CloseKeyWindow(_ event.Event) {
	wnd := window.KeyWindow()
	if wnd != nil && wnd.Closable() {
		wnd.AttemptClose()
	}
}

// ValidateCloseKeyWindow marks the menu item invalid if the current key
// window is not closable.
func ValidateCloseKeyWindow(evt event.Event) {
	wnd := window.KeyWindow()
	if wnd == nil || !wnd.Closable() {
		evt.(*event.Validate).MarkInvalid()
	}
}
