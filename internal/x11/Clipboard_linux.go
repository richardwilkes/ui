// Copyright (c) 2016 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

package x11

import (
	"fmt"
	"github.com/richardwilkes/ui/clipboard/datatypes"
	"unsafe"
	// #cgo pkg-config: x11
	// #include <X11/Xlib.h>
	// #include <X11/Xatom.h>
	"C"
)

var (
	clipboardWindow     Window
	acquiredTime        C.Time
	clipData            = make(map[string][]byte)
	clipDataChangeCount int
)

func initClipboard() {
	clipboardWindow = Window(C.XCreateWindow(display, C.Window(DefaultRootWindow()), 0, 0, 1, 1, 0, C.CopyFromParent, C.InputOnly, nil, 0, nil))
	clipboardWindow.SelectInput(PropertyNotifyType | SelectionClearType | SelectionRequestType | SelectionNotifyType)
}

func ownsClipboard() bool {
	return acquiredTime != 0 && C.XGetSelectionOwner(display, C.Atom(clipboardAtom)) == C.Window(clipboardWindow)
}

func ClipboardChangeCount() int {
	if !ownsClipboard() {
		// Someone else owns the clipboard, so increment the counter as we can't tell what the
		// real state is.
		clipDataChangeCount++
	}
	return clipDataChangeCount
}

func ClipboardClear() {
	clipData = make(map[string][]byte)
	clipDataChangeCount++
	claimSelectionOwnership()
}

func ClipboardTypes() []string {
	if ownsClipboard() {
		types := make([]string, len(clipData))
		i := 0
		for key := range clipData {
			types[i] = key
			i++
		}
		return types
	}

	clipboardWindow.DeleteProperty(clipboardAtom)
	C.XConvertSelection(display, C.Atom(clipboardAtom), C.Atom(targetsAtom), C.Atom(clipboardAtom), C.Window(clipboardWindow), lastEventTime)
	for {
		evt := clipboardWindow.NextEventOfType(SelectionNotifyType)
		if evt != nil {
			selEvt := evt.ToSelectionEvent()
			prop := selEvt.Property()
			if prop == C.None {
				return make([]string, 0)
			}
			actualType, format, count, data := clipboardWindow.Property(prop, C.AnyPropertyType)
			result := make([]string, 0, count)
			if data != nil {
				if actualType == C.XA_ATOM && format == 32 {
					names := make([]*C.char, count)
					C.XGetAtomNames(display, (*C.Atom)(data), C.int(count), &names[0])
					var seenStringType bool
					for i := 0; i < count; i++ {
						name := C.GoString(names[i])
						switch name {
						case "TIMESTAMP", "TARGETS", "MULTIPLE", "SAVE_TARGETS":
						case "UTF8_STRING", "COMPOUND_TEXT", "STRING", "TEXT":
							if !seenStringType {
								seenStringType = true
								result = append(result, datatypes.PlainText)
							}
						default:
							result = append(result, name)
						}
						C.XFree(unsafe.Pointer(names[i]))
					}
				}
				C.XFree(data)
			}
			clipboardWindow.DeleteProperty(prop)
			for _, t := range result {
				fmt.Println(t)
			}
			return result
		}
	}
}

func GetClipboard(dataType string) []byte {
	// RAW: Implement for Linux (i.e. cross-app support)
	return clipData[dataType]
}

func SetClipboard(data []datatypes.Data) {
	clipData = make(map[string][]byte)
	clipDataChangeCount++
	for _, one := range data {
		clipData[one.MimeType] = one.Bytes
	}
	claimSelectionOwnership()
}

func claimSelectionOwnership() {
	C.XSetSelectionOwner(display, C.Atom(clipboardAtom), C.Window(clipboardWindow), lastEventTime)
}

func ProcessSelectionClearEvent(evt *SelectionClearEvent) {
	fmt.Println("ProcessSelectionClearEvent")
	if evt.Window() == clipboardWindow {
		clipData = make(map[string][]byte)
		clipDataChangeCount++
		acquiredTime = 0
	}
}

func processSelectionNotifyEvent(evt *SelectionEvent) {
	fmt.Println("ProcessSelectionNotifyEvent")
}

func ProcessSelectionRequestEvent(evt *SelectionRequestEvent) {
	fmt.Println("ProcessSelectionRequestEvent")
	if evt.Owner() == clipboardWindow && evt.Selection() == clipboardAtom {
		when := evt.When()
		bad := evt.Property() == C.None || (when != C.CurrentTime && when < acquiredTime)
		if !bad {
			fmt.Printf("From %v\n", evt.Requestor())
			switch evt.Target() {
			case targetsAtom:
				// Send list of supported targets
				fmt.Println("got targetsAtom")
			case multipleAtom:
				// Multiple conversions were requested
				fmt.Println("got multipleAtom")
			default:
				// Convert data to requested format
				fmt.Printf("got atom %v\n", evt.Target())
			}
		}
		evt.Requestor().Send(NoEventMask, evt.NewNotify(bad))
	}
}

//    int i;
//    const Atom formats[] = { _glfw.x11.UTF8_STRING,
//                             _glfw.x11.COMPOUND_STRING,
//                             XA_STRING };
//    const int formatCount = sizeof(formats) / sizeof(formats[0]);
//
//    if (request->target == _glfw.x11.TARGETS)
//    {
//        // The list of supported targets was requested
//
//        const Atom targets[] = { _glfw.x11.TARGETS,
//                                 _glfw.x11.MULTIPLE,
//                                 _glfw.x11.UTF8_STRING,
//                                 _glfw.x11.COMPOUND_STRING,
//                                 XA_STRING };
//
//        XChangeProperty(_glfw.x11.display,
//                        request->requestor,
//                        request->property,
//                        XA_ATOM,
//                        32,
//                        PropModeReplace,
//                        (unsigned char*) targets,
//                        sizeof(targets) / sizeof(targets[0]));
//
//        return request->property;
//    }
//
//    if (request->target == _glfw.x11.MULTIPLE)
//    {
//        // Multiple conversions were requested
//
//        Atom* targets;
//        unsigned long i, count;
//
//        count = _glfwGetWindowProperty(request->requestor,
//                                       request->property,
//                                       _glfw.x11.ATOM_PAIR,
//                                       (unsigned char**) &targets);
//
//        for (i = 0;  i < count;  i += 2)
//        {
//            int j;
//
//            for (j = 0;  j < formatCount;  j++)
//            {
//                if (targets[i] == formats[j])
//                    break;
//            }
//
//            if (j < formatCount)
//            {
//                XChangeProperty(_glfw.x11.display,
//                                request->requestor,
//                                targets[i + 1],
//                                targets[i],
//                                8,
//                                PropModeReplace,
//                                (unsigned char*) _glfw.x11.selection.string,
//                                strlen(_glfw.x11.selection.string));
//            }
//            else
//                targets[i + 1] = None;
//        }
//
//        XChangeProperty(_glfw.x11.display,
//                        request->requestor,
//                        request->property,
//                        _glfw.x11.ATOM_PAIR,
//                        32,
//                        PropModeReplace,
//                        (unsigned char*) targets,
//                        count);
//
//        XFree(targets);
//
//        return request->property;
//    }
//
//    if (request->target == _glfw.x11.SAVE_TARGETS)
//    {
//        // The request is a check whether we support SAVE_TARGETS
//        // It should be handled as a no-op side effect target
//
//        XChangeProperty(_glfw.x11.display,
//                        request->requestor,
//                        request->property,
//                        XInternAtom(_glfw.x11.display, "NULL", False),
//                        32,
//                        PropModeReplace,
//                        NULL,
//                        0);
//
//        return request->property;
//    }
//
//    // Conversion to a data target was requested
//
//    for (i = 0;  i < formatCount;  i++)
//    {
//        if (request->target == formats[i])
//        {
//            // The requested target is one we support
//
//            XChangeProperty(_glfw.x11.display,
//                            request->requestor,
//                            request->property,
//                            request->target,
//                            8,
//                            PropModeReplace,
//                            (unsigned char*) _glfw.x11.selection.string,
//                            strlen(_glfw.x11.selection.string));
//
//            return request->property;
//        }
//    }
//
//    // The requested target is not supported
//
//    return None
