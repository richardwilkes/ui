// Copyright (c) 2016 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

package event

import (
	"bytes"
	"fmt"
)

// Modified is generated when some widgets are modified.
type Modified struct {
	target   Target
	finished bool
}

// NewModified creates a new Modified event. 'target' is the widget that was modified.
func NewModified(target Target) *Modified {
	return &Modified{target: target}
}

// Type returns the event type ID.
func (e *Modified) Type() Type {
	return ModifiedType
}

// Target the original target of the event.
func (e *Modified) Target() Target {
	return e.target
}

// Cascade returns true if this event should be passed to its target's parent if not marked done.
func (e *Modified) Cascade() bool {
	return false
}

// Finished returns true if this event has been handled and should no longer be processed.
func (e *Modified) Finished() bool {
	return e.finished
}

// Finish marks this event as handled and no longer eligible for processing.
func (e *Modified) Finish() {
	e.finished = true
}

// String implements the fmt.Stringer interface.
func (e *Modified) String() string {
	var buffer bytes.Buffer
	buffer.WriteString(fmt.Sprintf("Modified[Target: %v", e.target))
	if e.finished {
		buffer.WriteString(", Finished")
	}
	buffer.WriteString("]")
	return buffer.String()
}
