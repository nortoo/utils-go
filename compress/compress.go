package compress

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/nortoo/utils-go/shell"
	"path"
	"strconv"
	"strings"
)

func GetUncompressedFileSize(file string) (int64, error) {
	out := new(bytes.Buffer)
	var err error
	ext := path.Ext(file)
	switch ext {
	case ".zip":
		args := fmt.Sprintf("unzip -l %s | awk 'END{print $1}'", file)
		out, err = shell.Cmd("sh", "-c", args)
	case ".gz":
		args := fmt.Sprintf("gzip -l %s | awk 'END{print $1}'", file)
		out, err = shell.Cmd("sh", "-c", args)
	case ".tar":
		args := fmt.Sprintf("ls -al %s | awk '{print $5}'", file)
		out, err = shell.Cmd("sh", "-c", args)
	default:
		return 0, errors.New("unsupported file")
	}

	if err != nil {
		return 0, err
	}

	return strconv.ParseInt(strings.TrimSuffix(out.String(), "\n"), 10, 64)
}
