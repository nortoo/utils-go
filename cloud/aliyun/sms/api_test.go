package sms

import (
	"testing"
)

func TestSendSms(t *testing.T) {
	c, err := NewClient("", "", "")
	if err != nil {
		t.Fatal(err)
	}

	resp, err := c.Send(&SMS{
		Mobile:        "",
		SignatureName: "",
		TemplateCode:  "",
		Content:       nil,
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("response is %#v\n", resp)
}
