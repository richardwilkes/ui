package app

import (
	"os"
	"path/filepath"

	"github.com/richardwilkes/ui/event"
	"github.com/richardwilkes/ui/internal/x11"
	"github.com/richardwilkes/ui/menu/custom"
	"github.com/richardwilkes/ui/window"
)

func platformAppStart() {
	x11.OpenDisplay()
	window.LastWindowClosed = func() {
		if ShouldQuitAfterLastWindowClosed() {
			AttemptQuit()
		}
	}
	custom.Install()
	event.SendAppWillFinishStartup()
	event.SendAppDidFinishStartup()
	if window.Count() == 0 && ShouldQuitAfterLastWindowClosed() {
		AttemptQuit()
	}
	window.RunEventLoop()
}

func platformAppName() string {
	return filepath.Base(os.Args[0])
}

func platformHideApp() {
	for _, wnd := range window.Windows() {
		wnd.Minimize()
	}
}

func platformHideOtherApps() {
	// Not supported
}

func platformShowAllApps() {
	// Not supported
}

func platformAttemptQuit() {
	switch ShouldQuit() {
	case Cancel:
	case Later:
		window.DeferQuit()
	default:
		window.StartQuit()
	}
}

func platformMayQuitNow(quit bool) {
	window.ResumeQuit(quit)
}
