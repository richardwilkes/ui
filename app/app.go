package app

import (
	"fmt"
	"runtime"

	"github.com/richardwilkes/ui/event"
	"github.com/richardwilkes/ui/object"
)

// Possible termination responses
const (
	Cancel QuitResponse = iota
	Now
	Later // Must make a call to MayQuitNow() at some point in the future.
)

var (
	// App provides the top-level event distribution point. Events that
	// cascade will flow from the widgets, to their parents, to their window,
	// then finally to this app if not handled somewhere along the line.
	App *Application
)

// QuitResponse is used to respond to requests for app termination.
type QuitResponse int

type app interface {
	Start()
	Name() string
	Hide()
	HideOthers()
	ShowAll()
	AttemptQuit()
	MayQuitNow(quit bool)
}

// Application represents the overall application.
type Application struct {
	object.Base
	app           app
	eventHandlers event.Handlers
}

func init() {
	App = &Application{app: osApp}
	App.InitTypeAndID(App)
	event.SetGlobalTarget(App)
}

// Start the user interface. Locks the calling goroutine to its current OS
// thread. Does not return.
func (app *Application) Start() {
	runtime.LockOSThread()
	app.app.Start()
}

// Name returns the application's name.
func (app *Application) Name() string {
	return app.app.Name()
}

// Hide attempts to hide this application.
func (app *Application) Hide() {
	app.app.Hide()
}

// HideOthers attempts to hide other applications, leaving just this
// application visible.
func (app *Application) HideOthers() {
	app.app.HideOthers()
}

// ShowAll attempts to show all applications that are currently hidden.
func (app *Application) ShowAll() {
	app.app.ShowAll()
}

func (app *Application) String() string {
	return fmt.Sprintf("Application #%d", app.ID())
}

// EventHandlers implements the event.Target interface.
func (app *Application) EventHandlers() *event.Handlers {
	return &app.eventHandlers
}

// ParentTarget implements the event.Target interface.
func (app *Application) ParentTarget() event.Target {
	return nil
}

// AttemptQuit initiates the termination sequence.
func (app *Application) AttemptQuit() {
	app.app.AttemptQuit()
}

// MayQuitNow resumes the termination sequence that was delayed by calling
// Delay() on the AppTerminationRequested event.
func (app *Application) MayQuitNow(quit bool) {
	app.app.MayQuitNow(quit)
}

// ShouldQuit is called when a request to quit the application is made.
func (app *Application) ShouldQuit() QuitResponse {
	e := event.NewAppQuitRequested(event.GlobalTarget())
	event.Dispatch(e)
	if e.Canceled() {
		return Cancel
	}
	if e.Delayed() {
		return Later
	}
	return Now
}

// ShouldQuitAfterLastWindowClosed is called when the last window is closed
// to determine if the application should quit as a result.
func (app *Application) ShouldQuitAfterLastWindowClosed() bool {
	return event.SendAppLastWindowClosed().Quit()
}
