package goart_test

import (
	"regexp"
	"testing"
	"testing/fstest"

	"github.com/abhemanyus/goart"
)

var (
	fs = fstest.MapFS{
		"one.png":        {},
		"safe/one.png":   {},
		"unsafe/one.jpg": {},
		"two.png":        {},
		"two.txt":        {},
	}
)

func TestImageStore(t *testing.T) {
	store, err := goart.CreateImageStore(fs)
	assertError(t, err, nil)
	// returns only 1 image out of 4 when offset 3 three, regardless of limit
	assertLength(t, 1, len(store.GetImages(2, 3)))
}

func TestGetImages(t *testing.T) {
	images, _ := goart.GetImages(fs)
	t.Run("number of images read", func(t *testing.T) {
		assertLength(t, 4, len(images))
	})

	t.Run("entries are images", func(t *testing.T) {
		// supported formats are png, webp, jpeg
		imgReg := regexp.MustCompile(goart.ImgRegex)
		for _, file := range images {
			if !imgReg.MatchString(file) {
				t.Errorf("%q is not a supported image", file)
			}
		}
	})

	t.Run("images are accessible", func(t *testing.T) {
		// images aren't blocked by perms
		for _, file := range images {
			_, err := fs.Stat(file)
			assertError(t, err, nil)
		}
	})
}
