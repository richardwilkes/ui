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
#include "Cursor.h"

uiCursor systemCursor(int id) {
	switch (id) {
		case arrowID:
			return [NSCursor arrowCursor];
		case textID:
			return [NSCursor IBeamCursor];
		case verticalTextID:
			return [NSCursor IBeamCursorForVerticalLayout];
		case crossHairID:
			return [NSCursor crosshairCursor];
		case closedHandID:
			return [NSCursor closedHandCursor];
		case openHandID:
			return [NSCursor openHandCursor];
		case pointingHandID:
			return [NSCursor pointingHandCursor];
		case resizeLeftID:
			return [NSCursor resizeLeftCursor];
		case resizeRightID:
			return [NSCursor resizeRightCursor];
		case resizeLeftRightID:
			return [NSCursor resizeLeftRightCursor];
		case resizeUpID:
			return [NSCursor resizeUpCursor];
		case resizeDownID:
			return [NSCursor resizeDownCursor];
		case resizeUpDownID:
			return [NSCursor resizeUpDownCursor];
		case disappearingItemID:
			return [NSCursor disappearingItemCursor];
		case notAllowedID:
			return [NSCursor operationNotAllowedCursor];
		case dragLinkID:
			return [NSCursor dragLinkCursor];
		case dragCopyID:
			return [NSCursor dragCopyCursor];
		case contextMenuID:
			return [NSCursor contextualMenuCursor];
		default:
			return NULL;
	}
}

uiCursor newCursor(void *img, float hotX, float hotY) {
	return [[[NSCursor alloc] initWithImage:img hotSpot:NSMakePoint(hotX,hotY)] retain];
}

void disposeCursor(uiCursor cursor) {
	[((NSCursor *)cursor) release];
}
