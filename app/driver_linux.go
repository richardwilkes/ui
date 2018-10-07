package app

import (
	"os"
	"path/filepath"

	"github.com/richardwilkes/ui/event"
	"github.com/richardwilkes/ui/internal/x11"
	"github.com/richardwilkes/ui/menu/custom"
	"github.com/richardwilkes/ui/window"
)

var osDriver driver = &linuxDriver{}

type linuxDriver struct {
}

func (d *linuxDriver) Start() {
	x11.OpenDisplay()
	window.LastWindowClosed = func() {
		if App.ShouldQuitAfterLastWindowClosed() {
			App.AttemptQuit()
		}
	}
	custom.Install()
	event.SendAppWillFinishStartup()
	event.SendAppDidFinishStartup()
	if window.Count() == 0 && App.ShouldQuitAfterLastWindowClosed() {
		App.AttemptQuit()
	}
	window.RunEventLoop()
}

func (d *linuxDriver) Name() string {
	return filepath.Base(os.Args[0])
}

func (d *linuxDriver) Hide() {
	for _, wnd := range window.Windows() {
		wnd.Minimize()
	}
}

func (d *linuxDriver) HideOthers() {
	panic("unimplemented")
}

func (d *linuxDriver) ShowAll() {
	panic("unimplemented")
}

func (d *linuxDriver) AttemptQuit() {
	switch App.ShouldQuit() {
	case Cancel:
	case Later:
		window.DeferQuit()
	default:
		window.StartQuit()
	}
}

func (d *linuxDriver) MayQuitNow(quit bool) {
	window.ResumeQuit(quit)
}
