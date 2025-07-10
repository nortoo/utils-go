package ses

import (
	"fmt"
	"os"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	_aws "github.com/nortoo/utils-go/cloud/aws"
)

var c *Client

func TestMain(m *testing.M) {
	var err error

	var (
		accessID  = "<YOUR-SES-ACCESS-ID>"
		secretKey = "<YOUR-SES-SECRET-KEY>"
		region    = "<YOUR-SES-REGION>"
	)

	s, err := _aws.NewSession(&aws.Config{
		Credentials: credentials.NewStaticCredentials(accessID, secretKey, ""),
		Region:      aws.String(region),
	})
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	c = New(s)
	m.Run()
}

func TestClient_Send(t *testing.T) {
	from := "<THE-EMAIL-ADDRESS-OF-THE-SENDER>"
	to := "<THE-TARGET-EMAIL-ADDRESS>"

	content := &Content{
		Subject: "This is a test email.",
		HTML:    "This email is to test if the ses is working.",
		Text:    "",
	}

	err := c.Send(from, []string{to}, content)
	if err != nil {
		t.Fatalf("Failed to send email: %v", err)
	}
}
