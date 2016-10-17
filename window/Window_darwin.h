// Copyright (c) 2016 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

#include <cairo.h>
#include "Types.h"

#ifndef __RW_GOUI_WINDOW__
#define __RW_GOUI_WINDOW__

platformWindow platformGetKeyWindow();
void platformBringAllWindowsToFront();
void platformHideCursorUntilMouseMoves();
platformWindow platformNewWindow(double x, double y, double width, double height, int styleMask);
void platformCloseWindow(platformWindow window);
const char *platformGetWindowTitle(platformWindow window);
void platformSetWindowTitle(platformWindow window, const char *title);
void platformGetWindowFrame(platformWindow window, double *x, double *y, double *width, double *height);
void platformSetWindowFrame(platformWindow window, double x, double y, double width, double height);
void platformGetWindowContentFrame(platformWindow window, double *x, double *y, double *width, double *height);
void platformBringWindowToFront(platformWindow window);
void platformRepaintWindow(platformWindow window, double x, double y, double width, double height);
void platformFlushPainting(platformWindow window);
float platformGetWindowScalingFactor(platformWindow window);
void platformMinimizeWindow(platformWindow window);
void platformZoomWindow(platformWindow window);
void platformSetToolTip(platformWindow window, const char *tooltip);
void platformSetCursor(platformWindow window, void *cursor);
cairo_t *platformGraphics(platformWindow window);
void platformInvoke(unsigned long id);
void platformInvokeAfter(unsigned long id, long afterNanos);

#endif // __RW_GOUI_WINDOW__
