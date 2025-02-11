package file

import (
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"
)

// Exists returns whether the path exists
func Exists(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}

// IsValidPath returns whether the filepath is valid
func IsValidPath(filepath string) bool {
	reg := regexp.MustCompile(`\.{2,}`)
	pl := strings.Split(filepath, "/")
	for _, p := range pl {
		if reg.MatchString(p) {
			return false
		}
	}
	return true
}

// IsDir returns whether the path is a directory.
func IsDir(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return s.IsDir()
}

// Size returns the disk usage(bytes) of the specific file or directory.
func Size(path string) (int64, error) {
	var size int64
	err := filepath.Walk(path, func(_ string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			size += info.Size()
		}
		return nil
	})
	return size, err
}

// DateStylePathGenerator generates the date-style directory structure.
// e.g. 2006/01/02
func DateStylePathGenerator() string {
	return time.Now().Format("2006/01/02")
}
