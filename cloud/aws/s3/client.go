package s3

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

type Client struct {
	s3         *s3.S3
	uploader   *s3manager.Uploader
	downloader *s3manager.Downloader
}

// NewS3 returns a s3 service collection including S3, Uploader, and Downloader.
func NewS3(s *session.Session) *Client {
	return &Client{
		s3:         s3.New(s),
		uploader:   s3manager.NewUploader(s),
		downloader: s3manager.NewDownloader(s),
	}
}

// S3 returns S3 client.
func (c *Client) S3() *s3.S3 {
	return c.s3
}

// Uploader returns Uploader manager.
func (c *Client) Uploader() *s3manager.Uploader {
	return c.uploader
}

// Downloader returns Downloader manager.
func (c *Client) Downloader() *s3manager.Downloader {
	return c.downloader
}
