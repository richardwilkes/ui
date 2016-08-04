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
	arrowID,
	textID,
	verticalTextID,
	crossHairID,
	closedHandID,
	openHandID,
	pointingHandID,
	resizeLeftID,
	resizeRightID,
	resizeLeftRightID,
	resizeUpID,
	resizeDownID,
	resizeUpDownID,
	disappearingItemID,
	notAllowedID,
	dragLinkID,
	dragCopyID,
	contextMenuID,
	customID
};

uiCursor systemCursor(int id);
uiCursor newCursor(void *img, float hotX, float hotY);
void disposeCursor(uiCursor cursor);

#endif // __RW_GOUI_CURSOR__
