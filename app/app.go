package app

import (
	"fmt"
	"runtime"

	"github.com/richardwilkes/ui/event"
	"github.com/richardwilkes/ui/object"
)

var (
	// App provides the top-level event distribution point. Events that
	// cascade will flow from the widgets, to their parents, to their window,
	// then finally to this app if not handled somewhere along the line.
	App app
)

// Application represents the overall application.
type app struct {
	object.Base
	eventHandlers event.Handlers
}

func init() {
	App.InitTypeAndID(&App)
	event.SetGlobalTarget(&App)
}

func (a *app) String() string {
	return fmt.Sprintf("Application #%d", a.ID())
}

func (a *app) EventHandlers() *event.Handlers {
	return &a.eventHandlers
}

func (a *app) ParentTarget() event.Target {
	return nil
}

// Start the user interface. Locks the calling goroutine to its current OS
// thread. Does not return.
func Start() {
	runtime.LockOSThread()
	platformAppStart()
}

// Name returns the application's name.
func Name() string {
	return platformAppName()
}

// Hide attempts to hide this application.
func Hide() {
	platformHideApp()
}

// HideOthers attempts to hide other applications, leaving just this
// application visible.
func HideOthers() {
	platformHideOtherApps()
}

// ShowAll attempts to show all applications that are currently hidden.
func ShowAll() {
	platformShowAllApps()
}

// EventHandlers is a shortcut for calling app.App.EventHandlers().
func EventHandlers() *event.Handlers {
	return App.EventHandlers()
}
