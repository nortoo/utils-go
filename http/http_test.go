package http

import (
	"net/http"
	"testing"
)

func TestHTTPDownloader(t *testing.T) {
	err := Downloader("https://cdimage.debian.org/debian-cd/current/amd64/iso-cd/debian-12.9.0-amd64-netinst.iso",
		"/tmp/debian-12.9.0-amd64-netinst.iso",
		0755, 3)
	if err != nil {
		t.Fatal(err)
	}
}

func TestURLDownloader(t *testing.T) {
	err := URLDownloader("https://cdimage.debian.org/debian-cd/current/amd64/iso-cd/debian-12.9.0-amd64-netinst.iso",
		"/tmp/debian-12.9.0-amd64-netinst.iso",
		3)
	if err != nil {
		t.Fatal(err)
	}
}

func TestRequest(t *testing.T) {
	url := "https://echo.free.beeceptor.com/"
	req, _ := http.NewRequest("GET", url, nil)
	resp, err := RequestWithRetry(&http.Client{}, req, 3)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("echo -> %s", resp.Status)
}
