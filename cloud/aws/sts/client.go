package client

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sts"
)

type Client struct {
	sts *sts.STS
}

func NewSTS(s *session.Session) *Client {
	return &Client{
		sts: sts.New(s),
	}
}

func (c *Client) GetSTS() *sts.STS {
	return c.sts
}
