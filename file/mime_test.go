package file

import (
	"os"
	"testing"
)

var samples = map[string]string{
	"testdata/davechild_regular-expressions.pdf": "application/pdf",
}

func TestGetMIMEByFilename(t *testing.T) {
	for s, mt := range samples {
		if GetMIMEByFilename(s) != mt {
			t.Fatalf("mime type is incorrect: %s", s)
		}
	}
}

func TestGetMIMEByReadingFile(t *testing.T) {
	for s, mt := range samples {
		f, err := os.Open(s)
		if err != nil {
			t.Fatal(err)
		}
		if GetMIMEByReadingFile(*f) != mt {
			t.Fatalf("mime type is incorrect: %s", s)
			f.Close()
		}
		f.Close()
	}
}
