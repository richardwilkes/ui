package macmenus

import (
	"fmt"

	"github.com/richardwilkes/ui/event"
	"github.com/richardwilkes/ui/keys"
	"github.com/richardwilkes/ui/menu"
)

// NewSeparator creates a new separator item.
func NewSeparator() menu.Item {
	item := &platformItem{item: platformNewSeparator()}
	item.InitTypeAndID(item)
	itemMap[item.item] = item
	return item
}

// NewItem creates a new item with no key accelerator.
func NewItem(title string, handler event.Handler) menu.Item {
	return NewItemWithKey(title, 0, handler)
}

// NewItemWithKey creates a new item with a key accelerator using the platform-default modifiers.
func NewItemWithKey(title string, keyCode int, handler event.Handler) menu.Item {
	return NewItemWithKeyAndModifiers(title, keyCode, keys.PlatformMenuModifier(), handler)
}

// NewItemWithKeyAndModifiers creates a new item.
func NewItemWithKeyAndModifiers(title string, keyCode int, modifiers keys.Modifiers, handler event.Handler) menu.Item {
	item := &platformItem{item: platformNewItem(title, keyCode, modifiers), title: title, keyCode: keyCode, keyModifiers: modifiers, enabled: true}
	item.InitTypeAndID(item)
	if handler != nil {
		item.EventHandlers().Add(event.SelectionType, handler)
	}
	itemMap[item.item] = item
	return item
}

func (item *platformItem) String() string {
	return fmt.Sprintf("menu.Item #%d (%s)", item.ID(), item.Title())
}

// EventHandlers returns the handler mappings for this item.
func (item *platformItem) EventHandlers() *event.Handlers {
	if item.eventHandlers == nil {
		item.eventHandlers = &event.Handlers{}
	}
	return item.eventHandlers
}

// ParentTarget returns the parent target of this item, or nil.
func (item *platformItem) ParentTarget() event.Target {
	return event.GlobalTarget()
}

// Title returns this item's title.
func (item *platformItem) Title() string {
	return item.title
}

// KeyCode returns the key code that can be used to trigger this item. A value of 0 indicates no
// key is attached.
func (item *platformItem) KeyCode() int {
	return item.keyCode
}

// KeyModifiers returns the key modifiers that are required to trigger this item.
func (item *platformItem) KeyModifiers() keys.Modifiers {
	return item.keyModifiers
}

// SubMenu returns a sub-menu attached to this item or nil.
func (item *platformItem) SubMenu() menu.Menu {
	if menu, ok := menuMap[item.platformSubMenu()]; ok {
		return menu
	}
	return nil
}

// Enabled returns true if this item is enabled.
func (item *platformItem) Enabled() bool {
	return item.enabled
}

func (item *platformItem) Dispose() {
	if _, ok := itemMap[item.item]; ok {
		if subMenu := item.SubMenu(); subMenu != nil {
			subMenu.Dispose()
		}
		delete(itemMap, item.item)
		item.platformDispose()
	}
}
