#include <stdlib.h>
#include <Cocoa/Cocoa.h>
#include <Quartz/Quartz.h>
#include <cairo.h>
#include <cairo-quartz.h>
#include <dispatch/dispatch.h>

typedef void *platformWindow;

platformWindow getKeyWindow();
void bringAllWindowsToFront();
void hideCursorUntilMouseMoves();
void closeWindow(platformWindow window);
const char *getWindowTitle(platformWindow window);
void setWindowTitle(platformWindow window, const char *title);
void bringWindowToFront(platformWindow window);
void repaintWindow(platformWindow window, double x, double y, double width, double height);
void flushPainting(platformWindow window);
void minimizeWindow(platformWindow window);
void zoomWindow(platformWindow window);
void setCursor(platformWindow window, void *cursor);
void invoke(unsigned long id);
void invokeAfter(unsigned long id, long afterNanos);
platformWindow newWindow(double x, double y, double width, double height, int styleMask);
void getWindowFrame(platformWindow window, double *x, double *y, double *width, double *height);
void setWindowFrame(platformWindow window, double x, double y, double width, double height);
void getWindowContentFrame(platformWindow window, double *x, double *y, double *width, double *height);
