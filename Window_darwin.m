// Copyright (c) 2016 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

#include <Cocoa/Cocoa.h>
#include <Quartz/Quartz.h>
#include "_cgo_export.h"
#include "Window.h"

@interface drawingView : NSView
@end

@interface windowDelegate : NSObject<NSWindowDelegate>
@end

platformWindow platformNewWindow(platformRect bounds, int styleMask) {
	NSRect contentRect = NSMakeRect(0, 0, bounds.width, bounds.height);
	NSWindow *window = [[NSWindow alloc] initWithContentRect:contentRect styleMask:styleMask backing:NSBackingStoreBuffered defer:YES];
	[window setFrameTopLeftPoint:NSMakePoint(bounds.x, [[NSScreen mainScreen] visibleFrame].size.height - bounds.y)];
	[window disableCursorRects];
	drawingView *rootView = [drawingView new];
	[window setContentView:rootView];
	[window setDelegate: [windowDelegate new]];
	[rootView addTrackingArea:[[NSTrackingArea alloc] initWithRect:contentRect options:NSTrackingMouseEnteredAndExited | NSTrackingMouseMoved | NSTrackingActiveInKeyWindow | NSTrackingInVisibleRect | NSTrackingCursorUpdate owner:rootView userInfo:nil]];
	return (platformWindow)window;
}

void platformCloseWindow(platformWindow window) {
	// RAW: Implement platformCloseWindow for Mac
}

const char *platformGetWindowTitle(platformWindow window) {
	return [[((NSWindow *)window) title] UTF8String];
}

void platformSetWindowTitle(platformWindow window, const char *title) {
	[((NSWindow *)window) setTitle:[NSString stringWithUTF8String:title]];
}

platformRect platformGetWindowFrame(platformWindow window) {
	platformRect rect;
	CGRect frame = [((NSWindow *)window) frame];
	rect.x = frame.origin.x;
	rect.y = [[NSScreen mainScreen] visibleFrame].size.height - (frame.origin.y + frame.size.height);
	rect.width = frame.size.width;
	rect.height = frame.size.height;
	return rect;
}

platformPoint platformGetWindowPosition(platformWindow window) {
	platformPoint pt;
	CGRect frame = [((NSWindow *)window) frame];
	pt.x = frame.origin.x;
	pt.y = [[NSScreen mainScreen] visibleFrame].size.height - (frame.origin.y + frame.size.height);
	return pt;
}

platformSize platformGetWindowSize(platformWindow window) {
	CGSize cgSize = [((NSWindow *)window) frame].size;
	platformSize size;
	size.width = cgSize.width;
	size.height = cgSize.height;
	return size;
}

platformRect platformGetWindowContentFrame(platformWindow window) {
	NSWindow *win = (NSWindow *)window;
	CGRect frame = [[win contentView] frame];
	frame.origin = [win frame].origin;
	CGRect windowFrame = [win frameRectForContentRect:frame];
	frame.origin.x += frame.origin.x - windowFrame.origin.x;
	frame.origin.y += frame.origin.y - windowFrame.origin.y;
	platformRect rect;
	rect.x = frame.origin.x;
	rect.y = [[NSScreen mainScreen] visibleFrame].size.height - (frame.origin.y + frame.size.height);
	rect.width = frame.size.width;
	rect.height = frame.size.height;
	return rect;
}

platformPoint platformGetWindowContentPosition(platformWindow window) {
	NSWindow *win = (NSWindow *)window;
	CGRect frame = [[win contentView] frame];
	frame.origin = [win frame].origin;
	CGRect windowFrame = [win frameRectForContentRect:frame];
	frame.origin.x += frame.origin.x - windowFrame.origin.x;
	frame.origin.y += frame.origin.y - windowFrame.origin.y;
	platformPoint pt;
	pt.x = frame.origin.x;
	pt.y = [[NSScreen mainScreen] visibleFrame].size.height - (frame.origin.y + frame.size.height);
	return pt;
}

platformSize platformGetWindowContentSize(platformWindow window) {
	CGSize cgSize = [[((NSWindow *)window) contentView] frame].size;
	platformSize size;
	size.width = cgSize.width;
	size.height = cgSize.height;
	return size;
}

void platformSetWindowPosition(platformWindow window, float x, float y) {
	NSWindow *win = (NSWindow *)window;
	[win setFrameOrigin:NSMakePoint(x, [[NSScreen mainScreen] visibleFrame].size.height - (y + [win frame].size.height))];
}

void platformSetWindowSize(platformWindow window, float width, float height) {
	NSWindow *win = (NSWindow *)window;
	CGRect frame = [win frame];
	[win setFrame:NSMakeRect(frame.origin.x, frame.origin.y + (frame.size.height - height), width, height) display:YES];
}

void platformSetWindowContentPosition(platformWindow window, float x, float y) {
	platformPoint pos = platformGetWindowContentPosition(window);
	platformPoint outerPos = platformGetWindowPosition(window);
	platformSetWindowPosition(window, x + outerPos.x - pos.x, y + outerPos.y - pos.y);
}

void platformSetWindowContentSize(platformWindow window, float width, float height) {
	platformPoint origin = platformGetWindowPosition(window);
	[((NSWindow *)window) setContentSize:NSMakeSize(width, height)];
	platformSetWindowPosition(window, origin.x, origin.y);
}

float platformGetWindowScalingFactor(platformWindow window) {
	NSView *view = [((NSWindow *)window) contentView];
	CGRect bounds = [view bounds];
	CGFloat width = bounds.size.width;
	if (width <= 0) {
		return [((NSWindow *)window) backingScaleFactor];
	}
    return [view convertRectToBacking:bounds].size.width / width;
}

void platformMinimizeWindow(platformWindow window) {
	[((NSWindow *)window) performMiniaturize:nil];
}

void platformZoomWindow(platformWindow window) {
	[((NSWindow *)window) performZoom:nil];
}

void platformBringWindowToFront(platformWindow window) {
	[((NSWindow *)window) makeKeyAndOrderFront:nil];
}

void platformBringAllWindowsToFront() {
	[[NSRunningApplication currentApplication] activateWithOptions:NSApplicationActivateAllWindows | NSApplicationActivateIgnoringOtherApps];
}

platformWindow platformGetKeyWindow() {
	return (platformWindow)[NSApp keyWindow];
}

void platformRepaintWindow(platformWindow window, platformRect bounds) {
	[[((NSWindow *)window) contentView] setNeedsDisplayInRect:NSMakeRect(bounds.x, bounds.y, bounds.width, bounds.height)];
}

void platformFlushPainting(platformWindow window) {
	[CATransaction flush];
}

void platformSetToolTip(platformWindow window, const char *tooltip) {
	NSView *view = [((NSWindow *)window) contentView];
	// We always clear the old one out first. Failure to do so results in new tooltips not always showing up.
	[view setToolTip:nil];
	if (tooltip) {
		[view setToolTip:[NSString stringWithUTF8String:tooltip]];
	}
}

void platformSetCursor(platformWindow window, void *cursor) {
	[((NSCursor *)cursor) set];
}

void platformHideCursorUntilMouseMoves() {
	[NSCursor setHiddenUntilMouseMoves:YES];
}

@implementation drawingView

-(BOOL)isFlipped {
	return YES;
}

-(void)viewDidEndLiveResize {
	[self setNeedsDisplayInRect:[self bounds]];
}

-(void)drawRect:(NSRect)dirtyRect {
	platformRect bounds;
	bounds.x = dirtyRect.origin.x;
	bounds.y = dirtyRect.origin.y;
	bounds.width = dirtyRect.size.width;
	bounds.height = dirtyRect.size.height;
	drawWindow((platformWindow)[self window], [[NSGraphicsContext currentContext] CGContext], bounds, [self inLiveResize]);
}

-(int)getModifiers:(NSEvent *)theEvent {
	// macOS uses the same modifier mask bit order as we do, but it is shifted up by 16 bits
	return (theEvent.modifierFlags & (NSAlphaShiftKeyMask | NSShiftKeyMask | NSControlKeyMask | NSAlternateKeyMask | NSCommandKeyMask)) >> 16;
}

-(void)deliverMouseEvent:(NSEvent *)theEvent ofType:(unsigned char)type {
	unsigned char clickCount = 0;
	if (type != platformMouseEntered && type != platformMouseExited) {
		clickCount = theEvent.clickCount;
	}
	NSPoint where = [self convertPoint:theEvent.locationInWindow fromView:nil];
	handleWindowMouseEvent((platformWindow)[self window], type, [self getModifiers:theEvent], theEvent.buttonNumber, clickCount, where.x, where.y);
}

-(void)mouseDown:(NSEvent *)theEvent {
	[self deliverMouseEvent:theEvent ofType:platformMouseDown];
}

-(void)rightMouseDown:(NSEvent *)theEvent {
	[self deliverMouseEvent:theEvent ofType:platformMouseDown];
}

-(void)otherMouseDown:(NSEvent *)theEvent {
	[self deliverMouseEvent:theEvent ofType:platformMouseDown];
}

-(void)mouseDragged:(NSEvent *)theEvent {
	[self deliverMouseEvent:theEvent ofType:platformMouseDragged];
}

-(void)rightMouseDragged:(NSEvent *)theEvent {
	[self deliverMouseEvent:theEvent ofType:platformMouseDragged];
}

-(void)otherMouseDragged:(NSEvent *)theEvent {
	[self deliverMouseEvent:theEvent ofType:platformMouseDragged];
}

-(void)mouseUp:(NSEvent *)theEvent {
	[self deliverMouseEvent:theEvent ofType:platformMouseUp];
}

-(void)rightMouseUp:(NSEvent *)theEvent {
	[self deliverMouseEvent:theEvent ofType:platformMouseUp];
}

-(void)otherMouseUp:(NSEvent *)theEvent {
	[self deliverMouseEvent:theEvent ofType:platformMouseUp];
}

-(void)mouseMoved:(NSEvent *)theEvent {
	[self deliverMouseEvent:theEvent ofType:platformMouseMoved];
}

-(void)mouseEntered:(NSEvent *)theEvent {
	[self deliverMouseEvent:theEvent ofType:platformMouseEntered];
}

-(void)mouseExited:(NSEvent *)theEvent {
	[self deliverMouseEvent:theEvent ofType:platformMouseExited];
}

-(void)cursorUpdate:(NSEvent *)event {
	handleCursorUpdateEvent((platformWindow)[self window]);
}

-(void)scrollWheel:(NSEvent *)theEvent {
	NSPoint where = [self convertPoint:theEvent.locationInWindow fromView:nil];
	handleWindowMouseWheelEvent((platformWindow)[self window], platformMouseWheel, [self getModifiers:theEvent], where.x, where.y, theEvent.scrollingDeltaX, theEvent.scrollingDeltaY);
}

-(BOOL)acceptsFirstResponder {
	return YES;
}

-(void)deliverKeyEvent:(NSEvent *)theEvent ofType:(unsigned char)type {
	handleWindowKeyEvent((platformWindow)[self window], type, [self getModifiers:theEvent], theEvent.keyCode, (char *)[theEvent.characters UTF8String], theEvent.ARepeat);
}

- (void)keyDown:(NSEvent *)theEvent {
	[self deliverKeyEvent:theEvent ofType:platformKeyDown];
}

- (void)keyUp:(NSEvent *)theEvent {
	[self deliverKeyEvent:theEvent ofType:platformKeyUp];
}

- (void)flagsChanged:(NSEvent *)theEvent {
	int modifiers = [self getModifiers:theEvent];
	unsigned char type;
	switch (theEvent.keyCode) {
		case 57:	// Caps Lock
			type = (modifiers & platformCapsLockKeyMask) == 0 ? platformKeyUp : platformKeyDown;
			break;
		case 56:	// Left Shift
		case 60:	// Right Shift
			type = (modifiers & platformShiftKeyMask) == 0 ? platformKeyUp : platformKeyDown;
			break;
		case 59:	// Left Control
		case 62:	// Right Control
			type = (modifiers & platformControlKeyMask) == 0 ? platformKeyUp : platformKeyDown;
			break;
		case 58:	// Left Option
		case 61:	// Right Option
			type = (modifiers & platformOptionKeyMask) == 0 ? platformKeyUp : platformKeyDown;
			break;
		case 54:	// Right Cmd
		case 55:	// Left Cmd
			type = (modifiers & platformCommandKeyMask) == 0 ? platformKeyUp : platformKeyDown;
			break;
		default:
			type = platformKeyDown;
			break;
	}
	handleWindowKeyEvent((platformWindow)[self window], type, modifiers, theEvent.keyCode, nil, NO);
}

@end

@implementation windowDelegate

- (void)windowDidResize:(NSNotification *)notification {
	windowResized((platformWindow)[notification object]);
}

- (BOOL)windowShouldClose:(id)sender {
	return (BOOL)windowShouldClose((platformWindow)sender);
}

- (void)windowWillClose:(NSNotification *)notification {
	windowDidClose((platformWindow)[notification object]);
}

@end