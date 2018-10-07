package app

import (
	// #cgo CFLAGS: -x objective-c
	// #cgo LDFLAGS: -framework Cocoa
	// #include "app_darwin.h"
	"C"

	"github.com/richardwilkes/ui/event"
	"github.com/richardwilkes/ui/menu/macmenus"
)

var osApp app = &darwinApp{}

type darwinApp struct {
}

func (a *darwinApp) Start() {
	C.startUserInterface()
}

func (a *darwinApp) Name() string {
	return C.GoString(C.appName())
}

func (a *darwinApp) Hide() {
	C.hideApp()
}

func (a *darwinApp) HideOthers() {
	C.hideOtherApps()
}

func (a *darwinApp) ShowAll() {
	C.showAllApps()
}

func (a *darwinApp) AttemptQuit() {
	C.attemptQuit()
}

func (a *darwinApp) MayQuitNow(quit bool) {
	var mayQuit C.int
	if quit {
		mayQuit = 1
	}
	C.appMayQuitNow(mayQuit)
}

//export cbAppShouldQuit
func cbAppShouldQuit() int {
	return int(App.ShouldQuit())
}

//export cbAppShouldQuitAfterLastWindowClosed
func cbAppShouldQuitAfterLastWindowClosed() bool {
	return App.ShouldQuitAfterLastWindowClosed()
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
