package sms

import (
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	_sms "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/sms/v20190711"
)

func (c *Client) Send(sms *SMS) (*_sms.SendSmsResponse, error) {
	if err := sms.check(); err != nil {
		return nil, err
	}

	req := _sms.NewSendSmsRequest()
	req.SmsSdkAppid = common.StringPtr(c.smsSdkId)
	req.Sign = common.StringPtr(sms.SignatureName)
	//req.SenderId = common.StringPtr(senderId)
	req.TemplateID = common.StringPtr(sms.TemplateCode)
	req.TemplateParamSet = common.StringPtrs(sms.Content)
	req.PhoneNumberSet = common.StringPtrs([]string{sms.Mobile})

	return c.c.SendSms(req)
}
