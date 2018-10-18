package macmenus

import (
	// #cgo CFLAGS: -x objective-c
	// #cgo LDFLAGS: -framework Cocoa
	// #include "menus_darwin.h"
	"C"
	"strings"
	"unsafe"

	"github.com/richardwilkes/toolbox/xmath/geom"
	"github.com/richardwilkes/ui"
	"github.com/richardwilkes/ui/event"
	"github.com/richardwilkes/ui/keys"
	"github.com/richardwilkes/ui/menu"
	"github.com/richardwilkes/ui/object"
)

type platformMenu struct {
	menu  C.Menu // Must be first element in struct!
	title string
}

type platformItem struct {
	item C.Item // Must be first element in struct!
	object.Base
	eventHandlers *event.Handlers
	title         string
	keyCode       int
	keyModifiers  keys.Modifiers
	enabled       bool
}

var (
	menuMap = make(map[C.Menu]*platformMenu)
	itemMap = make(map[C.Item]*platformItem)
)

func platformSetBar(bar C.Menu) {
	C.setBar(bar)
}

func platformNewMenu(title string) C.Menu {
	cTitle := C.CString(title)
	defer C.free(unsafe.Pointer(cTitle))
	return C.newMenu(cTitle)
}

func platformNewSeparator() C.Item {
	return C.newSeparator()
}

func platformNewItem(title string, keyCode int, modifiers keys.Modifiers) C.Item {
	var keyCodeStr string
	if keyCode != 0 {
		mapping := keys.MappingForKeyCode(keyCode)
		if mapping.KeyChar != 0 {
			keyCodeStr = strings.ToLower(string(mapping.KeyChar))
		}
	}
	cTitle := C.CString(title)
	cKey := C.CString(keyCodeStr)
	defer C.free(unsafe.Pointer(cTitle))
	defer C.free(unsafe.Pointer(cKey))
	return C.newItem(cTitle, cKey, C.int(modifiers))
}

func (menu *platformMenu) platformDispose() {
	C.disposeMenu(menu.menu)
}

func (menu *platformMenu) platformItemCount() int {
	return int(C.itemCount(menu.menu))
}

func (menu *platformMenu) platformItem(index int) C.Item {
	return C.item(menu.menu, C.int(index))
}

func (menu *platformMenu) platformInsertItem(item C.Item, index int) {
	C.insertItem(menu.menu, item, C.int(index))
}

func (menu *platformMenu) platformRemove(index int) {
	C.removeItem(menu.menu, C.int(index))
}

func (menu *platformMenu) platformPopup(window ui.Window, where geom.Point, item C.Item) {
	C.popup(window.PlatformPtr(), menu.menu, C.double(where.X), C.double(where.Y), item)
}

func (item *platformItem) platformDispose() {
	C.disposeItem(item.item)
}

func (item *platformItem) platformSubMenu() C.Menu {
	return C.subMenu(item.item)
}

func (item *platformItem) platformSetSubMenu(subMenu C.Menu) {
	C.setSubMenu(item.item, subMenu)
}

// SetServicesMenu designates which menu is the services menu.
func SetServicesMenu(menu menu.Menu) {
	C.setServicesMenu(menu.(*platformMenu).menu)
}

// SetWindowMenu designates which menu is the window menu.
func SetWindowMenu(menu menu.Menu) {
	C.setWindowMenu(menu.(*platformMenu).menu)
}

// SetHelpMenu designates which menu is the help menu.
func SetHelpMenu(menu menu.Menu) {
	C.setHelpMenu(menu.(*platformMenu).menu)
}

//export validateMenuItemCallback
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
func handleMenuItemCallback(menuItem C.Item) {
	if item, ok := itemMap[menuItem]; ok {
		event.Dispatch(event.NewSelection(item))
	}
}
