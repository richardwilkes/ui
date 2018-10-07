package x11

import (
	"unsafe"

	"github.com/richardwilkes/ui/clipboard/datatypes"

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

func localClipboardTypes() []string {
	types := make([]string, len(clipData))
	i := 0
	for key := range clipData {
		types[i] = key
		i++
	}
	return types
}

func ClipboardTypes() []string {
	if ownsClipboard() {
		return localClipboardTypes()
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
					var seenRTFType bool
					for i := 0; i < count; i++ {
						name := C.GoString(names[i])
						switch name {
						case "TIMESTAMP", "TARGETS", "MULTIPLE", "SAVE_TARGETS":
						case datatypes.PlainText, "UTF8_STRING", "COMPOUND_TEXT", "STRING", "TEXT":
							if !seenStringType {
								seenStringType = true
								result = append(result, datatypes.PlainText)
							}
						case datatypes.RTFText, "TEXT/RTF", "application/rtf":
							if !seenRTFType {
								seenRTFType = true
								result = append(result, datatypes.RTFText)
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
			return result
		}
	}
}

func GetClipboard(dataType string) []byte {
	if ownsClipboard() {
		return clipData[dataType]
	}
	clipboardWindow.DeleteProperty(clipboardAtom)
	C.XConvertSelection(display, C.Atom(clipboardAtom), C.Atom(InternAtom(dataType)), C.Atom(clipboardAtom), C.Window(clipboardWindow), lastEventTime)
	for {
		evt := clipboardWindow.NextEventOfType(SelectionNotifyType)
		if evt != nil {
			selEvt := evt.ToSelectionEvent()
			prop := selEvt.Property()
			if prop == C.None {
				switch dataType {
				case datatypes.PlainText:
					return GetClipboard("UTF8_STRING")
				case "UTF8_STRING":
					return GetClipboard("STRING")
				default:
					return nil
				}
			}
			_, format, count, data := clipboardWindow.Property(prop, C.AnyPropertyType)
			length := count * format / 8
			result := make([]byte, length)
			if data != nil {
				raw := (*[1 << 30]byte)(data)
				copy(result, raw[:length])
				C.XFree(data)
			}
			clipboardWindow.DeleteProperty(prop)
			return result
		}
	}
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
	if evt.Window() == clipboardWindow {
		clipData = make(map[string][]byte)
		clipDataChangeCount++
		acquiredTime = 0
	}
}

func ProcessSelectionRequestEvent(evt *SelectionRequestEvent) {
	if evt.Owner() == clipboardWindow && evt.Selection() == clipboardAtom {
		when := evt.When()
		prop := evt.Property()
		bad := prop == C.None || (when != C.CurrentTime && when < acquiredTime)
		if !bad {
			target := evt.Target()
			switch target {
			case targetsAtom:
				// Send list of supported targets
				available := localClipboardTypes()
				atoms := make([]Atom, len(available)+1)
				atoms[0] = targetsAtom
				for i, one := range available {
					atoms[i+1] = InternAtom(one)
				}
				evt.Requestor().ChangeProperty(prop, C.XA_ATOM, 32, PropModeReplace, unsafe.Pointer(&atoms[0]), len(atoms))
			case multipleAtom:
				// Multiple conversions were requested
				// Not supported
				bad = true
			case saveTargetsAtom:
				// Save Targets requested
				// Not supported
				bad = true
			default:
				// Convert data to requested format
				requested := target.Name()
				var adjustedRequest string
				switch requested {
				case "UTF8_STRING", "COMPOUND_TEXT", "STRING", "TEXT":
					adjustedRequest = datatypes.PlainText
				case "TEXT/RTF", "application/rtf":
					adjustedRequest = datatypes.RTFText
				default:
					adjustedRequest = requested
				}
				bad = true
				for _, one := range localClipboardTypes() {
					if one == adjustedRequest {
						bad = false
						bytes := clipData[one]
						evt.Requestor().ChangeProperty(prop, target, 8, PropModeReplace, unsafe.Pointer(&bytes[0]), len(bytes))
						break
					}
				}
			}
		}
		evt.Requestor().Send(NoEventMask, evt.NewNotify(bad))
	}
}

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
