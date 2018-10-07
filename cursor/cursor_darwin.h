#include <Cocoa/Cocoa.h>

void *ArrowCursor();
void *TextCursor();
void *VerticalTextCursor();
void *CrossHairCursor();
void *ClosedHandCursor();
void *OpenHandCursor();
void *PointingHandCursor();
void *ResizeLeftCursor();
void *ResizeRightCursor();
void *ResizeLeftRightCursor();
void *ResizeUpCursor();
void *ResizeDownCursor();
void *ResizeUpDownCursor();
void *DisappearingItemCursor();
void *NotAllowedCursor();
void *DragLinkCursor();
void *DragCopyCursor();
void *ContextMenuCursor();
void *NewCursor(void *img, float hotX, float hotY);
void DisposeCursor(void *cursor);
