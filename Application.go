// Copyright (c) 2016 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

package ui

import (
	"fmt"
	"github.com/richardwilkes/ui/event"
	"github.com/richardwilkes/ui/id"
	"github.com/richardwilkes/ui/menu"
	"runtime"
)

// Application represents the overall application.
type Application struct {
	id            int64
	eventHandlers event.Handlers
}

var (
	// App provides the top-level event distribution point. Events that cascade will flow from the
	// widgets, to their parents, to their window, then finally to this app if not handled somewhere
	// along the line.
	App Application
)

func (app *Application) String() string {
	return fmt.Sprintf("Application #%d", app.ID())
}

// ID returns the unique ID for this application.
func (app *Application) ID() int64 {
	if app.id == 0 {
		app.id = id.NextID()
	}
	return app.id
}

// EventHandlers implements the event.Target interface.
func (app *Application) EventHandlers() *event.Handlers {
	return &app.eventHandlers
}

// ParentTarget implements the event.Target interface.
func (app *Application) ParentTarget() event.Target {
	return nil
}

// StartUserInterface starts the user interface. Locks the calling goroutine to its current OS
// thread. Does not return.
func StartUserInterface() {
	runtime.LockOSThread()
	menu.SetParentTarget(&App)
	platformStartUserInterface()
}

// AppName returns the application's name.
func AppName() string {
	return platformAppName()
}

// HideApp attempts to hide this application.
func HideApp() {
	platformHideApp()
}

// HideOtherApps attempts to hide other applications, leaving just this application visible.
func HideOtherApps() {
	platformHideOtherApps()
}

// ShowAllApps attempts to show all applications that are currently hidden.
func ShowAllApps() {
	platformShowAllApps()
}

// AttemptQuit initiates the termination sequence.
func AttemptQuit() {
	platformAttemptQuit()
}

// MayQuitNow resumes the termination sequence that was delayed by calling Delay() on the
// AppTerminationRequested event.
func MayQuitNow(quit bool) {
	platformAppMayQuitNow(quit)
}
