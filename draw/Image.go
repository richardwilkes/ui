// Copyright (c) 2016 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

package draw

import (
	"fmt"
	"github.com/richardwilkes/errs"
	"github.com/richardwilkes/ui/color"
	"github.com/richardwilkes/ui/geom"
	"net/http"
	"os"
	"sync"
	"unsafe"
)

type imgRef struct {
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

// ImageData is the raw information that makes up an Image.
type ImageData struct {
	Width  int
	Height int
	Pixels []color.Color
}

var (
	imageRegistryLock sync.Mutex
	nextImageID       = 1
	imageRegistry     = make(map[interface{}]*imgRef)
)

// AcquireImageFromFile attempts to load an image from the file system.
func AcquireImageFromFile(fs http.FileSystem, path string) (img *Image, e error) {
	imageRegistryLock.Lock()
	defer imageRegistryLock.Unlock()
	var r *imgRef
	var ok bool
	k := fsKey{fs: fs, path: path}
	if r, ok = imageRegistry[k]; !ok {
		var file http.File
		if file, e = fs.Open(path); e != nil {
			return nil, errs.Wrap(e)
		}
		defer file.Close()
		var fi os.FileInfo
		if fi, e = file.Stat(); e != nil {
			return nil, errs.Wrap(e)
		}
		size := int(fi.Size())
		buffer := make([]byte, size)
		var n int
		if n, e = file.Read(buffer); e != nil {
			return nil, errs.Wrap(e)
		}
		if n != size {
			return nil, errs.New(fmt.Sprintf("Read %d bytes from file, expected %d", n, size))
		}
		img := platformNewImageFromBytes(buffer)
		if img == nil {
			return nil, errs.New(fmt.Sprintf("Unable to load image from %s", path))
		}
		img.key = k
		img.id = nextImageID
		nextImageID++
		r = &imgRef{img: img}
		imageRegistry[img.key] = r
	}
	r.count++
	return r.img, nil
}

// AcquireImageFromURL attempts to load an image from a URL.
func AcquireImageFromURL(url string) (img *Image, e error) {
	imageRegistryLock.Lock()
	defer imageRegistryLock.Unlock()
	var r *imgRef
	var ok bool
	if r, ok = imageRegistry[url]; !ok {
		img := platformNewImageFromURL(url)
		if img == nil {
			return nil, errs.New(fmt.Sprintf("Unable to load image from %s", url))
		}
		img.key = url
		img.id = nextImageID
		nextImageID++
		r = &imgRef{img: img}
		imageRegistry[img.key] = r
	}
	r.count++
	return r.img, nil
}

// AcquireImageFromID attempts to find an already loaded image by its ID and return it. Returns nil
// if it cannot be found.
func AcquireImageFromID(id int) *Image {
	imageRegistryLock.Lock()
	defer imageRegistryLock.Unlock()
	if r, ok := imageRegistry[id]; ok {
		r.count++
		return r.img
	}
	return nil
}

// AcquireImageFromData creates a new image from the specified data.
func AcquireImageFromData(data *ImageData) (img *Image, e error) {
	img = platformNewImageFromData(data)
	if img == nil {
		return nil, errs.New("Unable to load image from data")
	}
	imageRegistryLock.Lock()
	defer imageRegistryLock.Unlock()
	img.id = nextImageID
	img.key = nextImageID
	nextImageID++
	r := &imgRef{img: img}
	imageRegistry[img.key] = r
	r.count++
	return img, nil
}

// AcquireImageBounds creates a new image from a region within this image.
func (img *Image) AcquireImageBounds(bounds geom.Rect) (image *Image, e error) {
	image = platformNewImageFromImage(img, bounds)
	if image == nil {
		return nil, errs.New("Unable to create image")
	}
	imageRegistryLock.Lock()
	defer imageRegistryLock.Unlock()
	image.id = nextImageID
	image.key = nextImageID
	nextImageID++
	r := &imgRef{img: image}
	imageRegistry[image.key] = r
	r.count++
	return image, nil
}

// AcquireDisabled returns an image based on this image which is desaturated and ghosted to
// represent a disabled state.
func (img *Image) AcquireDisabled() (image *Image, e error) {
	image = AcquireImageFromID(img.disabledID)
	if image != nil {
		return image, nil
	}
	data := img.Data()
	for i := range data.Pixels {
		p := data.Pixels[i]
		v := int((p.Luminance() * 255) + 0.5)
		data.Pixels[i] = color.RGBA(v, v, v, p.AlphaIntensity()*0.4)
	}
	if image, e = AcquireImageFromData(data); e == nil {
		img.disabledID = image.id
	}
	return image, e
}

// ID returns the underlying ID of the image.
func (img *Image) ID() int {
	return img.id
}

// Size returns the size of the image.
func (img *Image) Size() geom.Size {
	return img.size
}

// Data extracts the raw image data.
func (img *Image) Data() *ImageData {
	return img.platformData()
}

// PlatformPtr returns a pointer to the underlying platform-specific data.
func (img *Image) PlatformPtr() unsafe.Pointer {
	return img.img
}

// Release releases the image. If no other client is using the image, then the underlying OS
// resources for the image will be disposed of.
func (img *Image) Release() {
	imageRegistryLock.Lock()
	defer imageRegistryLock.Unlock()
	if ref, ok := imageRegistry[img.key]; ok {
		ref.count--
		if ref.count > 0 {
			return
		}
		delete(imageRegistry, img.key)
	}
	if img.img != nil {
		img.platformDispose()
	}
}
