## Deprecated. Future work moved to github.com/richardwilkes/ux

## A user interface for Go.

### C Dependencies:
- [Cairo](https://www.cairographics.org)
- [Pango](http://www.pango.org)
- X11 (linux only)

#### On macOS:
```
brew install cairo pango
```

#### On Linux:
```
sudo apt install pkg-config libcairo2-dev libpango1.0-dev libx11-dev libxcursor-dev
```

#### On Windows:
Install [MSYS2](http://www.msys2.org/) then install pkg-config and gtk3
```
pacman -S mingw64/mingw-w64-x86_64-pkg-config mingw64/mingw-w64-x86_64-gtk3
```

### Go Dependencies:
```
go get -u github.com/richardwilkes/toolbox
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
- [ ] Web View (only macOS implemented at the moment)

Top-level windows and dialogs:

- [ ] Dialog
- [ ] FileDialog
- [x] Window
