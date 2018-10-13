package windowmenu

import (
	"github.com/richardwilkes/toolbox/i18n"
	"github.com/richardwilkes/ui/event"
	"github.com/richardwilkes/ui/keys"
	"github.com/richardwilkes/ui/menu"
	"github.com/richardwilkes/ui/window"
)

// Install adds a standard 'Window' menu to the end of the menu bar.
func Install(bar menu.Bar) {
	windowMenu := menu.NewMenu(i18n.Text("Window"))

	item := menu.NewItemWithKey(i18n.Text("Minimize"), keys.VirtualKeyM, func(evt event.Event) {
		wnd := window.KeyWindow()
		if wnd != nil {
			wnd.Minimize()
		}
	})
	item.EventHandlers().Add(event.ValidateType, func(evt event.Event) {
		w := window.KeyWindow()
		if w == nil || !w.Minimizable() {
			evt.(*event.Validate).MarkInvalid()
		}
	})
	windowMenu.AppendItem(item)

	item = menu.NewItemWithKeyAndModifiers(i18n.Text("Zoom"), keys.VirtualKeyZ, keys.ShiftModifier|keys.PlatformMenuModifier(), func(evt event.Event) {
		wnd := window.KeyWindow()
		if wnd != nil {
			wnd.Zoom()
		}
	})
	item.EventHandlers().Add(event.ValidateType, func(evt event.Event) {
		w := window.KeyWindow()
		if w == nil || !w.Resizable() {
			evt.(*event.Validate).MarkInvalid()
		}
	})
	windowMenu.AppendItem(item)
	windowMenu.AppendItem(menu.NewSeparator())

	windowMenu.AppendItem(menu.NewItem(i18n.Text("Bring All to Front"), func(evt event.Event) { window.AllWindowsToFront() }))

	bar.AppendMenu(windowMenu)
	bar.SetupSpecialMenu(menu.WindowMenu, windowMenu)
}
