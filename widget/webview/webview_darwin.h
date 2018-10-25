#include <stdlib.h>
#include <Cocoa/Cocoa.h>
#include <WebKit/Webkit.h>
#include <Security/Security.h>

typedef void *platformWebView;

platformWebView newWebView(void *wnd);
void disposeWebView(platformWebView webview);
void setWebViewFrame(platformWebView webview, double x, double y, double width, double height);
void loadWebViewURL(platformWebView webview, const char *url);
