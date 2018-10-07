package app

import (
	"os"
	"path/filepath"

	"github.com/richardwilkes/ui/event"
	"github.com/richardwilkes/ui/internal/x11"
	"github.com/richardwilkes/ui/menu/custom"
	"github.com/richardwilkes/ui/window"
)

var osApp app = &linuxApp{}

type linuxApp struct {
}

func (a *linuxApp) Start() {
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

func (a *linuxApp) Name() string {
	return filepath.Base(os.Args[0])
}

func (a *linuxApp) Hide() {
	for _, wnd := range window.Windows() {
		wnd.Minimize()
	}
}

func (a *linuxApp) HideOthers() {
	panic("unimplemented")
}

func (a *linuxApp) ShowAll() {
	panic("unimplemented")
}

func (a *linuxApp) AttemptQuit() {
	switch App.ShouldQuit() {
	case Cancel:
	case Later:
		window.DeferQuit()
	default:
		window.StartQuit()
	}
}

func (a *linuxApp) MayQuitNow(quit bool) {
	window.ResumeQuit(quit)
}
