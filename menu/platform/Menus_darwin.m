// Copyright (c) 2016 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

#include <Cocoa/Cocoa.h>
#include "Menus_darwin.h"
#include "_cgo_export.h"

@interface ItemDelegate : NSObject
@end

@implementation ItemDelegate
- (BOOL)validateMenuItem:(NSMenuItem *)menuItem {
    return platformValidateMenuItem(menuItem);
}

- (void)handleMenuItem:(id)sender {
	platformHandleMenuItem(sender);
}
@end

static ItemDelegate *itemDelegate = nil;

Menu platformBar() {
    return [NSApp mainMenu];
}

void platformSetBar(Menu bar) {
    [NSApp setMainMenu:(NSMenu *)bar];
}

Menu platformNewMenu(const char *title) {
    NSMenu *menu = [[[NSMenu alloc] initWithTitle:[NSString stringWithUTF8String:title]] retain];
    return menu;
}

Item platformNewSeparator() {
	return [[NSMenuItem separatorItem] retain];
}

Item platformNewItem(const char *title, const char *key, int modifiers) {
	NSMenuItem *item = [[[NSMenuItem alloc] initWithTitle:[NSString stringWithUTF8String:title] action:@selector(handleMenuItem:) keyEquivalent:[NSString stringWithUTF8String:key]] retain];
	// macOS uses the same modifier mask bit order as we do, but it is shifted up by 16 bits
    [item setKeyEquivalentModifierMask:modifiers << 16];
    if (!itemDelegate) {
    	itemDelegate = [ItemDelegate new];
    }
    [item setTarget:itemDelegate];
    return item;
}

void platformDisposeItem(Item item) {
    [((NSMenuItem *)item) release];
}

void platformDisposeMenu(Menu menu) {
    [((NSMenu *)menu) release];
}

Menu platformSubMenu(Item item) {
    NSMenuItem *mitem = (NSMenuItem *)item;
    if ([mitem hasSubmenu]) {
        return [mitem submenu];
    }
    return nil;
}

void platformSetSubMenu(Item item, Menu subMenu) {
	[((NSMenuItem *)item) setSubmenu: subMenu];
}

int platformItemCount(Menu menu) {
    return [((NSMenu *)menu) numberOfItems];
}

Item platformItem(Menu menu, int index) {
    if (index < 0 || index >= platformItemCount(menu)) {
        return nil;
    }
    return [((NSMenu *)menu) itemAtIndex:index];
}

void platformAddItem(Menu menu, Item item) {
	[((NSMenu *)menu) addItem:item];
}

void platformInsertItem(Menu menu, Item item, int index) {
	[((NSMenu *)menu) insertItem:item atIndex:index];
}

void platformRemove(Menu menu, int index) {
	[((NSMenu *)menu) removeItemAtIndex:index];
}

void platformSetServicesMenu(Menu menu) {
	[NSApp setServicesMenu:menu];
}

void platformSetWindowMenu(Menu menu) {
	[NSApp setWindowsMenu:menu];
}

void platformSetHelpMenu(Menu menu) {
	[NSApp setHelpMenu:menu];
}
