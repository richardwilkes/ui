package app

import (
	"github.com/richardwilkes/ui/event"
)

// Possible termination responses
const (
	Cancel QuitResponse = iota
	Now
	Later // Must make a call to MayQuitNow() at some point in the future.
)

// QuitResponse is used to respond to requests for app termination.
type QuitResponse int

// AttemptQuit initiates the termination sequence.
func AttemptQuit() {
	platformAttemptQuit()
}

// MayQuitNow resumes the termination sequence that was delayed by calling
// Delay() on the AppTerminationRequested event.
func MayQuitNow(quit bool) {
	platformMayQuitNow(quit)
}

// ShouldQuit is called when a request to quit the application is made.
func ShouldQuit() QuitResponse {
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
func ShouldQuitAfterLastWindowClosed() bool {
	return event.SendAppLastWindowClosed().Quit()
}
