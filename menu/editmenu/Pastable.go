package editmenu

import (
	"github.com/richardwilkes/toolbox/i18n"
	"github.com/richardwilkes/ui/event"
	"github.com/richardwilkes/ui/keys"
	"github.com/richardwilkes/ui/menu"
	"github.com/richardwilkes/ui/window"
)

// Pastable defines the methods required of objects that can respond to the Paste menu item.
type Pastable interface {
	// CanPaste returns true if Paste() can be called successfully.
	CanPaste() bool
	// Paste the data from the clipboard into the object.
	Paste()
}

// AppendPasteItem appendss the standard Paste menu item to the specified menu.
func AppendPasteItem(m menu.Menu) {
	InsertPasteItem(m, -1)
}

// InsertPasteItem adds the standard Paste menu item to the specified menu.
func InsertPasteItem(m menu.Menu, index int) {
	item := menu.NewItemWithKey(i18n.Text("Paste"), keys.VirtualKeyV, Paste)
	item.EventHandlers().Add(event.ValidateType, CanPaste)
	m.InsertItem(item, index)
}

// Paste the data from the clipboard into the current keyboard focus.
func Paste(evt event.Event) {
	wnd := window.KeyWindow()
	if wnd != nil {
		focus := wnd.Focus()
		if p, ok := focus.(Pastable); ok {
			p.Paste()
		}
	}
}

// CanPaste returns true if Paste() can be called successfully.
func CanPaste(evt event.Event) {
	valid := false
	wnd := window.KeyWindow()
	if wnd != nil {
		focus := wnd.Focus()
		if p, ok := focus.(Pastable); ok {
			valid = p.CanPaste()
		}
	}
	if !valid {
		evt.(*event.Validate).MarkInvalid()
	}
}
