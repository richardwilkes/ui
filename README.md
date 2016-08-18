# go-ui
A user interface for Go.

Dependencies:
```
go get -u github.com/richardwilkes/errs
go get -u github.com/richardwilkes/i18n
go get -u github.com/richardwilkes/xmath
```

This is very much a work in progress. My intent is to make this work for Mac, Linux & Windows.
At the moment, the Mac implementation works well. I am in the midst of implementing the Linux
version.

The Linux version requires that you have the X11 and Cairo development packages installed.

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
