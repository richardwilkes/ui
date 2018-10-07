package app

import (
	"os"
	"path/filepath"

	"github.com/richardwilkes/ui/menu/custom"
	"github.com/richardwilkes/ui/window"
)

var osDriver driver = &windowsDriver{}

type windowsDriver struct {
}

func (d *windowsDriver) Start() {
	custom.Install()
	panic("unimplemented")
}

func (d *windowsDriver) Name() string {
	return filepath.Base(os.Args[0])
}

func (d *windowsDriver) Hide() {
	for _, wnd := range window.Windows() {
		wnd.Minimize()
	}
}

func (d *windowsDriver) HideOthers() {
	panic("unimplemented")
}

func (d *windowsDriver) ShowAll() {
	panic("unimplemented")
}

func (d *windowsDriver) AttemptQuit() {
	panic("unimplemented")
}

func (d *windowsDriver) MayQuitNow(quit bool) {
	panic("unimplemented")
}
