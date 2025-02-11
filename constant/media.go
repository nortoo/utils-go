package constant

type MediaType string

func (t MediaType) String() string {
	return string(t)
}

// common image extensions
const (
	PNG  MediaType = ".png"
	JPG  MediaType = ".jpg"
	JPEG MediaType = ".jpeg"
	GIF  MediaType = ".gif"
	BMP  MediaType = ".bmp"
	TIFF MediaType = ".tiff"
	TIF  MediaType = ".tif"
	Webp MediaType = ".webp"
)

// common video extensions
const (
	MP4  MediaType = ".mp4"
	AVI  MediaType = ".avi"
	Webm MediaType = ".webm"
)

// common documents extensions
const (
	PDF  MediaType = ".pdf"
	Doc  MediaType = ".doc"
	Docx MediaType = ".docx"
	Xls  MediaType = ".xls"
	Xlsx MediaType = ".xlsx"
)
