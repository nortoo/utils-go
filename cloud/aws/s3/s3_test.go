package s3

import (
	"fmt"
	"os"
	"testing"
	"time"

	_aws "github.com/nortoo/utils-go/cloud/aws"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
)

var c *Client

func TestMain(m *testing.M) {
	var err error

	var (
		accessID        = ""
		secretKey       = ""
		endpoint        = ""
		region          = ""
		disableSSL      = true
		enablePathStyle = true
	)

	s, err := _aws.NewSession(&aws.Config{
		Credentials:      credentials.NewStaticCredentials(accessID, secretKey, ""),
		Endpoint:         aws.String(endpoint),
		Region:           aws.String(region),
		DisableSSL:       aws.Bool(disableSSL),
		S3ForcePathStyle: aws.Bool(enablePathStyle),
	})
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	c = NewS3(s)
	m.Run()
}

func TestClient_Upload(t *testing.T) {
	fp := ""
	bucket := ""
	key := ""
	err := c.Upload(fp, bucket, key, PublicRead, 3)
	if err != nil {
		t.Fatal(err)
	}
}

func TestClient_Download(t *testing.T) {
	bucket := ""
	key := ""
	fp, err := c.Download("", bucket, key, 3)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(fp)
}

func TestClient_UploadByMultipart(t *testing.T) {
	fp := ""
	bucket := ""
	key := ""
	err := c.UploadByMultipart(fp, bucket, key, PublicRead, 3)
	if err != nil {
		t.Fatal(err)
	}
}

func TestClient_GetObjectUrl(t *testing.T) {
	bucket := ""
	key := ""

	url, err := c.GetObjectUrl(bucket, key, 60*time.Minute)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(url)
}

func TestClient_DeleteObject(t *testing.T) {
	bucket := ""
	key := ""

	err := c.DeleteObject(bucket, key)
	if err != nil {
		t.Fatal(err)
	}
}

func TestClient_IsObjectExist(t *testing.T) {
	bucket := ""
	key := ""

	result := c.IsObjectExist(bucket, key)
	t.Log("result: ", result)
}

func TestClient_HeadObject(t *testing.T) {
	bucket := ""
	key := ""

	result, err := c.HeadObject(bucket, key)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("result: %v\n", result)
}
