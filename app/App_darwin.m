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
#include "app.h"

@interface appDelegate : NSObject<NSApplicationDelegate>
@end

const char *uiAppName() {
    return strdup([[[NSProcessInfo processInfo] processName] UTF8String]);
}

void uiStart() {
    @autoreleasepool {
        [NSApplication sharedApplication];

        // Required for apps without bundle & Info.plist
        [NSApp setActivationPolicy:NSApplicationActivationPolicyRegular];
        [[NSRunningApplication currentApplication] activateWithOptions:NSApplicationActivateIgnoringOtherApps];

        [NSApp setDelegate:[appDelegate new]];
        [NSApp run];
    }
}

void uiAttemptTerminate() {
    [NSApp terminate:nil];
}

void uiAppMayTerminateNow(int terminate) {
    [NSApp replyToApplicationShouldTerminate:terminate];
}

void uiHideApp() {
    [NSApp hide:nil];
}

void uiHideOtherApps() {
    [NSApp hideOtherApplications:NSApp];
}

void uiShowAllApps() {
    [NSApp unhideAllApplications:NSApp];
}

@implementation appDelegate

- (void)applicationWillFinishLaunching:(NSNotification *)aNotification {
    willFinishStartup();
}

- (void)applicationDidFinishLaunching:(NSNotification *)aNotification {
    didFinishStartup();
}

- (NSApplicationTerminateReply)applicationShouldTerminate:(NSApplication *)sender {
    return shouldTerminate();
}

- (BOOL)applicationShouldTerminateAfterLastWindowClosed:(NSApplication *)theApplication {
    return shouldTerminateAfterLastWindowClosed();
}

- (void)applicationWillTerminate:(NSNotification *)aNotification {
    return willTerminate();
}

- (void)applicationWillBecomeActive:(NSNotification *)aNotification {
    return willBecomeActive();
}

- (void)applicationDidBecomeActive:(NSNotification *)aNotification {
    return didBecomeActive();
}

- (void)applicationWillResignActive:(NSNotification *)aNotification {
    return willResignActive();
}

- (void)applicationDidResignActive:(NSNotification *)aNotification {
    return didResignActive();
}

@end
