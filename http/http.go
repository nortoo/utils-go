package http

import (
	"bytes"
	"fmt"
	"io"
	"math"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/cavaliergopher/grab/v3"
	"github.com/pkg/errors"
)

// Downloader downloads a file to a specific path from an url.
// Deprecated
func Downloader(url, output string, perm os.FileMode, retry uint) error {
	baseDir := filepath.Base(output)
	err := os.MkdirAll(baseDir, perm)
	if err != nil {
		return err
	}

	// Create the file
	out, err := os.Create(output)
	if err != nil {
		return err
	}
	defer out.Close()

	// Get the data
	var resp *http.Response
	for i := 1; i <= int(retry); i++ {
		resp, err = http.Get(url)
		if err != nil {
			time.Sleep(time.Duration(math.Pow(2, float64(i))) * time.Second)
			continue
		} else {
			break
		}
	}
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}

func doURLDownloader(url, output string) error {
	client := grab.NewClient()
	req, _ := grab.NewRequest(output, url)
	fmt.Printf("[URLDownloader] Info: downloading %v...\n", req.URL())
	resp := client.Do(req)
	if resp.HTTPResponse == nil {
		fmt.Printf("[URLDownloader] Error: http return a nil response\n")
		return errors.New("[URLDownloader] Error: http return a nil response")
	}
	fmt.Printf("[URLDownloader] Info: url response status %v\n", resp.HTTPResponse.Status)

	t := time.NewTicker(500 * time.Millisecond)
	defer t.Stop()

Loop:
	for {
		select {
		case <-t.C:
			fmt.Printf("[URLDownloader] Info: transferred %v / %v bytes (%.2f%%)\n",
				resp.BytesComplete(),
				resp.Size(),
				100*resp.Progress())

		case <-resp.Done:
			// download is complete
			break Loop
		}
	}

	// check for errors
	if err := resp.Err(); err != nil {
		return err
	}

	fmt.Printf("[URLDownloader] Info: file saved to %v \n", resp.Filename)
	return nil
}

// URLDownloader downloads a file to a specific path from an url.
func URLDownloader(url, output string, retry uint) (err error) {
	for i := 1; i <= int(retry); i++ {
		err = doURLDownloader(url, output)
		if err != nil {
			time.Sleep(time.Duration(math.Pow(2, float64(i))) * time.Second)
			fmt.Printf("[URLDownloader] Info: download interrupted, retrying... \n")
			continue
		} else {
			break
		}
	}
	return err
}

// RequestWithRetry will resend an HTTP request when a failure occurs until the maximum number of retries is exceeded.
func RequestWithRetry(client *http.Client, req *http.Request, retry int) (resp *http.Response, err error) {
	reqTimes := 1
	contents := make([]byte, 0)
	if req.Body != nil {
		contents, _ = io.ReadAll(req.Body)
	}

	for {
		req.Body = io.NopCloser(bytes.NewReader(contents))
		resp, err = client.Do(req)
		if err != nil {
			if reqTimes > retry {
				return
			}
			reqTimes++
			time.Sleep(time.Duration(2*reqTimes) * time.Millisecond)
			continue
		}
		return
	}
}
