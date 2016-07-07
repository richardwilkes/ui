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
#include "Window.h"

@interface drawingView : NSView
@end

@interface windowDelegate : NSObject<NSWindowDelegate>
@end

uiWindow uiNewWindow(uiRect bounds, int styleMask) {
	NSRect contentRect = NSMakeRect(0, 0, bounds.width, bounds.height);
	NSWindow *window = [[NSWindow alloc] initWithContentRect:contentRect styleMask:styleMask backing:NSBackingStoreBuffered defer:YES];
	[window setFrameTopLeftPoint:NSMakePoint(bounds.x, [[NSScreen mainScreen] visibleFrame].size.height - bounds.y)];
	drawingView *rootView = [drawingView new];
	[window setContentView:rootView];
	[window setDelegate: [windowDelegate new]];
	[rootView addTrackingArea:[[NSTrackingArea alloc] initWithRect:contentRect options:NSTrackingMouseEnteredAndExited | NSTrackingMouseMoved | NSTrackingActiveInKeyWindow | NSTrackingInVisibleRect owner:rootView userInfo:nil]];
	return (uiWindow)window;
}

const char *uiGetWindowTitle(uiWindow window) {
	return strdup([[((NSWindow *)window) title] UTF8String]);
}

void uiSetWindowTitle(uiWindow window, const char *title) {
	[((NSWindow *)window) setTitle:[NSString stringWithUTF8String:title]];
}

uiRect uiGetWindowFrame(uiWindow window) {
	uiRect rect;
	CGRect frame = [((NSWindow *)window) frame];
	rect.x = frame.origin.x;
	rect.y = [[NSScreen mainScreen] visibleFrame].size.height - (frame.origin.y + frame.size.height);
	rect.width = frame.size.width;
	rect.height = frame.size.height;
	return rect;
}

uiPoint uiGetWindowPosition(uiWindow window) {
	uiPoint pt;
	CGRect frame = [((NSWindow *)window) frame];
	pt.x = frame.origin.x;
	pt.y = [[NSScreen mainScreen] visibleFrame].size.height - (frame.origin.y + frame.size.height);
	return pt;
}

uiSize uiGetWindowSize(uiWindow window) {
	CGSize cgSize = [((NSWindow *)window) frame].size;
	uiSize size;
	size.width = cgSize.width;
	size.height = cgSize.height;
	return size;
}

uiRect uiGetWindowContentFrame(uiWindow window) {
	NSWindow *win = (NSWindow *)window;
	CGRect frame = [[win contentView] frame];
	frame.origin = [win frame].origin;
	CGRect windowFrame = [win frameRectForContentRect:frame];
	frame.origin.x += frame.origin.x - windowFrame.origin.x;
	frame.origin.y += frame.origin.y - windowFrame.origin.y;
	uiRect rect;
	rect.x = frame.origin.x;
	rect.y = [[NSScreen mainScreen] visibleFrame].size.height - (frame.origin.y + frame.size.height);
	rect.width = frame.size.width;
	rect.height = frame.size.height;
	return rect;
}

uiPoint uiGetWindowContentPosition(uiWindow window) {
	NSWindow *win = (NSWindow *)window;
	CGRect frame = [[win contentView] frame];
	frame.origin = [win frame].origin;
	CGRect windowFrame = [win frameRectForContentRect:frame];
	frame.origin.x += frame.origin.x - windowFrame.origin.x;
	frame.origin.y += frame.origin.y - windowFrame.origin.y;
	uiPoint pt;
	pt.x = frame.origin.x;
	pt.y = [[NSScreen mainScreen] visibleFrame].size.height - (frame.origin.y + frame.size.height);
	return pt;
}

uiSize uiGetWindowContentSize(uiWindow window) {
	CGSize cgSize = [[((NSWindow *)window) contentView] frame].size;
	uiSize size;
	size.width = cgSize.width;
	size.height = cgSize.height;
	return size;
}

void uiSetWindowPosition(uiWindow window, float x, float y) {
	NSWindow *win = (NSWindow *)window;
	[win setFrameOrigin:NSMakePoint(x, [[NSScreen mainScreen] visibleFrame].size.height - (y + [win frame].size.height))];
}

void uiSetWindowSize(uiWindow window, float width, float height) {
	NSWindow *win = (NSWindow *)window;
	CGRect frame = [win frame];
	[win setFrame:NSMakeRect(frame.origin.x, frame.origin.y + (frame.size.height - height), width, height) display:YES];
}

void uiSetWindowContentPosition(uiWindow window, float x, float y) {
	uiPoint pos = uiGetWindowContentPosition(window);
	uiPoint outerPos = uiGetWindowPosition(window);
	uiSetWindowPosition(window, x + outerPos.x - pos.x, y + outerPos.y - pos.y);
}

void uiSetWindowContentSize(uiWindow window, float width, float height) {
	uiPoint origin = uiGetWindowPosition(window);
	[((NSWindow *)window) setContentSize:NSMakeSize(width, height)];
	uiSetWindowPosition(window, origin.x, origin.y);
}

float uiGetWindowScalingFactor(uiWindow window) {
	NSView *view = [((NSWindow *)window) contentView];
	CGRect bounds = [view bounds];
	CGFloat width = bounds.size.width;
	if (width <= 0) {
		return [((NSWindow *)window) backingScaleFactor];
	}
    return [view convertRectToBacking:bounds].size.width / width;
}

void uiMinimizeWindow(uiWindow window) {
	[((NSWindow *)window) performMiniaturize:nil];
}

void uiZoomWindow(uiWindow window) {
	[((NSWindow *)window) performZoom:nil];
}

void uiBringWindowToFront(uiWindow window) {
	[((NSWindow *)window) makeKeyAndOrderFront:nil];
}

void uiBringAllWindowsToFront() {
	[[NSRunningApplication currentApplication] activateWithOptions:NSApplicationActivateAllWindows | NSApplicationActivateIgnoringOtherApps];
}

uiWindow uiGetKeyWindow() {
	return (uiWindow)[NSApp keyWindow];
}

void uiRepaintWindow(uiWindow window, uiRect bounds) {
	[[((NSWindow *)window) contentView] setNeedsDisplayInRect:NSMakeRect(bounds.x, bounds.y, bounds.width, bounds.height)];
}

void uiSetToolTip(uiWindow window, const char *tooltip) {
	NSView *view = [((NSWindow *)window) contentView];
	// We always clear the old one out first. Failure to do so results in new tooltips not always showing up.
	[view setToolTip:nil];
	if (tooltip) {
		[view setToolTip:[NSString stringWithUTF8String:tooltip]];
	}
}

@implementation drawingView

-(BOOL)isFlipped {
	return YES;
}

-(void)viewDidEndLiveResize {
	[self setNeedsDisplayInRect:[self bounds]];
}

-(void)drawRect:(NSRect)dirtyRect {
	uiRect bounds;
	bounds.x = dirtyRect.origin.x;
	bounds.y = dirtyRect.origin.y;
	bounds.width = dirtyRect.size.width;
	bounds.height = dirtyRect.size.height;
	drawWindow((uiWindow)[self window], [[NSGraphicsContext currentContext] graphicsPort], bounds, [self inLiveResize]);
}

-(void)deliverMouseEvent:(NSEvent *)theEvent ofType:(unsigned char)type {
	unsigned char clickCount = 0;
	if (type != uiMouseEntered && type != uiMouseExited) {
		clickCount = theEvent.clickCount;
	}
	NSPoint where = [self convertPoint:theEvent.locationInWindow fromView:nil];
	// macOS uses the same modifier mask bit order as we do, but it is shifted up by 16 bits
	handleMouseEvent((uiWindow)[self window], type, (theEvent.modifierFlags & (NSAlphaShiftKeyMask | NSShiftKeyMask | NSControlKeyMask | NSAlternateKeyMask | NSCommandKeyMask)) >> 16, theEvent.buttonNumber, clickCount, where.x, where.y);
}

-(void)mouseDown:(NSEvent *)theEvent {
	[self deliverMouseEvent:theEvent ofType:uiMouseDown];
}

-(void)rightMouseDown:(NSEvent *)theEvent {
	[self deliverMouseEvent:theEvent ofType:uiMouseDown];
}

-(void)otherMouseDown:(NSEvent *)theEvent {
	[self deliverMouseEvent:theEvent ofType:uiMouseDown];
}

- (void)mouseDragged:(NSEvent *)theEvent {
	[self deliverMouseEvent:theEvent ofType:uiMouseDragged];
}

- (void)rightMouseDragged:(NSEvent *)theEvent {
	[self deliverMouseEvent:theEvent ofType:uiMouseDragged];
}

- (void)otherMouseDragged:(NSEvent *)theEvent {
	[self deliverMouseEvent:theEvent ofType:uiMouseDragged];
}

-(void)mouseUp:(NSEvent *)theEvent {
	[self deliverMouseEvent:theEvent ofType:uiMouseUp];
}

-(void)rightMouseUp:(NSEvent *)theEvent {
	[self deliverMouseEvent:theEvent ofType:uiMouseUp];
}

-(void)otherMouseUp:(NSEvent *)theEvent {
	[self deliverMouseEvent:theEvent ofType:uiMouseUp];
}

- (void)mouseMoved:(NSEvent *)theEvent {
	[self deliverMouseEvent:theEvent ofType:uiMouseMoved];
}

- (void)mouseEntered:(NSEvent *)theEvent {
	[self deliverMouseEvent:theEvent ofType:uiMouseEntered];
}

- (void)mouseExited:(NSEvent *)theEvent {
	[self deliverMouseEvent:theEvent ofType:uiMouseExited];
}

@end

@implementation windowDelegate

- (void)windowDidResize:(NSNotification *)notification {
	windowResized((uiWindow)[notification object]);
}

- (BOOL)windowShouldClose:(id)sender {
	return (BOOL)shouldClose((uiWindow)sender);
}

- (void)windowWillClose:(NSNotification *)notification {
	didClose((uiWindow)[notification object]);
}

@end
