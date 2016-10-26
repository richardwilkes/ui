# go-ui
A user interface for Go.

C Dependencies:
- [Cairo](https://www.cairographics.org)
- [Pango](http://www.pango.org)
- X11 (linux only)

Go Dependencies:
```
go get -u github.com/richardwilkes/errs
go get -u github.com/richardwilkes/i18n
go get -u github.com/richardwilkes/xmath
```

This is very much a work in progress. My intent is to make this work for Mac, Linux & Windows.
At the moment, the Mac and Linux implementations are functional, if not complete.

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

To run the Demo, you'll need to [follow the directions in the images folder](https://github.com/richardwilkes/ui/blob/master/Demo/images/README.md) to build the image resources.
