package ses

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"
)

type Client struct {
	sesCli *ses.SES
}

func New(session *session.Session) *Client {
	return &Client{sesCli: ses.New(session)}
}
