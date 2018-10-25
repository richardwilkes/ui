#include "display_darwin.h"

void getMainDisplayBounds(double *x, double *y, double *width, double *height) {
	NSRect bounds = [[NSScreen mainScreen] visibleFrame];
	*x = bounds.origin.x;
	*y = bounds.origin.y;
	*width = bounds.size.width;
	*height = bounds.size.height;
}
