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

#include <CoreGraphics/CoreGraphics.h>

typedef void *uiWindow;
typedef void *uiGraphicsContext;

typedef struct {
	float x, y;
} uiPoint;

typedef struct {
	float width, height;
} uiSize;

typedef struct {
	float x, y, width, height;
} uiRect;

enum {
	uiMouseDown = 0,
	uiMouseDragged,
	uiMouseUp,
	uiMouseEntered,
	uiMouseMoved,
	uiMouseExited,
	uiMouseWheel,
	uiKeyDown,
	uiKeyTyped,
	uiKeyUp
};

enum {
	uiCapsLockKeyMask	= 1 << 0,
	uiShiftKeyMask		= 1 << 1,
	uiControlKeyMask	= 1 << 2,
	uiOptionKeyMask		= 1 << 3,
	uiCommandKeyMask	= 1 << 4
};

uiWindow uiNewWindow(uiRect bounds, int styleMask);
uiRect uiGetWindowFrame(uiWindow window);
uiPoint uiGetWindowPosition(uiWindow window);
uiSize uiGetWindowSize(uiWindow window);
uiRect uiGetWindowContentFrame(uiWindow window);
uiPoint uiGetWindowContentPosition(uiWindow window);
uiSize uiGetWindowContentSize(uiWindow window);
const char *uiGetWindowTitle(uiWindow window);
void uiSetWindowTitle(uiWindow window, const char *title);
void uiSetWindowPosition(uiWindow window, float x, float y);
void uiSetWindowSize(uiWindow window, float width, float height);
void uiSetWindowContentPosition(uiWindow window, float x, float y);
void uiSetWindowContentSize(uiWindow window, float width, float height);
float uiGetWindowScalingFactor(uiWindow window);
void uiMinimizeWindow(uiWindow window);
void uiZoomWindow(uiWindow window);
void uiBringWindowToFront(uiWindow window);
void uiBringAllWindowsToFront();
uiWindow uiGetKeyWindow();
void uiRepaintWindow(uiWindow window, uiRect bounds);
void uiFlushPainting(uiWindow window);
void uiSetToolTip(uiWindow window, const char *tooltip);

#endif // __RW_GOUI_WINDOW__
