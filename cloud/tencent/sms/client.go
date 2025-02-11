package sms

import (
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	sms "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/sms/v20190711"
)

type Client struct {
	c        *sms.Client
	smsSdkId string
}

func NewClient(appid, secretKey, region, sdkId string) (*Client, error) {
	credential := common.NewCredential(appid, secretKey)
	cpf := profile.NewClientProfile()
	c, err := sms.NewClient(credential, region, cpf)
	if err != nil {
		return nil, err
	}

	return &Client{
		c:        c,
		smsSdkId: sdkId,
	}, nil
}
