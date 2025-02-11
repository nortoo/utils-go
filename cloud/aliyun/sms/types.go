package sms

import "github.com/pkg/errors"

// SMS is the essential parameter of the function Send.
type SMS struct {
	Mobile        string
	SignatureName string
	TemplateCode  string
	Content       map[string]interface{}
}

func (s *SMS) check() error {
	if s.Mobile == "" {
		return errors.New("parameter [Mobile] cannot be empty.")
	}
	if s.SignatureName == "" {
		return errors.New("parameter [SignatureName] cannot be empty.")
	}
	if s.TemplateCode == "" {
		return errors.New("parameter [TemplateCode] cannot be empty.")
	}
	if s.Content == nil {
		return errors.New("parameter [Content] cannot be empty.")
	}
	return nil
}
