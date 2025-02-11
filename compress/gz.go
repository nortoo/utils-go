package compress

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"path"
	"path/filepath"
	"strings"
)

func Gzip(src, dst string) error {
	src = strings.TrimPrefix(src, "./")
	d, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer d.Close()

	gw := gzip.NewWriter(d)
	defer gw.Close()

	tw := tar.NewWriter(gw)
	defer tw.Close()

	return filepath.Walk(src, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		var link string
		if info.Mode()&os.ModeSymlink == os.ModeSymlink {
			if link, err = os.Readlink(path); err != nil {
				return err
			}
		}

		fh, err := tar.FileInfoHeader(info, link)
		if err != nil {
			return err
		}

		fh.Name = strings.TrimPrefix(path, string(filepath.Separator))
		if info.IsDir() {
			fh.Name += "/"
		}
		err = tw.WriteHeader(fh)
		if err != nil {
			return err
		}

		if !info.Mode().IsRegular() {
			return nil
		}

		fr, err := os.Open(path)
		if err != nil {
			return err
		}
		defer fr.Close()

		_, err = io.Copy(tw, fr)
		return err
	})

}

func UnGzip(src, dst string) error {
	srcFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	if dst != "" {
		if err := os.MkdirAll(dst, 0755); err != nil {
			return err
		}
	}

	gr, err := gzip.NewReader(srcFile)
	if err != nil {
		return err
	}
	defer gr.Close()
	tr := tar.NewReader(gr)
	for {
		hdr, err := tr.Next()
		if err != nil {
			if err == io.EOF {
				break
			} else {
				return err
			}
		}

		filename := path.Join(dst, hdr.Name)
		if hdr.FileInfo().IsDir() {
			fmt.Println(filename)
			err := os.MkdirAll(filename, 0755)
			if err != nil {
				return err
			}
			continue
		}

		file, err := os.Create(filename)
		if err != nil {
			return err
		}
		io.Copy(file, tr)
	}
	return nil
}
