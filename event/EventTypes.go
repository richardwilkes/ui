// Copyright (c) 2016 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

package event

// The event types
const (
	AppWillFinishStartupType Type = iota
	AppDidFinishStartupType
	AppWillActivateType
	AppDidActivateType
	AppWillDeactivateType
	AppDidDeactivateType
	AppTerminationRequestedType
	AppWillTerminateType
	AppLastWindowClosedType
	PaintType
	MouseDownType
	MouseDraggedType
	MouseUpType
	MouseEnteredType
	MouseMovedType
	MouseExitedType
	MouseWheelType
	FocusGainedType
	FocusLostType
	KeyDownType
	KeyTypedType
	KeyUpType
	ToolTipType
	ResizedType
	ClosingType
	ClosedType
	// UserType should be used as the base value for custom application events.
	UserType = 10000
)

// Type holds a unique type ID for each event type.
type Type int
