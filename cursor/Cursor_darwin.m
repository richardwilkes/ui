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
#include "Cursor_darwin.h"

void *platformArrow() {
	return [NSCursor arrowCursor];
}

void *platformText() {
	return [NSCursor IBeamCursor];
}

void *platformVerticalText() {
	return [NSCursor IBeamCursorForVerticalLayout];
}

void *platformCrossHair() {
	return [NSCursor crosshairCursor];
}

void *platformClosedHand() {
	return [NSCursor closedHandCursor];
}

void *platformOpenHand() {
	return [NSCursor openHandCursor];
}

void *platformPointingHand() {
	return [NSCursor pointingHandCursor];
}

void *platformResizeLeft() {
	return [NSCursor resizeLeftCursor];
}

void *platformResizeRight() {
	return [NSCursor resizeRightCursor];
}

void *platformResizeLeftRight() {
	return [NSCursor resizeLeftRightCursor];
}

void *platformResizeUp() {
	return [NSCursor resizeUpCursor];
}

void *platformResizeDown() {
	return [NSCursor resizeDownCursor];
}

void *platformResizeUpDown() {
	return [NSCursor resizeUpDownCursor];
}

void *platformDisappearingItem() {
	return [NSCursor disappearingItemCursor];
}

void *platformNotAllowed() {
	return [NSCursor operationNotAllowedCursor];
}

void *platformDragLink() {
	return [NSCursor dragLinkCursor];
}

void *platformDragCopy() {
	return [NSCursor dragCopyCursor];
}

void *platformContextMenu() {
	return [NSCursor contextualMenuCursor];
}

void *platformNewCursor(void *img, float hotX, float hotY) {
	return [[[NSCursor alloc] initWithImage:img hotSpot:NSMakePoint(hotX,hotY)] retain];
}

void platformDisposeCursor(void *cursor) {
	[((NSCursor *)cursor) release];
}
