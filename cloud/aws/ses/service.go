package ses

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ses"
)

type Service interface {
	Send() error
}

func (c *Client) Send(from string, to []string, content *Content) error {
	receivers := make([]*string, len(to))
	for i, r := range to {
		receivers[i] = aws.String(r)
	}

	_, err := c.sesCli.SendEmail(&ses.SendEmailInput{
		ConfigurationSetName: nil,
		Destination: &ses.Destination{
			ToAddresses: receivers,
		},
		Message: &ses.Message{
			Body: &ses.Body{
				Html: &ses.Content{
					Charset: aws.String(ses.SNSActionEncodingUtf8),
					Data:    aws.String(content.HTML),
				},
				Text: &ses.Content{
					Charset: aws.String(ses.SNSActionEncodingUtf8),
					Data:    aws.String(content.Text),
				},
			},
			Subject: &ses.Content{
				Charset: aws.String(ses.SNSActionEncodingUtf8),
				Data:    aws.String(content.Subject),
			},
		},
		Source: aws.String(from),
	})

	return err
}
