// Copyright (c) 2016 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

package ui

import (
	"bytes"
	"runtime"
	"strconv"
	"strings"
)

// Error holds the detailed error message.
type Error struct {
	message string
	stack   []uintptr
	cause   error
}

var (
	sigpanic *runtime.Func
	// DetailedError should be set to true to enable detailed output that contains a stack trace as
	// well as chained errors. Defaults to true.
	DetailedError = true
	// TrimRuntime should be set to true to enable removal of code from the stack trace that is in
	// one of the golang the runtime packages. Defaults to true.
	TrimRuntime = true
)

func init() {
	var bad *int
	func() int {
		defer func() {
			if r := recover(); r != nil {
				var pcs [512]uintptr
				n := runtime.Callers(2, pcs[:])
				for _, pc := range pcs[:n] {
					f := runtime.FuncForPC(pc)
					if f.Name() == "runtime.sigpanic" {
						sigpanic = f
						break
					}
				}
			}
		}()
		return *bad // trigger a nil pointer dereference
	}()
}

// WrapError wraps an error and turns it into a detailed error. If error is already a detailed
// error, then it will be returned as-is.
func WrapError(cause error) error {
	if cause == nil {
		return nil
	}
	if err, ok := cause.(*Error); ok {
		return err
	}
	return &Error{message: cause.Error(), stack: callStack()}
}

// NewError creates a new detailed error with the specified message.
func NewError(message string) *Error {
	return &Error{message: message, stack: callStack()}
}

// NewErrorWithCause creates a new detailed error with the specified message and underlying cause.
func NewErrorWithCause(message string, cause error) *Error {
	return &Error{message: message, stack: callStack(), cause: cause}
}

func callStack() []uintptr {
	var pcs [512]uintptr
	n := runtime.Callers(3, pcs[:])
	cs := make([]uintptr, n)
	copy(cs, pcs[:n])
	return cs
}

// Message returns the message attached to this error.
func (err *Error) Message() string {
	return err.message
}

// Error implements the error interface.
func (err *Error) Error() string {
	if DetailedError {
		return err.Detail()
	}
	return err.message
}

// Detail returns the fully detailed error message, which includes the primary message, the call
// stack, and potentially one or more chained causes.
func (err *Error) Detail() string {
	var buffer bytes.Buffer
	err.detail(&buffer)
	return buffer.String()
}

func (err *Error) detail(buffer *bytes.Buffer) {
	var fn *runtime.Func
	buffer.WriteString(err.message)
	for _, pc := range err.stack {
		if fn != sigpanic {
			pc--
		}
		fn = runtime.FuncForPC(pc)
		if fn != nil {
			name := fn.Name()
			if TrimRuntime && strings.HasPrefix(name, "runtime.") {
				continue
			}
			buffer.WriteString("\n    [")
			buffer.WriteString(fn.Name())
			buffer.WriteString("] ")
			file, line := fn.FileLine(pc)
			buffer.WriteString(file)
			buffer.WriteByte(':')
			buffer.WriteString(strconv.Itoa(line))
		} else {
			buffer.WriteString("\n    [unknown]")
		}
	}
	if err.cause != nil {
		buffer.WriteString("\n  Caused by ")
		if detailed, ok := err.cause.(*Error); ok {
			detailed.detail(buffer)
		} else {
			buffer.WriteString(err.cause.Error())
		}
	}
}
