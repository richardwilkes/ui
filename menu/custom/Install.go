package custom

import (
	"github.com/richardwilkes/ui/event"
	"github.com/richardwilkes/ui/keys"
	"github.com/richardwilkes/ui/menu"
)

// Install the standard menu bar.
func Install() {
	menu.AppBar = func(id uint64) menu.Bar { return AppBar(id) }
	menu.Global = func() bool { return false }
	menu.NewMenu = func(title string) menu.Menu { return NewMenu(title) }
	menu.NewItem = func(title string, handler event.Handler) menu.Item { return NewItem(title, handler) }
	menu.NewItemWithKey = func(title string, keyCode int, handler event.Handler) menu.Item {
		return NewItemWithKey(title, keyCode, handler)
	}
	menu.NewItemWithKeyAndModifiers = func(title string, keyCode int, modifiers keys.Modifiers, handler event.Handler) menu.Item {
		return NewItemWithKeyAndModifiers(title, keyCode, modifiers, handler)
	}
	menu.NewSeparator = func() menu.Item { return NewSeparator() }
}
