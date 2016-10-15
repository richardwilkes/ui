// Copyright (c) 2016 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

package app

import (
	// #cgo CFLAGS: -x objective-c
	// #cgo LDFLAGS: -framework Cocoa
	// #include <Cocoa/Cocoa.h>
	//
	// NSApplicationTerminateReply callbackAppShouldQuit(void);
	// BOOL callbackAppShouldQuitAfterLastWindowClosed(void);
	// void callbackAppWillQuit(void);
	// void callbackAppWillFinishStartup(void);
	// void callbackAppDidFinishStartup(void);
	// void callbackAppWillBecomeActive(void);
	// void callbackAppDidBecomeActive(void);
	// void callbackAppWillResignActive(void);
	// void callbackAppDidResignActive(void);
	//
	// @interface appDelegate : NSObject<NSApplicationDelegate>
	// @end
	//
	// @implementation appDelegate
	// - (void)applicationWillFinishLaunching:(NSNotification *)aNotification { callbackAppWillFinishStartup(); }
	// - (void)applicationDidFinishLaunching:(NSNotification *)aNotification { callbackAppDidFinishStartup(); }
	// - (NSApplicationTerminateReply)applicationShouldTerminate:(NSApplication *)sender { return callbackAppShouldQuit(); } // The Mac response codes map to the same values we use
	// - (BOOL)applicationShouldTerminateAfterLastWindowClosed:(NSApplication *)theApplication { return callbackAppShouldQuitAfterLastWindowClosed(); }
	// - (void)applicationWillTerminate:(NSNotification *)aNotification { callbackAppWillQuit(); }
	// - (void)applicationWillBecomeActive:(NSNotification *)aNotification { callbackAppWillBecomeActive(); }
	// - (void)applicationDidBecomeActive:(NSNotification *)aNotification { callbackAppDidBecomeActive(); }
	// - (void)applicationWillResignActive:(NSNotification *)aNotification { callbackAppWillResignActive(); }
	// - (void)applicationDidResignActive:(NSNotification *)aNotification { callbackAppDidResignActive(); }
	// @end
	//
	// void startUserInterface() {
	//    @autoreleasepool {
	//        [NSApplication sharedApplication];
	//        // Required for apps without bundle & Info.plist
	//        [NSApp setActivationPolicy:NSApplicationActivationPolicyRegular];
	//        [NSApp setDelegate:[appDelegate new]];
	//        // Required to use 'NSApplicationActivateIgnoringOtherApps' otherwise our windows end up in the background.
	//       	[[NSRunningApplication currentApplication] activateWithOptions:NSApplicationActivateAllWindows | NSApplicationActivateIgnoringOtherApps];
	//        [NSApp run];
	//    }
	// }
	//
	// const char *appName() { return [[[NSProcessInfo processInfo] processName] UTF8String]; }
	// void hideApp() { [NSApp hide:nil]; }
	// void hideOtherApps() { [NSApp hideOtherApplications:NSApp]; }
	// void showAllApps() { [NSApp unhideAllApplications:NSApp]; }
	"C"
)

func platformStartUserInterface() {
	C.startUserInterface()
}

func platformAppName() string {
	return C.GoString(C.appName())
}

func platformHideApp() {
	C.hideApp()
}

func platformHideOtherApps() {
	C.hideOtherApps()
}

func platformShowAllApps() {
	C.showAllApps()
}
