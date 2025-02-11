package http

import (
	"errors"
	"io"
	"mime/multipart"
	"os"

	_file "github.com/nortoo/utils-go/file"
)

func SaveFile(file *multipart.FileHeader, dst, srcMd5 string) error {
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, src)
	if err != nil {
		return err
	}

	if srcMd5 != "" {
		f, err := os.Open(dst)
		if err != nil {
			return errors.New("failed to verify file")
		}
		defer f.Close()
		if newMd5, err := _file.MD5(f); err != nil || newMd5 != srcMd5 {
			return errors.New("failed to verify file")
		}
	}

	return nil
}
