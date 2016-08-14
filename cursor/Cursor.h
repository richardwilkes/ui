// Copyright (c) 2016 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

#ifndef __RW_GOUI_CURSOR__
#define __RW_GOUI_CURSOR__

typedef void *uiCursor;

enum systemIDs {
	platformArrowID,
	platformTextID,
	platformVerticalTextID,
	platformCrossHairID,
	platformClosedHandID,
	platformOpenHandID,
	platformPointingHandID,
	platformResizeLeftID,
	platformResizeRightID,
	platformResizeLeftRightID,
	platformResizeUpID,
	platformResizeDownID,
	platformResizeUpDownID,
	platformDisappearingItemID,
	platformNotAllowedID,
	platformDragLinkID,
	platformDragCopyID,
	platformContextMenuID,
	platformCustomID
};

uiCursor platformSystemCursor(int id);
uiCursor platformNewCursor(void *img, float hotX, float hotY);
void platformDisposeCursor(uiCursor cursor);

#endif // __RW_GOUI_CURSOR__
