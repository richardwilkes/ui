// Copyright (c) 2016 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

#include <Cocoa/Cocoa.h>
#include "_cgo_export.h"
#include "Menu.h"

@interface menuDelegate : NSObject
@end

@implementation menuDelegate
- (BOOL)validateMenuItem:(NSMenuItem *)menuItem {
    return validateMenuItem((uiMenuItem)menuItem);
}

- (void)handleMenuItem:(id)sender {
    handleMenuItem((uiMenuItem)sender);
}
@end

uiMenu getMainMenu() {
    return (uiMenu)[NSApp mainMenu];
}

void setMainMenu(uiMenu menuBar) {
    [NSApp setMainMenu:(NSMenu *)menuBar];
}

uiMenu uiNewMenu(const char *title) {
    NSMenu *menu = [[[NSMenu alloc] initWithTitle:[NSString stringWithUTF8String:title]] retain];
    return (uiMenu)menu;
}

void uiDisposeMenu(uiMenu menu) {
    [((NSMenu *)menu) release];
}

int uiMenuItemCount(uiMenu menu) {
    return [((NSMenu *)menu) numberOfItems];
}

uiMenuItem uiGetMenuItem(uiMenu menu, int index) {
    if (index < 0 || index >= uiMenuItemCount(menu)) {
        return nil;
    }
    return (uiMenuItem)[((NSMenu *)menu) itemAtIndex:index];
}

uiMenuItem uiAddMenuItem(uiMenu menu, const char *title, const char *key) {
    NSMenuItem *item = [((NSMenu *)menu) addItemWithTitle:[NSString stringWithUTF8String:title] action:@selector(handleMenuItem:) keyEquivalent:[NSString stringWithUTF8String:key]];
    [item setTarget:[menuDelegate new]];
    return (uiMenuItem)item;
}

uiMenuItem uiAddSeparator(uiMenu menu) {
    NSMenuItem *item = [NSMenuItem separatorItem];
    [((NSMenu *)menu) addItem:item];
    return item;
}

void uiSetKeyModifierMask(uiMenuItem item, int mask) {
	// macOS uses the same modifier mask bit order as we do, but it is shifted up by 16 bits
    [((NSMenuItem *)item) setKeyEquivalentModifierMask:mask << 16];
}

uiMenu uiGetSubMenu(uiMenuItem item) {
    NSMenuItem *mitem = (NSMenuItem *)item;
    if ([mitem hasSubmenu]) {
        return (uiMenu)[mitem submenu];
    }
    return nil;
}

void uiSetSubMenu(uiMenuItem item, uiMenu subMenu) {
    [((NSMenuItem *)item) setSubmenu: subMenu];
}

void uiSetServicesMenu(uiMenu menu) {
	[NSApp setServicesMenu:(NSMenu *)menu];
}

void uiSetWindowMenu(uiMenu menu) {
	[NSApp setWindowsMenu:(NSMenu *)menu];
}

void uiSetHelpMenu(uiMenu menu) {
	[NSApp setHelpMenu:(NSMenu *)menu];
}

void uiPopupMenu(uiWindow window, uiMenu menu, float x, float y, uiMenuItem itemAtLocation) {
	[((NSMenu *)menu) popUpMenuPositioningItem:itemAtLocation atLocation:NSMakePoint(x,y) inView:[((NSWindow *)window) contentView]];
}
