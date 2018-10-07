package cursor

import (
	"fmt"
	"unsafe"

	"github.com/richardwilkes/toolbox/xmath/geom"
	"github.com/richardwilkes/ui/color"
	"github.com/richardwilkes/ui/draw"

	// #cgo CFLAGS: -x objective-c
	// #cgo LDFLAGS: -framework Cocoa
	// #include <Cocoa/Cocoa.h>
	//
	// void *ArrowCursor() { return [NSCursor arrowCursor]; }
	// void *TextCursor() { return [NSCursor IBeamCursor]; }
	// void *VerticalTextCursor() { return [NSCursor IBeamCursorForVerticalLayout]; }
	// void *CrossHairCursor() { return [NSCursor crosshairCursor]; }
	// void *ClosedHandCursor() { return [NSCursor closedHandCursor]; }
	// void *OpenHandCursor() { return [NSCursor openHandCursor]; }
	// void *PointingHandCursor() { return [NSCursor pointingHandCursor]; }
	// void *ResizeLeftCursor() { return [NSCursor resizeLeftCursor]; }
	// void *ResizeRightCursor() { return [NSCursor resizeRightCursor]; }
	// void *ResizeLeftRightCursor() { return [NSCursor resizeLeftRightCursor]; }
	// void *ResizeUpCursor() { return [NSCursor resizeUpCursor]; }
	// void *ResizeDownCursor() { return [NSCursor resizeDownCursor]; }
	// void *ResizeUpDownCursor() { return [NSCursor resizeUpDownCursor]; }
	// void *DisappearingItemCursor() { return [NSCursor disappearingItemCursor]; }
	// void *NotAllowedCursor() { return [NSCursor operationNotAllowedCursor]; }
	// void *DragLinkCursor() { return [NSCursor dragLinkCursor]; }
	// void *DragCopyCursor() { return [NSCursor dragCopyCursor]; }
	// void *ContextMenuCursor() { return [NSCursor contextualMenuCursor]; }
	// void *NewCursor(void *img, float hotX, float hotY) {
	//		NSImage *nsimg = [[[NSImage alloc] initWithCGImage:img size:NSZeroSize] retain];
	// 		return [[[NSCursor alloc] initWithImage:nsimg hotSpot:NSMakePoint(hotX,hotY)] retain];
	// }
	// void DisposeCursor(void *cursor) { [((NSCursor *)cursor) release]; }
	"C"
)

func platformSystemCursor(id int) unsafe.Pointer {
	switch id {
	case arrow:
		return C.ArrowCursor()
	case text:
		return C.TextCursor()
	case verticalText:
		return C.VerticalTextCursor()
	case crossHair:
		return C.CrossHairCursor()
	case closedHand:
		return C.ClosedHandCursor()
	case openHand:
		return C.OpenHandCursor()
	case pointingHand:
		return C.PointingHandCursor()
	case resizeLeft:
		return C.ResizeLeftCursor()
	case resizeRight:
		return C.ResizeRightCursor()
	case resizeLeftRight:
		return C.ResizeLeftRightCursor()
	case resizeUp:
		return C.ResizeUpCursor()
	case resizeDown:
		return C.ResizeDownCursor()
	case resizeUpDown:
		return C.ResizeUpDownCursor()
	case disappearingItem:
		return C.DisappearingItemCursor()
	case notAllowed:
		return C.NotAllowedCursor()
	case dragLink:
		return C.DragLinkCursor()
	case dragCopy:
		return C.DragCopyCursor()
	case contextMenu:
		return C.ContextMenuCursor()
	default:
		panic(fmt.Sprintf("Invalid system cursor ID (%d)", id))
	}
}

func platformNewCursor(imgData *draw.ImageData, hotSpot geom.Point) unsafe.Pointer {
	colorspace := C.CGColorSpaceCreateWithName(C.kCGColorSpaceGenericRGB)
	defer C.CGColorSpaceRelease(colorspace)

	buffer := make([]color.Color, len(imgData.Pixels))
	for i, pixel := range imgData.Pixels {
		buffer[i] = pixel.Premultiply()
	}
	provider := C.CGDataProviderCreateWithData(nil, unsafe.Pointer(&buffer[0]), C.size_t(len(buffer)*4), nil)
	defer C.CGDataProviderRelease(provider)
	image := C.CGImageCreate(C.size_t(imgData.Width), C.size_t(imgData.Height), 8, 32, C.size_t(imgData.Width*4), colorspace, C.kCGBitmapByteOrder32Host|C.kCGImageAlphaPremultipliedFirst, provider, nil, false, C.kCGRenderingIntentDefault)

	return C.NewCursor(unsafe.Pointer(image), C.float(hotSpot.X), C.float(hotSpot.Y))
}

func platformDisposeCursor(cursor *Cursor) {
	C.DisposeCursor(cursor.cursor)
}
