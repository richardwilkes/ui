// Copyright (c) 2016 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

#include <stdlib.h>
#include "_cgo_export.h"
#include "Clipboard.h"

int platformClipboardChangeCount() {
	// RAW: Implement platformClipboardChangeCount for Linux
	return 0;
}

void platformClearClipboard() {
	// RAW: Implement platformClearClipboard for Linux
}

const char **platformClipboardTypes() {
	// RAW: Implement platformClipboardTypes for Linux
	const char **result = malloc(sizeof(char *));
	result[0] = NULL;
	return result;
}

struct platformClipboardData platformClipboardData(char *type) {
	// RAW: Implement platformClipboardData for Linux
	struct platformClipboardData d;
	d.count = 0;
	d.data = NULL;
	return d;
}

void platformSetClipboardData(char *type, int size, void *bytes) {
	// RAW: Implement platformSetClipboardData for Linux
}
