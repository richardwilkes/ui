// Copyright (c) 2016 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

package event

// Target marks objects that can be the target of an event.
type Target interface {
	// EventHandlers returns the handler mappings for this Target.
	EventHandlers() *Handlers
	// ParentTarget returns the parent of this Target, or nil.
	ParentTarget() Target
}
