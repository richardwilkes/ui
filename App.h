// Copyright (c) 2016 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

#ifndef __RW_GOUI_APP__
#define __RW_GOUI_APP__

const char *platformAppName();
void platformStart();
void platformAttemptTerminate();
void platformAppMayTerminateNow(int terminate);
void platformHideApp();
void platformHideOtherApps();
void platformShowAllApps();

#endif // __RW_GOUI_APP__
