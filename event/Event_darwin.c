// Copyright (c) 2016 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

#include <stdlib.h>
#include <dispatch/dispatch.h>
#include "_cgo_export.h"

void invoke(uint64_t id) {
	uint64_t *data = malloc(sizeof(uint64_t));
	*data = id;
	dispatch_async_f(dispatch_get_main_queue(), data, (dispatch_function_t)dispatchInvocation);
}

void invokeAfter(uint64_t id, int64_t afterNanos) {
	uint64_t *data = malloc(sizeof(uint64_t));
	*data = id;
	dispatch_after_f(dispatch_time(DISPATCH_TIME_NOW, afterNanos), dispatch_get_main_queue(), data, (dispatch_function_t)dispatchInvocation);
}
