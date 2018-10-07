package clipboard

import (
	"github.com/richardwilkes/ui/clipboard/datatypes"
	"github.com/richardwilkes/ui/internal/x11"
)

func platformChangeCount() int {
	return x11.ClipboardChangeCount()
}

func platformClear() {
	x11.ClipboardClear()
}

func platformTypes() []string {
	return x11.ClipboardTypes()
}

func platformGetData(dataType string) []byte {
	return x11.GetClipboard(dataType)
}

func platformSetData(data []datatypes.Data) {
	x11.SetClipboard(data)
}
