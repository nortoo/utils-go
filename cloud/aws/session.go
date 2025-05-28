package aws

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
)

// NewSession creates a new aws session.
func NewSession(config *aws.Config) (*session.Session, error) {
	s, err := session.NewSession(config)
	if err != nil {
		return nil, err
	}
	return session.Must(s, nil), nil
}
