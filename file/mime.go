package file

import (
	"io"
	"net/http"
	"os"
	"path"
	"strings"

	. "github.com/nortoo/utils-go/constant"
)

// GetMIMEByFilename returns the corresponding MIME type by the extension of the given filename.
func GetMIMEByFilename(filename string) string {
	switch strings.ToLower(path.Ext(filename)) {
	case JPG.String(), JPEG.String():
		return "image/jpeg"
	case PNG.String():
		return "image/png"
	case GIF.String():
		return "image/gif"
	case Webp.String():
		return "image/webp"
	case BMP.String():
		return "image/bmp"
	case MP4.String():
		return "video/mp4"
	case AVI.String():
		return "video/x-msvideo"
	case Webm.String():
		return "video/webm"
	case PDF.String():
		return "application/pdf"
	case Doc.String():
		return "application/msword"
	case Docx.String():
		return "application/vnd.openxmlformats-officedocument.wordprocessingml.document"
	case Xls.String():
		return "application/vnd.ms-excel"
	case Xlsx.String():
		return "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet"
	default:
		return "application/octet-stream"
	}
}

// GetMIMEByReadingFile accepts an opened file, and returns the mime type of the file by reading the content.
// Notice:
//
//	This function does not close the file automatically, so you SHOULD close the file after invoking this function.
func GetMIMEByReadingFile(f os.File) string {
	buffer := make([]byte, 512)
	_, err := f.Read(buffer)
	if err != nil && err != io.EOF {
		return "application/octet-stream"
	}
	defer f.Seek(0, 0)

	return http.DetectContentType(buffer)
}
