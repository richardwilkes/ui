package clipboard

import (
	"github.com/richardwilkes/ui/clipboard/datatypes"
)

var (
	lastChangeCount = -1
	dataTypes       []string
)

// Clear the clipboard contents.
func Clear() {
	platformClear()
}

// HasType returns true if the specified data type exists on the clipboard.
func HasType(dataType string) bool {
	for _, one := range Types() {
		if one == dataType {
			return true
		}
	}
	return false
}

// Types returns the types of data currently on the clipboard.
func Types() []string {
	changeCount := platformChangeCount()
	if changeCount != lastChangeCount {
		lastChangeCount = changeCount
		dataTypes = platformTypes()
	}
	return dataTypes
}

// Data returns the bytes associated with the specified data type on the clipboard. An empty byte
// slice will be returned if no such data type is present.
func Data(dataType string) []byte {
	return platformGetData(dataType)
}

// SetData sets the data into the system clipboard.
func SetData(data ...datatypes.Data) {
	platformSetData(data)
}
