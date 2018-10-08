#include "menus_darwin.h"
#include "_cgo_export.h"

@interface ItemDelegate : NSObject
@end

@implementation ItemDelegate

- (BOOL)validateMenuItem:(NSMenuItem *)menuItem {
	return validateMenuItemCallback(menuItem);
}

- (void)handleMenuItem:(id)sender {
	handleMenuItemCallback(sender);
}

@end

static ItemDelegate *itemDelegate = nil;

Item newItem(const char *title, const char *key, int modifiers) {
	NSMenuItem *item = [[[NSMenuItem alloc] initWithTitle:[NSString stringWithUTF8String:title] action:@selector(handleMenuItem:) keyEquivalent:[NSString stringWithUTF8String:key]] retain];
	// macOS uses the same modifier mask bit order as we do, but it is shifted up by 16 bits
	[item setKeyEquivalentModifierMask:modifiers << 16];
	if (!itemDelegate) {
	itemDelegate = [ItemDelegate new];
	}
	[item setTarget:itemDelegate];
	return item;
}

Menu subMenu(Item item) {
	NSMenuItem *mitem = (NSMenuItem *)item;
	if ([mitem hasSubmenu]) {
		return [mitem submenu];
	}
	return nil;
}

void setBar(Menu bar) {
	[NSApp setMainMenu:bar];
}

Menu newMenu(const char *title) {
	return [[[NSMenu alloc] initWithTitle:[NSString stringWithUTF8String:title]] retain];
}

Item newSeparator() {
	return [[NSMenuItem separatorItem] retain];
}

void disposeItem(Item item) {
	[((NSMenuItem *)item) release];
}

void disposeMenu(Menu menu) {
	[((NSMenu *)menu) release];
}

void setSubMenu(Item item, Menu subMenu) {
	[((NSMenuItem *)item) setSubmenu: subMenu];
}

int itemCount(Menu menu) {
	return [((NSMenu *)menu) numberOfItems];
}

Item item(Menu menu, int index) {
	return (index < 0 || index >= itemCount(menu)) ?  nil : [((NSMenu *)menu) itemAtIndex:index];
}

void insertItem(Menu menu, Item item, int index) {
	[((NSMenu *)menu) insertItem:item atIndex:index];
}

void removeItem(Menu menu, int index) {
	[((NSMenu *)menu) removeItemAtIndex:index];
}

void popup(void *window, Menu menu, double x, double y, Item itemAtLocation) {
	[((NSMenu *)menu) popUpMenuPositioningItem:itemAtLocation atLocation:NSMakePoint(x,y) inView:[((NSWindow *)window) contentView]];
}

void setServicesMenu(Menu menu) {
	[NSApp setServicesMenu:menu];
}

void setWindowMenu(Menu menu) {
	[NSApp setWindowsMenu:menu];
}

void setHelpMenu(Menu menu) {
	[NSApp setHelpMenu:menu];
}
