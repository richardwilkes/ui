package editmenu

import (
	"github.com/richardwilkes/toolbox/i18n"
	"github.com/richardwilkes/ui/event"
	"github.com/richardwilkes/ui/keys"
	"github.com/richardwilkes/ui/menu"
	"github.com/richardwilkes/ui/window"
)

// Deletable defines the methods required of objects that can respond to the Delete menu item.
type Deletable interface {
	// CanDelete returns true if Delete() can be called successfully.
	CanDelete() bool
	// Delete the data from the object.
	Delete()
}

// AppendDeleteItem appends the standard Delete menu item to the specified menu.
func AppendDeleteItem(m menu.Menu) {
	InsertDeleteItem(m, -1)
}

// InsertDeleteItem adds the standard Delete menu item to the specified menu.
func InsertDeleteItem(m menu.Menu, index int) {
	item := menu.NewItemWithKeyAndModifiers(i18n.Text("Delete"), keys.VirtualKeyBackspace, 0, Delete)
	item.EventHandlers().Add(event.ValidateType, CanDelete)
	m.InsertItem(item, index)
}

// Delete the data from the current keyboard focus.
func Delete(evt event.Event) {
	wnd := window.KeyWindow()
	if wnd != nil {
		focus := wnd.Focus()
		if c, ok := focus.(Deletable); ok {
			c.Delete()
		}
	}
}

// CanDelete returns true if Delete() can be called successfully.
func CanDelete(evt event.Event) {
	valid := false
	wnd := window.KeyWindow()
	if wnd != nil {
		focus := wnd.Focus()
		if c, ok := focus.(Deletable); ok {
			valid = c.CanDelete()
		}
	}
	if !valid {
		evt.(*event.Validate).MarkInvalid()
	}
}
