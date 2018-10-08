package flex

import (
	"github.com/richardwilkes/toolbox/xmath/geom"
	"github.com/richardwilkes/ui"
	"github.com/richardwilkes/ui/draw/align"
	"github.com/richardwilkes/ui/layout"
)

// Data is used to control how an object is laid out by the Flex layout.
type Data struct {
	HSpan        int             // Number of columns the widget should span
	VSpan        int             // Number of rows the widget should span
	SizeHint     geom.Size       // Hint requesting a particular size of the widget
	MinSize      geom.Size       // Override for the minimum size of the widget
	HAlign       align.Alignment // Horizontal alignment of the widget within its space
	VAlign       align.Alignment // Vertical alignment of the widget within its space
	HGrab        bool            // Grab excess horizontal space if true
	VGrab        bool            // Grab excess vertical space if true
	cacheSize    geom.Size
	minCacheSize geom.Size
}

// NewData creates a new Data.
func NewData() *Data {
	return &Data{
		HSpan:    1,
		VSpan:    1,
		HAlign:   align.Start,
		VAlign:   align.Middle,
		SizeHint: layout.NoHintSize,
		MinSize:  layout.NoHintSize}
}

// Clone the Data.
func (data *Data) Clone() *Data {
	d := *data
	d.normalizeAndResetCache()
	return &d
}

func (data *Data) normalizeAndResetCache() {
	data.minCacheSize.Width = 0
	data.minCacheSize.Height = 0
	data.cacheSize.Width = 0
	data.cacheSize.Height = 0
	if data.HSpan < 1 {
		data.HSpan = 1
	}
	if data.VSpan < 1 {
		data.VSpan = 1
	}
	if data.SizeHint.Width < layout.NoHint {
		data.SizeHint.Width = layout.NoHint
	}
	if data.SizeHint.Height < layout.NoHint {
		data.SizeHint.Height = layout.NoHint
	}
	if data.MinSize.Width < layout.NoHint {
		data.MinSize.Width = layout.NoHint
	}
	if data.MinSize.Height < layout.NoHint {
		data.MinSize.Height = layout.NoHint
	}
}

func (data *Data) computeCacheSize(target ui.Widget, hint geom.Size, useMinimumSize bool) {
	data.normalizeAndResetCache()
	min, pref, max := ui.Sizes(target, hint)
	if hint.Width != layout.NoHint || hint.Height != layout.NoHint {
		if data.MinSize.Width != layout.NoHint {
			data.minCacheSize.Width = data.MinSize.Width
		} else {
			data.minCacheSize.Width = min.Width
		}
		if hint.Width != layout.NoHint && hint.Width < data.minCacheSize.Width {
			hint.Width = data.minCacheSize.Width
		}
		if hint.Width != layout.NoHint && hint.Width > max.Width {
			hint.Width = max.Width
		}
		if data.MinSize.Height != layout.NoHint {
			data.minCacheSize.Height = data.MinSize.Height
		} else {
			data.minCacheSize.Height = min.Height
		}
		if hint.Height != layout.NoHint && hint.Height < data.minCacheSize.Height {
			hint.Height = data.minCacheSize.Height
		}
		if hint.Height != layout.NoHint && hint.Height > max.Height {
			hint.Height = max.Height
		}
	}
	if useMinimumSize {
		data.cacheSize = min
		if data.MinSize.Width != layout.NoHint {
			data.minCacheSize.Width = data.MinSize.Width
		} else {
			data.minCacheSize.Width = min.Width
		}
		if data.MinSize.Height != layout.NoHint {
			data.minCacheSize.Height = data.MinSize.Height
		} else {
			data.minCacheSize.Height = min.Height
		}
	} else {
		data.cacheSize = pref
	}
	if hint.Width != layout.NoHint {
		data.cacheSize.Width = hint.Width
	}
	if data.MinSize.Width != layout.NoHint && data.cacheSize.Width < data.MinSize.Width {
		data.cacheSize.Width = data.MinSize.Width
	}
	if data.SizeHint.Width != layout.NoHint {
		data.cacheSize.Width = data.SizeHint.Width
	}
	if hint.Height != layout.NoHint {
		data.cacheSize.Height = hint.Height
	}
	if data.MinSize.Height != layout.NoHint && data.cacheSize.Height < data.MinSize.Height {
		data.cacheSize.Height = data.MinSize.Height
	}
	if data.SizeHint.Height != layout.NoHint {
		data.cacheSize.Height = data.SizeHint.Height
	}
}
