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
	"github.com/richardwilkes/go-ui/widget/menu"
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
	menu.AddAppMenu(createAboutWindow, nil)

	fileMenu := menu.Bar().AddMenu("File")
	fileMenu.AddItem("Open", "o", nil, nil)

	menu.Bar().AddMenu("Edit")

	menu.AddWindowMenu()
	menu.AddHelpMenu()
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

	sep := widget.NewSeparator(true)
	sep.SetLayoutData(layout.NewPrecisionData().SetHorizontalAlignment(layout.Fill))
	root.AddChild(&sep.Block)

	checkBoxPanel := createCheckBoxPanel()
	checkBoxPanel.SetLayoutData(layout.NewPrecisionData().SetHorizontalGrab(true))
	root.AddChild(checkBoxPanel)

	wnd.Pack()
	wnd.ToFront()
}

func createButtonsPanel() *widget.Block {
	panel := widget.NewBlock()
	panel.Layout = &layout.Flow{HGap: 5, VGap: 5}

	button := widget.NewButton("Press Me")
	button.OnClick = func() { fmt.Println("The button 'Press Me' was clicked.") }
	button.OnToolTip = func(where geom.Point) string { return "This is the tooltip for the 'Press Me' button." }
	panel.AddChild(&button.Block)

	button = widget.NewButton("Default")
	button.OnClick = func() { fmt.Println("The button 'Default' was clicked.") }
	button.OnToolTip = func(where geom.Point) string { return "This is the tooltip for the 'Default' button." }
	button.SetFocused(true)
	panel.AddChild(&button.Block)

	button = widget.NewButton("Disabled")
	button.OnToolTip = func(where geom.Point) string { return "This is the tooltip for the 'Disabled' button." }
	button.SetDisabled(true)
	panel.AddChild(&button.Block)

	img, err := image.AcquireFromFile(images.FS, "/home.png")
	if err == nil {
		size := img.Size()
		size.Width /= 2
		size.Height /= 2
		imgButton := widget.NewImageButtonWithImageSize(img, size)
		imgButton.OnClick = func() { fmt.Println("The button 'Home' was clicked.") }
		imgButton.OnToolTip = func(where geom.Point) string { return "Home" }
		panel.AddChild(&imgButton.Block)

		imgButton = widget.NewImageButtonWithImageSize(img, size)
		imgButton.OnToolTip = func(where geom.Point) string { return "Disabled Home" }
		imgButton.SetDisabled(true)
		panel.AddChild(&imgButton.Block)
	} else {
		fmt.Println(err)
	}

	img, err = image.AcquireFromFile(images.FS, "/classic-apple-logo.png")
	if err == nil {
		size := img.Size()
		size.Width /= 2
		size.Height /= 2
		imgButton := widget.NewImageButtonWithImageSize(img, size)
		imgButton.OnClick = func() { fmt.Println("The button 'Classic Apple Logo' was clicked.") }
		imgButton.OnToolTip = func(where geom.Point) string { return "Classic Apple Logo" }
		panel.AddChild(&imgButton.Block)

		imgButton = widget.NewImageButtonWithImageSize(img, size)
		imgButton.OnToolTip = func(where geom.Point) string { return "Disabled Classic Apple Logo" }
		imgButton.SetDisabled(true)
		panel.AddChild(&imgButton.Block)
	} else {
		fmt.Println(err)
	}

	return panel
}

func createCheckBoxPanel() *widget.Block {
	panel := widget.NewBlock()
	panel.Layout = layout.NewPrecision()

	checkbox := widget.NewCheckBox("Press Me")
	checkbox.OnClick = func() { fmt.Println("The checkbox 'Press Me' was clicked.") }
	checkbox.OnToolTip = func(where geom.Point) string { return "This is the tooltip for the 'Press Me' checkbox." }
	panel.AddChild(&checkbox.Block)

	checkbox = widget.NewCheckBox("Initially Mixed")
	checkbox.OnClick = func() { fmt.Println("The checkbox 'Initially Mixed' was clicked.") }
	checkbox.OnToolTip = func(where geom.Point) string { return "This is the tooltip for the 'Initially Mixed' checkbox." }
	checkbox.SetState(widget.Mixed)
	panel.AddChild(&checkbox.Block)

	checkbox = widget.NewCheckBox("Disabled")
	checkbox.OnToolTip = func(where geom.Point) string { return "This is the tooltip for the 'Disabled' checkbox." }
	checkbox.SetDisabled(true)
	panel.AddChild(&checkbox.Block)

	checkbox = widget.NewCheckBox("Disabled w/Check")
	checkbox.OnToolTip = func(where geom.Point) string { return "This is the tooltip for the 'Disabled w/Check' checkbox." }
	checkbox.SetDisabled(true)
	checkbox.SetState(widget.Checked)
	panel.AddChild(&checkbox.Block)

	return panel
}

func createAboutWindow(item *menu.Item) {
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
