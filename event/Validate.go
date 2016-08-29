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

// Validate is generated when a widget needs to validate its state.
type Validate struct {
	target   Target
	invalid  bool
	finished bool
}

// NewValidate creates a new Validate event. 'target' is the widget needing validation.
func NewValidate(target Target) *Validate {
	return &Validate{target: target}
}

// Type returns the event type ID.
func (e *Validate) Type() Type {
	return ValidateType
}

// Target the original target of the event.
func (e *Validate) Target() Target {
	return e.target
}

// Cascade returns true if this event should be passed to its target's parent if not marked done.
func (e *Validate) Cascade() bool {
	return false
}

// Finished returns true if this event has been handled and should no longer be processed.
func (e *Validate) Finished() bool {
	return e.finished
}

// Finish marks this event as handled and no longer eligible for processing.
func (e *Validate) Finish() {
	e.finished = true
}

// Valid returns true if Validate should not proceed.
func (e *Validate) Valid() bool {
	return !e.invalid
}

// MarkInvalid marks the event as finding invalid state.
func (e *Validate) MarkInvalid() {
	e.invalid = true
	e.finished = true
}

// String implements the fmt.Stringer interface.
func (e *Validate) String() string {
	var buffer bytes.Buffer
	buffer.WriteString("Validate[")
	if e.invalid {
		buffer.WriteString("Invalid")
	} else {
		buffer.WriteString("Valid")
	}
	buffer.WriteString(fmt.Sprintf(", Target: %v", e.target))
	if e.finished {
		buffer.WriteString(", Finished")
	}
	buffer.WriteString("]")
	return buffer.String()
}
