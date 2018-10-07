package macmenus

import (
	// typedef void *Item;
	"C"

	"github.com/richardwilkes/ui/event"
)

//export validateMenuItemCallback
// nolint: deadcode
func validateMenuItemCallback(menuItem C.Item) bool {
	if item, ok := itemMap[menuItem]; ok {
		evt := event.NewValidate(item)
		event.Dispatch(evt)
		item.enabled = evt.Valid()
		return item.enabled
	}
	return true
}

//export handleMenuItemCallback
// nolint: deadcode
func handleMenuItemCallback(menuItem C.Item) {
	if item, ok := itemMap[menuItem]; ok {
		event.Dispatch(event.NewSelection(item))
	}
}
