package editmenu

import (
	"github.com/richardwilkes/toolbox/i18n"
	"github.com/richardwilkes/ui/event"
	"github.com/richardwilkes/ui/keys"
	"github.com/richardwilkes/ui/menu"
	"github.com/richardwilkes/ui/window"
)

// Copyable defines the methods required of objects that can respond to the Copy menu item.
type Copyable interface {
	// CanCopy returns true if Copy() can be called successfully.
	CanCopy() bool
	// Copy the data to the clipboard.
	Copy()
}

// AppendCopyItem appends the standard Copy menu item to the specified menu.
func AppendCopyItem(m menu.Menu) {
	InsertCopyItem(m, -1)
}

// InsertCopyItem adds the standard Copy menu item to the specified menu.
func InsertCopyItem(m menu.Menu, index int) {
	item := menu.NewItemWithKey(i18n.Text("Copy"), keys.VirtualKeyC, Copy)
	item.EventHandlers().Add(event.ValidateType, CanCopy)
	m.InsertItem(item, index)
}

// Copy the data from the current keyboard focus to the clipboard.
func Copy(evt event.Event) {
	wnd := window.KeyWindow()
	if wnd != nil {
		focus := wnd.Focus()
		if c, ok := focus.(Copyable); ok {
			c.Copy()
		}
	}
}

// CanCopy returns true if Copy() can be called successfully.
func CanCopy(evt event.Event) {
	valid := false
	wnd := window.KeyWindow()
	if wnd != nil {
		focus := wnd.Focus()
		if c, ok := focus.(Copyable); ok {
			valid = c.CanCopy()
		}
	}
	if !valid {
		evt.(*event.Validate).MarkInvalid()
	}
}
