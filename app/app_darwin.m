#include "app_darwin.h"
#include "_cgo_export.h"

@interface appDelegate : NSObject<NSApplicationDelegate>
@end

@implementation appDelegate

- (void)applicationWillFinishLaunching:(NSNotification *)aNotification {
	cbAppWillFinishStartup();
}

- (void)applicationDidFinishLaunching:(NSNotification *)aNotification {
	cbAppDidFinishStartup();
}

- (NSApplicationTerminateReply)applicationShouldTerminate:(NSApplication *)sender {
	// The Mac response codes map to the same values we use
	return cbAppShouldQuit();
}

- (BOOL)applicationShouldTerminateAfterLastWindowClosed:(NSApplication *)theApplication {
	return cbAppShouldQuitAfterLastWindowClosed();
}

- (void)applicationWillTerminate:(NSNotification *)aNotification {
	cbAppWillQuit();
}

- (void)applicationWillBecomeActive:(NSNotification *)aNotification {
	cbAppWillBecomeActive();
}

- (void)applicationDidBecomeActive:(NSNotification *)aNotification {
	cbAppDidBecomeActive();
}

- (void)applicationWillResignActive:(NSNotification *)aNotification {
	cbAppWillResignActive();
}

- (void)applicationDidResignActive:(NSNotification *)aNotification {
	cbAppDidResignActive();
}

@end

void startUserInterface() {
	@autoreleasepool {
		[NSApplication sharedApplication];
		// Required for apps without bundle & Info.plist
		[NSApp setActivationPolicy:NSApplicationActivationPolicyRegular];
		[NSApp setDelegate:[appDelegate new]];
		// Required to use 'NSApplicationActivateIgnoringOtherApps' otherwise our windows end up in the background.
		[[NSRunningApplication currentApplication] activateWithOptions:NSApplicationActivateAllWindows | NSApplicationActivateIgnoringOtherApps];
		[NSApp run];
	}
}

const char *appName() {
	return [[[NSProcessInfo processInfo] processName] UTF8String];
}

void hideApp() {
	[NSApp hide:nil];
}

void hideOtherApps() {
	[NSApp hideOtherApplications:NSApp];
}

void showAllApps() {
	[NSApp unhideAllApplications:NSApp];
}

void attemptQuit() {
	[NSApp terminate:nil];
}

void appMayQuitNow(int quit) {
	[NSApp replyToApplicationShouldTerminate:quit];
}
