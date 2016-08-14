// Copyright (c) 2016 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

#include <stdlib.h>
#include "_cgo_export.h"
#include "Window.h"

platformWindow platformNewWindow(platformRect bounds, int styleMask) {
	// RAW: Implement platformNewWindow for Windows
	return NULL;
}

const char *platformGetWindowTitle(platformWindow window) {
	// RAW: Implement platformGetWindowTitle for Windows
	return NULL;
}

void platformSetWindowTitle(platformWindow window, const char *title) {
	// RAW: Implement platformSetWindowTitle for Windows
}

platformRect platformGetWindowFrame(platformWindow window) {
	// RAW: Implement platformGetWindowFrame for Windows
	platformRect rect;
	return rect;
}

platformPoint platformGetWindowPosition(platformWindow window) {
	// RAW: Implement platformGetWindowPosition for Windows
	platformPoint pt;
	return pt;
}

platformSize platformGetWindowSize(platformWindow window) {
	// RAW: Implement platformGetWindowSize for Windows
	platformSize size;
	return size;
}

platformRect platformGetWindowContentFrame(platformWindow window) {
	// RAW: Implement platformGetWindowContentFrame for Windows
	platformRect rect;
	return rect;
}

platformPoint platformGetWindowContentPosition(platformWindow window) {
	// RAW: Implement platformGetWindowContentPosition for Windows
	platformPoint pt;
	return pt;
}

platformSize platformGetWindowContentSize(platformWindow window) {
	// RAW: Implement platformGetWindowContentSize for Windows
	platformSize size;
	return size;
}

void platformSetWindowPosition(platformWindow window, float x, float y) {
	// RAW: Implement platformSetWindowPosition for Windows
}

void platformSetWindowSize(platformWindow window, float width, float height) {
	// RAW: Implement platformSetWindowSize for Windows
}

void platformSetWindowContentPosition(platformWindow window, float x, float y) {
	// RAW: Implement platformSetWindowContentPosition for Windows
}

void platformSetWindowContentSize(platformWindow window, float width, float height) {
	// RAW: Implement platformSetWindowContentSize for Windows
}

float platformGetWindowScalingFactor(platformWindow window) {
	// RAW: Implement platformGetWindowScalingFactor for Windows
	return 1;
}

void platformMinimizeWindow(platformWindow window) {
	// RAW: Implement platformMinimizeWindow for Windows
}

void platformZoomWindow(platformWindow window) {
	// RAW: Implement platformZoomWindow for Windows
}

void platformBringWindowToFront(platformWindow window) {
	// RAW: Implement platformBringWindowToFront for Windows
}

void platformBringAllWindowsToFront() {
	// RAW: Implement platformBringAllWindowsToFront for Windows
}

platformWindow platformGetKeyWindow() {
	// RAW: Implement platformGetKeyWindow for Windows
	return NULL;
}

void platformRepaintWindow(platformWindow window, platformRect bounds) {
	// RAW: Implement platformRepaintWindow for Windows
}

void platformFlushPainting(platformWindow window) {
	// RAW: Implement platformFlushPainting for Windows
}

void platformSetToolTip(platformWindow window, const char *tooltip) {
	// RAW: Implement platformSetToolTip for Windows
}

void platformSetCursor(platformWindow window, void *cursor) {
	// RAW: Implement platformSetCursor for Windows
}

void platformHideCursorUntilMouseMoves() {
	// RAW: Implement platformHideCursorUntilMouseMoves for Windows
}
