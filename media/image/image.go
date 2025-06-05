package image

import (
	"bytes"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"

	_file "github.com/nortoo/utils-go/file"
	"github.com/nortoo/utils-go/types"

	"github.com/disintegration/imaging"
	. "github.com/nortoo/utils-go/constant"
	"github.com/pkg/errors"
)

// IsSupportedImage returns whether the file is an image file.
func IsSupportedImage(filename string) bool {
	return types.Contains(supportedImageFormats, MediaType(strings.ToLower(path.Ext(filename))))
}

// Dimensions return the image's dimensions.
func Dimensions(file string) (width, height int, err error) {
	imgData, err := os.ReadFile(file)
	if err != nil {
		return 0, 0, err
	}
	buf := bytes.NewBuffer(imgData)
	image, err := imaging.Decode(buf)
	if err != nil {
		return 0, 0, err
	}

	b := image.Bounds()
	return b.Max.X, b.Max.Y, nil
}

// Resize the image's dimensions.
// If one of width or height is 0, the image aspect ratio is preserved.
func Resize(file, dstDir string, width, height, imageQuality int) error {
	imgData, err := os.ReadFile(file)
	if err != nil {
		return err
	}
	buf := bytes.NewBuffer(imgData)
	image, err := imaging.Decode(buf)
	if err != nil {
		return err
	}

	filename := path.Base(file)
	dst := path.Join(dstDir, filename)
	b := image.Bounds()
	if b.Max.X < width {
		if !_file.Exists(dst) {
			err = os.MkdirAll(dstDir, 0755)
			if err != nil {
				return err
			}
		}
		return imaging.Save(image, dst, imaging.JPEGQuality(imageQuality))
	}

	image = imaging.Resize(image, width, height, imaging.Lanczos)
	if !_file.Exists(dst) {
		err = os.MkdirAll(dstDir, 0755)
		if err != nil {
			return err
		}
	}
	return imaging.Save(image, dst, imaging.JPEGQuality(imageQuality))
}

// Convert an image to another format.
func Convert(file, dstDir string, format MediaType) error {
	targetFormat := format.String()

	if _file.IsDir(file) {
		return errors.Errorf("file is not an image or is not supported: %s", file)
	}
	if !IsSupportedImage(file) {
		return errors.Errorf("file format [%s] is not supported", strings.ToLower(filepath.Ext(file)))
	}
	if !IsSupportedImage(targetFormat) {
		return errors.Errorf("target format [%s] is not supported", targetFormat)
	}

	imageFormat, err := imaging.FormatFromExtension(targetFormat)
	if err != nil {
		return err
	}

	imgData, err := os.ReadFile(file)
	if err != nil {
		return err
	}
	buf := bytes.NewBuffer(imgData)
	image, err := imaging.Decode(buf)
	if err != nil {
		return err
	}

	if !_file.Exists(dstDir) {
		err = os.MkdirAll(dstDir, 0755)
		if err != nil {
			return err
		}
	}
	filename := filepath.Base(file)
	newName := fmt.Sprintf("%s%s", strings.TrimSuffix(filename, filepath.Ext(filename)), targetFormat)
	dst := filepath.Join(dstDir, newName)
	f, err := os.OpenFile(dst, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	return imaging.Encode(f, image, imageFormat)
}

// ResizeByWalkingIntoDir accepts a directory(src), resizes all images inside the directory(src), and saves them to another one(dst).
func ResizeByWalkingIntoDir(src, dst string, width, height, imageQuality int) error {
	return filepath.Walk(src, func(p string, i os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if i.IsDir() || !IsSupportedImage(i.Name()) {
			return nil
		}
		d := filepath.Join(dst, strings.TrimSuffix(strings.TrimPrefix(p, src), i.Name()))
		return Resize(p, d, width, height, imageQuality)
	})
}
