# go-ui
A user interface for Go.

C Dependencies:
```
Cairo
Pango
X11 (linux only)
```

Go Dependencies:
```
go get -u github.com/richardwilkes/errs
go get -u github.com/richardwilkes/i18n
go get -u github.com/richardwilkes/xmath
```

This is very much a work in progress. My intent is to make this work for Mac, Linux & Windows.
At the moment, the Mac implementation works well. I am in the midst of implementing the Linux
version.

Widgets that have been implemented:

- [x] Button
- [x] CheckBox
- [x] ImageButton
- [x] Label
- [x] List
- [x] Menus
- [x] PopupMenu
- [ ] ProgressBar
- [x] RadioButton
- [x] ScrollArea
- [x] ScrollBar
- [ ] Slider
- [x] Separator
- [ ] SplitPanel
- [ ] Table
- [ ] TabPanel
- [ ] TextArea
- [x] TextField
- [ ] ToolBar
- [ ] Tree

Top-level windows and dialogs:

- [ ] Dialog
- [ ] FileDialog
- [x] Window
