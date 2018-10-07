package app

import (
	"os"
	"path/filepath"

	"github.com/richardwilkes/ui/menu/custom"
	"github.com/richardwilkes/ui/window"
)

func platformAppStart() {
	custom.Install()
	panic("unimplemented")
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
	panic("unimplemented")
}

func platformMayQuitNow(quit bool) {
	panic("unimplemented")
}
