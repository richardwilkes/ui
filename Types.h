// Copyright (c) 2016 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

#ifndef __RW_GOUI_TYPES__
#define __RW_GOUI_TYPES__

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

enum {
	platformBorderlessWindowMask	= 0,
	platformTitledWindowMask		= 1 << 0,
	platformClosableWindowMask		= 1 << 1,
	platformMinimizableWindowMask	= 1 << 2,
	platformResizableWindowMask		= 1 << 3
};

typedef void *platformWindow;
typedef void *platformSurface;
typedef void *platformMenu;
typedef void *platformMenuItem;

typedef struct {
	float x;
	float y;
	float width;
	float height;
} platformRect;

#endif // __RW_GOUI_TYPES__
