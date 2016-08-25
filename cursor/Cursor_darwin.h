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

void *platformArrow();
void *platformText();
void *platformVerticalText();
void *platformCrossHair();
void *platformClosedHand();
void *platformOpenHand();
void *platformPointingHand();
void *platformResizeLeft();
void *platformResizeRight();
void *platformResizeLeftRight();
void *platformResizeUp();
void *platformResizeDown();
void *platformResizeUpDown();
void *platformDisappearingItem();
void *platformNotAllowed();
void *platformDragLink();
void *platformDragCopy();
void *platformContextMenu();
void *platformSystemCursor(int id);
void *platformNewCursor(void *img, float hotX, float hotY);
void platformDisposeCursor(void *cursor);

#endif // __RW_GOUI_CURSOR__
