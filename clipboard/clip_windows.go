package clipboard

import (
	"github.com/richardwilkes/ui/clipboard/datatypes"
)

var (
	clipData            = make(map[string][]byte)
	clipDataChangeCount int
)

func platformChangeCount() int {
	// RAW: Implement for Windows (i.e. cross-app support)
	return clipDataChangeCount
}

func platformClear() {
	// RAW: Implement for Windows (i.e. cross-app support)
	clipData = make(map[string][]byte)
	clipDataChangeCount++
}

func platformTypes() []string {
	// RAW: Implement for Windows (i.e. cross-app support)
	types := make([]string, len(clipData))
	i := 0
	for key := range clipData {
		types[i] = key
		i++
	}
	return types
}

func platformGetData(dataType string) []byte {
	// RAW: Implement for Windows (i.e. cross-app support)
	return clipData[dataType]
}

func platformSetData(data []datatypes.Data) {
	// RAW: Implement for Windows (i.e. cross-app support)
	platformClear()
	for _, one := range data {
		clipData[one.MimeType] = one.Bytes
	}
}
