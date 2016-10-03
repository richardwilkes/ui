// Copyright (c) 2016 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

package quit

import (
	"github.com/richardwilkes/ui/event"
)

// AttemptQuit initiates the termination sequence.
func AttemptQuit() {
	platformAttemptQuit()
}

// MayQuitNow resumes the termination sequence that was delayed by calling Delay() on the
// AppTerminationRequested event.
func MayQuitNow(quit bool) {
	platformAppMayQuitNow(quit)
}

func AppShouldQuit() Response {
	e := event.NewAppQuitRequested(event.GlobalTarget())
	event.Dispatch(e)
	if e.Cancelled() {
		return Cancel
	}
	if e.Delayed() {
		return Later
	}
	return Now
}

func AppShouldQuitAfterLastWindowClosed() bool {
	e := event.NewAppLastWindowClosed(event.GlobalTarget())
	event.Dispatch(e)
	return e.Quit()
}
