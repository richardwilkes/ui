// Copyright (c) 2016 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

#ifndef __RW_GOUI_WINDOW__
#define __RW_GOUI_WINDOW__

typedef void *platformWindow;

typedef struct {
	float x, y;
} platformPoint;

typedef struct {
	float width, height;
} platformSize;

typedef struct {
	float x, y, width, height;
} platformRect;

enum {
	platformMouseDown = 0,
	platformMouseDragged,
	platformMouseUp,
	platformMouseEntered,
	platformMouseMoved,
	platformMouseExited,
	platformMouseWheel,
	platformKeyDown,
	platformKeyTyped,
	platformKeyUp
};

enum {
	platformCapsLockKeyMask	= 1 << 0,
	platformShiftKeyMask	= 1 << 1,
	platformControlKeyMask	= 1 << 2,
	platformOptionKeyMask	= 1 << 3,
	platformCommandKeyMask	= 1 << 4
};

platformWindow platformNewWindow(platformRect bounds, int styleMask);
void platformCloseWindow(platformWindow window);
platformRect platformGetWindowFrame(platformWindow window);
platformPoint platformGetWindowPosition(platformWindow window);
platformSize platformGetWindowSize(platformWindow window);
platformRect platformGetWindowContentFrame(platformWindow window);
platformPoint platformGetWindowContentPosition(platformWindow window);
platformSize platformGetWindowContentSize(platformWindow window);
const char *platformGetWindowTitle(platformWindow window);
void platformSetWindowTitle(platformWindow window, const char *title);
void platformSetWindowPosition(platformWindow window, float x, float y);
void platformSetWindowSize(platformWindow window, float width, float height);
void platformSetWindowContentPosition(platformWindow window, float x, float y);
void platformSetWindowContentSize(platformWindow window, float width, float height);
float platformGetWindowScalingFactor(platformWindow window);
void platformMinimizeWindow(platformWindow window);
void platformZoomWindow(platformWindow window);
void platformBringWindowToFront(platformWindow window);
void platformBringAllWindowsToFront();
platformWindow platformGetKeyWindow();
void platformRepaintWindow(platformWindow window, platformRect bounds);
void platformFlushPainting(platformWindow window);
void platformSetToolTip(platformWindow window, const char *tooltip);
void platformSetCursor(platformWindow window, void *cursor);
void platformHideCursorUntilMouseMoves();

#endif // __RW_GOUI_WINDOW__
