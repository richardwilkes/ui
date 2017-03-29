// Copyright (c) 2016 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

package macmenus

import (
	"strings"
	"unsafe"

	"github.com/richardwilkes/geom"
	"github.com/richardwilkes/ui"
	"github.com/richardwilkes/ui/event"
	"github.com/richardwilkes/ui/keys"
	"github.com/richardwilkes/ui/menu"
	"github.com/richardwilkes/ui/object"
	// #cgo CFLAGS: -x objective-c
	// #cgo LDFLAGS: -framework Cocoa
	// #include <Cocoa/Cocoa.h>
	//
	// typedef void *Menu;
	// typedef void *Item;
	//
	// BOOL validateMenuItemCallback(Item menuItem);
	// void handleMenuItemCallback(Item menuItem);
	//
	// @interface ItemDelegate : NSObject
	// @end
	//
	// @implementation ItemDelegate
	// - (BOOL)validateMenuItem:(NSMenuItem *)menuItem { return validateMenuItemCallback(menuItem); }
	// - (void)handleMenuItem:(id)sender { handleMenuItemCallback(sender); }
	// @end
	//
	// static ItemDelegate *itemDelegate = nil;
	//
	// Item newItem(const char *title, const char *key, int modifiers) {
	//	NSMenuItem *item = [[[NSMenuItem alloc] initWithTitle:[NSString stringWithUTF8String:title] action:@selector(handleMenuItem:) keyEquivalent:[NSString stringWithUTF8String:key]] retain];
	//	// macOS uses the same modifier mask bit order as we do, but it is shifted up by 16 bits
	//    [item setKeyEquivalentModifierMask:modifiers << 16];
	//    if (!itemDelegate) {
	//    	itemDelegate = [ItemDelegate new];
	//    }
	//    [item setTarget:itemDelegate];
	//    return item;
	// }
	//
	// Menu subMenu(Item item) {
	//    NSMenuItem *mitem = (NSMenuItem *)item;
	//    if ([mitem hasSubmenu]) {
	//        return [mitem submenu];
	//    }
	//    return nil;
	// }
	//
	// void setBar(Menu bar) { [NSApp setMainMenu:bar]; }
	// Menu newMenu(const char *title) { return [[[NSMenu alloc] initWithTitle:[NSString stringWithUTF8String:title]] retain]; }
	// Item newSeparator() { return [[NSMenuItem separatorItem] retain]; }
	// void disposeItem(Item item) { [((NSMenuItem *)item) release]; }
	// void disposeMenu(Menu menu) { [((NSMenu *)menu) release]; }
	// void setSubMenu(Item item, Menu subMenu) { [((NSMenuItem *)item) setSubmenu: subMenu]; }
	// int itemCount(Menu menu) { return [((NSMenu *)menu) numberOfItems]; }
	// Item item(Menu menu, int index) { return (index < 0 || index >= itemCount(menu)) ?  nil : [((NSMenu *)menu) itemAtIndex:index]; }
	// void insertItem(Menu menu, Item item, int index) { [((NSMenu *)menu) insertItem:item atIndex:index]; }
	// void removeItem(Menu menu, int index) { [((NSMenu *)menu) removeItemAtIndex:index]; }
	// void popup(void *window, Menu menu, double x, double y, Item itemAtLocation) { [((NSMenu *)menu) popUpMenuPositioningItem:itemAtLocation atLocation:NSMakePoint(x,y) inView:[((NSWindow *)window) contentView]]; }
	// void setServicesMenu(Menu menu) { [NSApp setServicesMenu:menu]; }
	// void setWindowMenu(Menu menu) { [NSApp setWindowsMenu:menu]; }
	// void setHelpMenu(Menu menu) { [NSApp setHelpMenu:menu]; }
	"C"
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

func SetServicesMenu(menu menu.Menu) {
	C.setServicesMenu(menu.(*platformMenu).menu)
}

func SetWindowMenu(menu menu.Menu) {
	C.setWindowMenu(menu.(*platformMenu).menu)
}

func SetHelpMenu(menu menu.Menu) {
	C.setHelpMenu(menu.(*platformMenu).menu)
}
