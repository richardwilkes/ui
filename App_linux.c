// Copyright (c) 2016 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

#include <X11/Xlib.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include "_cgo_export.h"
#include "globals_linux.h"
#include "App.h"
#include "Window.h"

int getKeyMask(unsigned int state) {
	int keyMask = 0;
	if ((state & LockMask) == LockMask) {
		keyMask |= platformCapsLockKeyMask;
	}
	if ((state & ShiftMask) == ShiftMask) {
		keyMask |= platformShiftKeyMask;
	}
	if ((state & ControlMask) == ControlMask) {
		keyMask |= platformControlKeyMask;
	}
	if ((state & Mod1Mask) == Mod1Mask) {
		keyMask |= platformOptionKeyMask;
	}
	if ((state & Mod4Mask) == Mod4Mask) {
		keyMask |= platformCommandKeyMask;
	}
	return keyMask;
}

int isScrollWheelButton(unsigned int button) {
	return button > 3 && button < 8;
}

int getButton(unsigned int button) {
	if (button == 2) {
		return 2;
	}
	if (button == 3) {
		return 1;
	}
	return 0;
}

const char *platformAppName() {
	// RAW: Implement platformAppName for Linux
	return "<unknown>";
}

void platformStart() {
	memset(&AppGlobals, 0, sizeof(AppGlobals));
	AppGlobals.display = XOpenDisplay(NULL);
	if (!AppGlobals.display) {
		fprintf(stderr, "Failed to open the X11 display\n");
		exit(1);
	}
	AppGlobals.wmDeleteMessage = XInternAtom(AppGlobals.display, "WM_DELETE_WINDOW", False);
	appWillFinishStartup();
	AppGlobals.running = 1;
	appDidFinishStartup();
	if (AppGlobals.windowCount == 0 && appShouldTerminateAfterLastWindowClosed()) {
		platformAttemptTerminate();
	}
	Window lastMouseDownWindow = 0;
	int mouseDownButton = -1;
	while (AppGlobals.running) {
		XEvent event;
		XNextEvent(AppGlobals.display, &event);
		switch (event.type) {
			case KeyPress:
				fprintf(stderr, "KeyPress\n");
				break;
			case KeyRelease:
				fprintf(stderr, "KeyRelease\n");
				break;
			case ButtonPress:
				if (isScrollWheelButton(event.xbutton.button)) {
					int dx = 0;
					int dy = 0;
					switch (event.xbutton.button) {
						case 4: // Up
							dy = -1;
							break;
						case 5: // Down
							dy = 1;
							break;
						case 6:	// Left
							dx = -1;
							break;
						case 7:	// Right
							dx = 1;
							break;
					}
					handleWindowMouseWheelEvent((platformWindow)event.xbutton.window, platformMouseWheel, getKeyMask(event.xbutton.state), event.xbutton.x, event.xbutton.y, dx, dy);
				} else {
					mouseDownButton = getButton(event.xbutton.button);
					lastMouseDownWindow = event.xbutton.window;
					// RAW: Needs concept of click count
					handleWindowMouseEvent((platformWindow)event.xbutton.window, platformMouseDown, getKeyMask(event.xbutton.state), mouseDownButton, 0, event.xbutton.x, event.xbutton.y);
				}
				break;
			case ButtonRelease:
				if (!isScrollWheelButton(event.xbutton.button)) {
					mouseDownButton = -1;
					// RAW: Needs concept of click count
					handleWindowMouseEvent((platformWindow)event.xbutton.window, platformMouseUp, getKeyMask(event.xbutton.state), getButton(event.xbutton.button), 0, event.xbutton.x, event.xbutton.y);
				}
				break;
			case MotionNotify:
				if (mouseDownButton != -1) {
					if (event.xmotion.window != lastMouseDownWindow) {
						// RAW: Translate coordinates appropriately
						printf("need translation for mouse drag\n");
					}
					handleWindowMouseEvent((platformWindow)lastMouseDownWindow, platformMouseDragged, getKeyMask(event.xmotion.state), mouseDownButton, 0, event.xmotion.x, event.xmotion.y);
				} else {
					handleWindowMouseEvent((platformWindow)event.xmotion.window, platformMouseMoved, getKeyMask(event.xmotion.state), 0, 0, event.xmotion.x, event.xmotion.y);
				}
				break;
			case EnterNotify:
				handleWindowMouseEvent((platformWindow)event.xcrossing.window, platformMouseEntered, getKeyMask(event.xcrossing.state), 0, 0, event.xcrossing.x, event.xcrossing.y);
				break;
			case LeaveNotify:
				handleWindowMouseEvent((platformWindow)event.xcrossing.window, platformMouseExited, getKeyMask(event.xcrossing.state), 0, 0, event.xcrossing.x, event.xcrossing.y);
				break;
			case FocusIn:
				appWillBecomeActive();
				appDidBecomeActive();
				break;
			case FocusOut:
				appWillResignActive();
				appDidResignActive();
				break;
			case Expose:
				{
					platformRect bounds;
					bounds.x = event.xexpose.x;
					bounds.y = event.xexpose.y;
					bounds.width = event.xexpose.width;
					bounds.height = event.xexpose.height;
					XGCValues values;
					GC gc = XCreateGC(AppGlobals.display, event.xexpose.window, 0, &values);
					drawWindow((platformWindow)event.xexpose.window, gc, bounds, 0);
					XFreeGC(AppGlobals.display, gc);
				}
				break;
			case DestroyNotify:
				windowDidClose((platformWindow)event.xdestroywindow.window);
				if (AppGlobals.windowCount == 0 && appShouldTerminateAfterLastWindowClosed()) {
					platformAttemptTerminate();
				}
				break;
			case ConfigureNotify:
				windowResized((platformWindow)event.xconfigure.window);
				break;
			case ClientMessage:
				if (event.xclient.data.l[0] == AppGlobals.wmDeleteMessage) {
					if (windowShouldClose((platformWindow)event.xclient.window)) {
						platformCloseWindow((platformWindow)event.xclient.window);
					}
				} else {
					// fprintf(stderr, "Unhandled X11 ClientMessage\n");
				}
				break;
			default:
				// fprintf(stderr, "Unhandled event (type %d)\n", event.type);
				break;
		}
	}
}

void terminateNow() {
	appWillTerminate();
	// RAW: Go through and tell each open window to close, ignoring any that refuse to.
	AppGlobals.running = 0;
	XCloseDisplay(AppGlobals.display);
	AppGlobals.display = NULL;
	exit(0);
}

void platformAttemptTerminate() {
	switch (appShouldTerminate()) {
		case platformTerminateCancel:
			break;
		case platformTerminateLater:
			AppGlobals.awaitingTermination = 1;
			break;
		default:
			terminateNow();
			break;
	}
}

void platformAppMayTerminateNow(int terminate) {
	if (AppGlobals.awaitingTermination) {
		AppGlobals.awaitingTermination = 0;
		if (terminate) {
			terminateNow();
		}
	}
}

void platformHideApp() {
	// RAW: Implement platformHideApp for Linux
}

void platformHideOtherApps() {
	// RAW: Implement platformHideOtherApps for Linux
}

void platformShowAllApps() {
	// RAW: Implement platformShowAllApps for Linux
}
