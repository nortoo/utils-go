package sms

import (
	"fmt"
	"testing"
)

func TestSendSMS(t *testing.T) {
	c, err := NewClient("", "", "", "")
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
	fmt.Printf("response is %#v\n", resp)
}
