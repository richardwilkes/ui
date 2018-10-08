#include "window_darwin.h"
#include "_cgo_export.h"

@interface drawingView : NSView
@end

@implementation drawingView

-(void)drawRect:(NSRect)dirtyRect {
	platformWindow wnd = (platformWindow)[self window];
	CGRect rect = [self bounds];
	cairo_surface_t *surface = cairo_quartz_surface_create_for_cg_context([[NSGraphicsContext currentContext] CGContext], (unsigned int)rect.size.width, (unsigned int)rect.size.height);
	cairo_t *gc = cairo_create(surface);
	cairo_surface_destroy(surface); // surface won't actually be destroyed until the gc is destroyed
	drawWindow(wnd, gc, dirtyRect.origin.x, dirtyRect.origin.y, dirtyRect.size.width, dirtyRect.size.height);
	cairo_destroy(gc);
}

-(void)mouseDown:(NSEvent *)theEvent {
	NSPoint where = [self convertPoint:theEvent.locationInWindow fromView:nil];
	handleMouseDownEvent((platformWindow)[self window], where.x, where.y, theEvent.buttonNumber, theEvent.clickCount, [self getModifiers:theEvent]);
}

-(void)mouseDragged:(NSEvent *)theEvent {
	NSPoint where = [self convertPoint:theEvent.locationInWindow fromView:nil];
	handleMouseDraggedEvent((platformWindow)[self window], where.x, where.y, theEvent.buttonNumber, [self getModifiers:theEvent]);
}

-(void)mouseUp:(NSEvent *)theEvent {
	NSPoint where = [self convertPoint:theEvent.locationInWindow fromView:nil];
	handleMouseUpEvent((platformWindow)[self window], where.x, where.y, theEvent.buttonNumber, [self getModifiers:theEvent]);
}

-(void)mouseEntered:(NSEvent *)theEvent {
	NSPoint where = [self convertPoint:theEvent.locationInWindow fromView:nil];
	handleMouseEnteredEvent((platformWindow)[self window], where.x, where.y, [self getModifiers:theEvent]);
}

-(void)mouseMoved:(NSEvent *)theEvent {
	NSPoint where = [self convertPoint:theEvent.locationInWindow fromView:nil];
	handleMouseMovedEvent((platformWindow)[self window], where.x, where.y, [self getModifiers:theEvent]);
}

-(void)mouseExited:(NSEvent *)theEvent {
	NSPoint where = [self convertPoint:theEvent.locationInWindow fromView:nil];
	handleMouseExitedEvent((platformWindow)[self window], where.x, where.y, [self getModifiers:theEvent]);
}

-(void)cursorUpdate:(NSEvent *)theEvent {
	NSPoint where = [self convertPoint:theEvent.locationInWindow fromView:nil];
	handleCursorUpdateEvent((platformWindow)[self window], where.x, where.y, [self getModifiers:theEvent]);
}

-(void)scrollWheel:(NSEvent *)theEvent {
	NSPoint where = [self convertPoint:theEvent.locationInWindow fromView:nil];
	handleWindowMouseWheelEvent((platformWindow)[self window], where.x, where.y, theEvent.scrollingDeltaX, theEvent.scrollingDeltaY, [self getModifiers:theEvent]);
}

-(void)flagsChanged:(NSEvent *)theEvent {
	BOOL down;
	switch (theEvent.keyCode) {
		case 57:	// Caps Lock
			down = (theEvent.modifierFlags & NSEventModifierFlagCapsLock) != 0;
			break;
		case 56:	// Left Shift
		case 60:	// Right Shift
			down = (theEvent.modifierFlags & NSEventModifierFlagShift) != 0;
			break;
		case 59:	// Left Control
		case 62:	// Right Control
			down = (theEvent.modifierFlags & NSEventModifierFlagControl) != 0;
			break;
		case 58:	// Left Option
		case 61:	// Right Option
			down = (theEvent.modifierFlags & NSEventModifierFlagOption) != 0;
			break;
		case 54:	// Right Cmd
		case 55:	// Left Cmd
			down = (theEvent.modifierFlags & NSEventModifierFlagCommand) != 0;
			break;
		default:
			down = true;
			break;
	}
	handleWindowKeyEvent((platformWindow)[self window], theEvent.keyCode, nil, [self getModifiers:theEvent], down, false);
}

-(BOOL)acceptsFirstResponder {
	return YES;
}

-(BOOL)isFlipped {
	return YES;
}

-(void)viewDidEndLiveResize {
	[self setNeedsDisplayInRect:[self bounds]];
}

-(int)getModifiers:(NSEvent *)theEvent {
	// macOS uses the same modifier mask bit order as we do, but it is shifted up by 16 bits
	return (theEvent.modifierFlags & (NSEventModifierFlagCapsLock | NSEventModifierFlagShift | NSEventModifierFlagControl | NSEventModifierFlagOption | NSEventModifierFlagCommand)) >> 16;
}

-(void)rightMouseDown:(NSEvent *)theEvent {
	[self mouseDown:theEvent];
}

-(void)otherMouseDown:(NSEvent *)theEvent {
	[self mouseDown:theEvent];
}

-(void)rightMouseDragged:(NSEvent *)theEvent {
	[self mouseDragged:theEvent];
}

-(void)otherMouseDragged:(NSEvent *)theEvent {
	[self mouseDragged:theEvent];
}

-(void)rightMouseUp:(NSEvent *)theEvent {
	[self mouseUp:theEvent];
}

-(void)otherMouseUp:(NSEvent *)theEvent {
	[self mouseUp:theEvent];
}

-(void)deliverKeyEvent:(NSEvent *)theEvent isDown:(BOOL)down {
	handleWindowKeyEvent((platformWindow)[self window], theEvent.keyCode, (char *)[theEvent.characters UTF8String], [self getModifiers:theEvent], down, theEvent.ARepeat);
}

-(void)keyDown:(NSEvent *)theEvent {
	[self deliverKeyEvent:theEvent isDown:true];
}

-(void)keyUp:(NSEvent *)theEvent {
	[self deliverKeyEvent:theEvent isDown:false];
}

@end

@interface drawingWindow : NSWindow
@end

@implementation drawingWindow

-(BOOL)canBecomeKeyWindow {
	return YES;
}

@end

@interface windowDelegate : NSObject<NSWindowDelegate>
@end

@implementation windowDelegate

-(void)windowDidResize:(NSNotification *)notification {
	windowResized((platformWindow)[notification object]);
}

-(void)windowDidBecomeKey:(NSNotification *)notification {
	windowGainedKey((platformWindow)[notification object]);
}

-(void)windowDidResignKey:(NSNotification *)notification {
	windowLostKey((platformWindow)[notification object]);
}

-(BOOL)windowShouldClose:(id)sender {
	return (BOOL)windowShouldClose((platformWindow)sender);
}

-(void)windowWillClose:(NSNotification *)notification {
	windowDidClose((platformWindow)[notification object]);
}

@end

platformWindow getKeyWindow() {
	return (platformWindow)[NSApp keyWindow];
}

void bringAllWindowsToFront() {
	[[NSRunningApplication currentApplication] activateWithOptions:NSApplicationActivateAllWindows | NSApplicationActivateIgnoringOtherApps];
}

void hideCursorUntilMouseMoves() {
	[NSCursor setHiddenUntilMouseMoves:YES];
}

void closeWindow(platformWindow window) {
	[((NSWindow *)window) close];
}

const char *getWindowTitle(platformWindow window) {
	return [[((NSWindow *)window) title] UTF8String];
}

void setWindowTitle(platformWindow window, const char *title) {
	[((NSWindow *)window) setTitle:[NSString stringWithUTF8String:title]];
}

void bringWindowToFront(platformWindow window) {
	[((NSWindow *)window) makeKeyAndOrderFront:nil];
}

void repaintWindow(platformWindow window, double x, double y, double width, double height) {
	[[((NSWindow *)window) contentView] setNeedsDisplayInRect:NSMakeRect(x, y, width, height)];
}

void flushPainting(platformWindow window) {
	[CATransaction flush];
}

void minimizeWindow(platformWindow window) {
	[((NSWindow *)window) performMiniaturize:nil];
}

void zoomWindow(platformWindow window) {
	[((NSWindow *)window) performZoom:nil];
}

void setCursor(platformWindow window, void *cursor) {
	[((NSCursor *)cursor) set];
}

void invoke(unsigned long id) {
	dispatch_async_f(dispatch_get_main_queue(), (void *)id, (dispatch_function_t)dispatchTask);
}

void invokeAfter(unsigned long id, long afterNanos) {
	dispatch_after_f(dispatch_time(DISPATCH_TIME_NOW, afterNanos), dispatch_get_main_queue(), (void *)id, (dispatch_function_t)dispatchTask);
}

platformWindow newWindow(double x, double y, double width, double height, int styleMask) {
	// The styleMask bits match those that Mac OS uses
	NSRect contentRect = NSMakeRect(0, 0, width, height);
	NSWindow *window = [[drawingWindow alloc] initWithContentRect:contentRect styleMask:styleMask backing:NSBackingStoreBuffered defer:YES];
	[window setFrameTopLeftPoint:NSMakePoint(x, [[NSScreen mainScreen] visibleFrame].size.height - y)];
	[window disableCursorRects];
	drawingView *rootView = [drawingView new];
	[window setContentView:rootView];
	[window setDelegate: [windowDelegate new]];
	[rootView addTrackingArea:[[NSTrackingArea alloc] initWithRect:contentRect options:NSTrackingMouseEnteredAndExited | NSTrackingMouseMoved | NSTrackingActiveInKeyWindow | NSTrackingInVisibleRect | NSTrackingCursorUpdate owner:rootView userInfo:nil]];
	return (platformWindow)window;
}

void getWindowFrame(platformWindow window, double *x, double *y, double *width, double *height) {
	CGRect frame = [((NSWindow *)window) frame];
	*x = frame.origin.x;
	*y = [[NSScreen mainScreen] visibleFrame].size.height - (frame.origin.y + frame.size.height);
	*width = frame.size.width;
	*height = frame.size.height;
}

void setWindowFrame(platformWindow window, double x, double y, double width, double height) {
	NSWindow *win = (NSWindow *)window;
	CGRect frame = [win frame];
	[win setFrame:NSMakeRect(x, [[NSScreen mainScreen] visibleFrame].size.height - (y + height), width, height) display:YES];
}

void getWindowContentFrame(platformWindow window, double *x, double *y, double *width, double *height) {
	NSWindow *win = (NSWindow *)window;
	CGRect frame = [[win contentView] frame];
	frame.origin = [win frame].origin;
	CGRect windowFrame = [win frameRectForContentRect:frame];
	frame.origin.x += frame.origin.x - windowFrame.origin.x;
	frame.origin.y += frame.origin.y - windowFrame.origin.y;
	*x = frame.origin.x;
	*y = [[NSScreen mainScreen] visibleFrame].size.height - (frame.origin.y + frame.size.height);
	*width = frame.size.width;
	*height = frame.size.height;
}
