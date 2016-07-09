// Copyright (c) 2016 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

package main

import (
	"fmt"
	"github.com/richardwilkes/go-ui/Demo/images"
	"github.com/richardwilkes/go-ui/app"
	"github.com/richardwilkes/go-ui/border"
	"github.com/richardwilkes/go-ui/font"
	"github.com/richardwilkes/go-ui/geom"
	"github.com/richardwilkes/go-ui/image"
	"github.com/richardwilkes/go-ui/layout"
	"github.com/richardwilkes/go-ui/widget"
)

var (
	aboutWindow *widget.Window
)

func main() {
	app.WillFinishStartup = func() {
		createMenuBar()
		createButtonsWindow()
	}
	app.ShouldTerminateAfterLastWindowClosed = func() bool { return true }
	app.Start()
}

func createMenuBar() {
	app.AddAppMenu(createAboutWindow, createPreferencesWindow)

	fileMenu := widget.MenuBar().AddMenu("File")
	fileMenu.AddItem("Open", "o", nil, nil)

	widget.MenuBar().AddMenu("Edit")

	app.AddWindowMenu()
	app.AddHelpMenu()
}

func createButtonsWindow() {
	wnd := widget.NewWindow(geom.Point{}, widget.StdWindowMask)
	wnd.SetTitle("Buttons")

	root := wnd.RootBlock()
	root.Border = border.NewEmpty(geom.Insets{Top: 10, Left: 10, Bottom: 10, Right: 10})
	root.Layout = layout.NewPrecision().SetVerticalSpacing(10)

	buttonsPanel := createButtonsPanel()
	buttonsPanel.SetLayoutData(layout.NewPrecisionData().SetHorizontalGrab(true))
	root.AddChild(buttonsPanel)

	addSeparator(root)

	checkBoxPanel := createCheckBoxPanel()
	checkBoxPanel.SetLayoutData(layout.NewPrecisionData().SetHorizontalGrab(true))
	root.AddChild(checkBoxPanel)

	addSeparator(root)

	radioButtonsPanel := createRadioButtonsPanel()
	radioButtonsPanel.SetLayoutData(layout.NewPrecisionData().SetHorizontalGrab(true))
	root.AddChild(radioButtonsPanel)

	addSeparator(root)

	popupMenusPanel := createPopupMenusPanel()
	popupMenusPanel.SetLayoutData(layout.NewPrecisionData().SetHorizontalGrab(true))
	root.AddChild(popupMenusPanel)

	wnd.Pack()
	wnd.ToFront()
}

func addSeparator(root *widget.Block) {
	sep := widget.NewSeparator(true)
	sep.SetLayoutData(layout.NewPrecisionData().SetHorizontalAlignment(layout.Fill))
	root.AddChild(&sep.Block)
}

func createButtonsPanel() *widget.Block {
	panel := widget.NewBlock()
	panel.Layout = &layout.Flow{HGap: 5, VGap: 5}

	createButton("Press Me", panel)
	createButton("Default", panel).SetFocused(true)
	createButton("Disabled", panel).SetDisabled(true)

	img, err := image.AcquireFromFile(images.FS, "/home.png")
	if err == nil {
		createImageButton(img, "Home", panel)
		createImageButton(img, "Home (disabled)", panel).SetDisabled(true)
	} else {
		fmt.Println(err)
	}

	img, err = image.AcquireFromFile(images.FS, "/classic-apple-logo.png")
	if err == nil {
		createImageButton(img, "Classic Apple Logo", panel)
		createImageButton(img, "Classic Apple Logo (disabled)", panel).SetDisabled(true)
	} else {
		fmt.Println(err)
	}

	return panel
}

func createButton(title string, panel *widget.Block) *widget.Button {
	button := widget.NewButton(title)
	button.OnClick = func() { fmt.Printf("The button '%s' was clicked.\n", title) }
	button.OnToolTip = func(where geom.Point) string {
		return fmt.Sprintf("This is the tooltip for the '%s' button.", title)
	}
	panel.AddChild(&button.Block)
	return button
}

func createImageButton(img *image.Image, name string, panel *widget.Block) *widget.ImageButton {
	size := img.Size()
	size.Width /= 2
	size.Height /= 2
	button := widget.NewImageButtonWithImageSize(img, size)
	button.OnClick = func() { fmt.Printf("The button '%s' was clicked.\n", name) }
	button.OnToolTip = func(where geom.Point) string { return name }
	panel.AddChild(&button.Block)
	return button
}

func createCheckBoxPanel() *widget.Block {
	panel := widget.NewBlock()
	panel.Layout = layout.NewPrecision()
	createCheckBox("Press Me", panel)
	createCheckBox("Initially Mixed", panel).SetState(widget.Mixed)
	createCheckBox("Disabled", panel).SetDisabled(true)
	checkbox := createCheckBox("Disabled w/Check", panel)
	checkbox.SetDisabled(true)
	checkbox.SetState(widget.Checked)
	return panel
}

func createCheckBox(title string, panel *widget.Block) *widget.CheckBox {
	checkbox := widget.NewCheckBox(title)
	checkbox.OnClick = func() { fmt.Printf("The checkbox '%s' was clicked.\n", title) }
	checkbox.OnToolTip = func(where geom.Point) string {
		return fmt.Sprintf("This is the tooltip for the '%s' checkbox.", title)
	}
	panel.AddChild(&checkbox.Block)
	return checkbox
}

func createRadioButtonsPanel() *widget.Block {
	panel := widget.NewBlock()
	panel.Layout = layout.NewPrecision()

	group := widget.NewRadioButtonGroup()
	first := createRadioButton("First", panel, group)
	createRadioButton("Second", panel, group)
	createRadioButton("Third (disabled)", panel, group).SetDisabled(true)
	createRadioButton("Fourth", panel, group)
	group.Select(first)

	return panel
}

func createRadioButton(title string, panel *widget.Block, group *widget.RadioButtonGroup) *widget.RadioButton {
	rb := widget.NewRadioButton(title)
	rb.OnClick = func() { fmt.Printf("The radio button '%s' was clicked.\n", title) }
	rb.OnToolTip = func(where geom.Point) string {
		return fmt.Sprintf("This is the tooltip for the '%s' radio button.", title)
	}
	panel.AddChild(&rb.Block)
	group.Add(rb)
	return rb
}

func createPopupMenusPanel() *widget.Block {
	panel := widget.NewBlock()
	panel.Layout = layout.NewPrecision()

	createPopupMenu(panel, 1, "One", "Two", "Three", "", "Four", "Five", "Six")
	createPopupMenu(panel, 2, "Red", "Blue", "Green").SetDisabled(true)

	return panel
}

func createPopupMenu(panel *widget.Block, selection int, titles ...string) *widget.PopupMenu {
	p := widget.NewPopupMenu()
	p.OnToolTip = func(where geom.Point) string {
		return fmt.Sprintf("This is the tooltip for the PopupMenu with %d items.", len(titles))
	}
	for _, title := range titles {
		if title == "" {
			p.AddSeparator()
		} else {
			p.AddItem(title)
		}
	}
	p.SelectIndex(selection)
	p.OnSelection = func() {
		fmt.Printf("The '%v' item was selected from the PopupMenu.\n", p.Selected())
	}
	panel.AddChild(&p.Block)
	return p
}

func createAboutWindow(item *widget.MenuItem) {
	if aboutWindow == nil {
		aboutWindow = widget.NewWindow(geom.Point{}, widget.TitledWindowMask|widget.ClosableWindowMask)
		aboutWindow.DidClose = func() { aboutWindow = nil }
		aboutWindow.SetTitle("About " + app.Name)
		root := aboutWindow.RootBlock()
		root.Border = border.NewEmpty(geom.Insets{Top: 10, Left: 10, Bottom: 10, Right: 10})
		root.Layout = layout.NewPrecision()
		title := widget.NewLabelWithFont(app.Name, font.Acquire(font.EmphasizedSystem))
		ld := layout.NewPrecisionData()
		ld.HorizontalAlignment = layout.Middle
		title.SetLayoutData(ld)
		root.AddChild(&title.Block)
		desc := widget.NewLabel("Simple app to demonstrate the\ncapabilities of the ui framework.")
		root.AddChild(&desc.Block)
		aboutWindow.Pack()
	}
	aboutWindow.ToFront()
}

func createPreferencesWindow(item *widget.MenuItem) {
	fmt.Println("Preferences...")
}
