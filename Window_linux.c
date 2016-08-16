// Copyright (c) 2016 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

#include <X11/Xlib.h>
#include <stdlib.h>
#include <stdio.h>
#include <string.h>
#include "globals_linux.h"
#include "_cgo_export.h"
#include "Window.h"

platformWindow platformNewWindow(platformRect bounds, int styleMask) {
	int screen = DefaultScreen(AppGlobals.display);
	Window window = XCreateSimpleWindow(AppGlobals.display, RootWindow(AppGlobals.display, screen), bounds.x, bounds.y, bounds.width, bounds.height, 1, BlackPixel(AppGlobals.display, screen), WhitePixel(AppGlobals.display, screen));
	XSelectInput(AppGlobals.display, window, KeyPressMask | KeyReleaseMask | ButtonPressMask | ButtonReleaseMask | EnterWindowMask | LeaveWindowMask | ExposureMask | PointerMotionMask | ExposureMask | VisibilityChangeMask | StructureNotifyMask | FocusChangeMask);
	XSetWMProtocols(AppGlobals.display, window, &AppGlobals.wmDeleteMessage, 1);
	XMapWindow(AppGlobals.display, window);
	// Move it back to the original location, as the window manager might have set it somewhere else
	XMoveWindow(AppGlobals.display, window, bounds.x, bounds.y);
	AppGlobals.windowCount++;
	return (platformWindow)window;
}

void platformCloseWindow(platformWindow window) {
	AppGlobals.windowCount--;
	XDestroyWindow(AppGlobals.display, (Window)window);
}

const char *platformGetWindowTitle(platformWindow window) {
	char *result;
	XFetchName(AppGlobals.display, (Window)window, &result);
	if (result) {
		char *name = strdup(result);
		XFree(result);
		result = name;
	}
	return result;
}

void platformSetWindowTitle(platformWindow window, const char *title) {
	XStoreName(AppGlobals.display, (Window)window, title);
}

platformRect platformGetWindowFrame(platformWindow window) {
	Window root;
	int x, y;
	unsigned int width, height, border, depth;
	XGetGeometry(AppGlobals.display, (Window)window, &root, &x, &y, &width, &height, &border, &depth);
	platformRect rect;
	rect.x = x;
	rect.y = y;
	rect.width = width;
	rect.height = height;
	return rect;
}

platformPoint platformGetWindowPosition(platformWindow window) {
	Window root;
	int x, y;
	unsigned int width, height, border, depth;
	XGetGeometry(AppGlobals.display, (Window)window, &root, &x, &y, &width, &height, &border, &depth);
	platformPoint pt;
	pt.x = x;
	pt.y = y;
	return pt;
}

platformSize platformGetWindowSize(platformWindow window) {
	Window root;
	int x, y;
	unsigned int width, height, border, depth;
	XGetGeometry(AppGlobals.display, (Window)window, &root, &x, &y, &width, &height, &border, &depth);
	platformSize size;
	size.width = width;
	size.height = height;
	return size;
}

platformRect platformGetWindowContentFrame(platformWindow window) {
	Window root;
	int x, y;
	unsigned int width, height, border, depth;
	XGetGeometry(AppGlobals.display, (Window)window, &root, &x, &y, &width, &height, &border, &depth);
	platformRect rect;
	rect.x = x + border;
	rect.y = y + border;
	rect.width = width - border * 2;
	rect.height = height - border * 2;
	return rect;
}

platformPoint platformGetWindowContentPosition(platformWindow window) {
	Window root;
	int x, y;
	unsigned int width, height, border, depth;
	XGetGeometry(AppGlobals.display, (Window)window, &root, &x, &y, &width, &height, &border, &depth);
	platformPoint pt;
	pt.x = x + border;
	pt.y = y + border;
	return pt;
}

platformSize platformGetWindowContentSize(platformWindow window) {
	Window root;
	int x, y;
	unsigned int width, height, border, depth;
	XGetGeometry(AppGlobals.display, (Window)window, &root, &x, &y, &width, &height, &border, &depth);
	platformSize size;
	size.width = width - border * 2;
	size.height = height - border * 2;
	return size;
}

void platformSetWindowPosition(platformWindow window, float x, float y) {
	XMoveWindow(AppGlobals.display, (Window)window, x, y);
}

void platformSetWindowSize(platformWindow window, float width, float height) {
	XResizeWindow(AppGlobals.display, (Window)window, width, height);
}

void platformSetWindowContentPosition(platformWindow window, float x, float y) {
	XMoveWindow(AppGlobals.display, (Window)window, x, y);
}

void platformSetWindowContentSize(platformWindow window, float width, float height) {
	XResizeWindow(AppGlobals.display, (Window)window, width, height);
}

float platformGetWindowScalingFactor(platformWindow window) {
	// RAW: Implement platformGetWindowScalingFactor for Linux
	return 1;
}

void platformMinimizeWindow(platformWindow window) {
	XIconifyWindow(AppGlobals.display, (Window)window, DefaultScreen(AppGlobals.display));
}

void platformZoomWindow(platformWindow window) {
	// RAW: Implement platformZoomWindow for Linux
}

void platformBringWindowToFront(platformWindow window) {
	XRaiseWindow(AppGlobals.display, (Window)window);
}

void platformBringAllWindowsToFront() {
	// RAW: Implement platformBringAllWindowsToFront for Linux
}

platformWindow platformGetKeyWindow() {
	Window focus;
	int revert;
	XGetInputFocus(AppGlobals.display, &focus, &revert);
	return (platformWindow)focus;
}

void platformRepaintWindow(platformWindow window, platformRect bounds) {
	XExposeEvent event;
	memset(&event, 0, sizeof(event));
	event.type = Expose;
	event.window = (Window)window;
	event.x = bounds.x;
	event.y = bounds.y;
	event.width = bounds.width;
	event.height = bounds.height;
	XSendEvent(AppGlobals.display, (Window)window, 0, ExposureMask, (XEvent *)&event);
}

void platformFlushPainting(platformWindow window) {
	XFlush(AppGlobals.display);
}

void platformSetToolTip(platformWindow window, const char *tooltip) {
	// RAW: Implement platformSetToolTip for Linux
}

void platformSetCursor(platformWindow window, void *cursor) {
	// RAW: Implement platformSetCursor for Linux
}

void platformHideCursorUntilMouseMoves() {
	// RAW: Implement platformHideCursorUntilMouseMoves for Linux
}
