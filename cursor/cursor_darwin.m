#include "cursor_darwin.h"

void *ArrowCursor() {
	return [NSCursor arrowCursor];
}

void *TextCursor() {
	return [NSCursor IBeamCursor];
}

void *VerticalTextCursor() {
	return [NSCursor IBeamCursorForVerticalLayout];
}

void *CrossHairCursor() {
	return [NSCursor crosshairCursor];
}

void *ClosedHandCursor() {
	return [NSCursor closedHandCursor];
}

void *OpenHandCursor() {
	return [NSCursor openHandCursor];
}

void *PointingHandCursor() {
	return [NSCursor pointingHandCursor];
}

void *ResizeLeftCursor() {
	return [NSCursor resizeLeftCursor];
}

void *ResizeRightCursor() {
	return [NSCursor resizeRightCursor];
}

void *ResizeLeftRightCursor() {
	return [NSCursor resizeLeftRightCursor];
}

void *ResizeUpCursor() {
	return [NSCursor resizeUpCursor];
}

void *ResizeDownCursor() {
	return [NSCursor resizeDownCursor];
}

void *ResizeUpDownCursor() {
	return [NSCursor resizeUpDownCursor];
}

void *DisappearingItemCursor() {
	return [NSCursor disappearingItemCursor];
}

void *NotAllowedCursor() {
	return [NSCursor operationNotAllowedCursor];
}

void *DragLinkCursor() {
	return [NSCursor dragLinkCursor];
}

void *DragCopyCursor() {
	return [NSCursor dragCopyCursor];
}

void *ContextMenuCursor() {
	return [NSCursor contextualMenuCursor];
}

void *NewCursor(void *img, float hotX, float hotY) {
		NSImage *nsimg = [[[NSImage alloc] initWithCGImage:img size:NSZeroSize] retain];
		return [[[NSCursor alloc] initWithImage:nsimg hotSpot:NSMakePoint(hotX,hotY)] retain];
}

void DisposeCursor(void *cursor) {
	[((NSCursor *)cursor) release];
}
