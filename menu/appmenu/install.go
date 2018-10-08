package appmenu

import (
	"runtime"

	"github.com/richardwilkes/toolbox/i18n"
	"github.com/richardwilkes/ui/app"
	"github.com/richardwilkes/ui/event"
	"github.com/richardwilkes/ui/keys"
	"github.com/richardwilkes/ui/menu"
)

// Install adds a standard 'application' menu to the front of the menu bar.
func Install(bar menu.Bar) (appMenu menu.Menu, aboutItem menu.Item, prefsItem menu.Item) {
	name := app.Name()
	appMenu = menu.NewMenu(name)

	aboutItem = menu.NewItem(i18n.Text("About ")+name, nil)
	appMenu.AppendItem(aboutItem)

	appMenu.AppendItem(menu.NewSeparator())
	prefsItem = menu.NewItemWithKey(i18n.Text("Preferences…"), keys.VirtualKeyComma, nil)
	appMenu.AppendItem(prefsItem)

	if runtime.GOOS == "darwin" {
		appMenu.AppendItem(menu.NewSeparator())
		servicesMenu := menu.NewMenu(i18n.Text("Services"))
		appMenu.AppendMenu(servicesMenu)
		bar.SetupSpecialMenu(menu.ServicesMenu, servicesMenu)
	}

	appMenu.AppendItem(menu.NewSeparator())
	appMenu.AppendItem(menu.NewItemWithKey(i18n.Text("Hide ")+name, keys.VirtualKeyH, func(evt event.Event) { app.Hide() }))
	if runtime.GOOS == "darwin" {
		appMenu.AppendItem(menu.NewItemWithKeyAndModifiers(i18n.Text("Hide Others"), keys.VirtualKeyH, keys.OptionModifier|keys.PlatformMenuModifier(), func(evt event.Event) { app.HideOthers() }))
		appMenu.AppendItem(menu.NewItem(i18n.Text("Show All"), func(evt event.Event) { app.ShowAll() }))
	}

	appMenu.AppendItem(menu.NewSeparator())
	appMenu.AppendItem(menu.NewItemWithKey(i18n.Text("Quit ")+name, keys.VirtualKeyQ, func(evt event.Event) { app.AttemptQuit() }))

	bar.InsertMenu(appMenu, 0)

	return appMenu, aboutItem, prefsItem
}
