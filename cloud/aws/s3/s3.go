package s3

import (
	"fmt"
	"math"
	"os"
	"path"
	"path/filepath"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"

	_file "github.com/nortoo/utils-go/file"
	_http "github.com/nortoo/utils-go/http"
	uuid "github.com/satori/go.uuid"
)

// s3ObjectACL
// AWS S3 ACL
// refer: https://docs.aws.amazon.com/AmazonS3/latest/userguide/acl-overview.html
type s3ObjectACL string

const (
	Private         s3ObjectACL = "private"
	PublicRead      s3ObjectACL = "public-read"
	PublicReadWrite s3ObjectACL = "public-read-write"
	BucketOwnerRead s3ObjectACL = "bucket-owner-read"
)

// Upload one file to s3
func (c *Client) Upload(pt, bucket, key string, acl s3ObjectACL, retry uint) error {
	f, err := os.Open(pt)
	if err != nil {
		return err
	}
	defer f.Close()
	contentType := _file.GetMIMEByReadingFile(*f)

	fi, err := f.Stat()
	if err != nil {
		return err
	}

	var n uint
UPLOAD:
	n++
	_, err = c.S3().PutObject(&s3.PutObjectInput{
		Bucket:             aws.String(bucket),
		ACL:                aws.String(string(acl)),
		Key:                aws.String(key),
		Body:               f,
		ContentLength:      aws.Int64(fi.Size()),
		ContentType:        aws.String(contentType),
		ContentDisposition: aws.String(_http.BuildContentDisposition(filepath.Base(key))),
	})
	if err != nil {
		if n > retry {
			return err
		}
		_, _ = f.Seek(0, 0)
		time.Sleep(time.Duration(math.Pow(2, float64(n))) * time.Second)
		goto UPLOAD
	}
	return err
}

// Download one file from s3 to the local, and returns the local file path.
func (c *Client) Download(dstDir, bucket, key string, retry uint) (fp string, err error) {
	err = os.MkdirAll(dstDir, 0755)
	if err != nil {
		return "", err
	}

	filename := fmt.Sprintf("%s%s", uuid.NewV4().String(), path.Ext(key))
	fp = path.Join(dstDir, filename)
	file, err := os.Create(fp)
	if err != nil {
		return "", err
	}
	defer file.Close()

	// S3 download retry
	var n uint
DOWNLOAD:
	n++
	_, err = c.Downloader().Download(file, &s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	})
	if err != nil {
		if n > retry {
			return
		}
		time.Sleep(time.Duration(math.Pow(2, float64(n))) * time.Second)
		goto DOWNLOAD
	}
	return
}

// GetObjectUrl returns the object's url.
func (c *Client) GetObjectUrl(bucket, key string, expire time.Duration) (url string, err error) {
	req, _ := c.S3().GetObjectRequest(&s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	})
	url, _, err = req.PresignRequest(expire)
	return
}

// UploadByMultipart to upload big file using multipart mode.
func (c *Client) UploadByMultipart(pt, bucket, key string, acl s3ObjectACL, retry uint) error {
	f, err := os.Open(pt)
	if err != nil {
		return nil
	}
	defer f.Close()
	contentType := _file.GetMIMEByReadingFile(*f)

	var n uint
UPLOAD:
	n++
	_, err = c.Uploader().Upload(&s3manager.UploadInput{
		Bucket:             aws.String(bucket),
		Key:                aws.String(key),
		ACL:                aws.String(string(acl)),
		Body:               f,
		ContentType:        aws.String(contentType),
		ContentDisposition: aws.String(_http.BuildContentDisposition(filepath.Base(key))),
	})
	if err != nil {
		if n > retry {
			return err
		}
		_, _ = f.Seek(0, 0)
		time.Sleep(time.Duration(math.Pow(2, float64(n))) * time.Second)
		goto UPLOAD
	}
	return err
}

// DeleteObject deletes the specific object from s3.
func (c *Client) DeleteObject(bucket, key string) error {
	_, err := c.S3().DeleteObject(&s3.DeleteObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	})
	return err
}

// IsObjectExist returns whether the specific object exists.
func (c *Client) IsObjectExist(bucket, key string) bool {
	_, err := c.S3().HeadObject(&s3.HeadObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	})
	return err == nil
}

// HeadObject returns the specific object head info, including file size, file type, etc.
func (c *Client) HeadObject(bucket, key string) (*s3.HeadObjectOutput, error) {
	return c.S3().HeadObject(&s3.HeadObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	})
}
