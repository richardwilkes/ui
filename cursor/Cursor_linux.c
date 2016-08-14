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
#include "Cursor.h"

uiCursor platformSystemCursor(int id) {
	// RAW: Implement platformSystemCursor for Linux
	switch (id) {
		case platformArrowID:
		case platformTextID:
		case platformVerticalTextID:
		case platformCrossHairID:
		case platformClosedHandID:
		case platformOpenHandID:
		case platformPointingHandID:
		case platformResizeLeftID:
		case platformResizeRightID:
		case platformResizeLeftRightID:
		case platformResizeUpID:
		case platformResizeDownID:
		case platformResizeUpDownID:
		case platformDisappearingItemID:
		case platformNotAllowedID:
		case platformDragLinkID:
		case platformDragCopyID:
		case platformContextMenuID:
		default:
			return NULL;
	}
}

uiCursor platformNewCursor(void *img, float hotX, float hotY) {
	// RAW: Implement platformNewCursor for Linux
	return NULL;
}

void platformDisposeCursor(uiCursor cursor) {
	// RAW: Implement platformDisposeCursor for Linux
}
