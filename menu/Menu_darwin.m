// Copyright (c) 2016 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

#include <Cocoa/Cocoa.h>
#include "Menu_darwin.h"
#include "_cgo_export.h"

@interface menuDelegate : NSObject
@end

@implementation menuDelegate
- (BOOL)validateMenuItem:(NSMenuItem *)menuItem {
    return validateMenuItem((platformItem)menuItem);
}

- (void)handleMenuItem:(id)sender {
    handleMenuItem((platformItem)sender);
}
@end

platformMenu platformMenuBar() {
    return (platformMenu)[NSApp mainMenu];
}

void platformSetMenuBar(platformMenu menuBar) {
    [NSApp setMainMenu:(NSMenu *)menuBar];
}

platformMenu platformNewMenu(const char *title) {
    NSMenu *menu = [[[NSMenu alloc] initWithTitle:[NSString stringWithUTF8String:title]] retain];
    return (platformMenu)menu;
}

void platformDisposeMenu(platformMenu menu) {
    [((NSMenu *)menu) release];
}

int platformMenuItemCount(platformMenu menu) {
    return [((NSMenu *)menu) numberOfItems];
}

platformItem platformGetMenuItem(platformMenu menu, int index) {
    if (index < 0 || index >= platformMenuItemCount(menu)) {
        return nil;
    }
    return (platformItem)[((NSMenu *)menu) itemAtIndex:index];
}

platformItem platformAddMenuItem(platformMenu menu, const char *title, const char *key) {
    NSMenuItem *item = [((NSMenu *)menu) addItemWithTitle:[NSString stringWithUTF8String:title] action:@selector(handleMenuItem:) keyEquivalent:[NSString stringWithUTF8String:key]];
    [item setTarget:[menuDelegate new]];
    return (platformItem)item;
}

platformItem platformAddSeparator(platformMenu menu) {
    NSMenuItem *item = [NSMenuItem separatorItem];
    [((NSMenu *)menu) addItem:item];
    return item;
}

void platformSetServicesMenu(platformMenu menu) {
	[NSApp setServicesMenu:(NSMenu *)menu];
}

void platformSetWindowMenu(platformMenu menu) {
	[NSApp setWindowsMenu:(NSMenu *)menu];
}

void platformSetHelpMenu(platformMenu menu) {
	[NSApp setHelpMenu:(NSMenu *)menu];
}
