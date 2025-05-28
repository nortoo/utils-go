package image

import (
	"fmt"
	"testing"

	"github.com/nortoo/utils-go/constant"
)

var samples = map[string]string{
	//"testdata/album/sample.webp": "550x368",
	"testdata/album/sample.bmp":  "640x480",
	"testdata/album/sample.tiff": "640x426",
	"testdata/album/sample.jpg":  "2800x1867",
	"testdata/album/sample.png":  "2896x4204",
}

func TestIsSupportedImage(t *testing.T) {
	for i := range samples {
		if !IsSupportedImage(i) {
			t.Logf("[%s] is not supported.", i)
		}
	}
}
func TestDimensions(t *testing.T) {
	for i, d := range samples {
		w, h, err := Dimensions(i)
		if err != nil {
			t.Fatalf("[%s]: cannot get the image info: %s", i, err)
		}
		if fmt.Sprintf("%dx%d", w, h) != d {
			t.Fatalf("the dimension of the file [%s] is incorrect", i)
		}
	}
}

func TestResize(t *testing.T) {
	for i := range samples {
		err := Resize(i, "testdata/output", 128, 0, 80)
		if err != nil {
			t.Fatalf("[%s]: cannot be resized: %s", i, err)
		}
	}
}

func TestResizeByWalkingIntoDir(t *testing.T) {
	err := ResizeByWalkingIntoDir("testdata/album", "testdata/output", 128, 0, 80)
	if err != nil {
		t.Fatal(err)
	}
}

func TestConvert(t *testing.T) {
	var targets = map[string][]constant.MediaType{
		"testdata/album/sample.bmp": {
			constant.JPG,
			constant.JPEG,
			constant.PNG,
			constant.TIFF,
			constant.TIF,
		},
		"testdata/album/sample.jpg": {
			constant.BMP,
			constant.JPEG,
			constant.PNG,
			constant.TIFF,
			constant.TIF,
		},
		"testdata/album/sample.png": {
			constant.BMP,
			constant.JPEG,
			constant.TIFF,
			constant.TIF,
		},
		"testdata/album/sample.tiff": {
			constant.BMP,
			constant.JPEG,
			constant.PNG,
			constant.TIF,
		},
	}

	var n int
	for s, ts := range targets {
		for _, target := range ts {
			err := Convert(s, fmt.Sprintf("testdata/output/%d", n), target)
			if err != nil {
				t.Fatalf("failed to convert [%s] to %s: %v", s, target, err)
			}
		}
		n++
	}
}
