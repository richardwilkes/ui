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
	"github.com/richardwilkes/ui"
	"github.com/richardwilkes/ui/Demo/images"
	"github.com/richardwilkes/ui/border"
	"github.com/richardwilkes/ui/draw"
	"github.com/richardwilkes/ui/font"
)

var (
	aboutWindow *ui.Window
)

func main() {
	ui.AppWillFinishStartup = func() {
		createMenuBar()
		createButtonsWindow()
	}
	ui.AppShouldTerminateAfterLastWindowClosed = func() bool { return true }
	ui.Start()
}

func createMenuBar() {
	ui.AddAppMenu(createAboutWindow, createPreferencesWindow)

	fileMenu := ui.MenuBar().AddMenu("File")
	fileMenu.AddItem("Open", "o", nil, nil)

	ui.MenuBar().AddMenu("Edit")

	ui.AddWindowMenu()
	ui.AddHelpMenu()
}

func createButtonsWindow() {
	wnd := ui.NewWindow(draw.Point{}, ui.StdWindowMask)
	wnd.SetTitle("Demo")

	root := wnd.RootWidget()
	root.SetBorder(border.NewEmpty(draw.Insets{Top: 10, Left: 10, Bottom: 10, Right: 10}))
	ui.NewPrecisionLayout(root).SetVerticalSpacing(10)

	buttonsPanel := createButtonsPanel()
	buttonsPanel.SetLayoutData(ui.NewPrecisionData().SetHorizontalGrab(true))
	root.AddChild(buttonsPanel)

	addSeparator(root)

	checkBoxPanel := createCheckBoxPanel()
	checkBoxPanel.SetLayoutData(ui.NewPrecisionData().SetHorizontalGrab(true))
	root.AddChild(checkBoxPanel)

	addSeparator(root)

	radioButtonsPanel := createRadioButtonsPanel()
	radioButtonsPanel.SetLayoutData(ui.NewPrecisionData().SetHorizontalGrab(true))
	root.AddChild(radioButtonsPanel)

	addSeparator(root)

	popupMenusPanel := createPopupMenusPanel()
	popupMenusPanel.SetLayoutData(ui.NewPrecisionData().SetHorizontalGrab(true))
	root.AddChild(popupMenusPanel)

	addSeparator(root)

	img, err := draw.AcquireImageFromURL("http://allwallpapersnew.com/wp-content/gallery/stock-photos-for-free/grassy_field_sunset___free_stock_by_kevron2001-d5blgkr.jpg")
	if err == nil {
		content := ui.NewImageLabel(img)
		_, prefSize, _ := ui.ComputeSizes(content, ui.NoLayoutHintSize)
		content.SetSize(prefSize)
		scrollArea := ui.NewScrollArea(content)
		scrollArea.SetLayoutData(ui.NewPrecisionData().SetHorizontalAlignment(draw.AlignFill).SetVerticalAlignment(draw.AlignFill).SetHorizontalGrab(true).SetVerticalGrab(true))
		root.AddChild(scrollArea)
	} else {
		fmt.Println(err)
	}

	wnd.Pack()
	wnd.ToFront()
}

func addSeparator(root ui.Widget) {
	sep := ui.NewSeparator(true)
	sep.SetLayoutData(ui.NewPrecisionData().SetHorizontalAlignment(draw.AlignFill))
	root.AddChild(sep)
}

func createButtonsPanel() *ui.Block {
	panel := ui.NewBlock()
	ui.NewFlowLayout(panel).SetHorizontalSpacing(5).SetVerticalSpacing(5)

	createButton("Press Me", panel)
	createButton("Disabled", panel).SetEnabled(false)

	img, err := draw.AcquireImageFromFile(images.FS, "/home.png")
	if err == nil {
		createImageButton(img, "Home", panel)
		createImageButton(img, "Home (disabled)", panel).SetEnabled(false)
	} else {
		fmt.Println(err)
	}

	img, err = draw.AcquireImageFromFile(images.FS, "/classic-apple-logo.png")
	if err == nil {
		createImageButton(img, "Classic Apple Logo", panel)
		createImageButton(img, "Classic Apple Logo (disabled)", panel).SetEnabled(false)
	} else {
		fmt.Println(err)
	}

	return panel
}

func createButton(title string, panel *ui.Block) *ui.Button {
	button := ui.NewButton(title)
	button.OnClick = func() { fmt.Printf("The button '%s' was clicked.\n", title) }
	button.AddEventHandler(ui.ToolTipEvent, func(event *ui.Event) {
		event.ToolTip = fmt.Sprintf("This is the tooltip for the '%s' button.", title)
	})
	panel.AddChild(button)
	return button
}

func createImageButton(img *draw.Image, name string, panel *ui.Block) *ui.ImageButton {
	size := img.Size()
	size.Width /= 2
	size.Height /= 2
	button := ui.NewImageButtonWithImageSize(img, size)
	button.OnClick = func() { fmt.Printf("The button '%s' was clicked.\n", name) }
	button.AddEventHandler(ui.ToolTipEvent, func(event *ui.Event) { event.ToolTip = name })
	panel.AddChild(button)
	return button
}

func createCheckBoxPanel() *ui.Block {
	panel := ui.NewBlock()
	ui.NewPrecisionLayout(panel)
	createCheckBox("Press Me", panel)
	createCheckBox("Initially Mixed", panel).SetState(ui.Mixed)
	createCheckBox("Disabled", panel).SetEnabled(false)
	checkbox := createCheckBox("Disabled w/Check", panel)
	checkbox.SetEnabled(false)
	checkbox.SetState(ui.Checked)
	return panel
}

func createCheckBox(title string, panel *ui.Block) *ui.CheckBox {
	checkbox := ui.NewCheckBox(title)
	checkbox.OnClick = func() { fmt.Printf("The checkbox '%s' was clicked.\n", title) }
	checkbox.AddEventHandler(ui.ToolTipEvent, func(event *ui.Event) {
		event.ToolTip = fmt.Sprintf("This is the tooltip for the '%s' checkbox.", title)
	})
	panel.AddChild(checkbox)
	return checkbox
}

func createRadioButtonsPanel() *ui.Block {
	panel := ui.NewBlock()
	ui.NewPrecisionLayout(panel)

	group := ui.NewRadioButtonGroup()
	first := createRadioButton("First", panel, group)
	createRadioButton("Second", panel, group)
	createRadioButton("Third (disabled)", panel, group).SetEnabled(false)
	createRadioButton("Fourth", panel, group)
	group.Select(first)

	return panel
}

func createRadioButton(title string, panel *ui.Block, group *ui.RadioButtonGroup) *ui.RadioButton {
	rb := ui.NewRadioButton(title)
	rb.OnClick = func() { fmt.Printf("The radio button '%s' was clicked.\n", title) }
	rb.AddEventHandler(ui.ToolTipEvent, func(event *ui.Event) {
		event.ToolTip = fmt.Sprintf("This is the tooltip for the '%s' radio button.", title)
	})
	panel.AddChild(rb)
	group.Add(rb)
	return rb
}

func createPopupMenusPanel() *ui.Block {
	panel := ui.NewBlock()
	ui.NewPrecisionLayout(panel)

	createPopupMenu(panel, 1, "One", "Two", "Three", "", "Four", "Five", "Six")
	createPopupMenu(panel, 2, "Red", "Blue", "Green").SetEnabled(false)

	return panel
}

func createPopupMenu(panel *ui.Block, selection int, titles ...string) *ui.PopupMenu {
	p := ui.NewPopupMenu()
	p.AddEventHandler(ui.ToolTipEvent, func(event *ui.Event) {
		event.ToolTip = fmt.Sprintf("This is the tooltip for the PopupMenu with %d items.", len(titles))
	})
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
	panel.AddChild(p)
	return p
}

func createAboutWindow(item *ui.MenuItem) {
	if aboutWindow == nil {
		aboutWindow = ui.NewWindow(draw.Point{}, ui.TitledWindowMask|ui.ClosableWindowMask)
		aboutWindow.DidClose = func() { aboutWindow = nil }
		aboutWindow.SetTitle("About " + ui.AppName())
		root := aboutWindow.RootWidget()
		root.SetBorder(border.NewEmpty(draw.Insets{Top: 10, Left: 10, Bottom: 10, Right: 10}))
		ui.NewPrecisionLayout(root)
		title := ui.NewLabelWithFont(ui.AppName(), font.Acquire(font.EmphasizedSystemDesc))
		title.SetLayoutData(ui.NewPrecisionData().SetHorizontalAlignment(draw.AlignMiddle))
		root.AddChild(title)
		desc := ui.NewLabel("Simple app to demonstrate the\ncapabilities of the ui framework.")
		root.AddChild(desc)
		aboutWindow.Pack()
	}
	aboutWindow.ToFront()
}

func createPreferencesWindow(item *ui.MenuItem) {
	fmt.Println("Preferences...")
}

type scrollTarget struct {
	hpos float32
	vpos float32
}

// LineScrollAmount implements ui.Scrollable.
func (st *scrollTarget) LineScrollAmount(horizontal, towardsStart bool) float32 {
	return 1
}

// PageScrollAmount implements ui.Scrollable.
func (st *scrollTarget) PageScrollAmount(horizontal, towardsStart bool) float32 {
	return 10
}

// ScrolledPosition implements ui.Scrollable.
func (st *scrollTarget) ScrolledPosition(horizontal bool) float32 {
	if horizontal {
		return st.hpos
	}
	return st.vpos
}

// SetScrolledPosition implements ui.Scrollable.
func (st *scrollTarget) SetScrolledPosition(horizontal bool, position float32) {
	if horizontal {
		st.hpos = position
	} else {
		st.vpos = position
	}
}

// VisibleSize implements ui.Scrollable.
func (st *scrollTarget) VisibleSize(horizontal bool) float32 {
	return 10
}

// ContentSize implements ui.Scrollable.
func (st *scrollTarget) ContentSize(horizontal bool) float32 {
	return 1000
}
