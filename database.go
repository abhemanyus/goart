package goart

import (
	"fmt"
	"io/fs"
	"regexp"
)

const ImgRegex = "(?i).(png|jpe?g|webp)$"

type ImageList []string

type ImageDatabase struct {
	images ImageList
	root   fs.FS
}

func GetImages(f fs.FS) (ImageList, error) {
	imgReg := regexp.MustCompile(ImgRegex)
	var images ImageList
	addImage := func(p string, d fs.DirEntry, err error) error {
		if d.IsDir() && err != nil {
			return fs.SkipDir
		}
		if !d.IsDir() && imgReg.MatchString(d.Name()) {
			images = append(images, p)
		}
		return nil
	}
	err := fs.WalkDir(f, ".", addImage)
	return images, err
}

func CreateImageStore(f fs.FS) (*ImageDatabase, error) {
	images, err := GetImages(f)
	if err != nil {
		return nil, fmt.Errorf("can't create image store %w", err)
	}
	return &ImageDatabase{images, f}, nil
}

func (db *ImageDatabase) GetImages(limit, offset int) ImageList {
	var images ImageList
	for i := offset; i < len(db.images) && i < limit+offset; i++ {
		images = append(images, db.images[i])
	}
	return images
}
