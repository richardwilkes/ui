package editmenu

import (
	"github.com/richardwilkes/toolbox/i18n"
	"github.com/richardwilkes/ui/event"
	"github.com/richardwilkes/ui/keys"
	"github.com/richardwilkes/ui/menu"
	"github.com/richardwilkes/ui/window"
)

// SelectAllable defines the methods required of objects that can respond to the Select All menu
// item.
type SelectAllable interface {
	// CanSelectAll returns true if SelectAll() can be called successfully.
	CanSelectAll() bool
	// SelectAll expands the selection to encompass the entire available range.
	SelectAll()
}

// AppendSelectAllItem adds the standard Select All menu item to the specified menu.
func AppendSelectAllItem(m menu.Menu) {
	InsertSelectAllItem(m, -1)
}

// InsertSelectAllItem adds the standard Select All menu item to the specified menu.
func InsertSelectAllItem(m menu.Menu, index int) {
	item := menu.NewItemWithKey(i18n.Text("Select All"), keys.VirtualKeyA, SelectAll)
	item.EventHandlers().Add(event.ValidateType, CanSelectAll)
	m.InsertItem(item, index)
}

// SelectAll expands the selection of the current keyboard focus to encompass the entire available
// range.
func SelectAll(evt event.Event) {
	wnd := window.KeyWindow()
	if wnd != nil {
		focus := wnd.Focus()
		if sa, ok := focus.(SelectAllable); ok {
			sa.SelectAll()
		}
	}
}

// CanSelectAll returns true if SelectAll() can be called successfully.
func CanSelectAll(evt event.Event) {
	valid := false
	wnd := window.KeyWindow()
	if wnd != nil {
		focus := wnd.Focus()
		if sa, ok := focus.(SelectAllable); ok {
			valid = sa.CanSelectAll()
		}
	}
	if !valid {
		evt.(*event.Validate).MarkInvalid()
	}
}
