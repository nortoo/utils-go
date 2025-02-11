package sms

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"
)

// Send a short message to a mobile phone.
func (c *Client) Send(sms *SMS) (*dysmsapi.SendSmsResponse, error) {
	if err := sms.check(); err != nil {
		return nil, err
	}
	smsContent, err := json.Marshal(sms.Content)
	if err != nil {
		return nil, err
	}

	req := dysmsapi.CreateSendSmsRequest()
	req.Scheme = "https"
	req.PhoneNumbers = sms.Mobile
	req.SignName = sms.SignatureName
	req.TemplateCode = sms.TemplateCode
	req.TemplateParam = string(smsContent)

	return c.c.SendSms(req)
}
