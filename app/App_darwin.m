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
#include "App_darwin.h"

@interface appDelegate : NSObject<NSApplicationDelegate>
@end

void platformStartUserInterface() {
    @autoreleasepool {
        [NSApplication sharedApplication];
        // Required for apps without bundle & Info.plist
        [NSApp setActivationPolicy:NSApplicationActivationPolicyRegular];
        [NSApp setDelegate:[appDelegate new]];
        // Required to use 'NSApplicationActivateIgnoringOtherApps' otherwise our windows
        // end up in the background.
       	[[NSRunningApplication currentApplication] activateWithOptions:NSApplicationActivateAllWindows | NSApplicationActivateIgnoringOtherApps];
        [NSApp run];
    }
}

const char *platformAppName() {
    return [[[NSProcessInfo processInfo] processName] UTF8String];
}

void platformHideApp() {
    [NSApp hide:nil];
}

void platformHideOtherApps() {
    [NSApp hideOtherApplications:NSApp];
}

void platformShowAllApps() {
    [NSApp unhideAllApplications:NSApp];
}

@implementation appDelegate

- (void)applicationWillFinishLaunching:(NSNotification *)aNotification {
    appWillFinishStartup();
}

- (void)applicationDidFinishLaunching:(NSNotification *)aNotification {
    appDidFinishStartup();
}

- (NSApplicationTerminateReply)applicationShouldTerminate:(NSApplication *)sender {
	// The Mac response codes map to the same values we use
    return callbackAppShouldQuit();
}

- (BOOL)applicationShouldTerminateAfterLastWindowClosed:(NSApplication *)theApplication {
    return callbackAppShouldQuitAfterLastWindowClosed();
}

- (void)applicationWillTerminate:(NSNotification *)aNotification {
    return callbackAppWillQuit();
}

- (void)applicationWillBecomeActive:(NSNotification *)aNotification {
    return appWillBecomeActive();
}

- (void)applicationDidBecomeActive:(NSNotification *)aNotification {
    return appDidBecomeActive();
}

- (void)applicationWillResignActive:(NSNotification *)aNotification {
    return appWillResignActive();
}

- (void)applicationDidResignActive:(NSNotification *)aNotification {
    return appDidResignActive();
}

@end
