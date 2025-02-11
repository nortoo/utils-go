package sms

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"
)

type Client struct {
	c *dysmsapi.Client
}

func NewClient(appid, secretKey, region string) (*Client, error) {
	c, err := dysmsapi.NewClientWithAccessKey(region, appid, secretKey)
	if err != nil {
		return nil, err
	}

	return &Client{c: c}, nil
}
