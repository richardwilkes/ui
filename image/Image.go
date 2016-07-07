// Copyright (c) 2016 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

package image

import (
	"fmt"
	"github.com/richardwilkes/go-ui/color"
	"github.com/richardwilkes/go-ui/err"
	"github.com/richardwilkes/go-ui/geom"
	"net/http"
	"os"
	"sync"
	"unsafe"
)

type ref struct {
	img   *Image
	count int
}

type fsKey struct {
	fs   http.FileSystem
	path string
}

// Image represents a set of pixels that can be drawn to a graphics.Context.
type Image struct {
	id         int
	disabledID int
	size       geom.Size
	img        unsafe.Pointer
	key        interface{}
}

// Data is the raw information that makes up an Image.
type Data struct {
	Width  int
	Height int
	Pixels []color.Color
}

var (
	lock     sync.Mutex
	nextID   = 1
	registry = make(map[interface{}]*ref)
)

// AcquireFromFile attempts to load an image from the file system.
func AcquireFromFile(fs http.FileSystem, path string) (img *Image, e error) {
	lock.Lock()
	defer lock.Unlock()
	var r *ref
	var ok bool
	k := fsKey{fs: fs, path: path}
	if r, ok = registry[k]; !ok {
		var file http.File
		if file, e = fs.Open(path); e != nil {
			return nil, err.Wrap(e)
		}
		defer file.Close()
		var fi os.FileInfo
		if fi, e = file.Stat(); e != nil {
			return nil, err.Wrap(e)
		}
		size := int(fi.Size())
		buffer := make([]byte, size)
		var n int
		if n, e = file.Read(buffer); e != nil {
			return nil, err.Wrap(e)
		}
		if n != size {
			return nil, err.New(fmt.Sprintf("Read %d bytes from file, expected %d", n, size))
		}
		img := newImageFromBytes(buffer)
		if img == nil {
			return nil, err.New(fmt.Sprintf("Unable to load image from %s", path))
		}
		img.key = k
		img.id = nextID
		nextID++
		r = &ref{img: img}
		registry[img.key] = r
	}
	r.count++
	return r.img, nil
}

// AcquireFromURL attempts to load an image from a URL.
func AcquireFromURL(url string) (img *Image, e error) {
	lock.Lock()
	defer lock.Unlock()
	var r *ref
	var ok bool
	if r, ok = registry[url]; !ok {
		img := newImageFromURL(url)
		if img == nil {
			return nil, err.New(fmt.Sprintf("Unable to load image from %s", url))
		}
		img.key = url
		img.id = nextID
		nextID++
		r = &ref{img: img}
		registry[img.key] = r
	}
	r.count++
	return r.img, nil
}

// AcquireFromID attempts to find an already loaded image by its ID and return it. Returns nil
// if it cannot be found.
func AcquireFromID(id int) *Image {
	lock.Lock()
	defer lock.Unlock()
	if r, ok := registry[id]; ok {
		r.count++
		return r.img
	}
	return nil
}

// AcquireFromData creates a new image from the specified data.
func AcquireFromData(data *Data) (img *Image, e error) {
	img = newImageFromData(data)
	if img == nil {
		return nil, err.New("Unable to load image from data")
	}
	lock.Lock()
	defer lock.Unlock()
	img.id = nextID
	img.key = nextID
	nextID++
	r := &ref{img: img}
	registry[img.key] = r
	r.count++
	return img, nil
}

// AcquireBounds creates a new image from a region within this image.
func (img *Image) AcquireBounds(bounds geom.Rect) (image *Image, e error) {
	image = newImageFromImage(img, bounds)
	if image == nil {
		return nil, err.New("Unable to create image")
	}
	lock.Lock()
	defer lock.Unlock()
	image.id = nextID
	image.key = nextID
	nextID++
	r := &ref{img: image}
	registry[image.key] = r
	r.count++
	return image, nil
}

// AcquireDisabled returns an image based on this image which is desaturated and ghosted to
// represent a disabled state.
func (img *Image) AcquireDisabled() (image *Image, e error) {
	image = AcquireFromID(img.disabledID)
	if image != nil {
		return image, nil
	}
	data := img.Data()
	for i := range data.Pixels {
		p := data.Pixels[i]
		v := int((p.Luminance() * 255) + 0.5)
		data.Pixels[i] = color.RGBA(v, v, v, p.AlphaIntensity()*0.4)
	}
	if image, e = AcquireFromData(data); e == nil {
		img.disabledID = image.id
	}
	return image, e
}

// ID returns the underlying ID of the image.
func (img *Image) ID() int {
	return img.id
}

// PlatformPointer returns the underlying platform data structure for the image.
// Not intended for use outside of the github.com/richardwilkes/go-ui package and its descendants.
func (img *Image) PlatformPointer() unsafe.Pointer {
	return img.img
}

// Size returns the size of the image.
func (img *Image) Size() geom.Size {
	return img.size
}

// Release releases the image. If no other client is using the image, then the underlying OS
// resources for the image will be disposed of.
func (img *Image) Release() {
	lock.Lock()
	defer lock.Unlock()
	if ref, ok := registry[img.key]; ok {
		ref.count--
		if ref.count > 0 {
			return
		}
		delete(registry, img.key)
	}
	if img.img != nil {
		img.dispose()
	}
}
