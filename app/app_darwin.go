package app

import (
	// #cgo CFLAGS: -x objective-c
	// #cgo LDFLAGS: -framework Cocoa
	// #include "app_darwin.h"
	"C"

	"github.com/richardwilkes/ui/event"
	"github.com/richardwilkes/ui/menu/macmenus"
)

func platformAppStart() {
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

func platformAttemptQuit() {
	C.attemptQuit()
}

func platformMayQuitNow(quit bool) {
	var mayQuit C.int
	if quit {
		mayQuit = 1
	}
	C.appMayQuitNow(mayQuit)
}

//export cbAppShouldQuit
func cbAppShouldQuit() int {
	return int(ShouldQuit())
}

//export cbAppShouldQuitAfterLastWindowClosed
func cbAppShouldQuitAfterLastWindowClosed() bool {
	return ShouldQuitAfterLastWindowClosed()
}

//export cbAppWillQuit
func cbAppWillQuit() {
	event.SendAppWillQuit()
}

//export cbAppWillFinishStartup
func cbAppWillFinishStartup() {
	macmenus.Install()
	event.SendAppWillFinishStartup()
}

//export cbAppDidFinishStartup
func cbAppDidFinishStartup() {
	event.SendAppDidFinishStartup()
}

//export cbAppWillBecomeActive
func cbAppWillBecomeActive() {
	event.SendAppWillActivate()
}

//export cbAppDidBecomeActive
func cbAppDidBecomeActive() {
	event.SendAppDidActivate()
}

//export cbAppWillResignActive
func cbAppWillResignActive() {
	event.SendAppWillDeactivate()
}

//export cbAppDidResignActive
func cbAppDidResignActive() {
	event.SendAppDidDeactivate()
}
