package clipboard

import (
	"unsafe"

	"github.com/richardwilkes/ui/clipboard/datatypes"

	// #cgo CFLAGS: -x objective-c
	// #cgo LDFLAGS: -framework Cocoa
	// #include "clip_darwin.h"
	"C"
)

const (
	utf8 = "public.utf8-plain-text"
)

func platformChangeCount() int {
	return int(C.clipboardChangeCount())
}

func platformClear() {
	C.clearClipboard()
}

func platformTypes() []string {
	set := make(map[string]bool)
	clipTypes := (*[1 << 30]*C.char)(unsafe.Pointer(C.clipboardTypes()))
	i := 0
	for clipTypes[i] != nil {
		i++
	}
	types := make([]string, 0, i)
	for j := 0; j < i; j++ {
		str := convertUTItoMimeType(C.GoString(clipTypes[j]))
		if str != "" && !set[str] {
			set[str] = true
			types = append(types, str)
		}
	}
	C.free(unsafe.Pointer(clipTypes))
	return types
}

func platformGetData(dataType string) []byte {
	cstr := C.CString(convertMimeTypeToUTI(dataType))
	data := C.getClipboardData(cstr)
	C.free(unsafe.Pointer(cstr))
	count := int(data.count)
	result := make([]byte, count)
	bytes := (*[1 << 30]byte)(data.data)
	for i := 0; i < count; i++ {
		result[i] = bytes[i]
	}
	return result
}

func platformSetData(data []datatypes.Data) {
	C.clearClipboard()
	for _, one := range data {
		cstr := C.CString(convertMimeTypeToUTI(one.MimeType))
		C.setClipboardData(cstr, C.int(len(one.Bytes)), unsafe.Pointer(&one.Bytes[0]))
		C.free(unsafe.Pointer(cstr))
	}
}

func convertMimeTypeToUTI(mimeType string) string {
	if mimeType == datatypes.PlainText {
		return utf8
	}
	cfstr := newCFStringRef(mimeType)
	uti := C.UTTypeCreatePreferredIdentifierForTag(C.kUTTagClassMIMEType, cfstr, C.CFStringRef(C.NULL))
	disposeCFStringRef(cfstr)
	if uti == C.CFStringRef(C.NULL) {
		return ""
	}
	result := cfStringRefToString(uti)
	disposeCFStringRef(uti)
	return result
}

func convertUTItoMimeType(uti string) string {
	if uti == utf8 {
		return datatypes.PlainText
	}
	cfstr := newCFStringRef(uti)
	mimeType := C.UTTypeCopyPreferredTagWithClass(cfstr, C.kUTTagClassMIMEType)
	disposeCFStringRef(cfstr)
	if mimeType == C.CFStringRef(C.NULL) {
		return ""
	}
	result := cfStringRefToString(mimeType)
	disposeCFStringRef(mimeType)
	return result
}

func newCFStringRef(str string) C.CFStringRef {
	cstr := C.CString(str)
	cfstr := C.CFStringCreateWithCString(C.CFAllocatorRef(C.NULL), cstr, C.kCFStringEncodingUTF8)
	C.free(unsafe.Pointer(cstr))
	return cfstr
}

func disposeCFStringRef(str C.CFStringRef) {
	C.CFRelease(C.CFTypeRef(str))
}

func cfStringRefToString(str C.CFStringRef) string {
	length := C.CFStringGetLength(str) * 2
	buffer := C.malloc(C.size_t(length))
	C.CFStringGetCString(str, (*C.char)(buffer), length, C.kCFStringEncodingUTF8)
	result := C.GoString((*C.char)(buffer))
	C.free(buffer)
	return result
}
