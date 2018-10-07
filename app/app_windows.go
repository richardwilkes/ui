package app

import (
	"os"
	"path/filepath"

	"github.com/richardwilkes/ui/menu/custom"
	"github.com/richardwilkes/ui/window"
)

var osApp app = &windowsApp{}

type windowsApp struct {
}

func (a *windowsApp) Start() {
	custom.Install()
	panic("unimplemented")
}

func (a *windowsApp) Name() string {
	return filepath.Base(os.Args[0])
}

func (a *windowsApp) Hide() {
	for _, wnd := range window.Windows() {
		wnd.Minimize()
	}
}

func (a *windowsApp) HideOthers() {
	panic("unimplemented")
}

func (a *windowsApp) ShowAll() {
	panic("unimplemented")
}

func (a *windowsApp) AttemptQuit() {
	panic("unimplemented")
}

func (a *windowsApp) MayQuitNow(quit bool) {
	panic("unimplemented")
}
