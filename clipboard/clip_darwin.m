#include "clip_darwin.h"

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

struct clipboardData getClipboardData(char *type) {
	struct clipboardData d;
	NSData *nsd = [[NSPasteboard generalPasteboard] dataForType:[NSString stringWithUTF8String:type]];
	d.count = [nsd length];
	d.data = [nsd bytes];
	return d;
}

void setClipboardData(char *type, int size, void *bytes) {
	[[NSPasteboard generalPasteboard] setData:[NSData dataWithBytes:bytes length:size] forType:[NSString stringWithUTF8String:type]];
}
