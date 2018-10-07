package app

import (
	// #cgo CFLAGS: -x objective-c
	// #cgo LDFLAGS: -framework Cocoa
	// #include "driver_darwin.h"
	"C"

	"github.com/richardwilkes/ui/event"
	"github.com/richardwilkes/ui/menu/macmenus"
)

var osDriver driver = &darwinDriver{}

type darwinDriver struct {
}

func (d *darwinDriver) Start() {
	C.startUserInterface()
}

func (d *darwinDriver) Name() string {
	return C.GoString(C.appName())
}

func (d *darwinDriver) Hide() {
	C.hideApp()
}

func (d *darwinDriver) HideOthers() {
	C.hideOtherApps()
}

func (d *darwinDriver) ShowAll() {
	C.showAllApps()
}

func (d *darwinDriver) AttemptQuit() {
	C.attemptQuit()
}

func (d *darwinDriver) MayQuitNow(quit bool) {
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
