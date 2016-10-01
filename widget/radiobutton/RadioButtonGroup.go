// Copyright (c) 2016 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

package radiobutton

// Group is used to ensure only one RadioButton in a group is selected at a time.
type Group struct {
	members []*RadioButton
}

// NewGroup creates a new group for the specified buttons. Each button is removed from
// any other group it may be in and placed in the newly created one.
func NewGroup(buttons ...*RadioButton) *Group {
	group := &Group{members: buttons}
	for _, button := range buttons {
		group.Add(button)
	}
	return group
}

// Add a button to the group, removing it from any group it may have previously been associated with.
func (g *Group) Add(button *RadioButton) {
	if button.group != nil {
		button.group.Remove(button)
	}
	button.group = g
	g.members = append(g.members, button)
}

// Remove a button from the group.
func (g *Group) Remove(button *RadioButton) {
	if button.group == g {
		for i, one := range g.members {
			if one == button {
				copy(g.members[i:], g.members[i+1:])
				g.members[len(g.members)-1] = nil
				g.members = g.members[:len(g.members)-1]
				button.group = nil
				break
			}
		}
	}
}

// Select a button, deselecting all others in the group.
func (g *Group) Select(button *RadioButton) {
	if button.group == g {
		for _, one := range g.members {
			one.setSelected(one == button)
		}
	}
}
