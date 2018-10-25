#include "webview_darwin.h"


@interface nav : NSObject<WKNavigationDelegate>
@end

@implementation nav

- (void)webView:(WKWebView *)webView didFailNavigation:(WKNavigation *)navigation withError:(NSError *)error {
	// RAW: Provide a way to return this to the Go code
	printf("didFailNavigation: %s\n", [[error domain] UTF8String]);
}

- (void)webView:(WKWebView *)webView didFailProvisionalNavigation:(WKNavigation *)navigation withError:(NSError *)error {
	// RAW: Provide a way to return this to the Go code
	printf("didFailProvisionalNavigation: %s\n", [[error localizedDescription] UTF8String]);
}

- (void)webView:(WKWebView *)webView decidePolicyForNavigationAction:(WKNavigationAction *)navigationAction decisionHandler:(void (^)(WKNavigationActionPolicy))decisionHandler {
	decisionHandler(WKNavigationActionPolicyAllow);
}

- (void)webView:(WKWebView *)webView decidePolicyForNavigationResponse:(WKNavigationResponse *)navigationResponse decisionHandler:(void (^)(WKNavigationResponsePolicy))decisionHandler {
	decisionHandler(WKNavigationResponsePolicyAllow);
}

- (void)webView:(WKWebView *)webView didReceiveAuthenticationChallenge:(NSURLAuthenticationChallenge *)challenge completionHandler:(void (^)(NSURLSessionAuthChallengeDisposition disposition, NSURLCredential *credential))completionHandler {
	// Allow everything
	SecTrustRef serverTrust = challenge.protectionSpace.serverTrust;
	CFDataRef exceptions = SecTrustCopyExceptions(serverTrust);
	SecTrustSetExceptions(serverTrust, exceptions);
	CFRelease(exceptions);
	completionHandler(NSURLSessionAuthChallengeUseCredential, [NSURLCredential credentialForTrust:serverTrust]);
}

@end

platformWebView newWebView(void *wnd) {
	WKWebView *view = [[WKWebView alloc] initWithFrame:NSMakeRect(0, 0, 0, 0) configuration:[WKWebViewConfiguration new]];
	[view setNavigationDelegate:[nav new]];
	NSWindow *window = (NSWindow *)wnd;
	[[window contentView] addSubview:view];
	return (platformWebView)view;
}

void disposeWebView(platformWebView webview) {
	WKWebView *w = (WKWebView *)webview;
	[w removeFromSuperview];
	[w release];
}

void setWebViewFrame(platformWebView webview, double x, double y, double width, double height) {
	WKWebView *w = (WKWebView *)webview;
	[w setFrame:NSMakeRect(x, y, width, height)];
	[w setNeedsLayout:YES];
	[w setNeedsDisplay:YES];
}

void loadWebViewURL(platformWebView webview, const char *url) {
	[((WKWebView *)webview) loadRequest:[NSURLRequest requestWithURL:[NSURL URLWithString:[NSString stringWithUTF8String:url]]]];
}
