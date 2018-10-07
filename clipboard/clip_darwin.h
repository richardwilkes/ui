#include <Cocoa/Cocoa.h>

struct clipboardData {
	int count;
	const void *data;
};

int clipboardChangeCount();
void clearClipboard();
const char **clipboardTypes();
struct clipboardData getClipboardData(char *type);
void setClipboardData(char *type, int size, void *bytes);
