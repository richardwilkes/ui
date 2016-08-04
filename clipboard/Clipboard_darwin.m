// Copyright (c) 2016 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

#include <Cocoa/Cocoa.h>
#include <stdlib.h>
#include "_cgo_export.h"
#include "Clipboard.h"

int clipboardChangeCount() {
	return [[NSPasteboard generalPasteboard] changeCount];
}

void clearClipboard() {
	[[NSPasteboard generalPasteboard] clearContents];
}

const char **clipboardTypes() {
	NSArray<NSString *> *types = [[NSPasteboard generalPasteboard] types];
	NSUInteger count = [types count];
	const char **result = malloc(sizeof(char *) * (count + 1));
	result[count] = NULL;
	for (int i = 0; i < count; i++) {
		result[i] = [[types objectAtIndex:i] UTF8String];
	}
	return result;
}

struct data clipboardData(char *type) {
	struct data d;
	NSData *nsd = [[NSPasteboard generalPasteboard] dataForType:[NSString stringWithUTF8String:type]];
	d.count = [nsd length];
	d.data = [nsd bytes];
	return d;
}

void setClipboardData(char *type, int size, void *bytes) {
	[[NSPasteboard generalPasteboard] setData:[NSData dataWithBytes:bytes length:size] forType:[NSString stringWithUTF8String:type]];
}
